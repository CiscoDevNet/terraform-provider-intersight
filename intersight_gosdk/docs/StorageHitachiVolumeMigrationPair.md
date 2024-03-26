# StorageHitachiVolumeMigrationPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiVolumeMigrationPair"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiVolumeMigrationPair"]
**CopyMode** | Pointer to **string** | Copy mode. NotSynchronized or VolumeMigration is stored. | [optional] [readonly] 
**LocalCloneCopypairId** | Pointer to **string** | Object ID of the pair. The following informations of pair are output, separated by commas. &lt;copy group name&gt;, &lt;device group name for the P-VOL (source volume)&gt;, &lt;device group name for the S-VOL (target volume)&gt;, &lt;name of the pair&gt;. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of the P-VOL (source volume) with a decimal (base 10) number. | [optional] [readonly] 
**PvolStatus** | Pointer to **string** | Pair volume status of the P-VOL. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of the S-VOL (target volume) with a decimal (base 10) number. | [optional] [readonly] 
**SvolStatus** | Pointer to **string** | Pair volume status of the S-VOL. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiVolumeMigrationPair

`func NewStorageHitachiVolumeMigrationPair(classId string, objectType string, ) *StorageHitachiVolumeMigrationPair`

NewStorageHitachiVolumeMigrationPair instantiates a new StorageHitachiVolumeMigrationPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiVolumeMigrationPairWithDefaults

`func NewStorageHitachiVolumeMigrationPairWithDefaults() *StorageHitachiVolumeMigrationPair`

NewStorageHitachiVolumeMigrationPairWithDefaults instantiates a new StorageHitachiVolumeMigrationPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiVolumeMigrationPair) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiVolumeMigrationPair) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiVolumeMigrationPair) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiVolumeMigrationPair) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiVolumeMigrationPair) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiVolumeMigrationPair) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCopyMode

`func (o *StorageHitachiVolumeMigrationPair) GetCopyMode() string`

GetCopyMode returns the CopyMode field if non-nil, zero value otherwise.

### GetCopyModeOk

`func (o *StorageHitachiVolumeMigrationPair) GetCopyModeOk() (*string, bool)`

GetCopyModeOk returns a tuple with the CopyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMode

`func (o *StorageHitachiVolumeMigrationPair) SetCopyMode(v string)`

SetCopyMode sets CopyMode field to given value.

### HasCopyMode

`func (o *StorageHitachiVolumeMigrationPair) HasCopyMode() bool`

HasCopyMode returns a boolean if a field has been set.

### GetLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPair) GetLocalCloneCopypairId() string`

GetLocalCloneCopypairId returns the LocalCloneCopypairId field if non-nil, zero value otherwise.

### GetLocalCloneCopypairIdOk

`func (o *StorageHitachiVolumeMigrationPair) GetLocalCloneCopypairIdOk() (*string, bool)`

GetLocalCloneCopypairIdOk returns a tuple with the LocalCloneCopypairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPair) SetLocalCloneCopypairId(v string)`

SetLocalCloneCopypairId sets LocalCloneCopypairId field to given value.

### HasLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPair) HasLocalCloneCopypairId() bool`

HasLocalCloneCopypairId returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiVolumeMigrationPair) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetPvolStatus

`func (o *StorageHitachiVolumeMigrationPair) GetPvolStatus() string`

GetPvolStatus returns the PvolStatus field if non-nil, zero value otherwise.

### GetPvolStatusOk

`func (o *StorageHitachiVolumeMigrationPair) GetPvolStatusOk() (*string, bool)`

GetPvolStatusOk returns a tuple with the PvolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolStatus

`func (o *StorageHitachiVolumeMigrationPair) SetPvolStatus(v string)`

SetPvolStatus sets PvolStatus field to given value.

### HasPvolStatus

`func (o *StorageHitachiVolumeMigrationPair) HasPvolStatus() bool`

HasPvolStatus returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiVolumeMigrationPair) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiVolumeMigrationPair) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetSvolStatus

`func (o *StorageHitachiVolumeMigrationPair) GetSvolStatus() string`

GetSvolStatus returns the SvolStatus field if non-nil, zero value otherwise.

### GetSvolStatusOk

`func (o *StorageHitachiVolumeMigrationPair) GetSvolStatusOk() (*string, bool)`

GetSvolStatusOk returns a tuple with the SvolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolStatus

`func (o *StorageHitachiVolumeMigrationPair) SetSvolStatus(v string)`

SetSvolStatus sets SvolStatus field to given value.

### HasSvolStatus

`func (o *StorageHitachiVolumeMigrationPair) HasSvolStatus() bool`

HasSvolStatus returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiVolumeMigrationPair) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiVolumeMigrationPair) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiVolumeMigrationPair) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiVolumeMigrationPair) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPair) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiVolumeMigrationPair) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPair) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPair) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


