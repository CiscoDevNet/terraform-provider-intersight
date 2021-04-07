### Resource Creation

```hcl
resource "intersight_firmware_eula" "firmware_eula1" {
    account {
        object_type = "iam.Account"
        moid  = intersight_account_iam.iam1.id
    }
}
```