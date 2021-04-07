### Resource Creation

```hcl
resource "intersight_firmware_psu_descriptor" "firmware_psu_descriptor1" {
  description = "Cisco UCS S3260 1050W (AC) Power Supply Unit"
  model       = "UCSC-PSU1-1050W"
  vendor      = "Cisco Systems, Inc"
  revision    = "0"
}
```