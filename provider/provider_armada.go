package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type dbaasCreds struct {
	access_key string
	secret_key string
}

func armadaProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: armadaSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"armada_ec2":      armada_ec2(),
			"armada_rds":      armada_rds(),
			"armada_dynamodb": armada_dynamodb(),
		},
		ConfigureFunc: providerConfiguration,
	}
}

func providerConfiguration(d *schema.ResourceData) (interface{}, error) {
	return &dbaasCreds{
		access_key: d.Get("access_key").(string),
		secret_key: d.Get("secret_key").(string),
	}, nil
}
