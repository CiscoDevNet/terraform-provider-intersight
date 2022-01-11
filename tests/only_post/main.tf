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
               moid = "61dd4d1477696e2d315f374f"
       }
       targets {
               class_id = "server.Profile"
               object_type = "server.Profile"
               additional_properties = jsonencode({
			Name = "demotesting_DERIVED-4"
			Description = "Sample description"
		})
#               name = "demotesting_DERIVED-4"
#               description = "Sample description"
               tags = []
       }
}

#resource "intersight_feedback_feedback_post" "feedback1"{
#	feedback_data {
#		account_name = "merajash"
#		comment = "Testing feedback"
#		email = "merajash@cisco.com"
#		evaluation = "Good"
#		follow_up = "false"
#		type = "Evaluation"
#	}
#}
