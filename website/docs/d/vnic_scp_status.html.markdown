---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_scp_status"
description: |-
  An internal MO to check if a SCP can be deployed or not on a specific Server Profile.
---

# Data Source: intersight_vnic_scp_status
An internal MO to check if a SCP can be deployed or not on a specific Server Profile.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_scp_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reason`:(string) The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure. 
* `status`:(string) Indicates if the LCP is ready for Deploy or not.* `ok` - No issues with the LCP/SCP/VIF.* `error` - The LCP/SCP/VIF cannot be deployed due to error.* `validating` - Validation in progress for the LCP. 
 