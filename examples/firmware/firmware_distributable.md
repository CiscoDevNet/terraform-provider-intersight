### Resource Creation

```hcl
resource "intersight_firmware_distributable" "firmware_distributable" {
  component_meta {
    class_id         = "firmware.ComponentMeta"
    component_label  = "UCS-FI-6454"
    component_type   = "NXOS"
    description      = "Cisco UCS 6454 Fabric Interconnect"
    disruption       = "EndpointReboot"
    is_oob_supported = false
    model            = "UCS-FI-6454"
    object_type      = "firmware.ComponentMeta"
    oob_manageability = [
      "Update",
      "Inventory",
      "Activate"
    ]
    packed_version = "7.0(3)N2(4.12S16)"
    redfish_url    = "/path/to/redfish/url"
    vendor         = "Cisco Systems Inc."
  }
  component_meta {
    class_id         = "firmware.ComponentMeta"
    component_label  = "UCS-IOM-2204XP"
    component_type   = "IOM"
    description      = "Cisco UCS 2204XP"
    disruption       = "EndpointReboot"
    is_oob_supported = false
    model            = "UCS-IOM-2204XP"
    object_type      = "firmware.ComponentMeta"
    oob_manageability = [
      "Update",
      "Inventory",
      "Activate"
    ]
    packed_version = "4.1(2S16)"
    redfish_url    = "/path/to/redfish/url"
    vendor         = "Cisco Systems Inc."
  }
  description    = "Cisco Intersight Infrastructure Bundle"
  image_category = "Unknown"
  import_action  = "None"
  md5sum         = "f4906255ae8f57ab5990066adba9e2f3"
  mdfid          = "IS-4GFI"
  name           = "ucs-intersight-infra-4gfi.4.1.2S19.gbin"
  origin         = "System"
  size           = 1355933747
  supported_models = [
    "UCS-FI-6454"
  ]
  tags {
    key   = "cisco.meta.distributabletype"
    value = "IMMFABRIC"
  }
  tags {
    key   = "cisco.meta.repositorytype"
    value = "IntersightCloud"
  }
  vendor = "Cisco"
}
```