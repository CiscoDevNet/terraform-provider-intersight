### Resource Creation

```hcl
resource "intersight_capability_psu_descriptor" "capability_psu_descriptor1" {
  description = "capability psu descriptor"
  model       = "UCSC-PSU1-1050W"
  vendor      = "Cisco Systems Inc"

  capabilities {
    moid        = var.capability_psu_descriptor
    object_type = "capability.PsuDescriptor"
    class_id    = "capability.PsuDescriptor"
  }
}

variable "capability_psu_descriptor" {
  type        = string
  description = "moid of capability.PsuDescriptor Mo"
}
```