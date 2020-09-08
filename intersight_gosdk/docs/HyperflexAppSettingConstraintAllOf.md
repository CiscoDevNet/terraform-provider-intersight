# HyperflexAppSettingConstraintAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HxdpVersion** | Pointer to **string** | The supported HyperFlex Data Platform version in regex format. | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the HyperFlex cluster. * &#x60;ESXi&#x60; - ESXi hypervisor as specified by the user. * &#x60;HYPERV&#x60; - Hyperv hypervisor as specified by the user. * &#x60;KVM&#x60; - KVM hypervisor as specified by the user. | [optional] [default to "ESXi"]
**MgmtPlatform** | Pointer to **string** | The supported management platform for the HyperFlex Cluster. * &#x60;FI&#x60; - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * &#x60;EDGE&#x60; - The host servers used in the cluster deployment are standalone severs. | [optional] [default to "FI"]
**ServerModel** | Pointer to **string** | The supported server models in regex format. | [optional] 

## Methods

### NewHyperflexAppSettingConstraintAllOf

`func NewHyperflexAppSettingConstraintAllOf() *HyperflexAppSettingConstraintAllOf`

NewHyperflexAppSettingConstraintAllOf instantiates a new HyperflexAppSettingConstraintAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppSettingConstraintAllOfWithDefaults

`func NewHyperflexAppSettingConstraintAllOfWithDefaults() *HyperflexAppSettingConstraintAllOf`

NewHyperflexAppSettingConstraintAllOfWithDefaults instantiates a new HyperflexAppSettingConstraintAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHxdpVersion

`func (o *HyperflexAppSettingConstraintAllOf) GetHxdpVersion() string`

GetHxdpVersion returns the HxdpVersion field if non-nil, zero value otherwise.

### GetHxdpVersionOk

`func (o *HyperflexAppSettingConstraintAllOf) GetHxdpVersionOk() (*string, bool)`

GetHxdpVersionOk returns a tuple with the HxdpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersion

`func (o *HyperflexAppSettingConstraintAllOf) SetHxdpVersion(v string)`

SetHxdpVersion sets HxdpVersion field to given value.

### HasHxdpVersion

`func (o *HyperflexAppSettingConstraintAllOf) HasHxdpVersion() bool`

HasHxdpVersion returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HyperflexAppSettingConstraintAllOf) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexAppSettingConstraintAllOf) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexAppSettingConstraintAllOf) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexAppSettingConstraintAllOf) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *HyperflexAppSettingConstraintAllOf) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *HyperflexAppSettingConstraintAllOf) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *HyperflexAppSettingConstraintAllOf) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *HyperflexAppSettingConstraintAllOf) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetServerModel

`func (o *HyperflexAppSettingConstraintAllOf) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HyperflexAppSettingConstraintAllOf) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HyperflexAppSettingConstraintAllOf) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HyperflexAppSettingConstraintAllOf) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


