---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cifs_service"
description: |-
        NetApp CIFS service represents a CIFS server on a storage virtual machine.

---

# Data Source: intersight_storage_net_app_cifs_service
NetApp CIFS service represents a CIFS server on a storage virtual machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cifs_service.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active_directory_fqdn`:(string) The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs. 
* `ad_organizational_unit`:(string) Identifies the organizational unit within the Active Directory domain to associate with the CIFS server. 
* `comment`:(string) A descriptive text comment for the CIFS server. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(string) Indicates that the CIFS service is administratively enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_name`:(string) Name of the NetApp CIFS server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_uuid`:(string) Unique identifier for the NetApp Storage Virtual Machine. 
 
