# NtpPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ntp.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ntp.Policy"]
**AuthenticatedNtpServers** | Pointer to [**[]NtpAuthNtpServer**](NtpAuthNtpServer.md) |  | [optional] 
**Enabled** | Pointer to **bool** | State of NTP service on the endpoint. | [optional] [default to true]
**NtpServers** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **string** | Timezone of services on the endpoint. * &#x60;Pacific/Niue&#x60; -  * &#x60;Pacific/Pago_Pago&#x60; -  * &#x60;Pacific/Honolulu&#x60; -  * &#x60;Pacific/Rarotonga&#x60; -  * &#x60;Pacific/Tahiti&#x60; -  * &#x60;Pacific/Marquesas&#x60; -  * &#x60;America/Anchorage&#x60; -  * &#x60;Pacific/Gambier&#x60; -  * &#x60;America/Los_Angeles&#x60; -  * &#x60;America/Tijuana&#x60; -  * &#x60;America/Vancouver&#x60; -  * &#x60;America/Whitehorse&#x60; -  * &#x60;Pacific/Pitcairn&#x60; -  * &#x60;America/Dawson_Creek&#x60; -  * &#x60;America/Denver&#x60; -  * &#x60;America/Edmonton&#x60; -  * &#x60;America/Hermosillo&#x60; -  * &#x60;America/Mazatlan&#x60; -  * &#x60;America/Phoenix&#x60; -  * &#x60;America/Yellowknife&#x60; -  * &#x60;America/Belize&#x60; -  * &#x60;America/Chicago&#x60; -  * &#x60;America/Costa_Rica&#x60; -  * &#x60;America/El_Salvador&#x60; -  * &#x60;America/Guatemala&#x60; -  * &#x60;America/Managua&#x60; -  * &#x60;America/Mexico_City&#x60; -  * &#x60;America/Regina&#x60; -  * &#x60;America/Tegucigalpa&#x60; -  * &#x60;America/Winnipeg&#x60; -  * &#x60;Pacific/Galapagos&#x60; -  * &#x60;America/Bogota&#x60; -  * &#x60;America/Cancun&#x60; -  * &#x60;America/Cayman&#x60; -  * &#x60;America/Guayaquil&#x60; -  * &#x60;America/Havana&#x60; -  * &#x60;America/Iqaluit&#x60; -  * &#x60;America/Jamaica&#x60; -  * &#x60;America/Lima&#x60; -  * &#x60;America/Nassau&#x60; -  * &#x60;America/New_York&#x60; -  * &#x60;America/Panama&#x60; -  * &#x60;America/Port-au-Prince&#x60; -  * &#x60;America/Rio_Branco&#x60; -  * &#x60;America/Toronto&#x60; -  * &#x60;Pacific/Easter&#x60; -  * &#x60;America/Caracas&#x60; -  * &#x60;America/Asuncion&#x60; -  * &#x60;America/Barbados&#x60; -  * &#x60;America/Boa_Vista&#x60; -  * &#x60;America/Campo_Grande&#x60; -  * &#x60;America/Cuiaba&#x60; -  * &#x60;America/Curacao&#x60; -  * &#x60;America/Grand_Turk&#x60; -  * &#x60;America/Guyana&#x60; -  * &#x60;America/Halifax&#x60; -  * &#x60;America/La_Paz&#x60; -  * &#x60;America/Manaus&#x60; -  * &#x60;America/Martinique&#x60; -  * &#x60;America/Port_of_Spain&#x60; -  * &#x60;America/Porto_Velho&#x60; -  * &#x60;America/Puerto_Rico&#x60; -  * &#x60;America/Santo_Domingo&#x60; -  * &#x60;America/Thule&#x60; -  * &#x60;Atlantic/Bermuda&#x60; -  * &#x60;America/St_Johns&#x60; -  * &#x60;America/Araguaina&#x60; -  * &#x60;America/Argentina/Buenos_Aires&#x60; -  * &#x60;America/Bahia&#x60; -  * &#x60;America/Belem&#x60; -  * &#x60;America/Cayenne&#x60; -  * &#x60;America/Fortaleza&#x60; -  * &#x60;America/Godthab&#x60; -  * &#x60;America/Maceio&#x60; -  * &#x60;America/Miquelon&#x60; -  * &#x60;America/Montevideo&#x60; -  * &#x60;America/Paramaribo&#x60; -  * &#x60;America/Recife&#x60; -  * &#x60;America/Santiago&#x60; -  * &#x60;America/Sao_Paulo&#x60; -  * &#x60;Antarctica/Palmer&#x60; -  * &#x60;Antarctica/Rothera&#x60; -  * &#x60;Atlantic/Stanley&#x60; -  * &#x60;America/Noronha&#x60; -  * &#x60;Atlantic/South_Georgia&#x60; -  * &#x60;America/Scoresbysund&#x60; -  * &#x60;Atlantic/Azores&#x60; -  * &#x60;Atlantic/Cape_Verde&#x60; -  * &#x60;Africa/Abidjan&#x60; -  * &#x60;Africa/Accra&#x60; -  * &#x60;Africa/Bissau&#x60; -  * &#x60;Africa/Casablanca&#x60; -  * &#x60;Africa/El_Aaiun&#x60; -  * &#x60;Africa/Monrovia&#x60; -  * &#x60;America/Danmarkshavn&#x60; -  * &#x60;Atlantic/Canary&#x60; -  * &#x60;Atlantic/Faroe&#x60; -  * &#x60;Atlantic/Reykjavik&#x60; -  * &#x60;Etc/GMT&#x60; -  * &#x60;Europe/Dublin&#x60; -  * &#x60;Europe/Lisbon&#x60; -  * &#x60;Europe/London&#x60; -  * &#x60;Africa/Algiers&#x60; -  * &#x60;Africa/Ceuta&#x60; -  * &#x60;Africa/Lagos&#x60; -  * &#x60;Africa/Ndjamena&#x60; -  * &#x60;Africa/Tunis&#x60; -  * &#x60;Africa/Windhoek&#x60; -  * &#x60;Europe/Amsterdam&#x60; -  * &#x60;Europe/Andorra&#x60; -  * &#x60;Europe/Belgrade&#x60; -  * &#x60;Europe/Berlin&#x60; -  * &#x60;Europe/Brussels&#x60; -  * &#x60;Europe/Budapest&#x60; -  * &#x60;Europe/Copenhagen&#x60; -  * &#x60;Europe/Gibraltar&#x60; -  * &#x60;Europe/Luxembourg&#x60; -  * &#x60;Europe/Madrid&#x60; -  * &#x60;Europe/Malta&#x60; -  * &#x60;Europe/Monaco&#x60; -  * &#x60;Europe/Oslo&#x60; -  * &#x60;Europe/Paris&#x60; -  * &#x60;Europe/Prague&#x60; -  * &#x60;Europe/Rome&#x60; -  * &#x60;Europe/Stockholm&#x60; -  * &#x60;Europe/Tirane&#x60; -  * &#x60;Europe/Vienna&#x60; -  * &#x60;Europe/Warsaw&#x60; -  * &#x60;Europe/Zurich&#x60; -  * &#x60;Africa/Cairo&#x60; -  * &#x60;Africa/Johannesburg&#x60; -  * &#x60;Africa/Maputo&#x60; -  * &#x60;Africa/Tripoli&#x60; -  * &#x60;Asia/Amman&#x60; -  * &#x60;Asia/Beirut&#x60; -  * &#x60;Asia/Damascus&#x60; -  * &#x60;Asia/Gaza&#x60; -  * &#x60;Asia/Jerusalem&#x60; -  * &#x60;Asia/Nicosia&#x60; -  * &#x60;Europe/Athens&#x60; -  * &#x60;Europe/Bucharest&#x60; -  * &#x60;Europe/Chisinau&#x60; -  * &#x60;Europe/Helsinki&#x60; -  * &#x60;Europe/Istanbul&#x60; -  * &#x60;Europe/Kaliningrad&#x60; -  * &#x60;Europe/Kiev&#x60; -  * &#x60;Europe/Riga&#x60; -  * &#x60;Europe/Sofia&#x60; -  * &#x60;Europe/Tallinn&#x60; -  * &#x60;Europe/Vilnius&#x60; -  * &#x60;Africa/Khartoum&#x60; -  * &#x60;Africa/Nairobi&#x60; -  * &#x60;Antarctica/Syowa&#x60; -  * &#x60;Asia/Baghdad&#x60; -  * &#x60;Asia/Qatar&#x60; -  * &#x60;Asia/Riyadh&#x60; -  * &#x60;Europe/Minsk&#x60; -  * &#x60;Europe/Moscow&#x60; -  * &#x60;Asia/Tehran&#x60; -  * &#x60;Asia/Baku&#x60; -  * &#x60;Asia/Dubai&#x60; -  * &#x60;Asia/Tbilisi&#x60; -  * &#x60;Asia/Yerevan&#x60; -  * &#x60;Europe/Samara&#x60; -  * &#x60;Indian/Mahe&#x60; -  * &#x60;Indian/Mauritius&#x60; -  * &#x60;Indian/Reunion&#x60; -  * &#x60;Asia/Kabul&#x60; -  * &#x60;Antarctica/Mawson&#x60; -  * &#x60;Asia/Aqtau&#x60; -  * &#x60;Asia/Aqtobe&#x60; -  * &#x60;Asia/Ashgabat&#x60; -  * &#x60;Asia/Dushanbe&#x60; -  * &#x60;Asia/Karachi&#x60; -  * &#x60;Asia/Tashkent&#x60; -  * &#x60;Asia/Yekaterinburg&#x60; -  * &#x60;Indian/Kerguelen&#x60; -  * &#x60;Indian/Maldives&#x60; -  * &#x60;Asia/Calcutta&#x60; -  * &#x60;Asia/Kolkata&#x60; -  * &#x60;Asia/Colombo&#x60; -  * &#x60;Asia/Katmandu&#x60; -  * &#x60;Antarctica/Vostok&#x60; -  * &#x60;Asia/Almaty&#x60; -  * &#x60;Asia/Bishkek&#x60; -  * &#x60;Asia/Dhaka&#x60; -  * &#x60;Asia/Omsk&#x60; -  * &#x60;Asia/Thimphu&#x60; -  * &#x60;Indian/Chagos&#x60; -  * &#x60;Asia/Rangoon&#x60; -  * &#x60;Indian/Cocos&#x60; -  * &#x60;Antarctica/Davis&#x60; -  * &#x60;Asia/Bangkok&#x60; -  * &#x60;Asia/Hovd&#x60; -  * &#x60;Asia/Jakarta&#x60; -  * &#x60;Asia/Krasnoyarsk&#x60; -  * &#x60;Asia/Saigon&#x60; -  * &#x60;Indian/Christmas&#x60; -  * &#x60;Antarctica/Casey&#x60; -  * &#x60;Asia/Brunei&#x60; -  * &#x60;Asia/Choibalsan&#x60; -  * &#x60;Asia/Hong_Kong&#x60; -  * &#x60;Asia/Irkutsk&#x60; -  * &#x60;Asia/Kuala_Lumpur&#x60; -  * &#x60;Asia/Macau&#x60; -  * &#x60;Asia/Makassar&#x60; -  * &#x60;Asia/Manila&#x60; -  * &#x60;Asia/Shanghai&#x60; -  * &#x60;Asia/Singapore&#x60; -  * &#x60;Asia/Taipei&#x60; -  * &#x60;Asia/Ulaanbaatar&#x60; -  * &#x60;Australia/Perth&#x60; -  * &#x60;Asia/Pyongyang&#x60; -  * &#x60;Asia/Dili&#x60; -  * &#x60;Asia/Jayapura&#x60; -  * &#x60;Asia/Seoul&#x60; -  * &#x60;Asia/Tokyo&#x60; -  * &#x60;Asia/Yakutsk&#x60; -  * &#x60;Pacific/Palau&#x60; -  * &#x60;Australia/Adelaide&#x60; -  * &#x60;Australia/Darwin&#x60; -  * &#x60;Antarctica/DumontDUrville&#x60; -  * &#x60;Asia/Magadan&#x60; -  * &#x60;Asia/Vladivostok&#x60; -  * &#x60;Australia/Brisbane&#x60; -  * &#x60;Australia/Hobart&#x60; -  * &#x60;Australia/Sydney&#x60; -  * &#x60;Pacific/Chuuk&#x60; -  * &#x60;Pacific/Guam&#x60; -  * &#x60;Pacific/Port_Moresby&#x60; -  * &#x60;Pacific/Efate&#x60; -  * &#x60;Pacific/Guadalcanal&#x60; -  * &#x60;Pacific/Kosrae&#x60; -  * &#x60;Pacific/Norfolk&#x60; -  * &#x60;Pacific/Noumea&#x60; -  * &#x60;Pacific/Pohnpei&#x60; -  * &#x60;Asia/Kamchatka&#x60; -  * &#x60;Pacific/Auckland&#x60; -  * &#x60;Pacific/Fiji&#x60; -  * &#x60;Pacific/Funafuti&#x60; -  * &#x60;Pacific/Kwajalein&#x60; -  * &#x60;Pacific/Majuro&#x60; -  * &#x60;Pacific/Nauru&#x60; -  * &#x60;Pacific/Tarawa&#x60; -  * &#x60;Pacific/Wake&#x60; -  * &#x60;Pacific/Wallis&#x60; -  * &#x60;Pacific/Apia&#x60; -  * &#x60;Pacific/Enderbury&#x60; -  * &#x60;Pacific/Fakaofo&#x60; -  * &#x60;Pacific/Tongatapu&#x60; -  * &#x60;Pacific/Kiritimati&#x60; - | [optional] [default to "Pacific/Niue"]
**ApplianceAccount** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewNtpPolicyAllOf

`func NewNtpPolicyAllOf(classId string, objectType string, ) *NtpPolicyAllOf`

NewNtpPolicyAllOf instantiates a new NtpPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpPolicyAllOfWithDefaults

`func NewNtpPolicyAllOfWithDefaults() *NtpPolicyAllOf`

NewNtpPolicyAllOfWithDefaults instantiates a new NtpPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NtpPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NtpPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NtpPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NtpPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NtpPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NtpPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticatedNtpServers

`func (o *NtpPolicyAllOf) GetAuthenticatedNtpServers() []NtpAuthNtpServer`

GetAuthenticatedNtpServers returns the AuthenticatedNtpServers field if non-nil, zero value otherwise.

### GetAuthenticatedNtpServersOk

`func (o *NtpPolicyAllOf) GetAuthenticatedNtpServersOk() (*[]NtpAuthNtpServer, bool)`

GetAuthenticatedNtpServersOk returns a tuple with the AuthenticatedNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedNtpServers

`func (o *NtpPolicyAllOf) SetAuthenticatedNtpServers(v []NtpAuthNtpServer)`

SetAuthenticatedNtpServers sets AuthenticatedNtpServers field to given value.

### HasAuthenticatedNtpServers

`func (o *NtpPolicyAllOf) HasAuthenticatedNtpServers() bool`

HasAuthenticatedNtpServers returns a boolean if a field has been set.

### SetAuthenticatedNtpServersNil

`func (o *NtpPolicyAllOf) SetAuthenticatedNtpServersNil(b bool)`

 SetAuthenticatedNtpServersNil sets the value for AuthenticatedNtpServers to be an explicit nil

### UnsetAuthenticatedNtpServers
`func (o *NtpPolicyAllOf) UnsetAuthenticatedNtpServers()`

UnsetAuthenticatedNtpServers ensures that no value is present for AuthenticatedNtpServers, not even an explicit nil
### GetEnabled

`func (o *NtpPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NtpPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NtpPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NtpPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNtpServers

`func (o *NtpPolicyAllOf) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *NtpPolicyAllOf) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *NtpPolicyAllOf) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *NtpPolicyAllOf) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *NtpPolicyAllOf) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *NtpPolicyAllOf) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetTimezone

`func (o *NtpPolicyAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NtpPolicyAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NtpPolicyAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *NtpPolicyAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetApplianceAccount

`func (o *NtpPolicyAllOf) GetApplianceAccount() IamAccountRelationship`

GetApplianceAccount returns the ApplianceAccount field if non-nil, zero value otherwise.

### GetApplianceAccountOk

`func (o *NtpPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool)`

GetApplianceAccountOk returns a tuple with the ApplianceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceAccount

`func (o *NtpPolicyAllOf) SetApplianceAccount(v IamAccountRelationship)`

SetApplianceAccount sets ApplianceAccount field to given value.

### HasApplianceAccount

`func (o *NtpPolicyAllOf) HasApplianceAccount() bool`

HasApplianceAccount returns a boolean if a field has been set.

### GetOrganization

`func (o *NtpPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NtpPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NtpPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *NtpPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *NtpPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *NtpPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *NtpPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *NtpPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *NtpPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *NtpPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


