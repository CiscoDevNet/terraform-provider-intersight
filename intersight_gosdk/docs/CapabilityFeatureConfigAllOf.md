# CapabilityFeatureConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FeatureConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FeatureConfig"]
**FeatureName** | Pointer to **string** | Name of the feature that identifies the specific adapter configuration. * &#x60;RoCEv2&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 2. * &#x60;RoCEv1&#x60; - Capability indicator of the RDMA over Converged Ethernet (RoCE) feature version 1. * &#x60;VMQ&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) feature. * &#x60;VMMQ&#x60; - Capability indicator of the Virtual Machine Multi-Queue (VMMQ) feature. * &#x60;VMQInterrupts&#x60; - Capability indicator of the Virtual Machine Queue (VMQ) Interrupts feature. * &#x60;NVGRE&#x60; - Capability indicator of the Network Virtualization using Generic Routing Encapsulation (NVGRE) feature. * &#x60;ARFS&#x60; - Capability indicator of the Accelerated Receive Flow Steering (ARFS) feature. * &#x60;VXLAN&#x60; - Capability indicator of the Virtual Extensible LAN (VXLAN) feature. * &#x60;usNIC&#x60; - Capability indicator of the User Space NIC (usNIC) feature. * &#x60;Advanced Filter&#x60; - Capability indicator of the Advanced Filter feature. * &#x60;Azure Stack Host QOS&#x60; - Capability indicator of the Azure Stack Host QOS feature. * &#x60;QinQ&#x60; - Capability indicator of the QinQ feature. * &#x60;SRIOV&#x60; - Capability indicator of the Single Root Input Output Virtualization (SR-IOV). | [optional] [default to "RoCEv2"]
**MinAdapterFwVersion** | Pointer to **string** | Firmware version of Adapter from which support for this feature is available. | [optional] 
**MinFwVersion** | Pointer to **string** | Firmware version of BMC from which support for this feature is available. | [optional] 
**SupportedFwVersions** | Pointer to **[]string** |  | [optional] 
**SupportedInAdapters** | Pointer to **[]string** |  | [optional] 
**SupportedInGenerations** | Pointer to **[]int32** |  | [optional] 
**UnsupportedFeatureMatrix** | Pointer to [**[]CapabilityUnsupportedFeatureConfig**](CapabilityUnsupportedFeatureConfig.md) |  | [optional] 
**ValidationAction** | Pointer to **string** | Action to be taken when validation does not succeed. * &#x60;Error&#x60; - Stop workflow execution by throwing error. * &#x60;Skip&#x60; - Remove the feature from configuration and continue workflow execution. | [optional] [default to "Error"]

## Methods

### NewCapabilityFeatureConfigAllOf

`func NewCapabilityFeatureConfigAllOf(classId string, objectType string, ) *CapabilityFeatureConfigAllOf`

NewCapabilityFeatureConfigAllOf instantiates a new CapabilityFeatureConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFeatureConfigAllOfWithDefaults

`func NewCapabilityFeatureConfigAllOfWithDefaults() *CapabilityFeatureConfigAllOf`

NewCapabilityFeatureConfigAllOfWithDefaults instantiates a new CapabilityFeatureConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFeatureConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFeatureConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFeatureConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFeatureConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFeatureConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFeatureConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureName

`func (o *CapabilityFeatureConfigAllOf) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *CapabilityFeatureConfigAllOf) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *CapabilityFeatureConfigAllOf) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *CapabilityFeatureConfigAllOf) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetMinAdapterFwVersion

`func (o *CapabilityFeatureConfigAllOf) GetMinAdapterFwVersion() string`

GetMinAdapterFwVersion returns the MinAdapterFwVersion field if non-nil, zero value otherwise.

### GetMinAdapterFwVersionOk

`func (o *CapabilityFeatureConfigAllOf) GetMinAdapterFwVersionOk() (*string, bool)`

GetMinAdapterFwVersionOk returns a tuple with the MinAdapterFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAdapterFwVersion

`func (o *CapabilityFeatureConfigAllOf) SetMinAdapterFwVersion(v string)`

SetMinAdapterFwVersion sets MinAdapterFwVersion field to given value.

### HasMinAdapterFwVersion

`func (o *CapabilityFeatureConfigAllOf) HasMinAdapterFwVersion() bool`

HasMinAdapterFwVersion returns a boolean if a field has been set.

### GetMinFwVersion

`func (o *CapabilityFeatureConfigAllOf) GetMinFwVersion() string`

GetMinFwVersion returns the MinFwVersion field if non-nil, zero value otherwise.

### GetMinFwVersionOk

`func (o *CapabilityFeatureConfigAllOf) GetMinFwVersionOk() (*string, bool)`

GetMinFwVersionOk returns a tuple with the MinFwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFwVersion

`func (o *CapabilityFeatureConfigAllOf) SetMinFwVersion(v string)`

SetMinFwVersion sets MinFwVersion field to given value.

### HasMinFwVersion

`func (o *CapabilityFeatureConfigAllOf) HasMinFwVersion() bool`

HasMinFwVersion returns a boolean if a field has been set.

### GetSupportedFwVersions

`func (o *CapabilityFeatureConfigAllOf) GetSupportedFwVersions() []string`

GetSupportedFwVersions returns the SupportedFwVersions field if non-nil, zero value otherwise.

### GetSupportedFwVersionsOk

`func (o *CapabilityFeatureConfigAllOf) GetSupportedFwVersionsOk() (*[]string, bool)`

GetSupportedFwVersionsOk returns a tuple with the SupportedFwVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFwVersions

`func (o *CapabilityFeatureConfigAllOf) SetSupportedFwVersions(v []string)`

SetSupportedFwVersions sets SupportedFwVersions field to given value.

### HasSupportedFwVersions

`func (o *CapabilityFeatureConfigAllOf) HasSupportedFwVersions() bool`

HasSupportedFwVersions returns a boolean if a field has been set.

### SetSupportedFwVersionsNil

`func (o *CapabilityFeatureConfigAllOf) SetSupportedFwVersionsNil(b bool)`

 SetSupportedFwVersionsNil sets the value for SupportedFwVersions to be an explicit nil

### UnsetSupportedFwVersions
`func (o *CapabilityFeatureConfigAllOf) UnsetSupportedFwVersions()`

UnsetSupportedFwVersions ensures that no value is present for SupportedFwVersions, not even an explicit nil
### GetSupportedInAdapters

`func (o *CapabilityFeatureConfigAllOf) GetSupportedInAdapters() []string`

GetSupportedInAdapters returns the SupportedInAdapters field if non-nil, zero value otherwise.

### GetSupportedInAdaptersOk

`func (o *CapabilityFeatureConfigAllOf) GetSupportedInAdaptersOk() (*[]string, bool)`

GetSupportedInAdaptersOk returns a tuple with the SupportedInAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInAdapters

`func (o *CapabilityFeatureConfigAllOf) SetSupportedInAdapters(v []string)`

SetSupportedInAdapters sets SupportedInAdapters field to given value.

### HasSupportedInAdapters

`func (o *CapabilityFeatureConfigAllOf) HasSupportedInAdapters() bool`

HasSupportedInAdapters returns a boolean if a field has been set.

### SetSupportedInAdaptersNil

`func (o *CapabilityFeatureConfigAllOf) SetSupportedInAdaptersNil(b bool)`

 SetSupportedInAdaptersNil sets the value for SupportedInAdapters to be an explicit nil

### UnsetSupportedInAdapters
`func (o *CapabilityFeatureConfigAllOf) UnsetSupportedInAdapters()`

UnsetSupportedInAdapters ensures that no value is present for SupportedInAdapters, not even an explicit nil
### GetSupportedInGenerations

`func (o *CapabilityFeatureConfigAllOf) GetSupportedInGenerations() []int32`

GetSupportedInGenerations returns the SupportedInGenerations field if non-nil, zero value otherwise.

### GetSupportedInGenerationsOk

`func (o *CapabilityFeatureConfigAllOf) GetSupportedInGenerationsOk() (*[]int32, bool)`

GetSupportedInGenerationsOk returns a tuple with the SupportedInGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInGenerations

`func (o *CapabilityFeatureConfigAllOf) SetSupportedInGenerations(v []int32)`

SetSupportedInGenerations sets SupportedInGenerations field to given value.

### HasSupportedInGenerations

`func (o *CapabilityFeatureConfigAllOf) HasSupportedInGenerations() bool`

HasSupportedInGenerations returns a boolean if a field has been set.

### SetSupportedInGenerationsNil

`func (o *CapabilityFeatureConfigAllOf) SetSupportedInGenerationsNil(b bool)`

 SetSupportedInGenerationsNil sets the value for SupportedInGenerations to be an explicit nil

### UnsetSupportedInGenerations
`func (o *CapabilityFeatureConfigAllOf) UnsetSupportedInGenerations()`

UnsetSupportedInGenerations ensures that no value is present for SupportedInGenerations, not even an explicit nil
### GetUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfigAllOf) GetUnsupportedFeatureMatrix() []CapabilityUnsupportedFeatureConfig`

GetUnsupportedFeatureMatrix returns the UnsupportedFeatureMatrix field if non-nil, zero value otherwise.

### GetUnsupportedFeatureMatrixOk

`func (o *CapabilityFeatureConfigAllOf) GetUnsupportedFeatureMatrixOk() (*[]CapabilityUnsupportedFeatureConfig, bool)`

GetUnsupportedFeatureMatrixOk returns a tuple with the UnsupportedFeatureMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfigAllOf) SetUnsupportedFeatureMatrix(v []CapabilityUnsupportedFeatureConfig)`

SetUnsupportedFeatureMatrix sets UnsupportedFeatureMatrix field to given value.

### HasUnsupportedFeatureMatrix

`func (o *CapabilityFeatureConfigAllOf) HasUnsupportedFeatureMatrix() bool`

HasUnsupportedFeatureMatrix returns a boolean if a field has been set.

### SetUnsupportedFeatureMatrixNil

`func (o *CapabilityFeatureConfigAllOf) SetUnsupportedFeatureMatrixNil(b bool)`

 SetUnsupportedFeatureMatrixNil sets the value for UnsupportedFeatureMatrix to be an explicit nil

### UnsetUnsupportedFeatureMatrix
`func (o *CapabilityFeatureConfigAllOf) UnsetUnsupportedFeatureMatrix()`

UnsetUnsupportedFeatureMatrix ensures that no value is present for UnsupportedFeatureMatrix, not even an explicit nil
### GetValidationAction

`func (o *CapabilityFeatureConfigAllOf) GetValidationAction() string`

GetValidationAction returns the ValidationAction field if non-nil, zero value otherwise.

### GetValidationActionOk

`func (o *CapabilityFeatureConfigAllOf) GetValidationActionOk() (*string, bool)`

GetValidationActionOk returns a tuple with the ValidationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationAction

`func (o *CapabilityFeatureConfigAllOf) SetValidationAction(v string)`

SetValidationAction sets ValidationAction field to given value.

### HasValidationAction

`func (o *CapabilityFeatureConfigAllOf) HasValidationAction() bool`

HasValidationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


