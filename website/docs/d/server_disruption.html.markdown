---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_disruption"
description: |-
        The Disruption object documents disruption events related to server profiles. This provides details on interruptions or issues that may impact server operation or configuration.
        #### Purpose
        Disruption is designed to inform administrators and systems of disruptions, supporting incident response and root cause analysis.
        #### Key Concepts
        - **Event Documentation:** Captures details about disruptions for subsequent review and action.
        - **Operational Awareness:** Integrates with monitoring and alerting systems to provide timely information on disruptions.
        - **Audit Trail:** Maintains a record of all disruptions affecting server profiles.

---

# Data Source: intersight_server_disruption
The Disruption object documents disruption events related to server profiles. This provides details on interruptions or issues that may impact server operation or configuration.
#### Purpose
Disruption is designed to inform administrators and systems of disruptions, supporting incident response and root cause analysis.
#### Key Concepts
- **Event Documentation:** Captures details about disruptions for subsequent review and action.
- **Operational Awareness:** Integrates with monitoring and alerting systems to provide timely information on disruptions.
- **Audit Trail:** Maintains a record of all disruptions affecting server profiles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_disruption.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) Description of the disruption. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Disruption that happens when the policy is deployed. 
* `severity`:(string) Severity of the disruption.* `Minor` - A disruption of minor severity.* `Major` - A disruption of major severity.* `Moderate` - A disruption of moderate severity. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of disruption that happens when the policy is deployed.* `Informational` - Disruptions categorized as informational do not require any user action, nor do they cause a reboot.* `ActionRequired` - User action is required to deploy the profile. 
 
