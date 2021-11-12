### Resource Creation

```hcl
resource "intersight_vmrc_console" "vmrc_console1" {
   session {
     moid        = var.iam_session
     object_type = "iam.Session"
   }
  virtual_machine {
    object_type = "virtualization.VmwareVirtualMachine"
  }
}

 variable "iam_session" {
   type = string
   description = "moid for iam session"
 }
```