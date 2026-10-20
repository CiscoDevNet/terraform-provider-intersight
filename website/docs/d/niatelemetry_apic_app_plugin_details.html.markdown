---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_app_plugin_details"
description: |-
        The ApicAppPluginDetails object tracks details of application plugins installed in APIC environments. It supports monitoring, permission management, and operational analytics for app plugins.
        #### Purpose
        ApicAppPluginDetails provides visibility into plugin deployment, permissions, and operational state, aiding system extensibility and control.
        #### Key Concepts
        - **Plugin Inventory**: Documents app IDs and permissions.
        - **Operational Management**: Tracks plugin status and permission levels.
        - **Extensibility**: Supports integration and management of diverse APIC plugins.

---

# Data Source: intersight_niatelemetry_apic_app_plugin_details
The ApicAppPluginDetails object tracks details of application plugins installed in APIC environments. It supports monitoring, permission management, and operational analytics for app plugins.
#### Purpose
ApicAppPluginDetails provides visibility into plugin deployment, permissions, and operational state, aiding system extensibility and control.
#### Key Concepts
- **Plugin Inventory**: Documents app IDs and permissions.
- **Operational Management**: Tracks plugin status and permission levels.
- **Extensibility**: Supports integration and management of diverse APIC plugins.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_app_plugin_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `app_id`:(string) App Id of the plugin in APIC. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `permission`:(string) Permissions to the app plugin in APIC. 
* `permission_level`:(string) Permission Level of the app plugin in APIC. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
 
