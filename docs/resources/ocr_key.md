---
page_title: "chainlink_ocr_key Resource - terraform-provider-chainlink"
subcategory: ""
description: |-
  Chainlink OCR key manages the lifecycle of a Chainlink OCR key
---

# Resource `chainlink_ocr_key`

`chainlink_ocr_key` manages a Chainlink OCR key

## Example Usage

```terraform

resource "chainlink_ocr_key" "this" {}

```

## Schema

### Required

None

### Optional

- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider

### Read-only

- **id** (String, Read-only) the ID of the OCR key
- **config_public_key** (String, Read-only) the config public key used
- **offchain_public_key** (String, Read-only) the off-chain public key
- **onchain_signing_address** (String, Read-only) the address used for off-chain signing of messages
