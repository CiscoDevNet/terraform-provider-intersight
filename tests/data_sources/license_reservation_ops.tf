data "intersight_license_license_reservation_op" "auth_code_installed_false" {
  auth_code_installed = false
}

output "auth_code_installed_false" {
  value = data.intersight_license_license_reservation_op.auth_code_installed_false.results.0
}

data "intersight_license_license_reservation_op" "license_reservation_op" {
}

output "license_reservation_op" {
  value = data.intersight_license_license_reservation_op.license_reservation_op.results.0
}
