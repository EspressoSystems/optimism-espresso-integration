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

> nix develop


## Docker

In order to download the docker images required by this project you may need to authenticate using a PAT.

Create a [Github Personal Access Token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic) following Creating a personal access token (classic).

Provide Docker with the PAT.

```
> export CR_PAT=<your PAT>
> echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
```

### Run the tests

Run the Espresso smoke tests:

> just smoke-tests


Run the Espresso integration tests. Note, this can take up to 30min.

> just espresso-tests


To run all the standard OP stack (w/o Espresso integration) tests (slow):

> just tests

To run a subset of the tests above (fast):

> just fast-tests



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


### Misc commands

In order to run the go linter do:
```
just golint
```

Generate the bindings for the contracts:
```
just gen-bindings
```

If some containers are still running (due to failed tests) run this command to stop and delete all the Espresso containers:

> just remove-containers


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
In practice, you will be provided a `key.pem` file, and you can connect like this:
```shell
chmod 400 key.pem
ssh -i "key.pem" ec2-user@<aws_instance_dns>
```

Note that the command above can be found in the AWS Console by selecting the instance and clicking on the button "Connect".


##### 3. Install dependencies

* Nix
```
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon`
source ~/.bashrc
```

* Git, Nitro, Docker
```
 sudo yum update
 sudo yum install git
 sudo usermod -a -G docker ec2-user
 sudo chown ec2-user /var/run/docker.sock
 sudo service docker start
 sudo dnf install aws-nitro-enclaves-cli -y
 sudo systemctl start nitro-enclaves-allocator.service
```

* Clone repository and update submodules
```
git clone https://github.com/EspressoSystems/optimism-espresso-integration.git
cd optimism-espresso-integration
git submodule update --init --recursive
```


* Enter the nix shell and run the enclave tests
```
cd optimism-espresso-integration
nix --extra-experimental-features "nix-command flakes" develop
just espresso-enclave-tests
```
