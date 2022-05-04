### Resource Creation

```hcl
resource "intersight_ippool_pool" "ippool_pool1" {
  name             = "ippool_pool1"
  description      = "ippool pool"
  assignment_order = "sequential"
  ip_v4_config {
    object_type = "ippool.IpV4Config"
    gateway     = "10.1.1.1"
    netmask     = "255.0.0.0"
    primary_dns = "8.8.8.8"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
  type        = string
  description = "value for organization"
}
```