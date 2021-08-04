resource "intersight_storage_storage_policy" "tf_storage_storage" {
  name = "tf_storage_storage_policy"
  description = "storage policy test"
  retain_policy_virtual_drives = true
  unused_disks_state = "UnconfiguredGood"
  virtual_drives {
    name = "RAID0_1"
    boot_drive = true
    drive_cache = "Default"
    expand_to_available = false
    io_policy = "Direct"
    access_policy = "ReadWrite"
    disk_group_policy = intersight_storage_disk_group_policy.tf_storage_disk_group.id
    read_policy = "NoReadAhead"
    size = 285148
    write_policy = "WriteThrough"
  }
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}