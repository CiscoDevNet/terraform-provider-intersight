### Resource Creation

```hcl
resource "intersight_virtualization_virtual_machine" "virtualization_virtual_machine1" {
  name           = "virtualization_virtual_machine1"
  power_state    = "PowerOn"
  provision_type = "OVA"
  cpu            = 16
  memory         = 3755356
   host {
     object_type = "virtualization.BaseHost.relationship"
     moid        = var.virtualization_base_host_relationship
   }
   inventory {
     object_type = "virtualization.BaseVirtualMachine.relationship"
     moid        = var.virtualization_base_virtual_machine_relationship
   }
   registered_device {
     object_type = "asset.DeviceRegistration.relationship"
     moid        = var.DeviceRegistration_relationship
   }
   workflow_info {
     object_type = "workflow.WorkflowInfos.relationship"
     moid        = var.workflow_WorkflowInfo_relationship
   }
}

 variable "kubernetes_VirtualMachineInfrastructureProvider" {
   type = string
   description = "moid for kubernetes virtaul machine infrastructure provider"
 }

 variable "virtualization_base_host_relationship" {
   type = string
   description = "moid for kubernetes virtualization  base host relationship"
 }

 variable "virtualization_base_virtual_machine_relationship" {
   type = string
   description = "moid for  virtualization base virtualization machine relationship"
 }

 variable "DeviceRegistration_relationship" {
   type = string
   description = "moid for device registration relationship"
 }

 variable "workflow_WorkflowInfo_relationship" {
   type = string
   description = "moid for workflow workflowInfo relationship"
 }
```