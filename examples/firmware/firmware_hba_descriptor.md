### Resource Creation

```hcl
resource "intersight_firmware_hba_descriptor" "firmware_hba_descriptor1" {
  description = "firmware hba descriptor"
  model       = "QLE2562"
  version     = "8.08.207-1"
  vendor      = "Cisco"
  revision    = "0"
}
```