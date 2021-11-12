### Resource Creation

```hcl
resource "intersight_hyperflex_cisco_hypervisor_manager" "hyperflex_cisco_hypervisor_manager1" {
  name = "hyperflex_cisco_hypervisor_manager1"
  account {
    object_type = "iam.Account"
    moid        = var.iam_account
  }
}

variable "iam_account" {
  type        = string
  description = "Moid of Iam account"
}
```