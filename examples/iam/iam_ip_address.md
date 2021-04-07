### Resource Creation

```hcl
resource "intersight_iam_ip_address" "iam_ip_address1" {
  address     = "12.13.14.15-12.13.14.200"
  description = "Trusted IP address range."
  ip_access_management {
    moid        = var.ip_access_management
    object_type = "iam.IpAccessManagement"
    enable      = true
  }
}
```