---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_policy"
description: |-
  BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
---

# Resource: intersight_appliance_backup_policy
BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
## Usage Example
### Resource Creation

```hcl
resource "intersight_appliance_backup_policy" "appliance_backup_policies1" {
  filename      = "default_filename1"
  protocol      = "scp"
  remote_host   = "host.remote"
  remote_path   = "path/to/remote/host"
  remote_port   = 0
  username      = "user_1"
  manual_backup = true
  schedule {
    object_type     = "onprem.Schedule"
    day_of_month    = 1
    day_of_week     = 1
    month_of_year   = 1
    repeat_interval = 0
    time_of_day     = 0
    time_zone       = "Pacific/Niue"
    week_of_month   = 1
  }
  account {
    object_type = "iam.Account"
    moid        = var.account
  }

}

variable "account" {
  type        = string
  description = "Moid of iam.Account Mo"
}
```
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `backup_time`:(string)(ReadOnly) The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_password_set`:(bool)(ReadOnly) Indicates whether the value of the 'password' property has been set. 
* `manual_backup`:(bool) Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `password`:(string) Password to authenticate the file server. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp or sftp).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `schedule`:(HashMap) - Schedule to create a backup of the Intersight Appliance. Manualbackup field must be set to 'false' for this schedule to be active. 
This complex property has following sub-properties:
  + `day_of_month`:(int) Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied. 
  + `day_of_week`:(int) Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified. 
  + `month_of_year`:(int) Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `repeat_interval`:(int) Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval. 
  + `time_of_day`:(int) Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified. 
  + `time_zone`:(string) Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC.* `Pacific/Niue` - * `Pacific/Pago_Pago` - * `Pacific/Honolulu` - * `Pacific/Rarotonga` - * `Pacific/Tahiti` - * `Pacific/Marquesas` - * `America/Anchorage` - * `Pacific/Gambier` - * `America/Los_Angeles` - * `America/Tijuana` - * `America/Vancouver` - * `America/Whitehorse` - * `Pacific/Pitcairn` - * `America/Dawson_Creek` - * `America/Denver` - * `America/Edmonton` - * `America/Hermosillo` - * `America/Mazatlan` - * `America/Phoenix` - * `America/Yellowknife` - * `America/Belize` - * `America/Chicago` - * `America/Costa_Rica` - * `America/El_Salvador` - * `America/Guatemala` - * `America/Managua` - * `America/Mexico_City` - * `America/Regina` - * `America/Tegucigalpa` - * `America/Winnipeg` - * `Pacific/Galapagos` - * `America/Bogota` - * `America/Cancun` - * `America/Cayman` - * `America/Guayaquil` - * `America/Havana` - * `America/Iqaluit` - * `America/Jamaica` - * `America/Lima` - * `America/Nassau` - * `America/New_York` - * `America/Panama` - * `America/Port-au-Prince` - * `America/Rio_Branco` - * `America/Toronto` - * `Pacific/Easter` - * `America/Caracas` - * `America/Asuncion` - * `America/Barbados` - * `America/Boa_Vista` - * `America/Campo_Grande` - * `America/Cuiaba` - * `America/Curacao` - * `America/Grand_Turk` - * `America/Guyana` - * `America/Halifax` - * `America/La_Paz` - * `America/Manaus` - * `America/Martinique` - * `America/Port_of_Spain` - * `America/Porto_Velho` - * `America/Puerto_Rico` - * `America/Santo_Domingo` - * `America/Thule` - * `Atlantic/Bermuda` - * `America/St_Johns` - * `America/Araguaina` - * `America/Argentina/Buenos_Aires` - * `America/Bahia` - * `America/Belem` - * `America/Cayenne` - * `America/Fortaleza` - * `America/Godthab` - * `America/Maceio` - * `America/Miquelon` - * `America/Montevideo` - * `America/Paramaribo` - * `America/Recife` - * `America/Santiago` - * `America/Sao_Paulo` - * `Antarctica/Palmer` - * `Antarctica/Rothera` - * `Atlantic/Stanley` - * `America/Noronha` - * `Atlantic/South_Georgia` - * `America/Scoresbysund` - * `Atlantic/Azores` - * `Atlantic/Cape_Verde` - * `Africa/Abidjan` - * `Africa/Accra` - * `Africa/Bissau` - * `Africa/Casablanca` - * `Africa/El_Aaiun` - * `Africa/Monrovia` - * `America/Danmarkshavn` - * `Atlantic/Canary` - * `Atlantic/Faroe` - * `Atlantic/Reykjavik` - * `Etc/GMT` - * `Europe/Dublin` - * `Europe/Lisbon` - * `Europe/London` - * `Africa/Algiers` - * `Africa/Ceuta` - * `Africa/Lagos` - * `Africa/Ndjamena` - * `Africa/Tunis` - * `Africa/Windhoek` - * `Europe/Amsterdam` - * `Europe/Andorra` - * `Europe/Belgrade` - * `Europe/Berlin` - * `Europe/Brussels` - * `Europe/Budapest` - * `Europe/Copenhagen` - * `Europe/Gibraltar` - * `Europe/Luxembourg` - * `Europe/Madrid` - * `Europe/Malta` - * `Europe/Monaco` - * `Europe/Oslo` - * `Europe/Paris` - * `Europe/Prague` - * `Europe/Rome` - * `Europe/Stockholm` - * `Europe/Tirane` - * `Europe/Vienna` - * `Europe/Warsaw` - * `Europe/Zurich` - * `Africa/Cairo` - * `Africa/Johannesburg` - * `Africa/Maputo` - * `Africa/Tripoli` - * `Asia/Amman` - * `Asia/Beirut` - * `Asia/Damascus` - * `Asia/Gaza` - * `Asia/Jerusalem` - * `Asia/Nicosia` - * `Europe/Athens` - * `Europe/Bucharest` - * `Europe/Chisinau` - * `Europe/Helsinki` - * `Europe/Istanbul` - * `Europe/Kaliningrad` - * `Europe/Kiev` - * `Europe/Riga` - * `Europe/Sofia` - * `Europe/Tallinn` - * `Europe/Vilnius` - * `Africa/Khartoum` - * `Africa/Nairobi` - * `Antarctica/Syowa` - * `Asia/Baghdad` - * `Asia/Qatar` - * `Asia/Riyadh` - * `Europe/Minsk` - * `Europe/Moscow` - * `Asia/Tehran` - * `Asia/Baku` - * `Asia/Dubai` - * `Asia/Tbilisi` - * `Asia/Yerevan` - * `Europe/Samara` - * `Indian/Mahe` - * `Indian/Mauritius` - * `Indian/Reunion` - * `Asia/Kabul` - * `Antarctica/Mawson` - * `Asia/Aqtau` - * `Asia/Aqtobe` - * `Asia/Ashgabat` - * `Asia/Dushanbe` - * `Asia/Karachi` - * `Asia/Tashkent` - * `Asia/Yekaterinburg` - * `Indian/Kerguelen` - * `Indian/Maldives` - * `Asia/Calcutta` - * `Asia/Kolkata` - * `Asia/Colombo` - * `Asia/Katmandu` - * `Antarctica/Vostok` - * `Asia/Almaty` - * `Asia/Bishkek` - * `Asia/Dhaka` - * `Asia/Omsk` - * `Asia/Thimphu` - * `Indian/Chagos` - * `Asia/Rangoon` - * `Indian/Cocos` - * `Antarctica/Davis` - * `Asia/Bangkok` - * `Asia/Hovd` - * `Asia/Jakarta` - * `Asia/Krasnoyarsk` - * `Asia/Saigon` - * `Indian/Christmas` - * `Antarctica/Casey` - * `Asia/Brunei` - * `Asia/Choibalsan` - * `Asia/Hong_Kong` - * `Asia/Irkutsk` - * `Asia/Kuala_Lumpur` - * `Asia/Macau` - * `Asia/Makassar` - * `Asia/Manila` - * `Asia/Shanghai` - * `Asia/Singapore` - * `Asia/Taipei` - * `Asia/Ulaanbaatar` - * `Australia/Perth` - * `Asia/Pyongyang` - * `Asia/Dili` - * `Asia/Jayapura` - * `Asia/Seoul` - * `Asia/Tokyo` - * `Asia/Yakutsk` - * `Pacific/Palau` - * `Australia/Adelaide` - * `Australia/Darwin` - * `Antarctica/DumontDUrville` - * `Asia/Magadan` - * `Asia/Vladivostok` - * `Australia/Brisbane` - * `Australia/Hobart` - * `Australia/Sydney` - * `Pacific/Chuuk` - * `Pacific/Guam` - * `Pacific/Port_Moresby` - * `Pacific/Efate` - * `Pacific/Guadalcanal` - * `Pacific/Kosrae` - * `Pacific/Norfolk` - * `Pacific/Noumea` - * `Pacific/Pohnpei` - * `Asia/Kamchatka` - * `Pacific/Auckland` - * `Pacific/Fiji` - * `Pacific/Funafuti` - * `Pacific/Kwajalein` - * `Pacific/Majuro` - * `Pacific/Nauru` - * `Pacific/Tarawa` - * `Pacific/Wake` - * `Pacific/Wallis` - * `Pacific/Apia` - * `Pacific/Enderbury` - * `Pacific/Fakaofo` - * `Pacific/Tongatapu` - * `Pacific/Kiritimati` - 
  + `week_of_month`:(int) Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to authenticate the fileserver. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_appliance_backup_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_backup_policy.example 1234567890987654321abcde
``` 