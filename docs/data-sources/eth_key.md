---
page_title: "chainlink_wallet Data Source - terraform-provider-chainlink"
subcategory: ""
description: |-
  Get an active ETH key as shown within the keys page.
---

# Data Source `chainlink_eth_key`

Get an ETH wallet address.

## Example Usage

```terraform
data "chainlink_eth_key" "this" {}
```

## Schema

### Optional

- **index** (Integer, Optional, Default: 0) the index within the returned keys to return 

---
- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider

### Read-only

- **id** (String, Read-only) the Ethereum address of the ETH key
