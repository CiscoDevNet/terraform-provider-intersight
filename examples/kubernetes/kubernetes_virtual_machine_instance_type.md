### Resource Creation

```hcl
resource "intersight_kubernetes_virtual_machine_instance_type" "kubernetes_virtual_machine_instance_type1" {
    description = "kubernetes virtual machine instance type"
    name = "kubernetes_virtual_machine_instance_type1"
    cpu = 4
    disk_size = 10
    memory = 4096
    organization {
        object_type = "organization.Organization"
        moid = var.organization
    }
}
```