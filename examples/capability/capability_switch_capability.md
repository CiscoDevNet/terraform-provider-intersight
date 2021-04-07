### Resource Creation

```hcl
resource "intersight_capability_switch_capability" "capability_switch_capability1" {
  name                     = "capability_switch_capability1"
  pid                      = "UCS-FI-6454"
  sku                      = "CON-NCF4P-FI6454CH"
  vid                      = "V00"
  dynamic_vifs_supported   = true
  fan_modules_supported    = true
  locator_beacon_supported = true
}
```