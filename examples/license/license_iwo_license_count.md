### Resource Creation

```hcl
resource "intersight_license_iwo_license_count" "license_iwo_license_count1" {
    active_admin = true
    active_license_type = "Base"
    enable_trial = true
    evaluation_period = 30
    extra_evaluation = 15
}
```