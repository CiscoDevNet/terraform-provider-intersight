# StorageStorageUtilizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataReduction** | Pointer to **float32** | Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Snapshot** | Pointer to **int64** | Physical space occupied by the snapshots, represented in bytes. | [optional] [readonly] 
**ThinProvisioned** | Pointer to **float32** | Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Volume** | Pointer to **int64** | Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size id represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageStorageUtilizationAllOf

`func NewStorageStorageUtilizationAllOf() *StorageStorageUtilizationAllOf`

NewStorageStorageUtilizationAllOf instantiates a new StorageStorageUtilizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStorageUtilizationAllOfWithDefaults

`func NewStorageStorageUtilizationAllOfWithDefaults() *StorageStorageUtilizationAllOf`

NewStorageStorageUtilizationAllOfWithDefaults instantiates a new StorageStorageUtilizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataReduction

`func (o *StorageStorageUtilizationAllOf) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StorageStorageUtilizationAllOf) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StorageStorageUtilizationAllOf) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StorageStorageUtilizationAllOf) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetSnapshot

`func (o *StorageStorageUtilizationAllOf) GetSnapshot() int64`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *StorageStorageUtilizationAllOf) GetSnapshotOk() (*int64, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *StorageStorageUtilizationAllOf) SetSnapshot(v int64)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *StorageStorageUtilizationAllOf) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *StorageStorageUtilizationAllOf) GetThinProvisioned() float32`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *StorageStorageUtilizationAllOf) GetThinProvisionedOk() (*float32, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *StorageStorageUtilizationAllOf) SetThinProvisioned(v float32)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *StorageStorageUtilizationAllOf) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StorageStorageUtilizationAllOf) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StorageStorageUtilizationAllOf) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StorageStorageUtilizationAllOf) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StorageStorageUtilizationAllOf) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetVolume

`func (o *StorageStorageUtilizationAllOf) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageStorageUtilizationAllOf) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageStorageUtilizationAllOf) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageStorageUtilizationAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


