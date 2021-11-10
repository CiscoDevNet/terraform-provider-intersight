
resource "intersight_ipmioverlan_policy" "ipmi1" {
  name        = "ipmi1"
  description = "demo ipmi policy"
  enabled     = true
  privilege   = "admin"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}