resource "intersight_storage_disk_group_policy" "tf_storage_disk_group" {
  name       = "tf_storage_disk_group"
  raid_level = "Raid0"
  use_jbods  = false
  span_groups {
    disks {
      slot_number = 1
    }
  }
  description = "Disk Group Test policy"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}