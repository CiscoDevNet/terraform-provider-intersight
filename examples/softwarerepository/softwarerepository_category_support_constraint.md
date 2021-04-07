### Resource Creation

```hcl
resource "intersight_softwarerepository_category_support_constraint" "softwarerepository_category_support_constraint" {
    name            = "softwarerepository_category_support_constraint1"
    parse_from_image_name       = true
    filtered_models     = {
        {
            object_type = "softwarerepository.ConstraintModels"
            name        = "softwarerepository_constraint_models1"
            min_version = "12.1(5)E2"
        }
    }
}
```