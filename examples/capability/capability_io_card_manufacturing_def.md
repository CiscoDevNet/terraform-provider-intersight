### Resource Creation

```hcl
resource "intersight_capability_io_card_manufacturing_def" "capability_io_card_manufacturing_def1" {
  caption      = "Cisco UCS VIC 1455"
  description  = "Cisco UCS VIC 1455 Quad Port 10/25G SFP28 CNA PCIE"
  name         = "vic_1455"
  pid          = "UCSC-PCIE-C25Q-04"
  product_name = "Cisco UCS VIC 1455 Quad Port"
  sku          = "UCSC-PCIE-C25Q-04"
  vid          = "V00"
}
```