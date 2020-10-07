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

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceSpecCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Chainlink)
	json := d.Get("json").(string)
	id, err := c.CreateSpec(json)
	if err != nil {
		return err
	}
	matcher := client.NewMatcher("spec", id)
	d.SetId(matcher.Id())
	if err := d.Set("spec_id", id); err != nil {
		return err
	}
	return nil
}

func resourceSpecRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Chainlink)
	spec, err := c.ReadSpec(d.Get("spec_id").(string))
	if err != nil {
		d.SetId("")
		return err
	}
	if err := d.Set("spec_id", spec.Data["id"]); err != nil {
		return err
	}
	return nil
}

func resourceSpecDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Chainlink)
	return c.DestroySpec(d.Get("spec_id").(string))
}