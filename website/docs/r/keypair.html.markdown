---
layout: "outscale"
page_title: "3DS OUTSCALE: outscale_keypair"
sidebar_current: "outscale-keypair"
description: |-
  [Manages a keypair.]
---

# outscale_keypair Resource

Manages a keypair.
For more information on this resource, see the [User Guide](https://wiki.outscale.net/display/EN/About+Keypairs).
For more information on this resource actions, see the [API documentation](https://docs.outscale.com/api#3ds-outscale-api-keypair).

## Example Usage

```

# Create a keypair

resource "outscale_keypair" "keypair01" {
	keypair_name = "terraform-keypair-create"
}

# Import keypairs

resource "outscale_keypair" "keypair02" {
	keypair_name = "terraform-keypair-import-file"
	public_key   = file("<PATH>")
}

resource "outscale_keypair" "keypair03" {
	keypair_name = "terraform-keypair-import-text"
	public_key   = "UFVCTElDIEtFWQ=="
}

```

## Argument Reference

The following arguments are supported:

* `keypair_name` - (Required) A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
* `public_key` - (Optional) The public key. It must be base64-encoded.

## Attribute Reference

The following attributes are exported:

* `keypair` - Information about the created keypair.
  * `keypair_fingerprint` - If you create a keypair, the SHA-1 digest of the DER encoded private key.<br />
If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.
  * `keypair_name` - The name of the keypair.
  * `private_key` - The private key.


## Import

A keypair can be imported using its name. For example:

```

$ terraform import outscale_keypair.ImportedKeypair Name-of-the-Keypair

```