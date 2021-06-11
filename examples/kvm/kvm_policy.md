### Resource Creation

```hcl
resource "intersight_kvm_policy" "kvm1" {
  name                      = "kvm1"
  description               = "demo kvm policy"
  enabled                   = true
  maximum_sessions          = 3
  remote_port               = 2069
  enable_video_encryption   = true
  enable_local_server_video = true
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```