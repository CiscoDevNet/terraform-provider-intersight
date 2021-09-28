data "intersight_cloud_sku_database_type" "cloud_sku_database_active_true_autoDiscovered_null" {
  name            = "db.m3.large"
  is_active = true
}

output "cloud_sku_database_active_true_autoDiscovered_false" {
  value = data.intersight_cloud_sku_database_type.cloud_sku_database_active_true_autoDiscovered_false.results.0
}


data "intersight_cloud_sku_database_type" "cloud_sku_database_active_true_autoDiscovered_false" {
  name            = "db.m3.large"
  is_active       = true
  is_auto_discovered = false
}

output "cloud_sku_database_active_true_autoDiscovered_null" {
  value = data.intersight_cloud_sku_database_type.cloud_sku_database_active_true_autoDiscovered_null.results.0
}


data "intersight_cloud_sku_database_type" "cloud_sku_database_active_autoDiscovered_null" {
 name            = "db.m3.large"
}

output "cloud_sku_database_active_autoDiscovered_null" {
  value = data.intersight_cloud_sku_database_type.cloud_sku_database_active_autoDiscovered_null.results.0
}


data "intersight_cloud_sku_database_type" "cloud_sku_database_active_null_autoDiscovered_false" {
  name            = "db.m3.large"
  is_auto_discovered = false
}

output "cloud_sku_database_active_null_autoDiscovered_false" {
  value = data.intersight_cloud_sku_database_type.cloud_sku_database_active_null_autoDiscovered_false.results.0
}












