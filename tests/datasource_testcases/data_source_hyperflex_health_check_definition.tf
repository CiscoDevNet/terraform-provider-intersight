data "intersight_hyperflex_health_check_definition" "script_execution_on_compute_nodes_false" {
  name            = "NTP Time Drift from Reference Clock"
  script_execution_on_compute_nodes = false
}

output "script_execution_on_compute_nodes_false" {
  value = data.intersight_hyperflex_health_check_definition.script_execution_on_compute_nodes_false.results.0
}


data "intersight_hyperflex_health_check_definition" "script_execution_on_compute_nodes_null" {
  name            = "NTP Time Drift from Reference Clock"
}

output "script_execution_on_compute_nodes_null" {
  value = data.intersight_hyperflex_health_check_definition.script_execution_on_compute_nodes_null.results.0
}