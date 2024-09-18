
resource "intersight_server_profile" "server1" {
  name   = "server1"
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
