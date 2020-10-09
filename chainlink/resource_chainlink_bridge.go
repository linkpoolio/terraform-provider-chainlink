package chainlink

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
	"strings"
)

func ResourceChainlinkBridgeType() *schema.Resource {
	return &schema.Resource{
		Create: resourceBridgeTypeCreate,
		Read:   resourceBridgeTypeRead,
		Update: resourceBridgeTypeUpdate,
		Delete: resourceBridgeTypeDelete,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func resourceBridgeTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := NewClientFromModel(d, m)
	name := d.Get("name").(string)
	if name != strings.ToLower(name) {
		return fmt.Errorf("name must not contain any capitals")
	}
	err := c.CreateBridge(name, d.Get("url").(string))
	if err != nil {
		return err
	}
	matcher := client.NewMatcher("bridge", name)
	d.SetId(matcher.Id())
	return nil
}

func resourceBridgeTypeRead(d *schema.ResourceData, m interface{}) error {
	c := NewClientFromModel(d, m)
	bT, err := c.ReadBridge(d.Get("name").(string))
	if err == client.ErrNotFound {
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	if err := d.Set("name", bT.Data.Attributes.Name); err != nil {
		return err
	}
	if err := d.Set("url", bT.Data.Attributes.URL); err != nil {
		return err
	}
	return nil
}

func resourceBridgeTypeUpdate(d *schema.ResourceData, m interface{}) error {
	if err := resourceBridgeTypeDelete(d, m); err != nil {
		return err
	}
	return resourceBridgeTypeCreate(d, m)
}

func resourceBridgeTypeDelete(d *schema.ResourceData, m interface{}) error {
	return NewClientFromModel(d, m).DeleteBridge(d.Get("name").(string))
}