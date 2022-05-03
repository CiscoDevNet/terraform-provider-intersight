### Resource Creation

```hcl
resource "intersight_firmware_switch_upgrade" "firmware_switch_upgrade1" {
  direct_download {
    object_type   = "firmware.DirectDownload"
    image_source  = "cisco"
    password      = "ChangeMe"
    upgradeoption = "chassis_upgrade_full"
    username      = "admin1"
  }
  network_share {
    cifs_server {
      file_location = "10.1.1.1"
      object_type   = "firmware.CifsServer"
      mount_options = "none"
    }
    map_type      = "cifs"
    upgradeoption = "nw_upgrade_full"
    username      = "admin1"
    object_type   = "firmware.NetworkShare"
  }
  skip_estimate_impact = false
  status               = "SUCCESSFUL"
  upgrade_type         = "network_upgrade"
}
```