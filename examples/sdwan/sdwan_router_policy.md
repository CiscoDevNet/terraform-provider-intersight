### Resource Creation

```hcl
resource "intersight_sdwan_router_policy" "sdwan_router_policy1" {
  name                 = "sdwan_router_policy1"
  description          = "sdwan_router_policy"
  wan_count            = 2
  wan_termination_type = "Single"
  deployment_size      = "Typical"
   organization {
     object_type = "organization.Organization"
     moid        = var.organization
   }
   profiles {
     object_type = "sdwan.Profile"
     moid        = var.sdwan_profile1
   }
   solution_image {
     object_type = "software.SolutionDistributable"
     moid        = var.solution_distributable
   }

}
variable "organization" {
   type = string
   description = "<value for organization>"
 }

 variable "sdwan_profile1" {
     type = string
     description = "value for sdwan_profile"  
   }

 variable "solution_distributable" {
     type = string
     description = "value for solution_distributable"  
   }
```