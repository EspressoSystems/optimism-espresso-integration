# Espresso Code Sync Procedure

## Schedule

- Refer to the “Recurrent: Celo Code Sync” section on Asana.
- If the time doesn’t work for you, let the team know (ideally one week in advance) so we can adjust the schedule.
- We may update the schedule and procedure for “cherry-pick to the integration branch”, pending discussion after we have code completion.
- Starting from December 2025, each Celo sync will also include syncing with Kona repositories.

## Terminologies

- *Celo integration repo*: [optimism-espresso-integration](https://github.com/EspressoSystems/optimism-espresso-integration) repo.
- *Celo tip branch*, or *tip branch*: `celo-tip-rebase-x` branch in the Celo integration repo, where `x` corresponds to the index Celo uses in their `celo-rebase-x` branch name. The Celo tip branch is directly synced from Celo.
- *Celo integration branch*: `celo-integration-rebase-x.y` branch in the Celo integration repo, where `x` corresponds to the index in the tip branch, and `y` corresponds to the index of our biweekly sync. The Celo integration branch contains our changes and Celo’s.
- *Terraform repo*: [tee-op-deploy](https://github.com/EspressoSystems/tee-op-deploy) repo, deployment code based on the Celo integration branch.
- *Kona fork repo*: [kona-celo-fork](https://github.com/EspressoSystems/kona-celo-fork) repo, forked from the `celo-org/kona` repo which is a fork of `op-rs/kona`, and contains our derivation changes.
- *Celo-Kona fork repo*: [celo-kona](https://github.com/EspressoSystems/celo-kona) repo, forked from the `celo-org/celo-kona` repo.
- *Succinct repo*: [op-succinct](https://github.com/EspressoSystems/op-succinct) repo, forked from the `celo-org/op-succinct` repo and dependent on the Kona fork and Celo Kona fork repos.

(Refer to [op-succinct-repos.png](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/docs/op-succinct-repos.png) for the relationship among Espresso and Celo repos.)

## 1 Procedure: Sync with Succinct

- (When: typically every other Friday, before syncing with Celo following [2 Procedure: Sync with Celo](#2-procedure-sync-with-celo).)
- Set a cutoff time and let the team know about this.
    - This is to prevent the case where a team member is working on something necessary to be merged to the default branch ASAP, but the code syncing process may block that.

### 1.1 Sync Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add kona-upstream https://github.com/celo-org/kona
git fetch kona-upstream
```

- If Celo's [default Kona branch](https://github.com/celo-org/kona/tree/replace-max-sequencer-drift-v1.1.7) has no updates since our last code sync, proceed to [1.2 Sync Celo-Kona Fork Repo](#12-sync-celo-kona-fork-repo).
    - Note: The default upstream branch is `replace-max-sequencer-drift-v1.1.7` as mentioned on [Slack](https://espressosys.slack.com/archives/C06LEU0LCN8/p1765799738195899?thread_ts=1765209556.168279&cid=C06LEU0LCN8).
- Otherwise, if the default Kona fork branch is `espresso-integration-x`, create a new branch `espresso-integration-y` where `y = x + 1`.

```bash
git checkout -b espresso-integration-y kona-upstream/replace-max-sequencer-drift-v1.1.7
```

- Cherry-pick commits from the default Kona fork branch onto Celo’s Kona branch.

```bash
git cherry-pick espresso-integration-x ^kona-upstream/replace-max-sequencer-drift-v1.1.7
```

- Follow the prompt to fix any cherry-pick issues.
- Double-check the commit history.

- Push the new branch *directly*. Add `--force` if needed.

```bash
git push -u origin espresso-integration-y
```

- Set the new branch as the default branch in GitHub repository settings.

### 1.2 Sync Celo-Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add celo-kona-upstream https://github.com/celo-org/celo-kona
git fetch celo-kona-upstream
```

- If Celo's [Celo-Kona release](https://github.com/celo-org/celo-kona/releases) has no updates since our last code sync and Celo has not informed us of a new version, proceed to [1.3 Sync Succinct Repo](#13-sync-succinct-repo).
  - Note: The release we should use is `v1.0.0` as mentioned on [Slack](https://espressosys.slack.com/archives/C06LEU0LCN8/p1769002077899559?thread_ts=1765209556.168279&cid=C06LEU0LCN8).
- Otherwise, if the default Celo-Kona fork branch is `espresso-integration-x`, create a new branch `espresso-integration-y` where `y = x + 1`.

```bash
git checkout -b espresso-integration-y celo-kona-upstream/release/v1.0.0-rc.4
```

- Cherry-pick commits from the default Celo-Kona fork branch onto Celo’s Celo-Kona branch.

```bash
git cherry-pick espresso-integration-x ^celo-kona-upstream/release/v1.0.0-rc.4
```

- Follow the prompt to fix any cherry-pick issues.
- Double-check the commit history.

- Push the new branch *directly*. Add `--force` if needed.

```bash
git push -u origin espresso-integration-y
```

- Set the new branch as the default branch.

### 1.3 Sync Succinct Repo

- Fetch the latest from upstream (if not done already).

```
git remote add succinct-upstream https://github.com/celo-org/op-succinct.git
git fetch succinct-upstream
```

- If Celo’s [default OP Succinct branch](https://github.com/celo-org/op-succinct) has no updates since our last code sync, proceed to [1.4 Update Imports in Succinct Repo](#14-update-imports-in-succinct-repo).
- Otherwise, if the default Succinct branch is `espresso-integration-x`, create a new branch `espresso-integration-y` where `y = x + 1`.

```bash
git checkout -b espresso-integration-y succinct-upstream/develop
```

- Cherry-pick commits from the default Succinct branch onto Celo’s Succinct branch.

```bash
git cherry-pick espresso-integration-x ^succinct-upstream/develop
```

- Follow the prompt to fix any cherry-pick issues.
- Double-check the commit history.

- Push the new branch *directly*. Add `--force` if needed.

```bash
git push -u origin espresso-integration-y
```

- Set the new branch as the default branch in GitHub repository settings.

- Start the [Build & push Celo fault-proof images](https://github.com/EspressoSystems/op-succinct/actions/workflows/fault-proof-celo-docker-build.yaml) CI workflow.
  - Make sure to use the link above since there is another CI workflow with the same name.

- Set the new default branches.

### 1.4 Update Imports in Succinct Repo

- If the kona and celo-kona repos were not updated in [1.1 Sync Kona Fork Repo](#11-sync-kona-fork-repo) and [1.2 Sync Celo-Kona Fork Repo](#12-sync-celo-kona-fork-repo), get the latest SHA of the [op-succinct-lite-proposer-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-proposer-celo) and [op-succinct-lite-challenger-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-challenger-celo) images and proceed to [2 Procedure: Sync with Celo](#2-procedure-sync-with-celo).
- Otherwise, update the `kona-*` and `celo-*` imports in `Cargo.toml`.
- Push the change to the new default branch, or if there is no such branch, create a PR and push to the original default branch.
- Get the latest SHA of the [op-succinct-lite-proposer-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-proposer-celo) and [op-succinct-lite-challenger-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-challenger-celo) images.

## 2 Procedure: Sync with Celo

- (When: typically every other Friday, after syncing with Kona repos following [1 Procedure: Sync with Succinct](#1-procedure-sync-with-succinct).)
- Set a cutoff time and let the team know about this.
    - This is to prevent the case where a team member is working on something necessary to be merged to the default branch ASAP, but the code syncing process may block that.

### 2.1 Update Celo integration

- Sync the Celo tip branch with the latest version at https://github.com/celo-org/optimism.
    - Note: Don’t use the “Sync fork” button because it will sync with Optimism’s `develop` branch.
    - Fetch the latest from upstream (if not done already).

    ```
    git remote add celo-upstream https://github.com/celo-org/optimism.git
    git fetch celo-upstream
    ```

    - If Celo’s [default branch](https://github.com/celo-org/optimism) has no updates since our last code sync, proceed to [2.2 Update Images in Celo Integration Repo ](#22-update-images-in-celo-integration-repo).
    - Otherwise, if Celo’s branch is on `x` and our tip branch is on `x.y`, create a new tip branch `celo-rebase-x.y'` where `y' = y + 1`.

    ```bash
    git checkout -b celo-tip-rebase-x.y' celo-upstream/celo-rebase-x
    git push origin celo-tip-rebase-x.y'
    ```

    - Otherwise, if Celo’s branch is on `x'` where `x' > x` and our tip branch is on `x.y`, create a new tip branch `celo-rebase-x'.0`.

    ```bash
    git checkout -b celo-tip-rebase-x'.0 celo-upstream/celo-rebase-x'
    git push origin celo-tip-rebase-x'.0
    ```

- Rebase the Celo integration branch onto the Celo tip branch.

    - Fetch the origin (if not done already).

    ```bash
    git fetch origin
    # --prune if you have any local setting
    ```

    - Create a new Celo integration branch `celo-integration-rebase-X.Y` matching the new tip branch created in the previous step.
        - Example: If you created tip branch `celo-tip-rebase-14.3`, create integration branch `celo-integration-rebase-14.3`

    ```bash
    # Replace X.Y with the version matching your new tip branch
    git checkout -b celo-integration-rebase-X.Y celo-tip-rebase-X.Y
    ```

    - Cherry-pick commits from the original Celo integration branch onto the tip branch.

    ```bash
    # Replace x.y with old version (e.g., 14.2) and X.Y with new version (e.g., 14.3)
    # This cherry-picks all Espresso-specific commits onto the new base
    git cherry-pick celo-tip-rebase-x.y..celo-integration-rebase-x.y
    ```

    - Follow the prompt to fix any cherry-pick issues.

    - Double-check the commit history.

    - Push the new branch *directly*. Add `--force` if needed.

    ```bash
    # Replace X.Y with your new branch version
    git push -u origin celo-integration-rebase-X.Y
    ```

### 2.2 Update Images in Celo Integration Repo

- If the Succinct images were not updated in [1.3 Sync Succinct Repo](#13-sync-succinct-repo), get the latest commit on the default branch and proceed to [2.3 Update Images in Terraform Repo](#23-update-images-in-terraform-repo).
- Otherwise, replace the image SHA of the `succinct-proposer` and `succinct-challenger` services in `docker-compose.yml`.
- Push the change to the new default branch, or if there is no such branch, create a PR and push to the original default branch.
- Get the latest commit on the default branch.

### 2.3 Update Images in Terraform Repo

- If the Celo integration repo is not updated with a new default branch or new Succinct images, proceed to [4 Procedure: Summary and Notification](#4-procedure-summary-and-notification).
- Otherwise, replace the `image_version` and `succinct_image_version` in `locals.tf`.
- Create a PR with the image update.
- After the PR is merged, proceed to [4 Procedure: Summary and Notification](#4-procedure-summary-and-notification).

# 3 Procedure: Cherry-Pick to Celo’s Upstreams

Note: This has not started yet. Eventually (perhaps after the testnet), we need a process to make sure Celo is updating its repos based on its upstreams, i.e., Optimism, Kona, and Succinct.

# 4 Procedure: Summary and Notification

- Document the new branches in [Code Sync Record](https://www.notion.so/espressosys/Code-Sync-Record-2e92431b68e98028901dc48c71aa8c3a).
- Let the team know that the Celo and Succinct sync is complete and they should be prepared to use the new default branches.

# 5 Making an Espresso-Specific Derivation Change

This section covers the workflow for making a change to our own derivation logic (not an upstream sync) and propagating it through the dependency chain.

## Repository Overview

The diagram below shows the relationship between the repositories.

![OP Succinct repos](op-succinct-repos.png)

| External | Celo (repo/branch) | Espresso (repo/branch) |
| :------: | :----------------: | :--------------------: |
| [kona](https://github.com/op-rs/kona) | [Celo/kona](https://github.com/celo-org/kona)/[palango/kona-1.1.7-celo](https://github.com/celo-org/kona/tree/palango/kona-1.1.7-celo) | [Espresso/kona-celo-fork](https://github.com/EspressoSystems/kona-celo-fork)/[espresso-integration](https://github.com/EspressoSystems/kona-celo-fork/tree/espresso-integration) |
| | [Celo/celo-kona](https://github.com/celo-org/celo-kona)/[main](https://github.com/celo-org/celo-kona/tree/main) | [Espresso/celo-kona](https://github.com/EspressoSystems/celo-kona)/[espresso-integration](https://github.com/EspressoSystems/celo-kona/tree/espresso-integration) |
| [op-succinct](https://github.com/succinctlabs/op-succinct) | [Celo/op-succinct](https://github.com/celo-org/op-succinct)/[develop](https://github.com/celo-org/op-succinct/tree/develop) | [Espresso/op-succinct](https://github.com/EspressoSystems/op-succinct)/[espresso-integration](https://github.com/EspressoSystems/op-succinct/tree/espresso-integration) |

## Propagation Steps

Changes to the derivation pipeline are made in the [kona-celo-fork](https://github.com/EspressoSystems/kona-celo-fork/tree/espresso-integration) repository, then propagated as follows:

1. Merge your PR into [kona-celo-fork](https://github.com/EspressoSystems/kona-celo-fork/tree/espresso-integration). E.g.: [bfabb62](https://github.com/EspressoSystems/kona-celo-fork/commit/bfabb62754bc53317ecb93442bb09d347cd6aad9).

2. Create a PR against [celo-kona](https://github.com/EspressoSystems/celo-kona/tree/espresso-integration) updating `Cargo.toml` to reference the updated kona version. E.g.: [a94b317](https://github.com/EspressoSystems/celo-kona/commit/a94b3172b1248a7cd650d692226c9d17b832eec9).

3. Create a PR against [op-succinct/espresso-integration](https://github.com/EspressoSystems/op-succinct/tree/espresso-integration) updating `Cargo.toml` to reference the updated kona and celo-kona versions. E.g.: [41780a3](https://github.com/EspressoSystems/op-succinct/pull/3/commits/41780a339bb1e177281957fcfe0383dfa41eff15).

4. After CI, check for new Docker images for the proposer and challenger:
   - [op-succinct-lite-proposer-celo](https://github.com/espressosystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-proposer-celo)
   - [op-succinct-lite-challenger-celo](https://github.com/espressosystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-challenger-celo)

   Update the image tags in [`espresso/docker-compose.yml`](../../espresso/docker-compose.yml). E.g.: [bd90858](https://github.com/EspressoSystems/optimism-espresso-integration/pull/293/commits/bd90858b0f871441785d4ac6437ff78b76d4b1f8).
