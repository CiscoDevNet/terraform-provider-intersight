
resource "intersight_vnic_san_connectivity_policy" "vnic_san1" {
  name = "vnic_san1"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  profiles {
    moid        = intersight_server_profile.tf_server_common.id
    object_type = "server.Profile"
  }
}
