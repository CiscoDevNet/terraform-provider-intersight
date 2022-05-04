### Resource Creation

```hcl
resource "intersight_firmware_driver_distributable" "firmware_driver_distributable1" {
  description   = "firmware driver distributable"
  import_action = "Extract"
  name          = "firmware_driver_distributable1"
  component_meta {
    class_id        = "firmware.ComponentMeta"
    component_label = "UCS-IOM-2204XP"
    #image_path            = null
    model          = "UCS-IOM-2204XP"
    packed_version = "4.1(2S16)"
    redfish_url    = "/path/to/redfish/url"
    vendor         = "Cisco Systems Inc."
    object_type    = "firmware.ComponentMeta"
    component_type = "ALL"
    description    = "firmware meta component"
    #disruption            = "None"
    is_oob_supported = "false"
    oob_manageability = [
      "Update",
      "Inventory",
      "Activate"
    ]
  }
  category = "Win64"
  catalog {
    moid        = var.softwarerepository_catalog
    object_type = "softwarerepository.Catalog"
  }
}

variable "softwarerepository_catalog" {
  type        = string
  description = "Moid of softwarerepository_catalog"
}
```