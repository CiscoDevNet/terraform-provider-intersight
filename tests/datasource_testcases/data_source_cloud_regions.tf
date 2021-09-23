data "intersight_cloud_regions" "cloud_region_active_billing_true" {
  name            = "EU (Milan)"
  is_active       = true
  is_billing_only = true
}

output "cloud_region_active_billing_true" {
  value = data.intersight_cloud_regions.cloud_region_active_billing_true.results.0
}


data "intersight_cloud_regions" "cloud_region_active_null_biiling_true" {
  name            = "EU (Milan)"
  is_billing_only = true
}

output "cloud_region_active_null_biiling_true" {
  value = data.intersight_cloud_regions.cloud_region_active_null_biiling_true.results.0
}


data "intersight_cloud_regions" "cloud_region_active_true_billing_null" {
  name      = "EU (Milan)"
  is_active = true
}

output "cloud_region_active_true_billing_null" {
  value = data.intersight_cloud_regions.cloud_region_active_true_billing_null.results.0
}


data "intersight_cloud_regions" "cloud_region_active_billing_null" {
  name = "EU (Milan)"
}

output "cloud_region_active_billing_null" {
  value = data.intersight_cloud_regions.cloud_region_active_billing_null.results.0
}








