### Resource Creation

```hcl
resource "intersight_capability_chassis_descriptor" "capability_chassis_descriptor1" {
  capabilities = [
    {
      moid        = var.capability_chassis_manufacturing_def
      object_type = "capability.ChassisManufacturingDef"
    }
  ]
  description  = "capability chassis descriptor"
  model        = "N20-C6508"
  revision     = "0"
  shared_scope = "shared"
  vendor       = "Cisco Systems Inc"
  version      = "ge 4.2(2.1)"
}
```