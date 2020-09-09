# HyperflexSysConfigPolicyRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**Name** | Pointer to **string** | Name of the concrete policy. | [optional] 
**DnsDomainName** | Pointer to **string** | The DNS Search Domain Name. This setting applies to HyperFlex Data Platform 3.0 or later only. | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **string** | The timezone of the HyperFlex cluster&#39;s system clock. * &#x60;Pacific/Niue&#x60; -  * &#x60;Pacific/Pago_Pago&#x60; -  * &#x60;Pacific/Honolulu&#x60; -  * &#x60;Pacific/Rarotonga&#x60; -  * &#x60;Pacific/Tahiti&#x60; -  * &#x60;Pacific/Marquesas&#x60; -  * &#x60;America/Anchorage&#x60; -  * &#x60;Pacific/Gambier&#x60; -  * &#x60;America/Los_Angeles&#x60; -  * &#x60;America/Tijuana&#x60; -  * &#x60;America/Vancouver&#x60; -  * &#x60;America/Whitehorse&#x60; -  * &#x60;Pacific/Pitcairn&#x60; -  * &#x60;America/Dawson_Creek&#x60; -  * &#x60;America/Denver&#x60; -  * &#x60;America/Edmonton&#x60; -  * &#x60;America/Hermosillo&#x60; -  * &#x60;America/Mazatlan&#x60; -  * &#x60;America/Phoenix&#x60; -  * &#x60;America/Yellowknife&#x60; -  * &#x60;America/Belize&#x60; -  * &#x60;America/Chicago&#x60; -  * &#x60;America/Costa_Rica&#x60; -  * &#x60;America/El_Salvador&#x60; -  * &#x60;America/Guatemala&#x60; -  * &#x60;America/Managua&#x60; -  * &#x60;America/Mexico_City&#x60; -  * &#x60;America/Regina&#x60; -  * &#x60;America/Tegucigalpa&#x60; -  * &#x60;America/Winnipeg&#x60; -  * &#x60;Pacific/Galapagos&#x60; -  * &#x60;America/Bogota&#x60; -  * &#x60;America/Cancun&#x60; -  * &#x60;America/Cayman&#x60; -  * &#x60;America/Guayaquil&#x60; -  * &#x60;America/Havana&#x60; -  * &#x60;America/Iqaluit&#x60; -  * &#x60;America/Jamaica&#x60; -  * &#x60;America/Lima&#x60; -  * &#x60;America/Nassau&#x60; -  * &#x60;America/New_York&#x60; -  * &#x60;America/Panama&#x60; -  * &#x60;America/Port-au-Prince&#x60; -  * &#x60;America/Rio_Branco&#x60; -  * &#x60;America/Toronto&#x60; -  * &#x60;Pacific/Easter&#x60; -  * &#x60;America/Caracas&#x60; -  * &#x60;America/Asuncion&#x60; -  * &#x60;America/Barbados&#x60; -  * &#x60;America/Boa_Vista&#x60; -  * &#x60;America/Campo_Grande&#x60; -  * &#x60;America/Cuiaba&#x60; -  * &#x60;America/Curacao&#x60; -  * &#x60;America/Grand_Turk&#x60; -  * &#x60;America/Guyana&#x60; -  * &#x60;America/Halifax&#x60; -  * &#x60;America/La_Paz&#x60; -  * &#x60;America/Manaus&#x60; -  * &#x60;America/Martinique&#x60; -  * &#x60;America/Port_of_Spain&#x60; -  * &#x60;America/Porto_Velho&#x60; -  * &#x60;America/Puerto_Rico&#x60; -  * &#x60;America/Santo_Domingo&#x60; -  * &#x60;America/Thule&#x60; -  * &#x60;Atlantic/Bermuda&#x60; -  * &#x60;America/St_Johns&#x60; -  * &#x60;America/Araguaina&#x60; -  * &#x60;America/Argentina/Buenos_Aires&#x60; -  * &#x60;America/Bahia&#x60; -  * &#x60;America/Belem&#x60; -  * &#x60;America/Cayenne&#x60; -  * &#x60;America/Fortaleza&#x60; -  * &#x60;America/Godthab&#x60; -  * &#x60;America/Maceio&#x60; -  * &#x60;America/Miquelon&#x60; -  * &#x60;America/Montevideo&#x60; -  * &#x60;America/Paramaribo&#x60; -  * &#x60;America/Recife&#x60; -  * &#x60;America/Santiago&#x60; -  * &#x60;America/Sao_Paulo&#x60; -  * &#x60;Antarctica/Palmer&#x60; -  * &#x60;Antarctica/Rothera&#x60; -  * &#x60;Atlantic/Stanley&#x60; -  * &#x60;America/Noronha&#x60; -  * &#x60;Atlantic/South_Georgia&#x60; -  * &#x60;America/Scoresbysund&#x60; -  * &#x60;Atlantic/Azores&#x60; -  * &#x60;Atlantic/Cape_Verde&#x60; -  * &#x60;Africa/Abidjan&#x60; -  * &#x60;Africa/Accra&#x60; -  * &#x60;Africa/Bissau&#x60; -  * &#x60;Africa/Casablanca&#x60; -  * &#x60;Africa/El_Aaiun&#x60; -  * &#x60;Africa/Monrovia&#x60; -  * &#x60;America/Danmarkshavn&#x60; -  * &#x60;Atlantic/Canary&#x60; -  * &#x60;Atlantic/Faroe&#x60; -  * &#x60;Atlantic/Reykjavik&#x60; -  * &#x60;Etc/GMT&#x60; -  * &#x60;Europe/Dublin&#x60; -  * &#x60;Europe/Lisbon&#x60; -  * &#x60;Europe/London&#x60; -  * &#x60;Africa/Algiers&#x60; -  * &#x60;Africa/Ceuta&#x60; -  * &#x60;Africa/Lagos&#x60; -  * &#x60;Africa/Ndjamena&#x60; -  * &#x60;Africa/Tunis&#x60; -  * &#x60;Africa/Windhoek&#x60; -  * &#x60;Europe/Amsterdam&#x60; -  * &#x60;Europe/Andorra&#x60; -  * &#x60;Europe/Belgrade&#x60; -  * &#x60;Europe/Berlin&#x60; -  * &#x60;Europe/Brussels&#x60; -  * &#x60;Europe/Budapest&#x60; -  * &#x60;Europe/Copenhagen&#x60; -  * &#x60;Europe/Gibraltar&#x60; -  * &#x60;Europe/Luxembourg&#x60; -  * &#x60;Europe/Madrid&#x60; -  * &#x60;Europe/Malta&#x60; -  * &#x60;Europe/Monaco&#x60; -  * &#x60;Europe/Oslo&#x60; -  * &#x60;Europe/Paris&#x60; -  * &#x60;Europe/Prague&#x60; -  * &#x60;Europe/Rome&#x60; -  * &#x60;Europe/Stockholm&#x60; -  * &#x60;Europe/Tirane&#x60; -  * &#x60;Europe/Vienna&#x60; -  * &#x60;Europe/Warsaw&#x60; -  * &#x60;Europe/Zurich&#x60; -  * &#x60;Africa/Cairo&#x60; -  * &#x60;Africa/Johannesburg&#x60; -  * &#x60;Africa/Maputo&#x60; -  * &#x60;Africa/Tripoli&#x60; -  * &#x60;Asia/Amman&#x60; -  * &#x60;Asia/Beirut&#x60; -  * &#x60;Asia/Damascus&#x60; -  * &#x60;Asia/Gaza&#x60; -  * &#x60;Asia/Jerusalem&#x60; -  * &#x60;Asia/Nicosia&#x60; -  * &#x60;Europe/Athens&#x60; -  * &#x60;Europe/Bucharest&#x60; -  * &#x60;Europe/Chisinau&#x60; -  * &#x60;Europe/Helsinki&#x60; -  * &#x60;Europe/Istanbul&#x60; -  * &#x60;Europe/Kaliningrad&#x60; -  * &#x60;Europe/Kiev&#x60; -  * &#x60;Europe/Riga&#x60; -  * &#x60;Europe/Sofia&#x60; -  * &#x60;Europe/Tallinn&#x60; -  * &#x60;Europe/Vilnius&#x60; -  * &#x60;Africa/Khartoum&#x60; -  * &#x60;Africa/Nairobi&#x60; -  * &#x60;Antarctica/Syowa&#x60; -  * &#x60;Asia/Baghdad&#x60; -  * &#x60;Asia/Qatar&#x60; -  * &#x60;Asia/Riyadh&#x60; -  * &#x60;Europe/Minsk&#x60; -  * &#x60;Europe/Moscow&#x60; -  * &#x60;Asia/Tehran&#x60; -  * &#x60;Asia/Baku&#x60; -  * &#x60;Asia/Dubai&#x60; -  * &#x60;Asia/Tbilisi&#x60; -  * &#x60;Asia/Yerevan&#x60; -  * &#x60;Europe/Samara&#x60; -  * &#x60;Indian/Mahe&#x60; -  * &#x60;Indian/Mauritius&#x60; -  * &#x60;Indian/Reunion&#x60; -  * &#x60;Asia/Kabul&#x60; -  * &#x60;Antarctica/Mawson&#x60; -  * &#x60;Asia/Aqtau&#x60; -  * &#x60;Asia/Aqtobe&#x60; -  * &#x60;Asia/Ashgabat&#x60; -  * &#x60;Asia/Dushanbe&#x60; -  * &#x60;Asia/Karachi&#x60; -  * &#x60;Asia/Tashkent&#x60; -  * &#x60;Asia/Yekaterinburg&#x60; -  * &#x60;Indian/Kerguelen&#x60; -  * &#x60;Indian/Maldives&#x60; -  * &#x60;Asia/Calcutta&#x60; -  * &#x60;Asia/Kolkata&#x60; -  * &#x60;Asia/Colombo&#x60; -  * &#x60;Asia/Katmandu&#x60; -  * &#x60;Antarctica/Vostok&#x60; -  * &#x60;Asia/Almaty&#x60; -  * &#x60;Asia/Bishkek&#x60; -  * &#x60;Asia/Dhaka&#x60; -  * &#x60;Asia/Omsk&#x60; -  * &#x60;Asia/Thimphu&#x60; -  * &#x60;Indian/Chagos&#x60; -  * &#x60;Asia/Rangoon&#x60; -  * &#x60;Indian/Cocos&#x60; -  * &#x60;Antarctica/Davis&#x60; -  * &#x60;Asia/Bangkok&#x60; -  * &#x60;Asia/Hovd&#x60; -  * &#x60;Asia/Jakarta&#x60; -  * &#x60;Asia/Krasnoyarsk&#x60; -  * &#x60;Asia/Saigon&#x60; -  * &#x60;Indian/Christmas&#x60; -  * &#x60;Antarctica/Casey&#x60; -  * &#x60;Asia/Brunei&#x60; -  * &#x60;Asia/Choibalsan&#x60; -  * &#x60;Asia/Hong_Kong&#x60; -  * &#x60;Asia/Irkutsk&#x60; -  * &#x60;Asia/Kuala_Lumpur&#x60; -  * &#x60;Asia/Macau&#x60; -  * &#x60;Asia/Makassar&#x60; -  * &#x60;Asia/Manila&#x60; -  * &#x60;Asia/Shanghai&#x60; -  * &#x60;Asia/Singapore&#x60; -  * &#x60;Asia/Taipei&#x60; -  * &#x60;Asia/Ulaanbaatar&#x60; -  * &#x60;Australia/Perth&#x60; -  * &#x60;Asia/Pyongyang&#x60; -  * &#x60;Asia/Dili&#x60; -  * &#x60;Asia/Jayapura&#x60; -  * &#x60;Asia/Seoul&#x60; -  * &#x60;Asia/Tokyo&#x60; -  * &#x60;Asia/Yakutsk&#x60; -  * &#x60;Pacific/Palau&#x60; -  * &#x60;Australia/Adelaide&#x60; -  * &#x60;Australia/Darwin&#x60; -  * &#x60;Antarctica/DumontDUrville&#x60; -  * &#x60;Asia/Magadan&#x60; -  * &#x60;Asia/Vladivostok&#x60; -  * &#x60;Australia/Brisbane&#x60; -  * &#x60;Australia/Hobart&#x60; -  * &#x60;Australia/Sydney&#x60; -  * &#x60;Pacific/Chuuk&#x60; -  * &#x60;Pacific/Guam&#x60; -  * &#x60;Pacific/Port_Moresby&#x60; -  * &#x60;Pacific/Efate&#x60; -  * &#x60;Pacific/Guadalcanal&#x60; -  * &#x60;Pacific/Kosrae&#x60; -  * &#x60;Pacific/Norfolk&#x60; -  * &#x60;Pacific/Noumea&#x60; -  * &#x60;Pacific/Pohnpei&#x60; -  * &#x60;Asia/Kamchatka&#x60; -  * &#x60;Pacific/Auckland&#x60; -  * &#x60;Pacific/Fiji&#x60; -  * &#x60;Pacific/Funafuti&#x60; -  * &#x60;Pacific/Kwajalein&#x60; -  * &#x60;Pacific/Majuro&#x60; -  * &#x60;Pacific/Nauru&#x60; -  * &#x60;Pacific/Tarawa&#x60; -  * &#x60;Pacific/Wake&#x60; -  * &#x60;Pacific/Wallis&#x60; -  * &#x60;Pacific/Apia&#x60; -  * &#x60;Pacific/Enderbury&#x60; -  * &#x60;Pacific/Fakaofo&#x60; -  * &#x60;Pacific/Tongatapu&#x60; -  * &#x60;Pacific/Kiritimati&#x60; - | [optional] [default to "Pacific/Niue"]
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexSysConfigPolicyRelationship

`func NewHyperflexSysConfigPolicyRelationship(classId string, objectType string, ) *HyperflexSysConfigPolicyRelationship`

NewHyperflexSysConfigPolicyRelationship instantiates a new HyperflexSysConfigPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSysConfigPolicyRelationshipWithDefaults

`func NewHyperflexSysConfigPolicyRelationshipWithDefaults() *HyperflexSysConfigPolicyRelationship`

NewHyperflexSysConfigPolicyRelationshipWithDefaults instantiates a new HyperflexSysConfigPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexSysConfigPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexSysConfigPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexSysConfigPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexSysConfigPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexSysConfigPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSysConfigPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSysConfigPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexSysConfigPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexSysConfigPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexSysConfigPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexSysConfigPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexSysConfigPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexSysConfigPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexSysConfigPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexSysConfigPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexSysConfigPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexSysConfigPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexSysConfigPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexSysConfigPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexSysConfigPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexSysConfigPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexSysConfigPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexSysConfigPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexSysConfigPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSysConfigPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSysConfigPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexSysConfigPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexSysConfigPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexSysConfigPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexSysConfigPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexSysConfigPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexSysConfigPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexSysConfigPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexSysConfigPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexSysConfigPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexSysConfigPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexSysConfigPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexSysConfigPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexSysConfigPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexSysConfigPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexSysConfigPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexSysConfigPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexSysConfigPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexSysConfigPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexSysConfigPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexSysConfigPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexSysConfigPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexSysConfigPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexSysConfigPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexSysConfigPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexSysConfigPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexSysConfigPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexSysConfigPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexSysConfigPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexSysConfigPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexSysConfigPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexSysConfigPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexSysConfigPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexSysConfigPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexSysConfigPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexSysConfigPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexSysConfigPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexSysConfigPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexSysConfigPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *HyperflexSysConfigPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexSysConfigPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexSysConfigPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexSysConfigPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HyperflexSysConfigPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexSysConfigPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexSysConfigPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexSysConfigPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDnsDomainName

`func (o *HyperflexSysConfigPolicyRelationship) GetDnsDomainName() string`

GetDnsDomainName returns the DnsDomainName field if non-nil, zero value otherwise.

### GetDnsDomainNameOk

`func (o *HyperflexSysConfigPolicyRelationship) GetDnsDomainNameOk() (*string, bool)`

GetDnsDomainNameOk returns a tuple with the DnsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainName

`func (o *HyperflexSysConfigPolicyRelationship) SetDnsDomainName(v string)`

SetDnsDomainName sets DnsDomainName field to given value.

### HasDnsDomainName

`func (o *HyperflexSysConfigPolicyRelationship) HasDnsDomainName() bool`

HasDnsDomainName returns a boolean if a field has been set.

### GetDnsServers

`func (o *HyperflexSysConfigPolicyRelationship) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *HyperflexSysConfigPolicyRelationship) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *HyperflexSysConfigPolicyRelationship) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *HyperflexSysConfigPolicyRelationship) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetNtpServers

`func (o *HyperflexSysConfigPolicyRelationship) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *HyperflexSysConfigPolicyRelationship) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *HyperflexSysConfigPolicyRelationship) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *HyperflexSysConfigPolicyRelationship) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetTimezone

`func (o *HyperflexSysConfigPolicyRelationship) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *HyperflexSysConfigPolicyRelationship) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *HyperflexSysConfigPolicyRelationship) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *HyperflexSysConfigPolicyRelationship) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexSysConfigPolicyRelationship) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexSysConfigPolicyRelationship) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexSysConfigPolicyRelationship) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexSysConfigPolicyRelationship) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexSysConfigPolicyRelationship) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexSysConfigPolicyRelationship) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexSysConfigPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexSysConfigPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexSysConfigPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexSysConfigPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


