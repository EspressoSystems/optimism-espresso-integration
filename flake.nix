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
      espresso_go_lib_version = "v0.0.35";
      pkgs = import inputs.nixpkgs { inherit overlays system;};
      espressoGoLibFile = if system == "x86_64-linux"
                then pkgs.fetchurl {
                  url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/${espresso_go_lib_version}/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a";
                  sha256 = "sha256:07yfsrphfpq7w40x2rnldswzzbd4j0p5jdmm74132cqbf02pn8y8";
                }
                else if system == "x86_64-darwin" then
                  pkgs.fetchurl {
                    url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/${espresso_go_lib_version}/libespresso_crypto_helper-x86_64-apple-darwin.a";
                    sha256 = "sha256:1va49y81p3yrf9z61srw6rfysmbbk2vix0r7l8i2mz8b3ln0gsgy";
                  }
                else # aarch64-darwin
                  pkgs.fetchurl {
                        url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/${espresso_go_lib_version}/libespresso_crypto_helper-aarch64-apple-darwin.a";
                        sha256 = "sha256:1fp0v9d3b41lkfpva6rz35xi832xq4355pw5785ym2jm69pcsnnn";
                  }
                  ;
      cgo_ld_flags = if system == "x86_64-linux"
                      then "-L/tmp -lespresso_crypto_helper-x86_64-unknown-linux-gnu"
                      else if system == "x86_64-darwin" then "-L/tmp -lespresso_crypto_helper-x86_64-apple-darwin -framework Foundation -framework SystemConfiguration"
                      else "-L/tmp -lespresso_crypto_helper-aarch64-apple-darwin -framework Foundation -framework SystemConfiguration" # aarch64-darwin
      ;

      target_link =  if system == "x86_64-linux" then  "/tmp/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a"
                      else if system == "x86_64-darwin" then "/tmp/libespresso_crypto_helper-x86_64-apple-darwin.a"
                      else "/tmp/libespresso_crypto_helper-aarch64-apple-darwin.a" # aarch64-darwin
                      ;

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
            pkgs.go-ethereum
            pkgs.gotestsum
          ];
          shellHook = ''
            export FOUNDRY_DISABLE_NIGHTLY_WARNING=1
            export DOWNLOADED_FILE_PATH=${espressoGoLibFile}
            echo "Espresso go library ${espresso_go_lib_version} stored at $DOWNLOADED_FILE_PATH"
            ln -sf ${espressoGoLibFile} ${target_link}
            export CGO_LDFLAGS="${cgo_ld_flags}"
          '';
        };
      }
    );
}
