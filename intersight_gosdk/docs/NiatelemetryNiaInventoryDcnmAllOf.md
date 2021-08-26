# NiatelemetryNiaInventoryDcnmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryDcnm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryDcnm"]
**Dev** | Pointer to **bool** | Returns the value of the dev Field. | [optional] 
**EpldImageCount** | Pointer to **int64** | Number of EPLD images uploaded to DCNM. | [optional] 
**HaEnabled** | Pointer to **bool** | Returns the value of the haEnabled field. | [optional] 
**HaReplicationStatus** | Pointer to **string** | Returns the value of the haReplicationStatus field. | [optional] 
**Install** | Pointer to **string** | Returns the value of the install field. | [optional] 
**IsIsnConfigured** | Pointer to **bool** | Returns true if ISN is configured. | [optional] 
**IsMediaController** | Pointer to **bool** | Returns the value of the isMediaController field. | [optional] 
**IsSmartLicenseEnabled** | Pointer to **bool** | Returns true if the Smart license is enabled and is in use. | [optional] 
**NumFabrics** | Pointer to **int64** | Returns total number of fabrics in DCNM set-up. | [optional] 
**NumFabricsInMsd** | Pointer to **int64** | Returns the number of fabrics in msd. | [optional] 
**NumIngressReplicationFabrics** | Pointer to **int64** | Returns the number of fabrics that have ingress replication type. | [optional] 
**NumLocalUsers** | Pointer to **int64** | Returns the number of local users other than admin user. | [optional] 
**NumMsd** | Pointer to **int64** | Returns the number of MSD fabrics. | [optional] 
**NumSviVrfCount** | Pointer to **int64** | Returns the number of svi interfaces configured for VRF vlans. | [optional] 
**NumTrmEnabledCount** | Pointer to **int64** | Returns the number of links where TRM is enabled. | [optional] 
**NumUpgUsers** | Pointer to **int64** | Number of users who have upgrade privileges excluding the admin. | [optional] 
**NxosImageCount** | Pointer to **int64** | Number of NXOS images uploaded to DCNM. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**SiteName** | Pointer to **string** | Name of fabric domain of the controller. | [optional] 
**UnderlayPeeringActiveLinksCount** | Pointer to **int64** | Returns the number of underlay peering active links. | [optional] 
**UpgJobCount** | Pointer to **int64** | Number of upgrade jobs configured on DCNM. | [optional] 
**UpgStatus** | Pointer to **string** | Upgrade status of jobs created on DCNM. | [optional] 
**Version** | Pointer to **string** | Returns the value of the version field. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryDcnmAllOf

`func NewNiatelemetryNiaInventoryDcnmAllOf(classId string, objectType string, ) *NiatelemetryNiaInventoryDcnmAllOf`

NewNiatelemetryNiaInventoryDcnmAllOf instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults

`func NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults() *NiatelemetryNiaInventoryDcnmAllOf`

NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDev

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDev() bool`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDevOk() (*bool, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetDev(v bool)`

SetDev sets Dev field to given value.

### HasDev

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetEpldImageCount() int64`

GetEpldImageCount returns the EpldImageCount field if non-nil, zero value otherwise.

### GetEpldImageCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetEpldImageCountOk() (*int64, bool)`

GetEpldImageCountOk returns a tuple with the EpldImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetEpldImageCount(v int64)`

SetEpldImageCount sets EpldImageCount field to given value.

### HasEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasEpldImageCount() bool`

HasEpldImageCount returns a boolean if a field has been set.

### GetHaEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatus() string`

GetHaReplicationStatus returns the HaReplicationStatus field if non-nil, zero value otherwise.

### GetHaReplicationStatusOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatusOk() (*string, bool)`

GetHaReplicationStatusOk returns a tuple with the HaReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaReplicationStatus(v string)`

SetHaReplicationStatus sets HaReplicationStatus field to given value.

### HasHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaReplicationStatus() bool`

HasHaReplicationStatus returns a boolean if a field has been set.

### GetInstall

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsIsnConfigured() bool`

GetIsIsnConfigured returns the IsIsnConfigured field if non-nil, zero value otherwise.

### GetIsIsnConfiguredOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsIsnConfiguredOk() (*bool, bool)`

GetIsIsnConfiguredOk returns a tuple with the IsIsnConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsIsnConfigured(v bool)`

SetIsIsnConfigured sets IsIsnConfigured field to given value.

### HasIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsIsnConfigured() bool`

HasIsIsnConfigured returns a boolean if a field has been set.

### GetIsMediaController

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaController() bool`

GetIsMediaController returns the IsMediaController field if non-nil, zero value otherwise.

### GetIsMediaControllerOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaControllerOk() (*bool, bool)`

GetIsMediaControllerOk returns a tuple with the IsMediaController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMediaController

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsMediaController(v bool)`

SetIsMediaController sets IsMediaController field to given value.

### HasIsMediaController

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsMediaController() bool`

HasIsMediaController returns a boolean if a field has been set.

### GetIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsSmartLicenseEnabled() bool`

GetIsSmartLicenseEnabled returns the IsSmartLicenseEnabled field if non-nil, zero value otherwise.

### GetIsSmartLicenseEnabledOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsSmartLicenseEnabledOk() (*bool, bool)`

GetIsSmartLicenseEnabledOk returns a tuple with the IsSmartLicenseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsSmartLicenseEnabled(v bool)`

SetIsSmartLicenseEnabled sets IsSmartLicenseEnabled field to given value.

### HasIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsSmartLicenseEnabled() bool`

HasIsSmartLicenseEnabled returns a boolean if a field has been set.

### GetNumFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabrics() int64`

GetNumFabrics returns the NumFabrics field if non-nil, zero value otherwise.

### GetNumFabricsOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsOk() (*int64, bool)`

GetNumFabricsOk returns a tuple with the NumFabrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabrics(v int64)`

SetNumFabrics sets NumFabrics field to given value.

### HasNumFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabrics() bool`

HasNumFabrics returns a boolean if a field has been set.

### GetNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsd() int64`

GetNumFabricsInMsd returns the NumFabricsInMsd field if non-nil, zero value otherwise.

### GetNumFabricsInMsdOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsdOk() (*int64, bool)`

GetNumFabricsInMsdOk returns a tuple with the NumFabricsInMsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabricsInMsd(v int64)`

SetNumFabricsInMsd sets NumFabricsInMsd field to given value.

### HasNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabricsInMsd() bool`

HasNumFabricsInMsd returns a boolean if a field has been set.

### GetNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumIngressReplicationFabrics() int64`

GetNumIngressReplicationFabrics returns the NumIngressReplicationFabrics field if non-nil, zero value otherwise.

### GetNumIngressReplicationFabricsOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumIngressReplicationFabricsOk() (*int64, bool)`

GetNumIngressReplicationFabricsOk returns a tuple with the NumIngressReplicationFabrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumIngressReplicationFabrics(v int64)`

SetNumIngressReplicationFabrics sets NumIngressReplicationFabrics field to given value.

### HasNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumIngressReplicationFabrics() bool`

HasNumIngressReplicationFabrics returns a boolean if a field has been set.

### GetNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsers() int64`

GetNumLocalUsers returns the NumLocalUsers field if non-nil, zero value otherwise.

### GetNumLocalUsersOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsersOk() (*int64, bool)`

GetNumLocalUsersOk returns a tuple with the NumLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumLocalUsers(v int64)`

SetNumLocalUsers sets NumLocalUsers field to given value.

### HasNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumLocalUsers() bool`

HasNumLocalUsers returns a boolean if a field has been set.

### GetNumMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumMsd() int64`

GetNumMsd returns the NumMsd field if non-nil, zero value otherwise.

### GetNumMsdOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumMsdOk() (*int64, bool)`

GetNumMsdOk returns a tuple with the NumMsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumMsd(v int64)`

SetNumMsd sets NumMsd field to given value.

### HasNumMsd

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumMsd() bool`

HasNumMsd returns a boolean if a field has been set.

### GetNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumSviVrfCount() int64`

GetNumSviVrfCount returns the NumSviVrfCount field if non-nil, zero value otherwise.

### GetNumSviVrfCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumSviVrfCountOk() (*int64, bool)`

GetNumSviVrfCountOk returns a tuple with the NumSviVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumSviVrfCount(v int64)`

SetNumSviVrfCount sets NumSviVrfCount field to given value.

### HasNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumSviVrfCount() bool`

HasNumSviVrfCount returns a boolean if a field has been set.

### GetNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumTrmEnabledCount() int64`

GetNumTrmEnabledCount returns the NumTrmEnabledCount field if non-nil, zero value otherwise.

### GetNumTrmEnabledCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumTrmEnabledCountOk() (*int64, bool)`

GetNumTrmEnabledCountOk returns a tuple with the NumTrmEnabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumTrmEnabledCount(v int64)`

SetNumTrmEnabledCount sets NumTrmEnabledCount field to given value.

### HasNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumTrmEnabledCount() bool`

HasNumTrmEnabledCount returns a boolean if a field has been set.

### GetNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumUpgUsers() int64`

GetNumUpgUsers returns the NumUpgUsers field if non-nil, zero value otherwise.

### GetNumUpgUsersOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumUpgUsersOk() (*int64, bool)`

GetNumUpgUsersOk returns a tuple with the NumUpgUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumUpgUsers(v int64)`

SetNumUpgUsers sets NumUpgUsers field to given value.

### HasNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumUpgUsers() bool`

HasNumUpgUsers returns a boolean if a field has been set.

### GetNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNxosImageCount() int64`

GetNxosImageCount returns the NxosImageCount field if non-nil, zero value otherwise.

### GetNxosImageCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNxosImageCountOk() (*int64, bool)`

GetNxosImageCountOk returns a tuple with the NxosImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNxosImageCount(v int64)`

SetNxosImageCount sets NxosImageCount field to given value.

### HasNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNxosImageCount() bool`

HasNxosImageCount returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUnderlayPeeringActiveLinksCount() int64`

GetUnderlayPeeringActiveLinksCount returns the UnderlayPeeringActiveLinksCount field if non-nil, zero value otherwise.

### GetUnderlayPeeringActiveLinksCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUnderlayPeeringActiveLinksCountOk() (*int64, bool)`

GetUnderlayPeeringActiveLinksCountOk returns a tuple with the UnderlayPeeringActiveLinksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUnderlayPeeringActiveLinksCount(v int64)`

SetUnderlayPeeringActiveLinksCount sets UnderlayPeeringActiveLinksCount field to given value.

### HasUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUnderlayPeeringActiveLinksCount() bool`

HasUnderlayPeeringActiveLinksCount returns a boolean if a field has been set.

### GetUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgJobCount() int64`

GetUpgJobCount returns the UpgJobCount field if non-nil, zero value otherwise.

### GetUpgJobCountOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgJobCountOk() (*int64, bool)`

GetUpgJobCountOk returns a tuple with the UpgJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUpgJobCount(v int64)`

SetUpgJobCount sets UpgJobCount field to given value.

### HasUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUpgJobCount() bool`

HasUpgJobCount returns a boolean if a field has been set.

### GetUpgStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgStatus() string`

GetUpgStatus returns the UpgStatus field if non-nil, zero value otherwise.

### GetUpgStatusOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgStatusOk() (*string, bool)`

GetUpgStatusOk returns a tuple with the UpgStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUpgStatus(v string)`

SetUpgStatus sets UpgStatus field to given value.

### HasUpgStatus

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUpgStatus() bool`

HasUpgStatus returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnmAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


