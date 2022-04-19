# VnicEthAdapterPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthAdapterPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthAdapterPolicyInventory"]
**AdvancedFilter** | Pointer to **bool** | Enables advanced filtering on the interface. | [optional] [readonly] [default to false]
**ArfsSettings** | Pointer to [**NullableVnicArfsSettings**](VnicArfsSettings.md) |  | [optional] 
**CompletionQueueSettings** | Pointer to [**NullableVnicCompletionQueueSettings**](VnicCompletionQueueSettings.md) |  | [optional] 
**GeneveEnabled** | Pointer to **bool** | GENEVE offload protocol allows you to create logical networks that span physical network boundaries by allowing any information to be encoded in a packet and passed between tunnel endpoints. | [optional] [readonly] [default to false]
**InterruptScaling** | Pointer to **bool** | Enables Interrupt Scaling on the interface. | [optional] [readonly] [default to false]
**InterruptSettings** | Pointer to [**NullableVnicEthInterruptSettings**](VnicEthInterruptSettings.md) |  | [optional] 
**NvgreSettings** | Pointer to [**NullableVnicNvgreSettings**](VnicNvgreSettings.md) |  | [optional] 
**PtpSettings** | Pointer to [**NullableVnicPtpSettings**](VnicPtpSettings.md) |  | [optional] 
**RoceSettings** | Pointer to [**NullableVnicRoceSettings**](VnicRoceSettings.md) |  | [optional] 
**RssHashSettings** | Pointer to [**NullableVnicRssHashSettings**](VnicRssHashSettings.md) |  | [optional] 
**RssSettings** | Pointer to **bool** | Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. | [optional] [readonly] [default to true]
**RxQueueSettings** | Pointer to [**NullableVnicEthRxQueueSettings**](VnicEthRxQueueSettings.md) |  | [optional] 
**TcpOffloadSettings** | Pointer to [**NullableVnicTcpOffloadSettings**](VnicTcpOffloadSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**NullableVnicEthTxQueueSettings**](VnicEthTxQueueSettings.md) |  | [optional] 
**UplinkFailbackTimeout** | Pointer to **int64** | Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC. | [optional] [readonly] [default to 5]
**VxlanSettings** | Pointer to [**NullableVnicVxlanSettings**](VnicVxlanSettings.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicEthAdapterPolicyInventory

`func NewVnicEthAdapterPolicyInventory(classId string, objectType string, ) *VnicEthAdapterPolicyInventory`

NewVnicEthAdapterPolicyInventory instantiates a new VnicEthAdapterPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthAdapterPolicyInventoryWithDefaults

`func NewVnicEthAdapterPolicyInventoryWithDefaults() *VnicEthAdapterPolicyInventory`

NewVnicEthAdapterPolicyInventoryWithDefaults instantiates a new VnicEthAdapterPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthAdapterPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthAdapterPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthAdapterPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthAdapterPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthAdapterPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthAdapterPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvancedFilter

`func (o *VnicEthAdapterPolicyInventory) GetAdvancedFilter() bool`

GetAdvancedFilter returns the AdvancedFilter field if non-nil, zero value otherwise.

### GetAdvancedFilterOk

`func (o *VnicEthAdapterPolicyInventory) GetAdvancedFilterOk() (*bool, bool)`

GetAdvancedFilterOk returns a tuple with the AdvancedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFilter

`func (o *VnicEthAdapterPolicyInventory) SetAdvancedFilter(v bool)`

SetAdvancedFilter sets AdvancedFilter field to given value.

### HasAdvancedFilter

`func (o *VnicEthAdapterPolicyInventory) HasAdvancedFilter() bool`

HasAdvancedFilter returns a boolean if a field has been set.

### GetArfsSettings

`func (o *VnicEthAdapterPolicyInventory) GetArfsSettings() VnicArfsSettings`

GetArfsSettings returns the ArfsSettings field if non-nil, zero value otherwise.

### GetArfsSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetArfsSettingsOk() (*VnicArfsSettings, bool)`

GetArfsSettingsOk returns a tuple with the ArfsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArfsSettings

`func (o *VnicEthAdapterPolicyInventory) SetArfsSettings(v VnicArfsSettings)`

SetArfsSettings sets ArfsSettings field to given value.

### HasArfsSettings

`func (o *VnicEthAdapterPolicyInventory) HasArfsSettings() bool`

HasArfsSettings returns a boolean if a field has been set.

### SetArfsSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetArfsSettingsNil(b bool)`

 SetArfsSettingsNil sets the value for ArfsSettings to be an explicit nil

### UnsetArfsSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetArfsSettings()`

UnsetArfsSettings ensures that no value is present for ArfsSettings, not even an explicit nil
### GetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventory) GetCompletionQueueSettings() VnicCompletionQueueSettings`

GetCompletionQueueSettings returns the CompletionQueueSettings field if non-nil, zero value otherwise.

### GetCompletionQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool)`

GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventory) SetCompletionQueueSettings(v VnicCompletionQueueSettings)`

SetCompletionQueueSettings sets CompletionQueueSettings field to given value.

### HasCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventory) HasCompletionQueueSettings() bool`

HasCompletionQueueSettings returns a boolean if a field has been set.

### SetCompletionQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetCompletionQueueSettingsNil(b bool)`

 SetCompletionQueueSettingsNil sets the value for CompletionQueueSettings to be an explicit nil

### UnsetCompletionQueueSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetCompletionQueueSettings()`

UnsetCompletionQueueSettings ensures that no value is present for CompletionQueueSettings, not even an explicit nil
### GetGeneveEnabled

`func (o *VnicEthAdapterPolicyInventory) GetGeneveEnabled() bool`

GetGeneveEnabled returns the GeneveEnabled field if non-nil, zero value otherwise.

### GetGeneveEnabledOk

`func (o *VnicEthAdapterPolicyInventory) GetGeneveEnabledOk() (*bool, bool)`

GetGeneveEnabledOk returns a tuple with the GeneveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneveEnabled

`func (o *VnicEthAdapterPolicyInventory) SetGeneveEnabled(v bool)`

SetGeneveEnabled sets GeneveEnabled field to given value.

### HasGeneveEnabled

`func (o *VnicEthAdapterPolicyInventory) HasGeneveEnabled() bool`

HasGeneveEnabled returns a boolean if a field has been set.

### GetInterruptScaling

`func (o *VnicEthAdapterPolicyInventory) GetInterruptScaling() bool`

GetInterruptScaling returns the InterruptScaling field if non-nil, zero value otherwise.

### GetInterruptScalingOk

`func (o *VnicEthAdapterPolicyInventory) GetInterruptScalingOk() (*bool, bool)`

GetInterruptScalingOk returns a tuple with the InterruptScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptScaling

`func (o *VnicEthAdapterPolicyInventory) SetInterruptScaling(v bool)`

SetInterruptScaling sets InterruptScaling field to given value.

### HasInterruptScaling

`func (o *VnicEthAdapterPolicyInventory) HasInterruptScaling() bool`

HasInterruptScaling returns a boolean if a field has been set.

### GetInterruptSettings

`func (o *VnicEthAdapterPolicyInventory) GetInterruptSettings() VnicEthInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicEthAdapterPolicyInventory) SetInterruptSettings(v VnicEthInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicEthAdapterPolicyInventory) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### SetInterruptSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetInterruptSettingsNil(b bool)`

 SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil

### UnsetInterruptSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetInterruptSettings()`

UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
### GetNvgreSettings

`func (o *VnicEthAdapterPolicyInventory) GetNvgreSettings() VnicNvgreSettings`

GetNvgreSettings returns the NvgreSettings field if non-nil, zero value otherwise.

### GetNvgreSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetNvgreSettingsOk() (*VnicNvgreSettings, bool)`

GetNvgreSettingsOk returns a tuple with the NvgreSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvgreSettings

`func (o *VnicEthAdapterPolicyInventory) SetNvgreSettings(v VnicNvgreSettings)`

SetNvgreSettings sets NvgreSettings field to given value.

### HasNvgreSettings

`func (o *VnicEthAdapterPolicyInventory) HasNvgreSettings() bool`

HasNvgreSettings returns a boolean if a field has been set.

### SetNvgreSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetNvgreSettingsNil(b bool)`

 SetNvgreSettingsNil sets the value for NvgreSettings to be an explicit nil

### UnsetNvgreSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetNvgreSettings()`

UnsetNvgreSettings ensures that no value is present for NvgreSettings, not even an explicit nil
### GetPtpSettings

`func (o *VnicEthAdapterPolicyInventory) GetPtpSettings() VnicPtpSettings`

GetPtpSettings returns the PtpSettings field if non-nil, zero value otherwise.

### GetPtpSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetPtpSettingsOk() (*VnicPtpSettings, bool)`

GetPtpSettingsOk returns a tuple with the PtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpSettings

`func (o *VnicEthAdapterPolicyInventory) SetPtpSettings(v VnicPtpSettings)`

SetPtpSettings sets PtpSettings field to given value.

### HasPtpSettings

`func (o *VnicEthAdapterPolicyInventory) HasPtpSettings() bool`

HasPtpSettings returns a boolean if a field has been set.

### SetPtpSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetPtpSettingsNil(b bool)`

 SetPtpSettingsNil sets the value for PtpSettings to be an explicit nil

### UnsetPtpSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetPtpSettings()`

UnsetPtpSettings ensures that no value is present for PtpSettings, not even an explicit nil
### GetRoceSettings

`func (o *VnicEthAdapterPolicyInventory) GetRoceSettings() VnicRoceSettings`

GetRoceSettings returns the RoceSettings field if non-nil, zero value otherwise.

### GetRoceSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetRoceSettingsOk() (*VnicRoceSettings, bool)`

GetRoceSettingsOk returns a tuple with the RoceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoceSettings

`func (o *VnicEthAdapterPolicyInventory) SetRoceSettings(v VnicRoceSettings)`

SetRoceSettings sets RoceSettings field to given value.

### HasRoceSettings

`func (o *VnicEthAdapterPolicyInventory) HasRoceSettings() bool`

HasRoceSettings returns a boolean if a field has been set.

### SetRoceSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetRoceSettingsNil(b bool)`

 SetRoceSettingsNil sets the value for RoceSettings to be an explicit nil

### UnsetRoceSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetRoceSettings()`

UnsetRoceSettings ensures that no value is present for RoceSettings, not even an explicit nil
### GetRssHashSettings

`func (o *VnicEthAdapterPolicyInventory) GetRssHashSettings() VnicRssHashSettings`

GetRssHashSettings returns the RssHashSettings field if non-nil, zero value otherwise.

### GetRssHashSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetRssHashSettingsOk() (*VnicRssHashSettings, bool)`

GetRssHashSettingsOk returns a tuple with the RssHashSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssHashSettings

`func (o *VnicEthAdapterPolicyInventory) SetRssHashSettings(v VnicRssHashSettings)`

SetRssHashSettings sets RssHashSettings field to given value.

### HasRssHashSettings

`func (o *VnicEthAdapterPolicyInventory) HasRssHashSettings() bool`

HasRssHashSettings returns a boolean if a field has been set.

### SetRssHashSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetRssHashSettingsNil(b bool)`

 SetRssHashSettingsNil sets the value for RssHashSettings to be an explicit nil

### UnsetRssHashSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetRssHashSettings()`

UnsetRssHashSettings ensures that no value is present for RssHashSettings, not even an explicit nil
### GetRssSettings

`func (o *VnicEthAdapterPolicyInventory) GetRssSettings() bool`

GetRssSettings returns the RssSettings field if non-nil, zero value otherwise.

### GetRssSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetRssSettingsOk() (*bool, bool)`

GetRssSettingsOk returns a tuple with the RssSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSettings

`func (o *VnicEthAdapterPolicyInventory) SetRssSettings(v bool)`

SetRssSettings sets RssSettings field to given value.

### HasRssSettings

`func (o *VnicEthAdapterPolicyInventory) HasRssSettings() bool`

HasRssSettings returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) GetRxQueueSettings() VnicEthRxQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) SetRxQueueSettings(v VnicEthRxQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### SetRxQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetRxQueueSettingsNil(b bool)`

 SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil

### UnsetRxQueueSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetRxQueueSettings()`

UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
### GetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventory) GetTcpOffloadSettings() VnicTcpOffloadSettings`

GetTcpOffloadSettings returns the TcpOffloadSettings field if non-nil, zero value otherwise.

### GetTcpOffloadSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool)`

GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventory) SetTcpOffloadSettings(v VnicTcpOffloadSettings)`

SetTcpOffloadSettings sets TcpOffloadSettings field to given value.

### HasTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventory) HasTcpOffloadSettings() bool`

HasTcpOffloadSettings returns a boolean if a field has been set.

### SetTcpOffloadSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetTcpOffloadSettingsNil(b bool)`

 SetTcpOffloadSettingsNil sets the value for TcpOffloadSettings to be an explicit nil

### UnsetTcpOffloadSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetTcpOffloadSettings()`

UnsetTcpOffloadSettings ensures that no value is present for TcpOffloadSettings, not even an explicit nil
### GetTxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) GetTxQueueSettings() VnicEthTxQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) SetTxQueueSettings(v VnicEthTxQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicEthAdapterPolicyInventory) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### SetTxQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetTxQueueSettingsNil(b bool)`

 SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil

### UnsetTxQueueSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetTxQueueSettings()`

UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
### GetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventory) GetUplinkFailbackTimeout() int64`

GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field if non-nil, zero value otherwise.

### GetUplinkFailbackTimeoutOk

`func (o *VnicEthAdapterPolicyInventory) GetUplinkFailbackTimeoutOk() (*int64, bool)`

GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventory) SetUplinkFailbackTimeout(v int64)`

SetUplinkFailbackTimeout sets UplinkFailbackTimeout field to given value.

### HasUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventory) HasUplinkFailbackTimeout() bool`

HasUplinkFailbackTimeout returns a boolean if a field has been set.

### GetVxlanSettings

`func (o *VnicEthAdapterPolicyInventory) GetVxlanSettings() VnicVxlanSettings`

GetVxlanSettings returns the VxlanSettings field if non-nil, zero value otherwise.

### GetVxlanSettingsOk

`func (o *VnicEthAdapterPolicyInventory) GetVxlanSettingsOk() (*VnicVxlanSettings, bool)`

GetVxlanSettingsOk returns a tuple with the VxlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanSettings

`func (o *VnicEthAdapterPolicyInventory) SetVxlanSettings(v VnicVxlanSettings)`

SetVxlanSettings sets VxlanSettings field to given value.

### HasVxlanSettings

`func (o *VnicEthAdapterPolicyInventory) HasVxlanSettings() bool`

HasVxlanSettings returns a boolean if a field has been set.

### SetVxlanSettingsNil

`func (o *VnicEthAdapterPolicyInventory) SetVxlanSettingsNil(b bool)`

 SetVxlanSettingsNil sets the value for VxlanSettings to be an explicit nil

### UnsetVxlanSettings
`func (o *VnicEthAdapterPolicyInventory) UnsetVxlanSettings()`

UnsetVxlanSettings ensures that no value is present for VxlanSettings, not even an explicit nil
### GetTargetMo

`func (o *VnicEthAdapterPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicEthAdapterPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicEthAdapterPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicEthAdapterPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


