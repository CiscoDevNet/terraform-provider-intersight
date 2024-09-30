terraform {
  required_providers {
    intersight = {
      source = "CiscoDevNet/intersight"
      version = "1.0.56"
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


#Boot precision policy for creating server profile template
resource "intersight_boot_precision_policy" "boot_precision1" {
  name                     = "boot_precision2"
  description              = "test policy"
  configured_boot_mode     = "Legacy"
  enforce_uefi_secure_boot = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  boot_devices {
    enabled     = true
    name        = "scu-device-hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  boot_devices {
    enabled     = true
    name        = "NIIODCIMCDVD"
    object_type = "boot.VirtualMedia"
    additional_properties = jsonencode({
      Subtype = "cimc-mapped-dvd"
    })
  }
  boot_devices {
    enabled     = true
    name        = "hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
}


resource "intersight_server_profile_template" "template1" {
  name            = "server_profile_template1"
  description     = "demo server profile template"
  target_platform = "FIAttached"
   organization {
     object_type = "organization.Organization"
     moid        = data.intersight_organization_organization.default.results.0.moid
   }
   policy_bucket {
     moid        = intersight_boot_precision_policy.boot_precision1.moid
     object_type = "boot.PrecisionPolicy"
   }
}


#resource "intersight_hcl_compatibility_status" "comptibility_status" {
#  request_type = "CheckCompatibility"
#}

#resource "intersight_os_os_support" "os_support"{
#  os_version = "Ubuntu Server 18.04 LTS"
#}

resource "intersight_bulk_mo_cloner" "clone_server1"{
       sources {
               class_id = "server.ProfileTemplate"
               object_type = "server.ProfileTemplate"
               moid = intersight_server_profile_template.template1.moid
       }
       targets {
               class_id = "server.Profile"
               object_type = "server.Profile"
               additional_properties = jsonencode({
                Name = "demotesting_DERIVED-4"
                Description = "Sample description"
               })
               tags = []
       }
}

resource "intersight_server_profile" "server_profile"{
	depends_on = [intersight_bulk_mo_cloner.clone_server1]
	name = "demotesting_DERIVED-4"
	description = "Sample description"
}

resource "intersight_server_profile" "server_profile1"{
	name = "demotesting_DERIVED-5"
	description = "Sample description"
  organization {
     object_type = "organization.Organization"
     moid        = data.intersight_organization_organization.default.results.0.moid
   }
}

resource "intersight_bulk_mo_merger" "trigger_profile_update" {
  class_id     = "bulk.MoMerger"
  merge_action = "Merge"
  object_type  = "bulk.MoMerger"
  sources {
    class_id    = "server.ProfileTemplate"
    moid        = intersight_server_profile_template.template1.moid
    object_type = "server.ProfileTemplate"
  }
  targets {
    class_id    = "server.Profile"
    moid        = intersight_server_profile.server_profile1.moid
    object_type = "server.Profile"
  }
}
