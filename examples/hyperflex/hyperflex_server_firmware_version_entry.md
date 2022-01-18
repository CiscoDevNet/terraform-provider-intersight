### Resource Creation

```hcl
resource "intersight_hyperflex_server_firmware_version_entry" "hyperflex_server_firmware_version_entry1" {
  constraint {
    hxdp_version    = "4.0(2a)|4.0(2b)|4.0(2c)|4.0(2d)|4.0(2e)|4.5(1a)"
    hypervisor_type = "ESXi"
    mgmt_platform   = "EDGE"
    object_type     = "hyperflex.AppSettingConstraint"
    server_model    = "HX.*M5.*$"
  }
  parent {
    object_type = "hyperflex.ServerFirmwareVersion"
    moid        = var.hyperflex_server_firmware_version
  }
  server_platform = "M5"
}

variable "hyperflex_server_firmware_version" {
  type        = string
  description = "Moid of hyperflex_server_firmware_version"
}
```