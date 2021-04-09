### Resource Creation

```hcl
resource "intersight_firmware_iom_descriptor" "firmware_iom_descriptor1" {
  description = "firmware iom descriptor"
  model       = "UCS-IOM-2208XP"
  version     = "4.0(4c)"
  vendor      = "Cisco"
  revision    = "0"
}
```