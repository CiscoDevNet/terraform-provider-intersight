### Resource Creation

```hcl
resource "intersight_hyperflex_local_credential_policy" "hyperflex_local_credential_policy1" {
  hxdp_root_pwd               = "ChangeMe@123"
  hypervisor_admin            = "admin"
  hypervisor_admin_pwd        = "ChangeMe"
  factory_hypervisor_password = false
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  name = "hyperflex_local_credential_policy1"
}
```