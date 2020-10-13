package chainlink

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func NewClient(email, password, url string) (interface{}, error) {
	return client.NewChainlink(&client.Config{
		Email:    fmt.Sprint(email),
		Password: fmt.Sprint(password),
		URL:      fmt.Sprint(url),
	})
}

func ConfigureFunc(d *schema.ResourceData) (interface{}, error) {
	return NewClient(
		d.Get("email").(string),
		d.Get("password").(string),
		d.Get("url").(string),
	);
}

func NewClientFromModel(d *schema.ResourceData, m interface{}) (*client.Chainlink, error) {
	email := d.Get("chainlink_email").(string)
	password := d.Get("chainlink_password").(string)
	url := d.Get("chainlink_url").(string)

	if len(email) > 0 && len(password) > 0 && len(url) > 0 {
		obj, err := NewClient(
			d.Get("chainlink_email").(string),
			d.Get("chainlink_password").(string),
			d.Get("chainlink_url").(string),
		)
		if err != nil {
			return nil, err
		}
		return obj.(*client.Chainlink), err
	}
	return m.(*client.Chainlink), nil
}

func mergeSchemaWithNodeProperties(schemaMap map[string]*schema.Schema) map[string]*schema.Schema {
	schemaMap["chainlink_url"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	}
	schemaMap["chainlink_email"] = &schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		ForceNew:  false,
		Sensitive: true,
	}
	schemaMap["chainlink_password"] = &schema.Schema{
		Type:      schema.TypeString,
		Optional:  true,
		ForceNew:  false,
		Sensitive: true,
	}
	return schemaMap
}
