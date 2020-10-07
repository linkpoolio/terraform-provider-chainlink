---
layout: ""
page_title: "Provider: Chainlink"
description: |-
  The Chainlink provider provides the ability to interact with a Chainlink nodes REST API and perform various actions.
---

# Chainlink Provider

The Chainlink provider provides the ability to interact with a Chainlink nodes REST API and perform various actions. 
For example creating a job specification or bridge.

Requirements:

- A deployed Chainlink node with the sign-in email and password

## Example Usage

```
provider "chainlink" {
    email    = "admin@node.local"
    password = "twochains"
    url      = "http://localhost:6688"
}

resource "chainlink_spec" "httpget_uint256" {
    json = <<-EOF
{
  "initiators": [
    {
      "type": "runlog"
    }
  ],
  "tasks": [
    {
      "type": "httpget"
    },
    {
      "type": "jsonparse"
    },
    {
      "type": "multiply"
    },
    {
      "type": "ethuint256"
    },
    {
      "type": "ethtx"
    }
  ]
} 
EOF
}

```

## Schema

### Required

- **url** (String, Required) the URL of the Chainlink node
- **email** (String, Required) the sign-in email used to login to the Chainlink node
- **password** (String, Required) the password for the account used to sign-in
