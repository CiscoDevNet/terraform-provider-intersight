### Resource Creation

```hcl
resource "intersight_storage_storage_policy" "tf_storage_policy" {
  name = "tf_storage_policy"
  use_jbod_for_vd_creation = true
  unused_disks_state = "UnconfiguredGood"
  m2_virtual_drive {
    enable = false
  }
  global_hot_spares = "3"
  raid0_drive {
    enable = true
    virtual_drive_policy {
      strip_size = 64
      write_policy = "Default"
      read_policy = "Default"
      access_policy = "Default"
      drive_cache = "Default"
    }
  }
  organization {
    object_type = "organization.Organization"
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}
```