### Resource Creation

```hcl
resource "intersight_firmware_sas_expander_descriptor" "firmware_sas_expander_descriptor1" {
  description = "firmware sas expander descriptor"
  model       = "R210-SASXPAND"
  version     = "04.08.01_B056"
  vendor      = "Cisco Systems, Inc"
  revision    = "0"
}
```