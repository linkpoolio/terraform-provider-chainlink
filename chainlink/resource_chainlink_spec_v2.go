package chainlink

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func ResourceChainlinkSpecV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpecV2Create,
		Read:   resourceSpecV2Read,
		Delete: resourceSpecV2Delete,
		Update: resourceSpecV2Update,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"toml": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		}),
	}
}

func resourceSpecV2Create(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	toml := d.Get("toml").(string)
	spec, err := c.CreateSpecV2(toml)
	if err != nil {
		return err
	}
	d.SetId(spec.Data.ID)
	return nil
}

func resourceSpecV2Read(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	if err := c.ReadSpecV2(d.Id()); err == client.ErrNotFound || err == client.ErrUnprocessableEntity {
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func resourceSpecV2Update(d *schema.ResourceData, m interface{}) error {
	if err := resourceSpecV2Delete(d, m); err != nil {
		return err
	}
	return resourceSpecV2Create(d, m)
}

func resourceSpecV2Delete(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	return c.DeleteSpecV2(d.Id())
}
