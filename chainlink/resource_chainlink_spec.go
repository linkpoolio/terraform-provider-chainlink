package chainlink

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func ResourceChainlinkSpec() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpecCreate,
		Read:   resourceSpecRead,
		Delete: resourceSpecDelete,
		Update: resourceSpecUpdate,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"spec_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"json": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		}),
	}
}

func resourceSpecCreate(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	json := d.Get("json").(string)

	id, err := c.CreateSpec(json)
	if err != nil {
		return err
	}

	d.SetId(id)
	if err := d.Set("spec_id", id); err != nil {
		return err
	}
	return nil
}

func resourceSpecRead(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	spec, err := c.ReadSpec(d.Get("spec_id").(string))

	if err == client.ErrNotFound {
		d.SetId("")
	} else if err != nil {
		d.SetId("")
		return err
	} else if err := d.Set("spec_id", spec.Data["id"]); err != nil {
		return err
	}
	return nil
}

func resourceSpecUpdate(d *schema.ResourceData, m interface{}) error {
	if err := resourceSpecDelete(d, m); err != nil {
		return err
	}
	return resourceSpecCreate(d, m)
}

func resourceSpecDelete(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	return c.DeleteSpec(d.Get("spec_id").(string))
}
