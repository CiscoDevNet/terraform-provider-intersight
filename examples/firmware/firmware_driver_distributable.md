### Resource Creation

```hcl
resource "intersight_firmware_driver_distributable" "firmware_driver_distributable1" {
  description   = "firmware driver distributable"
  import_action = "Extract"
  name          = "firmware_driver_distributable1"
  component_meta = [
    {
      object_type      = "firmware.ComponentMeta"
      component_type   = "ALL"
      description      = "firmware meta component"
      disruption       = "None"
      is_oob_supported = "false"
      oob_manageability = [
        "Update",
        "Inventory",
        "Activate"
      ]
    }
  ]
  catalog {
    moid        = var.softwarerepository_catalog
    object_type = "softwarerepository.Catalog"
  }
}
```