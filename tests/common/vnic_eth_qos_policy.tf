
resource "intersight_vnic_eth_qos_policy" "v_eth_qos1" {
  name           = "v_eth_qos1"
  description    = "demo vnic eth qos policy"
  mtu            = 1500
  rate_limit     = 0
  cos            = 0
  burst          = 1024
  priority       = "Best Effort"
  trust_host_cos = false
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
