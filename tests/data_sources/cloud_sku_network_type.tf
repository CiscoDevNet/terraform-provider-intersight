data "intersight_cloud_sku_network_type" "cloud_sku_network_active_true_autoDiscovered_null" {
  name            = "Load Balancer-Network"
  is_active = true
}

output "cloud_sku_network_active_true_autoDiscovered_null" {
  value = data.intersight_cloud_sku_network_type.cloud_sku_network_active_true_autoDiscovered_null.results.0
}


data "intersight_cloud_sku_network_type" "cloud_sku_network_active_true_autoDiscovered_false" {
  name            = "Load Balancer-Network"
  is_active       = true
  is_auto_discovered = false
}

output "cloud_sku_network_active_true_autoDiscovered_false" {
  value = data.intersight_cloud_sku_network_type.cloud_sku_network_active_true_autoDiscovered_false.results.0
}


data "intersight_cloud_sku_network_type" "cloud_sku_network_active_autoDiscovered_null" {
 name            = "Load Balancer-Network"
}

output "cloud_sku_network_active_autoDiscovered_null" {
  value = data.intersight_cloud_sku_network_type.cloud_sku_network_active_autoDiscovered_null.results.0
}


data "intersight_cloud_sku_network_type" "cloud_sku_network_active_null_autoDiscovered_false" {
  name            = "Load Balancer-Network"
  is_auto_discovered = false
}

output "cloud_sku_network_active_null_autoDiscovered_false" {
  value = data.intersight_cloud_sku_network_type.cloud_sku_network_active_null_autoDiscovered_false.results.0
}