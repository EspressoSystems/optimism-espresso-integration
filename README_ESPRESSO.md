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

*

### Run the tests

To run the tests:

> just tests
