resource "intersight_storage_disk_group_policy" "storage_disk_group1" {
  name       = "storage_disk_group1"
  raid_level = "Raid0"
  use_jbods  = false
  span_groups {
    disks {
      slot_number = 1
    }
  }
  description = "Disk Group Test policy"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_storage_disk_group_policy" "storage_disk_group3" {
  name        = "storage_disk_group2"
  description = "Disk Group Test policy"
  raid_level  = "Raid1"
  use_jbods   = true
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

/*
SAMPLE PAYLOAD
-----------------
{
    "Name": "AUTO_DG_POLICY_RAID0_CRR",
    "Description": "Disk Group Test policy",
    "Tags": [{"Key": "dg", "Value": "dg_policy"}],
    "RaidLevel": "Raid0",
    "UseJbods": True,
    "SpanGroups": [{"Disks": [{"SlotNumber": 1}]}]
}

{
    "Name": "AUTO_DG_POLICY_RAID1_CRR",
    "Description": "Disk Group Test policy",
    "Tags": [{"Key": "dg", "Value": "dg_policy"}],
    "RaidLevel": "Raid1",
    "UseJbods": True,
    "SpanGroups": [{"Disks": [{"SlotNumber": 1}, {"SlotNumber": 2}]}],
    "DedicatedHotSpares": [{"SlotNumber": 3}]
}
*/