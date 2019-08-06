provider "chainlink" {
  url      = "http://localhost:6688"
  email    = "admin@node.local"
  password = "twochains"
}

resource "chainlink_bridge" "this" {
  name = "${var.bridge_name}"
  url  = "${var.bridge_url}"
}

data "template_file" "this" {
  template = "${file("${path.module}/templates/spec.json")}"

  vars {
    bridge = "${var.bridge_name}"
  }
}

resource "chainlink_spec" "this" {
  json = "${data.template_file.this.rendered}"

  depends_on = ["chainlink_bridge.this"]
}