
resource "intersight_hyperflex_auto_support_policy" "hyperflex_auto_support_policy1" {
  admin_state               = true
  service_ticket_receipient = "test@example.com"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  name = "hyperflex_auto_support_policy1"
}
