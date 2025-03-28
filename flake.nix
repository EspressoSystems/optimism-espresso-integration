{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    foundry.url = "github:shazow/foundry.nix/main";
  };


  outputs = inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system:
      let
      overlays = [
        inputs.foundry.overlay
      ];
      pkgs = import inputs.nixpkgs { inherit overlays system;};
      downloadedFile = pkgs.fetchurl {
              url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/v0.0.34/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a";
              sha256 = "sha256:1c7ybrqjrp1709j08fk7zcr5q8hyfakvgv0m64zn2fywlqfdpszs";
            };
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
            pkgs.go_1_22
            pkgs.gotools
          ];
          shellHook = ''
                    export DOWNLOADED_FILE_PATH=${downloadedFile}
                    echo "Downloaded file is at $DOWNLOADED_FILE_PATH"
                    ln -sf /nix/store/8b5ranvnlb7sjrzpdpbb75vdp0gsyb1x-libespresso_crypto_helper-x86_64-unknown-linux-gnu.a /tmp/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a
                    export CGO_LDFLAGS="-L/tmp -lespresso_crypto_helper-x86_64-unknown-linux-gnu"
                  '';
        };
      }
    );
}
