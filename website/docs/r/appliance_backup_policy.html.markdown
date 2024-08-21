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
* `protocol`:(string) Communication protocol used by the file server (e.g. scp, sftp, or CIFS).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory or share name to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `retention_count`:(int) The number of backups before earliest backup is overwritten. Requires cleanup policy to be enabled. 
* `retention_policy_enabled`:(bool) If backup rotate policy is set, older backups will automatically be overwritten. The number of backups before overwriting is defined by the retentionCount property. 
* `schedule`:(HashMap) - Schedule to create a backup of the Intersight Appliance. Manualbackup field must be set to 'false' for this schedule to be active. 
This complex property has following sub-properties:
  + `day_of_month`:(int) Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied. 
  + `day_of_week`:(int) Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified. 
  + `month_of_year`:(int) Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `repeat_interval`:(int) Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval. 
  + `time_of_day`:(int) Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified. 
  + `time_zone`:(string) Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC.* `Pacific/Niue` - * `Africa/Abidjan` - * `Africa/Accra` - * `Africa/Addis_Ababa` - * `Africa/Algiers` - * `Africa/Asmara` - * `Africa/Bamako` - * `Africa/Bangui` - * `Africa/Banjul` - * `Africa/Bissau` - * `Africa/Blantyre` - * `Africa/Brazzaville` - * `Africa/Bujumbura` - * `Africa/Cairo` - * `Africa/Casablanca` - * `Africa/Ceuta` - * `Africa/Conakry` - * `Africa/Dakar` - * `Africa/Dar_es_Salaam` - * `Africa/Djibouti` - * `Africa/Douala` - * `Africa/El_Aaiun` - * `Africa/Freetown` - * `Africa/Gaborone` - * `Africa/Harare` - * `Africa/Johannesburg` - * `Africa/Juba` - * `Africa/Kampala` - * `Africa/Khartoum` - * `Africa/Kigali` - * `Africa/Kinshasa` - * `Africa/Lagos` - * `Africa/Libreville` - * `Africa/Lome` - * `Africa/Luanda` - * `Africa/Lubumbashi` - * `Africa/Lusaka` - * `Africa/Malabo` - * `Africa/Maputo` - * `Africa/Maseru` - * `Africa/Mbabane` - * `Africa/Mogadishu` - * `Africa/Monrovia` - * `Africa/Nairobi` - * `Africa/Ndjamena` - * `Africa/Niamey` - * `Africa/Nouakchott` - * `Africa/Ouagadougou` - * `Africa/Porto-Novo` - * `Africa/Sao_Tome` - * `Africa/Tripoli` - * `Africa/Tunis` - * `Africa/Windhoek` - * `America/Adak` - * `America/Anchorage` - * `America/Anguilla` - * `America/Antigua` - * `America/Araguaina` - * `America/Argentina/Buenos_Aires` - * `America/Argentina/Catamarca` - * `America/Argentina/Cordoba` - * `America/Argentina/Jujuy` - * `America/Argentina/La_Rioja` - * `America/Argentina/Mendoza` - * `America/Argentina/Rio_Gallegos` - * `America/Argentina/Salta` - * `America/Argentina/San_Juan` - * `America/Argentina/San_Luis` - * `America/Argentina/Tucuman` - * `America/Argentina/Ushuaia` - * `America/Aruba` - * `America/Asuncion` - * `America/Atikokan` - * `America/Bahia` - * `America/Bahia_Banderas` - * `America/Barbados` - * `America/Belem` - * `America/Belize` - * `America/Blanc-Sablon` - * `America/Boa_Vista` - * `America/Bogota` - * `America/Boise` - * `America/Cambridge_Bay` - * `America/Campo_Grande` - * `America/Cancun` - * `America/Caracas` - * `America/Cayenne` - * `America/Cayman` - * `America/Chicago` - * `America/Chihuahua` - * `America/Costa_Rica` - * `America/Creston` - * `America/Cuiaba` - * `America/Curacao` - * `America/Danmarkshavn` - * `America/Dawson` - * `America/Dawson_Creek` - * `America/Denver` - * `America/Detroit` - * `America/Dominica` - * `America/Edmonton` - * `America/Eirunepe` - * `America/El_Salvador` - * `America/Fortaleza` - * `America/Glace_Bay` - * `America/Godthab` - * `America/Goose_Bay` - * `America/Grand_Turk` - * `America/Grenada` - * `America/Guadeloupe` - * `America/Guatemala` - * `America/Guayaquil` - * `America/Guyana` - * `America/Halifax` - * `America/Havana` - * `America/Hermosillo` - * `America/Indiana/Indianapolis` - * `America/Indiana/Knox` - * `America/Indiana/Marengo` - * `America/Indiana/Petersburg` - * `America/Indiana/Tell_City` - * `America/Indiana/Vevay` - * `America/Indiana/Vincennes` - * `America/Indiana/Winamac` - * `America/Inuvik` - * `America/Iqaluit` - * `America/Jamaica` - * `America/Juneau` - * `America/Kentucky/Louisville` - * `America/Kentucky/Monticello` - * `America/Kralendijk` - * `America/La_Paz` - * `America/Lima` - * `America/Los_Angeles` - * `America/Lower_Princes` - * `America/Maceio` - * `America/Managua` - * `America/Manaus` - * `America/Marigot` - * `America/Martinique` - * `America/Matamoros` - * `America/Mazatlan` - * `America/Menominee` - * `America/Merida` - * `America/Metlakatla` - * `America/Mexico_City` - * `America/Miquelon` - * `America/Moncton` - * `America/Monterrey` - * `America/Montevideo` - * `America/Montreal` - * `America/Montserrat` - * `America/Nassau` - * `America/New_York` - * `America/Nipigon` - * `America/Nome` - * `America/Noronha` - * `America/North_Dakota/Beulah` - * `America/North_Dakota/Center` - * `America/North_Dakota/New_Salem` - * `America/Ojinaga` - * `America/Panama` - * `America/Pangnirtung` - * `America/Paramaribo` - * `America/Phoenix` - * `America/Port-au-Prince` - * `America/Port_of_Spain` - * `America/Porto_Velho` - * `America/Puerto_Rico` - * `America/Rainy_River` - * `America/Rankin_Inlet` - * `America/Recife` - * `America/Regina` - * `America/Resolute` - * `America/Rio_Branco` - * `America/Santa_Isabel` - * `America/Santarem` - * `America/Santiago` - * `America/Santo_Domingo` - * `America/Sao_Paulo` - * `America/Scoresbysund` - * `America/Shiprock` - * `America/Sitka` - * `America/St_Barthelemy` - * `America/St_Johns` - * `America/St_Kitts` - * `America/St_Lucia` - * `America/St_Thomas` - * `America/St_Vincent` - * `America/Swift_Current` - * `America/Tegucigalpa` - * `America/Thule` - * `America/Thunder_Bay` - * `America/Tijuana` - * `America/Toronto` - * `America/Tortola` - * `America/Vancouver` - * `America/Whitehorse` - * `America/Winnipeg` - * `America/Yakutat` - * `America/Yellowknife` - * `Antarctica/Casey` - * `Antarctica/Davis` - * `Antarctica/DumontDUrville` - * `Antarctica/Macquarie` - * `Antarctica/Mawson` - * `Antarctica/McMurdo` - * `Antarctica/Palmer` - * `Antarctica/Rothera` - * `Antarctica/South_Pole` - * `Antarctica/Syowa` - * `Antarctica/Troll` - * `Antarctica/Vostok` - * `Arctic/Longyearbyen` - * `Asia/Aden` - * `Asia/Almaty` - * `Asia/Amman` - * `Asia/Anadyr` - * `Asia/Aqtau` - * `Asia/Aqtobe` - * `Asia/Ashgabat` - * `Asia/Baghdad` - * `Asia/Bahrain` - * `Asia/Baku` - * `Asia/Bangkok` - * `Asia/Beirut` - * `Asia/Bishkek` - * `Asia/Brunei` - * `Asia/Calcutta` - * `Asia/Choibalsan` - * `Asia/Chongqing` - * `Asia/Colombo` - * `Asia/Damascus` - * `Asia/Dhaka` - * `Asia/Dili` - * `Asia/Dubai` - * `Asia/Dushanbe` - * `Asia/Gaza` - * `Asia/Harbin` - * `Asia/Hebron` - * `Asia/Ho_Chi_Minh` - * `Asia/Hong_Kong` - * `Asia/Hovd` - * `Asia/Irkutsk` - * `Asia/Jakarta` - * `Asia/Jayapura` - * `Asia/Jerusalem` - * `Asia/Kabul` - * `Asia/Kamchatka` - * `Asia/Karachi` - * `Asia/Kashgar` - * `Asia/Kathmandu` - * `Asia/Katmandu` - * `Asia/Khandyga` - * `Asia/Kolkata` - * `Asia/Krasnoyarsk` - * `Asia/Kuala_Lumpur` - * `Asia/Kuching` - * `Asia/Kuwait` - * `Asia/Macau` - * `Asia/Magadan` - * `Asia/Makassar` - * `Asia/Manila` - * `Asia/Muscat` - * `Asia/Nicosia` - * `Asia/Novokuznetsk` - * `Asia/Novosibirsk` - * `Asia/Omsk` - * `Asia/Oral` - * `Asia/Phnom_Penh` - * `Asia/Pontianak` - * `Asia/Pyongyang` - * `Asia/Qatar` - * `Asia/Qyzylorda` - * `Asia/Rangoon` - * `Asia/Riyadh` - * `Asia/Saigon` - * `Asia/Sakhalin` - * `Asia/Samarkand` - * `Asia/Seoul` - * `Asia/Shanghai` - * `Asia/Singapore` - * `Asia/Taipei` - * `Asia/Tashkent` - * `Asia/Tbilisi` - * `Asia/Tehran` - * `Asia/Thimphu` - * `Asia/Tokyo` - * `Asia/Ulaanbaatar` - * `Asia/Urumqi` - * `Asia/Ust-Nera` - * `Asia/Vientiane` - * `Asia/Vladivostok` - * `Asia/Yakutsk` - * `Asia/Yekaterinburg` - * `Asia/Yerevan` - * `Atlantic/Azores` - * `Atlantic/Bermuda` - * `Atlantic/Canary` - * `Atlantic/Cape_Verde` - * `Atlantic/Faroe` - * `Atlantic/Madeira` - * `Atlantic/Reykjavik` - * `Atlantic/South_Georgia` - * `Atlantic/St_Helena` - * `Atlantic/Stanley` - * `Australia/Adelaide` - * `Australia/Brisbane` - * `Australia/Broken_Hill` - * `Australia/Currie` - * `Australia/Darwin` - * `Australia/Eucla` - * `Australia/Hobart` - * `Australia/Lindeman` - * `Australia/Lord_Howe` - * `Australia/Melbourne` - * `Australia/Perth` - * `Australia/Sydney` - * `Etc/GMT` - * `Europe/Amsterdam` - * `Europe/Andorra` - * `Europe/Athens` - * `Europe/Belgrade` - * `Europe/Berlin` - * `Europe/Bratislava` - * `Europe/Brussels` - * `Europe/Bucharest` - * `Europe/Budapest` - * `Europe/Busingen` - * `Europe/Chisinau` - * `Europe/Copenhagen` - * `Europe/Dublin` - * `Europe/Gibraltar` - * `Europe/Guernsey` - * `Europe/Helsinki` - * `Europe/Isle_of_Man` - * `Europe/Istanbul` - * `Europe/Jersey` - * `Europe/Kaliningrad` - * `Europe/Kiev` - * `Europe/Lisbon` - * `Europe/Ljubljana` - * `Europe/London` - * `Europe/Luxembourg` - * `Europe/Madrid` - * `Europe/Malta` - * `Europe/Mariehamn` - * `Europe/Minsk` - * `Europe/Monaco` - * `Europe/Moscow` - * `Europe/Oslo` - * `Europe/Paris` - * `Europe/Podgorica` - * `Europe/Prague` - * `Europe/Riga` - * `Europe/Rome` - * `Europe/Samara` - * `Europe/San_Marino` - * `Europe/Sarajevo` - * `Europe/Simferopol` - * `Europe/Skopje` - * `Europe/Sofia` - * `Europe/Stockholm` - * `Europe/Tallinn` - * `Europe/Tirane` - * `Europe/Uzhgorod` - * `Europe/Vaduz` - * `Europe/Vatican` - * `Europe/Vienna` - * `Europe/Vilnius` - * `Europe/Volgograd` - * `Europe/Warsaw` - * `Europe/Zagreb` - * `Europe/Zaporozhye` - * `Europe/Zurich` - * `Indian/Antananarivo` - * `Indian/Chagos` - * `Indian/Christmas` - * `Indian/Cocos` - * `Indian/Comoro` - * `Indian/Kerguelen` - * `Indian/Mahe` - * `Indian/Maldives` - * `Indian/Mauritius` - * `Indian/Mayotte` - * `Indian/Reunion` - * `Pacific/Apia` - * `Pacific/Auckland` - * `Pacific/Chatham` - * `Pacific/Chuuk` - * `Pacific/Easter` - * `Pacific/Efate` - * `Pacific/Enderbury` - * `Pacific/Fakaofo` - * `Pacific/Fiji` - * `Pacific/Funafuti` - * `Pacific/Galapagos` - * `Pacific/Gambier` - * `Pacific/Guadalcanal` - * `Pacific/Guam` - * `Pacific/Honolulu` - * `Pacific/Johnston` - * `Pacific/Kiritimati` - * `Pacific/Kosrae` - * `Pacific/Kwajalein` - * `Pacific/Majuro` - * `Pacific/Marquesas` - * `Pacific/Midway` - * `Pacific/Nauru` - * `Pacific/Norfolk` - * `Pacific/Noumea` - * `Pacific/Pago_Pago` - * `Pacific/Palau` - * `Pacific/Pitcairn` - * `Pacific/Pohnpei` - * `Pacific/Port_Moresby` - * `Pacific/Rarotonga` - * `Pacific/Saipan` - * `Pacific/Tahiti` - * `Pacific/Tarawa` - * `Pacific/Tongatapu` - * `Pacific/Wake` - * `Pacific/Wallis` - * `UTC` - 
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


## Import
`intersight_appliance_backup_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_backup_policy.example 1234567890987654321abcde
``` 
