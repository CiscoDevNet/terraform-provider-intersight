
---
layout: "intersight"
page_title: "Intersight: intersight_sdwan_vmanage_account_policy"
sidebar_current: "docs-intersight-data-source-sdwan-vmanage-account-policy"
description: |-
A policy specifying vManage account details.
---

# Data Source: intersight_sdwan._vmanage_account_policy
A policy specifying vManage account details.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `endpoint_address`:(string) VManage server hostname (FQDN) that the acccount holds information for. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) Local password for authenticating with the vManage server. 
* `port`:(int) VManage Port number on which the application is running. 
* `username`:(string) Local username for authenticating with the vManage server. 
