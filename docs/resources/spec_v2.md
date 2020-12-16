---
page_title: "chainlink_spec_v2 Resource - terraform-provider-chainlink"
subcategory: ""
description: |-
  Chainlink spec manages the lifecycle of a V2 Chainlink job spec.
---

# Resource `chainlink_spec_v2`

`chainlink_spec_v2` manages a V2 Chainlink job specification.

## Example Usage

```terraform

resource "chainlink_spec_v2" "bootstrap" {
  toml = "isBootstrapNode = true"
}

```

## Schema

### Required

- **toml** (String, Required) The TOML encoded object of the job specification.

### Optional

- **chainlink_url** (String, Optional) equivalent to `url` in the provider configuration, takes precedence over the provider
- **chainlink_email** (String, Optional) equivalent to `email` in the provider configuration, takes precedence over the provider
- **chainlink_password** (String, Optional) equivalent to `password` in the provider configuration, takes precedence over the provider

### Read-only

- **id** (String, Read-only) the ID of the v2 job spec
