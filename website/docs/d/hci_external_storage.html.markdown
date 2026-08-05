---
subcategory: "hci"
layout: "intersight"
page_title: "Intersight: intersight_hci_external_storage"
description: |-
        An external storage instance represents storage provided by external vendors such as Pure Storage that is used by Nutanix clusters.
        #### Purpose
        The ExternalStorage object provides a view of external storage configuration and capacity, enabling administrators to monitor Pure Storage FlashArrays, pods, and realms integrated with Nutanix clusters.
        #### Key Concepts
        - **External Storage Integration:** Enables Nutanix clusters to utilize volumes from Pure Storage FlashArrays as underlying storage.
        - **Pod and Realm Support:** Tracks Pure Storage pods (consistency groups) and realms (logical groupings for ActiveCluster).
        - **Capacity Monitoring:** Reports total and free capacity of external storage resources.

---

# Data Source: intersight_hci_external_storage
An external storage instance represents storage provided by external vendors such as Pure Storage that is used by Nutanix clusters.
#### Purpose
The ExternalStorage object provides a view of external storage configuration and capacity, enabling administrators to monitor Pure Storage FlashArrays, pods, and realms integrated with Nutanix clusters.
#### Key Concepts
- **External Storage Integration:** Enables Nutanix clusters to utilize volumes from Pure Storage FlashArrays as underlying storage.
- **Pod and Realm Support:** Tracks Pure Storage pods (consistency groups) and realms (logical groupings for ActiveCluster).
- **Capacity Monitoring:** Reports total and free capacity of external storage resources.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hci_external_storage.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_storage_ext_id`:(string) The external identifier of the external storage instance. 
* `free_capacity_bytes`:(int) Free capacity of the external storage in bytes. 
* `health`:(string) Health status of the external storage. Iris API does not populate health, Janus API does. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the external-storage instance. 
* `pod_ext_id`:(string) Id of the pod this external-storage connects to. 
* `pod_name`:(string) Name of the pod this external-storage connects to. 
* `provider_type`:(string) The type of storage provider. Normalized to EVERPURE_FLASHARRAY for Pure StorageFlashArray (API may return either PURE_STORAGE_FLASHARRAY or EVERPURE_FLASHARRAY). 
* `realm_name`:(string) Realm of the pod this external-storage connects to. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `total_capacity_bytes`:(int) Total capacity of the external storage in bytes. 
 
