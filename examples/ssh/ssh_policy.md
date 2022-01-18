### Resource Creation

```hcl
resource "intersight_ssh_policy" "ssh_policy1" {
  name        = "ssh_policy1"
  description = "ssh policy"
  enabled     = true
  port        = 22
  timeout     = 1800
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  profiles {
    moid        = var.profile
    object_type = "server.Profile"
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }

 variable "profile"{
    type = string
    description = "Moid of server.Profile"
 }
```