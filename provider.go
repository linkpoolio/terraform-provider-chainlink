package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/chainlink"
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
		ConfigureFunc: chainlink.ConfigureFunc,
		ResourcesMap: map[string]*schema.Resource{
			"chainlink_bridge":  chainlink.ResourceChainlinkBridgeType(),
			"chainlink_spec":    chainlink.ResourceChainlinkSpec(),
			"chainlink_spec_v2": chainlink.ResourceChainlinkSpecV2(),
			"chainlink_ocr_key": chainlink.ResourceChainlinkOCRKey(),
			"chainlink_p2p_key": chainlink.ResourceChainlinkP2PKey(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"chainlink_wallet": chainlink.DataSourceChainlinkWallet(),
		},
	}
}
