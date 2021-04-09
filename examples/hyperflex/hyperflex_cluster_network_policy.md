### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_network_policy" "hyperflex_cluster_network_policy1" {
  mgmt_vlan {
    name    = "hx-inband-mgmt"
    vlan_id = 301
  }
  jumbo_frame  = true
  uplink_speed = "10G"
  mac_prefix_range {
    end_addr   = "00:25:B5:D5"
    start_addr = "00:25:B5:D5"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
  name = "hyperflex_cluster_network_policy1"
}
```