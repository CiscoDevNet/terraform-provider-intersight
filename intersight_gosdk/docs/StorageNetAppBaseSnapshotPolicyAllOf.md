# StorageNetAppBaseSnapshotPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Copies** | Pointer to [**[]StorageNetAppSnapshotPolicySchedule**](StorageNetAppSnapshotPolicySchedule.md) |  | [optional] 
**Enabled** | Pointer to **string** | Snapshot policy is enabled or not. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the NetApp Snapshot Policy. | [optional] [readonly] 
**Scope** | Pointer to **string** | Identifies whether the Snapshot Policy is owned by the Storage Virtual Machine or the cluster. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the NetApp Snapshot Policy. | [optional] [readonly] 

## Methods

### NewStorageNetAppBaseSnapshotPolicyAllOf

`func NewStorageNetAppBaseSnapshotPolicyAllOf(classId string, objectType string, ) *StorageNetAppBaseSnapshotPolicyAllOf`

NewStorageNetAppBaseSnapshotPolicyAllOf instantiates a new StorageNetAppBaseSnapshotPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppBaseSnapshotPolicyAllOfWithDefaults

`func NewStorageNetAppBaseSnapshotPolicyAllOfWithDefaults() *StorageNetAppBaseSnapshotPolicyAllOf`

NewStorageNetAppBaseSnapshotPolicyAllOfWithDefaults instantiates a new StorageNetAppBaseSnapshotPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCopies

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetCopies() []StorageNetAppSnapshotPolicySchedule`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetCopiesOk() (*[]StorageNetAppSnapshotPolicySchedule, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetCopies(v []StorageNetAppSnapshotPolicySchedule)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### SetCopiesNil

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetCopiesNil(b bool)`

 SetCopiesNil sets the value for Copies to be an explicit nil

### UnsetCopies
`func (o *StorageNetAppBaseSnapshotPolicyAllOf) UnsetCopies()`

UnsetCopies ensures that no value is present for Copies, not even an explicit nil
### GetEnabled

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppBaseSnapshotPolicyAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


