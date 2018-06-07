Terraform Provider
==================

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/linkpoolio/terraform-provider-clnode`

```sh
$ mkdir -p $GOPATH/src/github.com/linkpoolio; cd $GOPATH/src/github.com/linkpoolio
$ git clone git@github.com:linkpoolio/terraform-provider-clnode
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/linkpoolio/terraform-provider-clnode
$ dep ensure && go build -o terraform-provider-cl-node_v.0.1
```

Using the provider
----------------------

You can configure the provider with the following:
```
provider "clnode" {
    username = "chainlink"
    password = "twochains"
    protocol = "http"
}
```
The attributes above are all the default values and are optional. The username/password is the credentials to be used to authenticate against the node and the protocol to be used can only be http/https.

The address of the node is specified in each resource to aid managing clusters of nodes.

Available resources:

- clnode_bridge_type

#### clnode_bridge_type

This will create and manage bridge types (external adaptors) on the ChainLink node, an example:

```
resource "clnode_bridge_type" "asset_price" {
    node_addr  = "localhost:6688"
    name       = "assetprice"
    url        = "http://localhost:8080/price"
}
```