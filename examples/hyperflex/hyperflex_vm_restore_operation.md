### Resource Creation

```hcl
resource "intersight_hyperflex_vm_restore_operation" "hyperflex_vm_restore_operation1" {
    new_name = "virtual_machine_1"
    power_on = true
    organization {
        object_type = "organization.Organization"
        moid = var.organization_organization
    }
}
```