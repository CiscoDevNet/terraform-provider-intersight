
resource "intersight_syslog_policy" "syslog1" {
  name        = "syslog1"
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
  profiles {
    moid        = intersight_server_profile.tf_server_common.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
}
