---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile"
description: |-
  This specifies the configuration policies for a cluster of switches.
---

# Data Source: intersight_fabric_switch_cluster_profile
This specifies the configuration policies for a cluster of switches.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_switch_cluster_profile.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `switch_profiles_count`:(int) Number of switch profiles that are part of this cluster profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
 