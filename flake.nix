{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    foundry = {
      # Pinned to be roughly the same as in mise.toml
      url = "github:shazow/foundry.nix/33a209625b9e31227a5f11417e95a3ac7264d811";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs = inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system:
      let
      overlays = [
        inputs.foundry.overlay
      ];
      pkgs = import inputs.nixpkgs { inherit overlays system;};
      in
      {
        devShell = pkgs.mkShell {
          packages = [
            pkgs.jq
            pkgs.yq-go
            pkgs.uv
            pkgs.shellcheck
            pkgs.python311
            pkgs.foundry-bin
            pkgs.just
            pkgs.go
            pkgs.gotools
          ];
        };
      }
    );
}
