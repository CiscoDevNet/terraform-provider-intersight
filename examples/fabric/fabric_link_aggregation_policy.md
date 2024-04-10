### Resource Creation

```hcl
resource "intersight_fabric_link_aggregation_policy" "fabric_link_aggregation_policy1" {
    name               = "fabric_link_aggregation_policy1"
    suspend_individual = "false"
    lacp_rate          = "normal"

  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```
