### Resource Creation

```hcl
resource "intersight_firmware_bios_descriptor" "firmware_bios_descriptor1" {
  description = "firmware bios descriptor"
  model       = "N20-B6620-1"
  version     = "S5500.86B.01.00.0036-191.0.1620091126"
  vendor      = "Cisco Systems Inc"
  revision    = "0"
}
```