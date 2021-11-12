### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_replication_network_policy_deployment" "hyperflex_cluster_replication_network_policy_deployment1" {
  discovered = true
  replication_vlan {
    name        = "replication_vlan1"
    vlan_id     = 100
    object_type = "replication.NamedVlan"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```