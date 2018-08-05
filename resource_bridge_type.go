package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-clnode/client"
	"strings"
	"fmt"
)

func resourceBridgeType() *schema.Resource {
	return &schema.Resource{
		Create: resourceBridgeTypeCreate,
		Read:   resourceBridgeTypeRead,
		Update: resourceBridgeTypeUpdate,
		Delete: resourceBridgeTypeDelete,

		Schema: map[string]*schema.Schema{
			"node_addr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBridgeTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.NodeClient)
	name := d.Get("name").(string)
	if name != strings.ToLower(name) {
		return fmt.Errorf("name must not contain any capitals")
	}
	err := c.CreateBridgeType(
		d.Get("node_addr").(string),
		name,
		d.Get("url").(string))
	if err != nil {
		return err
	}
	matcher := client.NewMatcher(d.Get("node_addr").(string), name)
	d.SetId(matcher.Id())
	return nil
}

func resourceBridgeTypeRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.NodeClient)
	bT, err := c.ReadBridgeType(d.Id())
	if err != nil {
		return err
	}
	d.Set("name", bT.Data.Attributes.Name)
	d.Set("url", bT.Data.Attributes.Url)
	return nil
}

func resourceBridgeTypeUpdate(d *schema.ResourceData, m interface{}) error {
	err := resourceBridgeTypeDelete(d, m)
	if err != nil {
		return err
	}
	err = resourceBridgeTypeCreate(d, m)
	if err != nil {
		return err
	}
	return nil
}

func resourceBridgeTypeDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.NodeClient)
	err := c.DeleteBridgeType(d.Id())
	if err != nil {
		return err
	}
	return nil
}