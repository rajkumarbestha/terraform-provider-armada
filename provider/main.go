package main

import (
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: armadaProvider,
	}
	plugin.Serve(&opts)
}
