resource "intersight_storage_storage_policy" "storage_storage1" {
  name                         = "storage_storage_policy1"
  description                  = "storage policy test"
  retain_policy_virtual_drives = true
  unused_disks_state           = "UnconfiguredGood"
  virtual_drives {
    object_type = "storage.VirtualDriveConfig"
    boot_drive = true
    drive_cache = "NoChange"
    expand_to_available = false
    io_policy = "Direct"
    name = "RAID0_1"
    access_policy = "ReadWrite"
    disk_group_policy = intersight_storage_disk_group_policy.storage_disk_group1.id //92e
    read_policy = "NoReadAhead"
    size = 285148
    write_policy = "WriteThrough"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_storage_storage_policy" "storage_storage2" {
  name                         = "storage_storage_policy2"
  description                  = "storage policy test"
  retain_policy_virtual_drives = true
  unused_disks_state           = "Jbod"
  virtual_drives {
    object_type = "storage.VirtualDriveConfig"
    disk_group_policy = intersight_storage_disk_group_policy.storage_disk_group3.id //3e3
    drive_cache = "Enable"
    access_policy = "ReadWrite"
    boot_drive = false
    expand_to_available = false
    io_policy = "Direct"
    name = "vdrive1"
    read_policy = "Default"
    size = 102400
    write_policy = "AlwaysWriteBack"
  }
  virtual_drives {
    object_type = "storage.VirtualDriveConfig"
    access_policy = "ReadWrite"
    boot_drive = false
    io_policy = "Direct"
    name = "vdrive2"
    size = 100
    disk_group_policy = intersight_storage_disk_group_policy.storage_disk_group3.id //3de
    drive_cache = "Enable"
    expand_to_available = false
    read_policy = "Default"
    write_policy = "AlwaysWriteBack"
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
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
    "Name": "AUTO_STORAGE_POLICY_CRR",
    "Description": "storage policy test",
    "Tags": [{"Key": "policy", "Value": "storage"}],
    "RetainPolicyVirtualDrives": False,
    "EnableDriveSecurity": False,
    "UnusedDisksState": "Jbod",
    "VirtualDrives": [{
        "Name": "vdrive1",
        "Size": 102400,
        "DiskGroupPolicy": "",
        "AccessPolicy": "ReadWrite",
        "ReadPolicy": "Default",
        "WritePolicy": "AlwaysWriteBack",
        "IoPolicy": "Direct",
        "DriveCache": "Enable",
        "ExpandToAvailable": False,
        "BootDrive": False
    },
        {
        "Name": "vdrive2",
        "Size": 100,
        "DiskGroupPolicy": "",
        "AccessPolicy": "ReadWrite",
        "ReadPolicy": "Default",
        "WritePolicy": "AlwaysWriteBack",
        "IoPolicy": "Direct",
        "DriveCache": "Enable",
        "ExpandToAvailable": False,
        "BootDrive": False
    }]
}
*/