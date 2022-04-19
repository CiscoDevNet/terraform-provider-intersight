# VnicEthAdapterPolicyInventoryAllOf

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

### NewVnicEthAdapterPolicyInventoryAllOf

`func NewVnicEthAdapterPolicyInventoryAllOf(classId string, objectType string, ) *VnicEthAdapterPolicyInventoryAllOf`

NewVnicEthAdapterPolicyInventoryAllOf instantiates a new VnicEthAdapterPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthAdapterPolicyInventoryAllOfWithDefaults

`func NewVnicEthAdapterPolicyInventoryAllOfWithDefaults() *VnicEthAdapterPolicyInventoryAllOf`

NewVnicEthAdapterPolicyInventoryAllOfWithDefaults instantiates a new VnicEthAdapterPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvancedFilter

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetAdvancedFilter() bool`

GetAdvancedFilter returns the AdvancedFilter field if non-nil, zero value otherwise.

### GetAdvancedFilterOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetAdvancedFilterOk() (*bool, bool)`

GetAdvancedFilterOk returns a tuple with the AdvancedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFilter

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetAdvancedFilter(v bool)`

SetAdvancedFilter sets AdvancedFilter field to given value.

### HasAdvancedFilter

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasAdvancedFilter() bool`

HasAdvancedFilter returns a boolean if a field has been set.

### GetArfsSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetArfsSettings() VnicArfsSettings`

GetArfsSettings returns the ArfsSettings field if non-nil, zero value otherwise.

### GetArfsSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetArfsSettingsOk() (*VnicArfsSettings, bool)`

GetArfsSettingsOk returns a tuple with the ArfsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArfsSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetArfsSettings(v VnicArfsSettings)`

SetArfsSettings sets ArfsSettings field to given value.

### HasArfsSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasArfsSettings() bool`

HasArfsSettings returns a boolean if a field has been set.

### SetArfsSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetArfsSettingsNil(b bool)`

 SetArfsSettingsNil sets the value for ArfsSettings to be an explicit nil

### UnsetArfsSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetArfsSettings()`

UnsetArfsSettings ensures that no value is present for ArfsSettings, not even an explicit nil
### GetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetCompletionQueueSettings() VnicCompletionQueueSettings`

GetCompletionQueueSettings returns the CompletionQueueSettings field if non-nil, zero value otherwise.

### GetCompletionQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool)`

GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetCompletionQueueSettings(v VnicCompletionQueueSettings)`

SetCompletionQueueSettings sets CompletionQueueSettings field to given value.

### HasCompletionQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasCompletionQueueSettings() bool`

HasCompletionQueueSettings returns a boolean if a field has been set.

### SetCompletionQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetCompletionQueueSettingsNil(b bool)`

 SetCompletionQueueSettingsNil sets the value for CompletionQueueSettings to be an explicit nil

### UnsetCompletionQueueSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetCompletionQueueSettings()`

UnsetCompletionQueueSettings ensures that no value is present for CompletionQueueSettings, not even an explicit nil
### GetGeneveEnabled

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetGeneveEnabled() bool`

GetGeneveEnabled returns the GeneveEnabled field if non-nil, zero value otherwise.

### GetGeneveEnabledOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetGeneveEnabledOk() (*bool, bool)`

GetGeneveEnabledOk returns a tuple with the GeneveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneveEnabled

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetGeneveEnabled(v bool)`

SetGeneveEnabled sets GeneveEnabled field to given value.

### HasGeneveEnabled

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasGeneveEnabled() bool`

HasGeneveEnabled returns a boolean if a field has been set.

### GetInterruptScaling

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetInterruptScaling() bool`

GetInterruptScaling returns the InterruptScaling field if non-nil, zero value otherwise.

### GetInterruptScalingOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetInterruptScalingOk() (*bool, bool)`

GetInterruptScalingOk returns a tuple with the InterruptScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptScaling

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetInterruptScaling(v bool)`

SetInterruptScaling sets InterruptScaling field to given value.

### HasInterruptScaling

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasInterruptScaling() bool`

HasInterruptScaling returns a boolean if a field has been set.

### GetInterruptSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetInterruptSettings() VnicEthInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetInterruptSettings(v VnicEthInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### SetInterruptSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetInterruptSettingsNil(b bool)`

 SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil

### UnsetInterruptSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetInterruptSettings()`

UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
### GetNvgreSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetNvgreSettings() VnicNvgreSettings`

GetNvgreSettings returns the NvgreSettings field if non-nil, zero value otherwise.

### GetNvgreSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetNvgreSettingsOk() (*VnicNvgreSettings, bool)`

GetNvgreSettingsOk returns a tuple with the NvgreSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvgreSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetNvgreSettings(v VnicNvgreSettings)`

SetNvgreSettings sets NvgreSettings field to given value.

### HasNvgreSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasNvgreSettings() bool`

HasNvgreSettings returns a boolean if a field has been set.

### SetNvgreSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetNvgreSettingsNil(b bool)`

 SetNvgreSettingsNil sets the value for NvgreSettings to be an explicit nil

### UnsetNvgreSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetNvgreSettings()`

UnsetNvgreSettings ensures that no value is present for NvgreSettings, not even an explicit nil
### GetPtpSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetPtpSettings() VnicPtpSettings`

GetPtpSettings returns the PtpSettings field if non-nil, zero value otherwise.

### GetPtpSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetPtpSettingsOk() (*VnicPtpSettings, bool)`

GetPtpSettingsOk returns a tuple with the PtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetPtpSettings(v VnicPtpSettings)`

SetPtpSettings sets PtpSettings field to given value.

### HasPtpSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasPtpSettings() bool`

HasPtpSettings returns a boolean if a field has been set.

### SetPtpSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetPtpSettingsNil(b bool)`

 SetPtpSettingsNil sets the value for PtpSettings to be an explicit nil

### UnsetPtpSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetPtpSettings()`

UnsetPtpSettings ensures that no value is present for PtpSettings, not even an explicit nil
### GetRoceSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRoceSettings() VnicRoceSettings`

GetRoceSettings returns the RoceSettings field if non-nil, zero value otherwise.

### GetRoceSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRoceSettingsOk() (*VnicRoceSettings, bool)`

GetRoceSettingsOk returns a tuple with the RoceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoceSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRoceSettings(v VnicRoceSettings)`

SetRoceSettings sets RoceSettings field to given value.

### HasRoceSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasRoceSettings() bool`

HasRoceSettings returns a boolean if a field has been set.

### SetRoceSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRoceSettingsNil(b bool)`

 SetRoceSettingsNil sets the value for RoceSettings to be an explicit nil

### UnsetRoceSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetRoceSettings()`

UnsetRoceSettings ensures that no value is present for RoceSettings, not even an explicit nil
### GetRssHashSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRssHashSettings() VnicRssHashSettings`

GetRssHashSettings returns the RssHashSettings field if non-nil, zero value otherwise.

### GetRssHashSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRssHashSettingsOk() (*VnicRssHashSettings, bool)`

GetRssHashSettingsOk returns a tuple with the RssHashSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssHashSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRssHashSettings(v VnicRssHashSettings)`

SetRssHashSettings sets RssHashSettings field to given value.

### HasRssHashSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasRssHashSettings() bool`

HasRssHashSettings returns a boolean if a field has been set.

### SetRssHashSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRssHashSettingsNil(b bool)`

 SetRssHashSettingsNil sets the value for RssHashSettings to be an explicit nil

### UnsetRssHashSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetRssHashSettings()`

UnsetRssHashSettings ensures that no value is present for RssHashSettings, not even an explicit nil
### GetRssSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRssSettings() bool`

GetRssSettings returns the RssSettings field if non-nil, zero value otherwise.

### GetRssSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRssSettingsOk() (*bool, bool)`

GetRssSettingsOk returns a tuple with the RssSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRssSettings(v bool)`

SetRssSettings sets RssSettings field to given value.

### HasRssSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasRssSettings() bool`

HasRssSettings returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRxQueueSettings() VnicEthRxQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRxQueueSettings(v VnicEthRxQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### SetRxQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetRxQueueSettingsNil(b bool)`

 SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil

### UnsetRxQueueSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetRxQueueSettings()`

UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
### GetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTcpOffloadSettings() VnicTcpOffloadSettings`

GetTcpOffloadSettings returns the TcpOffloadSettings field if non-nil, zero value otherwise.

### GetTcpOffloadSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool)`

GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetTcpOffloadSettings(v VnicTcpOffloadSettings)`

SetTcpOffloadSettings sets TcpOffloadSettings field to given value.

### HasTcpOffloadSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasTcpOffloadSettings() bool`

HasTcpOffloadSettings returns a boolean if a field has been set.

### SetTcpOffloadSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetTcpOffloadSettingsNil(b bool)`

 SetTcpOffloadSettingsNil sets the value for TcpOffloadSettings to be an explicit nil

### UnsetTcpOffloadSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetTcpOffloadSettings()`

UnsetTcpOffloadSettings ensures that no value is present for TcpOffloadSettings, not even an explicit nil
### GetTxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTxQueueSettings() VnicEthTxQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetTxQueueSettings(v VnicEthTxQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### SetTxQueueSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetTxQueueSettingsNil(b bool)`

 SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil

### UnsetTxQueueSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetTxQueueSettings()`

UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
### GetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetUplinkFailbackTimeout() int64`

GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field if non-nil, zero value otherwise.

### GetUplinkFailbackTimeoutOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetUplinkFailbackTimeoutOk() (*int64, bool)`

GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetUplinkFailbackTimeout(v int64)`

SetUplinkFailbackTimeout sets UplinkFailbackTimeout field to given value.

### HasUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasUplinkFailbackTimeout() bool`

HasUplinkFailbackTimeout returns a boolean if a field has been set.

### GetVxlanSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetVxlanSettings() VnicVxlanSettings`

GetVxlanSettings returns the VxlanSettings field if non-nil, zero value otherwise.

### GetVxlanSettingsOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetVxlanSettingsOk() (*VnicVxlanSettings, bool)`

GetVxlanSettingsOk returns a tuple with the VxlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetVxlanSettings(v VnicVxlanSettings)`

SetVxlanSettings sets VxlanSettings field to given value.

### HasVxlanSettings

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasVxlanSettings() bool`

HasVxlanSettings returns a boolean if a field has been set.

### SetVxlanSettingsNil

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetVxlanSettingsNil(b bool)`

 SetVxlanSettingsNil sets the value for VxlanSettings to be an explicit nil

### UnsetVxlanSettings
`func (o *VnicEthAdapterPolicyInventoryAllOf) UnsetVxlanSettings()`

UnsetVxlanSettings ensures that no value is present for VxlanSettings, not even an explicit nil
### GetTargetMo

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicEthAdapterPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicEthAdapterPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicEthAdapterPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


