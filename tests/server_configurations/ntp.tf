resource "intersight_ntp_policy" "tf_ntp1" {
  name    = "tf_ntp1"
  enabled = true
  ntp_servers = [
    "ntp.esl.cisco.com",
    "time-a-g.nist.gov",
    "time-b-g.nist.gov"
  ]
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

resource "intersight_ntp_policy" "tf_ntp2" {
  name    = "tf_ntp2"
  enabled = true
  ntp_servers = [
    "10.10.10.10",
    "10.10.10.11",
    "10.10.10.12",
    "10.10.10.13"
  ]
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
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