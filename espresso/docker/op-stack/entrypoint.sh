#!/bin/sh

export "${ENV_PREFIX}_GAME_FACTORY_ADDRESS"=$(jq -r '.opChainDeployments[0].disputeGameFactoryProxyAddress' ./deployer/state.json)

"$@"
