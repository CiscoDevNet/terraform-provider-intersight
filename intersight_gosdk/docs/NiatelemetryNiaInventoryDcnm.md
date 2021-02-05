# NiatelemetryNiaInventoryDcnm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryDcnm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryDcnm"]
**Dev** | Pointer to **bool** | Returns the value of the dev Field. | [optional] 
**HaEnabled** | Pointer to **bool** | Returns the value of the haEnabled field. | [optional] 
**HaReplicationStatus** | Pointer to **string** | Returns the value of the haReplicationStatus field. | [optional] 
**Install** | Pointer to **string** | Returns the value of the install field. | [optional] 
**IsMediaController** | Pointer to **bool** | Returns the value of the isMediaController field. | [optional] 
**NumFabrics** | Pointer to **int64** | Returns total number of fabrics in DCNM set-up. | [optional] 
**NumFabricsInMsd** | Pointer to **int64** | Returns the number of fabrics in msd. | [optional] 
**NumLocalUsers** | Pointer to **int64** | Returns the number of local users other than admin user. | [optional] 
**Version** | Pointer to **string** | Returns the value of the version field. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

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


