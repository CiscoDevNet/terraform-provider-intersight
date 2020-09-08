
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
