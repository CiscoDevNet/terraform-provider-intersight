---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_state"
description: |-
        DeviceState keeps tracks the Intersight Appliance's current state. Intersight Appliance's
        device connector queries its DeviceState managed object during the startup. Device connector
        also updates DeviceState managed object periodically by pushing modified DeviceState managed
        object from the Intersight Appliance to the Intersight.

---

# Data Source: intersight_appliance_device_state
DeviceState keeps tracks the Intersight Appliance's current state. Intersight Appliance's
device connector queries its DeviceState managed object during the startup. Device connector
also updates DeviceState managed object periodically by pushing modified DeviceState managed
object from the Intersight Appliance to the Intersight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_device_state.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `blocked_version`:(string) Version string of the current software bundle that is blocked for upgrade in the Intersight Appliance. It is used by UI to show banner message for blocked upgrade. 
* `certificate`:(string) Certificate to be used for verifying software upgrade bundles. Intersight's upgrade service sets the certificate dynamically when the Intersight Appliance queries DeviceState. 
* `certificate_not_after`:(string) Expiration date of the software bundle verification certificate. 
* `connection_status`:(string) Intersight Appliance's connectivity status. ConnectionStatus field is updated infrequently, and value may not be up to date. However, upgrade service will populate this field with actual value when queried.* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.* `NotConnected` - Intersight is unable to establish a connection to the target.* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.* `UnclaimInProgress` - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported.* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. 
* `create_time`:(string) The time when this managed object was created. 
* `current_fingerprint`:(string) Fingerprint of the software bundle that is currently installed in the Intersight Appliance. 
* `current_version`:(string) Version string of the current software bundle that is installed in the Intersight Appliance. 
* `dc_version`:(string) Version string of the Intersight Appliance's device connector. Device connector reports version number during the initial handshake. 
* `desired_version`:(string) The desired software bundle version of the Intersight Appliance. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `downloaded_fingerprint`:(string) Fingerprint of the downloaded software bundle. 
* `downloaded_version`:(string) Intersight appliance software bundle version downloaded on the endpoint.Once Intersight Appliance upgrade service starts processing the version, it is updated as pending version. 
* `hostname`:(string) Hostname of the Intersight Appliance. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pending_fingerprint`:(string) Fingerprint of the pending software bundle. 
* `pending_version`:(string) Version string of the pending software bundle that the Intersight Appliance will install. 
* `serial_id`:(string) SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. SerialId is a unique UUID string, and it will not change for the life time of the Intersight Appliance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `upgrade_blocked`:(bool) Flag to indicate whether upgrade on this Intersight Appliance is blocked. 
 
