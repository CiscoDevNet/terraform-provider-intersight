### Resource Creation

```hcl
resource "intersight_capability_chassis_manufacturing_def" "capability_chassis_manufacturing_def1" {
  caption           = "Cisco UCS 5108 DC2 Blade Server Chassis"
  chassis_code_name = "Blade Server Chassis"
  description       = "Cisco DC Blade Server Chassis 6U with Eight Blade Server Slots"
  name              = "chm2"
  pid               = "UCSB-5108-DC2"
  product_name      = "Cisco UCS 5108 DC2 Chassis"
  sku               = "UCSB-5108-DC2"
  vid               = "V01"
}
```