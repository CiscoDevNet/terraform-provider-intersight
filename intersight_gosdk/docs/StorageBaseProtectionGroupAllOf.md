# StorageBaseProtectionGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the protection Group. | [optional] 
**Prefix** | Pointer to **string** | Prefix used for all generated snapshots from the protection group. | [optional] 
**ReplicationEnabled** | Pointer to **bool** | Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. | [optional] 
**SnapshotEnabled** | Pointer to **bool** | Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. | [optional] 

## Methods

### NewStorageBaseProtectionGroupAllOf

`func NewStorageBaseProtectionGroupAllOf() *StorageBaseProtectionGroupAllOf`

NewStorageBaseProtectionGroupAllOf instantiates a new StorageBaseProtectionGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseProtectionGroupAllOfWithDefaults

`func NewStorageBaseProtectionGroupAllOfWithDefaults() *StorageBaseProtectionGroupAllOf`

NewStorageBaseProtectionGroupAllOfWithDefaults instantiates a new StorageBaseProtectionGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


