### Resource Creation

```hcl
resource "intersight_firmware_distributable" "firmware_distributable" {
  component_meta = [
    {
      image_path            = null
      additional_properties = ""
      class_id              = "firmware.ComponentMeta"
      component_label       = "UCS-FI-6454"
      component_type        = "NXOS"
      description           = "Cisco UCS 6454 Fabric Interconnect"
      disruption            = "EndpointReboot"
      is_oob_supported      = false
      model                 = "UCS-FI-6454"
      object_type           = "firmware.ComponentMeta"
      oob_manageability = [
        "Update",
        "Inventory",
        "Activate"
      ]
      packed_version = "7.0(3)N2(4.12S16)"
      redfish_url    = "/path/to/redfish/url"
      vendor         = "Cisco Systems Inc."
    },
    {
      image_path            = null
      additional_properties = ""
      class_id              = "firmware.ComponentMeta"
      component_label       = "UCS-IOM-2204XP"
      component_type        = "IOM"
      description           = "Cisco UCS 2204XP"
      disruption            = "EndpointReboot"
      is_oob_supported      = false
      model                 = "UCS-IOM-2204XP"
      object_type           = "firmware.ComponentMeta"
      oob_manageability = [
        "Update",
        "Inventory",
        "Activate"
      ]
      packed_version = "4.1(2S16)"
      redfish_url    = "/path/to/redfish/url"
      vendor         = "Cisco Systems Inc."
    }
  ]
  description      = "Cisco Intersight Infrastructure Bundle"
  download_count   = 0
  image_category   = "Unknown"
  import_action    = "None"
  import_state     = "Imported"
  imported_time    = "0001-01-01T00:00:00Z"
  last_access_time = "0001-01-01T00:00:00Z"
  md5sum           = "f4906255ae8f57ab5990066adba9e2f3"
  mdfid            = "IS-4GFI"
  mod_time         = "2020-07-03T15:30:39.506Z"
  name             = "ucs-intersight-infra-4gfi.4.1.2S19.gbin"
  origin           = "System"
  size             = 1355933747
  supported_models = [
    "UCS-FI-6454"
  ]
  tags = [
    {
      additional_properties = ""
      key                   = "cisco.meta.distributabletype"
      value                 = "IMMFABRIC"
    },
    {
      additional_properties = ""
      key                   = "cisco.meta.repositorytype"
      value                 = "IntersightCloud"
    }
  ]
  vendor = "Cisco"
}
```