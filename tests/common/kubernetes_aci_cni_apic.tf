
resource "intersight_kubernetes_aci_cni_apic" "kubernetes_aci_cni_apic1" {
  num_aci_cni_profiles = 3
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
