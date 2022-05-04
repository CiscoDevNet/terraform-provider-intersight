resource "intersight_sol_policy" "tf_sol" {
  name      = "tf_sol"
  enabled   = false
  baud_rate = 9600
  com_port  = "com1"
  ssh_port  = 1096
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  tags {
    key   = "source"
    value = "terraform"
  }
}