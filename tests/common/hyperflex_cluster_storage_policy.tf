
resource "intersight_hyperflex_cluster_storage_policy" "hyperflex_cluster_storage_policy1" {
  disk_partition_cleanup = true
  vdi_optimization       = true
  logical_avalability_zone_config {
    auto_config = false
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  name = "hyperflex_cluster_storage_policy1"
}
