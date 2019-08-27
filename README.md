Terraform Provider Armada
=========================
This is a custom plugin implementation of a Terraform Provider Armada made using go.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Developing the Provider
---------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](https://github.com/rajkumarbestha/terraform-provider-customplugin#requirements) before proceeding).

Clone repository to: `somepath/development/terraform-providers/`

```
mkdir -p somepath/development/terraform-providers/; cd somepath/development/terraform-providers/
git clone https://github.com/rajkumarbestha/terraform-provider-customplugin.git
```
To compile the provider, enter the provider directory in the repo and run the below command. This will build the provider and put the provider binary in the current directory.

```
go build -o terraform-provider-armada.exe
```

Using the Provider
----------------------

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider (as above) in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) Or else for a quick workthrough, place the plugin(.exe) from the repo in the same directory as the terraform binary or place it in your current directory from where you are running the terraform scripts, and run `terraform init` to initialize it.

For either installation method, documentation about the provider specific configuration options can be found on the [provider's website](https://www.terraform.io/docs/providers/aws/index.html).



