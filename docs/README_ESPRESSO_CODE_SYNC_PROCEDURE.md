# Espresso Code Sync Procesure

## Schedule

- Refer to the “Recurrent: Celo Code Sync” section on Asana.
- If the time doesn’t work for you, let the team know (ideally one week in advance) so we can adjust the schedule.
- We may update the schedule and procedure for “cherry-pick to the integration branch”, pending discussion after we have code completion.
- Starting from December 2025, each Celo sync will also include syncing with Kona repositories.

## Terminologies

- *Celo tip branch*, or *tip branch*: `celo-tip-rebase-x` branch, where `x` corresponds to the index Celo uses in their `celo-rebase-x` branch name. The Celo tip branch is directly synced from Celo.
- *Celo integration branch*: `celo-integration-rebase-x.y` branch, where `x` corresponds to the index in the tip branch, and `y` corresponds to the index of our biweekly sync. The Celo integration branch contains our changes and Celo’s.
- *Kona fork repo*: `kona` repo, forked from the `op-rs/kona` repo and contains our derivation changes.
- *Celo-Kona fork repo*: `kona-celo-fork` repo, forked from the `celo-org/kona` repo, which is a fork of `op-rs/kona`.
- *Succinct repo*: `op-succinct` repo, forked from the `celo-org/op-succinct` repo and imports the Kona fork and Celo Kona fork repos.

(Refer to [op-succinct-repos.png](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/docs/op-succinct-repos.png) for the relationship among Espresso and Celo repos.)

## Procedure: Sync with Celo

- (When: every other Friday, before syncing with Kona repos following [Procedure: Sync with Succinct](#procedure-sync-with-succinct).)
- Set a cutoff time and let the team know about this.
    - This is to prevent the case where a team member is working on something necessary to be merged to the default branch ASAP, but the code syncing process may block that.
- Sync the Celo tip branch with the latest version at https://github.com/celo-org/optimism.
    - Note: Don’t use the “Sync fork” button because it will sync with Optimism’s `develop` branch.
    - Fetch the latest from upstream (if not done already).

    ```
    git remote add celo-upstream https://github.com/celo-org/optimism.git
    git fetch celo-upstream
    ```

    - If Celo’s [default branch](https://github.com/celo-org/optimism) has no updates since our last code sync, proceed to [Procedure: Sync with Succinct](#procedure-sync-with-succinct).
    - Otherwise, if Celo’s branch is on `x` and our tip branch is on `x.y`, create a new tip branch `celo-rebase-x.y'` where `y' = y + 1`.

    ```
    git checkout -b celo-tip-rebase-x.y' celo-upstream/celo-rebase-x
    git push origin celo-tip-rebase-x.y'
    ```

    - Otherwise, if Celo’s branch is on `x'` where `x' > x` and our tip branch is on `x.y`, create a new tip branch `celo-rebase-x'.0`.

    ```
    git checkout -b celo-tip-rebase-x'.0 celo-upstream/celo-rebase-y
    git push origin celo-tip-rebase-x'.0
    ```

- Rebase the Celo integration branch onto the Celo tip branch.


    - Fetch the origin (if not done already).

    ```bash
    git fetch origin
    # --prune if you have any local setting
    ```

    - Fetch the old tip and the new tip

    ```bash
    git branch -a | grep celo-tip-rebase-x.y
    git branch -a | grep celo-tip-rebase-x'.y'
    # make sure you track the old tip locally
    git switch -c celo-tip-rebase-x.y --track origin/celo-tip-rebase-x.y
    ```

    - Create a new integration branch from the current integration branch

    ```bash
    git switch celo-integration-rebase-x.y
    git switch -c celo-integration-rebase-x'.y'
    ```

    - Rebase the integration branch onto the new tip branch.

    ```bash
    # rebase to the new tip with any changes not in the old tip
    git rebase --rebase-merges --onto celo-tip-rebase-x'.y' celo-tip-rebase-x.y
    ```

    - Resolve conflicts, if any.

    ```bash
    git status

    # Manually resolve conflicts. Some useful cmds:
    git rebase --skip # skip this commit if you see a duplicate one
    git rebase --edit-todo # check the following commits and update to `drop` or `squash` or `pick` if needed
    cat .git/rebase-merge/done # check the commits you've already done

    # run the following cmd after each conflict resolve
    git add . # or stage specific file change
    git rebase --continue
    ```

    - When the rebase finishes, you’ll see

    ```bash
    Successfully rebased and updated refs/heads/celo-integration-rebase-x'.y'.
    ```

    - Make sure the code compiles, then push the new branch *directly*.

    ```bash
    git push -u origin $(git branch --show-current)
    ```

    - Fix new errors. Make sure the CI passes.
    - An example

    ```bash
    git fetch origin --prune
    git branch -a | grep celo-tip-rebase-13.2
    git switch -c celo-tip-rebase-13.2 --track origin/celo-tip-rebase-13.2
    git switch celo-integration-rebase-13.2
    git switch -c celo-integration-rebase-14.1
    git rebase --rebase-merges --onto celo-tip-rebase-14.1 celo-tip-rebase-13.2
    git push -u origin $(git branch --show-current)
    ```


## Procedure: Sync with Succinct

- (When: every other Friday, after syncing with Celo following [Procedure: Sync with Celo](#procedure-sync-with-celo).)
- Set a cutoff time and let the team know about this.
    - This is to prevent the case where a team member is working on something necessary to be merged to the default branch ASAP, but the code syncing process may block that.

### 1. Sync Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add kona-upstream https://github.com/celo-org/kona
git fetch kona-upstream
```

- If Celo’s [default Kona branch](https://github.com/celo-org/kona/tree/replace-max-sequencer-drift-v1.1.7) has no updates since our last code sync, proceed to [2. Sync Celo-Kona Fork Repo](#2-sync-celo-kona-fork-repo).
    - Note: The default branch is `replace-max-sequencer-drift-v1.1.7` as mentioned on [Slack](https://espressosys.slack.com/archives/C06LEU0LCN8/p1765799738195899?thread_ts=1765209556.168279&cid=C06LEU0LCN8).
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Kona branch.

```
git checkout -b espresso-integration-x kona-upstream/main
```

- Rebase the original Kona fork branch `espresso-integration-x` onto Celo’s Succinct branch.

```jsx
git rebase espresso-integration-x
```

- Resolve conflicts, if any.
- Push the new branch *directly*.

### 2. Sync Celo-Kona Fork Repo

- Fetch the latest from upstream (if not done already).

```
git remote add celo-kona-upstream https://github.com/celo-org/celo-kona
git fetch celo-kona-upstream
```

- If Celo’s [default Celo-Kona branch](https://github.com/celo-org/celo-kona) has no updates since our last code sync, proceed to [3. Sync Succinct Repo](#3-sync-succinct-repo).
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Celo-Kona branch.

```
git checkout -b espresso-integration-x celo-kona-upstream/main
```

- Rebase the original Celo-Kona fork branch `espresso-integration-x` onto Celo’s Celo-Kona branch.

```jsx
git rebase espresso-integration-x
```

- Resolve conflicts, if any.
- Make sure the CI passes.
- Push the new branch *directly*.

### 3. Sync Succinct Repo

- Fetch the latest from upstream (if not done already).

```
git remote add succinct-upstream https://github.com/celo-org/op-succinct.git
git fetch succinct-upstream
```

- If Celo’s [default OP Succinct branch](https://github.com/celo-org/op-succinct) has no updates since our last code sync, skip this week and let the team know.
- Otherwise, create a sync branch `espresso-integration-y` where `y` is the commit on Celo’s Succinct branch.

```
git checkout -b espresso-integration-x succinct-upstream/develop
```

- Rebase the original Succinct branch `espresso-integration-x` onto Celo’s Succinct branch.

```jsx
git rebase espresso-integration-x
```

- Resolve conflicts, if any.
- Push the new branch *directly*.
- Make sure the CI passes.
- Let the team know the Celo and Succinct sync is complete and update the default branches.
    - It is expected to be done by EOD next Monday, but we do not usually have a hard deadline for this, so just make sure to communicate with the team about the progress.

# Procedure: Cherry-Pick to Celo’s Upstreams

Note: This has not started yet. Eventually (perhaps after the testnet), we need a process to make sure Celo is updating its repos based on its upstreams, i.e., Optimism, Kona, and Succinct.
