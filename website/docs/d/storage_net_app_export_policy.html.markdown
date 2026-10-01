---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_export_policy"
description: |-
        The NetAppExportPolicies object defines the rules that control client access to volumes via NFS.
        ####  Purpose
        This provides granular security control, specifying which clients can access volumes and what level of access (Read-Only/Read-Write) they are granted.
        ####  Key Concepts
        - **Rule-Based Access:** Uses client matching and security types (e.g., sys, krb5) to grant access.
        - **Security Style:** Defines how permissions are enforced (UNIX vs. NTFS).
        - **Policy Management:** Groups export rules into policies that can be applied to volumes.

---

# Data Source: intersight_storage_net_app_export_policy
The NetAppExportPolicies object defines the rules that control client access to volumes via NFS.
####  Purpose
This provides granular security control, specifying which clients can access volumes and what level of access (Read-Only/Read-Write) they are granted.
####  Key Concepts
- **Rule-Based Access:** Uses client matching and security types (e.g., sys, krb5) to grant access.
- **Security Style:** Defines how permissions are enforced (UNIX vs. NTFS).
- **Policy Management:** Groups export rules into policies that can be applied to volumes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_export_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_uuid`:(string) Unique identity of the device. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the NFS export in storage array. 
* `policy_id`:(int) ID for the Export Policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_name`:(string) The storage virtual machine name for the export policy. 
* `uuid`:(string) The uuid of this NFS export. 
 
