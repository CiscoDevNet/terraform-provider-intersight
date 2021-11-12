### Resource Creation

```hcl
resource "intersight_tam_advisory_definition" "tam_advisory_definition" {
  name  = "tam_advisory_definition"
  state = "ready"
  severity = [{
    object_type           = "tam.SecurityAdvisoryDetails"
    additional_properties = ""
    class_id              = ""
  }]
  actions = [{
    additional_properties = ""
    class_id              = ""
    operation_type        = "create"
    name                  = "tam_security_advisories1"
    queries               = null
    type                  = "restApi"
    affected_object_type  = ""
    identifiers           = null
    object_type           = "tam.SecurityAdvisoryDetails"
    alert_type            = "psirt"
  }]
  type = "securityAdvisory"
  api_data_sources {
    object_type = "tam.ApiDataSource"
    name        = "api_data_source_1"
    type        = "intersightApi"

  }
  advisory_details = [{
    object_type           = "tam.SecurityAdvisoryDetails"
    description           = "tam security advisory"
    additional_properties = ""
    class_id              = "tam.SecurityAdvisoryDetails"
  }]

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