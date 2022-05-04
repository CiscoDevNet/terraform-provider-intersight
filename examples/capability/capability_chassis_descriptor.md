### Resource Creation

```hcl
resource "intersight_capability_chassis_descriptor" "capability_chassis_descriptor1" {
  capabilities {
    moid        = var.capability_chassis_manufacturing_def
    object_type = "capability.ChassisManufacturingDef"
    class_id    = "capability.ChassisManufacturingDef"
  }
  description = "capability chassis descriptor"
  model       = "N20-C6508"
  revision    = "0"
  vendor      = "Cisco Systems Inc"
}

variable "capability_chassis_manufacturing_def" {
  type        = string
  description = "Moid of capability.ChassisManufacturingDef Mo"
}
```