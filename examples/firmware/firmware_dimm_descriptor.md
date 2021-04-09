### Resource Creation

```hcl
resource "intersight_firmware_dimm_descriptor" "firmware_dimm_descriptor1" {
  description = "firmware dimm descriptor"
  model       = "N01-M308GB2"
  version     = "0x46185EC2"
  vendor      = "Samsung Electronics, Inc."
  revision    = "0"
}
```