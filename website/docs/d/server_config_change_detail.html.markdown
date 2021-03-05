---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_config_change_detail"
description: |-
  The configuration change details are captured here.
---

# Data Source: intersight_server_config_change_detail
The configuration change details are captured here.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_config_change_detail.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `config_change_flag`:(string) Config change flag to differentiate Pending-changes and Config-drift.* `Pending-changes` - Config change flag represents changes are due to not deployed changes from Intersight.* `Drift-changes` - Config change flag represents changes are due to endpoint configuration changes. 
* `message`:(string) Detailed description of the config change. 
* `mod_status`:(string) Modification status of the mo in this config change.* `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.* `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.* `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully.* `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 