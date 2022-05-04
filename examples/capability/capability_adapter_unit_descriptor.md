### Resource Creation

```hcl
resource "intersight_capability_adapter_unit_descriptor" "capability_adapter_unit_descriptor1" {
  description = "capability adapter unit descriptor"
  model       = "Cisco Systems Inc"
  capabilities {
    moid        = var.capability_adapter_unit_descriptor
    object_type = "capability.AdapterUnitDescriptor"
    class_id    = "capability.AdapterUnitDescriptor"
  }
  connectivity_order  = "sequential"
  ethernet_port_speed = 40
}

variable "capability_adapter_unit_descriptor" {
  type        = string
  description = "Moid of capability.AdapterUnitDescriptor Mo"
}
```