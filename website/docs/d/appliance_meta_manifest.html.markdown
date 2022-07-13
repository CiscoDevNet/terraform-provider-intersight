---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_meta_manifest"
description: |-
        MetaManifest managed object describes the files by downloading from the S3 and without having required
        user to run the upgrade for the appliance. This MO will be polled periodically from UI to display the
        installation of the current and past history of Metada Manifest ImageBundle.

---

# Data Source: intersight_appliance_meta_manifest
MetaManifest managed object describes the files by downloading from the S3 and without having required
user to run the upgrade for the appliance. This MO will be polled periodically from UI to display the
installation of the current and past history of Metada Manifest ImageBundle.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_meta_manifest.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `creation_date`:(string) Record creation date for explicitly tracking all collections for purging the old records. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_description`:(string) Metadata Manifest ImageBundle file description. 
* `file_name`:(string) File names of the Metadata Manifest ImageBundle as reported by the pvapp portal. 
* `filechk_sum`:(string) The md5 checksum of the Metadata Manifest ImageBundle file. 
* `install_date`:(string) Install date of the Metadata Manifest ImageBundle. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status reported for the successful installation of the meta data files. 
 
