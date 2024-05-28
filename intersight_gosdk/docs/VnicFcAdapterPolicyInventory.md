# VnicFcAdapterPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcAdapterPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcAdapterPolicyInventory"]
**ErrorDetectionTimeout** | Pointer to **int64** | Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. | [optional] [readonly] [default to 2000]
**ErrorRecoverySettings** | Pointer to [**NullableVnicFcErrorRecoverySettings**](VnicFcErrorRecoverySettings.md) |  | [optional] 
**FlogiSettings** | Pointer to [**NullableVnicFlogiSettings**](VnicFlogiSettings.md) |  | [optional] 
**InterruptSettings** | Pointer to [**NullableVnicFcInterruptSettings**](VnicFcInterruptSettings.md) |  | [optional] 
**IoThrottleCount** | Pointer to **int64** | The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. | [optional] [readonly] [default to 512]
**LunCount** | Pointer to **int64** | The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. Lun Count value can exceed 1024 only for vHBA of type &#39;FC Initiator&#39; and on servers having supported firmware version. | [optional] [readonly] [default to 1024]
**LunQueueDepth** | Pointer to **int64** | The number of commands that the HBA can send and receive in a single transmission per LUN. | [optional] [readonly] [default to 20]
**PlogiSettings** | Pointer to [**NullableVnicPlogiSettings**](VnicPlogiSettings.md) |  | [optional] 
**ResourceAllocationTimeout** | Pointer to **int64** | Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. | [optional] [readonly] [default to 10000]
**RxQueueSettings** | Pointer to [**NullableVnicFcQueueSettings**](VnicFcQueueSettings.md) |  | [optional] 
**ScsiQueueSettings** | Pointer to [**NullableVnicScsiQueueSettings**](VnicScsiQueueSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**NullableVnicFcQueueSettings**](VnicFcQueueSettings.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicFcAdapterPolicyInventory

`func NewVnicFcAdapterPolicyInventory(classId string, objectType string, ) *VnicFcAdapterPolicyInventory`

NewVnicFcAdapterPolicyInventory instantiates a new VnicFcAdapterPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcAdapterPolicyInventoryWithDefaults

`func NewVnicFcAdapterPolicyInventoryWithDefaults() *VnicFcAdapterPolicyInventory`

NewVnicFcAdapterPolicyInventoryWithDefaults instantiates a new VnicFcAdapterPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcAdapterPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcAdapterPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcAdapterPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcAdapterPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcAdapterPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcAdapterPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyInventory) GetErrorDetectionTimeout() int64`

GetErrorDetectionTimeout returns the ErrorDetectionTimeout field if non-nil, zero value otherwise.

### GetErrorDetectionTimeoutOk

`func (o *VnicFcAdapterPolicyInventory) GetErrorDetectionTimeoutOk() (*int64, bool)`

GetErrorDetectionTimeoutOk returns a tuple with the ErrorDetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyInventory) SetErrorDetectionTimeout(v int64)`

SetErrorDetectionTimeout sets ErrorDetectionTimeout field to given value.

### HasErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyInventory) HasErrorDetectionTimeout() bool`

HasErrorDetectionTimeout returns a boolean if a field has been set.

### GetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyInventory) GetErrorRecoverySettings() VnicFcErrorRecoverySettings`

GetErrorRecoverySettings returns the ErrorRecoverySettings field if non-nil, zero value otherwise.

### GetErrorRecoverySettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetErrorRecoverySettingsOk() (*VnicFcErrorRecoverySettings, bool)`

GetErrorRecoverySettingsOk returns a tuple with the ErrorRecoverySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyInventory) SetErrorRecoverySettings(v VnicFcErrorRecoverySettings)`

SetErrorRecoverySettings sets ErrorRecoverySettings field to given value.

### HasErrorRecoverySettings

`func (o *VnicFcAdapterPolicyInventory) HasErrorRecoverySettings() bool`

HasErrorRecoverySettings returns a boolean if a field has been set.

### SetErrorRecoverySettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetErrorRecoverySettingsNil(b bool)`

 SetErrorRecoverySettingsNil sets the value for ErrorRecoverySettings to be an explicit nil

### UnsetErrorRecoverySettings
`func (o *VnicFcAdapterPolicyInventory) UnsetErrorRecoverySettings()`

UnsetErrorRecoverySettings ensures that no value is present for ErrorRecoverySettings, not even an explicit nil
### GetFlogiSettings

`func (o *VnicFcAdapterPolicyInventory) GetFlogiSettings() VnicFlogiSettings`

GetFlogiSettings returns the FlogiSettings field if non-nil, zero value otherwise.

### GetFlogiSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetFlogiSettingsOk() (*VnicFlogiSettings, bool)`

GetFlogiSettingsOk returns a tuple with the FlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlogiSettings

`func (o *VnicFcAdapterPolicyInventory) SetFlogiSettings(v VnicFlogiSettings)`

SetFlogiSettings sets FlogiSettings field to given value.

### HasFlogiSettings

`func (o *VnicFcAdapterPolicyInventory) HasFlogiSettings() bool`

HasFlogiSettings returns a boolean if a field has been set.

### SetFlogiSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetFlogiSettingsNil(b bool)`

 SetFlogiSettingsNil sets the value for FlogiSettings to be an explicit nil

### UnsetFlogiSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetFlogiSettings()`

UnsetFlogiSettings ensures that no value is present for FlogiSettings, not even an explicit nil
### GetInterruptSettings

`func (o *VnicFcAdapterPolicyInventory) GetInterruptSettings() VnicFcInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetInterruptSettingsOk() (*VnicFcInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicFcAdapterPolicyInventory) SetInterruptSettings(v VnicFcInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicFcAdapterPolicyInventory) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### SetInterruptSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetInterruptSettingsNil(b bool)`

 SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil

### UnsetInterruptSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetInterruptSettings()`

UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
### GetIoThrottleCount

`func (o *VnicFcAdapterPolicyInventory) GetIoThrottleCount() int64`

GetIoThrottleCount returns the IoThrottleCount field if non-nil, zero value otherwise.

### GetIoThrottleCountOk

`func (o *VnicFcAdapterPolicyInventory) GetIoThrottleCountOk() (*int64, bool)`

GetIoThrottleCountOk returns a tuple with the IoThrottleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoThrottleCount

`func (o *VnicFcAdapterPolicyInventory) SetIoThrottleCount(v int64)`

SetIoThrottleCount sets IoThrottleCount field to given value.

### HasIoThrottleCount

`func (o *VnicFcAdapterPolicyInventory) HasIoThrottleCount() bool`

HasIoThrottleCount returns a boolean if a field has been set.

### GetLunCount

`func (o *VnicFcAdapterPolicyInventory) GetLunCount() int64`

GetLunCount returns the LunCount field if non-nil, zero value otherwise.

### GetLunCountOk

`func (o *VnicFcAdapterPolicyInventory) GetLunCountOk() (*int64, bool)`

GetLunCountOk returns a tuple with the LunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunCount

`func (o *VnicFcAdapterPolicyInventory) SetLunCount(v int64)`

SetLunCount sets LunCount field to given value.

### HasLunCount

`func (o *VnicFcAdapterPolicyInventory) HasLunCount() bool`

HasLunCount returns a boolean if a field has been set.

### GetLunQueueDepth

`func (o *VnicFcAdapterPolicyInventory) GetLunQueueDepth() int64`

GetLunQueueDepth returns the LunQueueDepth field if non-nil, zero value otherwise.

### GetLunQueueDepthOk

`func (o *VnicFcAdapterPolicyInventory) GetLunQueueDepthOk() (*int64, bool)`

GetLunQueueDepthOk returns a tuple with the LunQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunQueueDepth

`func (o *VnicFcAdapterPolicyInventory) SetLunQueueDepth(v int64)`

SetLunQueueDepth sets LunQueueDepth field to given value.

### HasLunQueueDepth

`func (o *VnicFcAdapterPolicyInventory) HasLunQueueDepth() bool`

HasLunQueueDepth returns a boolean if a field has been set.

### GetPlogiSettings

`func (o *VnicFcAdapterPolicyInventory) GetPlogiSettings() VnicPlogiSettings`

GetPlogiSettings returns the PlogiSettings field if non-nil, zero value otherwise.

### GetPlogiSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetPlogiSettingsOk() (*VnicPlogiSettings, bool)`

GetPlogiSettingsOk returns a tuple with the PlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlogiSettings

`func (o *VnicFcAdapterPolicyInventory) SetPlogiSettings(v VnicPlogiSettings)`

SetPlogiSettings sets PlogiSettings field to given value.

### HasPlogiSettings

`func (o *VnicFcAdapterPolicyInventory) HasPlogiSettings() bool`

HasPlogiSettings returns a boolean if a field has been set.

### SetPlogiSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetPlogiSettingsNil(b bool)`

 SetPlogiSettingsNil sets the value for PlogiSettings to be an explicit nil

### UnsetPlogiSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetPlogiSettings()`

UnsetPlogiSettings ensures that no value is present for PlogiSettings, not even an explicit nil
### GetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyInventory) GetResourceAllocationTimeout() int64`

GetResourceAllocationTimeout returns the ResourceAllocationTimeout field if non-nil, zero value otherwise.

### GetResourceAllocationTimeoutOk

`func (o *VnicFcAdapterPolicyInventory) GetResourceAllocationTimeoutOk() (*int64, bool)`

GetResourceAllocationTimeoutOk returns a tuple with the ResourceAllocationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyInventory) SetResourceAllocationTimeout(v int64)`

SetResourceAllocationTimeout sets ResourceAllocationTimeout field to given value.

### HasResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyInventory) HasResourceAllocationTimeout() bool`

HasResourceAllocationTimeout returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) GetRxQueueSettings() VnicFcQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetRxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) SetRxQueueSettings(v VnicFcQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### SetRxQueueSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetRxQueueSettingsNil(b bool)`

 SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil

### UnsetRxQueueSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetRxQueueSettings()`

UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
### GetScsiQueueSettings

`func (o *VnicFcAdapterPolicyInventory) GetScsiQueueSettings() VnicScsiQueueSettings`

GetScsiQueueSettings returns the ScsiQueueSettings field if non-nil, zero value otherwise.

### GetScsiQueueSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetScsiQueueSettingsOk() (*VnicScsiQueueSettings, bool)`

GetScsiQueueSettingsOk returns a tuple with the ScsiQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiQueueSettings

`func (o *VnicFcAdapterPolicyInventory) SetScsiQueueSettings(v VnicScsiQueueSettings)`

SetScsiQueueSettings sets ScsiQueueSettings field to given value.

### HasScsiQueueSettings

`func (o *VnicFcAdapterPolicyInventory) HasScsiQueueSettings() bool`

HasScsiQueueSettings returns a boolean if a field has been set.

### SetScsiQueueSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetScsiQueueSettingsNil(b bool)`

 SetScsiQueueSettingsNil sets the value for ScsiQueueSettings to be an explicit nil

### UnsetScsiQueueSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetScsiQueueSettings()`

UnsetScsiQueueSettings ensures that no value is present for ScsiQueueSettings, not even an explicit nil
### GetTxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) GetTxQueueSettings() VnicFcQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicFcAdapterPolicyInventory) GetTxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) SetTxQueueSettings(v VnicFcQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicFcAdapterPolicyInventory) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### SetTxQueueSettingsNil

`func (o *VnicFcAdapterPolicyInventory) SetTxQueueSettingsNil(b bool)`

 SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil

### UnsetTxQueueSettings
`func (o *VnicFcAdapterPolicyInventory) UnsetTxQueueSettings()`

UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
### GetTargetMo

`func (o *VnicFcAdapterPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicFcAdapterPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicFcAdapterPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicFcAdapterPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *VnicFcAdapterPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *VnicFcAdapterPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


