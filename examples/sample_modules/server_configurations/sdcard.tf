resource "intersight_sdcard_policy" "sdcard1" {
  name        = "sdcard1"
  description = "test policy"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  partitions {
    type        = "OS"
    object_type = "sdcard.Partition"

    virtual_drives {
      enable      = true
      object_type = "sdcard.OperatingSystem"
      additional_properties = jsonencode({
        Name = "Hypervisor"
      })
    }
  }

  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}


/*
SAMPLE PAYLOAD
-----------------
SdcardPolicyApi: {
  "Name": "AUTO_TEST_SDCARD_MASTER",
  "Description": "SD Card Policy",
  "Tags": [{"Key": "policy", "Value": "sd_card"}],
  "Partitions": [{
    "ObjectType": "sdcard.Partition",
    "Type": "Utility",
    "VirtualDrives": [{"Enable": False,
      "ObjectType": "sdcard.Diagnostics"},
      {"Enable": False,
        "ObjectType": "sdcard.Drivers"},
      {"Enable": False,
        "ObjectType": "sdcard.HostUpgradeUtility"},
      {"Enable": False,
        "ObjectType": "sdcard.ServerConfigurationUtility"},
      {"Enable": False,
        "ObjectType": "sdcard.UserPartition"}]
  }]
}
*/