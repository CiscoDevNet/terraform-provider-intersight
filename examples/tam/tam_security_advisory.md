### Resource Creation

```hcl
resource "intersight_tam_security_advisory" "tam_security_advisory1" {
  name           = "tam_security_advisories1"
  operation_type = "create"
  state          = "ready"
  severity = {
    object_type = "tam.SecurityAdvisoryDetails"
  }
  actions = {
    object_type = "tam.SecurityAdvisoryDetails"
    alert_type  = "psirt"
  }
  type   = "psirt"
  status = "final"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```