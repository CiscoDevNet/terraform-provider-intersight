---
subcategory: "scheduler"
layout: "intersight"
page_title: "Intersight: intersight_scheduler_task_schedule"
description: |-
        Metadata used to schedule one-time or repeated tasks.

---

# Resource: intersight_scheduler_task_schedule
Metadata used to schedule one-time or repeated tasks.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) The action of the scheduled task such as suspend or resume.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_object`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) A description to describe the schedule for easier identification. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `last_action`:(string)(ReadOnly) The last action for the scheduled task is saved in this field. Set to none if there was no action.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A schedule name for easier identification (not required to be unique). 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `schedule_params`:(HashMap) - According to the schedule type this property is evaluated. If the property Type is set to OneTime, then the ObjectType must be scheduler.OneTimeScheduleParams. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [scheduler.OneTimeScheduleParams](#schedulerOneTimeScheduleParams)
[scheduler.RecurringScheduleParams](#schedulerRecurringScheduleParams)
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `start_time`:(string) The schedule start time. A future time is required. When the start time is updated, it is mandatory to specify the corresponding timeZone property as well. 
  + `time_zone`:(string) The timezone for the startTime specified. It is a mandatory input property when start time is provided.* `Pacific/Niue` - * `Pacific/Pago_Pago` - * `Pacific/Honolulu` - * `Pacific/Rarotonga` - * `Pacific/Tahiti` - * `Pacific/Marquesas` - * `America/Anchorage` - * `Pacific/Gambier` - * `America/Los_Angeles` - * `America/Tijuana` - * `America/Vancouver` - * `America/Whitehorse` - * `Pacific/Pitcairn` - * `America/Dawson_Creek` - * `America/Denver` - * `America/Edmonton` - * `America/Hermosillo` - * `America/Mazatlan` - * `America/Phoenix` - * `America/Yellowknife` - * `America/Belize` - * `America/Chicago` - * `America/Costa_Rica` - * `America/El_Salvador` - * `America/Guatemala` - * `America/Managua` - * `America/Mexico_City` - * `America/Regina` - * `America/Tegucigalpa` - * `America/Winnipeg` - * `Pacific/Galapagos` - * `America/Bogota` - * `America/Cancun` - * `America/Cayman` - * `America/Guayaquil` - * `America/Havana` - * `America/Iqaluit` - * `America/Jamaica` - * `America/Lima` - * `America/Nassau` - * `America/New_York` - * `America/Nuuk` - * `America/Panama` - * `America/Port-au-Prince` - * `America/Rio_Branco` - * `America/Toronto` - * `Pacific/Easter` - * `America/Caracas` - * `America/Asuncion` - * `America/Barbados` - * `America/Boa_Vista` - * `America/Campo_Grande` - * `America/Cuiaba` - * `America/Curacao` - * `America/Grand_Turk` - * `America/Guyana` - * `America/Halifax` - * `America/La_Paz` - * `America/Manaus` - * `America/Martinique` - * `America/Port_of_Spain` - * `America/Porto_Velho` - * `America/Puerto_Rico` - * `America/Santo_Domingo` - * `America/Thule` - * `Atlantic/Bermuda` - * `America/St_Johns` - * `America/Araguaina` - * `America/Argentina/Buenos_Aires` - * `America/Bahia` - * `America/Belem` - * `America/Cayenne` - * `America/Fortaleza` - * `America/Godthab` - * `America/Maceio` - * `America/Miquelon` - * `America/Montevideo` - * `America/Paramaribo` - * `America/Recife` - * `America/Santiago` - * `America/Sao_Paulo` - * `Antarctica/Palmer` - * `Antarctica/Rothera` - * `Atlantic/Stanley` - * `America/Noronha` - * `Atlantic/South_Georgia` - * `America/Scoresbysund` - * `Atlantic/Azores` - * `Atlantic/Cape_Verde` - * `Africa/Abidjan` - * `Africa/Accra` - * `Africa/Bissau` - * `Africa/Casablanca` - * `Africa/El_Aaiun` - * `Africa/Monrovia` - * `America/Danmarkshavn` - * `Atlantic/Canary` - * `Atlantic/Faroe` - * `Atlantic/Reykjavik` - * `Etc/GMT` - * `Europe/Dublin` - * `Europe/Lisbon` - * `Europe/London` - * `Africa/Algiers` - * `Africa/Ceuta` - * `Africa/Lagos` - * `Africa/Ndjamena` - * `Africa/Tunis` - * `Africa/Windhoek` - * `Europe/Amsterdam` - * `Europe/Andorra` - * `Europe/Belgrade` - * `Europe/Berlin` - * `Europe/Brussels` - * `Europe/Budapest` - * `Europe/Copenhagen` - * `Europe/Gibraltar` - * `Europe/Luxembourg` - * `Europe/Madrid` - * `Europe/Malta` - * `Europe/Monaco` - * `Europe/Oslo` - * `Europe/Paris` - * `Europe/Prague` - * `Europe/Rome` - * `Europe/Stockholm` - * `Europe/Tirane` - * `Europe/Vienna` - * `Europe/Warsaw` - * `Europe/Zurich` - * `Africa/Cairo` - * `Africa/Johannesburg` - * `Africa/Maputo` - * `Africa/Tripoli` - * `Asia/Amman` - * `Asia/Beirut` - * `Asia/Damascus` - * `Asia/Gaza` - * `Asia/Jerusalem` - * `Asia/Nicosia` - * `Europe/Athens` - * `Europe/Bucharest` - * `Europe/Chisinau` - * `Europe/Helsinki` - * `Europe/Istanbul` - * `Europe/Kaliningrad` - * `Europe/Kiev` - * `Europe/Riga` - * `Europe/Sofia` - * `Europe/Tallinn` - * `Europe/Vilnius` - * `Africa/Khartoum` - * `Africa/Nairobi` - * `Antarctica/Syowa` - * `Asia/Baghdad` - * `Asia/Qatar` - * `Asia/Riyadh` - * `Europe/Minsk` - * `Europe/Moscow` - * `Asia/Tehran` - * `Asia/Baku` - * `Asia/Dubai` - * `Asia/Tbilisi` - * `Asia/Yerevan` - * `Europe/Samara` - * `Indian/Mahe` - * `Indian/Mauritius` - * `Indian/Reunion` - * `Asia/Kabul` - * `Antarctica/Mawson` - * `Asia/Aqtau` - * `Asia/Aqtobe` - * `Asia/Ashgabat` - * `Asia/Dushanbe` - * `Asia/Karachi` - * `Asia/Tashkent` - * `Asia/Yekaterinburg` - * `Indian/Kerguelen` - * `Indian/Maldives` - * `Asia/Calcutta` - * `Asia/Kolkata` - * `Asia/Colombo` - * `Asia/Kathmandu` - * `Asia/Katmandu` - * `Antarctica/Vostok` - * `Asia/Almaty` - * `Asia/Bishkek` - * `Asia/Dhaka` - * `Asia/Omsk` - * `Asia/Thimphu` - * `Indian/Chagos` - * `Asia/Rangoon` - * `Indian/Cocos` - * `Antarctica/Davis` - * `Asia/Bangkok` - * `Asia/Ho_Chi_Minh` - * `Asia/Hovd` - * `Asia/Jakarta` - * `Asia/Krasnoyarsk` - * `Asia/Saigon` - * `Indian/Christmas` - * `Antarctica/Casey` - * `Asia/Brunei` - * `Asia/Choibalsan` - * `Asia/Hong_Kong` - * `Asia/Irkutsk` - * `Asia/Kuala_Lumpur` - * `Asia/Macau` - * `Asia/Makassar` - * `Asia/Manila` - * `Asia/Shanghai` - * `Asia/Singapore` - * `Asia/Taipei` - * `Asia/Ulaanbaatar` - * `Australia/Perth` - * `Asia/Pyongyang` - * `Asia/Dili` - * `Asia/Jayapura` - * `Asia/Seoul` - * `Asia/Tokyo` - * `Asia/Yakutsk` - * `Asia/Yangon` - * `Pacific/Palau` - * `Australia/Adelaide` - * `Australia/Darwin` - * `Antarctica/DumontDUrville` - * `Asia/Magadan` - * `Asia/Vladivostok` - * `Australia/Brisbane` - * `Australia/Hobart` - * `Australia/Sydney` - * `Pacific/Chuuk` - * `Pacific/Guam` - * `Pacific/Port_Moresby` - * `Pacific/Efate` - * `Pacific/Guadalcanal` - * `Pacific/Kosrae` - * `Pacific/Norfolk` - * `Pacific/Noumea` - * `Pacific/Pohnpei` - * `Asia/Kamchatka` - * `Pacific/Auckland` - * `Pacific/Fiji` - * `Pacific/Funafuti` - * `Pacific/Kwajalein` - * `Pacific/Majuro` - * `Pacific/Nauru` - * `Pacific/Tarawa` - * `Pacific/Wake` - * `Pacific/Wallis` - * `Pacific/Apia` - * `Pacific/Enderbury` - * `Pacific/Fakaofo` - * `Pacific/Tongatapu` - * `Pacific/Kiritimati` - * `UTC` - 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(HashMap) -(ReadOnly) Status of the current scheduled task. 
This complex property has following sub-properties:
  + `nr_count`:(int)(ReadOnly) The task completion count, which includes both successful executions and any failures. 
  + `current_status`:(string)(ReadOnly) The status of the current task.* `None` - No status is set (default).* `Scheduled` - The status is set when a task is scheduled.* `Running` - The status is set when a task is running.* `Completed` - The status is set when a task is complete.* `Failed` - The status is set when a task fails.* `Suspended` - The status is set when a task is suspended.* `Skipped` - The status is set when a task is skipped because the previous task is still running. 
  + `is_system_suspended`:(bool)(ReadOnly) Indicates if this task was suspended by the system. 
  + `next_run_start_time`:(string)(ReadOnly) The next run time for a recurrently scheduled the task. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `prev_run_end_time`:(string)(ReadOnly) The time when the last occurrence of scheduled task completed. 
  + `prev_run_start_time`:(string)(ReadOnly) The previous time the scheduled task was run. 
  + `reason`:(string)(ReadOnly) The reason why the task failed or suspended. 
* `suspend_end_time`:(string) Suspend a task until an end date. this applies only to the action suspendTill. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `task_request`:(HashMap) - According to the schedule type the recurring params are populated. 
This complex property has following sub-properties:
  + `body`:(JSON as string) The request body that is sent as part of this API request. 
  + `headers`:(JSON as string) Collection of key value pairs to set in the request header. 
  + `method`:(string) The supported values are POST, PUT, DELETE, PATCH. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `protocol`:(string) The accepted web protocol values are http and https. 
  + `response`:(JSON as string) The response obtained for the scheduled API service. 
  + `timeout`:(int) Upper limit on the execution time of a scheduled task. Helps purge run-away scheduled tasks. 
  + `url`:(string) The URL of the resource in the target to which the API request is made. 
* `type`:(string) An Enum describing the type of scheduler to use.* `None` - No value was set for the schedule type (Enum value None).* `OneTime` - Define a one-time task execution time that will not automatically repeat.* `Recurring` - Specify a recurring task cadence based on a predefined pattern, such as daily, weekly, monthly, yearly, or every <interval> pattern. This option is not currently supported. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `workflow_definition`:(HashMap) -(ReadOnly) A reference to a workflowWorkflowDefinition resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_scheduler_task_schedule` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_scheduler_task_schedule.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [scheduler.OneTimeScheduleParams](#argument-reference)
The parameters for configuring a onetime schedule.

### [scheduler.RecurringScheduleParams](#argument-reference)
The parameters for configuring a recurring schedule.
  
