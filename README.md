Terraform Provider Armada
=========================

[![Build Status](https://dev.azure.com/rajrocksdeworld/Sample_GO/_apis/build/status/Sample_GO-Go%20(preview)-CI?branchName=master)](https://dev.azure.com/rajrocksdeworld/Sample_GO/_build/latest?definitionId=3&branchName=master)

- Website: https://www.terraform.io
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">
This is a custom plugin implementation of a Terraform Provider Armada, made using go.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

Developing the Provider
---------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](https://github.com/rajkumarbestha/terraform-provider-customplugin#requirements) before proceeding).

To build the plugin through the pipeline, you can use [azurepipeline.yml](https://github.com/rajkumarbestha/terraform-provider-armada/blob/master/azurebuildpipeline/azure-pipelines.yml) and download the Artifact (plugin) directly after the build succeeds.

Or if you wish to build it locally, clone the repo, install the dependencies and build.

Clone repository to: `somepath/terraform-providers/`

```
mkdir -p somepath/terraform-providers/; cd somepath/terraform-providers/
git clone https://github.com/rajkumarbestha/terraform-provider-armada.git
```
To compile the provider, enter the provider directory in the repo and run the below command. This will build the provider and put the provider binary in the current directory.

```
go build -o terraform-provider-armada.exe
```

Using the Provider
----------------------

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider (as above) in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) Or else for a quick workthrough, place the plugin(.exe) from the repo in the same directory as the terraform binary or place it in your current directory from where you are running the terraform scripts, and run `terraform init` to initialize it.

After initializing the custom plugin, you might want to write terraform scripts, these simple points will make you familiar with this plugin.

1. Provider Block:

```
provider "armada"{
   // AK and SK.
}
```

2. Resource Block:

```
resource "armada_ec2" "dev_ec2__test"{
   // required fields.
}
```

To know about the resources and the fields this plugin supports, please have a look at this [sample terraform scripts](https://github.com/rajkumarbestha/terraform-provider-armada/tree/master/examples). These scripts will make you familiar with the plugin easily.

Happy Terraforming! :)

Contributing
---------------------------

Terraform is the work of thousands of contributors. We appreciate your help!

To contribute, please reach out to Rasheed/Rajkumar.



