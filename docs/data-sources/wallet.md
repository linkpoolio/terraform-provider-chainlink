---
page_title: "chainlink_wallet Data Source - terraform-provider-chainlink"
subcategory: ""
description: |-
  Get the active nodes wallet address as shown in Chainlink config
---

# Data Source `chainlink_wallet`

Get the active wallet address as shown in config.

## Example Usage

```terraform
data "chainlink_wallet" "this" {}
```

## Schema

### Optional

- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider

### Read-only

- **address** (String, Read-only) the Ethereum address of the Chainlink node
