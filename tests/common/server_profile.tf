
resource "intersight_server_profile" "server1" {
  name   = "server1"
  target_platform = "Standalone"
  server_family   = "UCSC2XX/4XX"
  action = "No-op"
  tags {
    key   = "server"
    value = "demo"
  }
  tags {
    key   = "project"
    value = "cloud_migration"
  }
  tags {
    key   = "Environment"
    value = "production"
  }
   tags {
    key   = "Application"
    value = "WebService"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
