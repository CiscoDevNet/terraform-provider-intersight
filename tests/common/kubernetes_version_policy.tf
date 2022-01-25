
resource "intersight_kubernetes_version_policy" "kubernetes_version_policy1" {
  description = "kubernetes version policy"
  name        = "kubernetes_version_policy1"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
