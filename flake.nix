{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    foundry.url = "github:shazow/foundry.nix/main";
  };

  outputs =
    inputs:
    inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        overlays = [
          inputs.foundry.overlay
        ];

        go_1_22_7 = pkgs.go_1_22.overrideAttrs (oldAttrs: rec {
          version = "1.22.7";

          src = pkgs.fetchurl {
            url = "https://go.dev/dl/go1.22.7.src.tar.gz";
            sha256 = "sha256-ZkMth9heDPrD7f/mN9WTD8Td9XkzE/4R5KDzMwI8h58=";
          };
        });

        pkgs = import inputs.nixpkgs { inherit overlays system; };
        espressoGoLibVersion = "0.2.1";
        baseUrl = "https://github.com/EspressoSystems/espresso-network/releases/download/sdks%2Fgo%2Fv${espressoGoLibVersion}";
        espressoGoLibFile =
          if system == "x86_64-linux" then
            pkgs.fetchurl {
              url = baseUrl + "/libespresso_crypto_helper-x86_64-unknown-linux-gnu.so";
              sha256 = "sha256:b3e28f7dc755d72b27a2a43c2bcfdc0e4e82096e03596a01447bd8f406e6653c";
            }
          else if system == "x86_64-darwin" then
            pkgs.fetchurl {
              url = baseUrl + "/libespresso_crypto_helper-x86_64-apple-darwin.dylib";
              sha256 = "sha256:716cb9eb548222ed1c7b5d1585bd5f03d0680cbae3f8db14cbf37837f54b9788";
            }
          # aarch64-darwin
          else
            pkgs.fetchurl {
              url = baseUrl + "/libespresso_crypto_helper-aarch64-apple-darwin.dylib";
              sha256 = "sha256:6c74ec631ccd9d23258ff99a8060068a548740fac814633ceab2ad7c7dc90a74";
            };
        cgo_ld_flags =
          if system == "x86_64-linux" then
            "-L/tmp -lespresso_crypto_helper-x86_64-unknown-linux-gnu"
          else if system == "x86_64-darwin" then
            "-L/tmp -lespresso_crypto_helper-x86_64-apple-darwin -framework Foundation -framework SystemConfiguration"
          else
            "-L/tmp -lespresso_crypto_helper-aarch64-apple-darwin -framework Foundation -framework SystemConfiguration" # aarch64-darwin
        ;

        target_link =
          if system == "x86_64-linux" then
            "/tmp/libespresso_crypto_helper-x86_64-unknown-linux-gnu.so"
          else if system == "x86_64-darwin" then
            "/tmp/libespresso_crypto_helper-x86_64-apple-darwin.dylib"
          else
            "/tmp/libespresso_crypto_helper-aarch64-apple-darwin.dylib" # aarch64-darwin
        ;

        enclaver = pkgs.rustPlatform.buildRustPackage rec {
          pname = "enclaver";
          version = "0.5.0";

          src = pkgs.fetchFromGitHub {
            owner = "enclaver-io";
            repo = pname;
            rev = "v${version}";
            hash = "sha256-gfzfgcnVDRqywAJ/SC2Af6VfHPELDkoVlkhaKElMP2g=";
          };

          useFetchCargoVendor = true;
          cargoHash = "sha256-o+CzTn5++Mj6SP9yFeTOBn4feapnL2m1EsYmXQBqTuc=";
          cargoRoot = "enclaver";
          buildAndTestSubdir = cargoRoot;
        };

      in
      {

        formatter = pkgs.nixfmt-rfc-style;

        devShells = {
          default = pkgs.mkShell {
            packages = [
              pkgs.zlib
              enclaver
              pkgs.jq
              pkgs.yq-go
              pkgs.uv
              pkgs.shellcheck
              pkgs.python311
              pkgs.foundry-bin
              pkgs.just
              go_1_22_7
              pkgs.gotools
              pkgs.go-ethereum
              pkgs.golangci-lint
              pkgs.awscli2
              pkgs.just
            ];
            shellHook = ''
              export FOUNDRY_DISABLE_NIGHTLY_WARNING=1
              export DOWNLOADED_FILE_PATH=${espressoGoLibFile}
              echo "Espresso go library v${espressoGoLibVersion} stored at $DOWNLOADED_FILE_PATH"
              ln -sf ${espressoGoLibFile} ${target_link}
              export CGO_LDFLAGS="${cgo_ld_flags} -L${pkgs.zlib}/lib"
              export LD_LIBRARY_PATH=/tmp:${pkgs.zlib}/lib:$LD_LIBRARY_PATH
              export MACOSX_DEPLOYMENT_TARGET=14.5
            '';
          };
        };
      }
    );
}
