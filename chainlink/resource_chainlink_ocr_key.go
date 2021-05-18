package chainlink

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceChainlinkOCRKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceOCRKeyCreate,
		Read:   resourceOCRKeyRead,
		Delete: resourceOCRKeyDelete,
		Update: resourceOCRKeyUpdate,

		Schema: mergeSchemaWithNodeProperties(map[string]*schema.Schema{
			"config_public_key": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"offchain_public_key": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"onchain_signing_address": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		}),
	}
}

func resourceOCRKeyCreate(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	key, err := c.CreateOCRKey()
	if err != nil {
		return err
	}

	d.SetId(key.Data.ID)
	if err := d.Set("config_public_key", key.Data.Attributes.ConfigPublicKey); err != nil {
		return err
	} else if err := d.Set("offchain_public_key", key.Data.Attributes.OffChainPublicKey); err != nil {
		return err
	} else if err := d.Set("onchain_signing_address", key.Data.Attributes.OnChainSigningAddress); err != nil {
		return err
	}
	return nil
}

func resourceOCRKeyRead(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}

	keys, err := c.ReadOCRKeys()
	if err != nil {
		return err
	}

	id := d.Id()
	found := false
	for _, key := range keys.Data {
		if key.ID == id {
			found = true
			break
		}
	}
	if !found {
		d.SetId("")
	}
	return nil
}

func resourceOCRKeyUpdate(d *schema.ResourceData, m interface{}) error {
	if err := resourceOCRKeyDelete(d, m); err != nil {
		return err
	}
	return resourceOCRKeyCreate(d, m)
}

func resourceOCRKeyDelete(d *schema.ResourceData, m interface{}) error {
	c, err := NewClientFromModel(d, m)
	if err != nil {
		return err
	}
	return c.DeleteOCRKey(d.Id())
}
