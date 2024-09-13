data "intersight_organization_organization" "organization" {
  name = "default"
}

output "default_organization" {
  value = data.intersight_organization_organization.organization.results.0
}

data "intersight_softwarerepository_catalog" "sw_catalog" {
  organization {
    moid = data.intersight_organization_organization.organization.id
  }
  account_moid = data.intersight_organization_organization.organization.results.0.account_moid
}

output "catalog_in_default_organization" {
  value = data.intersight_softwarerepository_catalog.sw_catalog.results.0
}

data "intersight_search_search_item" "os_linux" {
  object_type = "compute.PhysicalSummary"
}

output "search_compute_physical" {
  value = data.intersight_search_search_item.os_linux.results.0
}