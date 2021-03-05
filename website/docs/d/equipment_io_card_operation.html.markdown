---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_io_card_operation"
description: |-
  Models the configurable properties of a iomodule in Intersight.
---

# Data Source: intersight_equipment_io_card_operation
Models the configurable properties of a iomodule in Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_io_card_operation.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_power_state`:(string) User configured power state of the iomodule.* `None` - Placeholder default value for iom power state property.* `Reboot` - IO Module reboot state property value. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis iomodule. Applying - This state denotes that the settings are being applied in the target chassis iomodule. Failed - This state denotes that the settings could not be applied in the target chassis iomodule.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 