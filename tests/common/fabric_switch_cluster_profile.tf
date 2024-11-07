
resource "intersight_fabric_switch_cluster_profile" "fabric_switch_cluster_profile1" {
  name        = "fabric_switch_cluster_profile"
  description = "demo fabric switch cluster profile"
  nr_type        = "instance"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
