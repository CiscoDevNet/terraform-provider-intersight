# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_active_true" {
#   name            = "Provisioned IOPS"
#   is_active = true
# }

# output "cloud_sku_volume_active_true" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_active_true.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_autoDiscovered_false" {
#   name            = "Provisioned IOPS"
#   is_auto_discovered = false
# }

# output "cloud_sku_volume_autoDiscovered_false" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_autoDiscovered_false.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_isBootable_true" {
#   name            = "Provisioned IOPS"
#   is_bootable = true
# }

# output "cloud_sku_volume_isBootable_true" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_isBootable_true.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_isDefault_false" {
#   name            = "Provisioned IOPS"
#   is_default = false
# }

# output "cloud_sku_volume_isDefault_false" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_isDefault_false.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_isProvisionedIops_true" {
#   name            = "Provisioned IOPS"
#   is_provisioned_iops = true
# }

# output "cloud_sku_volume_isProvisionedIops_true" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_isProvisionedIops_true.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume_active_true_autoDiscovered_false_isBootable_false_isDefault_false_isProvisionedIOPS_false" {
#   name            = "Provisioned IOPS"
#   is_active       = true
#   is_bootable = true
#   is_auto_discovered = false
#   is_default = false
#   is_provisioned_iops = true
# }

# output "cloud_sku_volume_active_true_autoDiscovered_false_isBootable_false_isDefault_false_isProvisionedIOPS_false" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume_active_true_autoDiscovered_false_isBootable_false_isDefault_false_isProvisionedIOPS_false.results.0
# }


# data "intersight_cloud_sku_volume_type" "cloud_sku_volume" {
#  name            = "Provisioned IOPS"
# }

# output "cloud_sku_volume" {
#   value = data.intersight_cloud_sku_volume_type.cloud_sku_volume.results.0
# }