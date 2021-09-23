data "intersight_capability_sioc_module_capability_def" "capability_sioc_module_dcSupported_false" {
  name            = "sioc0"
  dc_supported       = false
}

output "capability_sioc_module_dcSupported_false" {
  value = data.intersight_capability_sioc_module_capability_def.capability_sioc_module_dcSupported_false.results.0
}


data "intersight_capability_sioc_module_capability_def" "capability_sioc_module_dcSupported_null" {
  name            = "sioc0"
}

output "capability_sioc_module_dcSupported_null" {
  value = data.intersight_capability_sioc_module_capability_def.capability_sioc_module_dcSupported_null.results.0
}