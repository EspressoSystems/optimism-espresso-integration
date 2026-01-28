# Espresso Code Sync Procesure

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
- *Kona fork repo*: [kona-celo-fork](https://github.com/EspressoSystems/kona-celo-fork/tree/espresso-integration) repo, forked from the `celo-org/kona` repo which is a fork of `op-rs/kona`, and contains our derivation changes.
- *Celo-Kona fork repo*: [celo-kona](https://github.com/EspressoSystems/celo-kona/tree/espresso-integration) repo, forked from the `celo-org/celo-kona` repo.
- *Succinct repo*: [op-succinct](https://github.com/EspressoSystems/op-succinct/tree/espresso-integration) repo, forked from the `celo-org/op-succinct` repo and dependent on the Kona fork and Celo Kona fork repos.

(Refer to [op-succinct-repos.png](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/docs/op-succinct-repos.png) for the relationship among Espresso and Celo repos.)

## 1 Procedure: Sync with Succinct

- (When: typically every other Friday, before syncing with Celo following [2 Procedure: Sync with Celo](#2-procedure-sync-with-celo).)
- Set a cutoff time and let the team know about this.
    - This is to prevent the case where a team member is working on something necessary to be merged to the default branch ASAP, but the code syncing process may block that.

### 1.1. Sync Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add kona-upstream https://github.com/celo-org/kona
git fetch kona-upstream
```

- If Celo’s [default Kona branch](https://github.com/celo-org/kona/tree/replace-max-sequencer-drift-v1.1.7) has no updates since our last code sync, proceed to [1-2 Sync Celo-Kona Fork Repo](#1-2-sync-celo-kona-fork-repo).
    - Note: The default upstream branch is `replace-max-sequencer-drift-v1.1.7` as mentioned on [Slack](https://espressosys.slack.com/archives/C06LEU0LCN8/p1765799738195899?thread_ts=1765209556.168279&cid=C06LEU0LCN8).
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Kona branch.

```bash
git checkout -b espresso-integration-x kona-upstream/replace-max-sequencer-drift-v1.1.7
```

- Cherry-pick commits from the original Kona branch `espresso-integration-x` onto Celo’s Kona branch.

```bash
git cherry-pick espresso-integration-x ^kona-upstream/replace-max-sequencer-drift-v1.1.7
```

- Follow the prompt to fix any cherry-pick issues.
- Double-check the commit history.

- Push the new branch *directly*. Add `--force` if needed.

```bash
git push -u origin espresso-integration-y
```

- Set the new branch as the default branch.

### 1.2 Sync Celo-Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add celo-kona-upstream https://github.com/celo-org/celo-kona
git fetch celo-kona-upstream
```

- If Celo’s [default Celo-Kona branch](https://github.com/celo-org/celo-kona/tree/release/v1.0.0-rc.4) has no updates since our last code sync, proceed to [1-3 Sync Succinct Repo](#1-3-sync-succinct-repo).
  - Note: The default upstream branch we are using is `release/v1.0.0-rc.4` as of January, 2026.
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Celo-Kona branch.

```bash
git checkout -b espresso-integration-x celo-kona-upstream/release/v1.0.0-rc.4
```

- Cherry-pick commits from the original Celo-Kona fork branch `espresso-integration-x` onto Celo’s Celo-Kona branch.

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

- If Celo’s [default OP Succinct branch](https://github.com/celo-org/op-succinct) has no updates since our last code sync, proceed to [2 Procedure: Sync with Celo](#2-procedure-sync-with-celo).
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Succinct branch.

```bash
git checkout -b espresso-integration-x succinct-upstream/develop
```

- Cherry-pick commits from the original Succinct branch `espresso-integration-x` onto Celo’s Succinct branch.

```bash
git cherry-pick espresso-integration-x ^succinct-upstream/develop
```

- Follow the prompt to fix any cherry-pick issues.
- Double-check the commit history.

- Push the new branch *directly*. Add `--force` if needed.

```bash
git push -u origin espresso-integration-y
```

- Set the new branch as the default branch.

- Start the [Build & push Celo fault-proof images](https://github.com/EspressoSystems/op-succinct/actions/workflows/fault-proof-celo-docker-build.yaml) CI workflow.
  - Make sure to use the link above since there is another CI workflow with the same name.

- After the CI completes, get the latest SHA of the [op-succinct-lite-proposer-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-proposer-celo) and [op-succinct-lite-challenger-celo](https://github.com/EspressoSystems/op-succinct/pkgs/container/op-succinct%2Fop-succinct-lite-challenger-celo) and proceed to [2 Procedure: Sync with Celo](#2-procedure-sync-with-celo).

- Set the new default branches.

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

    - If Celo’s [default branch](https://github.com/celo-org/optimism) has no updates since our last code sync, proceed to [2.2 Update Images in Celo Integration Repo ](#2-2-update-images-in-celo-integration-repo).
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

    - Create a new Celo integration branch `celo-integration-rebase-x'-y'` where `x` and `y` are consistent with the tip branch created in the previous step.

    ```bash
    git checkout -b celo-integration-rebase-x'-y' celo-tip-rebase-x'-y'
    ```

    - Cherry-pick commits from the original Celo integration branch onto the tip branch.

    ```bash
    git cherry-pick celo-integration-rebase-x-y ^celo-tip-rebase-x'-y'
    ```

    - Follow the prompt to fix any cherry-pick issues.

    - Double-check the commit history.

    - Push the new branch *directly*. Add `--force` if needed.

    ```bash
    git push -u origin celo-integration-rebase-x'-y'
    ```

### 2.2 Update Images in Celo Integration Repo

- If the Succinct images were not updated in [1-3 Sync Succinct Repo](#1-3-sync-succinct-repo), get the latest commit on the default branch and proceed to [2-3 Update Images in Terraform Repo](#2-3-update-images-in-terraform-repo).
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
    - It is expected to be done by EOD next Monday, but we do not usually have a hard deadline for this, so just make sure to communicate with the team about the progress.
