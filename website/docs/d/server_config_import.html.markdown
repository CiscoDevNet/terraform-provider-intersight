---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_config_import"
description: |-
  Configuration import action will import the existing configuration from physical server and generate Intersight policies and server profile from it. At end of successful import, server will be assigned to the generated profile which has policies associated with it. No server profile or policies will be generated if configuration import fails.
---

# Data Source: intersight_server_config_import
Configuration import action will import the existing configuration from physical server and generate Intersight policies and server profile from it. At end of successful import, server will be assigned to the generated profile which has policies associated with it. No server profile or policies will be generated if configuration import fails.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_config_import.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the imported profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `policy_prefix`:(string) Policy prefix for the policies of the imported server profile. 
* `profile_name`:(string) Profile name for the imported server profile. 
 