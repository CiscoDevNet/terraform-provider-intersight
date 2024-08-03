# StorageHitachiSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiSnapshot"]
**ConcordanceRate** | Pointer to **int64** | Concordance rate for pairs. | [optional] [readonly] 
**IsConsistencyGroup** | Pointer to **bool** | Whether the pair was created in the consistency group mode (CTG mode). | [optional] [readonly] 
**IsMultistageable** | Pointer to **bool** | Whether the pair can be a multistage pair. | [optional] [readonly] 
**IsRedirectOnWrite** | Pointer to **bool** | Returns whether it is a Thin Image Advanced pair. | [optional] [readonly] 
**MuNumber** | Pointer to **int64** | MU number of the primary volume. | [optional] [readonly] 
**Name** | Pointer to **string** | Object ID of the pair for snapshot data. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of the primary volume. | [optional] [readonly] 
**SnapshotGroupName** | Pointer to **string** | Name of the snapshot group that contains the pair for snapshot data. | [optional] [readonly] 
**SnapshotPoolId** | Pointer to **int64** | ID of the pool in which the snapshot data is created. | [optional] [readonly] 
**SplitTime** | Pointer to **string** | Time when snapshot data was created. | [optional] [readonly] 
**Status** | Pointer to **string** | Pair status. Pair status changes according to the pair operation. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of the secondary volume. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiSnapshot

`func NewStorageHitachiSnapshot(classId string, objectType string, ) *StorageHitachiSnapshot`

NewStorageHitachiSnapshot instantiates a new StorageHitachiSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiSnapshotWithDefaults

`func NewStorageHitachiSnapshotWithDefaults() *StorageHitachiSnapshot`

NewStorageHitachiSnapshotWithDefaults instantiates a new StorageHitachiSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiSnapshot) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiSnapshot) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiSnapshot) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiSnapshot) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiSnapshot) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiSnapshot) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConcordanceRate

`func (o *StorageHitachiSnapshot) GetConcordanceRate() int64`

GetConcordanceRate returns the ConcordanceRate field if non-nil, zero value otherwise.

### GetConcordanceRateOk

`func (o *StorageHitachiSnapshot) GetConcordanceRateOk() (*int64, bool)`

GetConcordanceRateOk returns a tuple with the ConcordanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcordanceRate

`func (o *StorageHitachiSnapshot) SetConcordanceRate(v int64)`

SetConcordanceRate sets ConcordanceRate field to given value.

### HasConcordanceRate

`func (o *StorageHitachiSnapshot) HasConcordanceRate() bool`

HasConcordanceRate returns a boolean if a field has been set.

### GetIsConsistencyGroup

`func (o *StorageHitachiSnapshot) GetIsConsistencyGroup() bool`

GetIsConsistencyGroup returns the IsConsistencyGroup field if non-nil, zero value otherwise.

### GetIsConsistencyGroupOk

`func (o *StorageHitachiSnapshot) GetIsConsistencyGroupOk() (*bool, bool)`

GetIsConsistencyGroupOk returns a tuple with the IsConsistencyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConsistencyGroup

`func (o *StorageHitachiSnapshot) SetIsConsistencyGroup(v bool)`

SetIsConsistencyGroup sets IsConsistencyGroup field to given value.

### HasIsConsistencyGroup

`func (o *StorageHitachiSnapshot) HasIsConsistencyGroup() bool`

HasIsConsistencyGroup returns a boolean if a field has been set.

### GetIsMultistageable

`func (o *StorageHitachiSnapshot) GetIsMultistageable() bool`

GetIsMultistageable returns the IsMultistageable field if non-nil, zero value otherwise.

### GetIsMultistageableOk

`func (o *StorageHitachiSnapshot) GetIsMultistageableOk() (*bool, bool)`

GetIsMultistageableOk returns a tuple with the IsMultistageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultistageable

`func (o *StorageHitachiSnapshot) SetIsMultistageable(v bool)`

SetIsMultistageable sets IsMultistageable field to given value.

### HasIsMultistageable

`func (o *StorageHitachiSnapshot) HasIsMultistageable() bool`

HasIsMultistageable returns a boolean if a field has been set.

### GetIsRedirectOnWrite

`func (o *StorageHitachiSnapshot) GetIsRedirectOnWrite() bool`

GetIsRedirectOnWrite returns the IsRedirectOnWrite field if non-nil, zero value otherwise.

### GetIsRedirectOnWriteOk

`func (o *StorageHitachiSnapshot) GetIsRedirectOnWriteOk() (*bool, bool)`

GetIsRedirectOnWriteOk returns a tuple with the IsRedirectOnWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedirectOnWrite

`func (o *StorageHitachiSnapshot) SetIsRedirectOnWrite(v bool)`

SetIsRedirectOnWrite sets IsRedirectOnWrite field to given value.

### HasIsRedirectOnWrite

`func (o *StorageHitachiSnapshot) HasIsRedirectOnWrite() bool`

HasIsRedirectOnWrite returns a boolean if a field has been set.

### GetMuNumber

`func (o *StorageHitachiSnapshot) GetMuNumber() int64`

GetMuNumber returns the MuNumber field if non-nil, zero value otherwise.

### GetMuNumberOk

`func (o *StorageHitachiSnapshot) GetMuNumberOk() (*int64, bool)`

GetMuNumberOk returns a tuple with the MuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuNumber

`func (o *StorageHitachiSnapshot) SetMuNumber(v int64)`

SetMuNumber sets MuNumber field to given value.

### HasMuNumber

`func (o *StorageHitachiSnapshot) HasMuNumber() bool`

HasMuNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiSnapshot) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiSnapshot) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiSnapshot) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiSnapshot) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetSnapshotGroupName

`func (o *StorageHitachiSnapshot) GetSnapshotGroupName() string`

GetSnapshotGroupName returns the SnapshotGroupName field if non-nil, zero value otherwise.

### GetSnapshotGroupNameOk

`func (o *StorageHitachiSnapshot) GetSnapshotGroupNameOk() (*string, bool)`

GetSnapshotGroupNameOk returns a tuple with the SnapshotGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotGroupName

`func (o *StorageHitachiSnapshot) SetSnapshotGroupName(v string)`

SetSnapshotGroupName sets SnapshotGroupName field to given value.

### HasSnapshotGroupName

`func (o *StorageHitachiSnapshot) HasSnapshotGroupName() bool`

HasSnapshotGroupName returns a boolean if a field has been set.

### GetSnapshotPoolId

`func (o *StorageHitachiSnapshot) GetSnapshotPoolId() int64`

GetSnapshotPoolId returns the SnapshotPoolId field if non-nil, zero value otherwise.

### GetSnapshotPoolIdOk

`func (o *StorageHitachiSnapshot) GetSnapshotPoolIdOk() (*int64, bool)`

GetSnapshotPoolIdOk returns a tuple with the SnapshotPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPoolId

`func (o *StorageHitachiSnapshot) SetSnapshotPoolId(v int64)`

SetSnapshotPoolId sets SnapshotPoolId field to given value.

### HasSnapshotPoolId

`func (o *StorageHitachiSnapshot) HasSnapshotPoolId() bool`

HasSnapshotPoolId returns a boolean if a field has been set.

### GetSplitTime

`func (o *StorageHitachiSnapshot) GetSplitTime() string`

GetSplitTime returns the SplitTime field if non-nil, zero value otherwise.

### GetSplitTimeOk

`func (o *StorageHitachiSnapshot) GetSplitTimeOk() (*string, bool)`

GetSplitTimeOk returns a tuple with the SplitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitTime

`func (o *StorageHitachiSnapshot) SetSplitTime(v string)`

SetSplitTime sets SplitTime field to given value.

### HasSplitTime

`func (o *StorageHitachiSnapshot) HasSplitTime() bool`

HasSplitTime returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiSnapshot) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiSnapshot) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiSnapshot) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiSnapshot) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiSnapshot) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiSnapshot) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiSnapshot) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiSnapshot) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageHitachiSnapshot) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageHitachiSnapshot) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StorageHitachiSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiSnapshot) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StorageHitachiSnapshot) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StorageHitachiSnapshot) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


