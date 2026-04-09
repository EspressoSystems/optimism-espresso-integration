# Optimism Espresso Integration

> See also: [deployment configuration](espresso/docs/README_ESPRESSO_DEPLOY_CONFIG.md) · [code sync procedure](espresso/docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md)

## Development environment

### Clone the repository and initialize the submodules

```console
> git clone git@github.com:EspressoSystems/optimism-espresso-integration.git
> git submodule update --init --recursive
```

### Nix shell

* Install nix following the instructions at <https://nixos.org/download/>

* Enter the nix shell of this project:

```console
> nix develop .
```

### Configuring Docker

In order to download the docker images required by this project you may need to authenticate using a PAT.

Create a [Github Personal Access Token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic) following Creating a personal access token (classic).

Provide Docker with the PAT:

```console
> export CR_PAT=<your PAT>
> echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
```

On Linux, run docker as a non-root user (not needed on macOS with Docker Desktop):

```console
> sudo groupadd docker
> sudo usermod -aG docker $USER
```

### Run the tests

Run the Espresso smoke tests:

```console
> just smoke-tests
```

Run the Espresso integration tests. Note, this can take up to 30min.

```console
> just espresso-tests
```

To run all the standard OP stack (w/o Espresso integration) tests (slow):

```console
> just tests
```

To run a subset of the tests above (fast):

```console
> just fast-tests
```

To run the devnet tests:

```console
> just devnet-tests
```

#### Espresso Attestation Verifier
The Espresso Attestation Verifier is utilized to register Attestations for the
Builder. For the E2E Testnet (utilized in local testing) this is provided as
an opt-in configuration.  The live Batcher is expected to run with this
enabled and configured.  If it is not configured and we have an Attestation
a warning log will be issued.

In order to enable the Espresso Attestation Verifier in the local E2E tests
you merely need to include the relevant option in the configuration:

```go
system, _, err := launcher.StartE2eDevnet(ctx, t,
    env.WithEspressoAttestationVerifierService(),
)
```

> **NOTE:** This configuration has default values configured for convenience
> and to make the Attestation Verifier Service launch without any external
> configuration being required. However, the option itself can also take
> options that allow the values to be overridden.
> Additionally, to preserve the previous behavior of the Attestation Verifier
> it also supports configuration via the same Environment Variables that
> were previously required.

These environment variables are set in the [espresso/.env](espresso/.env) file for reference.
However for clarity they are listed here.  These values will only be used
in the testing when they are populated to a non-empty value.

```env
ESPRESSO_ATTESTATION_VERIFIER_PORT=<The port to host the verifier service on>
ESPRESSO_ATTESTATION_VERIFIER_RPC_URL=<The RPC URL to communicate with>
ESPRESSO_ATTESTATION_VERIFIER_SP1_PROVER=<The SP1 Prover mode to operate in>
ESPRESSO_ATTESTATION_VERIFIER_NITRO_VERIFIER_ADDRESS=<The nitro verifier address>
ESPRESSO_ATTESTATION_VERIFIER_SKIP_TIME_VALIDITY_CHECK=<whether or not to enable the validity check>
ESPRESSO_ATTESTATION_VERIFIER_HOST=<The host to listen for. Meant to be the bind address, best if "0.0.0.0" is used>
ESPRESSO_ATTESTATION_VERIFIER_NETWORK_PRIVATE_KEY=<The hex encoded private key to utilize>
ESPRESSO_ATTESTATION_VERIFIER_NETWORK_RPC_URL=<The Network RPC URL to utilize>
ESPRESSO_ATTESTATION_VERIFIER_NETWORK_USE_DOCKER=<Whether or not to use docker for the attestation verifier. "1" or "0">
ESPRESSO_ATTESTATION_VERIFIER_RUST_LOG=<The RUST_LOG level to pass to the service>
ESPRESSO_ATTESTATION_VERIFIER_DOCKER_IMAGE=<The Docker Image to utilize for the attestation verifier service>
```

### Misc commands

In order to run the go linter do:

```console
just golint
```

Generate the bindings for the contracts:

```console
just gen-bindings
```

Stop all devnet containers (docker compose services):

```console
just stop-containers
```

Remove a stuck Espresso dev node container left over from a failed integration test (matches by image ancestor):

```console
just remove-espresso-containers
```

### Guide: setting up an enclave-enabled Nitro EC2 instance

This guide explains how to prepare an enclave-enabled parent EC2 instance.

You can follow the official AWS Enclaves setup guide: <https://docs.aws.amazon.com/enclaves/latest/user/getting-started.html>.

#### Step-by-step instructions

##### 1. Launch the EC2 instance

Use the AWS Management Console or AWS CLI to launch a new EC2 instance.

Make sure to:

* **Enable Enclaves**
  * In the CLI: set the `--enclave-options` flag to `true`
  * In the Console: select `Enabled` under the **Enclave** section

* **Use the following configuration:**
  * **Architecture:** x86_64
  * **AMI:** Amazon Linux 2023
  * **Instance Type:** `m6a.2xlarge`
  * **Volume Size:** 100 GB

##### 2. Connect to the instance

Once the instance is running, connect to it via the AWS Console or CLI.
In practice, you will be provided a `key.pem` file, and you can connect like this:

```console
chmod 400 key.pem
ssh -i "key.pem" ec2-user@<aws_instance_dns>
```

Note that the command above can be found in the AWS Console by selecting the instance and clicking on the button "Connect".

##### 3. Install dependencies

* Nix:

```console
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon
source ~/.bashrc
```

* Git, Docker:

```console
sudo yum update
sudo yum install git
sudo yum install docker
sudo usermod -a -G docker ec2-user
sudo service docker start
sudo chown ec2-user /var/run/docker.sock
```

* Nitro

These commands install, configure, and start the Nitro enclave service:

```console
sudo yum install -y aws-nitro-enclaves-cli-1.4.2
sudo systemctl stop nitro-enclaves-allocator.service || true
echo -e '---\nmemory_mib: 4096\ncpu_count: 2' | sudo tee /etc/nitro_enclaves/allocator.yaml
sudo systemctl start nitro-enclaves-allocator.service
```

* Clone repository and update submodules:

```console
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git submodule update --init --recursive
```

* Enter the nix shell and run the enclave tests:

```console
nix --extra-experimental-features "nix-command flakes" develop
just compile-contracts
just espresso-enclave-tests
```

#### Building, running and registering enclave images

`op-batcher/enclave-tools` provides a command-line utility for common operations on batcher enclave images.
Before using it, set your AWS instance as described in the guide above, then build the tool:

```console
cd op-batcher/
just enclave-tools
```

This should create the `op-batcher/bin/enclave-tools` binary. For available commands and flags:

```console
./op-batcher/bin/enclave-tools --help
```

##### Building a batcher image

To build a batcher enclave image, and tag it with specified tag:

```console
./op-batcher/bin/enclave-tools build --op-root ./ --tag op-batcher-enclave
```

On success this command will output PCR measurements of the enclave image, which can then be registered with the BatchAuthenticator contract.

##### Running a batcher image

To run the enclave image built by the previous command:

```console
./op-batcher/bin/enclave-tools run --image op-batcher-enclave --args --argument-1,value-1,--argument-2,value-2
```

Arguments will be forwarded to the op-batcher.

##### Registering a batcher image

To register PCR0 of the batcher enclave image built by the previous command:

```console
./op-batcher/bin/enclave-tools register --l1-url example.com:1234 --authenticator 0x123..def --private-key 0x123..def --pcr0 0x123..def
```

You will need to provide the L1 URL, the contract address of BatchAuthenticator, private key of L1 account used to deploy BatchAuthenticator and PCR0 obtained when building the image.

# Local devnet

This section describes how to run a local devnet. There are two paths: using the convenience scripts (quick start) or driving docker compose manually (more control).

## Quick start via scripts

From the `espresso/scripts` directory:

```console
cd espresso/scripts
```

Prebuild everything and start all services (note `l2-genesis` takes ~2 minutes):

```console
./startup.sh
```

Or with AWS Nitro Enclave as the TEE:

```console
USE_TEE=true ./startup.sh
```

View logs for a service. There are 17 services tracked by `logs.sh`. Some names have convenient aliases (e.g. `sequencer` for `op-node-sequencer`). Add `-tee` to `batcher` and `proposer` when running with TEE.

```console
./logs.sh dev-node
./logs.sh sequencer
./logs.sh verifier
./logs.sh caff-node
./logs.sh batcher
./logs.sh proposer
```

Shut down all services:

```console
./shutdown.sh
```

## Manual setup via Docker Compose

* Ensure that your Docker Compose, Engine, and plugins are up-to-date. Particularly, if the Docker
Compose version is `2.37.3` or the Docker Engine version is `27.4.0`, and the Docker build hangs,
you may need to upgrade the version.

* Enter the Nix shell in the repo root:

  ```console
  nix develop .
  ```

* Build the op-deployer (re-run if the op-deployer is modified):

  ```console
  cd op-deployer
  just
  cd ../
  ```

* Build the contracts (re-run if the contracts are modified):

  ```console
  just compile-contracts
  ```

* Go to the `espresso` directory:

  ```console
  cd espresso
  ```

* Shut down all containers:

  ```console
  docker compose down -v --remove-orphans
  ```

  Or, if there are remaining containers from your last TEE run:

  ```console
  COMPOSE_PROFILES=tee docker compose down -v --remove-orphans
  docker rm -f $(docker ps -aq --filter "ancestor=op-batcher-enclavetool")
  ```

* Prepare OP contract allocations (re-run only when the OP contracts are modified):

  ```console
  ./scripts/prepare-allocs.sh
  ```

* Build and start all services in the background:

  ```console
  docker compose up --build -d
  ```

  If you're on a machine with [AWS Nitro Enclaves enabled](#guide-setting-up-an-enclave-enabled-nitro-ec2-instance), use the `tee` profile instead to start the enclave batcher:

  ```console
  COMPOSE_PROFILES=tee docker compose up --build -d
  ```

* Run the services and check the log:

  ```console
  docker compose logs -f
  ```

## Blockscout

Blockscout (block explorer reading from the caff node) is available at `http://localhost:3000`.

## Log monitoring

For a selection of important metrics and corresponding log lines see `espresso/docs/metrics.md`.

## Investigate a service

* Shut down all containers, then build and start only the specific service.

  ```console
  docker compose down
  docker compose up <service-name>
  ```

* If the environment variable setting is not picked up, pass it explicitly.

  ```console
  docker compose --env-file .env up <service-name>
  ```

## Apply a change

* In most cases, simply remove all containers and run commands as normal.

  ```console
  docker compose down
  ```

  For TEE, use the same shutdown command from the [Manual setup via Docker Compose](#manual-setup-via-docker-compose) section above.

* To start the project fresh, remove containers, volumes, and network, from this project.

  ```console
  docker compose down -v
  ```

* To start the system fresh, remove all volumes.

  ```console
  docker volume prune -a
  ```

* If encountering an issue related to outdated deployment files, remove those files before
restarting.
  * Go to the scripts directory.

  ```console
  cd espresso/scripts
  ```

  * Run the script.

  ```console
  ./cleanup.sh
  ```

* If you have changed OP contracts, you will have to start the devnet fresh and re-generate
  the genesis allocations by running `prepare-allocs.sh`.

## Continuous integration environment

### Running enclave tests in EC2

In order to run the tests for the enclave in EC2 via GitHub actions one must create an AWS user that supports the following policy:

```json
{
 "Version": "2012-10-17",
 "Statement": [
  {
   "Effect": "Allow",
   "Action": [
    "ec2:AuthorizeSecurityGroupIngress",
    "ec2:RunInstances",
    "ec2:DescribeInstances",
    "ec2:TerminateInstances",
    "ec2:DescribeImages",
    "ec2:CreateTags",
    "ec2:DescribeSecurityGroups",
    "ec2:DescribeKeyPairs",
    "ec2:ImportKeyPair",
    "ec2:DescribeInstanceStatus"
   ],
   "Resource": "*"
  }
 ]
}
```

Currently, the GitHub workflow in `.github/workflows/espresso-enclave.yaml` relies on AWS AMI with id `ami-0d259f3ae020af5f9` under `arn:aws:iam::324783324287`.
In order to refresh this AMI one needs to:

1. Create an AWS EC2 instance with the characteristics described in (see `.github/workflows/espresso-enclave.yaml` *Launch EC2 Instance* job).
2. Copy the script `espresso/scripts/enclave-prepare-ami.sh` in the EC2 instance (e.g. using scp) and run it.
3. [Export the AMI instance](https://docs.aws.amazon.com/toolkit-for-visual-studio/latest/user-guide/tkv-create-ami-from-instance.html).

# OP Succinct Lite and derivation pipeline dependencies

For the OP Succinct repository overview, branch table, and the procedure for propagating derivation pipeline changes through kona → celo-kona → op-succinct, see [`espresso/docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md`](espresso/docs/README_ESPRESSO_CODE_SYNC_PROCEDURE.md).

# Testnet migration

For deployment configuration parameters and the procedure for generating `allocs.json`, see [`espresso/docs/README_ESPRESSO_DEPLOY_CONFIG.md`](espresso/docs/README_ESPRESSO_DEPLOY_CONFIG.md).
