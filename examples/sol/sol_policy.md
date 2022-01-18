### Resource Creation

```hcl
resource "intersight_sol_policy" "sol1" {
  name        = "sol2"
  description = "demo serial over lan policy"
  enabled     = false
  baud_rate   = 9600
  com_port    = "com1"
  ssh_port    = 1096
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
variable "organization" {
   type = string
   description = "<value for organization>"
 }
```