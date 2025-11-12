### Resource Creation

```hcl
resource "intersight_workflow_error_response_handler" "workflow_error_response_handler1" {
  name          = "workflow_error_response_handler1"
  platform_type = "UCSD"
  parameters {
    class_id            = "content.TextParameter"
    secure              = false
    accept_single_value = true
    object_type         = "content.TextParameter"
    name                = "show-pure"
    item_type           = "string"
    nr_type                = "string"
  }
  types {
    object_type = "content.ComplexType"
    parameters {
      class_id            = "content.TextParameter"
      secure              = false
      accept_single_value = true
      object_type         = "content.TextParameter"
      name                = "show-hitachi"
      item_type           = "string"
      nr_type                = "string"
    }
  }
  catalog {
    object_type = "workflow.Catalog"
    moid        = var.workflow_catalog
  }

}

variable "workflow_catalog" {
  type        = string
  description = "moid for workflow catalog"
}
```
