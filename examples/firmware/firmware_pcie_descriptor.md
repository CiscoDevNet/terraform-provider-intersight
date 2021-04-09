### Resource Creation

```hcl
resource "intersight_firmware_pcie_descriptor" "firmware_pcie_descriptor1" {
  description = "firmware iom descriptor"
  model       = "UCSC-PCIE-IRJ45"
  vendor      = "Cisco Systems, Inc"
  revision    = "0"
}
```