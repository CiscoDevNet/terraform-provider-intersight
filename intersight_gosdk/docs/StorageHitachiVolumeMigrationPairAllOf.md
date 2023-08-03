# StorageHitachiVolumeMigrationPairAllOf

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

### NewStorageHitachiVolumeMigrationPairAllOf

`func NewStorageHitachiVolumeMigrationPairAllOf(classId string, objectType string, ) *StorageHitachiVolumeMigrationPairAllOf`

NewStorageHitachiVolumeMigrationPairAllOf instantiates a new StorageHitachiVolumeMigrationPairAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiVolumeMigrationPairAllOfWithDefaults

`func NewStorageHitachiVolumeMigrationPairAllOfWithDefaults() *StorageHitachiVolumeMigrationPairAllOf`

NewStorageHitachiVolumeMigrationPairAllOfWithDefaults instantiates a new StorageHitachiVolumeMigrationPairAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCopyMode

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetCopyMode() string`

GetCopyMode returns the CopyMode field if non-nil, zero value otherwise.

### GetCopyModeOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetCopyModeOk() (*string, bool)`

GetCopyModeOk returns a tuple with the CopyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyMode

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetCopyMode(v string)`

SetCopyMode sets CopyMode field to given value.

### HasCopyMode

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasCopyMode() bool`

HasCopyMode returns a boolean if a field has been set.

### GetLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetLocalCloneCopypairId() string`

GetLocalCloneCopypairId returns the LocalCloneCopypairId field if non-nil, zero value otherwise.

### GetLocalCloneCopypairIdOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetLocalCloneCopypairIdOk() (*string, bool)`

GetLocalCloneCopypairIdOk returns a tuple with the LocalCloneCopypairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetLocalCloneCopypairId(v string)`

SetLocalCloneCopypairId sets LocalCloneCopypairId field to given value.

### HasLocalCloneCopypairId

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasLocalCloneCopypairId() bool`

HasLocalCloneCopypairId returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetPvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetPvolStatus() string`

GetPvolStatus returns the PvolStatus field if non-nil, zero value otherwise.

### GetPvolStatusOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetPvolStatusOk() (*string, bool)`

GetPvolStatusOk returns a tuple with the PvolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetPvolStatus(v string)`

SetPvolStatus sets PvolStatus field to given value.

### HasPvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasPvolStatus() bool`

HasPvolStatus returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetSvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetSvolStatus() string`

GetSvolStatus returns the SvolStatus field if non-nil, zero value otherwise.

### GetSvolStatusOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetSvolStatusOk() (*string, bool)`

GetSvolStatusOk returns a tuple with the SvolStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetSvolStatus(v string)`

SetSvolStatus sets SvolStatus field to given value.

### HasSvolStatus

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasSvolStatus() bool`

HasSvolStatus returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiVolumeMigrationPairAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPairAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiVolumeMigrationPairAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


