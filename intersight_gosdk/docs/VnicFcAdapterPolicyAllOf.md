# VnicFcAdapterPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcAdapterPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcAdapterPolicy"]
**ErrorDetectionTimeout** | Pointer to **int64** | Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. | [optional] [default to 2000]
**ErrorRecoverySettings** | Pointer to [**NullableVnicFcErrorRecoverySettings**](VnicFcErrorRecoverySettings.md) |  | [optional] 
**FlogiSettings** | Pointer to [**NullableVnicFlogiSettings**](VnicFlogiSettings.md) |  | [optional] 
**InterruptSettings** | Pointer to [**NullableVnicFcInterruptSettings**](VnicFcInterruptSettings.md) |  | [optional] 
**IoThrottleCount** | Pointer to **int64** | The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. | [optional] [default to 512]
**LunCount** | Pointer to **int64** | The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. | [optional] [default to 1024]
**LunQueueDepth** | Pointer to **int64** | The number of commands that the HBA can send and receive in a single transmission per LUN. | [optional] [default to 20]
**PlogiSettings** | Pointer to [**NullableVnicPlogiSettings**](VnicPlogiSettings.md) |  | [optional] 
**ResourceAllocationTimeout** | Pointer to **int64** | Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. | [optional] [default to 10000]
**RxQueueSettings** | Pointer to [**NullableVnicFcQueueSettings**](VnicFcQueueSettings.md) |  | [optional] 
**ScsiQueueSettings** | Pointer to [**NullableVnicScsiQueueSettings**](VnicScsiQueueSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**NullableVnicFcQueueSettings**](VnicFcQueueSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicFcAdapterPolicyAllOf

`func NewVnicFcAdapterPolicyAllOf(classId string, objectType string, ) *VnicFcAdapterPolicyAllOf`

NewVnicFcAdapterPolicyAllOf instantiates a new VnicFcAdapterPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcAdapterPolicyAllOfWithDefaults

`func NewVnicFcAdapterPolicyAllOfWithDefaults() *VnicFcAdapterPolicyAllOf`

NewVnicFcAdapterPolicyAllOfWithDefaults instantiates a new VnicFcAdapterPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcAdapterPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcAdapterPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcAdapterPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcAdapterPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcAdapterPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcAdapterPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyAllOf) GetErrorDetectionTimeout() int64`

GetErrorDetectionTimeout returns the ErrorDetectionTimeout field if non-nil, zero value otherwise.

### GetErrorDetectionTimeoutOk

`func (o *VnicFcAdapterPolicyAllOf) GetErrorDetectionTimeoutOk() (*int64, bool)`

GetErrorDetectionTimeoutOk returns a tuple with the ErrorDetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyAllOf) SetErrorDetectionTimeout(v int64)`

SetErrorDetectionTimeout sets ErrorDetectionTimeout field to given value.

### HasErrorDetectionTimeout

`func (o *VnicFcAdapterPolicyAllOf) HasErrorDetectionTimeout() bool`

HasErrorDetectionTimeout returns a boolean if a field has been set.

### GetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyAllOf) GetErrorRecoverySettings() VnicFcErrorRecoverySettings`

GetErrorRecoverySettings returns the ErrorRecoverySettings field if non-nil, zero value otherwise.

### GetErrorRecoverySettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetErrorRecoverySettingsOk() (*VnicFcErrorRecoverySettings, bool)`

GetErrorRecoverySettingsOk returns a tuple with the ErrorRecoverySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRecoverySettings

`func (o *VnicFcAdapterPolicyAllOf) SetErrorRecoverySettings(v VnicFcErrorRecoverySettings)`

SetErrorRecoverySettings sets ErrorRecoverySettings field to given value.

### HasErrorRecoverySettings

`func (o *VnicFcAdapterPolicyAllOf) HasErrorRecoverySettings() bool`

HasErrorRecoverySettings returns a boolean if a field has been set.

### SetErrorRecoverySettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetErrorRecoverySettingsNil(b bool)`

 SetErrorRecoverySettingsNil sets the value for ErrorRecoverySettings to be an explicit nil

### UnsetErrorRecoverySettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetErrorRecoverySettings()`

UnsetErrorRecoverySettings ensures that no value is present for ErrorRecoverySettings, not even an explicit nil
### GetFlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) GetFlogiSettings() VnicFlogiSettings`

GetFlogiSettings returns the FlogiSettings field if non-nil, zero value otherwise.

### GetFlogiSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetFlogiSettingsOk() (*VnicFlogiSettings, bool)`

GetFlogiSettingsOk returns a tuple with the FlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) SetFlogiSettings(v VnicFlogiSettings)`

SetFlogiSettings sets FlogiSettings field to given value.

### HasFlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) HasFlogiSettings() bool`

HasFlogiSettings returns a boolean if a field has been set.

### SetFlogiSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetFlogiSettingsNil(b bool)`

 SetFlogiSettingsNil sets the value for FlogiSettings to be an explicit nil

### UnsetFlogiSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetFlogiSettings()`

UnsetFlogiSettings ensures that no value is present for FlogiSettings, not even an explicit nil
### GetInterruptSettings

`func (o *VnicFcAdapterPolicyAllOf) GetInterruptSettings() VnicFcInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetInterruptSettingsOk() (*VnicFcInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicFcAdapterPolicyAllOf) SetInterruptSettings(v VnicFcInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicFcAdapterPolicyAllOf) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### SetInterruptSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetInterruptSettingsNil(b bool)`

 SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil

### UnsetInterruptSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetInterruptSettings()`

UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
### GetIoThrottleCount

`func (o *VnicFcAdapterPolicyAllOf) GetIoThrottleCount() int64`

GetIoThrottleCount returns the IoThrottleCount field if non-nil, zero value otherwise.

### GetIoThrottleCountOk

`func (o *VnicFcAdapterPolicyAllOf) GetIoThrottleCountOk() (*int64, bool)`

GetIoThrottleCountOk returns a tuple with the IoThrottleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoThrottleCount

`func (o *VnicFcAdapterPolicyAllOf) SetIoThrottleCount(v int64)`

SetIoThrottleCount sets IoThrottleCount field to given value.

### HasIoThrottleCount

`func (o *VnicFcAdapterPolicyAllOf) HasIoThrottleCount() bool`

HasIoThrottleCount returns a boolean if a field has been set.

### GetLunCount

`func (o *VnicFcAdapterPolicyAllOf) GetLunCount() int64`

GetLunCount returns the LunCount field if non-nil, zero value otherwise.

### GetLunCountOk

`func (o *VnicFcAdapterPolicyAllOf) GetLunCountOk() (*int64, bool)`

GetLunCountOk returns a tuple with the LunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunCount

`func (o *VnicFcAdapterPolicyAllOf) SetLunCount(v int64)`

SetLunCount sets LunCount field to given value.

### HasLunCount

`func (o *VnicFcAdapterPolicyAllOf) HasLunCount() bool`

HasLunCount returns a boolean if a field has been set.

### GetLunQueueDepth

`func (o *VnicFcAdapterPolicyAllOf) GetLunQueueDepth() int64`

GetLunQueueDepth returns the LunQueueDepth field if non-nil, zero value otherwise.

### GetLunQueueDepthOk

`func (o *VnicFcAdapterPolicyAllOf) GetLunQueueDepthOk() (*int64, bool)`

GetLunQueueDepthOk returns a tuple with the LunQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunQueueDepth

`func (o *VnicFcAdapterPolicyAllOf) SetLunQueueDepth(v int64)`

SetLunQueueDepth sets LunQueueDepth field to given value.

### HasLunQueueDepth

`func (o *VnicFcAdapterPolicyAllOf) HasLunQueueDepth() bool`

HasLunQueueDepth returns a boolean if a field has been set.

### GetPlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) GetPlogiSettings() VnicPlogiSettings`

GetPlogiSettings returns the PlogiSettings field if non-nil, zero value otherwise.

### GetPlogiSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetPlogiSettingsOk() (*VnicPlogiSettings, bool)`

GetPlogiSettingsOk returns a tuple with the PlogiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) SetPlogiSettings(v VnicPlogiSettings)`

SetPlogiSettings sets PlogiSettings field to given value.

### HasPlogiSettings

`func (o *VnicFcAdapterPolicyAllOf) HasPlogiSettings() bool`

HasPlogiSettings returns a boolean if a field has been set.

### SetPlogiSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetPlogiSettingsNil(b bool)`

 SetPlogiSettingsNil sets the value for PlogiSettings to be an explicit nil

### UnsetPlogiSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetPlogiSettings()`

UnsetPlogiSettings ensures that no value is present for PlogiSettings, not even an explicit nil
### GetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyAllOf) GetResourceAllocationTimeout() int64`

GetResourceAllocationTimeout returns the ResourceAllocationTimeout field if non-nil, zero value otherwise.

### GetResourceAllocationTimeoutOk

`func (o *VnicFcAdapterPolicyAllOf) GetResourceAllocationTimeoutOk() (*int64, bool)`

GetResourceAllocationTimeoutOk returns a tuple with the ResourceAllocationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyAllOf) SetResourceAllocationTimeout(v int64)`

SetResourceAllocationTimeout sets ResourceAllocationTimeout field to given value.

### HasResourceAllocationTimeout

`func (o *VnicFcAdapterPolicyAllOf) HasResourceAllocationTimeout() bool`

HasResourceAllocationTimeout returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) GetRxQueueSettings() VnicFcQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetRxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) SetRxQueueSettings(v VnicFcQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### SetRxQueueSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetRxQueueSettingsNil(b bool)`

 SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil

### UnsetRxQueueSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetRxQueueSettings()`

UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
### GetScsiQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) GetScsiQueueSettings() VnicScsiQueueSettings`

GetScsiQueueSettings returns the ScsiQueueSettings field if non-nil, zero value otherwise.

### GetScsiQueueSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetScsiQueueSettingsOk() (*VnicScsiQueueSettings, bool)`

GetScsiQueueSettingsOk returns a tuple with the ScsiQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) SetScsiQueueSettings(v VnicScsiQueueSettings)`

SetScsiQueueSettings sets ScsiQueueSettings field to given value.

### HasScsiQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) HasScsiQueueSettings() bool`

HasScsiQueueSettings returns a boolean if a field has been set.

### SetScsiQueueSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetScsiQueueSettingsNil(b bool)`

 SetScsiQueueSettingsNil sets the value for ScsiQueueSettings to be an explicit nil

### UnsetScsiQueueSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetScsiQueueSettings()`

UnsetScsiQueueSettings ensures that no value is present for ScsiQueueSettings, not even an explicit nil
### GetTxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) GetTxQueueSettings() VnicFcQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicFcAdapterPolicyAllOf) GetTxQueueSettingsOk() (*VnicFcQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) SetTxQueueSettings(v VnicFcQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicFcAdapterPolicyAllOf) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### SetTxQueueSettingsNil

`func (o *VnicFcAdapterPolicyAllOf) SetTxQueueSettingsNil(b bool)`

 SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil

### UnsetTxQueueSettings
`func (o *VnicFcAdapterPolicyAllOf) UnsetTxQueueSettings()`

UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
### GetOrganization

`func (o *VnicFcAdapterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicFcAdapterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicFcAdapterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicFcAdapterPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


