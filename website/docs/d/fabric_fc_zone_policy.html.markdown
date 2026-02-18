---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_fc_zone_policy"
description: |-
        The FcZonePolicy object manages FC (Fibre Channel) zoning configurations within the system, enabling the organization and control of SAN (Storage Area Network) environments. It is crucial for maintaining efficient and secure storage network operations.
        #### Purpose
        FcZonePolicy is designed to facilitate the configuration of FC zones, supporting structured and secure SAN management. This provides a framework for organizing initiator and target endpoints, optimizing storage network efficiency and security.
        #### Key Concepts
        - **Zoning Configuration:** Enables the definition and management of FC zones, supporting various zoning types and configurations.
        - **SAN Optimization:** Organizes and defines SAN endpoint interaction rules, enabling consistent and predictable SAN communication behavior.
        - **Security:** Supports secure SAN operations by defining access controls and zoning parameters, protecting data integrity.
        - **Policy Integration:** Works in concert with other SAN policies and configurations, ensuring integrated and streamlined storage management.

---

# Data Source: intersight_fabric_fc_zone_policy
The FcZonePolicy object manages FC (Fibre Channel) zoning configurations within the system, enabling the organization and control of SAN (Storage Area Network) environments. It is crucial for maintaining efficient and secure storage network operations.
#### Purpose
FcZonePolicy is designed to facilitate the configuration of FC zones, supporting structured and secure SAN management. This provides a framework for organizing initiator and target endpoints, optimizing storage network efficiency and security.
#### Key Concepts
- **Zoning Configuration:** Enables the definition and management of FC zones, supporting various zoning types and configurations.
- **SAN Optimization:** Organizes and defines SAN endpoint interaction rules, enabling consistent and predictable SAN communication behavior.
- **Security:** Supports secure SAN operations by defining access controls and zoning parameters, protecting data integrity.
- **Policy Integration:** Works in concert with other SAN policies and configurations, ensuring integrated and streamlined storage management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_fc_zone_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fc_target_zoning_type`:(string) Type of FC zoning. Allowed values are SIST, SIMT and None.* `SIST` - The system automatically creates one zone for each vHBA and storage port pair. Each zone has two members.* `SIMT` - The system automatically creates one zone for each vHBA. Configure this type of zoning if the number of zones created is likely to exceed the maximum supported number of zones.* `None` - FC zoning is not configured. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
