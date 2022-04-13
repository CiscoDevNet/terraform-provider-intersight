### Resource Creation

```hcl
resource "intersight_storage_drive_group" "tf_drive_gp" {
  name       = "tf_drive_gp"
  raid_level = "Raid0"
  manual_drive_group {
    span_groups {
      slots = "2"
    }
  }
  virtual_drives {
    name                = "tf_drive_gp-vd"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
    }
  }
  virtual_drives {
    name                = "drive_gp-vd-01"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
    }
  }
  storage_policy {
    moid = var.tf_storage_policy
  }
}

variable "tf_storage_policy" {
  type        = string
  description = "value for tf_storage_policy"
}
```