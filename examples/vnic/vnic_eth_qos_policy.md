### Resource Creation

```hcl
resource "intersight_vnic_eth_qos_policy" "v_eth_qos1" {
  name           = "v_eth_qos1"
  mtu            = 1500
  rate_limit     = 0
  cos            = 0
  trust_host_cos = false
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```