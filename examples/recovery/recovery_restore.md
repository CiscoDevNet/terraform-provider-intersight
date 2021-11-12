### Resource Creation

```hcl
resource "intersight_recovery_restore" "recovery_restore1" {
   device {
     object_type = "asset.DeviceRegistration"
     moid        = var.asset_device_registtration
   }
   workflow_info {
     object_type = "workflow.WorkflowInfos.relationship"
     moid        = var.workflow_WorkflowInfo_relationship
   }
   organization {
     object_type = "organization.Organization"
     moid        = var.organization
   }

}
variable "workflow_WorkflowInfo_relationship" {
   type = string
   description = "value for workflow_WorkflowInfo_relationship"
 }

 variable "organization" {
   type = string
   description = "value for organization"
 }

 variable "asset_device_registtration" {
   type = string
   description = "value for asset_device_registtration"
 }
```