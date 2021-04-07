### Resource Creation

```hcl
resource "intersight_vnic_san_connectivity_policy" "vnic_san1" {
  name = "vnic_san1"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
```