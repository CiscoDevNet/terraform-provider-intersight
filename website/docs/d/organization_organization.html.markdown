---
subcategory: "organization"
layout: "intersight"
page_title: "Intersight: intersight_organization_organization"
description: |-
  Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
---

# Data Source: intersight_organization_organization
Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) The informative description about the usage of this organization. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the organization. There can be multiple organizations under an account. 
