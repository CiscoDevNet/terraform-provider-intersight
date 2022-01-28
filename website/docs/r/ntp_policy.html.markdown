---
subcategory: "ntp"
layout: "intersight"
page_title: "Intersight: intersight_ntp_policy"
description: |-
  Policy to configure the NTP Servers.
---

# Resource: intersight_ntp_policy
Policy to configure the NTP Servers.
## Usage Example
### Resource Creation

```hcl
resource "intersight_ntp_policy" "ntp1" {
  name        = "ntp1"
  description = "test policy"
  enabled     = true
  ntp_servers = [
    "ntp.esl.cisco.com",
    "time-a-g.nist.gov",
    "time-b-g.nist.gov"
  ]
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}variable "organization" {
   type = string
   description = "value for organization"
 }
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `appliance_account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `authenticated_ntp_servers`:(Array)
This complex property has following sub-properties:
  + `key_type`:(string) Type of symmetric key to use for this server.* `SHA1` - Key type used by the authentication is SHA1. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `server_name`:(string) Server hostname or IP address. 
  + `sym_key_id`:(int) The key ID is a positive integer that identifies a cryptographic key used to authenticate NTP messages. 
  + `sym_key_value`:(string) The value of the symmetric key. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of NTP service on the endpoint. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `ntp_servers`:
                (Array of schema.TypeString) -
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
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
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `timezone`:(string) Timezone of services on the endpoint.* `Pacific/Niue` - * `Pacific/Pago_Pago` - * `Pacific/Honolulu` - * `Pacific/Rarotonga` - * `Pacific/Tahiti` - * `Pacific/Marquesas` - * `America/Anchorage` - * `Pacific/Gambier` - * `America/Los_Angeles` - * `America/Tijuana` - * `America/Vancouver` - * `America/Whitehorse` - * `Pacific/Pitcairn` - * `America/Dawson_Creek` - * `America/Denver` - * `America/Edmonton` - * `America/Hermosillo` - * `America/Mazatlan` - * `America/Phoenix` - * `America/Yellowknife` - * `America/Belize` - * `America/Chicago` - * `America/Costa_Rica` - * `America/El_Salvador` - * `America/Guatemala` - * `America/Managua` - * `America/Mexico_City` - * `America/Regina` - * `America/Tegucigalpa` - * `America/Winnipeg` - * `Pacific/Galapagos` - * `America/Bogota` - * `America/Cancun` - * `America/Cayman` - * `America/Guayaquil` - * `America/Havana` - * `America/Iqaluit` - * `America/Jamaica` - * `America/Lima` - * `America/Nassau` - * `America/New_York` - * `America/Panama` - * `America/Port-au-Prince` - * `America/Rio_Branco` - * `America/Toronto` - * `Pacific/Easter` - * `America/Caracas` - * `America/Asuncion` - * `America/Barbados` - * `America/Boa_Vista` - * `America/Campo_Grande` - * `America/Cuiaba` - * `America/Curacao` - * `America/Grand_Turk` - * `America/Guyana` - * `America/Halifax` - * `America/La_Paz` - * `America/Manaus` - * `America/Martinique` - * `America/Port_of_Spain` - * `America/Porto_Velho` - * `America/Puerto_Rico` - * `America/Santo_Domingo` - * `America/Thule` - * `Atlantic/Bermuda` - * `America/St_Johns` - * `America/Araguaina` - * `America/Argentina/Buenos_Aires` - * `America/Bahia` - * `America/Belem` - * `America/Cayenne` - * `America/Fortaleza` - * `America/Godthab` - * `America/Maceio` - * `America/Miquelon` - * `America/Montevideo` - * `America/Paramaribo` - * `America/Recife` - * `America/Santiago` - * `America/Sao_Paulo` - * `Antarctica/Palmer` - * `Antarctica/Rothera` - * `Atlantic/Stanley` - * `America/Noronha` - * `Atlantic/South_Georgia` - * `America/Scoresbysund` - * `Atlantic/Azores` - * `Atlantic/Cape_Verde` - * `Africa/Abidjan` - * `Africa/Accra` - * `Africa/Bissau` - * `Africa/Casablanca` - * `Africa/El_Aaiun` - * `Africa/Monrovia` - * `America/Danmarkshavn` - * `Atlantic/Canary` - * `Atlantic/Faroe` - * `Atlantic/Reykjavik` - * `Etc/GMT` - * `Europe/Dublin` - * `Europe/Lisbon` - * `Europe/London` - * `Africa/Algiers` - * `Africa/Ceuta` - * `Africa/Lagos` - * `Africa/Ndjamena` - * `Africa/Tunis` - * `Africa/Windhoek` - * `Europe/Amsterdam` - * `Europe/Andorra` - * `Europe/Belgrade` - * `Europe/Berlin` - * `Europe/Brussels` - * `Europe/Budapest` - * `Europe/Copenhagen` - * `Europe/Gibraltar` - * `Europe/Luxembourg` - * `Europe/Madrid` - * `Europe/Malta` - * `Europe/Monaco` - * `Europe/Oslo` - * `Europe/Paris` - * `Europe/Prague` - * `Europe/Rome` - * `Europe/Stockholm` - * `Europe/Tirane` - * `Europe/Vienna` - * `Europe/Warsaw` - * `Europe/Zurich` - * `Africa/Cairo` - * `Africa/Johannesburg` - * `Africa/Maputo` - * `Africa/Tripoli` - * `Asia/Amman` - * `Asia/Beirut` - * `Asia/Damascus` - * `Asia/Gaza` - * `Asia/Jerusalem` - * `Asia/Nicosia` - * `Europe/Athens` - * `Europe/Bucharest` - * `Europe/Chisinau` - * `Europe/Helsinki` - * `Europe/Istanbul` - * `Europe/Kaliningrad` - * `Europe/Kiev` - * `Europe/Riga` - * `Europe/Sofia` - * `Europe/Tallinn` - * `Europe/Vilnius` - * `Africa/Khartoum` - * `Africa/Nairobi` - * `Antarctica/Syowa` - * `Asia/Baghdad` - * `Asia/Qatar` - * `Asia/Riyadh` - * `Europe/Minsk` - * `Europe/Moscow` - * `Asia/Tehran` - * `Asia/Baku` - * `Asia/Dubai` - * `Asia/Tbilisi` - * `Asia/Yerevan` - * `Europe/Samara` - * `Indian/Mahe` - * `Indian/Mauritius` - * `Indian/Reunion` - * `Asia/Kabul` - * `Antarctica/Mawson` - * `Asia/Aqtau` - * `Asia/Aqtobe` - * `Asia/Ashgabat` - * `Asia/Dushanbe` - * `Asia/Karachi` - * `Asia/Tashkent` - * `Asia/Yekaterinburg` - * `Indian/Kerguelen` - * `Indian/Maldives` - * `Asia/Calcutta` - * `Asia/Kolkata` - * `Asia/Colombo` - * `Asia/Katmandu` - * `Antarctica/Vostok` - * `Asia/Almaty` - * `Asia/Bishkek` - * `Asia/Dhaka` - * `Asia/Omsk` - * `Asia/Thimphu` - * `Indian/Chagos` - * `Asia/Rangoon` - * `Indian/Cocos` - * `Antarctica/Davis` - * `Asia/Bangkok` - * `Asia/Hovd` - * `Asia/Jakarta` - * `Asia/Krasnoyarsk` - * `Asia/Saigon` - * `Indian/Christmas` - * `Antarctica/Casey` - * `Asia/Brunei` - * `Asia/Choibalsan` - * `Asia/Hong_Kong` - * `Asia/Irkutsk` - * `Asia/Kuala_Lumpur` - * `Asia/Macau` - * `Asia/Makassar` - * `Asia/Manila` - * `Asia/Shanghai` - * `Asia/Singapore` - * `Asia/Taipei` - * `Asia/Ulaanbaatar` - * `Australia/Perth` - * `Asia/Pyongyang` - * `Asia/Dili` - * `Asia/Jayapura` - * `Asia/Seoul` - * `Asia/Tokyo` - * `Asia/Yakutsk` - * `Pacific/Palau` - * `Australia/Adelaide` - * `Australia/Darwin` - * `Antarctica/DumontDUrville` - * `Asia/Magadan` - * `Asia/Vladivostok` - * `Australia/Brisbane` - * `Australia/Hobart` - * `Australia/Sydney` - * `Pacific/Chuuk` - * `Pacific/Guam` - * `Pacific/Port_Moresby` - * `Pacific/Efate` - * `Pacific/Guadalcanal` - * `Pacific/Kosrae` - * `Pacific/Norfolk` - * `Pacific/Noumea` - * `Pacific/Pohnpei` - * `Asia/Kamchatka` - * `Pacific/Auckland` - * `Pacific/Fiji` - * `Pacific/Funafuti` - * `Pacific/Kwajalein` - * `Pacific/Majuro` - * `Pacific/Nauru` - * `Pacific/Tarawa` - * `Pacific/Wake` - * `Pacific/Wallis` - * `Pacific/Apia` - * `Pacific/Enderbury` - * `Pacific/Fakaofo` - * `Pacific/Tongatapu` - * `Pacific/Kiritimati` - 
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
`intersight_ntp_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_ntp_policy.example 1234567890987654321abcde
``` 