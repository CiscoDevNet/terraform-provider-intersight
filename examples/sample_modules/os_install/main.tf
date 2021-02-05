terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "0.1.6"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}
