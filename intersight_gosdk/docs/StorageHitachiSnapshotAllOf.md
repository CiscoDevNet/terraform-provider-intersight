# StorageHitachiSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiSnapshot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiSnapshot"]
**ConcordanceRate** | Pointer to **int64** | Concordance rate for pairs. | [optional] [readonly] 
**IsConsistencyGroup** | Pointer to **bool** | Whether the pair was created in the consistency group mode (CTG mode). | [optional] [readonly] 
**IsMultistageable** | Pointer to **bool** | Whether the pair can be a multistage pair. | [optional] [readonly] 
**MuNumber** | Pointer to **int64** | MU number of the primary volume. | [optional] [readonly] 
**Name** | Pointer to **string** | Object ID of the pair for snapshot data. | [optional] [readonly] 
**PvolLdevId** | Pointer to **int64** | LDEV number of the primary volume. | [optional] [readonly] 
**SnapshotGroupName** | Pointer to **string** | Name of the snapshot group that contains the pair for snapshot data. | [optional] [readonly] 
**SnapshotPoolId** | Pointer to **int64** | ID of the pool in which the snapshot data is created. | [optional] [readonly] 
**SplitTime** | Pointer to **string** | Time when snapshot data was created. | [optional] [readonly] 
**Status** | Pointer to **string** | Pair status. Pair status changes according to the pair operation. | [optional] [readonly] 
**SvolLdevId** | Pointer to **int64** | LDEV number of the secondary volume. | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiSnapshotAllOf

`func NewStorageHitachiSnapshotAllOf(classId string, objectType string, ) *StorageHitachiSnapshotAllOf`

NewStorageHitachiSnapshotAllOf instantiates a new StorageHitachiSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiSnapshotAllOfWithDefaults

`func NewStorageHitachiSnapshotAllOfWithDefaults() *StorageHitachiSnapshotAllOf`

NewStorageHitachiSnapshotAllOfWithDefaults instantiates a new StorageHitachiSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiSnapshotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiSnapshotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiSnapshotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiSnapshotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiSnapshotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiSnapshotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConcordanceRate

`func (o *StorageHitachiSnapshotAllOf) GetConcordanceRate() int64`

GetConcordanceRate returns the ConcordanceRate field if non-nil, zero value otherwise.

### GetConcordanceRateOk

`func (o *StorageHitachiSnapshotAllOf) GetConcordanceRateOk() (*int64, bool)`

GetConcordanceRateOk returns a tuple with the ConcordanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcordanceRate

`func (o *StorageHitachiSnapshotAllOf) SetConcordanceRate(v int64)`

SetConcordanceRate sets ConcordanceRate field to given value.

### HasConcordanceRate

`func (o *StorageHitachiSnapshotAllOf) HasConcordanceRate() bool`

HasConcordanceRate returns a boolean if a field has been set.

### GetIsConsistencyGroup

`func (o *StorageHitachiSnapshotAllOf) GetIsConsistencyGroup() bool`

GetIsConsistencyGroup returns the IsConsistencyGroup field if non-nil, zero value otherwise.

### GetIsConsistencyGroupOk

`func (o *StorageHitachiSnapshotAllOf) GetIsConsistencyGroupOk() (*bool, bool)`

GetIsConsistencyGroupOk returns a tuple with the IsConsistencyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConsistencyGroup

`func (o *StorageHitachiSnapshotAllOf) SetIsConsistencyGroup(v bool)`

SetIsConsistencyGroup sets IsConsistencyGroup field to given value.

### HasIsConsistencyGroup

`func (o *StorageHitachiSnapshotAllOf) HasIsConsistencyGroup() bool`

HasIsConsistencyGroup returns a boolean if a field has been set.

### GetIsMultistageable

`func (o *StorageHitachiSnapshotAllOf) GetIsMultistageable() bool`

GetIsMultistageable returns the IsMultistageable field if non-nil, zero value otherwise.

### GetIsMultistageableOk

`func (o *StorageHitachiSnapshotAllOf) GetIsMultistageableOk() (*bool, bool)`

GetIsMultistageableOk returns a tuple with the IsMultistageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultistageable

`func (o *StorageHitachiSnapshotAllOf) SetIsMultistageable(v bool)`

SetIsMultistageable sets IsMultistageable field to given value.

### HasIsMultistageable

`func (o *StorageHitachiSnapshotAllOf) HasIsMultistageable() bool`

HasIsMultistageable returns a boolean if a field has been set.

### GetMuNumber

`func (o *StorageHitachiSnapshotAllOf) GetMuNumber() int64`

GetMuNumber returns the MuNumber field if non-nil, zero value otherwise.

### GetMuNumberOk

`func (o *StorageHitachiSnapshotAllOf) GetMuNumberOk() (*int64, bool)`

GetMuNumberOk returns a tuple with the MuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuNumber

`func (o *StorageHitachiSnapshotAllOf) SetMuNumber(v int64)`

SetMuNumber sets MuNumber field to given value.

### HasMuNumber

`func (o *StorageHitachiSnapshotAllOf) HasMuNumber() bool`

HasMuNumber returns a boolean if a field has been set.

### GetName

`func (o *StorageHitachiSnapshotAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageHitachiSnapshotAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageHitachiSnapshotAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageHitachiSnapshotAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPvolLdevId

`func (o *StorageHitachiSnapshotAllOf) GetPvolLdevId() int64`

GetPvolLdevId returns the PvolLdevId field if non-nil, zero value otherwise.

### GetPvolLdevIdOk

`func (o *StorageHitachiSnapshotAllOf) GetPvolLdevIdOk() (*int64, bool)`

GetPvolLdevIdOk returns a tuple with the PvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvolLdevId

`func (o *StorageHitachiSnapshotAllOf) SetPvolLdevId(v int64)`

SetPvolLdevId sets PvolLdevId field to given value.

### HasPvolLdevId

`func (o *StorageHitachiSnapshotAllOf) HasPvolLdevId() bool`

HasPvolLdevId returns a boolean if a field has been set.

### GetSnapshotGroupName

`func (o *StorageHitachiSnapshotAllOf) GetSnapshotGroupName() string`

GetSnapshotGroupName returns the SnapshotGroupName field if non-nil, zero value otherwise.

### GetSnapshotGroupNameOk

`func (o *StorageHitachiSnapshotAllOf) GetSnapshotGroupNameOk() (*string, bool)`

GetSnapshotGroupNameOk returns a tuple with the SnapshotGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotGroupName

`func (o *StorageHitachiSnapshotAllOf) SetSnapshotGroupName(v string)`

SetSnapshotGroupName sets SnapshotGroupName field to given value.

### HasSnapshotGroupName

`func (o *StorageHitachiSnapshotAllOf) HasSnapshotGroupName() bool`

HasSnapshotGroupName returns a boolean if a field has been set.

### GetSnapshotPoolId

`func (o *StorageHitachiSnapshotAllOf) GetSnapshotPoolId() int64`

GetSnapshotPoolId returns the SnapshotPoolId field if non-nil, zero value otherwise.

### GetSnapshotPoolIdOk

`func (o *StorageHitachiSnapshotAllOf) GetSnapshotPoolIdOk() (*int64, bool)`

GetSnapshotPoolIdOk returns a tuple with the SnapshotPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPoolId

`func (o *StorageHitachiSnapshotAllOf) SetSnapshotPoolId(v int64)`

SetSnapshotPoolId sets SnapshotPoolId field to given value.

### HasSnapshotPoolId

`func (o *StorageHitachiSnapshotAllOf) HasSnapshotPoolId() bool`

HasSnapshotPoolId returns a boolean if a field has been set.

### GetSplitTime

`func (o *StorageHitachiSnapshotAllOf) GetSplitTime() string`

GetSplitTime returns the SplitTime field if non-nil, zero value otherwise.

### GetSplitTimeOk

`func (o *StorageHitachiSnapshotAllOf) GetSplitTimeOk() (*string, bool)`

GetSplitTimeOk returns a tuple with the SplitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitTime

`func (o *StorageHitachiSnapshotAllOf) SetSplitTime(v string)`

SetSplitTime sets SplitTime field to given value.

### HasSplitTime

`func (o *StorageHitachiSnapshotAllOf) HasSplitTime() bool`

HasSplitTime returns a boolean if a field has been set.

### GetStatus

`func (o *StorageHitachiSnapshotAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageHitachiSnapshotAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageHitachiSnapshotAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageHitachiSnapshotAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSvolLdevId

`func (o *StorageHitachiSnapshotAllOf) GetSvolLdevId() int64`

GetSvolLdevId returns the SvolLdevId field if non-nil, zero value otherwise.

### GetSvolLdevIdOk

`func (o *StorageHitachiSnapshotAllOf) GetSvolLdevIdOk() (*int64, bool)`

GetSvolLdevIdOk returns a tuple with the SvolLdevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvolLdevId

`func (o *StorageHitachiSnapshotAllOf) SetSvolLdevId(v int64)`

SetSvolLdevId sets SvolLdevId field to given value.

### HasSvolLdevId

`func (o *StorageHitachiSnapshotAllOf) HasSvolLdevId() bool`

HasSvolLdevId returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiSnapshotAllOf) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiSnapshotAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiSnapshotAllOf) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiSnapshotAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiSnapshotAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiSnapshotAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiSnapshotAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiSnapshotAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


