---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_profile_template"
description: |-
        The Chassis Profile Template consists of common chassis profile configurations, which can be reused across multiple profiles. Chassis profiles can be created from the template using the Derive operation. Additionally, an existing profile can be attached to a template to use the configuration set in the template.
        To derive chassis profiles from a chassis profile template, you must use the synchronous /v1/bulk/MoCloners bulk API.
        Deriving profiles from a Chassis Profile Template
        URL: /v1/bulk/MoCloners
        Method: POST
        Body: >
        {
        Sources:[
        {
        Moid:64fb5d17656e6f301e43045b,
        ObjectType:chassis.ProfileTemplate
        }],
        Targets:[
        {
        Name:template1_DERIVED-1”,
        ObjectType:chassis.Profile,
        Organization:
        {
        ObjectType:organization.Organization,
        Moid:64b0b9ef697265301e52ea0c
        },
        Description:,
        Tags:[],
        AssignedChassis:
        {
        Moid:65efe097617675301ecf186f,
        ObjectType:equipment.Chassis
        }
        }]
        }
        The API response includes the derived Chassis profile MO details.
        Template Updates
        When the profile template is updated, a call to the /v1/bulk/MoMergers API is to be made by the client, to synchronize the template changes to all derived profile instances.
        Updating profiles from a Chassis Profile Template
        URL: /v1/bulk/MoMergers
        Method: POST
        Body: >
        {
        Sources:[
        {
        Moid:64fb5d17656e6f301e43045b,
        ObjectType:chassis.ProfileTemplate
        }],
        Targets:[
        {
        Moid:6502ffc8656e6f301e5e9f6b,
        ObjectType:chassis.Profile
        }],
        MergeAction:Replace
        }
        The response of the MoMerger API call would contain the changed profiles.

---

# Data Source: intersight_chassis_profile_template
The Chassis Profile Template consists of common chassis profile configurations, which can be reused across multiple profiles. Chassis profiles can be created from the template using the Derive operation. Additionally, an existing profile can be attached to a template to use the configuration set in the template.
To derive chassis profiles from a chassis profile template, you must use the synchronous /v1/bulk/MoCloners bulk API.
Deriving profiles from a Chassis Profile Template
URL: /v1/bulk/MoCloners
Method: POST
Body: >
 {
    "Sources":[
      {
        "Moid":"64fb5d17656e6f301e43045b",
        "ObjectType":"chassis.ProfileTemplate"
      }],
    "Targets":[
      {
        "Name":"template1_DERIVED-1”,
        "ObjectType":"chassis.Profile",
        "Organization":
          {
            "ObjectType":"organization.Organization",
            "Moid":"64b0b9ef697265301e52ea0c"
          },
        "Description":"",
        "Tags":[],
        "AssignedChassis":
          {
            "Moid":"65efe097617675301ecf186f",
            "ObjectType":"equipment.Chassis
          }
      }]
 }
The API response includes the derived Chassis profile MO details.
Template Updates
When the profile template is updated, a call to the /v1/bulk/MoMergers API is to be made by the client, to synchronize the template changes to all derived profile instances.
Updating profiles from a Chassis Profile Template
URL: /v1/bulk/MoMergers
Method: POST
Body: >
 {
    "Sources":[
      {
        "Moid":"64fb5d17656e6f301e43045b",
        "ObjectType":"chassis.ProfileTemplate"
      }],
    "Targets":[
      {
        "Moid":"6502ffc8656e6f301e5e9f6b",
        "ObjectType":"chassis.Profile"
      }],
    "MergeAction":"Replace"
 }
The response of the MoMerger API call would contain the changed profiles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_profile_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_platform`:(string) The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `usage`:(int) The count of the chassis profiles derived from the template. 
 
