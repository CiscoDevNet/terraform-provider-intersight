# StorageHitachiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.HitachiPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.HitachiPool"]
**BlockingModeBlockade** | Pointer to **string** | Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. | [optional] [readonly] 
**BlockingModeFull** | Pointer to **string** | Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available. | [optional] [readonly] 
**DepletionThreshold** | Pointer to **string** | The depletion threshold set for the pool (%). | [optional] [readonly] 
**IsShrinking** | Pointer to **bool** | Whether the pool is shrinking is output. | [optional] [readonly] 
**MonitoringMode** | Pointer to **string** | Performance monitoring execution mode (monitor mode). * &#x60;N/A&#x60; - Performance monitoring is not available. * &#x60;Period mode&#x60; - Period mode is the default setting. If Period mode is enabled, tier range values and page relocations are determined based solely on the monitoring data from the last complete cycle. * &#x60;Continuous mode&#x60; - When Continuous mode is enabled, the weighted average efficiency is calculated using the latest monitoring information and the collected monitoring information in the past cycles. Page relocations are determined using this weighted average efficiency. | [optional] [readonly] [default to "N/A"]
**MonitoringStatus** | Pointer to **string** | Status of monitor information. | [optional] [readonly] 
**PoolActionMode** | Pointer to **string** | Execution mode for the pool. * &#x60;N/A&#x60; - Execution Mode is not available for the pool. * &#x60;Auto&#x60; - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator). * &#x60;Manual&#x60; - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC. | [optional] [readonly] [default to "N/A"]
**ProgressOfReplacing** | Pointer to **string** | Displays the status of the tier relocation processing. | [optional] [readonly] 
**TotalReservedCapacity** | Pointer to **int64** | Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool. | [optional] [readonly] 
**WarningThreshold** | Pointer to **int64** | The warning threshold set for the pool (%). | [optional] [readonly] 
**Array** | Pointer to [**StorageHitachiArrayRelationship**](StorageHitachiArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStorageHitachiPool

`func NewStorageHitachiPool(classId string, objectType string, ) *StorageHitachiPool`

NewStorageHitachiPool instantiates a new StorageHitachiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageHitachiPoolWithDefaults

`func NewStorageHitachiPoolWithDefaults() *StorageHitachiPool`

NewStorageHitachiPoolWithDefaults instantiates a new StorageHitachiPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageHitachiPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageHitachiPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageHitachiPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageHitachiPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageHitachiPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageHitachiPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockingModeBlockade

`func (o *StorageHitachiPool) GetBlockingModeBlockade() string`

GetBlockingModeBlockade returns the BlockingModeBlockade field if non-nil, zero value otherwise.

### GetBlockingModeBlockadeOk

`func (o *StorageHitachiPool) GetBlockingModeBlockadeOk() (*string, bool)`

GetBlockingModeBlockadeOk returns a tuple with the BlockingModeBlockade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingModeBlockade

`func (o *StorageHitachiPool) SetBlockingModeBlockade(v string)`

SetBlockingModeBlockade sets BlockingModeBlockade field to given value.

### HasBlockingModeBlockade

`func (o *StorageHitachiPool) HasBlockingModeBlockade() bool`

HasBlockingModeBlockade returns a boolean if a field has been set.

### GetBlockingModeFull

`func (o *StorageHitachiPool) GetBlockingModeFull() string`

GetBlockingModeFull returns the BlockingModeFull field if non-nil, zero value otherwise.

### GetBlockingModeFullOk

`func (o *StorageHitachiPool) GetBlockingModeFullOk() (*string, bool)`

GetBlockingModeFullOk returns a tuple with the BlockingModeFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingModeFull

`func (o *StorageHitachiPool) SetBlockingModeFull(v string)`

SetBlockingModeFull sets BlockingModeFull field to given value.

### HasBlockingModeFull

`func (o *StorageHitachiPool) HasBlockingModeFull() bool`

HasBlockingModeFull returns a boolean if a field has been set.

### GetDepletionThreshold

`func (o *StorageHitachiPool) GetDepletionThreshold() string`

GetDepletionThreshold returns the DepletionThreshold field if non-nil, zero value otherwise.

### GetDepletionThresholdOk

`func (o *StorageHitachiPool) GetDepletionThresholdOk() (*string, bool)`

GetDepletionThresholdOk returns a tuple with the DepletionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepletionThreshold

`func (o *StorageHitachiPool) SetDepletionThreshold(v string)`

SetDepletionThreshold sets DepletionThreshold field to given value.

### HasDepletionThreshold

`func (o *StorageHitachiPool) HasDepletionThreshold() bool`

HasDepletionThreshold returns a boolean if a field has been set.

### GetIsShrinking

`func (o *StorageHitachiPool) GetIsShrinking() bool`

GetIsShrinking returns the IsShrinking field if non-nil, zero value otherwise.

### GetIsShrinkingOk

`func (o *StorageHitachiPool) GetIsShrinkingOk() (*bool, bool)`

GetIsShrinkingOk returns a tuple with the IsShrinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShrinking

`func (o *StorageHitachiPool) SetIsShrinking(v bool)`

SetIsShrinking sets IsShrinking field to given value.

### HasIsShrinking

`func (o *StorageHitachiPool) HasIsShrinking() bool`

HasIsShrinking returns a boolean if a field has been set.

### GetMonitoringMode

`func (o *StorageHitachiPool) GetMonitoringMode() string`

GetMonitoringMode returns the MonitoringMode field if non-nil, zero value otherwise.

### GetMonitoringModeOk

`func (o *StorageHitachiPool) GetMonitoringModeOk() (*string, bool)`

GetMonitoringModeOk returns a tuple with the MonitoringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringMode

`func (o *StorageHitachiPool) SetMonitoringMode(v string)`

SetMonitoringMode sets MonitoringMode field to given value.

### HasMonitoringMode

`func (o *StorageHitachiPool) HasMonitoringMode() bool`

HasMonitoringMode returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *StorageHitachiPool) GetMonitoringStatus() string`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *StorageHitachiPool) GetMonitoringStatusOk() (*string, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *StorageHitachiPool) SetMonitoringStatus(v string)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *StorageHitachiPool) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### GetPoolActionMode

`func (o *StorageHitachiPool) GetPoolActionMode() string`

GetPoolActionMode returns the PoolActionMode field if non-nil, zero value otherwise.

### GetPoolActionModeOk

`func (o *StorageHitachiPool) GetPoolActionModeOk() (*string, bool)`

GetPoolActionModeOk returns a tuple with the PoolActionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolActionMode

`func (o *StorageHitachiPool) SetPoolActionMode(v string)`

SetPoolActionMode sets PoolActionMode field to given value.

### HasPoolActionMode

`func (o *StorageHitachiPool) HasPoolActionMode() bool`

HasPoolActionMode returns a boolean if a field has been set.

### GetProgressOfReplacing

`func (o *StorageHitachiPool) GetProgressOfReplacing() string`

GetProgressOfReplacing returns the ProgressOfReplacing field if non-nil, zero value otherwise.

### GetProgressOfReplacingOk

`func (o *StorageHitachiPool) GetProgressOfReplacingOk() (*string, bool)`

GetProgressOfReplacingOk returns a tuple with the ProgressOfReplacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressOfReplacing

`func (o *StorageHitachiPool) SetProgressOfReplacing(v string)`

SetProgressOfReplacing sets ProgressOfReplacing field to given value.

### HasProgressOfReplacing

`func (o *StorageHitachiPool) HasProgressOfReplacing() bool`

HasProgressOfReplacing returns a boolean if a field has been set.

### GetTotalReservedCapacity

`func (o *StorageHitachiPool) GetTotalReservedCapacity() int64`

GetTotalReservedCapacity returns the TotalReservedCapacity field if non-nil, zero value otherwise.

### GetTotalReservedCapacityOk

`func (o *StorageHitachiPool) GetTotalReservedCapacityOk() (*int64, bool)`

GetTotalReservedCapacityOk returns a tuple with the TotalReservedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReservedCapacity

`func (o *StorageHitachiPool) SetTotalReservedCapacity(v int64)`

SetTotalReservedCapacity sets TotalReservedCapacity field to given value.

### HasTotalReservedCapacity

`func (o *StorageHitachiPool) HasTotalReservedCapacity() bool`

HasTotalReservedCapacity returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *StorageHitachiPool) GetWarningThreshold() int64`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *StorageHitachiPool) GetWarningThresholdOk() (*int64, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *StorageHitachiPool) SetWarningThreshold(v int64)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *StorageHitachiPool) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.

### GetArray

`func (o *StorageHitachiPool) GetArray() StorageHitachiArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageHitachiPool) GetArrayOk() (*StorageHitachiArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageHitachiPool) SetArray(v StorageHitachiArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageHitachiPool) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *StorageHitachiPool) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StorageHitachiPool) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StorageHitachiPool) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StorageHitachiPool) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


