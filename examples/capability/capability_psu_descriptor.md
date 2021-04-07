### Resource Creation

```hcl
resource "intersight_capability_psu_descriptor" "capability_psu_descriptor1" {
  description = "capability psu descriptor"
  model       = "UCSC-PSU1-1050W"
  vendor      = "Cisco Systems Inc"
  version     = "3.0(2)"
  capabilities = [
    {
      moid        = var.capability_psu_descriptor
      object_type = "capability.PsuDescriptor"
    }
  ]
}
```