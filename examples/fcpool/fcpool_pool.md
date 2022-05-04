### Resource Creation

```hcl
resource "intersight_fcpool_pool" "fcpool_pool1" {
  name             = "fcpool_pool1"
  description      = "fcpool pool"
  assignment_order = "sequential"
  id_blocks {
    object_type = "fcpool.Block"
    from        = "50:00:00:00:00:00:00:00"
    to          = "50:50:00:00:00:00:00:00"
    class_id    = "fcpool.Block"
  }
  pool_purpose = "WWPN"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```