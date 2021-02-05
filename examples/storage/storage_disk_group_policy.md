### Resource Creation

```hcl
resource "intersight_storage_disk_group_policy" "storage_disk_group3" {
  name = "storage_disk_group2"
  description = "Disk Group Test policy"
  raid_level = "Raid1"
  use_jbods = true
  span_groups {
    disks {
      slot_number = 2
    }
    disks {
      slot_number = 4
    }
  }
  dedicated_hot_spares {
    slot_number = 5
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}
```