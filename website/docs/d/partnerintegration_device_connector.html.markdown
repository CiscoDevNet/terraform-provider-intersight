---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_device_connector"
description: |-
        Recipe for device connector build and deploy.

---

# Data Source: intersight_partnerintegration_device_connector
Recipe for device connector build and deploy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_partnerintegration_device_connector.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be performed on the device connector recipe.* `None` - Default Value of the action, i.e. do nothing.* `Build` - Build the device connector service image.* `Deploy` - Deploy the device connector service on the appliance.* `Upload` - Upload a file to the Partner Integration Appliance bucket. 
* `build_start_time`:(string) Time when build was triggered. 
* `build_status`:(string) Status of build for device connector recipe.* `None` - Default value of the status. i.e. done nothing.* `BackendInProgress` - The backend build is in progress.* `BackendFailed` - The backend build has failed.* `DockerInProgress` - The docker build is in progress.* `DockerFailed` - The docker build has failed.* `Completed` - The operation completed successfully. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `image_name`:(string) Name of the docker image that is built. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the device connector recipe. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_bucket`:(string) Name of the bucket to pick up the file from. 
* `src_file_name`:(string) Name of source file to upload. 
 
