# StoragePureFlashBladeUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureFlashBladeUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureFlashBladeUtilization"]
**AvailableProvisioned** | Pointer to **int64** | The amount of space left that the current object can grow before writes are stopped due to a hard limit quota being hit. The value is total_provisioned minus the virtual space used. | [optional] [readonly] 
**AvailableRatio** | Pointer to **float32** | The ratio of the available space versus the total provisioned space. | [optional] [readonly] 
**DataReduction** | Pointer to **float32** | Reduction of data. Ratio of mapped sectors versus the amount of physical space the data occupies after data compression and deduplication. | [optional] [readonly] 
**Destroyed** | Pointer to **int64** | Unique physical space (excluding snapshots) occupied by destroyed data within the child containers, in bytes. | [optional] [readonly] 
**DestroyedVirtual** | Pointer to **int64** | The amount of destroyed logically written data within the child containers, in bytes. | [optional] [readonly] 
**Parity** | Pointer to **float32** | A representation of data redundancy on the array. Data redundancy is rebuilt automatically by the system whenever parity is less than 1.0. | [optional] [readonly] 
**Snapshot** | Pointer to **int64** | Physical usage by snapshots, other than unique in bytes. | [optional] [readonly] 
**TotalProvisioned** | Pointer to **int64** | The effective provisioned size of the container, at which a hard limit will be applied. For the array, this is the sum of the file systems and accounts. Measured in bytes. | [optional] [readonly] 
**Unique** | Pointer to **int64** | Unique physical space occupied by customer data, in bytes. Excludes shared space, snapshots, and metadata. | [optional] [readonly] 
**Virtual** | Pointer to **int64** | Logically consumed space for the object, in bytes. | [optional] [readonly] 

## Methods

### NewStoragePureFlashBladeUtilization

`func NewStoragePureFlashBladeUtilization(classId string, objectType string, ) *StoragePureFlashBladeUtilization`

NewStoragePureFlashBladeUtilization instantiates a new StoragePureFlashBladeUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureFlashBladeUtilizationWithDefaults

`func NewStoragePureFlashBladeUtilizationWithDefaults() *StoragePureFlashBladeUtilization`

NewStoragePureFlashBladeUtilizationWithDefaults instantiates a new StoragePureFlashBladeUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureFlashBladeUtilization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureFlashBladeUtilization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureFlashBladeUtilization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureFlashBladeUtilization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureFlashBladeUtilization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureFlashBladeUtilization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableProvisioned

`func (o *StoragePureFlashBladeUtilization) GetAvailableProvisioned() int64`

GetAvailableProvisioned returns the AvailableProvisioned field if non-nil, zero value otherwise.

### GetAvailableProvisionedOk

`func (o *StoragePureFlashBladeUtilization) GetAvailableProvisionedOk() (*int64, bool)`

GetAvailableProvisionedOk returns a tuple with the AvailableProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProvisioned

`func (o *StoragePureFlashBladeUtilization) SetAvailableProvisioned(v int64)`

SetAvailableProvisioned sets AvailableProvisioned field to given value.

### HasAvailableProvisioned

`func (o *StoragePureFlashBladeUtilization) HasAvailableProvisioned() bool`

HasAvailableProvisioned returns a boolean if a field has been set.

### GetAvailableRatio

`func (o *StoragePureFlashBladeUtilization) GetAvailableRatio() float32`

GetAvailableRatio returns the AvailableRatio field if non-nil, zero value otherwise.

### GetAvailableRatioOk

`func (o *StoragePureFlashBladeUtilization) GetAvailableRatioOk() (*float32, bool)`

GetAvailableRatioOk returns a tuple with the AvailableRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRatio

`func (o *StoragePureFlashBladeUtilization) SetAvailableRatio(v float32)`

SetAvailableRatio sets AvailableRatio field to given value.

### HasAvailableRatio

`func (o *StoragePureFlashBladeUtilization) HasAvailableRatio() bool`

HasAvailableRatio returns a boolean if a field has been set.

### GetDataReduction

`func (o *StoragePureFlashBladeUtilization) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StoragePureFlashBladeUtilization) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StoragePureFlashBladeUtilization) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StoragePureFlashBladeUtilization) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetDestroyed

`func (o *StoragePureFlashBladeUtilization) GetDestroyed() int64`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureFlashBladeUtilization) GetDestroyedOk() (*int64, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureFlashBladeUtilization) SetDestroyed(v int64)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureFlashBladeUtilization) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetDestroyedVirtual

`func (o *StoragePureFlashBladeUtilization) GetDestroyedVirtual() int64`

GetDestroyedVirtual returns the DestroyedVirtual field if non-nil, zero value otherwise.

### GetDestroyedVirtualOk

`func (o *StoragePureFlashBladeUtilization) GetDestroyedVirtualOk() (*int64, bool)`

GetDestroyedVirtualOk returns a tuple with the DestroyedVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyedVirtual

`func (o *StoragePureFlashBladeUtilization) SetDestroyedVirtual(v int64)`

SetDestroyedVirtual sets DestroyedVirtual field to given value.

### HasDestroyedVirtual

`func (o *StoragePureFlashBladeUtilization) HasDestroyedVirtual() bool`

HasDestroyedVirtual returns a boolean if a field has been set.

### GetParity

`func (o *StoragePureFlashBladeUtilization) GetParity() float32`

GetParity returns the Parity field if non-nil, zero value otherwise.

### GetParityOk

`func (o *StoragePureFlashBladeUtilization) GetParityOk() (*float32, bool)`

GetParityOk returns a tuple with the Parity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParity

`func (o *StoragePureFlashBladeUtilization) SetParity(v float32)`

SetParity sets Parity field to given value.

### HasParity

`func (o *StoragePureFlashBladeUtilization) HasParity() bool`

HasParity returns a boolean if a field has been set.

### GetSnapshot

`func (o *StoragePureFlashBladeUtilization) GetSnapshot() int64`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *StoragePureFlashBladeUtilization) GetSnapshotOk() (*int64, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *StoragePureFlashBladeUtilization) SetSnapshot(v int64)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *StoragePureFlashBladeUtilization) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetTotalProvisioned

`func (o *StoragePureFlashBladeUtilization) GetTotalProvisioned() int64`

GetTotalProvisioned returns the TotalProvisioned field if non-nil, zero value otherwise.

### GetTotalProvisionedOk

`func (o *StoragePureFlashBladeUtilization) GetTotalProvisionedOk() (*int64, bool)`

GetTotalProvisionedOk returns a tuple with the TotalProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProvisioned

`func (o *StoragePureFlashBladeUtilization) SetTotalProvisioned(v int64)`

SetTotalProvisioned sets TotalProvisioned field to given value.

### HasTotalProvisioned

`func (o *StoragePureFlashBladeUtilization) HasTotalProvisioned() bool`

HasTotalProvisioned returns a boolean if a field has been set.

### GetUnique

`func (o *StoragePureFlashBladeUtilization) GetUnique() int64`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *StoragePureFlashBladeUtilization) GetUniqueOk() (*int64, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *StoragePureFlashBladeUtilization) SetUnique(v int64)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *StoragePureFlashBladeUtilization) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetVirtual

`func (o *StoragePureFlashBladeUtilization) GetVirtual() int64`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *StoragePureFlashBladeUtilization) GetVirtualOk() (*int64, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *StoragePureFlashBladeUtilization) SetVirtual(v int64)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *StoragePureFlashBladeUtilization) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


