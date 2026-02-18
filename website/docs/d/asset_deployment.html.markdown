---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_deployment"
description: |-
        The Deployment object is an essential element of the Cisco consumption-based subscription ecosystem, providing detailed information about deployments associated with subscriptions. It ensures seamless integration and updates through Cisco Install Base, supporting effective deployment tracking and management.
        #### Purpose
        Deployment serves as a detailed repository for tracking and managing deployments under consumption-based subscriptions, facilitating automated updates and ensuring accurate data representation. This information is also queried by downstream OPUS team for billing purposes.
        #### Key Concepts
        - **Integration:** Automates deployment management through synchronization with Cisco Install Base, ensuring data accuracy.
        - **Detailed Information:** Provides comprehensive data on deployment status, customer details, and associated metrics.
        - **Access Management:** Enforces secure access through defined privileges, ensuring only authorized users can interact with deployment data.
        - **Relationship and Dependency Handling:** Manages complex relationships between deployments, devices, and subscriptions, supporting efficient management.
        - **Consumption metering Alerts:** It keeps track of  Deployment level consumption based metering related alerts.

---

# Data Source: intersight_asset_deployment
The Deployment object is an essential element of the Cisco consumption-based subscription ecosystem, providing detailed information about deployments associated with subscriptions. It ensures seamless integration and updates through Cisco Install Base, supporting effective deployment tracking and management.
#### Purpose
Deployment serves as a detailed repository for tracking and managing deployments under consumption-based subscriptions, facilitating automated updates and ensuring accurate data representation. This information is also queried by downstream OPUS team for billing purposes. 
#### Key Concepts
- **Integration:** Automates deployment management through synchronization with Cisco Install Base, ensuring data accuracy.
- **Detailed Information:** Provides comprehensive data on deployment status, customer details, and associated metrics.
- **Access Management:** Enforces secure access through defined privileges, ensuring only authorized users can interact with deployment data.
- **Relationship and Dependency Handling:** Manages complex relationships between deployments, devices, and subscriptions, supporting efficient management.
- **Consumption metering Alerts:** It keeps track of  Deployment level consumption based metering related alerts.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_deployment.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_ref_id`:(string) Identifies the consumption-based subscription's deployment. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_date`:(string) End Date for the consumption-based subscription's deployment. 
* `license_type`:(string) Active license tier for the purchased Cisco device's deployment.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `mlb_offer_type`:(string) Identifier for consumption based pricing. MLB refers to MultiLine Bundle. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_date`:(string) Start Date for the consumption-based subscription's deployment. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
 
