### Resource Creation

```hcl
resource "intersight_workflow_error_response_handler" "workflow_error_response_handler1" {
  name          = "workflow_error_response_handler1"
  platform_type = "UCSD"
  parameters = [{
    object_type         = "content.TextParameter"
    accept_single_value = true
    name                = "show-pure"
    item_type           = "string"

  }]
  types {
    object_type = "content.ComplexType"
    parameters = [{
      object_type = "content.TextParameter"
      name        = "show-hitachi"
      item_type   = "string"
      type        = "string"
    }]
  }
  catalog {
    object_type = "workflow.Catalog"
    moid        = var.workflow_catalog
  }

}
```