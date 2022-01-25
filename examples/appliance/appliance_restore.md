### Resource Creation

```hcl
resource "intersight_appliance_restore" "appliance_restore1" {
  filename    = "name_of_the_file.txt"
  protocol    = "SFTP"
  remote_host = "appliance-remote.hostname.com"
  remote_path = "path/to/the/remote/host"
  remote_port = 22
  username    = "authenicate_user"
  password    = "ChangeMe"
  messages = [
    "Message generated during restore process"
  ]
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