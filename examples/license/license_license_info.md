### Resource Creation

```hcl
resource "intersight_license_license_info" "license_license_info1" {
  evaluation_period = 30
  extra_evaluation  = 15
  tags {
    key   = "main_license"
    value = "license_license_info1"
  }
}
```