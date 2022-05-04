resource "intersight_boot_precision_policy" "tf_boot_precision" {
  name                     = "tf_boot_precision"
  description              = "test policy"
  configured_boot_mode     = "Legacy"
  enforce_uefi_secure_boot = false
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  boot_devices {
    enabled     = true
    name        = "scu-device-hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  boot_devices {
    enabled     = true
    name        = "NIIODCIMCDVD"
    object_type = "boot.VirtualMedia"
    additional_properties = jsonencode({
      Subtype = "cimc-mapped-dvd"
    })
  }
  boot_devices {
    enabled     = true
    name        = "hdd"
    object_type = "boot.LocalDisk"
    additional_properties = jsonencode({
      Slot = "MRAID"
      Bootloader = {
        Description = ""
        Name        = ""
        ObjectType  = "boot.Bootloader"
        Path        = ""
      }
    })
  }
  tags {
    key   = "source"
    value = "terraform"
  }
}

/*
SAMPLE PAYLOAD
-----------------
BootPrecisionPolicyApi: {"Name": "AUTO_PBO_POLICY_CRR",
  "ConfiguredBootMode": "Uefi",
  "Description": "testing boot order policy",
  "Tags": [{"Key": "policy", "Value": "pbo"}],
  "EnforceUefiSecureBoot": True,
  "BootDevices": [{"Enabled": True,
    "Name": "iSCSI",
    "ObjectType": "boot.Iscsi",
    "Slot": "MLOM"},
    {"Enabled": True,
      "Name": "USB1",
      "ObjectType": "boot.Usb",
      "SubType": "usb-fdd"},
    {"Enabled": True,
      "Name": "Vmedia1",
      "ObjectType": "boot.VirtualMedia",
      "SubType": "cimc-mapped-hdd"},
    {"Enabled": True,
      "Name": "PCHstorage1",
      "ObjectType": "boot.PchStorage",
      "Lun": 3},
      {"Enabled": True,
       "Name": "ldhdd",
       "ObjectType": "boot.LocalDisk",
       "Slot": "MRAID",
       "Bootloader": {
           "Description": "san-boot-desc",
           "Name": "san-boot",
           "ObjectType": "boot.Bootloader",
           "Path": "/rhel/mnt/"
       }},
      {"Enabled": False,
       "Name": "Pxe",
       "ObjectType": "boot.Pxe",
       "MacAddress": "6a:00:00:5a:53:fe",
       "Port": 255,
       "Slot": "MLOM",
       "IpType": "IPv4"},
      {"Enabled": True,
       "Name": "San",
       "ObjectType": "boot.San",
       "Slot": "MLOM",
       "Bootloader": {
           "Description": "san-boot-desc",
           "Name": "san-boot",
           "ObjectType": "boot.Bootloader",
           "Path": "/rhel/mnt/"
       }},
       {"Enabled": False,
       "Name": "UEFIShell",
       "ObjectType": "boot.UefiShell"},
      {"Enabled": True,
       "Name": "SDCard",
       "ObjectType": "boot.SdCard",
       "SubType": "flex-util"},
      {"Enabled": False,
       "Name": "NVMe",
       "ObjectType": "boot.Nvme"}
   ]
}
*/