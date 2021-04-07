### Resource Creation

```hcl
resource "intersight_capability_io_card_descriptor" "capability_io_card_descriptor1" {   
    description = "capability io card descriptor"
    model = "UCSC-PCIE-C25Q-04"
    vendor = "Cisco Systems Inc"
    version = "4.0(4l)"
    capabilities [
      {
        object_type = "capability.IoCardDescriptor"
        moid = var.capability_io_card_descriptor
      }
    ]
    uif_connectivity = "inline"
}
```