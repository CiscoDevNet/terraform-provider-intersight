---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_device_connector_manager"
description: |-
        ### Overview
        The DeviceConnectorManager object is a central component in managing registered Intersight Assist Appliance Devices within the Intersight ecosystem. This provides essential information and controls pertaining to device registration and management.
        #### Purpose
        DeviceConnectorManager is responsible for maintaining and overseeing the registration status of Intersight Assist appliances. It helps manage the lifecycle and connectivity of devices, ensuring they are effectively integrated and operational within the system.
        #### Key Concepts
        - **Device Registration:** - Manages the registration of Intersight Assist appliances, ensuring they are correctly enrolled and identified within the system.
        - **Lifecycle Management:** - Facilitates the management of the device lifecycle, including registration, updates, and potential deregistration.
        - **Relationship Management:** - Establishes and maintains relationships with registered devices, supporting cascading updates and ensuring that connected systems are informed of changes.
        - **Access Control:** - Utilizes privilege sets to control access and operations, ensuring that only authorized users can perform device management tasks.

---

# Data Source: intersight_asset_device_connector_manager
### Overview
The DeviceConnectorManager object is a central component in managing registered Intersight Assist Appliance Devices within the Intersight ecosystem. This provides essential information and controls pertaining to device registration and management.
 #### Purpose
DeviceConnectorManager is responsible for maintaining and overseeing the registration status of Intersight Assist appliances. It helps manage the lifecycle and connectivity of devices, ensuring they are effectively integrated and operational within the system.
#### Key Concepts
- **Device Registration:** - Manages the registration of Intersight Assist appliances, ensuring they are correctly enrolled and identified within the system.
- **Lifecycle Management:** - Facilitates the management of the device lifecycle, including registration, updates, and potential deregistration.
- **Relationship Management:** - Establishes and maintains relationships with registered devices, supporting cascading updates and ensuring that connected systems are informed of changes.
- **Access Control:** - Utilizes privilege sets to control access and operations, ensuring that only authorized users can perform device management tasks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_device_connector_manager.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
