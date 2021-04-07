### Resource Creation

```hcl
resource "intersight_vmrc_console" "vmrc_console1" {
    session {
        moid = var.iam_session
        object_type = "iam.Session"
    }
    virtual_machine {
        object_type = "vmrc.Console"
        hypervisor_type = "Hyper-V"
        power_state = "PoweredOn"
        dhcp_enabled = true
        guest_state = "Running"
        is_template = false
    }
}
```