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
* `action`:(string) The action of the scheduled task such as suspend or resume.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. Not supported in this release. 
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
* `execution_statuses`:(Array)
This complex property has following sub-properties:
  + `name`:(string)(ReadOnly) Name of the schedule defined in SchedulePolicy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(HashMap) -(ReadOnly) The status of the scheduled task. 
This complex property has following sub-properties:
    + `completed_count`:(int)(ReadOnly) The count of tasks that ran to successful completion. 
    + `consecutive_failures`:(int)(ReadOnly) The number of consecutive times the task has failed. 
    + `nr_count`:(int)(ReadOnly) The task completion count, which includes both successful executions and any failures. 
    + `current_status`:(string)(ReadOnly) The status of the current task.* `None` - No status is set (default).* `Scheduled` - The status is set when a task is scheduled.* `Running` - The status is set when a task is running.* `Completed` - The status is set when a task is complete.* `Failed` - The status is set when a task fails.* `Suspended` - The status is set when a task is suspended.* `Skipped` - The status is set when a task is skipped because the previous task is still running. 
    + `failed_count`:(int)(ReadOnly) The count of tasks that failed. 
    + `is_system_suspended`:(bool)(ReadOnly) The flag if set to true means it was suspended by the system. 
    + `last_run_status`:(string)(ReadOnly) The last task completion status, which includes both successful executions and any failures.* `None` - No status is set (default).* `Scheduled` - The status is set when a task is scheduled.* `Running` - The status is set when a task is running.* `Completed` - The status is set when a task is complete.* `Failed` - The status is set when a task fails.* `Suspended` - The status is set when a task is suspended.* `Skipped` - The status is set when a task is skipped because the previous task is still running. 
    + `next_run_start_time`:(string)(ReadOnly) The next run time for a recurrently scheduled the task. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `prev_run_end_time`:(string)(ReadOnly) The time when the last occurrence of scheduled task completed. 
    + `prev_run_start_time`:(string)(ReadOnly) The previous time the scheduled task was run. 
    + `reason`:(string)(ReadOnly) The reason why the task failed or suspended. 
    + `skipped_count`:(int)(ReadOnly) The count of tasks that were skipped. 
* `last_action`:(string)(ReadOnly) The last action for the scheduled task is saved in this field. Set to none if there was no action.* `None` - No action is set (default).* `Suspend` - Suspend a scheduled task indefinitely.* `Resume` - Resume a suspended scheduled task.* `SuspendTill` - Suspend the scheduled task until a specified end-date. Not supported in this release. 
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
* `policy`:(HashMap) - A reference to a schedulerSchedulePolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `schedule_params`:(HashMap) - According to the schedule type this property is evaluated. If the property Type is set to OneTime, then the ObjectType must be scheduler.OneTimeScheduleParams. If the Type is Recurring, then the ObjectType must be scheduler.RecurringScheduleParams. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [scheduler.OneTimeScheduleParams](#schedulerOneTimeScheduleParams)
[scheduler.RecurringScheduleParams](#schedulerRecurringScheduleParams)
  + `duration`:(string) The duration of the schedule. Its syntax is specified at https://www.w3.org/TR/xmlschema11-2/#nt-durationRep For example, P20DT10H5M2.3S is for 20 days, 10 hours, 5 minutes and 2.3 seconds. It is a mandatory input property for Policy based schedules. 
  + `name`:(string) The name of the schedule. It is a mandatory input property for Policy based schedules. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `start_time`:(string) The schedule start time. A future time is required. 
  + `time_zone`:(string) The timezone for the startTime specified. It is a mandatory input property when start time is provided.* `Pacific/Niue` - * `Africa/Abidjan` - * `Africa/Accra` - * `Africa/Addis_Ababa` - * `Africa/Algiers` - * `Africa/Asmara` - * `Africa/Bamako` - * `Africa/Bangui` - * `Africa/Banjul` - * `Africa/Bissau` - * `Africa/Blantyre` - * `Africa/Brazzaville` - * `Africa/Bujumbura` - * `Africa/Cairo` - * `Africa/Casablanca` - * `Africa/Ceuta` - * `Africa/Conakry` - * `Africa/Dakar` - * `Africa/Dar_es_Salaam` - * `Africa/Djibouti` - * `Africa/Douala` - * `Africa/El_Aaiun` - * `Africa/Freetown` - * `Africa/Gaborone` - * `Africa/Harare` - * `Africa/Johannesburg` - * `Africa/Juba` - * `Africa/Kampala` - * `Africa/Khartoum` - * `Africa/Kigali` - * `Africa/Kinshasa` - * `Africa/Lagos` - * `Africa/Libreville` - * `Africa/Lome` - * `Africa/Luanda` - * `Africa/Lubumbashi` - * `Africa/Lusaka` - * `Africa/Malabo` - * `Africa/Maputo` - * `Africa/Maseru` - * `Africa/Mbabane` - * `Africa/Mogadishu` - * `Africa/Monrovia` - * `Africa/Nairobi` - * `Africa/Ndjamena` - * `Africa/Niamey` - * `Africa/Nouakchott` - * `Africa/Ouagadougou` - * `Africa/Porto-Novo` - * `Africa/Sao_Tome` - * `Africa/Tripoli` - * `Africa/Tunis` - * `Africa/Windhoek` - * `America/Adak` - * `America/Anchorage` - * `America/Anguilla` - * `America/Antigua` - * `America/Araguaina` - * `America/Argentina/Buenos_Aires` - * `America/Argentina/Catamarca` - * `America/Argentina/Cordoba` - * `America/Argentina/Jujuy` - * `America/Argentina/La_Rioja` - * `America/Argentina/Mendoza` - * `America/Argentina/Rio_Gallegos` - * `America/Argentina/Salta` - * `America/Argentina/San_Juan` - * `America/Argentina/San_Luis` - * `America/Argentina/Tucuman` - * `America/Argentina/Ushuaia` - * `America/Aruba` - * `America/Asuncion` - * `America/Atikokan` - * `America/Bahia` - * `America/Bahia_Banderas` - * `America/Barbados` - * `America/Belem` - * `America/Belize` - * `America/Blanc-Sablon` - * `America/Boa_Vista` - * `America/Bogota` - * `America/Boise` - * `America/Cambridge_Bay` - * `America/Campo_Grande` - * `America/Cancun` - * `America/Caracas` - * `America/Cayenne` - * `America/Cayman` - * `America/Chicago` - * `America/Chihuahua` - * `America/Costa_Rica` - * `America/Creston` - * `America/Cuiaba` - * `America/Curacao` - * `America/Danmarkshavn` - * `America/Dawson` - * `America/Dawson_Creek` - * `America/Denver` - * `America/Detroit` - * `America/Dominica` - * `America/Edmonton` - * `America/Eirunepe` - * `America/El_Salvador` - * `America/Fortaleza` - * `America/Glace_Bay` - * `America/Godthab` - * `America/Goose_Bay` - * `America/Grand_Turk` - * `America/Grenada` - * `America/Guadeloupe` - * `America/Guatemala` - * `America/Guayaquil` - * `America/Guyana` - * `America/Halifax` - * `America/Havana` - * `America/Hermosillo` - * `America/Indiana/Indianapolis` - * `America/Indiana/Knox` - * `America/Indiana/Marengo` - * `America/Indiana/Petersburg` - * `America/Indiana/Tell_City` - * `America/Indiana/Vevay` - * `America/Indiana/Vincennes` - * `America/Indiana/Winamac` - * `America/Inuvik` - * `America/Iqaluit` - * `America/Jamaica` - * `America/Juneau` - * `America/Kentucky/Louisville` - * `America/Kentucky/Monticello` - * `America/Kralendijk` - * `America/La_Paz` - * `America/Lima` - * `America/Los_Angeles` - * `America/Lower_Princes` - * `America/Maceio` - * `America/Managua` - * `America/Manaus` - * `America/Marigot` - * `America/Martinique` - * `America/Matamoros` - * `America/Mazatlan` - * `America/Menominee` - * `America/Merida` - * `America/Metlakatla` - * `America/Mexico_City` - * `America/Miquelon` - * `America/Moncton` - * `America/Monterrey` - * `America/Montevideo` - * `America/Montreal` - * `America/Montserrat` - * `America/Nassau` - * `America/New_York` - * `America/Nipigon` - * `America/Nome` - * `America/Noronha` - * `America/North_Dakota/Beulah` - * `America/North_Dakota/Center` - * `America/North_Dakota/New_Salem` - * `America/Ojinaga` - * `America/Panama` - * `America/Pangnirtung` - * `America/Paramaribo` - * `America/Phoenix` - * `America/Port-au-Prince` - * `America/Port_of_Spain` - * `America/Porto_Velho` - * `America/Puerto_Rico` - * `America/Rainy_River` - * `America/Rankin_Inlet` - * `America/Recife` - * `America/Regina` - * `America/Resolute` - * `America/Rio_Branco` - * `America/Santa_Isabel` - * `America/Santarem` - * `America/Santiago` - * `America/Santo_Domingo` - * `America/Sao_Paulo` - * `America/Scoresbysund` - * `America/Shiprock` - * `America/Sitka` - * `America/St_Barthelemy` - * `America/St_Johns` - * `America/St_Kitts` - * `America/St_Lucia` - * `America/St_Thomas` - * `America/St_Vincent` - * `America/Swift_Current` - * `America/Tegucigalpa` - * `America/Thule` - * `America/Thunder_Bay` - * `America/Tijuana` - * `America/Toronto` - * `America/Tortola` - * `America/Vancouver` - * `America/Whitehorse` - * `America/Winnipeg` - * `America/Yakutat` - * `America/Yellowknife` - * `Antarctica/Casey` - * `Antarctica/Davis` - * `Antarctica/DumontDUrville` - * `Antarctica/Macquarie` - * `Antarctica/Mawson` - * `Antarctica/McMurdo` - * `Antarctica/Palmer` - * `Antarctica/Rothera` - * `Antarctica/South_Pole` - * `Antarctica/Syowa` - * `Antarctica/Troll` - * `Antarctica/Vostok` - * `Arctic/Longyearbyen` - * `Asia/Aden` - * `Asia/Almaty` - * `Asia/Amman` - * `Asia/Anadyr` - * `Asia/Aqtau` - * `Asia/Aqtobe` - * `Asia/Ashgabat` - * `Asia/Baghdad` - * `Asia/Bahrain` - * `Asia/Baku` - * `Asia/Bangkok` - * `Asia/Beirut` - * `Asia/Bishkek` - * `Asia/Brunei` - * `Asia/Calcutta` - * `Asia/Choibalsan` - * `Asia/Chongqing` - * `Asia/Colombo` - * `Asia/Damascus` - * `Asia/Dhaka` - * `Asia/Dili` - * `Asia/Dubai` - * `Asia/Dushanbe` - * `Asia/Gaza` - * `Asia/Harbin` - * `Asia/Hebron` - * `Asia/Ho_Chi_Minh` - * `Asia/Hong_Kong` - * `Asia/Hovd` - * `Asia/Irkutsk` - * `Asia/Jakarta` - * `Asia/Jayapura` - * `Asia/Jerusalem` - * `Asia/Kabul` - * `Asia/Kamchatka` - * `Asia/Karachi` - * `Asia/Kashgar` - * `Asia/Kathmandu` - * `Asia/Katmandu` - * `Asia/Khandyga` - * `Asia/Kolkata` - * `Asia/Krasnoyarsk` - * `Asia/Kuala_Lumpur` - * `Asia/Kuching` - * `Asia/Kuwait` - * `Asia/Macau` - * `Asia/Magadan` - * `Asia/Makassar` - * `Asia/Manila` - * `Asia/Muscat` - * `Asia/Nicosia` - * `Asia/Novokuznetsk` - * `Asia/Novosibirsk` - * `Asia/Omsk` - * `Asia/Oral` - * `Asia/Phnom_Penh` - * `Asia/Pontianak` - * `Asia/Pyongyang` - * `Asia/Qatar` - * `Asia/Qyzylorda` - * `Asia/Rangoon` - * `Asia/Riyadh` - * `Asia/Saigon` - * `Asia/Sakhalin` - * `Asia/Samarkand` - * `Asia/Seoul` - * `Asia/Shanghai` - * `Asia/Singapore` - * `Asia/Taipei` - * `Asia/Tashkent` - * `Asia/Tbilisi` - * `Asia/Tehran` - * `Asia/Thimphu` - * `Asia/Tokyo` - * `Asia/Ulaanbaatar` - * `Asia/Urumqi` - * `Asia/Ust-Nera` - * `Asia/Vientiane` - * `Asia/Vladivostok` - * `Asia/Yakutsk` - * `Asia/Yekaterinburg` - * `Asia/Yerevan` - * `Atlantic/Azores` - * `Atlantic/Bermuda` - * `Atlantic/Canary` - * `Atlantic/Cape_Verde` - * `Atlantic/Faroe` - * `Atlantic/Madeira` - * `Atlantic/Reykjavik` - * `Atlantic/South_Georgia` - * `Atlantic/St_Helena` - * `Atlantic/Stanley` - * `Australia/Adelaide` - * `Australia/Brisbane` - * `Australia/Broken_Hill` - * `Australia/Currie` - * `Australia/Darwin` - * `Australia/Eucla` - * `Australia/Hobart` - * `Australia/Lindeman` - * `Australia/Lord_Howe` - * `Australia/Melbourne` - * `Australia/Perth` - * `Australia/Sydney` - * `Etc/GMT` - * `Europe/Amsterdam` - * `Europe/Andorra` - * `Europe/Athens` - * `Europe/Belgrade` - * `Europe/Berlin` - * `Europe/Bratislava` - * `Europe/Brussels` - * `Europe/Bucharest` - * `Europe/Budapest` - * `Europe/Busingen` - * `Europe/Chisinau` - * `Europe/Copenhagen` - * `Europe/Dublin` - * `Europe/Gibraltar` - * `Europe/Guernsey` - * `Europe/Helsinki` - * `Europe/Isle_of_Man` - * `Europe/Istanbul` - * `Europe/Jersey` - * `Europe/Kaliningrad` - * `Europe/Kiev` - * `Europe/Lisbon` - * `Europe/Ljubljana` - * `Europe/London` - * `Europe/Luxembourg` - * `Europe/Madrid` - * `Europe/Malta` - * `Europe/Mariehamn` - * `Europe/Minsk` - * `Europe/Monaco` - * `Europe/Moscow` - * `Europe/Oslo` - * `Europe/Paris` - * `Europe/Podgorica` - * `Europe/Prague` - * `Europe/Riga` - * `Europe/Rome` - * `Europe/Samara` - * `Europe/San_Marino` - * `Europe/Sarajevo` - * `Europe/Simferopol` - * `Europe/Skopje` - * `Europe/Sofia` - * `Europe/Stockholm` - * `Europe/Tallinn` - * `Europe/Tirane` - * `Europe/Uzhgorod` - * `Europe/Vaduz` - * `Europe/Vatican` - * `Europe/Vienna` - * `Europe/Vilnius` - * `Europe/Volgograd` - * `Europe/Warsaw` - * `Europe/Zagreb` - * `Europe/Zaporozhye` - * `Europe/Zurich` - * `Indian/Antananarivo` - * `Indian/Chagos` - * `Indian/Christmas` - * `Indian/Cocos` - * `Indian/Comoro` - * `Indian/Kerguelen` - * `Indian/Mahe` - * `Indian/Maldives` - * `Indian/Mauritius` - * `Indian/Mayotte` - * `Indian/Reunion` - * `Pacific/Apia` - * `Pacific/Auckland` - * `Pacific/Chatham` - * `Pacific/Chuuk` - * `Pacific/Easter` - * `Pacific/Efate` - * `Pacific/Enderbury` - * `Pacific/Fakaofo` - * `Pacific/Fiji` - * `Pacific/Funafuti` - * `Pacific/Galapagos` - * `Pacific/Gambier` - * `Pacific/Guadalcanal` - * `Pacific/Guam` - * `Pacific/Honolulu` - * `Pacific/Johnston` - * `Pacific/Kiritimati` - * `Pacific/Kosrae` - * `Pacific/Kwajalein` - * `Pacific/Majuro` - * `Pacific/Marquesas` - * `Pacific/Midway` - * `Pacific/Nauru` - * `Pacific/Norfolk` - * `Pacific/Noumea` - * `Pacific/Pago_Pago` - * `Pacific/Palau` - * `Pacific/Pitcairn` - * `Pacific/Pohnpei` - * `Pacific/Port_Moresby` - * `Pacific/Rarotonga` - * `Pacific/Saipan` - * `Pacific/Tahiti` - * `Pacific/Tarawa` - * `Pacific/Tongatapu` - * `Pacific/Wake` - * `Pacific/Wallis` - * `UTC` - 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(HashMap) -(ReadOnly) Status of the current scheduled task. 
This complex property has following sub-properties:
  + `completed_count`:(int)(ReadOnly) The count of tasks that ran to successful completion. 
  + `consecutive_failures`:(int)(ReadOnly) The number of consecutive times the task has failed. 
  + `nr_count`:(int)(ReadOnly) The task completion count, which includes both successful executions and any failures. 
  + `current_status`:(string)(ReadOnly) The status of the current task.* `None` - No status is set (default).* `Scheduled` - The status is set when a task is scheduled.* `Running` - The status is set when a task is running.* `Completed` - The status is set when a task is complete.* `Failed` - The status is set when a task fails.* `Suspended` - The status is set when a task is suspended.* `Skipped` - The status is set when a task is skipped because the previous task is still running. 
  + `failed_count`:(int)(ReadOnly) The count of tasks that failed. 
  + `is_system_suspended`:(bool)(ReadOnly) The flag if set to true means it was suspended by the system. 
  + `last_run_status`:(string)(ReadOnly) The last task completion status, which includes both successful executions and any failures.* `None` - No status is set (default).* `Scheduled` - The status is set when a task is scheduled.* `Running` - The status is set when a task is running.* `Completed` - The status is set when a task is complete.* `Failed` - The status is set when a task fails.* `Suspended` - The status is set when a task is suspended.* `Skipped` - The status is set when a task is skipped because the previous task is still running. 
  + `next_run_start_time`:(string)(ReadOnly) The next run time for a recurrently scheduled the task. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `prev_run_end_time`:(string)(ReadOnly) The time when the last occurrence of scheduled task completed. 
  + `prev_run_start_time`:(string)(ReadOnly) The previous time the scheduled task was run. 
  + `reason`:(string)(ReadOnly) The reason why the task failed or suspended. 
  + `skipped_count`:(int)(ReadOnly) The count of tasks that were skipped. 
* `suspend_end_time`:(string) Suspend a task until an end date. this applies only to the action suspendTill. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `task_request`:(HashMap) - According to the schedule type the recurring params are populated. 
This complex property has following sub-properties:
  + `body`:(JSON as string) The request body that is sent as part of this API request. 
  + `headers`:(JSON as string) Collection of key value pairs to set in the request header. 
  + `method`:(string) The supported values are POST, PUT, DELETE, PATCH. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `protocol`:(string) The accepted web protocol values are http and https. 
  + `response`:(JSON as string) The response obtained for the scheduled API service. 
  + `timeout`:(int) Upper limit on the execution time of a scheduled task. Helps purge run-away scheduled tasks.Not supported in this release. 
  + `url`:(string) The URL of the resource in the target to which the API request is made. 
* `type`:(string) An Enum describing the type of scheduler to use.* `None` - No value was set for the schedule type (Enum value None).* `OneTime` - Define a one-time task execution time that will not automatically repeat.* `Recurring` - Specify a recurring task cadence based on a predefined pattern, such as daily, weekly, monthly, or every <interval> pattern. 
* `use_policy`:(bool) Indicates if the schedule is policy based or not. 
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
* `cadence`:(string) Allowed values for a recurring schedule cadence.* `None` - No value set for the cadence type (Enum value None).* `Every` - Use the 'Every' cadence for tasks that need to be run frequently and are relatively small or quick to execute. This could include tasks such as checking the status of a service every 15 minutes, or updating a counter.* `Daily` - A Daily cadence allows for a scheduled task to be run every day or every n-interval days.* `Weekly` - A Weekly cadence allows for a scheduled task to be run every week or every n-interval weeks on specific days.* `Monthly` - A Montly cadence allows for a scheduled task to be run every month on specific days. 
* `end_after_occurrences`:(int) Specify the number of occurrences (instead of an end-time) for a recurring schedule. 
* `end_time`:(string) End time for the recurring schedule. The schedule will not run beyond this time. If using the endAfterOccurrences parameter instead, this field should be set to zero time, i.e, 0001-01-01T00:00:00Z. 
* `failure_threshold`:(int) The maximum number of consecutive failures until the recurring scheduled task is suspended by the system. The default is 1. 
* `params`:(HashMap) - The cadence for the recurring schedule. Different parameters are used depending on the schedule type. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  
