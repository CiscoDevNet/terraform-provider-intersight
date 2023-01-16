# CapabilityAdapterFirmwareRequirementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.AdapterFirmwareRequirement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.AdapterFirmwareRequirement"]
**AdapterSeries** | Pointer to **string** | Series of adapter. Example - Cruz, Bodega. | [optional] [readonly] 
**IgnoreEmptyCurrentVersion** | Pointer to **bool** | Do not update if the current version is reported as empty. | [optional] [readonly] 
**MinimumAdapterVersion** | Pointer to **string** | The minimum adapter version that supports this adapter upgrade. | [optional] [readonly] 
**MinimumBmcVersion** | Pointer to **string** | The minimum BMC version that supports this adapter upgrade. | [optional] [readonly] 
**RecommendedBmcVersion** | Pointer to **string** | The recommended BMC version that supports this adapter upgrade. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityAdapterFirmwareRequirementAllOf

`func NewCapabilityAdapterFirmwareRequirementAllOf(classId string, objectType string, ) *CapabilityAdapterFirmwareRequirementAllOf`

NewCapabilityAdapterFirmwareRequirementAllOf instantiates a new CapabilityAdapterFirmwareRequirementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityAdapterFirmwareRequirementAllOfWithDefaults

`func NewCapabilityAdapterFirmwareRequirementAllOfWithDefaults() *CapabilityAdapterFirmwareRequirementAllOf`

NewCapabilityAdapterFirmwareRequirementAllOfWithDefaults instantiates a new CapabilityAdapterFirmwareRequirementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdapterSeries

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetAdapterSeries() string`

GetAdapterSeries returns the AdapterSeries field if non-nil, zero value otherwise.

### GetAdapterSeriesOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetAdapterSeriesOk() (*string, bool)`

GetAdapterSeriesOk returns a tuple with the AdapterSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterSeries

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetAdapterSeries(v string)`

SetAdapterSeries sets AdapterSeries field to given value.

### HasAdapterSeries

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasAdapterSeries() bool`

HasAdapterSeries returns a boolean if a field has been set.

### GetIgnoreEmptyCurrentVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetIgnoreEmptyCurrentVersion() bool`

GetIgnoreEmptyCurrentVersion returns the IgnoreEmptyCurrentVersion field if non-nil, zero value otherwise.

### GetIgnoreEmptyCurrentVersionOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetIgnoreEmptyCurrentVersionOk() (*bool, bool)`

GetIgnoreEmptyCurrentVersionOk returns a tuple with the IgnoreEmptyCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEmptyCurrentVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetIgnoreEmptyCurrentVersion(v bool)`

SetIgnoreEmptyCurrentVersion sets IgnoreEmptyCurrentVersion field to given value.

### HasIgnoreEmptyCurrentVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasIgnoreEmptyCurrentVersion() bool`

HasIgnoreEmptyCurrentVersion returns a boolean if a field has been set.

### GetMinimumAdapterVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetMinimumAdapterVersion() string`

GetMinimumAdapterVersion returns the MinimumAdapterVersion field if non-nil, zero value otherwise.

### GetMinimumAdapterVersionOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetMinimumAdapterVersionOk() (*string, bool)`

GetMinimumAdapterVersionOk returns a tuple with the MinimumAdapterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAdapterVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetMinimumAdapterVersion(v string)`

SetMinimumAdapterVersion sets MinimumAdapterVersion field to given value.

### HasMinimumAdapterVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasMinimumAdapterVersion() bool`

HasMinimumAdapterVersion returns a boolean if a field has been set.

### GetMinimumBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetMinimumBmcVersion() string`

GetMinimumBmcVersion returns the MinimumBmcVersion field if non-nil, zero value otherwise.

### GetMinimumBmcVersionOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetMinimumBmcVersionOk() (*string, bool)`

GetMinimumBmcVersionOk returns a tuple with the MinimumBmcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetMinimumBmcVersion(v string)`

SetMinimumBmcVersion sets MinimumBmcVersion field to given value.

### HasMinimumBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasMinimumBmcVersion() bool`

HasMinimumBmcVersion returns a boolean if a field has been set.

### GetRecommendedBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetRecommendedBmcVersion() string`

GetRecommendedBmcVersion returns the RecommendedBmcVersion field if non-nil, zero value otherwise.

### GetRecommendedBmcVersionOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetRecommendedBmcVersionOk() (*string, bool)`

GetRecommendedBmcVersionOk returns a tuple with the RecommendedBmcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetRecommendedBmcVersion(v string)`

SetRecommendedBmcVersion sets RecommendedBmcVersion field to given value.

### HasRecommendedBmcVersion

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasRecommendedBmcVersion() bool`

HasRecommendedBmcVersion returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityAdapterFirmwareRequirementAllOf) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityAdapterFirmwareRequirementAllOf) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityAdapterFirmwareRequirementAllOf) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityAdapterFirmwareRequirementAllOf) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


