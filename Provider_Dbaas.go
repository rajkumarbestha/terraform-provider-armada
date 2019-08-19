package main

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/schema"
)

type dbaasCreds struct {
	access_key string
	secret_key string
}

func dbaasProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:      	dbaasSchema(),
		ResourcesMap: map[string]*schema.Resource{
                        "dbaas_EC2": dbaasEC2(),
    },
		ConfigureFunc: providerConfiguration,
	}
}

func providerConfiguration(d *schema.ResourceData) (interface{}, error){
	return &dbaasCreds{
		access_key: d.Get("access_key").(string),
		secret_key: d.Get("secret_key").(string),
	}, nil
}
