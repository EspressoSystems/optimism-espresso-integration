# Celo Migration Guide

## Summary

This section outlines the proposed migration process and highlights key questions for the Celo team. The goal is to ensure both teams are aligned on requirements and responsibilities, so that the migration can be planned and executed smoothly.

### Proposed Migration Steps

#### Deploying New L1 Contracts

Celo's team will deploy the two new contracts on the L1.

1. Deploy your two contracts: Deploy BatchAuthenticator and BatchInbox, Espresso team will provide instructions and scripts
2. Update the rollup config:

```json
{
  "batch_inbox_address": "0x<your-new-BatchInbox-address>",
  ...
}
```

3. Restart op-node and op-batcher with the new config

**Done!** No L1 contract upgrades needed.

We should ask Celo if they plan to follow these steps for the migration or they have other guides. Espresso should provide a script containing the contract artifacts (Solidity code or compiled binaries and ABIs). Once deployed, contract address should be recorded and contract should be verified.

---

#### New Config for `op-batcher` and `op-node`

Celo provides batcher key and contract address. Updates to the espresso-integrated release version for the component needs updates. And configure the activation timestamp.

#### Restarting `op-node` and `op-batcher`

Migration will be activated at the configured timestamp, more details [here in our design book](https://github.com/EspressoSystems/book/pull/90), OP also has similar instructions (see step 2 [here](https://docs.optimism.io/notices/upgrade-17#for-node-operators)).

### Espresso Team's Plan

- **Sync with Celo Team regularly:** Sync with the Celo Team every week on the migration design / implementation progress. Ensure both teams are aware of the status and align on next steps.
## Celo Deployment

### Prepare for the Deployment
* Go to the scripts directory.
```console
cd espresso/scripts
```

### Prebuild Everything and Start All Services
Note that `l2-genesis` is expected to take around 2 minutes.
```console
./startup.sh
```
Or build and start the devnet with AWS Nitro Enclave as the TEE:
```console
USE_TEE=true ./startup.sh
```

### View Logs
There are 17 services in total, as listed in `logs.sh`. Run the script with the service name to
view its logs, e.g., `./logs.sh op-geth-sequencer`. Note that some service names can be replaced
by more convenient alias, e.g., `sequencer` instead of `op-node-sequencer`, but it is also suported
to use their full names.

The following are common commands to view the logs of critical services. Add `-tee` to the batcher
and the proposer services if running with the TEE.
```console
./logs.sh dev-node
./logs.sh sequencer
./logs.sh verifier
./logs.sh caff-node
./logs.sh batcher
./logs.sh proposer
```

### Shut Down All Services
```console
./shutdown.sh
```

# Celo Migration Guide

## Summary

This section outlines the proposed migration process and highlights key questions for the Celo team. The goal is to ensure both teams are aligned on requirements and responsibilities, so that the migration can be planned and executed smoothly.

### Proposed Migration Steps

#### Deploying New L1 Contracts

Celo's team will deploy the two new contracts on the L1.

1. Deploy your two contracts: Deploy BatchAuthenticator and BatchInbox, Espresso team will provide instructions and scripts
2. Update the rollup config:

```json
{
  "batch_inbox_address": "0x<your-new-BatchInbox-address>",
  ...
}
```

3. Restart op-node and op-batcher with the new config

**Done!** No L1 contract upgrades needed.

We should ask Celo if they plan to follow these steps for the migration or they have other guides. Espresso should provide a script containing the contract artifacts (Solidity code or compiled binaries and ABIs). Once deployed, contract address should be recorded and contract should be verified.

---

#### New Config for `op-batcher` and `op-node`

Celo provides batcher key and contract address. Updates to the espresso-integrated release version for the component needs updates. And configure the activation timestamp.

#### Restarting `op-node` and `op-batcher`

Migration will be activated at the configured timestamp, more details [here in our design book](https://github.com/EspressoSystems/book/pull/90), OP also has similar instructions (see step 2 [here](https://docs.optimism.io/notices/upgrade-17#for-node-operators)).

### Espresso Team's Plan

- **Sync with Celo Team regularly:** Sync with the Celo Team every week on the migration design / implementation progress. Ensure both teams are aware of the status and align on next steps.

- **Devnet setup:** Leverage existing Espresso tooling for devnet deployment, see our [docker compose file](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/espresso/docker-compose.yml) to start it locally and [terraform deployment script (still private)](https://github.com/EspressoSystems/tee-op-deploy) to start it in AWS' remote services. Adapt tooling to Celo's targeted migration testnet (currently [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to create a devnet environment that closely mirrors the production setup.

- **Migration rehearsal:** Rehearse all migration codes, scripts, and operational steps:
  - Validate contract deployment scripts, make sure deployed contracts are verified
  - Test configuration updates for op-node and op-batcher
  - Verify activation timestamp logic
  - Test rollback procedures, ensure ability to revert

  Adjust them based on test results. Keep the issues/resolutions/scripts/steps updated in this readme.

- **Run migration on real testnet:** Apply migration to Celo's live testnet. Monitor and maintain the testnet for at least one week. Make sure all transactions are processing correctly through Espresso sequencer, batches are successfully posted on L1, collect feedback from Celo team and community.

- **Run migration on mainnet:** Following the proven procedures from testnet, coordinate closely with Celo team on timing and execution, have rollback plan ready, provide ongoing support and monitoring.

### Coordination with Celo Team (Key Questions to Ask)

- **Identify Points of Contact:** Who will be the point of contact at Celo for migration test and deployment.

- **Testnet Environment:** How Celo runs their testnets and plans to test the Espresso-enabled chain. Do they use an OP Stack Kurtosis devnet, or will they deploy a [public test network](https://forum.celo.org/t/celo-as-an-ethereum-l2-a-frontier-chain-for-global-impact/11376) for this upgrade? (Celo has recently launched "Eclair" testnet with other upgrades.) Whether we should use an existing testnet (like Alfajores/Eclair) or spin up a dedicated devnet (like [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to rehearse the migration.

- **Division of Responsibilities:** Confirm what is expected from each team during the migration. For example, Celo's DevOps team would likely handle deploying the new L1 contracts (BatchInbox and BatchAuthentication) and executing any on-chain upgrade transactions (needs confirmation), while Espresso's team provides the contract code, modified node software, and guidance. As the new batcher needs to be run in TEE, Espresso will run it post-upgrade. In such case, Espresso will need access to the batcher's private key and more back-and-forth with Celo's team, and we need to determine who will fund the batcher. If Celo is going to run it in the future, we should plan on supporting batcher running in TDX.

- **Migration Steps:** Does Celo have an internal procedure or guide from those upgrades that we can piggyback on? I haven't found any relevant recent doc on this yet. Or will Celo perform upgrade following [OP's guide](https://docs.optimism.io/op-stack/protocol/network-upgrades) (example with more details: [this](https://docs.optimism.io/notices/upgrade-17) and [this](https://docs.optimism.io/notices/fusaka-notice))?


- **Devnet setup:** Leverage existing Espresso tooling for devnet deployment, see our [docker compose file](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/espresso/docker-compose.yml) to start it locally and [terraform deployment script (still private)](https://github.com/EspressoSystems/tee-op-deploy) to start it in AWS' remote services. Adapt tooling to Celo's targeted migration testnet (currently [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to create a devnet environment that closely mirrors the production setup.

- **Migration rehearsal:** Rehearse all migration codes, scripts, and operational steps:
  - Validate contract deployment scripts, make sure deployed contracts are verified
  - Test configuration updates for op-node and op-batcher
  - Verify activation timestamp logic
  - Test rollback procedures, ensure ability to revert

  Adjust them based on test results. Keep the issues/resolutions/scripts/steps updated in this readme.

- **Run migration on real testnet:** Apply migration to Celo's live testnet. Monitor and maintain the testnet for at least one week. Make sure all transactions are processing correctly through Espresso sequencer, batches are successfully posted on L1, collect feedback from Celo team and community.

- **Run migration on mainnet:** Following the proven procedures from testnet, coordinate closely with Celo team on timing and execution, have rollback plan ready, provide ongoing support and monitoring.

### Coordination with Celo Team (Key Questions to Ask)
## Celo Deployment

### Prepare for the Deployment
* Go to the scripts directory.
```console
cd espresso/scripts
```

### Prebuild Everything and Start All Services
Note that `l2-genesis` is expected to take around 2 minutes.
```console
./startup.sh
```
Or build and start the devnet with AWS Nitro Enclave as the TEE:
```console
USE_TEE=true ./startup.sh
```

### View Logs
There are 17 services in total, as listed in `logs.sh`. Run the script with the service name to
view its logs, e.g., `./logs.sh op-geth-sequencer`. Note that some service names can be replaced
by more convenient alias, e.g., `sequencer` instead of `op-node-sequencer`, but it is also suported
to use their full names.

The following are common commands to view the logs of critical services. Add `-tee` to the batcher
and the proposer services if running with the TEE.
```console
./logs.sh dev-node
./logs.sh sequencer
./logs.sh verifier
./logs.sh caff-node
./logs.sh batcher
./logs.sh proposer
```

### Shut Down All Services
```console
./shutdown.sh
```

# Celo Migration Guide

## Summary

This section outlines the proposed migration process and highlights key questions for the Celo team. The goal is to ensure both teams are aligned on requirements and responsibilities, so that the migration can be planned and executed smoothly.

### Proposed Migration Steps

#### Deploying New L1 Contracts

Celo's team will deploy the two new contracts on the L1.

1. Deploy your two contracts: Deploy BatchAuthenticator and BatchInbox, Espresso team will provide instructions and scripts
2. Update the rollup config:

```json
{
  "batch_inbox_address": "0x<your-new-BatchInbox-address>",
  ...
}
```

3. Restart op-node and op-batcher with the new config

**Done!** No L1 contract upgrades needed.

We should ask Celo if they plan to follow these steps for the migration or they have other guides. Espresso should provide a script containing the contract artifacts (Solidity code or compiled binaries and ABIs). Once deployed, contract address should be recorded and contract should be verified.

---

#### New Config for `op-batcher` and `op-node`

Celo provides batcher key and contract address. Updates to the espresso-integrated release version for the component needs updates. And configure the activation timestamp.

#### Restarting `op-node` and `op-batcher`

Migration will be activated at the configured timestamp, more details [here in our design book](https://github.com/EspressoSystems/book/pull/90), OP also has similar instructions (see step 2 [here](https://docs.optimism.io/notices/upgrade-17#for-node-operators)).

### Espresso Team's Plan

- **Sync with Celo Team regularly:** Sync with the Celo Team every week on the migration design / implementation progress. Ensure both teams are aware of the status and align on next steps.

- **Devnet setup:** Leverage existing Espresso tooling for devnet deployment, see our [docker compose file](https://github.com/EspressoSystems/optimism-espresso-integration/blob/celo-integration-rebase-14.1/espresso/docker-compose.yml) to start it locally and [terraform deployment script (still private)](https://github.com/EspressoSystems/tee-op-deploy) to start it in AWS' remote services. Adapt tooling to Celo's targeted migration testnet (currently [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to create a devnet environment that closely mirrors the production setup.

- **Migration rehearsal:** Rehearse all migration codes, scripts, and operational steps:
  - Validate contract deployment scripts, make sure deployed contracts are verified
  - Test configuration updates for op-node and op-batcher
  - Verify activation timestamp logic
  - Test rollback procedures, ensure ability to revert

  Adjust them based on test results. Keep the issues/resolutions/scripts/steps updated in this readme.

- **Run migration on real testnet:** Apply migration to Celo's live testnet. Monitor and maintain the testnet for at least one week. Make sure all transactions are processing correctly through Espresso sequencer, batches are successfully posted on L1, collect feedback from Celo team and community.

- **Run migration on mainnet:** Following the proven procedures from testnet, coordinate closely with Celo team on timing and execution, have rollback plan ready, provide ongoing support and monitoring.

### Coordination with Celo Team (Key Questions to Ask)

- **Identify Points of Contact:** Who will be the point of contact at Celo for migration test and deployment.

- **Testnet Environment:** How Celo runs their testnets and plans to test the Espresso-enabled chain. Do they use an OP Stack Kurtosis devnet, or will they deploy a [public test network](https://forum.celo.org/t/celo-as-an-ethereum-l2-a-frontier-chain-for-global-impact/11376) for this upgrade? (Celo has recently launched "Eclair" testnet with other upgrades.) Whether we should use an existing testnet (like Alfajores/Eclair) or spin up a dedicated devnet (like [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to rehearse the migration.

- **Division of Responsibilities:** Confirm what is expected from each team during the migration. For example, Celo's DevOps team would likely handle deploying the new L1 contracts (BatchInbox and BatchAuthentication) and executing any on-chain upgrade transactions (needs confirmation), while Espresso's team provides the contract code, modified node software, and guidance. As the new batcher needs to be run in TEE, Espresso will run it post-upgrade. In such case, Espresso will need access to the batcher's private key and more back-and-forth with Celo's team, and we need to determine who will fund the batcher. If Celo is going to run it in the future, we should plan on supporting batcher running in TDX.

- **Migration Steps:** Does Celo have an internal procedure or guide from those upgrades that we can piggyback on? I haven't found any relevant recent doc on this yet. Or will Celo perform upgrade following [OP's guide](https://docs.optimism.io/op-stack/protocol/network-upgrades) (example with more details: [this](https://docs.optimism.io/notices/upgrade-17) and [this](https://docs.optimism.io/notices/fusaka-notice))?

lfajores/Eclair) or spin up a dedicated devnet (like [Jello](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)) to rehearse the migration.

- **Division of Responsibilities:** Confirm what is expected from each team during the migration. For example, Celo's DevOps team would likely handle deploying the new L1 contracts (BatchInbox and BatchAuthentication) and executing any on-chain upgrade transactions (needs confirmation), while Espresso's team provides the contract code, modified node software, and guidance. As the new batcher needs to be run in TEE, Espresso will run it post-upgrade. In such case, Espresso will need access to the batcher's private key and more back-and-forth with Celo's team, and we need to determine who will fund the batcher. If Celo is going to run it in the future, we should plan on supporting batcher running in TDX.

- **Migration Steps:** Does Celo have an internal procedure or guide from those upgrades that we can piggyback on? I haven't found any relevant recent doc on this yet. Or will Celo perform upgrade following [OP's guide](https://docs.optimism.io/op-stack/protocol/network-upgrades) (example with more details: [this](https://docs.optimism.io/notices/upgrade-17) and [this](https://docs.optimism.io/notices/fusaka-notice))?

