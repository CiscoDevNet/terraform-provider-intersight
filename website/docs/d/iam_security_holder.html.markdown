
---
layout: "intersight"
page_title: "Intersight: intersight_iam_security_holder"
sidebar_current: "docs-intersight-data-source-iam-security-holder"
description: |-
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
---

# Data Source: intersight_iam._security_holder
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
