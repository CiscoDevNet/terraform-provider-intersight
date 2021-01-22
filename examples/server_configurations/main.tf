terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "0.1.5"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkeyfile = var.secretkey_file
  endpoint = var.endpoint
}
