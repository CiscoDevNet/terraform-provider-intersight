### Resource Creation

```hcl
resource "intersight_capability_switch_descriptor" "capability_switch_descriptor1" {
  description = "capability switch descriptor"
  model       = "UCS-FI-6454"
  vendor      = "Cisco Systems Inc"
  version     = "3.1(2c)"
  capabilities = [
    {
      object_type = "capability.SwitchDescriptor"
      moid        = var.capability_switch_descriptor
    }
  ]
}
```