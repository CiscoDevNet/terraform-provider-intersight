### Resource Creation
```hcl
resource "intersight_virtualization_virtual_machine" "vm-test-tf-is" {
  name            = "vm-test-tf"
  action          = "Create"
  cpu             = 2
  memory          = 2048
  hypervisor_type = "ESXi"
  guest_os        = "linux"
  power_state     = "PowerOn"
  cluster {
      object_type           = "virtualization.VmwareCluster"
      moid                  = "changeMe"
    }
  vm_config {
    additional_properties : jsonencode({
      Compute = {
        ResourcePool = ""
        Annotation = ""
        ObjectType = "virtualization.EsxiVmComputeConfiguration"
      }
      Customspec = null
      Datacenter = "changeMe"
      Folder = ""
      Image = ""
      Network = {
        Interfaces = []
        ObjectType = "virtualization.EsxiVmNetworkConfiguration"
      }
      Storage = {
        Datastore = "changeMe"
        Disks = []
        ObjectType = "virtualization.EsxiVmStorageConfiguration"
      }
      Template = "changeMe"
    })
  }
  provision_type = "Template"
  cluster_esxi = "RMLAB - Workloads"
}
```