resource "intersight_smtp_policy" "smtp1" {
  enabled      = false
  name         = "smtp1"
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
    object_type = "organization.Organization"
    moid = var.organization
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

/*
SAMPLE PAYLOAD
-----------------
SmtpPolicyApi: {
    "Name": "AUTO_SMTP_POLICY_CRR",
    "Description": "testing smtp policy",
    "Tags": [{"Key": "policy", "Value": "smtp_policy"}],
    "Enabled": True,
    "SmtpServer": "10.10.10.1",
    "SmtpPort": 32,
    "MinSeverity": "critical",
    "SenderEmail": "IMCSQAAutomation@cisco.com",
    "SmtpRecipients": ["aw@cisco.com",
                       "bx@cisco.com",
                       "cy@cisco.com",
                       "dz@cisco.com"]
}
*/