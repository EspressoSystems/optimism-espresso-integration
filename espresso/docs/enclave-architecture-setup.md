# EIF Architecture: op-batcher in AWS Nitro Enclave

This document explains how the op-batcher runs inside an AWS Nitro Enclave — what
every image contains, how the pieces are built, and exactly what code runs at
runtime and in what order.

---

## Key Files

| File | Runs where | Purpose |
|------|-----------|---------|
| [`op-batcher/enclave-entrypoint.bash`](../../op-batcher/enclave-entrypoint.bash) | Inside Nitro enclave | Receives args, sets up socat proxies, starts op-batcher |
| `run-eif.sh` (`aws-nitro`) | Outer container (EC2 host) | Lives in the `aws-nitro` repo alongside `build-eif.yml`.  Cleans stale enclaves, polls `describe-enclaves`, performs 8338 readiness handshake, delivers args via nc, delegates shutdown to enclaver-run. |
| [`espresso/docker/op-batcher-tee/run-enclave.sh`](../docker/op-batcher-tee/run-enclave.sh) | Local dev only | Full build + register + run in one command |
| `op-batcher/enclave-tools/` | CI runner | Builds EIF images, registers PCR0 on-chain |
| `build-eif.yml` (`awd-nitro`) | `aws-nitro` repo CI workflow | Builds and pushes `op-batcher-eif:TAG` |

---

## The Two Images

Everything revolves around two Docker images.

### `op-batcher-enclave-app:TAG` — the inner image

Built by the espresso-integration CI.  This is the *application* that will run
**inside** the Nitro enclave.

Contents (defined in [`espresso/docker/op-stack/Dockerfile`](../docker/op-stack/Dockerfile)):

| Path | What it is |
|------|-----------|
| `/usr/local/bin/op-batcher` | The batcher binary |
| `/app/entrypoint.sh` → [`op-batcher/enclave-entrypoint.bash`](../../op-batcher/enclave-entrypoint.bash) | Startup script that runs inside the enclave |
| `socat`, `nc`, `trurl` | Networking tools used by the entrypoint for proxy setup |

**This image determines PCR0.**  Every byte of its filesystem is hashed into
PCR0 when the EIF is built.  If the app changes, PCR0 changes.

---

### `op-batcher-eif:TAG` — the outer runner image

Built by the `aws-nitro` repo `build-eif` workflow.  This is the only image
deployed to ECS/Terraform.  It is layered in three stages:

```
┌─────────────────────────────────────────────┐
│  /run-eif.sh            (ENTRYPOINT)         │  ← stage 3: run-eif.sh
│  /bin/busybox + symlinks (sh, nc, timeout…)  │     added by build-eif.yml
├─────────────────────────────────────────────┤
│  /enclave/application.eif   (the EIF)        │  ← stage 2: enclaver output
│  /usr/local/bin/enclaver-run                 │     FROM scratch + enclaver-run
│  /bin/nitro-cli                              │     + nitro-cli + glibc
├─────────────────────────────────────────────┤
│  (scratch — empty)                           │  ← stage 1: enclaver base
└─────────────────────────────────────────────┘
```

The enclaver base is literally `FROM scratch` (see
[runtimebase.dockerfile in enclaver v0.5.0](https://github.com/enclaver-io/enclaver/blob/v0.5.0/build/dockerfiles/runtimebase.dockerfile)).
It has no shell, no utilities — only `enclaver-run`, `nitro-cli`, and the
shared libraries `nitro-cli` needs.  That is why `busybox` and `run-eif.sh`
must be added on top in stage 3.

> **Note**: adding `run-eif.sh` and `busybox` as outer layers does **not**
> affect PCR0–2.  Those measurements are sealed inside the EIF at stage 2 and
> never change once the EIF is built.

---

## The Build Pipeline

Two independent CI runs produce `op-batcher-eif:TAG`:

```
espresso-integration CI
  │
  ├──▶ op-batcher-enclave-app:TAG   (app → goes into EIF, determines PCR0)
  └──▶ op-batcher-tee:TAG           (source of enclave-tools)
              │
              │   [infra repo: build-eif.yml + run-eif.sh]
              │
              ▼
        op-batcher-eif:TAG          (single self-contained runner)
```

### Step 1 — espresso-integration CI

Produces two images from the same commit:

- **`op-batcher-enclave-app`**: the inner app (what runs inside the enclave).
- **`op-batcher-tee`**: a build-helper image containing `enclave-tools` (a
  CLI for building EIFs).

### Step 2 — infra repo `build-eif-op` workflow

Runs with a `workflow_dispatch` trigger, takes a source tag as input.

1. **Pull and pin** both source images to their `sha256` digest — re-running
   the workflow will always use the exact same bytes.
2. **Extract `enclave-tools`** from `op-batcher-tee` so the CI runner can call
   it.
3. **`enclave-tools build-eif`** — calls enclaver to wrap
   `op-batcher-enclave-app` into an EIF Docker image locally.  Captures PCR0,
   PCR1, PCR2.
4. **Layer `busybox` + `run-eif.sh`** on top of the enclaver output, add OCI
   labels (including PCR values), set `ENTRYPOINT ["/run-eif.sh"]`.
   `run-eif.sh` is taken from the `aws-nitro` repo checkout (not from the app CI
   images) — edit it there and re-run `build-eif.yml` with the same app tag
   to update the runner without any espresso-integration CI rebuild.
5. **Push** the single `op-batcher-eif:TAG` to the registry.

> **Why two steps?**  The devops "team" controls when a new EIF is promoted to
> production, independently of the application CI. We use this approach to permit
> both local testing with the enclaver tools and use in production.

---

## Enclave Resources

`enclaver.yaml` is embedded inside the EIF, so CPU and memory are sealed into PCR0 at build
time.  Defaults: **2 vCPUs**, **4096 MiB** (overridable via `--cpu-count` / `--memory-mb`
passed to `enclave-tools build-eif`).  Changing either value requires rebuilding the EIF and
re-registering the new PCR0 on-chain.

---

## Runtime: What Happens When ECS Starts the Container

```
EC2 Host
└─ Docker container: op-batcher-eif:TAG
    │
    ├─ PID 1: /bin/sh /run-eif.sh
    │   │  1. validates env vars; terminates stale enclaves; asserts TCP:8337 free
    │   │  2. starts enclaver-run in background
    │   │  3. polls describe-enclaves until enclave ID appears (≤120 s)
    │   │  4. polls TCP:8338 until "READY" received (readiness handshake)
    │   │  5. sends NUL-separated batcher args to TCP:8337
    │   └─ waits on enclaver-run PID (forwards SIGTERM on shutdown)
    │
    └─ enclaver-run (background)
        │  reads /enclave/enclaver.yaml
        │  calls: nitro-cli run-enclave --eif-path /enclave/application.eif
        │  bridges: TCP:8337  ←──vsock──▶  port 8337 inside enclave
        │
        └─ [Nitro Enclave — isolated from host]
            └─ enclave-entrypoint.bash
                │  1. nc listener on port 8337 (background, before anything else)
                │  2. checks Odyn egress proxy is up
                │  3. sends "READY" on port 8338 (background, readiness handshake)
                │  4. waits for nc:8337 to finish (blocking up to 60 s)
                │  5. rewrites internal URLs via socat → Odyn
                └─ exec op-batcher [all assembled args]
```

### Detailed step-by-step

**1. `run-eif.sh` starts (PID 1)**
Validates all required environment variables (`L1_RPC_URL`, `OP_BATCHER_PRIVATE_KEY`,
etc.) and prints the configuration.  No sensitive values are logged.  Calls
`nitro-cli describe-enclaves` and terminates any enclave already running on this
host.  Then verifies TCP:8337 is not bound — a cheap pre-flight check that fails fast
with a clear error rather than letting `enclaver-run` start and then fail
silently with `EADDRINUSE` on vsock:17002.

**2. `enclaver-run` starts**
Reads `/enclave/enclaver.yaml` (which specifies CPU count, memory, and that
port 8337 is an ingress port).  First binds vsock:17002 for the egress proxy, then
calls `nitro-cli run-enclave` to boot the EIF inside the Nitro hypervisor.  After
the enclave registers, opens the TCP↔vsock bridge: anything connecting to
`127.0.0.1:8337` on the host is proxied into vsock port 8337 inside the enclave.

**3. Enclave readiness: `describe-enclaves` polling**
`run-eif.sh` polls `nitro-cli describe-enclaves` once per second (up to 120 s)
until the enclave ID appears.  This is the definitive signal that the Nitro
hypervisor has accepted the EIF and the enclave is running.

> **Why not poll TCP:8337?**  TCP:8337 is the host-side vsock bridge, which
> `enclaver-run` opens *after* the enclave registers — but *before* the enclave's
> internal `nc` listener is ready.  Connecting to TCP:8337 at this point would
> reach the ingress proxy, which would immediately try to forward to the enclave
> and get `ECONNREFUSED` — consuming the one-shot nc listener window or causing a
> failure.

**4. Readiness handshake on port 8338**
After capturing the enclave ID, `run-eif.sh` polls TCP:8338 in a retry loop
(up to 30 s).  Inside the enclave, `enclave-entrypoint.bash` sends the string
`"READY"` on port 8338 immediately after the Odyn check succeeds.  When
`run-eif.sh` receives that signal it knows two things: the nc:8337 arg listener
is open (it started before the Odyn check), and Odyn is verified.  This is a
deterministic handshake that replaces a fixed sleep and holds regardless of host
load.  If the signal is not received within 30 s, `run-eif.sh` logs a warning
and proceeds anyway so a stale 8338 listener never blocks arg delivery.

**5. Arg delivery (the NUL-separated protocol)**
`run-eif.sh` pipes every batcher CLI flag through `nc 127.0.0.1 8337` as a
stream of NUL-terminated strings, followed by a second NUL to signal
end-of-stream:

```
run-eif.sh sends:                enclave-entrypoint.bash reads:
  "--l1-eth-rpc=http://…\0"  →   arg 1
  "--l2-eth-rpc=http://…\0"  →   arg 2
  "--private-key=0xabc…\0"   →   arg 3  (logged as [REDACTED])
  …
  "\0"  (second NUL)          →   empty read → break loop
```

The protocol is implemented with `printf '%s\0'` on the sender side and
`IFS= read -r -d '' arg` in a loop on the receiver side.  The private key is
never logged on the outer container; `enclave-entrypoint.bash` redacts it in the
argument dump printed before `exec op-batcher`.

**6. Inside the enclave: nc listener is already open**
`enclave-entrypoint.bash` starts `nc -l -p 8337 -w 60` as a **background
process at its very first line**, before the Odyn check.  This eliminates the
race: by the time `run-eif.sh` sends args (step 5), the listener has been open
for several seconds.  An `EXIT` trap ensures the nc process and its tempfile are
cleaned up on any exit path, including early failures.

**7. Inside the enclave: proxy setup**
After `wait`ing for nc to finish receiving args, `enclave-entrypoint.bash`
inspects each URL argument.  URLs pointing to `localhost` or `127.0.0.1`
(internal services on the EC2 host) cannot be reached directly from the
enclave — for each one a `socat` process is started to forward connections
through **Odyn** (enclaver's egress proxy, which tunnels traffic over vsock).
External URLs (Espresso sequencer, L1 RPC) pass through the HTTPS proxy
directly, preserving the correct `Host` header and SNI.

**8. `op-batcher` starts**
After all proxies are ready, `exec op-batcher [all args]` replaces the
entrypoint process.  The batcher now runs inside the Nitro enclave with no
ability for the EC2 host to inspect its memory or state.

**9. Steady state and shutdown**
`enclaver-run` stays alive proxying vsock traffic for the lifetime of the
enclave.  `run-eif.sh` waits on its PID.  When ECS sends `SIGTERM` to PID 1:

```
1. SIGTERM → run-eif.sh (PID 1)
2. trap fires enclave_shutdown()
3. kill $ENCLAVER_PID  — sends SIGTERM to enclaver-run
4. enclaver-run's handler fires: calls nitro-cli terminate-enclave,
   which releases vsock:17002, then enclaver-run exits
5. wait $ENCLAVER_PID  — blocks until enclaver-run has fully exited,
   guaranteeing the vsock port is free before the container disappears
6. exit 0
```

`run-eif.sh` delegates enclave termination entirely to enclaver-run rather than
calling `nitro-cli terminate-enclave` directly.  Calling it directly causes a
double-terminate race — `run-eif.sh` terminates the enclave by ID, then
enclaver-run's SIGTERM handler also tries to terminate the same enclave and gets
`E11` (socket error, enclave already gone), producing errors in logs.
The correct pattern is: signal enclaver-run, wait for it — enclaver-run owns the
enclave lifecycle.

## Pre-built EIF vs. Docker-in-Docker

An earlier iteration of this system (`run-enclave.sh`) drove the batcher via
`enclave-tools run` — a wrapper around `docker run --privileged
--device=/dev/nitro_enclaves`.  This spins up a second Docker container whose
`enclaver-run` process boots a **real** Nitro enclave with **real** PCR0
measurements.  The `BatchAuthenticator` attestation works the same in both
approaches; the security model is equivalent.

The reason `run-enclave.sh` is not suitable for production is operational, not
cryptographic.

### Operational problems with Docker-in-Docker

**1. No SIGTERM trap in AWS mode — stale inner containers block the next start**

`run-enclave.sh` only installs a cleanup trap in `DEPLOYMENT_MODE=local`.  In
AWS mode (the default), when ECS sends SIGTERM to the outer task, the outer
process exits, but the inner `batcher-enclaver-*` Docker container keeps
running.  That container's `enclaver-run` holds `vsock:17002`.  When ECS starts
a replacement task on the same instance, `enclaver-run` fails with
`EADDRINUSE` because the vsock port is still occupied.  The only recovery is to
terminate the EC2 instance.

`run-eif.sh` traps SIGTERM → kills `enclaver-run` → `enclaver-run` terminates
the enclave → `vsock:17002` is released before the container exits.

**2. `describe-enclaves` session scoping makes outer-layer cleanup blind**

`nitro-cli describe-enclaves` lists enclaves started by the **current**
`/dev/nitro_enclaves` session.  In Docker-in-Docker the Nitro enclave is started
by the *inner* container's session, so the outer `run-enclave.sh` cannot see or
terminate it via `describe-enclaves`.  If the inner container exits uncleanly the
enclave keeps running and `vsock:17002` stays bound.

With `run-eif.sh` there is only one process layer — `enclaver-run` runs directly
in the outer container and owns the session.  `describe-enclaves` sees the enclave,
and `run-eif.sh` can clean up stale enclaves on startup.

**3. Docker daemon dependency**

`enclave-tools run` requires the Docker socket to be mounted into the outer
container.  `run-eif.sh` only requires the `/dev/nitro_enclaves` device —
no Docker daemon needed at runtime.

**4. EIF pre-built; no build or registration at startup**

`run-enclave.sh` calls `enclave-tools build` (full EIF build) and `enclave-tools
register` (on-chain transaction) on every cold start.  A build takes minutes; a
registration transaction takes a block.  `run-eif.sh` starts from a pre-built EIF
baked into the image — startup is seconds, not minutes.

**5. Complex monitoring loop vs `wait`**

`run-enclave.sh` polls `docker inspect` every 30 seconds in a monitoring loop,
writes JSON status files, and launches a background log-capture sub-process.
`run-eif.sh` simply calls `wait $ENCLAVER_PID` — `enclaver-run` stays alive for
the enclave lifetime, so waiting on it is the correct and minimal signal.

**6. Comma-separated env var arg delivery vs NUL-separated vsock protocol**

`run-enclave.sh` assembles all CLI flags into a single comma-separated
`BATCHER_ARGS` env var (e.g.
`--l1-eth-rpc=http://…,--l2-eth-rpc=http://…,--private-key=0xabc…`).  Inside the
enclave `enclave-entrypoint.bash` runs `eval set -- "$ENCLAVE_BATCHER_ARGS"` to
split them back out.  This has two problems: the private key was fully visible in
the env var (readable from `docker inspect` of the inner container), and commas
inside URL values would silently break argument parsing.

`run-eif.sh` sends args as NUL-separated bytes over the vsock/TCP bridge to port
8337 after the enclave is already running.  The private key is never visible in
any env var or `docker inspect` output.

**7. Deterministic readiness handshake vs `sleep 5`**

`run-enclave.sh` uses an unconditional `sleep 5` before looking for the inner
container.  Under load this races.  `run-eif.sh` waits for the explicit `READY`
signal on port 8338, which the enclave sends only after `nc:8337` is open and
Odyn is verified.

### When Docker-in-Docker is still appropriate

`run-enclave.sh` is the right tool for local development: it builds the EIF from
source, registers the PCR0 against a local devnet, and starts the batcher in one
command.  People can iterate on the binary and entrypoint without a
Nitro-capable EC2 instance.  The tradeoff is explicit — full build each run, no
pre-built artifact.


## Trust Boundary

The diagram below shows what is inside the hardware trust boundary and what is not.

```
┌─────────────────────────────────────────────────────────────────────┐
│  EC2 Instance  (root-accessible; AWS hardware-managed)              │
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │  Docker container  (ECS task)         ← OUTSIDE TEE           │  │
│  │                                                               │  │
│  │  run-eif.sh          enclaver-run                             │  │
│  │  ECS env vars:                                                │  │
│  │    OP_BATCHER_PRIVATE_KEY  ← visible to EC2 root              │  │
│  │    L1/L2/rollup RPC URLs   ← set by operator, not attested    │  │
│  │    ESPRESSO_LIGHT_CLIENT_ADDR                                 │  │
│  └──────────────────────┬──────────────────────────────────────┬─┘  │
│                         │ vsock:8337 (args in)                 │    │
│                         │ vsock:8338 (ready out)               │    │
│                         │ vsock:17002 (egress out)             │    │
│  ═══════════════════════╪══════════════════════════════════════│════│
│      TEE BOUNDARY       ▼                                      ▼    │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │  Nitro Enclave  (hardware-isolated; no debug access)          │  │
│  │                                                               │  │
│  │  enclave-entrypoint.bash                                      │  │
│  │  op-batcher binary                                            │  │
│  │  OP_BATCHER_PRIVATE_KEY  (in enclave memory only after entry) │  │
│  │  cpu_count=2, memory_mb=4096  ← sealed in EIF at build time   │  │
│  │                                                               │  │
│  │  PCR0 = SHA-384(entire EIF)  ← fixed at build time            │  │
│  │  Attestation doc signed by AWS Nitro hardware                 │  │
│  └───────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────┘
```

**What the TEE does not prove**: the correctness of the arguments passed to that
code.  RPC endpoints, the light client address, and origin heights all come from
ECS env vars outside the boundary.

---

## PCR Measurements

| Register | Measures | Changes when… |
|----------|---------|--------------|
| **PCR0** | SHA-384 of the entire EIF (kernel + rootfs + app) | `op-batcher-enclave-app` changes |
| **PCR1** | SHA-384 of the Linux kernel used by enclaver | Kernel version changes |
| **PCR2** | SHA-384 of the boot filesystem | Enclaver version changes |

PCR0 is the most important and only value registered in the espresso TEE contracts as of now but we also expose PCR 1 & 2 for later inclusion

The `build-eif.yml` workflow bakes all three PCR values — and a `keccak256` of
PCR0 (`enclave.hash`) — directly into the runner image as OCI labels:

```
enclave.pcr0=<sha384-hex>
enclave.pcr1=<sha384-hex>
enclave.pcr2=<sha384-hex>
enclave.hash=<keccak256-of-pcr0>   ← value to pass to enclave-tools register
enclave.tag=<source-tag>
enclave.app-image=<digest>
```

This means `docker inspect ghcr.io/…/op-batcher-eif:TAG` gives the PCR0 without
needing to re-run the build.

---

## PCR0 Change Procedure

### What triggers a PCR0 change

| Change | PCR0 affected? |
|--------|---------------|
| `op-batcher/enclave-entrypoint.bash` | Yes |
| `op-batcher` binary | Yes |
| `socat`, `nc`, `trurl` package versions (Alpine) | Yes |
| `TARGET_BASE_IMAGE` (Alpine version) | Yes |
| Ingress/egress port config in `enclaver.go` | Yes (changes embedded enclaver.yaml) |
| `--cpu-count` / `--memory-mb` passed to `build-eif` | Yes — manifest is embedded in EIF at `/etc/enclaver/enclaver.yaml` |
| Enclaver version | Yes (PCR1/PCR2; may affect PCR0) |
| `run-eif.sh` | **No** — outer layer, not in EIF |
| ECS task env vars | **No** — outside the EIF entirely |
| `build-eif.yml` workflow | **No** |
