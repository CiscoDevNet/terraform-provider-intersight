### Resource Creation

```hcl
resource "intersight_firmware_hba_descriptor" "firmware_hba_descriptor1" {
  description = "firmware hba descriptor"
  model       = "QLE2562"
  vendor      = "Cisco"
  revision    = "0"
}
```