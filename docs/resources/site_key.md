---
page_title: "Site Key"
description: |-
  Generates a site key pair
---

# site_key Resource/Data Source

Creates a site key

## Example Usage

```hcl
resource "site_key" "mykey" {
  
}

output "publickey" {
  filename = "/mypath/public_key.pem"
  content     = site_key.mykey.public_key
}

output "publickey" {
  filename = "/mypath/private_key.pem"
  content     = site_key.mykey.private_key
}

```


## Attribute Reference
* `private_key` - The private key 
* `public_key` - The public key 
