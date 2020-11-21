
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_setup_info"
sidebar_current: "docs-intersight-data-source-appliance-setup-info"
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

# Data Source: intersight_appliance._setup_info
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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `build_type`:(string) Build type of the Intersight Appliance setup (e.g. release or debug). 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cloud_url`:(string) URL of the Intersight to which this Intersight Appliance is connected to. 
* `deployment_mode`:(string) Indicates where Intersight Appliance is installed in air-gapped or connected mode.In connected mode, Intersight Appliance is claimed by Intesight SaaS.In air-gapped mode, Intersight Appliance does not connect to any Cisco services.* `Connected` - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services.* `Private` - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services. 
* `end_time`:(string) End date of the Intersight Appliance's initial setup. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `start_time`:(string) Start date of the Intersight Appliance's initial setup. 
