resource "intersight_smtp_policy" "tf_smtp" {
  enabled      = false
  name         = "tf_smtp"
  description  = "testing smtp policy"
  smtp_port    = 32
  min_severity = "critical"
  smtp_server  = "10.10.10.1"
  sender_email = "IMCSQAAutomation@cisco.com"
  smtp_recipients = [
    "aw@cisco.com",
    "cy@cisco.com",
  "dz@cisco.com"]
  organization {
    moid        = data.intersight_organization_organization.default.results.0.moid
    object_type = "organization.Organization"
  }
  tags {
    key   = "source"
    value = "terraform"
  }
}