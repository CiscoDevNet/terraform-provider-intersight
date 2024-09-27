resource "intersight_kubernetes_cluster_profile" "kcp1" {
  name = "dummy_kcp_tf"
  organization {
    object_type = "organization.Organization"
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

data "intersight_kubernetes_cluster_profile" "kcp1" {
  depends_on = [
    intersight_kubernetes_cluster_profile.kcp1
  ]
  name = "dummy_kcp_tf"
}

resource "intersight_kubernetes_addon_policy" "kap1" {
  name = "dummy-kap1"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_kubernetes_addon_policy" "kap2" {
  name = "dummy-kap2"
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}

resource "intersight_kubernetes_cluster_addon_profile" "iks-terra-addons" {
  name = "iks-terra-addons"
  addons {
    addon_policy {
      object_type = "kubernetes.AddonPolicy"
      moid = intersight_kubernetes_addon_policy.kap1.moid
    }
  }
  addons {
    addon_policy {
      object_type = "kubernetes.AddonPolicy"
      moid = intersight_kubernetes_addon_policy.kap2.moid
    }
  }
  organization {
    object_type = "organization.Organization"
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}