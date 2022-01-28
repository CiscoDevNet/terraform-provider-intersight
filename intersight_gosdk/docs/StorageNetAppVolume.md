# StorageNetAppVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppVolume"]
**AutosizeMode** | Pointer to **string** | The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink. * &#x60;off&#x60; - The volume will not grow or shrink in size in response to the amount of used space. * &#x60;grow&#x60; - The volume will automatically grow when used space in the volume is above the grow threshold. * &#x60;grow_shrink&#x60; - The volume will grow or shrink in size in response to the amount of used space. | [optional] [readonly] [default to "off"]
**AvgPerformanceMetrics** | Pointer to [**StorageNetAppPerformanceMetricsAverage**](StorageNetAppPerformanceMetricsAverage.md) |  | [optional] 
**ExportPolicyName** | Pointer to **string** | The name of the Export Policy. | [optional] [readonly] 
**Key** | Pointer to **string** | Unique identifier of NetApp Volume across data center. | [optional] [readonly] 
**SnapshotPolicyName** | Pointer to **string** | The name of the Snapshot Policy. | [optional] [readonly] 
**SnapshotPolicyUuid** | Pointer to **string** | The UUID of the Snapshot Policy. | [optional] [readonly] 
**SnapshotReservePercent** | Pointer to **int64** | The space that has been set aside as a reserve for Snapshot copy usage represented as a percent. | [optional] [readonly] 
**SnapshotUsed** | Pointer to **float64** | The total space used by Snapshot copies in the volume represented in bytes. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of a NetApp volume. * &#x60;offline&#x60; - Read and write access to the volume is not allowed. * &#x60;online&#x60; - Read and write access to the volume is allowed. * &#x60;error&#x60; - Storage volume state of error type. * &#x60;mixed&#x60; - The constituents of a FlexGroup volume are not all in the same state. | [optional] [readonly] [default to "offline"]
**Type** | Pointer to **string** | NetApp volume type. The volume type can be Read-write, Data-protection, or Load-sharing. * &#x60;data-protection&#x60; - Prevents modification of the data on the Volume. * &#x60;read-write&#x60; - Data on the Volume can be modified. * &#x60;load-sharing&#x60; - The volume type is Load Sharing DP. | [optional] [readonly] [default to "data-protection"]
**Uuid** | Pointer to **string** | Universally unique identifier of a NetApp Volume. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**DiskPool** | Pointer to [**[]StorageNetAppAggregateRelationship**](StorageNetAppAggregateRelationship.md) | An array of relationships to storageNetAppAggregate resources. | [optional] [readonly] 
**Events** | Pointer to [**[]StorageNetAppVolumeEventRelationship**](StorageNetAppVolumeEventRelationship.md) | An array of relationships to storageNetAppVolumeEvent resources. | [optional] [readonly] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppVolume

`func NewStorageNetAppVolume(classId string, objectType string, ) *StorageNetAppVolume`

NewStorageNetAppVolume instantiates a new StorageNetAppVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppVolumeWithDefaults

`func NewStorageNetAppVolumeWithDefaults() *StorageNetAppVolume`

NewStorageNetAppVolumeWithDefaults instantiates a new StorageNetAppVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutosizeMode

`func (o *StorageNetAppVolume) GetAutosizeMode() string`

GetAutosizeMode returns the AutosizeMode field if non-nil, zero value otherwise.

### GetAutosizeModeOk

`func (o *StorageNetAppVolume) GetAutosizeModeOk() (*string, bool)`

GetAutosizeModeOk returns a tuple with the AutosizeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutosizeMode

`func (o *StorageNetAppVolume) SetAutosizeMode(v string)`

SetAutosizeMode sets AutosizeMode field to given value.

### HasAutosizeMode

`func (o *StorageNetAppVolume) HasAutosizeMode() bool`

HasAutosizeMode returns a boolean if a field has been set.

### GetAvgPerformanceMetrics

`func (o *StorageNetAppVolume) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage`

GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field if non-nil, zero value otherwise.

### GetAvgPerformanceMetricsOk

`func (o *StorageNetAppVolume) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool)`

GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPerformanceMetrics

`func (o *StorageNetAppVolume) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage)`

SetAvgPerformanceMetrics sets AvgPerformanceMetrics field to given value.

### HasAvgPerformanceMetrics

`func (o *StorageNetAppVolume) HasAvgPerformanceMetrics() bool`

HasAvgPerformanceMetrics returns a boolean if a field has been set.

### GetExportPolicyName

`func (o *StorageNetAppVolume) GetExportPolicyName() string`

GetExportPolicyName returns the ExportPolicyName field if non-nil, zero value otherwise.

### GetExportPolicyNameOk

`func (o *StorageNetAppVolume) GetExportPolicyNameOk() (*string, bool)`

GetExportPolicyNameOk returns a tuple with the ExportPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicyName

`func (o *StorageNetAppVolume) SetExportPolicyName(v string)`

SetExportPolicyName sets ExportPolicyName field to given value.

### HasExportPolicyName

`func (o *StorageNetAppVolume) HasExportPolicyName() bool`

HasExportPolicyName returns a boolean if a field has been set.

### GetKey

`func (o *StorageNetAppVolume) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StorageNetAppVolume) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StorageNetAppVolume) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StorageNetAppVolume) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSnapshotPolicyName

`func (o *StorageNetAppVolume) GetSnapshotPolicyName() string`

GetSnapshotPolicyName returns the SnapshotPolicyName field if non-nil, zero value otherwise.

### GetSnapshotPolicyNameOk

`func (o *StorageNetAppVolume) GetSnapshotPolicyNameOk() (*string, bool)`

GetSnapshotPolicyNameOk returns a tuple with the SnapshotPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyName

`func (o *StorageNetAppVolume) SetSnapshotPolicyName(v string)`

SetSnapshotPolicyName sets SnapshotPolicyName field to given value.

### HasSnapshotPolicyName

`func (o *StorageNetAppVolume) HasSnapshotPolicyName() bool`

HasSnapshotPolicyName returns a boolean if a field has been set.

### GetSnapshotPolicyUuid

`func (o *StorageNetAppVolume) GetSnapshotPolicyUuid() string`

GetSnapshotPolicyUuid returns the SnapshotPolicyUuid field if non-nil, zero value otherwise.

### GetSnapshotPolicyUuidOk

`func (o *StorageNetAppVolume) GetSnapshotPolicyUuidOk() (*string, bool)`

GetSnapshotPolicyUuidOk returns a tuple with the SnapshotPolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyUuid

`func (o *StorageNetAppVolume) SetSnapshotPolicyUuid(v string)`

SetSnapshotPolicyUuid sets SnapshotPolicyUuid field to given value.

### HasSnapshotPolicyUuid

`func (o *StorageNetAppVolume) HasSnapshotPolicyUuid() bool`

HasSnapshotPolicyUuid returns a boolean if a field has been set.

### GetSnapshotReservePercent

`func (o *StorageNetAppVolume) GetSnapshotReservePercent() int64`

GetSnapshotReservePercent returns the SnapshotReservePercent field if non-nil, zero value otherwise.

### GetSnapshotReservePercentOk

`func (o *StorageNetAppVolume) GetSnapshotReservePercentOk() (*int64, bool)`

GetSnapshotReservePercentOk returns a tuple with the SnapshotReservePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotReservePercent

`func (o *StorageNetAppVolume) SetSnapshotReservePercent(v int64)`

SetSnapshotReservePercent sets SnapshotReservePercent field to given value.

### HasSnapshotReservePercent

`func (o *StorageNetAppVolume) HasSnapshotReservePercent() bool`

HasSnapshotReservePercent returns a boolean if a field has been set.

### GetSnapshotUsed

`func (o *StorageNetAppVolume) GetSnapshotUsed() float64`

GetSnapshotUsed returns the SnapshotUsed field if non-nil, zero value otherwise.

### GetSnapshotUsedOk

`func (o *StorageNetAppVolume) GetSnapshotUsedOk() (*float64, bool)`

GetSnapshotUsedOk returns a tuple with the SnapshotUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUsed

`func (o *StorageNetAppVolume) SetSnapshotUsed(v float64)`

SetSnapshotUsed sets SnapshotUsed field to given value.

### HasSnapshotUsed

`func (o *StorageNetAppVolume) HasSnapshotUsed() bool`

HasSnapshotUsed returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppVolume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppVolume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppVolume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppVolume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppVolume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppVolume) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppVolume) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppVolume) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppVolume) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetDiskPool

`func (o *StorageNetAppVolume) GetDiskPool() []StorageNetAppAggregateRelationship`

GetDiskPool returns the DiskPool field if non-nil, zero value otherwise.

### GetDiskPoolOk

`func (o *StorageNetAppVolume) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool)`

GetDiskPoolOk returns a tuple with the DiskPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPool

`func (o *StorageNetAppVolume) SetDiskPool(v []StorageNetAppAggregateRelationship)`

SetDiskPool sets DiskPool field to given value.

### HasDiskPool

`func (o *StorageNetAppVolume) HasDiskPool() bool`

HasDiskPool returns a boolean if a field has been set.

### SetDiskPoolNil

`func (o *StorageNetAppVolume) SetDiskPoolNil(b bool)`

 SetDiskPoolNil sets the value for DiskPool to be an explicit nil

### UnsetDiskPool
`func (o *StorageNetAppVolume) UnsetDiskPool()`

UnsetDiskPool ensures that no value is present for DiskPool, not even an explicit nil
### GetEvents

`func (o *StorageNetAppVolume) GetEvents() []StorageNetAppVolumeEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppVolume) GetEventsOk() (*[]StorageNetAppVolumeEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppVolume) SetEvents(v []StorageNetAppVolumeEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppVolume) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppVolume) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppVolume) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetTenant

`func (o *StorageNetAppVolume) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppVolume) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppVolume) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppVolume) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


