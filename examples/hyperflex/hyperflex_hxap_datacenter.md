### Resource Creation

```hcl
resource "intersight_hyperflex_hxap_datacenter" "hyperflex_hxap_datacenter1" {
  name = "hyperflex_hxap_datacenter1"
  account {
    object_type = "iam.Account"
    moid        = intersight_iam_account.account1.id
  }
  registered_device {
    moid        = intersight_registered_device.device1.id
    object_type = "asset.DeviceRegistrations"
  }
}
```