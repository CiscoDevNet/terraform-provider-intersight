### Resource Creation

```hcl
resource "intersight_virtualization_virtual_machine" "virtualization_virtual_machine1" {
    name = "virtualization_virtual_machine1"
    power_state = "PowerOn"
    provision_type = "OVA"
    cpu = 16
    memory = 3755356
    _0_virtual_machine_node_profile {
        object_type = "kubernetes.VirtualMachineInfrastructureProvider"
        moid = var.kubernetes_VirtualMachineInfrastructureProvider
    }
    host {
        object_type = "virtualization.BaseHost.relationship"
        moid = var.virtualization_base_host.relationship
    }
    inventory {
        object_type = "virtualization.BaseVirtualMachine.relationship"
        moid = var.virtualization_base_virtual_machine.relationship
    }
    registered_device {
        object_type = "asset.DeviceRegistration_relationship"
        moid = var.DeviceRegistration_relationship
    }
    workflow_info {
        object_type = "workflow.WorkflowInfos.relationship"
        moid = var.workflow_WorkflowInfo.relationship
    }
}
```