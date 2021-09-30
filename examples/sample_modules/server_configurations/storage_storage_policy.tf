resource "intersight_storage_storage_policy" "tf_storage_policy" {
  name               = "tf_storage_policy"
  use_jbod_for_vd_creation = true
  description        = "storage policy test"
  unused_disks_state = "UnconfiguredGood"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  global_hot_spares = "3"
  m2_virtual_drive {
    enable      = false
    object_type = "storage.M2VirtualDriveConfig"
  }
  raid0_drive {
    enable      = true
    drive_slots = "2"
    object_type = "storage.R0Drive"
    virtual_drive_policy {
      strip_size    = 64
      access_policy = "Default"
      read_policy   = "Default"
      write_policy  = "Default"
      drive_cache   = "Default"
      object_type   = "storage.VirtualDrivePolicy"
    }
  }
}