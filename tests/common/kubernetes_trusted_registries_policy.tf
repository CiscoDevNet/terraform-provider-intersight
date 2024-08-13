
resource "intersight_kubernetes_trusted_registries_policy" "kubernetes_trusted_registries_policy1" {
  description = "kubernetes trusted registries policy"
  name        = "kubernetes_trusted_registries_policy1"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}