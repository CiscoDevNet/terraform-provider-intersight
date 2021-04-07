### Resource Creation

```hcl
resource "intersight_hyperflex_node_config_policy" "hyperflex_node_config_policy1" {
  mgmt_ip_range {
    end_addr   = "10.225.68.240"
    gateway    = "10.225.68.1"
    netmask    = "255.255.255.0"
    start_addr = "10.225.68.238"
  }
  hxdp_ip_range {
    end_addr   = "10.225.68.243"
    gateway    = "10.225.68.1"
    netmask    = "255.255.255.0"
    start_addr = "10.225.68.241"
  }
  node_name_prefix = "EdgeM4"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
  name = "hyperflex_node_config_policy1"
}
```