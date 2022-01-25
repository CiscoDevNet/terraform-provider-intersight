### Resource Creation

```hcl
resource "intersight_storage_storage_policy" "tf_storage_policy" {
  name                     = "tf_storage_policy"
  use_jbod_for_vd_creation = true
  description              = "storage policy test"
  unused_disks_state       = "UnconfiguredGood"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  global_hot_spares = "3"
  m2_virtual_drive {
    enable      = false
    object_type = "storage.M2VirtualDriveConfig"
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```