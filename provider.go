package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-clnode/client"
	"fmt"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "chainlink",
				Description: "ChainLink Node API Username",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "twochains",
				Description: "ChainLink Node API Password",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http",
				Description: "ChainLink Node API Protocol (http or https)",
			},
		},
		ConfigureFunc: newClient,
		ResourcesMap: map[string]*schema.Resource{
			"clnode_bridge_type": resourceBridgeType(),
		},
	}
}

func newClient(d *schema.ResourceData) (interface{}, error) {
	protocol := d.Get("protocol").(string)
	if protocol != "http" && protocol != "https" {
		return nil, fmt.Errorf("protocol not supported, use http or https")
	}
	return client.NewNodeClient(client.Config{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Protocol: protocol,
	}), nil
}