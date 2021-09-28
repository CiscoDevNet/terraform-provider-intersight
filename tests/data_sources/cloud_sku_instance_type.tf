data "intersight_cloud_sku_instance_type" "cloud_sku_instance_active_true" {
  name            = "a1.xlarge"
  is_active = true
}

output "cloud_sku_instance_active_true" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_active_true.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_autoDiscovered_false" {
  name            = "a1.xlarge"
  is_auto_discovered = false
}

output "cloud_sku_instance_autoDiscovered_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_autoDiscovered_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_cudaSupport_false" {
  name            = "a1.xlarge"
  cuda_support = false
}

output "cloud_sku_instance_cudaSupport_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_cudaSupport_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_active_true_autoDiscovered_false" {
  name            = "a1.xlarge"
  is_active       = true
  is_auto_discovered = false
}

output "cloud_sku_instance_active_true_autoDiscovered_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_active_true_autoDiscovered_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_active_true_cudaSupport_false" {
  name            = "a1.xlarge"
  is_active       = true
  cuda_support = false
}

output "cloud_sku_instance_active_true_cudaSupport_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_active_true_cudaSupport_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_cudaSupport_false_autoDiscovered_false" {
  name            = "a1.xlarge"
  cuda_support = false
  is_auto_discovered = false
}

output "cloud_sku_instance_cudaSupport_false_autoDiscovered_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_cudaSupport_false_autoDiscovered_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance_active_true_autoDiscovered_false_cudaSupport_false" {
  name            = "a1.xlarge"
  is_active       = true
  cuda_support = false
  is_auto_discovered = false
}

output "cloud_sku_instance_active_true_autoDiscovered_false_cudaSupport_false" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance_active_true_autoDiscovered_false_cudaSupport_false.results.0
}


data "intersight_cloud_sku_instance_type" "cloud_sku_instance" {
 name            = "a1.xlarge"
}

output "cloud_sku_instance" {
  value = data.intersight_cloud_sku_instance_type.cloud_sku_instance.results.0
}

