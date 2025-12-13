# Celo Testnet Migration

## Overview

This document serves as the main collaboration tool between Celo Labs and Espresso Systems teams to agree on the migration process from Celo Testnet to its Espresso-integrated version.

Currently this document contains questions and will evolve into a detailed guide describing the concrete steps of the migration.

---

## Questions from Espresso to Celo

### Big Picture

#### Q1: What are Celo's goals that Espresso should be aware of?

In particular, which Celo Testnet should we target for migration?

*Awaiting response from Celo.*

#### Q2: Are there migration-related resources you can share with us?

For example: scripts, documentation, or runbooks from previous upgrades.

*Awaiting response from Celo.*

#### Q3: Can Celo share SLAs on all services provided to users?

It would be helpful to highlight any SLAs that are directly related to, or impacted by, the batch poster (e.g., finality guarantees, uptime commitments).

*Awaiting response from Celo.*

---

### Support Model

#### Q4: What are the deployment phases?

For example, we could have:

| Phase | Description | Duration |
|-------|-------------|----------|
| **Devnet** | Espresso & Celo testing only | Temporary |
| **Testnet** | Espresso, Celo, and community testing | Persistent |
| **Mainnet** | Production deployment | Permanent |


*Awaiting response from Celo.*

#### Q5: What reporting and logging does Celo have currently?

- What alerts are automated?
- What automations should Espresso replicate for the batcher?

*Awaiting response from Celo.*

#### Q6: What incidents should one team vs both teams be responsible for?

It would be helpful to identify:
- What errors require both teams to be on-call?
- When should Celo provide logs to Espresso?

**Proposed Incident Matrix:**

| Incident Type | Primary | Support |
|---------------|---------|---------|
| Batcher not posting batches | Espresso | Celo (if network issue) |
| op-node / op-geth issues | Celo | Espresso (if integration-related) |
| L1 contract issues | Both | Both |
| Chain liveness issues | Both | Both |

*Awaiting response from Celo.*

#### Q7: What are Celo's batch posting SLA expectations?

For example:
- Expected batch posting interval (e.g., every 30 minutes)?
- Acceptable finality lag threshold?
- Uptime requirements (e.g., 99.9%)?

*Awaiting response from Celo.*

#### Q8: What is Celo's disaster recovery approach?

- If the state is corrupted, does Celo capture snapshots of the rollup state that Espresso could use for recovery?
- If Celo loses a server, what happens? (This helps with our own testing.)

*Awaiting response from Celo.*

---

### Access Requirements

#### Q9: What access does Espresso need from Celo?

Espresso needs the following to operate the TEE batcher:

| Requirement | Description |
|-------------|-------------|
| Sequencer RPC URL | RPC access to the sequencer |
| Chain configuration | Chain config files for the rollup |
| Batcher configuration | Current `op-batcher` configuration |
| State directory access | Access to `op-node` state (if migrating existing state) |

*Awaiting response from Celo.*

---

## Questions from Celo to Espresso

*Please add questions here.*

---

## Resources

### Espresso Integration
- [Source code (optimism-espresso-integration)](https://github.com/EspressoSystems/optimism-espresso-integration)
- [Local devnet configuration guide](./README_ESPRESSO_DEPLOY_CONFIG.md)

### OP Stack References
- [OP Network Upgrades Guide](https://docs.optimism.io/op-stack/protocol/network-upgrades)
- [OP Upgrade 17 Notice](https://docs.optimism.io/notices/upgrade-17)

### Celo References
- [Celo L2 Forum Post](https://forum.celo.org/t/celo-as-an-ethereum-l2-a-frontier-chain-for-global-impact/11376)
- [Jello Hardfork Announcement](https://forum.celo.org/t/introducing-the-jello-hardfork-op-succinct-lite-now-live-on-celo-sepolia/12603)
