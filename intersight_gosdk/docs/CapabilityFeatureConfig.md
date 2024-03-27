# CapabilityFeatureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FeatureConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FeatureConfig"]
**FeatureName** | Pointer to **string** | Name of the feature that identifies the specific adapter configuration. * &#x60;RoCEv2&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2. * &#x60;RoCEv1&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1. * &#x60;VMQ&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) feature. * &#x60;VMMQ&#x60; - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature. * &#x60;VMQInterrupts&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature. * &#x60;NVGRE&#x60; - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature. * &#x60;ARFS&#x60; - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature. * &#x60;VXLAN&#x60; - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. * &#x60;usNIC&#x60; - Capability indicator of the User Space NIC (usNIC) feature. * &#x60;Advanced Filter&#x60; - Capability indicator of the Advanced Filter feature. * &#x60;Azure Stack Host QOS&#x60; - Capability indicator of the Azure Stack Host QOS feature. * &#x60;GENEVE Offload&#x60; - Capability indicator of the Generic Network Virtualization Encapsulation (Geneve) Offload feature. * &#x60;QinQ&#x60; - Capability indicator of the QinQ feature. * &#x60;SRIOV&#x60; - Capability indicator of the Single Root Input Output Virtualization (SR-IOV). | [optional] [default to "RoCEv2"]
**MinAdapterFwVersion** | Pointer to **string** | Firmware version of Adapter from which support for this feature is available. | [optional] 
**MinFwVersion** | Pointer to **string** | Firmware version of BMC from which support for this feature is available. | [optional] 
**SupportedFwVersions** | Pointer to **[]string** |  | [optional] 
**SupportedInAdapters** | Pointer to **[]string** |  | [optional] 
**SupportedInGenerations** | Pointer to **[]int32** |  | [optional] 
**UnsupportedFeatureMatrix** | Pointer to [**[]CapabilityUnsupportedFeatureConfig**](CapabilityUnsupportedFeatureConfig.md) |  | [optional] 
**ValidationAction** | Pointer to **string** | Action to be taken when validation does not succeed. * &#x60;Error&#x60; - Stop workflow execution by throwing error. * &#x60;Skip&#x60; - Remove the feature from configuration and continue workflow execution. | [optional] [default to "Error"]

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

### GetMinAdapterFwVersion

`func (o *CapabilityFeatureConfig) GetMinAdapterFwVersion() string`

GetMinAdapterFwVersion returns the MinAdapterFwVersion field if non-nil, zero value otherwise.

### GetMinAdapterFwVersionOk

`func (o *CapabilityFeatureConfig) GetMinAdapterFwVersionOk() (*string, bool)`

GetMinAdapterFwVersionOk returns a tuple with the MinAdapterFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAdapterFwVersion

`func (o *CapabilityFeatureConfig) SetMinAdapterFwVersion(v string)`

SetMinAdapterFwVersion sets MinAdapterFwVersion field to given value.

### HasMinAdapterFwVersion

`func (o *CapabilityFeatureConfig) HasMinAdapterFwVersion() bool`

HasMinAdapterFwVersion returns a boolean if a field has been set.

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
### GetUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfig) GetUnsupportedFeatureMatrix() []CapabilityUnsupportedFeatureConfig`

GetUnsupportedFeatureMatrix returns the UnsupportedFeatureMatrix field if non-nil, zero value otherwise.

### GetUnsupportedFeatureMatrixOk

`func (o *CapabilityFeatureConfig) GetUnsupportedFeatureMatrixOk() (*[]CapabilityUnsupportedFeatureConfig, bool)`

GetUnsupportedFeatureMatrixOk returns a tuple with the UnsupportedFeatureMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfig) SetUnsupportedFeatureMatrix(v []CapabilityUnsupportedFeatureConfig)`

SetUnsupportedFeatureMatrix sets UnsupportedFeatureMatrix field to given value.

### HasUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfig) HasUnsupportedFeatureMatrix() bool`

HasUnsupportedFeatureMatrix returns a boolean if a field has been set.

### SetUnsupportedFeatureMatrixNil

`func (o *CapabilityFeatureConfig) SetUnsupportedFeatureMatrixNil(b bool)`

 SetUnsupportedFeatureMatrixNil sets the value for UnsupportedFeatureMatrix to be an explicit nil

### UnsetUnsupportedFeatureMatrix
`func (o *CapabilityFeatureConfig) UnsetUnsupportedFeatureMatrix()`

UnsetUnsupportedFeatureMatrix ensures that no value is present for UnsupportedFeatureMatrix, not even an explicit nil
### GetValidationAction

`func (o *CapabilityFeatureConfig) GetValidationAction() string`

GetValidationAction returns the ValidationAction field if non-nil, zero value otherwise.

### GetValidationActionOk

`func (o *CapabilityFeatureConfig) GetValidationActionOk() (*string, bool)`

GetValidationActionOk returns a tuple with the ValidationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationAction

`func (o *CapabilityFeatureConfig) SetValidationAction(v string)`

SetValidationAction sets ValidationAction field to given value.

### HasValidationAction

`func (o *CapabilityFeatureConfig) HasValidationAction() bool`

HasValidationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


