data "intersight_iam_idp_reference" "idp_reference_multifactor_false" {
  name            = "Cisco"
  multi_factor_authentication       = false
}

output "idp_reference_multifactor_false" {
  value = data.intersight_iam_idp_reference.idp_reference_multifactor_false.results.0
}

data "intersight_iam_idp_reference" "idp_reference_multifactor_null" {
  name            = "Cisco"
}

output "idp_reference_multifactor_null" {
  value = data.intersight_iam_idp_reference.idp_reference_multifactor_null.results.0
}
