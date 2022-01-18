### Resource Creation

```hcl
resource "intersight_capability_fan_module_descriptor" "capability_fan_module_descriptor1" {
  capabilities = [
    {
      moid                  = var.capability_fan_module_descriptor
      object_type           = "capability.FanModuleDescriptor"
      class_id              = "capability.FanModuleDescriptor"
      additional_properties = ""
      selector              = ""
    }
  ]
  description = "capability of fan module descriptor"
  model       = "UCSC-C3X60-FANM"
  revision    = "0"
  vendor      = "Cisco Systems Inc"

}

variable "capability_fan_module_descriptor" {
  type        = string
  description = "moid of capability.FanModuleDescriptor Mo"
}
```