---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_vmedia"
description: |-
        The compute.vMedia object represents the virtual media configuration and inventory for a server. It acts as a container for virtual media settings and any currently mapped virtual media images.
        #### Purpose
        The vMedia object is used to manage and monitor the virtual media service on a server. This provides a central point to ensure that the service is enabled, if encryption is active, and what virtual media images are currently mapped. It is essential for remote OS installation, diagnostics, and maintenance.
        #### Key Concepts
        - **Service Configuration:** Provides key settings for the virtual media service, including whether it is enabled and if encryption is active for communications.
        - **Image Inventory:** Contains a collection of compute.Mapping objects, each representing a specific virtual media image that has been mounted to the server.
        - **Boot Integration:** The lowPowerUsb setting controls whether virtual drives appear in the host's boot selection menu, integrating virtual media directly into the boot process.
        - **Centralized Management:** Aggregates all virtual media information for a server into a single, manageable object.

---

# Data Source: intersight_compute_vmedia
The compute.vMedia object represents the virtual media configuration and inventory for a server. It acts as a container for virtual media settings and any currently mapped virtual media images.
#### Purpose
The vMedia object is used to manage and monitor the virtual media service on a server. This provides a central point to ensure that the service is enabled, if encryption is active, and what virtual media images are currently mapped. It is essential for remote OS installation, diagnostics, and maintenance.
#### Key Concepts
- **Service Configuration:** Provides key settings for the virtual media service, including whether it is enabled and if encryption is active for communications.
- **Image Inventory:** Contains a collection of compute.Mapping objects, each representing a specific virtual media image that has been mounted to the server.
- **Boot Integration:** The lowPowerUsb setting controls whether virtual drives appear in the host's boot selection menu, integrating virtual media directly into the boot process.
- **Centralized Management:** Aggregates all virtual media information for a server into a single, manageable object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_vmedia.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of the Virtual Media service on the server. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
