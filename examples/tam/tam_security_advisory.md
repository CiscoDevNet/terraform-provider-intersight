### Resource Creation

```hcl
resource "intersight_tam_security_advisory" "tam_security_advisory1" {
  name  = "tam_security_advisories1"
  state = "ready"

  severity = [{
    object_type           = "tam.SecurityAdvisoryDetails"
    additional_properties = ""
    class_id              = "tam.SecurityAdvisoryDetails"
  }]

  actions = [{
    additional_properties = ""
    class_id              = "tam.SecurityAdvisoryDetails"
    operation_type        = "create"
    name                  = "tam_security_advisories1"
    queries               = null
    object_type           = "tam.SecurityAdvisoryDetails"
    alert_type            = "psirt"
    type                  = "restApi"
    affected_object_type  = ""
    identifiers           = null
  }]

  status = "final"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

 variable "organization" {
   type =  string
   description = "value for moid"
 }
```