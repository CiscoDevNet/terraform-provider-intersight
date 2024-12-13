---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_alarm"
description: |-
        An alert from an HCI cluster reported from the Prism Central serviceability/Alerts API.

---

# Data Source: intersight_hci_alarm
An alert from an HCI cluster reported from the Prism Central serviceability/Alerts API.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_alarm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alarm_ext_id`:(string) The unique identifier for the alarm on the endpoint. 
* `alert_type`:(string) The code for the reported alarm. 
* `cluster_ext_id`:(string) The unique identifer for the cluster associated with the source entity on the endpoint. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_resolved`:(bool) The status of the alarm. If an alarm is resolved, this value is set as true on the endpoint. 
* `message`:(string) The description from the endpoint explaining the cause of the alarm. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `origin_acknowledge_time`:(string) The time the alarm was acknowledged on the endpoint. 
* `origin_creation_time`:(string) The time the alarm was created on the endpoint. 
* `origin_is_acknowledged`:(bool) The acknowledgement status of the alert on the endpoint. 
* `resolved_time`:(string) The time the alarm was resolved on the endpoint. 
* `severity`:(string) The severity of the alarm. Valid values are Critical, Warning, Info.* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_entity_ext_id`:(string) The unique identifer for the entity on the endpoint for which the alarm was raised. 
* `source_entity_name`:(string) The name of the entity on the endpoint for which the alarm was raised. 
* `source_entity_type`:(string) The object type for the entity corresponding to the objects inventoried. 
* `title`:(string) The title of the reported alarm. 
 
