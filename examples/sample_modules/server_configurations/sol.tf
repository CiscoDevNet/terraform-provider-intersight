resource "intersight_sol_policy" "sol1" {
  name      = "sol2"
  enabled   = false
  baud_rate = 9600
  com_port  = "com1"
  ssh_port  = 1096
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

/*
SAMPLE PAYLOAD
-----------------
SolPolicyApi: {
  "Name": "AUTO_SOL_POLICY_CRR",
  "Description": "SOL Policy Test",
  "Tags": [{"Key": "policy", "Value": "sol_policy"}],
  "Enabled": True,
  "BaudRate": 9600,
  "ComPort": 'com1',
  "SshPort": 2024
}
*/