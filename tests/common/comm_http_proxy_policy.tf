
resource "intersight_comm_http_proxy_policy" "comm_http_proxy_policy1" {
  name        = "comm_http_proxy_policy1"
  description = "comm http proxy policy"
  hostname    = "10.10.10.1"
  password    = "ChangeMe"
  port        = 53
  username    = "admin"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
