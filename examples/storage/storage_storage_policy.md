### Resource Creation

```hcl
resource "intersight_storage_storage_policy" "storage_storage1" {
  name = "storage_storage_policy1"
  description = "storage policy test"
  retain_policy_virtual_drives = true
  unused_disks_state = "UnconfiguredGood"
  virtual_drives {
    object_type = "storage.VirtualDriveConfig"
    boot_drive = true
    drive_cache = "NoChange"
    expand_to_available = false
    io_policy = "Direct"
    name = "RAID0_1"
    access_policy = "ReadWrite"
    disk_group_policy = intersight_storage_disk_group_policy.storage_disk_group1.id
    read_policy = "NoReadAhead"
    size = 285148
    write_policy = "WriteThrough"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```