
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `endpoint_address`:(string) VManage server hostname (FQDN) that the acccount holds information for. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `password`:(string) Local password for authenticating with the vManage server. 
* `port`:(int) VManage Port number on which the application is running. 
* `username`:(string) Local username for authenticating with the vManage server. 
