# NiatelemetryNiaInventoryDcnm

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

### NewNiatelemetryNiaInventoryDcnm

`func NewNiatelemetryNiaInventoryDcnm(classId string, objectType string, ) *NiatelemetryNiaInventoryDcnm`

NewNiatelemetryNiaInventoryDcnm instantiates a new NiatelemetryNiaInventoryDcnm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryDcnmWithDefaults

`func NewNiatelemetryNiaInventoryDcnmWithDefaults() *NiatelemetryNiaInventoryDcnm`

NewNiatelemetryNiaInventoryDcnmWithDefaults instantiates a new NiatelemetryNiaInventoryDcnm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryDcnm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryDcnm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryDcnm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryDcnm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryDcnm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryDcnm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDev

`func (o *NiatelemetryNiaInventoryDcnm) GetDev() bool`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *NiatelemetryNiaInventoryDcnm) GetDevOk() (*bool, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *NiatelemetryNiaInventoryDcnm) SetDev(v bool)`

SetDev sets Dev field to given value.

### HasDev

`func (o *NiatelemetryNiaInventoryDcnm) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnm) GetEpldImageCount() int64`

GetEpldImageCount returns the EpldImageCount field if non-nil, zero value otherwise.

### GetEpldImageCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetEpldImageCountOk() (*int64, bool)`

GetEpldImageCountOk returns a tuple with the EpldImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnm) SetEpldImageCount(v int64)`

SetEpldImageCount sets EpldImageCount field to given value.

### HasEpldImageCount

`func (o *NiatelemetryNiaInventoryDcnm) HasEpldImageCount() bool`

HasEpldImageCount returns a boolean if a field has been set.

### GetHaEnabled

`func (o *NiatelemetryNiaInventoryDcnm) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *NiatelemetryNiaInventoryDcnm) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *NiatelemetryNiaInventoryDcnm) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *NiatelemetryNiaInventoryDcnm) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnm) GetHaReplicationStatus() string`

GetHaReplicationStatus returns the HaReplicationStatus field if non-nil, zero value otherwise.

### GetHaReplicationStatusOk

`func (o *NiatelemetryNiaInventoryDcnm) GetHaReplicationStatusOk() (*string, bool)`

GetHaReplicationStatusOk returns a tuple with the HaReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnm) SetHaReplicationStatus(v string)`

SetHaReplicationStatus sets HaReplicationStatus field to given value.

### HasHaReplicationStatus

`func (o *NiatelemetryNiaInventoryDcnm) HasHaReplicationStatus() bool`

HasHaReplicationStatus returns a boolean if a field has been set.

### GetInstall

`func (o *NiatelemetryNiaInventoryDcnm) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *NiatelemetryNiaInventoryDcnm) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *NiatelemetryNiaInventoryDcnm) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *NiatelemetryNiaInventoryDcnm) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnm) GetIsIsnConfigured() bool`

GetIsIsnConfigured returns the IsIsnConfigured field if non-nil, zero value otherwise.

### GetIsIsnConfiguredOk

`func (o *NiatelemetryNiaInventoryDcnm) GetIsIsnConfiguredOk() (*bool, bool)`

GetIsIsnConfiguredOk returns a tuple with the IsIsnConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnm) SetIsIsnConfigured(v bool)`

SetIsIsnConfigured sets IsIsnConfigured field to given value.

### HasIsIsnConfigured

`func (o *NiatelemetryNiaInventoryDcnm) HasIsIsnConfigured() bool`

HasIsIsnConfigured returns a boolean if a field has been set.

### GetIsMediaController

`func (o *NiatelemetryNiaInventoryDcnm) GetIsMediaController() bool`

GetIsMediaController returns the IsMediaController field if non-nil, zero value otherwise.

### GetIsMediaControllerOk

`func (o *NiatelemetryNiaInventoryDcnm) GetIsMediaControllerOk() (*bool, bool)`

GetIsMediaControllerOk returns a tuple with the IsMediaController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMediaController

`func (o *NiatelemetryNiaInventoryDcnm) SetIsMediaController(v bool)`

SetIsMediaController sets IsMediaController field to given value.

### HasIsMediaController

`func (o *NiatelemetryNiaInventoryDcnm) HasIsMediaController() bool`

HasIsMediaController returns a boolean if a field has been set.

### GetIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnm) GetIsSmartLicenseEnabled() bool`

GetIsSmartLicenseEnabled returns the IsSmartLicenseEnabled field if non-nil, zero value otherwise.

### GetIsSmartLicenseEnabledOk

`func (o *NiatelemetryNiaInventoryDcnm) GetIsSmartLicenseEnabledOk() (*bool, bool)`

GetIsSmartLicenseEnabledOk returns a tuple with the IsSmartLicenseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnm) SetIsSmartLicenseEnabled(v bool)`

SetIsSmartLicenseEnabled sets IsSmartLicenseEnabled field to given value.

### HasIsSmartLicenseEnabled

`func (o *NiatelemetryNiaInventoryDcnm) HasIsSmartLicenseEnabled() bool`

HasIsSmartLicenseEnabled returns a boolean if a field has been set.

### GetNumFabrics

`func (o *NiatelemetryNiaInventoryDcnm) GetNumFabrics() int64`

GetNumFabrics returns the NumFabrics field if non-nil, zero value otherwise.

### GetNumFabricsOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumFabricsOk() (*int64, bool)`

GetNumFabricsOk returns a tuple with the NumFabrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFabrics

`func (o *NiatelemetryNiaInventoryDcnm) SetNumFabrics(v int64)`

SetNumFabrics sets NumFabrics field to given value.

### HasNumFabrics

`func (o *NiatelemetryNiaInventoryDcnm) HasNumFabrics() bool`

HasNumFabrics returns a boolean if a field has been set.

### GetNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnm) GetNumFabricsInMsd() int64`

GetNumFabricsInMsd returns the NumFabricsInMsd field if non-nil, zero value otherwise.

### GetNumFabricsInMsdOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumFabricsInMsdOk() (*int64, bool)`

GetNumFabricsInMsdOk returns a tuple with the NumFabricsInMsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnm) SetNumFabricsInMsd(v int64)`

SetNumFabricsInMsd sets NumFabricsInMsd field to given value.

### HasNumFabricsInMsd

`func (o *NiatelemetryNiaInventoryDcnm) HasNumFabricsInMsd() bool`

HasNumFabricsInMsd returns a boolean if a field has been set.

### GetNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnm) GetNumIngressReplicationFabrics() int64`

GetNumIngressReplicationFabrics returns the NumIngressReplicationFabrics field if non-nil, zero value otherwise.

### GetNumIngressReplicationFabricsOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumIngressReplicationFabricsOk() (*int64, bool)`

GetNumIngressReplicationFabricsOk returns a tuple with the NumIngressReplicationFabrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnm) SetNumIngressReplicationFabrics(v int64)`

SetNumIngressReplicationFabrics sets NumIngressReplicationFabrics field to given value.

### HasNumIngressReplicationFabrics

`func (o *NiatelemetryNiaInventoryDcnm) HasNumIngressReplicationFabrics() bool`

HasNumIngressReplicationFabrics returns a boolean if a field has been set.

### GetNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnm) GetNumLocalUsers() int64`

GetNumLocalUsers returns the NumLocalUsers field if non-nil, zero value otherwise.

### GetNumLocalUsersOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumLocalUsersOk() (*int64, bool)`

GetNumLocalUsersOk returns a tuple with the NumLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnm) SetNumLocalUsers(v int64)`

SetNumLocalUsers sets NumLocalUsers field to given value.

### HasNumLocalUsers

`func (o *NiatelemetryNiaInventoryDcnm) HasNumLocalUsers() bool`

HasNumLocalUsers returns a boolean if a field has been set.

### GetNumMsd

`func (o *NiatelemetryNiaInventoryDcnm) GetNumMsd() int64`

GetNumMsd returns the NumMsd field if non-nil, zero value otherwise.

### GetNumMsdOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumMsdOk() (*int64, bool)`

GetNumMsdOk returns a tuple with the NumMsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMsd

`func (o *NiatelemetryNiaInventoryDcnm) SetNumMsd(v int64)`

SetNumMsd sets NumMsd field to given value.

### HasNumMsd

`func (o *NiatelemetryNiaInventoryDcnm) HasNumMsd() bool`

HasNumMsd returns a boolean if a field has been set.

### GetNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnm) GetNumSviVrfCount() int64`

GetNumSviVrfCount returns the NumSviVrfCount field if non-nil, zero value otherwise.

### GetNumSviVrfCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumSviVrfCountOk() (*int64, bool)`

GetNumSviVrfCountOk returns a tuple with the NumSviVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnm) SetNumSviVrfCount(v int64)`

SetNumSviVrfCount sets NumSviVrfCount field to given value.

### HasNumSviVrfCount

`func (o *NiatelemetryNiaInventoryDcnm) HasNumSviVrfCount() bool`

HasNumSviVrfCount returns a boolean if a field has been set.

### GetNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnm) GetNumTrmEnabledCount() int64`

GetNumTrmEnabledCount returns the NumTrmEnabledCount field if non-nil, zero value otherwise.

### GetNumTrmEnabledCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumTrmEnabledCountOk() (*int64, bool)`

GetNumTrmEnabledCountOk returns a tuple with the NumTrmEnabledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnm) SetNumTrmEnabledCount(v int64)`

SetNumTrmEnabledCount sets NumTrmEnabledCount field to given value.

### HasNumTrmEnabledCount

`func (o *NiatelemetryNiaInventoryDcnm) HasNumTrmEnabledCount() bool`

HasNumTrmEnabledCount returns a boolean if a field has been set.

### GetNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnm) GetNumUpgUsers() int64`

GetNumUpgUsers returns the NumUpgUsers field if non-nil, zero value otherwise.

### GetNumUpgUsersOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNumUpgUsersOk() (*int64, bool)`

GetNumUpgUsersOk returns a tuple with the NumUpgUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnm) SetNumUpgUsers(v int64)`

SetNumUpgUsers sets NumUpgUsers field to given value.

### HasNumUpgUsers

`func (o *NiatelemetryNiaInventoryDcnm) HasNumUpgUsers() bool`

HasNumUpgUsers returns a boolean if a field has been set.

### GetNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnm) GetNxosImageCount() int64`

GetNxosImageCount returns the NxosImageCount field if non-nil, zero value otherwise.

### GetNxosImageCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetNxosImageCountOk() (*int64, bool)`

GetNxosImageCountOk returns a tuple with the NxosImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnm) SetNxosImageCount(v int64)`

SetNxosImageCount sets NxosImageCount field to given value.

### HasNxosImageCount

`func (o *NiatelemetryNiaInventoryDcnm) HasNxosImageCount() bool`

HasNxosImageCount returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryDcnm) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryDcnm) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryDcnm) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryDcnm) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryNiaInventoryDcnm) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryNiaInventoryDcnm) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryNiaInventoryDcnm) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryNiaInventoryDcnm) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnm) GetUnderlayPeeringActiveLinksCount() int64`

GetUnderlayPeeringActiveLinksCount returns the UnderlayPeeringActiveLinksCount field if non-nil, zero value otherwise.

### GetUnderlayPeeringActiveLinksCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetUnderlayPeeringActiveLinksCountOk() (*int64, bool)`

GetUnderlayPeeringActiveLinksCountOk returns a tuple with the UnderlayPeeringActiveLinksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnm) SetUnderlayPeeringActiveLinksCount(v int64)`

SetUnderlayPeeringActiveLinksCount sets UnderlayPeeringActiveLinksCount field to given value.

### HasUnderlayPeeringActiveLinksCount

`func (o *NiatelemetryNiaInventoryDcnm) HasUnderlayPeeringActiveLinksCount() bool`

HasUnderlayPeeringActiveLinksCount returns a boolean if a field has been set.

### GetUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnm) GetUpgJobCount() int64`

GetUpgJobCount returns the UpgJobCount field if non-nil, zero value otherwise.

### GetUpgJobCountOk

`func (o *NiatelemetryNiaInventoryDcnm) GetUpgJobCountOk() (*int64, bool)`

GetUpgJobCountOk returns a tuple with the UpgJobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnm) SetUpgJobCount(v int64)`

SetUpgJobCount sets UpgJobCount field to given value.

### HasUpgJobCount

`func (o *NiatelemetryNiaInventoryDcnm) HasUpgJobCount() bool`

HasUpgJobCount returns a boolean if a field has been set.

### GetUpgStatus

`func (o *NiatelemetryNiaInventoryDcnm) GetUpgStatus() string`

GetUpgStatus returns the UpgStatus field if non-nil, zero value otherwise.

### GetUpgStatusOk

`func (o *NiatelemetryNiaInventoryDcnm) GetUpgStatusOk() (*string, bool)`

GetUpgStatusOk returns a tuple with the UpgStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgStatus

`func (o *NiatelemetryNiaInventoryDcnm) SetUpgStatus(v string)`

SetUpgStatus sets UpgStatus field to given value.

### HasUpgStatus

`func (o *NiatelemetryNiaInventoryDcnm) HasUpgStatus() bool`

HasUpgStatus returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNiaInventoryDcnm) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNiaInventoryDcnm) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNiaInventoryDcnm) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNiaInventoryDcnm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryDcnm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryDcnm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


