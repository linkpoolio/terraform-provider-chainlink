package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func resourceSpec() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpecCreate,
		Read:   resourceSpecRead,
		Delete: resourceSpecDelete,

		Schema: map[string]*schema.Schema{
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
	return nil
}

func resourceSpecRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Chainlink)
	spec, err := c.ReadSpec(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	if err := d.Set("id", spec.Data["id"]); err != nil {
		return err
	}
	return nil
}

func resourceSpecDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}