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
sh <(curl --proto '=https' --tlsv1.2 -L https://nixos.org/nix/install) --daemon
source ~/.bashrc
```

* Git, Docker
```
 sudo yum update
 sudo yum install git
 sudo yum install docker
 sudo usermod -a -G docker ec2-user
 sudo service docker start
 sudo chown ec2-user /var/run/docker.sock
```

* Nitro

These commands install the dependencies for, start the service related to and configures the enclave.

```
sudo yum install -y aws-nitro-enclaves-cli-1.4.2
sudo systemctl stop nitro-enclaves-allocator.service || true
echo -e '---\nmemory_mib: 4096\ncpu_count: 2' | sudo tee /etc/nitro_enclaves/allocator.yaml
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
nix --extra-experimental-features "nix-command flakes" develop
just compile-contracts
just espresso-enclave-tests
```

#### Building, running and registering enclave images

`op-batcher/enclave-tools` provides a command-line utility for common operations on batcher enclave images.
Before using it, set your AWS instance as described in the guide above, then build the tool:

```
cd op-batcher/
just enclave-tools
```

This should create `op-batcher/bin/enclave-tools` binary. You can run
```
./op-batcher/bin/enclave-tools --help
```
to get information on available commands and flags.

##### Building a batcher image

To build a batcher enclave image, and tag it with specified tag:
```
./op-batcher/bin/enclave-tools build --op-root ./ --tag op-batcher-enclave
```
On success this command will output PCR measurements of the enclave image, which can then be registered with BatchAuthenticator
contract.

##### Running a batcher image
To run enclave image built by the previous command:
```
./op-batcher/bin/enclave-tools run --image op-batcher-enclave --args --argument-1,value-1,--argument-2,value-2
```
Arguments will be forwarded to the op-batcher

##### Registering a batcher image
To register PCR0 of the batcher enclave image built by the previous command:
```
./op-batcher/bin/enclave-tools register --l1-url example.com:1234 --authenticator 0x123..def --private-key 0x123..def --pcr0 0x123..def
```
You will need to provide the L1 URL, the contract address of BatchAuthenticator, private key of L1 account used to deploy BatchAuthenticator and PCR0 obtained when building the image.

## Docker Compose

### Run Docker Compose

* Ensure that your Docker Compose, Engine, and plugins are up-to-date. Particularly, if the Docker
Compose version is `2.37.3` or the Docker Engine version is `27.4.0`, and the Docker build hangs,
you may need to upgrade the version.

* Go to the `espresso` directory.
```
cd espresso
```

* Copy the example environment setting.
```
cp .env.example .env
```

* Shut down all containers.
```
docker compose down
```

* Build and start all services in the background.
```
docker compose up --build -d
```

* Run the services and check the log.
```
docker compose logs -f
```

### Investigate a Service

* Shut down all containers.
```
docker compose down
```

* Build and start the specific service and check the log.
```
docker compose up <service-name>
```

* If the environment variable setting is not picked up, pass it explicitly.
```
docker compose --env-file .env up <service-name>
```

* If there is a timing synchronization issue, update the `l2_time` field in `rollup-devnet.json`
with the current timestamp, convert the time to hex and update the `timestamp` fields in the two
genesis files, `l1-genesis-devnet.json` and `l2-genesis-devnet.json`, too.

### Apply a Change

* In most cases, simply remove all containers and run commands as normal.
```
docker compose down
```

* To start the project fresh, remove containers, volumes, and network, from this project.
```
docker compose down -v
```

* To start the system fresh, remove all volumes.
```
docker volume prune -a
```

* If a genesis file is updated, you may get a hash mismatch error when running a service that uses
the genesis file. Replace the corresponding `hash` field in `rollup-devnet.json`, then rerun the
failed command.

## Continuous Integration environment

### Running enclave tests in EC2

In order to run the tests for the enclave in EC2 via github actions one must create an AWS user that supports the following policy:

```json
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"ec2:AuthorizeSecurityGroupIngress",
				"ec2:RunInstances",
				"ec2:DescribeInstances",
				"ec2:TerminateInstances",
				"ec2:DescribeImages",
				"ec2:CreateTags",
				"ec2:DescribeSecurityGroups",
				"ec2:DescribeKeyPairs",
				"ec2:ImportKeyPair",
				"ec2:DescribeInstanceStatus"
			],
			"Resource": "*"
		}
	]
}
```

Currently, the github workflow in `.github/workflows/enclave.yaml` relies on a custom AWS AMI with id `ami-0ff5662328e9bbc2f`.
In order to refresh this AMI one needs to:
1. Create an AWS EC2 instance with the characteristics described in (see `.github/workflows/enclave.yaml` *Launch EC2 Instance* job).
2. Copy the script `espresso/scrips/enclave-prepare-ami.sh` in the EC2 instance (e.g. using scp) and run it.
3. [Export the AMI instance](https://docs.aws.amazon.com/toolkit-for-visual-studio/latest/user-guide/tkv-create-ami-from-instance.html).
