### Resource Creation

```hcl
resource "intersight_hyperflex_server_firmware_version" "hyperflex_server_firmware_version1" {
    app_catalog {
        object_type = "hyperflex.AppCatalog"
        moid = "hyperflex_app_catalog"
    }
    parent {
      object_type = "hyperflex.AppCatalog"
      moid = var.hyperflex_app_catalog
    }
    server_firmware_version_entries [
      {
        object_type = "hyperflex.ServerFirmwareVersionEntry"
        moid = var.hyperflex_server_firmware_version_entry
      }
    ]
}
```