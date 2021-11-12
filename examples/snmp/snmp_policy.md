### Resource Creation

```hcl
resource "intersight_snmp_policy" "snmp1" {
  name                    = "snmp1"
  description             = "testing smtp policy"
  enabled                 = true
  snmp_port               = 1983
  access_community_string = "dummy123"
  community_access        = "Disabled"
  trap_community          = "TrapCommunity"
  sys_contact             = "aanimish"
  sys_location            = "Karnataka"
  engine_id               = "vvb"
  snmp_users {
    name         = "demouser"
    privacy_type = "AES"
    #auth_password    = var.auth_password
    #privacy_password = var.privacy_password
    security_level = "AuthPriv"
    auth_type      = "SHA"
  }
  snmp_traps {
    destination = "10.10.10.1"
    enabled     = false
    port        = 660
    type        = "Trap"
    user        = "demouser"
    nr_version  = "V3"
  }
   profiles {
     moid        = var.profile
     object_type = "server.Profile"
   }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

 variable "auth_password" {
   type = string
   description = "value for auth_password"  
 }

 variable "privacy_password" {
   type =  string
   description = "value for privacy password"
 }
 
 variable "organization" {
 type = string
 description = "value for organization"  
 }

variable "profile"{
  type = string
  description = "Moid of server.Profile"
}
```