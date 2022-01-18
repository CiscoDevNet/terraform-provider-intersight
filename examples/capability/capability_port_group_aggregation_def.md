### Resource Creation

```hcl
resource "intersight_capability_port_group_aggregation_def" "capability_port_group_aggregation_def1" {
  name                  = "capability_port_group_aggregation_def1"
  aggregation_cap       = "operational"
  hw40_g_port_group_cap = false
}
```