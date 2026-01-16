---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_firmware_summary"
description: |-
        The FirmwareSummary object provides a comprehensive overview of the firmware running on each component within compute physical objects, such as rack or blade servers. It is an integral part of the inventory management system, ensuring that the firmware details are up-to-date and easily accessible for management and policy enforcement.
        #### Purpose
        The FirmwareSummary object is designed to collect and present detailed information about the firmware versions operating across various server components. It plays a crucial role in maintaining system integrity and performance by enabling administrators to monitor, manage, and enforce firmware policies effectively.
        #### Key Concepts
        - **Inventory Integration:** FirmwareSummary is deeply integrated with the inventory system, reflecting accurate and current firmware data for each managed server component.
        - **Policy Enforcement:** Acts as a foundational element for firmware policy implementation, ensuring that firmware versions align with organizational standards and compliance requirements.
        - **Comprehensive Collection:** Aggregates firmware information across multiple components, providing a holistic view of server firmware status, which aids in decision-making and troubleshooting.
        - **Access and Permissions:** Equipped with robust access controls, allowing authorized personnel to read, update, and manage firmware information securely.

---

# Data Source: intersight_firmware_firmware_summary
The FirmwareSummary object provides a comprehensive overview of the firmware running on each component within compute physical objects, such as rack or blade servers. It is an integral part of the inventory management system, ensuring that the firmware details are up-to-date and easily accessible for management and policy enforcement.
#### Purpose
The FirmwareSummary object is designed to collect and present detailed information about the firmware versions operating across various server components. It plays a crucial role in maintaining system integrity and performance by enabling administrators to monitor, manage, and enforce firmware policies effectively.
#### Key Concepts
- **Inventory Integration:** FirmwareSummary is deeply integrated with the inventory system, reflecting accurate and current firmware data for each managed server component.
- **Policy Enforcement:** Acts as a foundational element for firmware policy implementation, ensuring that firmware versions align with organizational standards and compliance requirements.
- **Comprehensive Collection:** Aggregates firmware information across multiple components, providing a holistic view of server firmware status, which aids in decision-making and troubleshooting.
- **Access and Permissions:** Equipped with robust access controls, allowing authorized personnel to read, update, and manage firmware information securely.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_firmware_summary.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bundle_version`:(string) Version details at the bundle level for the each of server. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
