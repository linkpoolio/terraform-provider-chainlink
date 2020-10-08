package chainlink

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linkpoolio/terraform-provider-chainlink/client"
)

func NewClient(d *schema.ResourceData) (interface{}, error) {
	return client.NewChainlink(&client.Config{
		Email:    fmt.Sprint(d.Get("email")),
		Password: fmt.Sprint(d.Get("password")),
		URL:      fmt.Sprint(d.Get("url")),
	})
}

func NewClientFromModel(d *schema.ResourceData, m interface{}) *client.Chainlink {
	if obj, err := NewClient(d); err == nil {
		return obj.(*client.Chainlink)
	} else {
		return m.(*client.Chainlink)
	}
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
