### Resource Creation

```hcl
resource "intersight_connectorpack_connector_pack_upgrade" "connectorpack_connector_pack_upgrade1" {
  connector_pack_op_type = "Install"
  workflow_info {
    object_type = "workflow.WorkflowInfo"
    moid        = var.workflow_workflow_info
  }
}
```