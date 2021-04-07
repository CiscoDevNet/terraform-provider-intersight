### Resource Creation

```hcl
resource "intersight_workflow_custom_data_type_definition" "workflow_custom_data_type_definition1" {
  name        = "workflow_custom_data_type_definition1"
  description = "captures a customized data type definition for workflow input/output"
  label       = "custom_data_dfn1"
  catalog {
    object_type = "workflow.Catalog"
    moid        = var.workflow_catalog
  }
  parameter_set = [{
    object_type       = "workflow.ParameterSet"
    name              = "show-netapp"
    condition         = "eq"
    control_parameter = "PlatformType"
    enable_parameters = [
      "ExpandedVolumeCapacity"
    ]
    value = "storage.NetAppCluster"

  }]
}
```