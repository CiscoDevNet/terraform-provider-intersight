resource "intersight_vmedia_policy" "tf_vmedia" {
  name          = "tf_vmedia"
  description   = "test policy"
  enabled       = true
  encryption    = true
  low_power_usb = true
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
}