# CapabilityFeatureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FeatureConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FeatureConfig"]
**FeatureName** | Pointer to **string** | Name of the feature that identifies the specific adapter configuration. * &#x60;RoCEv2&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2. * &#x60;RoCEv1&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1. * &#x60;VMQ&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) feature. * &#x60;VMMQ&#x60; - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature. * &#x60;VMQInterrupts&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature. * &#x60;NVGRE&#x60; - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature. * &#x60;ARFS&#x60; - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature. * &#x60;VXLAN&#x60; - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. | [optional] [default to "RoCEv2"]
**MinFwVersion** | Pointer to **string** | Firmware version from which support for this feature is available. | [optional] 
**SupportedFwVersions** | Pointer to **[]string** |  | [optional] 
**SupportedInAdapters** | Pointer to **[]string** |  | [optional] 
**SupportedInGenerations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCapabilityFeatureConfig

`func NewCapabilityFeatureConfig(classId string, objectType string, ) *CapabilityFeatureConfig`

NewCapabilityFeatureConfig instantiates a new CapabilityFeatureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFeatureConfigWithDefaults

`func NewCapabilityFeatureConfigWithDefaults() *CapabilityFeatureConfig`

NewCapabilityFeatureConfigWithDefaults instantiates a new CapabilityFeatureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFeatureConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFeatureConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFeatureConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFeatureConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFeatureConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFeatureConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureName

`func (o *CapabilityFeatureConfig) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *CapabilityFeatureConfig) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *CapabilityFeatureConfig) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *CapabilityFeatureConfig) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetMinFwVersion

`func (o *CapabilityFeatureConfig) GetMinFwVersion() string`

GetMinFwVersion returns the MinFwVersion field if non-nil, zero value otherwise.

### GetMinFwVersionOk

`func (o *CapabilityFeatureConfig) GetMinFwVersionOk() (*string, bool)`

GetMinFwVersionOk returns a tuple with the MinFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFwVersion

`func (o *CapabilityFeatureConfig) SetMinFwVersion(v string)`

SetMinFwVersion sets MinFwVersion field to given value.

### HasMinFwVersion

`func (o *CapabilityFeatureConfig) HasMinFwVersion() bool`

HasMinFwVersion returns a boolean if a field has been set.

### GetSupportedFwVersions

`func (o *CapabilityFeatureConfig) GetSupportedFwVersions() []string`

GetSupportedFwVersions returns the SupportedFwVersions field if non-nil, zero value otherwise.

### GetSupportedFwVersionsOk

`func (o *CapabilityFeatureConfig) GetSupportedFwVersionsOk() (*[]string, bool)`

GetSupportedFwVersionsOk returns a tuple with the SupportedFwVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFwVersions

`func (o *CapabilityFeatureConfig) SetSupportedFwVersions(v []string)`

SetSupportedFwVersions sets SupportedFwVersions field to given value.

### HasSupportedFwVersions

`func (o *CapabilityFeatureConfig) HasSupportedFwVersions() bool`

HasSupportedFwVersions returns a boolean if a field has been set.

### SetSupportedFwVersionsNil

`func (o *CapabilityFeatureConfig) SetSupportedFwVersionsNil(b bool)`

 SetSupportedFwVersionsNil sets the value for SupportedFwVersions to be an explicit nil

### UnsetSupportedFwVersions
`func (o *CapabilityFeatureConfig) UnsetSupportedFwVersions()`

UnsetSupportedFwVersions ensures that no value is present for SupportedFwVersions, not even an explicit nil
### GetSupportedInAdapters

`func (o *CapabilityFeatureConfig) GetSupportedInAdapters() []string`

GetSupportedInAdapters returns the SupportedInAdapters field if non-nil, zero value otherwise.

### GetSupportedInAdaptersOk

`func (o *CapabilityFeatureConfig) GetSupportedInAdaptersOk() (*[]string, bool)`

GetSupportedInAdaptersOk returns a tuple with the SupportedInAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInAdapters

`func (o *CapabilityFeatureConfig) SetSupportedInAdapters(v []string)`

SetSupportedInAdapters sets SupportedInAdapters field to given value.

### HasSupportedInAdapters

`func (o *CapabilityFeatureConfig) HasSupportedInAdapters() bool`

HasSupportedInAdapters returns a boolean if a field has been set.

### SetSupportedInAdaptersNil

`func (o *CapabilityFeatureConfig) SetSupportedInAdaptersNil(b bool)`

 SetSupportedInAdaptersNil sets the value for SupportedInAdapters to be an explicit nil

### UnsetSupportedInAdapters
`func (o *CapabilityFeatureConfig) UnsetSupportedInAdapters()`

UnsetSupportedInAdapters ensures that no value is present for SupportedInAdapters, not even an explicit nil
### GetSupportedInGenerations

`func (o *CapabilityFeatureConfig) GetSupportedInGenerations() []int32`

GetSupportedInGenerations returns the SupportedInGenerations field if non-nil, zero value otherwise.

### GetSupportedInGenerationsOk

`func (o *CapabilityFeatureConfig) GetSupportedInGenerationsOk() (*[]int32, bool)`

GetSupportedInGenerationsOk returns a tuple with the SupportedInGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInGenerations

`func (o *CapabilityFeatureConfig) SetSupportedInGenerations(v []int32)`

SetSupportedInGenerations sets SupportedInGenerations field to given value.

### HasSupportedInGenerations

`func (o *CapabilityFeatureConfig) HasSupportedInGenerations() bool`

HasSupportedInGenerations returns a boolean if a field has been set.

### SetSupportedInGenerationsNil

`func (o *CapabilityFeatureConfig) SetSupportedInGenerationsNil(b bool)`

 SetSupportedInGenerationsNil sets the value for SupportedInGenerations to be an explicit nil

### UnsetSupportedInGenerations
`func (o *CapabilityFeatureConfig) UnsetSupportedInGenerations()`

UnsetSupportedInGenerations ensures that no value is present for SupportedInGenerations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


