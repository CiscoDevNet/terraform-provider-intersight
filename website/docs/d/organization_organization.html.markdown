
---
layout: "intersight"
page_title: "Intersight: intersight_organization_organization"
sidebar_current: "docs-intersight-data-source-organization-organization"
description: |-
Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
---

# Data Source: intersight_organization._organization
Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) The informative description about the usage of this organization. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the organization. There can be multiple organizations under an account. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
