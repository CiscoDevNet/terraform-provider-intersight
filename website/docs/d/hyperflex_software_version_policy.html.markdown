
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_software_version_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-software-version-policy"
description: |-
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
---

# Data Source: intersight_hyperflex._software_version_policy
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `hxdp_version`:(string) Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster. 
* `hypervisor_version`:(string) Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `server_firmware_version`:(string) Desired server firmware version to apply on the HyperFlex Cluster. 
