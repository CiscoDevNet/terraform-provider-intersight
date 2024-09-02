
resource "intersight_server_profile_template" "tf_template_common" {
  name            = "tf_template_common"
  description     = "demo server profile template"
  target_platform = "FIAttached"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}