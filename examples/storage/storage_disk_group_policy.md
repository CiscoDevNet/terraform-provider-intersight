### Resource Creation

```hcl
resource "intersight_storage_disk_group_policy" "storage_disk_group3" {
  name        = "storage_disk_group2"
  description = "demo disk group policy"
  raid_level  = "Raid1"
  use_jbods   = true
  span_groups = [{
    additional_properties = ""
    class_id              = "storage.SpanGroup"
    disks = [{
      additional_properties = ""
      class_id              = "storage.LocalDisk"
      object_type           = "storage.LocalDisk"
      slot_number           = 1
      },
      {
        additional_properties = ""
        class_id              = "storage.LocalDisk"
        object_type           = "storage.LocalDisk"
        slot_number           = 2
      }
    ]
    object_type = "storage.SpanGroup"
  }]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```