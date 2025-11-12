
# resource "intersight_kubernetes_network_policy" "kubernetes_network_policy1" {
#   description = "kubernetes network policy"
#   name        = "kubernetes_network_policy1"
#   cni_type    = "Calico"
#   organization {
#     object_type = "organization.Organization"
#     moid        = data.intersight_organization_organization.default.results.0.moid
#   }
# }
