resource "intersight_storage_drive_group" "tf_drive_gp" {
  name       = "tf_drive_gp"
  raid_level = "Raid0"
  manual_drive_group {
    object_type = "storage.ManualDriveGroup"
    span_groups {
      slots       = "2"
      object_type = "storage.SpanDrives"
    }
  }
  virtual_drives {
    name                = "tf_drive_gp-vd"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    object_type         = "storage.VirtualDriveConfiguration"
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
  virtual_drives {
    name                = "drive_gp-vd-01"
    size                = 100
    expand_to_available = false
    boot_drive          = false
    object_type         = "storage.VirtualDriveConfiguration"
    virtual_drive_policy {
      strip_size    = 64
      write_policy  = "Default"
      read_policy   = "Default"
      access_policy = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
  storage_policy {
    moid        = intersight_storage_storage_policy.tf_storage_policy.moid
    object_type = "storage.StoragePolicy"
  }
}