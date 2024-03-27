---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_proactive_ha"
description: |-
        Vmware vCenter has a functionality to support 'proactive HA' in clusters. Common attributes of 'HA Provider' and custom alarms added for a provider in a vCenter.

---

# Data Source: intersight_virtualization_vmware_proactive_ha
Vmware vCenter has a functionality to support 'proactive HA' in clusters. Common attributes of 'HA Provider' and custom alarms added for a provider in a vCenter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_proactive_ha.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `last_acknowledged_alarm_time`:(string) Last recognized alarm time for a proactive HA alarm instance in a vCenter. 
* `last_sent_alarm_time`:(string) Time at which the last alarm was sent from cloud to the device connector. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
