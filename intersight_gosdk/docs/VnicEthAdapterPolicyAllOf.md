# VnicEthAdapterPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedFilter** | Pointer to **bool** | Enables advanced filtering on the interface. | [optional] 
**ArfsSettings** | Pointer to [**VnicArfsSettings**](vnic.ArfsSettings.md) |  | [optional] 
**CompletionQueueSettings** | Pointer to [**VnicCompletionQueueSettings**](vnic.CompletionQueueSettings.md) |  | [optional] 
**InterruptScaling** | Pointer to **bool** | Enables Interrupt Scaling on the interface. | [optional] 
**InterruptSettings** | Pointer to [**VnicEthInterruptSettings**](vnic.EthInterruptSettings.md) |  | [optional] 
**NvgreSettings** | Pointer to [**VnicNvgreSettings**](vnic.NvgreSettings.md) |  | [optional] 
**RoceSettings** | Pointer to [**VnicRoceSettings**](vnic.RoceSettings.md) |  | [optional] 
**RssHashSettings** | Pointer to [**VnicRssHashSettings**](vnic.RssHashSettings.md) |  | [optional] 
**RssSettings** | Pointer to **bool** | Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. | [optional] 
**RxQueueSettings** | Pointer to [**VnicEthRxQueueSettings**](vnic.EthRxQueueSettings.md) |  | [optional] 
**TcpOffloadSettings** | Pointer to [**VnicTcpOffloadSettings**](vnic.TcpOffloadSettings.md) |  | [optional] 
**TxQueueSettings** | Pointer to [**VnicEthTxQueueSettings**](vnic.EthTxQueueSettings.md) |  | [optional] 
**UplinkFailbackTimeout** | Pointer to **int64** | Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC. | [optional] 
**VxlanSettings** | Pointer to [**VnicVxlanSettings**](vnic.VxlanSettings.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewVnicEthAdapterPolicyAllOf

`func NewVnicEthAdapterPolicyAllOf() *VnicEthAdapterPolicyAllOf`

NewVnicEthAdapterPolicyAllOf instantiates a new VnicEthAdapterPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthAdapterPolicyAllOfWithDefaults

`func NewVnicEthAdapterPolicyAllOfWithDefaults() *VnicEthAdapterPolicyAllOf`

NewVnicEthAdapterPolicyAllOfWithDefaults instantiates a new VnicEthAdapterPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedFilter

`func (o *VnicEthAdapterPolicyAllOf) GetAdvancedFilter() bool`

GetAdvancedFilter returns the AdvancedFilter field if non-nil, zero value otherwise.

### GetAdvancedFilterOk

`func (o *VnicEthAdapterPolicyAllOf) GetAdvancedFilterOk() (*bool, bool)`

GetAdvancedFilterOk returns a tuple with the AdvancedFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFilter

`func (o *VnicEthAdapterPolicyAllOf) SetAdvancedFilter(v bool)`

SetAdvancedFilter sets AdvancedFilter field to given value.

### HasAdvancedFilter

`func (o *VnicEthAdapterPolicyAllOf) HasAdvancedFilter() bool`

HasAdvancedFilter returns a boolean if a field has been set.

### GetArfsSettings

`func (o *VnicEthAdapterPolicyAllOf) GetArfsSettings() VnicArfsSettings`

GetArfsSettings returns the ArfsSettings field if non-nil, zero value otherwise.

### GetArfsSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetArfsSettingsOk() (*VnicArfsSettings, bool)`

GetArfsSettingsOk returns a tuple with the ArfsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArfsSettings

`func (o *VnicEthAdapterPolicyAllOf) SetArfsSettings(v VnicArfsSettings)`

SetArfsSettings sets ArfsSettings field to given value.

### HasArfsSettings

`func (o *VnicEthAdapterPolicyAllOf) HasArfsSettings() bool`

HasArfsSettings returns a boolean if a field has been set.

### GetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) GetCompletionQueueSettings() VnicCompletionQueueSettings`

GetCompletionQueueSettings returns the CompletionQueueSettings field if non-nil, zero value otherwise.

### GetCompletionQueueSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool)`

GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) SetCompletionQueueSettings(v VnicCompletionQueueSettings)`

SetCompletionQueueSettings sets CompletionQueueSettings field to given value.

### HasCompletionQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) HasCompletionQueueSettings() bool`

HasCompletionQueueSettings returns a boolean if a field has been set.

### GetInterruptScaling

`func (o *VnicEthAdapterPolicyAllOf) GetInterruptScaling() bool`

GetInterruptScaling returns the InterruptScaling field if non-nil, zero value otherwise.

### GetInterruptScalingOk

`func (o *VnicEthAdapterPolicyAllOf) GetInterruptScalingOk() (*bool, bool)`

GetInterruptScalingOk returns a tuple with the InterruptScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptScaling

`func (o *VnicEthAdapterPolicyAllOf) SetInterruptScaling(v bool)`

SetInterruptScaling sets InterruptScaling field to given value.

### HasInterruptScaling

`func (o *VnicEthAdapterPolicyAllOf) HasInterruptScaling() bool`

HasInterruptScaling returns a boolean if a field has been set.

### GetInterruptSettings

`func (o *VnicEthAdapterPolicyAllOf) GetInterruptSettings() VnicEthInterruptSettings`

GetInterruptSettings returns the InterruptSettings field if non-nil, zero value otherwise.

### GetInterruptSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool)`

GetInterruptSettingsOk returns a tuple with the InterruptSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptSettings

`func (o *VnicEthAdapterPolicyAllOf) SetInterruptSettings(v VnicEthInterruptSettings)`

SetInterruptSettings sets InterruptSettings field to given value.

### HasInterruptSettings

`func (o *VnicEthAdapterPolicyAllOf) HasInterruptSettings() bool`

HasInterruptSettings returns a boolean if a field has been set.

### GetNvgreSettings

`func (o *VnicEthAdapterPolicyAllOf) GetNvgreSettings() VnicNvgreSettings`

GetNvgreSettings returns the NvgreSettings field if non-nil, zero value otherwise.

### GetNvgreSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetNvgreSettingsOk() (*VnicNvgreSettings, bool)`

GetNvgreSettingsOk returns a tuple with the NvgreSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvgreSettings

`func (o *VnicEthAdapterPolicyAllOf) SetNvgreSettings(v VnicNvgreSettings)`

SetNvgreSettings sets NvgreSettings field to given value.

### HasNvgreSettings

`func (o *VnicEthAdapterPolicyAllOf) HasNvgreSettings() bool`

HasNvgreSettings returns a boolean if a field has been set.

### GetRoceSettings

`func (o *VnicEthAdapterPolicyAllOf) GetRoceSettings() VnicRoceSettings`

GetRoceSettings returns the RoceSettings field if non-nil, zero value otherwise.

### GetRoceSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetRoceSettingsOk() (*VnicRoceSettings, bool)`

GetRoceSettingsOk returns a tuple with the RoceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoceSettings

`func (o *VnicEthAdapterPolicyAllOf) SetRoceSettings(v VnicRoceSettings)`

SetRoceSettings sets RoceSettings field to given value.

### HasRoceSettings

`func (o *VnicEthAdapterPolicyAllOf) HasRoceSettings() bool`

HasRoceSettings returns a boolean if a field has been set.

### GetRssHashSettings

`func (o *VnicEthAdapterPolicyAllOf) GetRssHashSettings() VnicRssHashSettings`

GetRssHashSettings returns the RssHashSettings field if non-nil, zero value otherwise.

### GetRssHashSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetRssHashSettingsOk() (*VnicRssHashSettings, bool)`

GetRssHashSettingsOk returns a tuple with the RssHashSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssHashSettings

`func (o *VnicEthAdapterPolicyAllOf) SetRssHashSettings(v VnicRssHashSettings)`

SetRssHashSettings sets RssHashSettings field to given value.

### HasRssHashSettings

`func (o *VnicEthAdapterPolicyAllOf) HasRssHashSettings() bool`

HasRssHashSettings returns a boolean if a field has been set.

### GetRssSettings

`func (o *VnicEthAdapterPolicyAllOf) GetRssSettings() bool`

GetRssSettings returns the RssSettings field if non-nil, zero value otherwise.

### GetRssSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetRssSettingsOk() (*bool, bool)`

GetRssSettingsOk returns a tuple with the RssSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSettings

`func (o *VnicEthAdapterPolicyAllOf) SetRssSettings(v bool)`

SetRssSettings sets RssSettings field to given value.

### HasRssSettings

`func (o *VnicEthAdapterPolicyAllOf) HasRssSettings() bool`

HasRssSettings returns a boolean if a field has been set.

### GetRxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) GetRxQueueSettings() VnicEthRxQueueSettings`

GetRxQueueSettings returns the RxQueueSettings field if non-nil, zero value otherwise.

### GetRxQueueSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool)`

GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) SetRxQueueSettings(v VnicEthRxQueueSettings)`

SetRxQueueSettings sets RxQueueSettings field to given value.

### HasRxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) HasRxQueueSettings() bool`

HasRxQueueSettings returns a boolean if a field has been set.

### GetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyAllOf) GetTcpOffloadSettings() VnicTcpOffloadSettings`

GetTcpOffloadSettings returns the TcpOffloadSettings field if non-nil, zero value otherwise.

### GetTcpOffloadSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool)`

GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpOffloadSettings

`func (o *VnicEthAdapterPolicyAllOf) SetTcpOffloadSettings(v VnicTcpOffloadSettings)`

SetTcpOffloadSettings sets TcpOffloadSettings field to given value.

### HasTcpOffloadSettings

`func (o *VnicEthAdapterPolicyAllOf) HasTcpOffloadSettings() bool`

HasTcpOffloadSettings returns a boolean if a field has been set.

### GetTxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) GetTxQueueSettings() VnicEthTxQueueSettings`

GetTxQueueSettings returns the TxQueueSettings field if non-nil, zero value otherwise.

### GetTxQueueSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool)`

GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) SetTxQueueSettings(v VnicEthTxQueueSettings)`

SetTxQueueSettings sets TxQueueSettings field to given value.

### HasTxQueueSettings

`func (o *VnicEthAdapterPolicyAllOf) HasTxQueueSettings() bool`

HasTxQueueSettings returns a boolean if a field has been set.

### GetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyAllOf) GetUplinkFailbackTimeout() int64`

GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field if non-nil, zero value otherwise.

### GetUplinkFailbackTimeoutOk

`func (o *VnicEthAdapterPolicyAllOf) GetUplinkFailbackTimeoutOk() (*int64, bool)`

GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyAllOf) SetUplinkFailbackTimeout(v int64)`

SetUplinkFailbackTimeout sets UplinkFailbackTimeout field to given value.

### HasUplinkFailbackTimeout

`func (o *VnicEthAdapterPolicyAllOf) HasUplinkFailbackTimeout() bool`

HasUplinkFailbackTimeout returns a boolean if a field has been set.

### GetVxlanSettings

`func (o *VnicEthAdapterPolicyAllOf) GetVxlanSettings() VnicVxlanSettings`

GetVxlanSettings returns the VxlanSettings field if non-nil, zero value otherwise.

### GetVxlanSettingsOk

`func (o *VnicEthAdapterPolicyAllOf) GetVxlanSettingsOk() (*VnicVxlanSettings, bool)`

GetVxlanSettingsOk returns a tuple with the VxlanSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanSettings

`func (o *VnicEthAdapterPolicyAllOf) SetVxlanSettings(v VnicVxlanSettings)`

SetVxlanSettings sets VxlanSettings field to given value.

### HasVxlanSettings

`func (o *VnicEthAdapterPolicyAllOf) HasVxlanSettings() bool`

HasVxlanSettings returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicEthAdapterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicEthAdapterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicEthAdapterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicEthAdapterPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


