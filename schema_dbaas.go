package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dbaasSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_key": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"secret_key": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
