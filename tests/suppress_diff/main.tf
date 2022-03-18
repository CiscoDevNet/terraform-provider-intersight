terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.25"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}

data "intersight_firmware_distributable" "recommended_version" {
    supported_models = ["UCSC-C240-M5SX"]
	#owners = ["5c4a74693935687776d94f98"]
    recommended_build = "Y"
}

output "versions" {
	value = length(data.intersight_firmware_distributable.recommended_version.results)
}
