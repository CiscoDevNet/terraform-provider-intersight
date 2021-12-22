terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.21"
    }
  }
}

provider "intersight" {
  apikey = var.api_key
  secretkey = var.secretkey
  endpoint = var.endpoint
}

data "intersight_organization_organization" "default" {
  name  = "default"
}

resource "intersight_bulk_mo_cloner" "clone_server"{
	sources {
		moid = "61c344db77696e2d314cc684"
		object_type = "server.ProfileTemplate"
	}
	targets {
		name = "demotest3_DERIVED-1"
		object_type = "server.Profile"
		description = ""
		tags = []
	}
}

#resource "intersight_hcl_compatibility_status" "comptibility_status" {
#  request_type = "CheckCompatibility"
#}

#resource "intersight_os_os_support" "os_support"{
#  os_version = "Ubuntu Server 18.04 LTS"
#}
