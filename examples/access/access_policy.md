### Resource Creation

```hcl
resource "intersight_access_policy" "access1" {
  name        = "access1"
  description = "demo imc access policy"
  inband_vlan = 19
  inband_ip_pool {
    object_type = "ippool.Pool"
    moid        = intersight_ip_pool.pool1.id
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```