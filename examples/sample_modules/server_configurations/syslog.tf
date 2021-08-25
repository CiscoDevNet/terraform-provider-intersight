resource "intersight_syslog_policy" "tf_syslog" {
  name        = "tf_syslog"
  description = "demo syslog policy"
  local_clients {
    min_severity = "emergency"
    object_type  = "syslog.LocalFileLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "10.10.10.10"
    port         = 514
    protocol     = "tcp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  remote_clients {
    enabled      = true
    hostname     = "2001:0db8:0a0b:12f0:0000:0000:0000:0004"
    port         = 64000
    protocol     = "udp"
    min_severity = "emergency"
    object_type  = "syslog.RemoteLoggingClient"
  }
  organization {
    moid = data.intersight_organization_organization.default.results.0.moid
  }
}

/*
SAMPLE PAYLOAD
-----------------
SyslogPolicyApi: {
  "Name": "AUTO_POLICY_SYSLOG_CRR",
  "Description": "Syslog Enable Testing",
  "Tags": [{"Key": "syslogenable", "Value": "111"}],
  "LocalClients": [{"MinSeverity": "emergency", "ObjectType": "syslog.LocalFileLoggingClient"}
  ],
  "RemoteClients": [{"Enabled": True, "Hostname": "10.10.10.1", "Port": 514,
    "Protocol": "tcp", "MinSeverity": "emergency", "ObjectType":
    "syslog.RemoteLoggingClient"},
    {"Enabled": True, "Hostname": "2001:0db8:0a0b:12f0:0000:0000:0000:0004", "Port": 64000,
      "Protocol": "udp", "MinSeverity": "emergency", "ObjectType": "syslog.RemoteLoggingClient"}

  ]
}
*/