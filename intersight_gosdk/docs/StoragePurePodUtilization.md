# StoragePurePodUtilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PurePodUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PurePodUtilization"]
**DataReduction** | Pointer to **float32** | Ratio of logical data to physical storage due to compression and deduplication only (excludes thin provisioning). For example, a ratio of 15.4 means 15.4 MB of logical data uses 1 MB of physical space. | [optional] [readonly] 
**Footprint** | Pointer to **int64** | Total physical space used by the pod on the array, including metadata overhead, snapshots, shared data, and system metadata. | [optional] [readonly] 
**Replication** | Pointer to **int64** | Amount of physical space used by replication activities for this pod. Value is 0 if replication is disabled. | [optional] [readonly] 
**Shared** | Pointer to **int64** | Amount of shared (deduplicated or referenced) data across volumes and snapshots in the pod. | [optional] [readonly] 
**Snapshots** | Pointer to **int64** | Number of snapshot objects present within the pod. | [optional] [readonly] 
**ThinProvisioning** | Pointer to **float32** | Efficiency ratio of thin provisioning. A value close to 1.0 indicates most provisioned space is used; a lower value indicates over-provisioning with less actual usage. | [optional] [readonly] 
**TotalPhysical** | Pointer to **int64** | Actual physical space consumed on flash by this pod&#39;s data. Equivalent to totalUsed. | [optional] [readonly] 
**TotalProvisioned** | Pointer to **int64** | Total logical size provisioned to volumes within the pod; amount of storage presented to hosts, regardless of how much is actually used. | [optional] [readonly] 
**TotalReduction** | Pointer to **float32** | The total data efficiency ratio achieved, factoring in both data reduction (compression and deduplication) and thin provisioning. For example, a ratio of 1800 means that 1800 units of logical data require only 1 unit of physical storage. | [optional] [readonly] 
**TotalUsed** | Pointer to **int64** | Total physical space used on the array by the pod, including unique, shared, and snapshot data. | [optional] [readonly] 
**Unique** | Pointer to **int64** | Amount of non-shared, unique data stored in the pod. | [optional] [readonly] 
**UsedProvisioned** | Pointer to **int64** | Amount of provisioned space that is currently being used by data written to volumes. | [optional] [readonly] 
**Virtual** | Pointer to **int64** | The total amount of logical data written to the volumes in the pod, as seen by the hosts. | [optional] [readonly] 

## Methods

### NewStoragePurePodUtilization

`func NewStoragePurePodUtilization(classId string, objectType string, ) *StoragePurePodUtilization`

NewStoragePurePodUtilization instantiates a new StoragePurePodUtilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePurePodUtilizationWithDefaults

`func NewStoragePurePodUtilizationWithDefaults() *StoragePurePodUtilization`

NewStoragePurePodUtilizationWithDefaults instantiates a new StoragePurePodUtilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePurePodUtilization) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePurePodUtilization) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePurePodUtilization) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePurePodUtilization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePurePodUtilization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePurePodUtilization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataReduction

`func (o *StoragePurePodUtilization) GetDataReduction() float32`

GetDataReduction returns the DataReduction field if non-nil, zero value otherwise.

### GetDataReductionOk

`func (o *StoragePurePodUtilization) GetDataReductionOk() (*float32, bool)`

GetDataReductionOk returns a tuple with the DataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReduction

`func (o *StoragePurePodUtilization) SetDataReduction(v float32)`

SetDataReduction sets DataReduction field to given value.

### HasDataReduction

`func (o *StoragePurePodUtilization) HasDataReduction() bool`

HasDataReduction returns a boolean if a field has been set.

### GetFootprint

`func (o *StoragePurePodUtilization) GetFootprint() int64`

GetFootprint returns the Footprint field if non-nil, zero value otherwise.

### GetFootprintOk

`func (o *StoragePurePodUtilization) GetFootprintOk() (*int64, bool)`

GetFootprintOk returns a tuple with the Footprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFootprint

`func (o *StoragePurePodUtilization) SetFootprint(v int64)`

SetFootprint sets Footprint field to given value.

### HasFootprint

`func (o *StoragePurePodUtilization) HasFootprint() bool`

HasFootprint returns a boolean if a field has been set.

### GetReplication

`func (o *StoragePurePodUtilization) GetReplication() int64`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *StoragePurePodUtilization) GetReplicationOk() (*int64, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *StoragePurePodUtilization) SetReplication(v int64)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *StoragePurePodUtilization) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetShared

`func (o *StoragePurePodUtilization) GetShared() int64`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StoragePurePodUtilization) GetSharedOk() (*int64, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StoragePurePodUtilization) SetShared(v int64)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StoragePurePodUtilization) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSnapshots

`func (o *StoragePurePodUtilization) GetSnapshots() int64`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *StoragePurePodUtilization) GetSnapshotsOk() (*int64, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *StoragePurePodUtilization) SetSnapshots(v int64)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *StoragePurePodUtilization) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetThinProvisioning

`func (o *StoragePurePodUtilization) GetThinProvisioning() float32`

GetThinProvisioning returns the ThinProvisioning field if non-nil, zero value otherwise.

### GetThinProvisioningOk

`func (o *StoragePurePodUtilization) GetThinProvisioningOk() (*float32, bool)`

GetThinProvisioningOk returns a tuple with the ThinProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioning

`func (o *StoragePurePodUtilization) SetThinProvisioning(v float32)`

SetThinProvisioning sets ThinProvisioning field to given value.

### HasThinProvisioning

`func (o *StoragePurePodUtilization) HasThinProvisioning() bool`

HasThinProvisioning returns a boolean if a field has been set.

### GetTotalPhysical

`func (o *StoragePurePodUtilization) GetTotalPhysical() int64`

GetTotalPhysical returns the TotalPhysical field if non-nil, zero value otherwise.

### GetTotalPhysicalOk

`func (o *StoragePurePodUtilization) GetTotalPhysicalOk() (*int64, bool)`

GetTotalPhysicalOk returns a tuple with the TotalPhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysical

`func (o *StoragePurePodUtilization) SetTotalPhysical(v int64)`

SetTotalPhysical sets TotalPhysical field to given value.

### HasTotalPhysical

`func (o *StoragePurePodUtilization) HasTotalPhysical() bool`

HasTotalPhysical returns a boolean if a field has been set.

### GetTotalProvisioned

`func (o *StoragePurePodUtilization) GetTotalProvisioned() int64`

GetTotalProvisioned returns the TotalProvisioned field if non-nil, zero value otherwise.

### GetTotalProvisionedOk

`func (o *StoragePurePodUtilization) GetTotalProvisionedOk() (*int64, bool)`

GetTotalProvisionedOk returns a tuple with the TotalProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProvisioned

`func (o *StoragePurePodUtilization) SetTotalProvisioned(v int64)`

SetTotalProvisioned sets TotalProvisioned field to given value.

### HasTotalProvisioned

`func (o *StoragePurePodUtilization) HasTotalProvisioned() bool`

HasTotalProvisioned returns a boolean if a field has been set.

### GetTotalReduction

`func (o *StoragePurePodUtilization) GetTotalReduction() float32`

GetTotalReduction returns the TotalReduction field if non-nil, zero value otherwise.

### GetTotalReductionOk

`func (o *StoragePurePodUtilization) GetTotalReductionOk() (*float32, bool)`

GetTotalReductionOk returns a tuple with the TotalReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReduction

`func (o *StoragePurePodUtilization) SetTotalReduction(v float32)`

SetTotalReduction sets TotalReduction field to given value.

### HasTotalReduction

`func (o *StoragePurePodUtilization) HasTotalReduction() bool`

HasTotalReduction returns a boolean if a field has been set.

### GetTotalUsed

`func (o *StoragePurePodUtilization) GetTotalUsed() int64`

GetTotalUsed returns the TotalUsed field if non-nil, zero value otherwise.

### GetTotalUsedOk

`func (o *StoragePurePodUtilization) GetTotalUsedOk() (*int64, bool)`

GetTotalUsedOk returns a tuple with the TotalUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsed

`func (o *StoragePurePodUtilization) SetTotalUsed(v int64)`

SetTotalUsed sets TotalUsed field to given value.

### HasTotalUsed

`func (o *StoragePurePodUtilization) HasTotalUsed() bool`

HasTotalUsed returns a boolean if a field has been set.

### GetUnique

`func (o *StoragePurePodUtilization) GetUnique() int64`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *StoragePurePodUtilization) GetUniqueOk() (*int64, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *StoragePurePodUtilization) SetUnique(v int64)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *StoragePurePodUtilization) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUsedProvisioned

`func (o *StoragePurePodUtilization) GetUsedProvisioned() int64`

GetUsedProvisioned returns the UsedProvisioned field if non-nil, zero value otherwise.

### GetUsedProvisionedOk

`func (o *StoragePurePodUtilization) GetUsedProvisionedOk() (*int64, bool)`

GetUsedProvisionedOk returns a tuple with the UsedProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedProvisioned

`func (o *StoragePurePodUtilization) SetUsedProvisioned(v int64)`

SetUsedProvisioned sets UsedProvisioned field to given value.

### HasUsedProvisioned

`func (o *StoragePurePodUtilization) HasUsedProvisioned() bool`

HasUsedProvisioned returns a boolean if a field has been set.

### GetVirtual

`func (o *StoragePurePodUtilization) GetVirtual() int64`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *StoragePurePodUtilization) GetVirtualOk() (*int64, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *StoragePurePodUtilization) SetVirtual(v int64)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *StoragePurePodUtilization) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


