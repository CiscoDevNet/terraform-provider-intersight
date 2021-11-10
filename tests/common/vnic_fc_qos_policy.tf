
resource "intersight_vnic_fc_qos_policy" "v_fc_qos1" {
  name                = "v_fc_qos1"
  rate_limit          = 10000
  cos                 = 6
  max_data_field_size = 2112
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
