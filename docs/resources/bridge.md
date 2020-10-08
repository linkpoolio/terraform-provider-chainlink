---
page_title: "chainlink_bridge Resource - terraform-provider-chainlink"
subcategory: ""
description: |-
  Chainlink spec manages the lifecycle of a Chainlink bridge.
---

# Resource `chainlink_bridge`

`chainlink_bridge` manages a Chainlink bridge.

## Example Usage

```terraform

resource "chainlink_bridge" "coinmarketcap" {
    name = "coinmarketcap"
    url  = "http://coinmarketcap.local:8080/"
}

```

## Schema

### Required

- **name** (String, Required) the name of the Chainlink bridge, which is then referenced in any job specs.
- **url** (String, Required) the URL of the bridge, called by the node on request 

### Optional

- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider
