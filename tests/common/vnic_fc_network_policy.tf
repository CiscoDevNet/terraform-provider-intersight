
resource "intersight_vnic_fc_network_policy" "v_fc_network1" {
  name = "v_fc_network1"
  vsan_settings {
    id = 100
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
