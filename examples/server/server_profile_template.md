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
     moid        = var.precision_policy
     object_type = "boot.PrecisionPolicy"
   }
   policy_bucket {
     moid = var.access_policy
     object_type = "access.Policy"
   }
   policy_bucket {
     moid = var.sol_policy
     object_type = "sol.Policy"
   }
}
variable "organization" {
   type = string
   description = "<value for organization>"
 }

variable "sol_policy"{
  type = string
  description = "Moid of sol.Policy"
}

variable "access_policy"{
  type = string
  description = "Moid of access.Policy"
}

variable "precision_policy"{
  type = string
  description = "Moid of boot.PrecisionPolicy"
}
```