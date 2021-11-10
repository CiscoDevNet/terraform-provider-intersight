
resource "intersight_ntp_policy" "ntp1" {
  name        = "ntp1"
  description = "test policy"
  enabled     = true
  ntp_servers = [
    "ntp.esl.cisco.com",
    "time-a-g.nist.gov",
    "time-b-g.nist.gov"
  ]
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
