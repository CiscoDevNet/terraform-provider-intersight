### Resource Creation

```hcl
resource "intersight_access_policy" "access1" {
  name                     = "access1"
  description              = "test policy"
  inband_vlan              = 19
  inband_ip_pool {
    object_type = "ippool.Pool"
    moid = intersight_ip_pools.pool1.id
  }
  inband_vrf {
    object_type = "vrf.Vrves"
    moid = intersight_vrf_vrves.vrves1.id
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }

}
```