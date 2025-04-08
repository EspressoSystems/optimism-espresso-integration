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
      espressoGoLibFile = if system == "x86_64-linux"
                then pkgs.fetchurl {
                  url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/v0.0.34/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a";
                  sha256 = "sha256:1c7ybrqjrp1709j08fk7zcr5q8hyfakvgv0m64zn2fywlqfdpszs";
                }
                else if system == "x86_64-darwin" then
                  pkgs.fetchurl {
                    url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/v0.0.34/libespresso_crypto_helper-x86_64-apple-darwin.a";
                    sha256 = "sha256:1fbijfam49c2i2l0d56i0zgczcbh2gljc6fh63g7qq3h7b7z5wc6";
                  }
                else # aarch64-darwin
                  pkgs.fetchurl {
                        url = "https://github.com/EspressoSystems/espresso-network-go/releases/download/v0.0.34/libespresso_crypto_helper-aarch64-apple-darwin.a";
                        sha256 = "sha256:18iqpqm3jmvj20vdd8zz05891lw5sxqy6vhfc8ghmg55czabip2q";
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
          ];
          shellHook = ''
          export FOUNDRY_DISABLE_NIGHTLY_WARNING=1
          export DOWNLOADED_FILE_PATH=${espressoGoLibFile}
          echo "Espresso go library stored at $DOWNLOADED_FILE_PATH"
          ln -sf ${espressoGoLibFile} ${target_link}
          export CGO_LDFLAGS="${cgo_ld_flags}"
                  '';
        };
      }
    );
}
