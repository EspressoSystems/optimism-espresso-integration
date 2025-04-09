# Optimism Espresso Integration

## Development environment

### Clone the repository and initialize the submodules

```
> git clone git@github.com:EspressoSystems/optimism-espresso-integration.git
> git submodule update --init --recursive
```

### Nix shell

* Install nix following the instructions at https://nixos.org/download/

* Enter the nix shell of this project

> nix develop .

### Mises

*  Install Mises

Follow the instructions for your own OS: https://mise.jdx.dev/getting-started.html
When executing the script some instruction will be printed in order to activate mises, for example in Ubuntu you will get a message like:
```
######################################################################## 100,0%
mise: installed successfully to /home/leloup/.local/bin/mise
mise: run the following to activate mise in your shell:
echo "eval \"\$(/home/leloup/.local/bin/mise activate bash)\"" >> ~/.bashrc
```

In this case you should run
```
echo "eval \"\$(/home/leloup/.local/bin/mise activate bash)\"" >> ~/.bashrc
```

And then open a new terminal or type:
```
> source ~/.bashrc
```

Finally, install all the dependencies:

```
> mise install
```

### Install Espresso go cryptographic library

This step is only needed if you use Mises as Nix automatically installs the Espresso go cryptographic library.

- Create a local directory for later use. Note it has to be created under the home directory by default.

  ```bash
  cd ..
  mkdir -p ~/local-lib
  ```

- Get `libespresso_crypto_helper.a`, by either method below.
  - Get it from the CI. See https://github.com/EspressoSystems/espresso-network-go/releases
    - Download `libespresso_crypto_helper-x86_64-apple-darwin.a` (or the one for linux).
    - Move the downloaded file to `local-lib`.

  - Build it locally.
    - Download the sequencer Go code via `https://github.com/EspressoSystems/espresso-network-go/archive/refs/tags/v0.0.34.tar.gz`.
      - Replace the version number if there’s a newer one.
    - Go to the downloaded folder.

      ```bash
      cd espresso-network-go-0.0.34
      ```

    - Build the verification code.
      - Make sure to not run this in the nix shell. Otherwise, the generated file will be in the wrong directory, which will cause `just` to fail later.
      - This may require `rustup update` if the Rust version is older than expected.

      ```bash
      cargo build --release --locked --manifest-path ./verification/rust/Cargo.toml
      ```

    - Copy the `libespresso_crypto_helper.a` file.
      - Linux:

        ```bash
        sudo cp ./espresso-network-go-0.0.34/verification/rust/target/release/libespresso_crypto_helper.a ~/local-lib/libespresso_crypto_helper-x86_64-unknown-linux-gnu.a
        ```

      - Mac:

        ```bash
        sudo cp ./espresso-network-go-0.0.34/verification/rust/target/release/libespresso_crypto_helper.a ~/local-lib/libespresso_crypto_helper-x86_64-apple-darwin.a
        ```

- Set the flag.
  - Linux:

      ```bash
      export CGO_LDFLAGS="-L$HOME/local-lib -lespresso_crypto_helper-x86_64-unknown-linux-gnu"
      ```

  - Mac:

      ```bash
      export CGO_LDFLAGS="-L$HOME/local-lib -lespresso_crypto_helper-x86_64-apple-darwin -framework Foundation -framework SystemConfiguration"
      ```

### Run the tests

To run all the tests (slow):

> just tests


To run a subset of the tests (fast):

> just fast-tests


Run the Espresso integration tests:

> just espresso-tests


If in the Nix environment, any `just` command fails with a tool version mismatch error such as
`version "go1.22.7" does not match go tool version "go1.22.12"`, use
`export GOROOT="$(dirname $(dirname $(which go)))/share/go"` to set the expected Go version.