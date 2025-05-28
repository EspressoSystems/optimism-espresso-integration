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

### Install Espresso go library

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

## Docker

In order to download the docker images required by this project you may need to authenticate using a PAT.

Create a [Github Personal Access Token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic) following Creating a personal access token (classic).

Provide Docker with the PAT.

```
> export CR_PAT=<your PAT>
> echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
```

### Run the tests

To run all the tests (slow):

> just tests


To run a subset of the tests (fast):

> just fast-tests


Run the Espresso smoke tests:

> just smoke-tests


Run the Espresso integration tests:

> just espresso-tests


If some containers are still running (due to failed tests) run this command to stop and delete all the Espresso containers:

> just remove-containers


If in the Nix environment, any `just` command fails with a tool version mismatch error such as
`version "go1.22.7" does not match go tool version "go1.22.12"`, use
`export GOROOT="$(dirname $(dirname $(which go)))/share/go"` to set the expected Go version.

### Run the Kurtosis devnet

- Install tools.
  - Install Docker Desktop via https://www.docker.com/products/docker-desktop/.
    - Or podman, colima, etc.
    - Verify Docker is installed:
      ```bash
      docker version
      ```

  - Install Kurtosis via https://docs.kurtosis.com/install/.

- Run the devnet.
  - In the Nix environment:
    ```bash
    cd kurtosis-devnet
    just espresso-devnet
    ```

  - If you get the `command not found` or the `"kurtosis": executable file not found in $PATH`
  error, add the Docker's binary directory to `PATH`. E.g., if the Docker CLI lives at
  `/Applications/Docker.app/Contents/Resources/bin/`, run:
    ```bash
    echo 'export PATH="/Applications/Docker.app/Contents/Resources/bin:$PATH"' >> ~/.bash_profile
    source ~/.bash_profile
    ```
    or:
    ```bash
    echo 'export PATH="/Applications/Docker.app/Contents/Resources/bin:$PATH"' >> ~/.zshrc
    source ~/.zshrc
    ```
    if you are using Zsh. Then restart the devnet test.


  - Kurtosis devnet can be quite slow to start, especially on the first run. Verify everything is
  running with:
    ```bash
    kurtosis enclave inspect espresso-devnet
    ```

  - Read logs:
    ```bash
    kurtosis service logs espresso-devnet <service-name>

    # show all the logs
    kurtosis service logs -a espresso-devnet <service-name>

    # frequently used commands
    kurtosis service logs -a espresso-devnet op-batcher-op-kurtosis
    kurtosis service logs -a espresso-devnet op-cl-1-op-node-op-geth-op-kurtosis
    ```

  - Clean up:
    ```bash
    kurtosis clean -a
    ```


### CI environment

We currently use Circle CI but this is temporary. In order to run the go linter do:
```
just golint
```

### Guide: Setting Up an Enclave-Enabled Nitro EC2 Instance

This guide explains how to prepare an enclave-enabled parent EC2 instance.

You can follow the official AWS Enclaves setup guide: https://docs.aws.amazon.com/enclaves/latest/user/getting-started.html.


#### Step-by-Step Instructions

##### 1. Launch the EC2 Instance

Use the AWS Management Console or AWS CLI to launch a new EC2 instance.

Make sure to:

- **Enable Enclaves**
  - In the CLI: set the `--enclave-options` flag to `true`
  - In the Console: select `Enabled` under the **Enclave** section

- **Use the following configuration:**
  - **Architecture:** x86_64
  - **AMI:** Amazon Linux 2023
  - **Instance Type:** `m6a.2xlarge`
  - **Volume Size:** 100 GB


##### 2. Connect to the Instance

Once the instance is running, connect to it via the AWS Console or CLI.
In practice, you will be provided a key.pem file and you can connect like this:
```shell
ssh -i "key.pem" ec2-user@<aws_instance_dns>
```

Note that the command above can be found in the AWS by selecting the instance and clicking on the button "Connect".


##### 3. Install dependencies

* Nix
`sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon`

* Git, Nitro
```
 sudo yum update
 sudo yum install git
 sudo dnf install aws-nitro-enclaves-cli -y
```

* Clone repository and update submodules
```
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git submodule update --init --recursive
```

* Configure enclave
```
sudo mkdir /etc/nitro_enclaves
touch /etc/nitro_enclaves/allocator.yaml
```

In the file `/etc/nitro_enclaves/allocator.yaml` put the following content:
```
---
# Enclave configuration file.
#
# How much memory to allocate for enclaves (in MiB).
memory_mib: 4096
#
# How many CPUs to reserve for enclaves.
cpu_count: 2
#
# Alternatively, the exact CPUs to be reserved for the enclave can be explicitly
# configured by using `cpu_pool` (like below), instead of `cpu_count`.
# Note: cpu_count and cpu_pool conflict with each other. Only use exactly one of them.
# Example of reserving CPUs 2, 3, and 6 through 9:
# cpu_pool: 2,3,6-9
[ec2-user@ip-172-31
```

Restart the enclave service
```
sudo systemctl start nitro-enclaves-allocator.service
```

* Enter the nix shell and run the enclave tests
```
cd optimism-espresso-integration
nix --extra-experimental-features "nix-command flakes" develop
just espresso-enclave-tests
```






