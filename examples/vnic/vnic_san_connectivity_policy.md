### Resource Creation

```hcl
resource "intersight_vnic_san_connectivity_policy" "vnic_san1" {
  name = "vnic_san1"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  profiles {
    moid        = var.profile
    object_type = "server.Profile"
  }
}

variable "profile"{
  type = string
  description = "Moid of server.Profile"
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```