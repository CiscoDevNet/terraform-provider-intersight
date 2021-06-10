### Resource Creation

```hcl
resource "intersight_server_profile_template" "template1" {
  name            = "server_profile_template1"
  description     = "demo server profile template"
  target_platform = "FIAttached"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
  # the following policy_bucket statements map different policies to this
  # template -- the object_type shows the policy type
  policy_bucket {
    moid        = intersight_boot_precision_policy.boot_precision1.moid
    object_type = "boot.PrecisionPolicy"
  }
  policy_bucket {
    moid = intersight_access_policy.access1.moid
    object_type = "access.Policy"
  }
  policy_bucket {
    moid = intersight_sol_policy.sol1.moid
    object_type = "sol.Policy"
  }
}
```