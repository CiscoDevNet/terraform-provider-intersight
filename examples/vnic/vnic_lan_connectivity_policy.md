### Resource Creation

```hcl
resource "intersight_vnic_lan_connectivity_policy" "vnic_lan1" {
  name = "vnic_lan1"
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