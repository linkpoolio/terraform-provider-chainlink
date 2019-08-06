Terraform Provider
==================

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/linkpoolio/terraform-provider-chainlink`

```sh
$ mkdir -p $GOPATH/src/github.com/linkpoolio; cd $GOPATH/src/github.com/linkpoolio
$ git clone git@github.com:linkpoolio/terraform-provider-chainlink
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/linkpoolio/terraform-provider-chainlink
$ go build -o terraform-provider-chainlink_v0.2
```

Using the provider
----------------------

You can configure the provider with the following:
```
provider "chainlink" {
    email    = "admin@node.local"
    password = "twochains"
    url      = "http://localhost:6688"
}
```

All the variables defined above are the defaults. Keep in-mind that if you write the password in clear text within your
HCL files, or fetch it via other methods, it will be stored in plain-text or in-state.

Available resources:

- cl_bridge
- cl_spec

#### cl_bridge

This will create and manage bridge types (external adaptors) on the Chainlink node, an example:

```
resource "cl_bridge" "asset_price" {
    name       = "assetprice"
    url        = "http://localhost:8080/price"
}
```

#### cl_spec

This will create and manage bridge types (external adaptors) on the Chainlink node, an example:

```
resource "cl_spec" "httpget_uint256" {
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