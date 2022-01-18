### Resource Creation

```hcl
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
    moid        = var.organization
  }
  profiles {
    moid        = var.profile
    object_type = "server.Profile"
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }

variable "profile"{
  type = string
  description = "Moid of server.Profile"
}
```