
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_capability_info"
sidebar_current: "docs-intersight-data-source-hyperflex-capability-info"
description: |-
A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
---

# Data Source: intersight_hyperflex._capability_info
A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the capability or feature set consisting of a collection of constraint rules and value. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `value`:(string) Capability Value which is valid only iff all specified constraints match. 
