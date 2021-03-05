---
subcategory: "sdwan"
layout: "intersight"
page_title: "Intersight: intersight_sdwan_vmanage_account_policy"
description: |-
  A policy specifying vManage account details.
---

# Data Source: intersight_sdwan_vmanage_account_policy
A policy specifying vManage account details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdwan_vmanage_account_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `endpoint_address`:(string) VManage server hostname (FQDN) that the acccount holds information for. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) Local password for authenticating with the vManage server. 
* `port`:(int) VManage Port number on which the application is running. 
* `username`:(string) Local username for authenticating with the vManage server. 
 