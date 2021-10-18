---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_setup_info"
description: |-
  SetupInfo will have only one managed object. SetupInfo managed object is to keep
track of the Intersight Appliance's setup information and guide the UI through
the initial configuration of the Intersight Appliance.
The SetupInfo managed object is created during the Intersight Appliance setup.
The Intersight UI uses this object to store the initial configuration states
that the user has completed. If the user closes the Intersight UI without
finishing all the initial configuration, then the Intersight UI will use this
managed object to display the next configuration that the user needs to complete
when the user uses the Intersight Appliance next time.
---

# Data Source: intersight_appliance_setup_info
SetupInfo will have only one managed object. SetupInfo managed object is to keep
track of the Intersight Appliance's setup information and guide the UI through
the initial configuration of the Intersight Appliance.
The SetupInfo managed object is created during the Intersight Appliance setup.
The Intersight UI uses this object to store the initial configuration states
that the user has completed. If the user closes the Intersight UI without
finishing all the initial configuration, then the Intersight UI will use this
managed object to display the next configuration that the user needs to complete
when the user uses the Intersight Appliance next time.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_setup_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_version`:(string) The version of Intersight Appliance backup which can restore to. 
* `build_type`:(string) Build type of the Intersight Appliance setup (e.g. release or debug). 
* `cloud_url`:(string) URL of the Intersight to which this Intersight Appliance is connected to. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_mode`:(string) Indicates where Intersight Appliance is installed in air-gapped or connected mode.In connected mode, Intersight Appliance is claimed by Intesight SaaS.In air-gapped mode, Intersight Appliance does not connect to any Cisco services.* `Connected` - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services.* `Private` - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) End date of the Intersight Appliance's initial setup. 
* `latest_version`:(string) The most recent version which Intersight Appliance can upgrade to. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the Intersight Appliance's initial setup. 
 