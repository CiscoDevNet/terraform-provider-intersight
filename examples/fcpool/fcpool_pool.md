### Resource Creation

```hcl
resource "intersight_fcpool_pool" "fcpool_pool1" {
    name = "fcpool_pool1"
    description = "fcpool pool"
    assignment_order = "sequential"
    id_blocks [{
        object_type = "fcpool.Block"
        from = "50:00:00:00:00:00:00:00"
        to = "5F:FF:FF:FF:FF:FF:FF:FF"
    }]
    pool_purpose = "To create fcpool"
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
}
```