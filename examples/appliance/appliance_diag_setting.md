### Resource Creation

```hcl
resource "intersight_appliance_diag_setting" "appliance_diag_setting1" {
  name                        = "appliance_diag_setting1"
  device_id = "WZP23350KMU"
  hostname = "10.106.233.221"
  is_password_set = true
  password = "xxxxxxxx"
  account {
        object_type = "iam.Account"
        moid = intersight_account_iam.iam1.id
  }
}
```