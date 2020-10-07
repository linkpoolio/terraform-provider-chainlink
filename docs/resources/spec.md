---
page_title: "chainlink_spec Resource - terraform-provider-chainlink"
subcategory: ""
description: |-
  Chainlink spec manages the lifecycle of a Chainlink job spec.
---

# Resource `chainlink_spec`

`chainlink_spec` manages a Chainlink job specification.

## Example Usage

```terraform

resource "chainlink_bridge" "coinmarketcap" {
    name = "coinmarketcap"
    url  = "http://coinmarketcap.local:8080/"
}

locals {
  http_get_uint256 = {
    initiators = [
      {
        type = "runlog"
      }
    ]
    tasks = [
      {
        type = chainlink_bridge.coinmarketcap.name
      },
      {
        type = "multiply"
      },
      {
        type = "ethuint256"
      },
      {
        type = "ethtx"
      }
    ]
  }
}

resource "chainlink_spec" "vrf" {
  json = jsonencode(local.http_get_uint256)
}

```

## Schema

### Required

- **json** (String, Required) The encoded JSON object of the job specification.

### Read-only

- **spec_id** (String, Read-only) the ID of the job spec
