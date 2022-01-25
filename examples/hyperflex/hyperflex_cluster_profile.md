### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_profile" "hyperflex_cluster_profile1" {
  storage_data_vlan {
    name    = "hx-storage-data"
    vlan_id = 27
  }
  mgmt_ip_address    = "10.225.68.237"
  mac_address_prefix = "00:25:B5:D5"
  mgmt_platform      = "EDGE"
  replication        = 3
  description        = "This is hyperflex cluster profile"
  tags {
    key   = "test"
    value = "ucsback-10G-3nodehx-cluster-"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  name = "hyperflex_cluster_profile1"
}
```