data "intersight_iam_idp" "iam_idp_bool_null" {
  name = "Cisco"
}

output "out_idp_bool_null" {
    value = data.intersight_iam_idp.iam_idp_bool_null.results.0
}


data "intersight_iam_idp" "iam_idp_bool_true" {
  name = "Cisco"
  enable_single_logout = true
}

output "out_idp_bool_true" {
    value = data.intersight_iam_idp.iam_idp_bool_true.results.0
}