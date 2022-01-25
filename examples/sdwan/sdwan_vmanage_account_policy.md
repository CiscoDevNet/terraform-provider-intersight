### Resource Creation

```hcl
resource "intersight_sdwan_vmanage_account_policy" "sdwan_vmanage_account_policy1" {
  name             = "sdwan_vmanage_account_policy1"
  description      = "sdwan_vmanage_account_policy"
  endpoint_address = "<server fqdn>"
   organization {
     object_type = "organization.Organization"
     moid        = var.organization
   }
   profiles {
     object_type = "sdwan.Profile"
     moid        = var.sdwan_profile1
   }
  username = "username to authenticate the vmanage_server"

}

 variable "organization" {
   type = string
   description = "value for organization"
 }
 
 variable "sdwan_profile1" {
     type = string
     description = "value for sdwan_profile"  
   }
```