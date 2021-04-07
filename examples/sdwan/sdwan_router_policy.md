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
    moid        = intersight_sdwan_profile.sdwan_profile1.id
  }
  solution_image {
    object_type = "software.SolutionDistributable"
    moid        = intersight_software.SolutionDistributable.id
  }

}
```