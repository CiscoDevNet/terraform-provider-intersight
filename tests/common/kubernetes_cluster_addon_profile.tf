
# resource "intersight_kubernetes_cluster_addon_profile" "iks1" {
#   name = "cluster_addon_profile1"

#   addons {
#     addon_policy {
#       object_type = "kubernetes.AddonPolicy"
#       moid = intersight_kubernetes_addon_policy.kubernetes_addon_policy1.moid
#     }
#   }
#   addons {
#     addon_policy {
#       object_type = "kubernetes.AddonPolicy"
#       moid = intersight_kubernetes_addon_policy.kubernetes_addon_policy1.moid
#     }
#   }
#   organization {
#     moid        = data.intersight_organization_organization.default.results.0.moid
#     object_type = "organization.Organization"
#   }
# }
