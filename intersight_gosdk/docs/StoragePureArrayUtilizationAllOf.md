# StoragePureArrayUtilizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureArrayUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureArrayUtilization"]
**DataReduction** | Pointer to **float32** | Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Parity** | Pointer to **float32** | Percentage of data that is fully protected. The percentage value will drop below 100% if the data is not fully protected. | [optional] [readonly] 
**Provisioned** | Pointer to **int64** | Total provisioned storage capacity in Pure FlashArray, represented in bytes. | [optional] [readonly] 
**Shared** | Pointer to **int64** | Physical space occupied by deduplicated data, represented in bytes. The space is shared with other volumes and snapshots as a result of data deduplication. | [optional] [readonly] 
**Snapshot** | Pointer to **int64** | Physical space occupied by the snapshots, represented in bytes. | [optional] [readonly] 
**System** | Pointer to **int64** | Physical space occupied by internal array metadata, represented in bytes. | [optional] [readonly] 
**ThinProvisioned** | Pointer to **float32** | Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array&#39;s flash modules. | [optional] [readonly] 
**Volume** | Pointer to **int64** | Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size is represented in bytes. | [optional] [readonly] 

## Methods

### NewStoragePureArrayUtilizationAllOf

`func NewStoragePureArrayUtilizationAllOf(classId string, objectType string, ) *StoragePureArrayUtilizationAllOf`

NewStoragePureArrayUtilizationAllOf instantiates a new StoragePureArrayUtilizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureArrayUtilizationAllOfWithDefaults

`func NewStoragePureArrayUtilizationAllOfWithDefaults() *StoragePureArrayUtilizationAllOf`

NewStoragePureArrayUtilizationAllOfWithDefaults instantiates a new StoragePureArrayUtilizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureArrayUtilizationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureArrayUtilizationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureArrayUtilizationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureArrayUtilizationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureArrayUtilizationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureArrayUtilizationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReduction

`func (o *StoragePureArrayUtilizationAllOf) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StoragePureArrayUtilizationAllOf) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StoragePureArrayUtilizationAllOf) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StoragePureArrayUtilizationAllOf) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetParity

`func (o *StoragePureArrayUtilizationAllOf) GetParity() float32`

GetParity returns the Parity field if non-nil, zero value otherwise.

### GetParityOk

`func (o *StoragePureArrayUtilizationAllOf) GetParityOk() (*float32, bool)`

GetParityOk returns a tuple with the Parity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParity

`func (o *StoragePureArrayUtilizationAllOf) SetParity(v float32)`

SetParity sets Parity field to given value.

### HasParity

`func (o *StoragePureArrayUtilizationAllOf) HasParity() bool`

HasParity returns a boolean if a field has been set.

### GetProvisioned

`func (o *StoragePureArrayUtilizationAllOf) GetProvisioned() int64`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *StoragePureArrayUtilizationAllOf) GetProvisionedOk() (*int64, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *StoragePureArrayUtilizationAllOf) SetProvisioned(v int64)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *StoragePureArrayUtilizationAllOf) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetShared

`func (o *StoragePureArrayUtilizationAllOf) GetShared() int64`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StoragePureArrayUtilizationAllOf) GetSharedOk() (*int64, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StoragePureArrayUtilizationAllOf) SetShared(v int64)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StoragePureArrayUtilizationAllOf) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSnapshot

`func (o *StoragePureArrayUtilizationAllOf) GetSnapshot() int64`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *StoragePureArrayUtilizationAllOf) GetSnapshotOk() (*int64, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *StoragePureArrayUtilizationAllOf) SetSnapshot(v int64)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *StoragePureArrayUtilizationAllOf) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSystem

`func (o *StoragePureArrayUtilizationAllOf) GetSystem() int64`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *StoragePureArrayUtilizationAllOf) GetSystemOk() (*int64, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *StoragePureArrayUtilizationAllOf) SetSystem(v int64)`

SetSystem sets System field to given value.

### HasSystem

`func (o *StoragePureArrayUtilizationAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetThinProvisioned

`func (o *StoragePureArrayUtilizationAllOf) GetThinProvisioned() float32`

GetThinProvisioned returns the ThinProvisioned field if non-nil, zero value otherwise.

### GetThinProvisionedOk

`func (o *StoragePureArrayUtilizationAllOf) GetThinProvisionedOk() (*float32, bool)`

GetThinProvisionedOk returns a tuple with the ThinProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioned

`func (o *StoragePureArrayUtilizationAllOf) SetThinProvisioned(v float32)`

SetThinProvisioned sets ThinProvisioned field to given value.

### HasThinProvisioned

`func (o *StoragePureArrayUtilizationAllOf) HasThinProvisioned() bool`

HasThinProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StoragePureArrayUtilizationAllOf) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StoragePureArrayUtilizationAllOf) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StoragePureArrayUtilizationAllOf) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StoragePureArrayUtilizationAllOf) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetVolume

`func (o *StoragePureArrayUtilizationAllOf) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StoragePureArrayUtilizationAllOf) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StoragePureArrayUtilizationAllOf) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StoragePureArrayUtilizationAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


