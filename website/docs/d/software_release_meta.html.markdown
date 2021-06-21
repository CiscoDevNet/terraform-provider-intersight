---
subcategory: "software"
layout: "intersight"
page_title: "Intersight: intersight_software_release_meta"
description: |-
  Release information for various software images. Gives information on the latest released version of a product.
---

# Data Source: intersight_software_release_meta
Release information for various software images. Gives information on the latest released version of a product.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_software_release_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `image_type`:(string) The subtype of the distributable image. For e.g. the firmware distributable is categorized according to the component it can upgrade - Standalone server, Intersight managed server or UCS Managed Fabric Interconnect. 
* `latest_file_name`:(string) The name of the latest image file uploaded for this software type. It is populated as part of the image import operation. 
* `latest_version`:(string) Latest version of the image avaiable for a specific software. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `software_type_id`:(string) The software type id of the image (For e.g. firmware.Distributable, software.ApplianceDistributable, software.HyperflexBundleDistributable, software.UcsdBundleDistributable). 
 