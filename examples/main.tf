terraform {
  required_providers {  
    cloudux-utils = {
      source  = "AvidCloud/cloudux-utils"
      version = ">= 0.0.5"
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
  value = cloudux-utils_site_key.example.public_key
}

output "private_key" {
  value = cloudux-utils_site_key.example.private_key
}