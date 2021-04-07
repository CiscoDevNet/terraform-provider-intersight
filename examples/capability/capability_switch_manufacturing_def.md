### Resource Creation

```hcl
resource "intersight_capability_switch_manufacturing_def" "capability_switch_manufacturing_def1" {
    name = "capability_switch_manufacturing_def"
    pid = "UCS-FI-6454"
    sku = "CON-NCF4P-FI6454CH"
    vid = "V00"
    caption = "UCS 6454 1RU FI, with no PSU"
    description = "Configured model: UCS 6454 1RU FI, with no PSU, with 54 ports and includes 18x10/25-Gbps and 2x40/100-Gbps port licenses"
    part_number = "UCS-FI-6454"
    product_name = "UCS 6454 1RU FI"
}
```