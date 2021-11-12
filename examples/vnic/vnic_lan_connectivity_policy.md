### Resource Creation

```hcl
resource "intersight_vnic_lan_connectivity_policy" "vnic_lan1" {
  name                = "vnic_lan1"
  description         = "demo vnic lan connectivity policy"
  iqn_allocation_type = "None"
  placement_mode      = "auto"
  target_platform     = "FIAttached"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```