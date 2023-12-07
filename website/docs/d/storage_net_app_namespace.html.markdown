---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_namespace"
description: |-
        NetApp Namespace is a collection of addressable logical blocks presented to hosts connected to the storage virtual machine using the NVMe over Fabrics protocol.

---

# Data Source: intersight_storage_net_app_namespace
NetApp Namespace is a collection of addressable logical blocks presented to hosts connected to the storage virtual machine using the NVMe over Fabrics protocol.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_namespace.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `container_state`:(string) The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_mapped`:(string) Reports if the NVMe namespace is mapped to an NVMe subsystem. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The base name component of the NVMe namespace. 
* `namespace_id`:(string) The NVMe namespace identifier. An identifier used by an NVMe controller to provide access to the NVMe namespace. The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \ h\ . 
* `path`:(string) The fully qualified path name of the NVMe namespace composed of a \ /vol\  prefix, the volume name, the (optional) qtree name and base name of the namespace. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of the NVMe namespace. Normal states for a namespace are online and offline. Other states indicate errors (nvfail, space_error). 
* `subsystem_name`:(string) The NVMe subsystem name to which the NVMe namespace is mapped. 
* `svm_name`:(string) The storage virtual machine name for the NVMe namespace. 
* `uuid`:(string) Universally unique identifier of the NVMe namespace. 
* `volume_name`:(string) The volume name in which the NVMe namespace is located. 
 
