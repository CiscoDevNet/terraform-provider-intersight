terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "0.1.1-8.appliance"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}
