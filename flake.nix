{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";

    # Pinned to commit that has Go 1.23.8
    pkgs-go.url = "github:nixos/nixpkgs/ebe4301";
    # Foundry 1.2.3 — pinned for the Solidity compiler (forge)
    pkgs-foundry.url = "github:nixos/nixpkgs/648f701";
    # Foundry 1.5.1 — newer anvil with built-in Beacon REST API
    pkgs-anvil.url = "github:nixos/nixpkgs/6201e203d09599479a3b3450ed24fa81537ebc4e";
  };

  outputs =
    inputs:
    inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import inputs.nixpkgs { inherit system; };

        go_1_23_8 = (import inputs.pkgs-go { inherit system; }).go_1_23;
        foundry_1_2_3 = (import inputs.pkgs-foundry { inherit system; }).foundry;

        # Newer Foundry (1.5.1) for anvil only — includes built-in Beacon REST API
        # needed by the devnet L1.  We extract just the anvil binary so it doesn't
        # shadow forge/cast from the pinned 1.2.3 package.
        anvil_1_5_1 = pkgs.runCommand "anvil-1.5.1" { } ''
          mkdir -p $out/bin
          ln -s ${(import inputs.pkgs-anvil { inherit system; }).foundry}/bin/anvil $out/bin/anvil
        '';

        espressoGoLibFile = pkgs.stdenv.mkDerivation rec {
          pname = "libespresso_crypto_helper";
          version = "0.2.1";

          baseUrl = "https://github.com/EspressoSystems/espresso-network/releases/download/sdks%2Fgo%2Fv${version}";
          source =
            {
              "x86_64-linux" = pkgs.fetchurl {
                url = baseUrl + "/libespresso_crypto_helper-x86_64-unknown-linux-gnu.so";
                sha256 = "sha256:b3e28f7dc755d72b27a2a43c2bcfdc0e4e82096e03596a01447bd8f406e6653c";
              };
              "x86_64-darwin" = pkgs.fetchurl {
                url = baseUrl + "/libespresso_crypto_helper-x86_64-apple-darwin.dylib";
                sha256 = "sha256:716cb9eb548222ed1c7b5d1585bd5f03d0680cbae3f8db14cbf37837f54b9788";
              };
              "aarch64-linux" = pkgs.fetchurl {
                url = baseUrl + "/libespresso_crypto_helper-aarch64-unknown-linux-gnu.so";
                sha256 = "sha256:886aef8aeaa0d5695abc6a9ae54f085899a031371c10755218e387442ecb331f";
              };
              "aarch64-darwin" = pkgs.fetchurl {
                url = baseUrl + "/libespresso_crypto_helper-aarch64-apple-darwin.dylib";
                sha256 = "sha256:6c74ec631ccd9d23258ff99a8060068a548740fac814633ceab2ad7c7dc90a74";
              };
            }
            ."${system}";

          dontUnpack = true;
          installPhase = ''
            mkdir -p $out/lib
            cp ${source} $out/lib/
          '';
        };

        eth-beacon-genesis = pkgs.buildGoModule rec {
          pname = "eth-beacon-genesis";
          version = "703e97a";

          src = pkgs.fetchFromGitHub {
            owner = "ethpandaops";
            repo = pname;
            rev = version;
            hash = "sha256-Toal70A8cnIAtR4iCacRQ5vT+MHUlMc81l1dzjj56mA=";
          };

          vendorHash = "sha256-keBJHjl42o6guAAAWoESJateXVG3hotdSnDv2pf1Lv4=";
          proxyVendor = true;

          doCheck = false;
        };

        eth2-val-tools = pkgs.buildGoModule rec {
          pname = "eth2-val-tools";
          version = "662955e";

          src = pkgs.fetchFromGitHub {
            owner = "protolambda";
            repo = pname;
            rev = version;
            hash = "sha256-UpQmCS/FrY667EnNH2XCTJhzhLOpsfS5GUhGvXGG65U=";
          };

          vendorHash = "sha256-IblAuZgk7EBkfcFoEugzb9pO454zfHq6RxIfgvUFBDo=";
          proxyVendor = true;

          doCheck = false;
        };

        # Fetch the pre-built enclaver binary from GitHub releases.
        enclaver =
          let
            version = "0.6.1";
            platformMap = {
              "x86_64-linux"   = { platform = "linux-x86_64";  sha256 = "sha256-lvKcruw4U+cJUAEOp/LeGiLJX7+SHngL6K3iFgflNy8="; };
              "aarch64-linux"  = { platform = "linux-aarch64"; sha256 = "sha256-qsUFqYQ445ngQtC+ncG8i/9LAr7uPrFJ5TAa9qKJcw8="; };
              "x86_64-darwin"  = { platform = "macos-x86_64";  sha256 = "sha256-Ru4s+mfy+jaerqDO2c2wQFjf5Ph+aZJa0+Y9YG1K3Sk="; };
              "aarch64-darwin" = { platform = "macos-aarch64"; sha256 = "sha256-LeLg9oxe5LTGXCLMDuaY3MujC4cu15oycd8hPLjcp+I="; };
            };
            platformInfo = platformMap."${system}";
          in
          pkgs.stdenv.mkDerivation {
            pname = "enclaver";
            inherit version;

            src = pkgs.fetchurl {
              url = "https://github.com/enclaver-io/enclaver/releases/download/v${version}/enclaver-${platformInfo.platform}-v${version}.tar.gz";
              sha256 = platformInfo.sha256;
            };

            sourceRoot = ".";

            installPhase = ''
              mkdir -p $out/bin
              cp enclaver-*/enclaver $out/bin/enclaver
              chmod +x $out/bin/enclaver
            '';
          };

      in
      let
        # Exclude enclaver because buildRustPackage requires fetching Rust crates from crates.io at
        # build time, which CI is routinely blocked from. Workflows that need enclaver use
        # devShells.enclave instead.
        commonBuildInputs = [
          pkgs.zlib
          espressoGoLibFile
        ];

        commonPackages = [
          go_1_23_8
          foundry_1_2_3
          anvil_1_5_1

          eth-beacon-genesis
          eth2-val-tools

          pkgs.awscli2
          pkgs.cargo
          (import inputs.pkgs-go { inherit system; }).dasel
          pkgs.go-ethereum
          pkgs.jq
          pkgs.just
          pkgs.pnpm
          pkgs.python311
          pkgs.shellcheck
          pkgs.uv
          pkgs.yq-go
          pkgs.tmux
          pkgs.golangci-lint
        ];

        commonShellHook = ''
          export MACOSX_DEPLOYMENT_TARGET=14.5
          export PATH=$PATH:$PWD/op-deployer/bin
        '';
      in
      {
        formatter = pkgs.nixfmt-rfc-style;

        devShells = {
          # Default shell.
          default = pkgs.mkShell {
            buildInputs = commonBuildInputs;
            packages = commonPackages;
            shellHook = commonShellHook;
          };

          # Enclave shell. Use with:
          #   nix develop .#enclave --command <cmd>
          enclave = pkgs.mkShell {
            buildInputs = commonBuildInputs;
            packages = commonPackages ++ [ enclaver ];
            shellHook = commonShellHook;
          };
        };
      }
    );
}
