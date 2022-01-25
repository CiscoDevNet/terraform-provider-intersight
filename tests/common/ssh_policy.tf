
resource "intersight_ssh_policy" "ssh_policy1" {
  name        = "ssh_policy1"
  description = "ssh policy"
  enabled     = true
  port        = 22
  timeout     = 1800
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
