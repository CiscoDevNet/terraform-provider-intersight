### Resource Creation

```hcl
resource "intersight_firmware_drive_descriptor" "firmware_drive_descriptor1" {
  description = "firmware drive descriptor"
  model       = "LT0400MO"
  version     = "C401"
  vendor      = "Dell EMC Enterprise SSDs"
  revision    = "0"
}
```