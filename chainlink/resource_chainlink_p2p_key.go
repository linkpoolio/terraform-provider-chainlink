package chainlink

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func ResourceChainlinkP2PKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceP2PKeyCreate,
		Read:   resourceP2PKeyRead,
		Delete: resourceP2PKeyDelete,
		Update: resourceP2PKeyUpdate,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"peer_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		}),
	}
}

func resourceP2PKeyCreate(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	key, err := c.CreateP2PKey()
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(key.Data.Attributes.ID))
	if err := d.Set("peer_id", key.Data.Attributes.PeerID); err != nil {
		return err
	} else if err := d.Set("public_key", key.Data.Attributes.PublicKey); err != nil {
		return err
	}
	return nil
}

func resourceP2PKeyRead(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	keys, err := c.ReadP2PKeys()
	if err != nil {
		return err
	}

	id := d.Id()
	found := false
	for _, key := range keys.Data {
		if fmt.Sprint(key.Attributes.ID) == id {
			found = true
			break
		}
	}
	if !found {
		d.SetId("")
	}
	return nil
}

func resourceP2PKeyUpdate(d *schema.ResourceData, m interface{}) error {
	if err := resourceP2PKeyDelete(d, m); err != nil {
		return err
	}
	return resourceP2PKeyCreate(d, m)
}

func resourceP2PKeyDelete(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	return c.DeleteP2PKey(int(id))
}
