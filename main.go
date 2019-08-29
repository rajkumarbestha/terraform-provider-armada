package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rajkumarbestha/terraform-provider-armada/armada"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: armada.Provider,
	}
	plugin.Serve(&opts)
}
