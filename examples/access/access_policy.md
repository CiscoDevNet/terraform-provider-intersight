### Resource Creation

```hcl
resource "intersight_access_policy" "access1" {
  name        = "access1"
  description = "demo imc access policy"
  inband_vlan = 19
  inband_ip_pool {
    object_type = "ippool.Pool"
    moid        = var.inband_ip_pool
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "inband_ip_pool" {
  type        = string
  description = "Moid of ippool.Pool Mo"
}
```