# StoragePureRealmUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureRealmUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureRealmUtilization"]
**DataReduction** | Pointer to **float32** | Ratio of logical data to physical storage due to compression and deduplication only (excludes thin provisioning). | [optional] [readonly] 
**Footprint** | Pointer to **int64** | The \&quot;as-if only\&quot; physical space consumed by the realm&#39;s data, including metadata overhead and system space. | [optional] [readonly] 
**Shared** | Pointer to **int64** | Physical space containing shared data (deduplicated or referenced) across volumes and snapshots in the realm. | [optional] [readonly] 
**Snapshots** | Pointer to **int64** | Physical space attributed exclusively to snapshot data within this realm. | [optional] [readonly] 
**System** | Pointer to **int64** | Physical space used by system metadata overhead and other internal realm structures. | [optional] [readonly] 
**ThinProvisioning** | Pointer to **float32** | Efficiency ratio of thin provisioning, indicating savings between used provisioned space and virtual usage. | [optional] [readonly] 
**TotalProvisioned** | Pointer to **int64** | Total logical capacity provisioned to all objects (e.g., volumes) within the realm. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | The combined data efficiency ratio achieved through both data reduction (compression/deduplication) and thin provisioning. | [optional] [readonly] 
**TotalUsed** | Pointer to **int64** | Total physical space used on the array by the realm, combining unique, shared, and snapshot data (when available). | [optional] [readonly] 
**Unique** | Pointer to **int64** | Unique physical data stored in the realm that is not shared or deduplicated across other objects. | [optional] [readonly] 
**UsedProvisioned** | Pointer to **int64** | Amount of the provisioned space that is currently in use by data written to realm objects (logical size). | [optional] [readonly] 
**Virtual** | Pointer to **int64** | The total logical data written to the realm, as seen by hosts. Used in calculating reduction approximations. | [optional] [readonly] 

## Methods

### NewStoragePureRealmUtilization

`func NewStoragePureRealmUtilization(classId string, objectType string, ) *StoragePureRealmUtilization`

NewStoragePureRealmUtilization instantiates a new StoragePureRealmUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureRealmUtilizationWithDefaults

`func NewStoragePureRealmUtilizationWithDefaults() *StoragePureRealmUtilization`

NewStoragePureRealmUtilizationWithDefaults instantiates a new StoragePureRealmUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureRealmUtilization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureRealmUtilization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureRealmUtilization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureRealmUtilization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureRealmUtilization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureRealmUtilization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReduction

`func (o *StoragePureRealmUtilization) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StoragePureRealmUtilization) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StoragePureRealmUtilization) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StoragePureRealmUtilization) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetFootprint

`func (o *StoragePureRealmUtilization) GetFootprint() int64`

GetFootprint returns the Footprint field if non-nil, zero value otherwise.

### GetFootprintOk

`func (o *StoragePureRealmUtilization) GetFootprintOk() (*int64, bool)`

GetFootprintOk returns a tuple with the Footprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootprint

`func (o *StoragePureRealmUtilization) SetFootprint(v int64)`

SetFootprint sets Footprint field to given value.

### HasFootprint

`func (o *StoragePureRealmUtilization) HasFootprint() bool`

HasFootprint returns a boolean if a field has been set.

### GetShared

`func (o *StoragePureRealmUtilization) GetShared() int64`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StoragePureRealmUtilization) GetSharedOk() (*int64, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StoragePureRealmUtilization) SetShared(v int64)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StoragePureRealmUtilization) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSnapshots

`func (o *StoragePureRealmUtilization) GetSnapshots() int64`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *StoragePureRealmUtilization) GetSnapshotsOk() (*int64, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *StoragePureRealmUtilization) SetSnapshots(v int64)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *StoragePureRealmUtilization) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetSystem

`func (o *StoragePureRealmUtilization) GetSystem() int64`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *StoragePureRealmUtilization) GetSystemOk() (*int64, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *StoragePureRealmUtilization) SetSystem(v int64)`

SetSystem sets System field to given value.

### HasSystem

`func (o *StoragePureRealmUtilization) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetThinProvisioning

`func (o *StoragePureRealmUtilization) GetThinProvisioning() float32`

GetThinProvisioning returns the ThinProvisioning field if non-nil, zero value otherwise.

### GetThinProvisioningOk

`func (o *StoragePureRealmUtilization) GetThinProvisioningOk() (*float32, bool)`

GetThinProvisioningOk returns a tuple with the ThinProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioning

`func (o *StoragePureRealmUtilization) SetThinProvisioning(v float32)`

SetThinProvisioning sets ThinProvisioning field to given value.

### HasThinProvisioning

`func (o *StoragePureRealmUtilization) HasThinProvisioning() bool`

HasThinProvisioning returns a boolean if a field has been set.

### GetTotalProvisioned

`func (o *StoragePureRealmUtilization) GetTotalProvisioned() int64`

GetTotalProvisioned returns the TotalProvisioned field if non-nil, zero value otherwise.

### GetTotalProvisionedOk

`func (o *StoragePureRealmUtilization) GetTotalProvisionedOk() (*int64, bool)`

GetTotalProvisionedOk returns a tuple with the TotalProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProvisioned

`func (o *StoragePureRealmUtilization) SetTotalProvisioned(v int64)`

SetTotalProvisioned sets TotalProvisioned field to given value.

### HasTotalProvisioned

`func (o *StoragePureRealmUtilization) HasTotalProvisioned() bool`

HasTotalProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StoragePureRealmUtilization) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StoragePureRealmUtilization) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StoragePureRealmUtilization) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StoragePureRealmUtilization) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetTotalUsed

`func (o *StoragePureRealmUtilization) GetTotalUsed() int64`

GetTotalUsed returns the TotalUsed field if non-nil, zero value otherwise.

### GetTotalUsedOk

`func (o *StoragePureRealmUtilization) GetTotalUsedOk() (*int64, bool)`

GetTotalUsedOk returns a tuple with the TotalUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsed

`func (o *StoragePureRealmUtilization) SetTotalUsed(v int64)`

SetTotalUsed sets TotalUsed field to given value.

### HasTotalUsed

`func (o *StoragePureRealmUtilization) HasTotalUsed() bool`

HasTotalUsed returns a boolean if a field has been set.

### GetUnique

`func (o *StoragePureRealmUtilization) GetUnique() int64`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *StoragePureRealmUtilization) GetUniqueOk() (*int64, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *StoragePureRealmUtilization) SetUnique(v int64)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *StoragePureRealmUtilization) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUsedProvisioned

`func (o *StoragePureRealmUtilization) GetUsedProvisioned() int64`

GetUsedProvisioned returns the UsedProvisioned field if non-nil, zero value otherwise.

### GetUsedProvisionedOk

`func (o *StoragePureRealmUtilization) GetUsedProvisionedOk() (*int64, bool)`

GetUsedProvisionedOk returns a tuple with the UsedProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedProvisioned

`func (o *StoragePureRealmUtilization) SetUsedProvisioned(v int64)`

SetUsedProvisioned sets UsedProvisioned field to given value.

### HasUsedProvisioned

`func (o *StoragePureRealmUtilization) HasUsedProvisioned() bool`

HasUsedProvisioned returns a boolean if a field has been set.

### GetVirtual

`func (o *StoragePureRealmUtilization) GetVirtual() int64`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *StoragePureRealmUtilization) GetVirtualOk() (*int64, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *StoragePureRealmUtilization) SetVirtual(v int64)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *StoragePureRealmUtilization) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


