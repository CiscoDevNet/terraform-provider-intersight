### Resource Creation

```hcl
resource "intersight_capability_switch_descriptor" "capability_switch_descriptor1" {
  description = "capability switch descriptor"
  model       = "UCS-FI-6454"
  vendor      = "Cisco Systems Inc"

  capabilities = [
    {
      object_type           = "capability.SwitchDescriptor"
      class_id              = "capability.SwitchDescriptor"
      moid                  = var.capability_switch_descriptor
      additional_properties = ""
      selector              = ""
    }
  ]
}

variable "capability_switch_descriptor" {
  type        = string
  description = "moid of capability.SwitchDescriptor Mo"
}
```