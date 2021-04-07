### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_replication_network_policy_deployment" "hyperflex_cluster_replication_network_policy_deployment1" {
  discovered = true
  replication_ip_ranges = [
    {
      end_addr   = "10.10.10.100"
      gateway    = "10.10.10.1"
      netmask    = "255.255.255.0"
      start_addr = "10.10.10.10"
    }
  ]
  replication_vlan {
    name        = "replication_vlan1"
    vlan_id     = 100
    object_type = "replication.NamedVlan"
    moid        = var.replication_named_vlan
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
}
```