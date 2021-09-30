data "intersight_softwarerepository_category_support_constraint" "parse_from_image_name_true" {
  name = "UcsmModel"
  parse_from_image_name = true
}

output "parse_from_image_name_true" {
  value = data.intersight_softwarerepository_category_support_constraint.parse_from_image_name_true.results.0
}


data "intersight_softwarerepository_category_support_constraint" "software_repository_category_support_constraint" {
  name = "UcsmModel"
}

output "software_repository_category_support_constraint" {
  value = data.intersight_softwarerepository_category_support_constraint.software_repository_category_support_constraint.results.0
}