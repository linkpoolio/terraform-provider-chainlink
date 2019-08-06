package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:6688",
				Description: "The node url address",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "admin@node.local",
				Description: "Node email address",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "twochains",
				Description: "Node password",
			},
		},
		ConfigureFunc: newClient,
		ResourcesMap: map[string]*schema.Resource{
			"chainlink_bridge": resourceBridgeType(),
			"chainlink_spec":   resourceSpec(),
		},
	}
}

func newClient(d *schema.ResourceData) (interface{}, error) {
	return client.NewChainlink(&client.Config{
		Email:    d.Get("email").(string),
		Password: d.Get("password").(string),
		URL:      d.Get("url").(string),
	})
}
