### Resource Creation

```hcl
resource "intersight_virtualization_virtual_disk" "virtualization_virtual_disk1" {
  name                 = "virtualization_virtual_disk1"
  source_certs         = "<Base64 encoded CA certificates of the https source>"
  source_disk_to_clone = "<Source disk from which the content is copied>"
  source_file_path     = "<image_path>.ova"
  cluster {
    object_type = "virtualization.BaseCluster.relationship"
    moid        = var.virtualization_BaseCluster_relastionship
  }
  inventory {
    object_type = "virtualization.BaseVirtualDisk.relationship"
    moid        = var.virtualization_base_virtual_disk.relationship
  }
  registered_device {
    object_type = "asset.DeviceRegistration_relationship"
    moid        = var.DeviceRegistration_relationship
  }
  workflow_info {
    object_type = "workflow.WorkflowInfos.relationship"
    moid        = var.workflow_WorkflowInfo.relationship
  }
}
```