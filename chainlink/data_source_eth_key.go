package chainlink

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func DataSourceETHKey() *schema.Resource {
	return &schema.Resource{
		Read: resourceETHKeyRead,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		}),
	}
}

func resourceETHKeyRead(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	keys, err := c.ReadETHKeys()
	if err != nil {
		return err
	}

	index, ok := d.Get("index").(int)
	if !ok {
		return fmt.Errorf("provided index of %s is not an integer", fmt.Sprint(index))
	}
	if index >= len(keys.Data) {
		return fmt.Errorf("provided index of %d exceeds %d amount of keys returned", index, len(keys.Data))
	}

	d.SetId(keys.Data[index].Attributes.Address)
	return nil
}
