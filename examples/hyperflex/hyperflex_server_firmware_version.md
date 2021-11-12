### Resource Creation

```hcl
resource "intersight_hyperflex_server_firmware_version" "hyperflex_server_firmware_version1" {
  app_catalog {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  parent {
    object_type = "hyperflex.AppCatalog"
    moid        = "6115d0746973682d31fe8d13"
  }
  server_firmware_version_entries = [
    {
      object_type           = "hyperflex.ServerFirmwareVersionEntry"
      class_id              = "hyperflex.ServerFirmwareVersionEntry"
      additional_properties = ""
      selector              = null
      moid                  = var.hyperflex_server_firmware_version_entry
    }
  ]
}

variable "hyperflex_server_firmware_version_entry" {
  type        = string
  description = "Moid of hyperflex_server_firmware_version_entry "
}
```