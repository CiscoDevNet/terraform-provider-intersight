### Resource Creation

```hcl
resource "intersight_capability_io_card_descriptor" "capability_io_card_descriptor1" {
  description = "capability io card descriptor"
  model       = "UCSC-PCIE-C25Q-04"
  vendor      = "Cisco Systems Inc"
  capabilities = [
    {
      object_type           = "capability.IoCardDescriptor"
      class_id              = "capability.IoCardDescriptor"
      moid                  = var.capability_io_card_descriptor
      additional_properties = ""
      selector              = ""
    }
  ]
  uif_connectivity = "inline"
}

variable "capability_io_card_descriptor" {
  type        = string
  description = "moid of capability.IoCardDescriptor Mo"
}
```