# StorageStorageUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**DataReduction** | Pointer to **float32** | Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Snapshot** | Pointer to **int64** | Physical space occupied by the snapshots, represented in bytes. | [optional] [readonly] 
**ThinProvisioned** | Pointer to **float32** | Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Volume** | Pointer to **int64** | Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size id represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageStorageUtilization

`func NewStorageStorageUtilization(classId string, objectType string, ) *StorageStorageUtilization`

NewStorageStorageUtilization instantiates a new StorageStorageUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStorageUtilizationWithDefaults

`func NewStorageStorageUtilizationWithDefaults() *StorageStorageUtilization`

NewStorageStorageUtilizationWithDefaults instantiates a new StorageStorageUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageStorageUtilization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageStorageUtilization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageStorageUtilization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageStorageUtilization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageStorageUtilization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageStorageUtilization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReduction

`func (o *StorageStorageUtilization) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StorageStorageUtilization) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StorageStorageUtilization) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StorageStorageUtilization) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetSnapshot

`func (o *StorageStorageUtilization) GetSnapshot() int64`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *StorageStorageUtilization) GetSnapshotOk() (*int64, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *StorageStorageUtilization) SetSnapshot(v int64)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *StorageStorageUtilization) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *StorageStorageUtilization) GetThinProvisioned() float32`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *StorageStorageUtilization) GetThinProvisionedOk() (*float32, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *StorageStorageUtilization) SetThinProvisioned(v float32)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *StorageStorageUtilization) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StorageStorageUtilization) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StorageStorageUtilization) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StorageStorageUtilization) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StorageStorageUtilization) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetVolume

`func (o *StorageStorageUtilization) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageStorageUtilization) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageStorageUtilization) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageStorageUtilization) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


