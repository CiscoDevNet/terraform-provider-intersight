### Resource Creation

```hcl
resource "intersight_capability_fan_module_manufacturing_def" "capability_fan_module_manufacturing_def1" {
  caption      = "Cisco UCS S3260 Fan"
  description  = "Cisco UCS S3260 Fan module containing 2x80mm fans FRU"
  name         = "fan1"
  pid          = "UCSC-C3X60-FANM"
  product_name = "Cisco UCS S3260 Fan module"
  sku          = "UCSC-C3X60-FANM"
  vid          = "V01"
}
```