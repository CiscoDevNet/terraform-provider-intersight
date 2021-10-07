data "intersight_license_license_reservation_op" "auth_code_installed_false" {
  auth_code_installed = false
}

output "auth_code_installed_false" {
  value = data.intersight_license_license_reservation_op.auth_code_installed_false.results.0
}


data "intersight_license_license_reservation_op" "generate_request_code_false" {
  generate_request_code = false
}

output "generate_request_code_false" {
  value = data.intersight_license_license_reservation_op.generate_request_code_false.results.0
}


data "intersight_license_license_reservation_op" "generate_return_code_false" {
  generate_return_code = false
}

output "generate_return_code_false" {
  value = data.intersight_license_license_reservation_op.generate_return_code_false.results.0
}


data "intersight_license_license_reservation_op" "license_reservation_op" {
}

output "license_reservation_op" {
  value = data.intersight_license_license_reservation_op.license_reservation_op.results.0
}