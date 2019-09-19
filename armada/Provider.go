package armada

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: armadaSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"armada_ec2":      armadaEc2(),
			"armada_rds":      armadaRds(),
			"armada_dynamodb": armadaDynamodb(),
		},
		ConfigureFunc: providerConfiguration,
	}
}

func providerConfiguration(d *schema.ResourceData) (interface{}, error) {
	return &armadaCreds{
		access_key: d.Get("access_key").(string),
		secret_key: d.Get("secret_key").(string),
		endpoint: d.Get("endpoint").(string),
	}, nil
}
