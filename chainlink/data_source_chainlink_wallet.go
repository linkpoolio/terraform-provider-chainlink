package chainlink

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func DataSourceChainlinkWallet() *schema.Resource {
	return &schema.Resource{
		Read: resourceDataWalletRead,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDataWalletRead(d *schema.ResourceData, m interface{}) error {
	if addr, err := m.(*client.Chainlink).ReadWallet(); err != nil {
		return err
	} else if err := d.Set("address", addr); err != nil {
		return err
	} else {
		m := client.NewMatcher("wallet", addr)
		d.SetId(m.Id())
	}
	return nil
}
