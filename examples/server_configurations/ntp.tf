resource "intersight_ntp_policy" "ntp1" {
  name    = "ntp1"
  enabled = true
  ntp_servers = [
    "ntp.esl.cisco.com",
    "time-a-g.nist.gov",
    "time-b-g.nist.gov"
  ]
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

resource "intersight_ntp_policy" "ntp2" {
  name    = "ntp2"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
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
NtpPolicyApi: {
  "Name": "AUTO_NTP_POLICY_CRR",
  "Description": "NTP Policy Test",
  "Tags": [{"Key": "policy", "Value": "ntp_policy"}],
  "Enabled": True,
  "NtpServers": [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
}
*/