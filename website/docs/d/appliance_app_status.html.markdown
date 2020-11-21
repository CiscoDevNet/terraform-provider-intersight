
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_app_status"
sidebar_current: "docs-intersight-data-source-appliance-app-status"
description: |-
Status of an application.
---

# Data Source: intersight_appliance._app_status
Status of an application.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `app_label`:(string) Unique label to identify the application. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `operational_status`:(string) Operational status of the application.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
* `ready_count`:(int) Number of replicas ready.  The number of instances ofthe application currently ready to perform its intendedfunctions. 
* `replica_count`:(int) Number of replicas provisioned. The number of instancesof the application provisioned to run on the Intersight appliance. 
* `restart_count1_hour`:(int) Number of instance restarts in the last hour. 
* `restart_count24_hours`:(int) Number of instance restarts in the last 24 hours. 
* `restart_count5_mins`:(int) Number of instance restarts in the last 5 minutes. 
* `restart_count_total`:(int) Total number of restarts since last deployment. 
* `running_count`:(int) Number of replicas running. The number of instances ofthe application currently running. 
