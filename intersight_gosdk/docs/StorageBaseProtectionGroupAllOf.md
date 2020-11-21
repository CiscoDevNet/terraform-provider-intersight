# StorageBaseProtectionGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureProtectionGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.PureProtectionGroup"]
**Name** | Pointer to **string** | Name of the protection Group. | [optional] 
**Prefix** | Pointer to **string** | Prefix used for all generated snapshots from the protection group. | [optional] 
**ReplicationEnabled** | Pointer to **bool** | Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. | [optional] 
**SnapshotEnabled** | Pointer to **bool** | Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. | [optional] 

## Methods

### NewStorageBaseProtectionGroupAllOf

`func NewStorageBaseProtectionGroupAllOf(classId string, objectType string, ) *StorageBaseProtectionGroupAllOf`

NewStorageBaseProtectionGroupAllOf instantiates a new StorageBaseProtectionGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseProtectionGroupAllOfWithDefaults

`func NewStorageBaseProtectionGroupAllOfWithDefaults() *StorageBaseProtectionGroupAllOf`

NewStorageBaseProtectionGroupAllOfWithDefaults instantiates a new StorageBaseProtectionGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageBaseProtectionGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageBaseProtectionGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageBaseProtectionGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageBaseProtectionGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageBaseProtectionGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageBaseProtectionGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *StorageBaseProtectionGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseProtectionGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseProtectionGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseProtectionGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *StorageBaseProtectionGroupAllOf) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *StorageBaseProtectionGroupAllOf) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *StorageBaseProtectionGroupAllOf) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *StorageBaseProtectionGroupAllOf) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReplicationEnabled

`func (o *StorageBaseProtectionGroupAllOf) GetReplicationEnabled() bool`

GetReplicationEnabled returns the ReplicationEnabled field if non-nil, zero value otherwise.

### GetReplicationEnabledOk

`func (o *StorageBaseProtectionGroupAllOf) GetReplicationEnabledOk() (*bool, bool)`

GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabled

`func (o *StorageBaseProtectionGroupAllOf) SetReplicationEnabled(v bool)`

SetReplicationEnabled sets ReplicationEnabled field to given value.

### HasReplicationEnabled

`func (o *StorageBaseProtectionGroupAllOf) HasReplicationEnabled() bool`

HasReplicationEnabled returns a boolean if a field has been set.

### GetSnapshotEnabled

`func (o *StorageBaseProtectionGroupAllOf) GetSnapshotEnabled() bool`

GetSnapshotEnabled returns the SnapshotEnabled field if non-nil, zero value otherwise.

### GetSnapshotEnabledOk

`func (o *StorageBaseProtectionGroupAllOf) GetSnapshotEnabledOk() (*bool, bool)`

GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnabled

`func (o *StorageBaseProtectionGroupAllOf) SetSnapshotEnabled(v bool)`

SetSnapshotEnabled sets SnapshotEnabled field to given value.

### HasSnapshotEnabled

`func (o *StorageBaseProtectionGroupAllOf) HasSnapshotEnabled() bool`

HasSnapshotEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


