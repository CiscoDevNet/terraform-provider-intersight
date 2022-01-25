### Resource Creation

```hcl
resource "intersight_hyperflex_proxy_setting_policy" "hyperflex_proxy_setting_policy1" {
  hostname    = "10.10.10.1"
  port        = 32628
  username    = ""
  password    = "ChangeMe"
  description = "This is autoProxy"
  tags {
    key   = "test"
    value = "autoProxy"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  name = "hyperflex_proxy_setting_policy1"
}
```