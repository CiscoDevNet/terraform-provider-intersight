### Resource Creation

```hcl
resource "intersight_appliance_backup" "appliance_backup1" {
  filename    = "default_filename1"
  protocol    = "scp"
  remote_host = "host.remote"
  remote_path = "path/to/remote/host"
  remote_port = 0
  username    = "user_1"
  password    = "ChangeMe"
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