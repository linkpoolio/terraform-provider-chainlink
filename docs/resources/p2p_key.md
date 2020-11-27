---
page_title: "chainlink_p2p_key Resource - terraform-provider-chainlink"
subcategory: ""
description: |-
  Chainlink P2P key manages the lifecycle of a Chainlink P2P key
---

# Resource `chainlink_p2p_key`

`chainlink_p2p_key` manages a Chainlink P2P key

## Example Usage

```terraform

resource "chainlink_p2p_key" "this" {}

```

## Schema

### Required

None

### Optional

- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider

### Read-only

- **id** (String, Read-only) the ID of the P2P key
- **peer_id** (String, Read-only) the ID of the peer
- **public_key** (String, Read-only) the public key used within libp2p communication
