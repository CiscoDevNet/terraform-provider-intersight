### Resource Creation

```hcl
resource "intersight_recovery_restore" "recovery_restore1" {
    device {
        object_type = "asset.DeviceRegistration"
        moid = var.asset_device_registtration
    }
    workflow_info {
        object_type = "workflow.WorkflowInfos.relationship"
        moid = var.workflow_WorkflowInfo.relationship
    }
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }

}
```