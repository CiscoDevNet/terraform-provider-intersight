### Resource Creation

```hcl
resource "intersight_vnic_iscsi_static_target_policy" "vnic_iscsi_static_target_policy" {
  name        = "vnic_iscsi_static_target_policy1"
  description = "vnic iscsi static target policy"
  ip_address  = "10.1.1.1"
  port        = 860
  lun {
    class_id    = "vnic.Lun"
    object_type = "vnic.Lun"
    bootable    = true
    lun_id      = 4
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
  type        = string
  description = "<value for organization>"
}
```