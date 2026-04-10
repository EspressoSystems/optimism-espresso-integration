#!/bin/sh

export "${ENV_PREFIX}_GAME_FACTORY_ADDRESS"="$(jq -r '.opChainDeployments[0].DisputeGameFactoryProxy' ./deployer/state.json)"

"$@"
