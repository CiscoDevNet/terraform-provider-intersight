### Resource Creation

```hcl
resource "intersight_fabric_multicast_policy" "fabric_multicast_policy1" {
    name = "fabric_multicast_policy1"
    description = "fabric multicast policy"
    querier_ip_address = "192.168.0.1"
    querier_state = "Enabled"
    snooping_state = "Enabled"
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
}
```