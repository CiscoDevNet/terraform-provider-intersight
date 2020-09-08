# VnicFcAdapterPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDetectionTimeout** | Pointer to **int64** | Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred. | [optional] 
**ErrorRecoverySettings** | Pointer to [**VnicFcErrorRecoverySettings**](vnic.FcErrorRecoverySettings.md) |  | [optional] 
**FlogiSettings** | Pointer to [**VnicFlogiSettings**](vnic.FlogiSettings.md) |  | [optional] 
**InterruptSettings** | Pointer to [**VnicFcInterruptSettings**](vnic.FcInterruptSettings.md) |  | [optional] 
**IoThrottleCount** | Pointer to **int64** | The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed. | [optional] 
**LunCount** | Pointer to **int64** | The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server. | [optional] 
**LunQueueDepth** | Pointer to **int64** | The number of commands that the HBA can send and receive in a single transmission per LUN. | [optional] 
**PlogiSettings** | Pointer to [**VnicPlogiSettings**](vnic.PlogiSettings.md) |  | [optional] 
**ResourceAllocationTimeout** | Pointer to **int64** | Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated. | [optional] 
**RxQueueSettings** | Pointer to [**VnicFcQueueSettings**](vnic.FcQueueSettings.md) |  | [optional] 
**ScsiQueueSettings** | Pointer to [**VnicScsiQueueSettings**](vnic.ScsiQueueSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**VnicFcQueueSettings**](vnic.FcQueueSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicFcAdapterPolicyAllOf

`func NewVnicFcAdapterPolicyAllOf() *VnicFcAdapterPolicyAllOf`

NewVnicFcAdapterPolicyAllOf instantiates a new VnicFcAdapterPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcAdapterPolicyAllOfWithDefaults

`func NewVnicFcAdapterPolicyAllOfWithDefaults() *VnicFcAdapterPolicyAllOf`

NewVnicFcAdapterPolicyAllOfWithDefaults instantiates a new VnicFcAdapterPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


