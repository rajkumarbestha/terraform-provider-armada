package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rajkumarbestha/terraform-provider-customplugin/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}