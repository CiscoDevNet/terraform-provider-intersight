terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = var.version
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}
