---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_info"
description: |-
  UCS Director accounts managed by Intersight.
---

# Data Source: intersight_iaas_ucsd_info
UCS Director accounts managed by Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_ucsd_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Moid of the UCS Director device connector's asset.DeviceRegistration. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `guid`:(string) Unique ID of UCS Director getting registerd with Intersight. 
* `host_name`:(string) The UCS Director hostname for management. 
* `ip`:(string) The UCS Director IP address for management. 
* `last_backup`:(string) Last successful backup created for this UCS Director appliance if backup is configured. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_type`:(string) NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node. 
* `product_name`:(string) The UCS Director product name. 
* `product_vendor`:(string) The UCS Director product vendor. 
* `product_version`:(string) The UCS Director product/platform version. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The UCS Director status. Possible values are Active, Inactive, Unknown. 
 