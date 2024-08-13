
resource "intersight_macpool_pool" "macpool_pool1" {
  name = "AUTO_test_macpool"
  mac_blocks {
    object_type = "macpool.Block"
    from        = "00:25:B5:9d:68:16"
    to          = "00:25:B5:9d:68:40"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
