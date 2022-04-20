terraform {
  required_providers {  
    cloudux-utils = {
      source  = "AvidCloud/cloudux-utils"
      version = ">= 0.0.3"
    }
  }
}

provider "cloudux-utils" {
  # Configuration options
}

resource "cloudux-utils_site_key" "example" {
    rsa_bits = 4096
}

output "public_key" {
  value = site_key.example.public_key
}

output "private_key" {
  value = site_key.example.private_key
}