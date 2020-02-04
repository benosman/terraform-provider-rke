package rke

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

//Schemas

func rkeClusterBastionHostFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"address": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Address of Bastion Host",
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "SSH User to Bastion Host",
		},
		"port": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "SSH Port of Bastion Host",
			ValidateFunc: validation.IntBetween(1, 65535),
			Computed:     true,
		},
		"ssh_agent_auth": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "SSH Agent Auth enable",
		},
		"ssh_cert": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Optional:    true,
			Description: "SSH Certificate Key",
			Computed:    true,
		},
		"ssh_cert_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SSH Certificate Key Path",
			Computed:    true,
		},
		"ssh_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "SSH Private Key",
			Computed:    true,
		},
		"ssh_key_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SSH Private Key Path",
			Computed:    true,
		},
	}
	return s
}
