
resource "intersight_kubernetes_addon_policy" "kubernetes_addon_policy1" {
  description = "kubernetes addon policy"
  name        = "kubernetes_addon_policy1"
  addon_configuration {
    install_strategy = "Always"
    release_name     = "helm-release-3"
    upgrade_strategy = "UpgradeOnly"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
