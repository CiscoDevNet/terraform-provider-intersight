### Resource Creation

```hcl
resource "intersight_hyperflex_cisco_hypervisor_manager" "hyperflex_cisco_hypervisor_manager1" {
  name = "hyperflex_cisco_hypervisor_manager1"
  registered_device {
    moid        = intersight_registered_device.device1.id
    object_type = "asset.DeviceRegistrations"
  }
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
}
```