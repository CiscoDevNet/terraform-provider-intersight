### Resource Creation

```hcl
resource "intersight_uuidpool_pool" "uuidpool_pool1" {
  name             = "uuidpool_pool1"
  description      = "uuidpool_pool"
  assignment_order = "default"
  prefix           = "123e4567-e89b-42d3"
  uuid_suffix_blocks {
    class_id    = "uuidpool_UuidBlock"
    object_type = "uuidpool.UuidBlock"
    from        = "123e4567-e89b-42d3"
    to          = "123e4567-e89b-84e6"
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