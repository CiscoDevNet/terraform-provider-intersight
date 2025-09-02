---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile_template"
description: |-
        The Switch Cluster Profile Template consists of common switch profile configurations, which can be reused across multiple profiles. Switch Cluster Profiles can be created from the template using the Derive operation. Additionally, an existing profile can be attached to a template to use the configuration set in the template.
        To derive switch cluster profiles from a switch cluster profile template, you must use the synchronous /v1/bulk/MoCloners bulk API.
        Deriving profiles from a Switch Profile Template
        URL: /v1/bulk/MoCloners
        Method: POST
        Body: >
        {
        Sources:[
        {
        Moid:64fb5d17656e6f301e43045b,
        ObjectType:fabric.SwitchClusterProfileTemplate
        }],
        Targets:[
        {
        Name:template1_DERIVED-1,
        ObjectType:fabric.SwitchClusterProfile,
        ClusterAssignments: [
        {
        SourceSwitchProfileOrTemplateName:template1-A,
        NetworkElement:
        {
        ObjectType: network.Element,
        Moid: 64fe8802617675301eb3bdaf
        }
        },
        {
        SourceSwitchProfileOrTemplateName:template1-B,
        NetworkElement:
        {
        ObjectType: network.Element,
        Moid: 64fe8802617675301eb3bdc0
        }
        }],
        Organization:
        {
        ObjectType:organization.Organization,
        Moid:64b0b9ef697265301e52ea0c
        },
        Description:,
        Tags:[]
        }]
        }
        The API response includes the derived SwitchClusterProfile object details.
        Template Updates
        When the profile template is updated, a call to the /v1/bulk/MoMergers API is to be made by the client to synchronize the template changes to all derived profile instances.
        Updating profiles from a Switch Profile Template
        URL: /v1/bulk/MoMergers
        Method: POST
        Body: >
        {
        Sources:[
        {
        Moid:64fb5d17656e6f301e43045b,
        ObjectType:fabric.SwitchClusterProfileTemplate
        }],
        Targets:[
        {
        Moid:6502ffc8656e6f301e5e9f6b,
        ObjectType:fabric.SwitchClusterProfile
        }],
        MergeAction:Replace
        }
        The response of the MoMerger API call would contain the changed profiles.

---

# Data Source: intersight_fabric_switch_cluster_profile_template
The Switch Cluster Profile Template consists of common switch profile configurations, which can be reused across multiple profiles. Switch Cluster Profiles can be created from the template using the Derive operation. Additionally, an existing profile can be attached to a template to use the configuration set in the template.
To derive switch cluster profiles from a switch cluster profile template, you must use the synchronous /v1/bulk/MoCloners bulk API.
Deriving profiles from a Switch Profile Template
URL: /v1/bulk/MoCloners
Method: POST
Body: >
 {
    "Sources":[
      {
        "Moid":"64fb5d17656e6f301e43045b",
        "ObjectType":"fabric.SwitchClusterProfileTemplate"
      }],
    "Targets":[
      {
        "Name":"template1_DERIVED-1",
        "ObjectType":"fabric.SwitchClusterProfile",
        "ClusterAssignments": [
          {
            "SourceSwitchProfileOrTemplateName":"template1-A",
            "NetworkElement":
              {
                "ObjectType": "network.Element",
                "Moid": "64fe8802617675301eb3bdaf"
              }
          },
          {
            "SourceSwitchProfileOrTemplateName":"template1-B",
            "NetworkElement":
              {
                "ObjectType": "network.Element",
                "Moid": "64fe8802617675301eb3bdc0"
              }
          }],
        "Organization":
          {
            "ObjectType":"organization.Organization",
            "Moid":"64b0b9ef697265301e52ea0c"
          },
        "Description":"",
        "Tags":[]
      }]
 }
The API response includes the derived SwitchClusterProfile object details.
Template Updates
When the profile template is updated, a call to the /v1/bulk/MoMergers API is to be made by the client to synchronize the template changes to all derived profile instances.
Updating profiles from a Switch Profile Template
URL: /v1/bulk/MoMergers
Method: POST
Body: >
 {
    "Sources":[
      {
        "Moid":"64fb5d17656e6f301e43045b",
        "ObjectType":"fabric.SwitchClusterProfileTemplate"
      }],
    "Targets":[
      {
        "Moid":"6502ffc8656e6f301e5e9f6b",
        "ObjectType":"fabric.SwitchClusterProfile"
      }],
    "MergeAction":"Replace"
 }
The response of the MoMerger API call would contain the changed profiles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_cluster_profile_template.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_profiles_count`:(int) Number of switch profiles that are part of this cluster profile. 
* `target_platform`:(string) Type of the profile. 'UcsDomain' profile for network and management configuration on UCS Fabric Interconnect. 'UnifiedEdge' profile for network, management and chassis configuration on Unified Edge.* `UCS Domain` - Profile/policy type for network and management configuration on UCS Fabric Interconnect.* `Unified Edge` - Profile/policy type for network, management and chassis configuration on Unified Edge. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `usage`:(int) The count of switch cluster profiles derived from the template. 
 
