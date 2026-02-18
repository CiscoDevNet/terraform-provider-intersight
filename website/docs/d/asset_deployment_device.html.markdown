---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_deployment_device"
description: |-
        The DeploymentDevice object is pivotal in managing device information within consumption-based subscriptions. This provides a detailed view of device installation status, replacements, and transaction history, facilitating seamless management through Cisco Install Base integration.
        #### Purpose
        DeploymentDevice acts as a detailed repository for device-related data within deployments, ensuring accurate tracking under consumption-based subscriptions. This information is also queried by downstream OPUS team for billing purposes.
        #### Key Concepts
        - **Device Management:** Centralizes device data, supporting efficient tracking and management within deployments.
        - **Integration with Install Base:** Listens to messages from Cisco Install Base for automated data updates, ensuring accuracy.
        - **Access Management:** Enforces secure access through defined privileges, ensuring only authorized users can interact with DeploymentDevice data.
        - **Heartbeat generation:** Active DeploymentDevice records are used to publish hourly heartbeats to track device level subscription usage for billing purposes.
        - **Relationship Management:** Manages associations between devices and deployments, supporting effective data management.
        - **Consumption metering Alerts:** It keeps track of  device level consumption based metering related alerts.

---

# Data Source: intersight_asset_deployment_device
The DeploymentDevice object is pivotal in managing device information within consumption-based subscriptions. This provides a detailed view of device installation status, replacements, and transaction history, facilitating seamless management through Cisco Install Base integration.
#### Purpose
DeploymentDevice acts as a detailed repository for device-related data within deployments, ensuring accurate tracking under consumption-based subscriptions. This information is also queried by downstream OPUS team for billing purposes. 
#### Key Concepts
- **Device Management:** Centralizes device data, supporting efficient tracking and management within deployments.
- **Integration with Install Base:** Listens to messages from Cisco Install Base for automated data updates, ensuring accuracy.
- **Access Management:** Enforces secure access through defined privileges, ensuring only authorized users can interact with DeploymentDevice data.
- **Heartbeat generation:** Active DeploymentDevice records are used to publish hourly heartbeats to track device level subscription usage for billing purposes.
- **Relationship Management:** Manages associations between devices and deployments, supporting effective data management.
- **Consumption metering Alerts:** It keeps track of  device level consumption based metering related alerts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_deployment_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Unique identifier of the Cisco device. 
* `device_pid`:(string) Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `product_subgroup`:(string) Product Subgroup type helps to determine if device subgroup within Product type has to be billed using consumption metering. example \ N9300 Series\  in Product type \ SWITCH\ . 
* `product_type`:(string) Product type helps to determine if device has to be billed using consumption metering. example \ SERVER\ . 
* `registered_device_moid`:(string) String reference to the device connector. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `virtualization_platform`:(string) Virtualization platform is used to identify the hypervisor type. example \ ESXi\ . 
* `workload`:(string) Workload/Usecase running on the device. 
 
