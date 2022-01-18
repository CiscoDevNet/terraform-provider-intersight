### Resource Creation

```hcl
resource "intersight_appliance_remote_file_import" "appliance_remote_file_import1" {
  filename = "appliance_file.txt"
  hostname = "localhost"
  password = "ChangeMe"
  path     = "/remote/host/"
  port     = 22
  protocol = "sftp"
  username = "admin"
  account {
    object_type = "iam.Account"
    moid        = var.account
  }
}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```