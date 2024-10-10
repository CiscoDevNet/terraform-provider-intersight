
resource "intersight_server_profile_template" "template1" {
  name            = "server_profile_template1"
  description     = "demo server profile template"
  target_platform = "FIAttached"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
