---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_mapping"
description: |-
        The compute.Mapping object represents a single virtual media image that has been uploaded or mapped to a server. This provides the inventory details for a specific virtual media connection.
        #### Purpose
        The Mapping object is used to track the state of a virtual media mapping. It inventories key details such as the imageName, the type of media (mediaTypes), and a server-assigned identifier. This allows administrators to see which images are currently available to the server for tasks like OS installation or diagnostics.
        #### Key Concepts
        - **Virtual Media Inventory:** Provides a record of a specific virtual media image connected to the server.
        - **Image Identification:** Stores the imageName and fileLocation (if applicable) to identify the source of the virtual media.
        - **State Tracking:** Although primarily for inventory, it represents a specific mapping session that can be created or removed.
        - **Hierarchical Relationship:** Is a child of the compute.Vmedia object, which represents the overall virtual media configuration for the server.

---

# Data Source: intersight_compute_mapping
The compute.Mapping object represents a single virtual media image that has been uploaded or mapped to a server. This provides the inventory details for a specific virtual media connection.
#### Purpose
The Mapping object is used to track the state of a virtual media mapping. It inventories key details such as the imageName, the type of media (mediaTypes), and a server-assigned identifier. This allows administrators to see which images are currently available to the server for tasks like OS installation or diagnostics.
#### Key Concepts
- **Virtual Media Inventory:** Provides a record of a specific virtual media image connected to the server.
- **Image Identification:** Stores the imageName and fileLocation (if applicable) to identify the source of the virtual media.
- **State Tracking:** Although primarily for inventory, it represents a specific mapping session that can be created or removed.
- **Hierarchical Relationship:** Is a child of the compute.Vmedia object, which represents the overall virtual media configuration for the server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_mapping.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_location`:(string) Remote image location from where the image is uploaded to server. 
* `identifier`:(string) The identity assigned to this Virtual Media Image by server. 
* `image_name`:(string) Image name of uploaded Virtual Media Image. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of Virtual Media mapping assigne by server. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
