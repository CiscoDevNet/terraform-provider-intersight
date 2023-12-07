# CapabilityUnsupportedFeatureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.UnsupportedFeatureConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.UnsupportedFeatureConfig"]
**Generation** | Pointer to **int32** | The adapter generations that support this feature. * &#x60;4&#x60; - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * &#x60;2&#x60; - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * &#x60;3&#x60; - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * &#x60;5&#x60; - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string. | [optional] [default to 4]
**UnsupportdFeatures** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCapabilityUnsupportedFeatureConfig

`func NewCapabilityUnsupportedFeatureConfig(classId string, objectType string, ) *CapabilityUnsupportedFeatureConfig`

NewCapabilityUnsupportedFeatureConfig instantiates a new CapabilityUnsupportedFeatureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityUnsupportedFeatureConfigWithDefaults

`func NewCapabilityUnsupportedFeatureConfigWithDefaults() *CapabilityUnsupportedFeatureConfig`

NewCapabilityUnsupportedFeatureConfigWithDefaults instantiates a new CapabilityUnsupportedFeatureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityUnsupportedFeatureConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityUnsupportedFeatureConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityUnsupportedFeatureConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityUnsupportedFeatureConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityUnsupportedFeatureConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityUnsupportedFeatureConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGeneration

`func (o *CapabilityUnsupportedFeatureConfig) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *CapabilityUnsupportedFeatureConfig) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *CapabilityUnsupportedFeatureConfig) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *CapabilityUnsupportedFeatureConfig) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetUnsupportdFeatures

`func (o *CapabilityUnsupportedFeatureConfig) GetUnsupportdFeatures() []string`

GetUnsupportdFeatures returns the UnsupportdFeatures field if non-nil, zero value otherwise.

### GetUnsupportdFeaturesOk

`func (o *CapabilityUnsupportedFeatureConfig) GetUnsupportdFeaturesOk() (*[]string, bool)`

GetUnsupportdFeaturesOk returns a tuple with the UnsupportdFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportdFeatures

`func (o *CapabilityUnsupportedFeatureConfig) SetUnsupportdFeatures(v []string)`

SetUnsupportdFeatures sets UnsupportdFeatures field to given value.

### HasUnsupportdFeatures

`func (o *CapabilityUnsupportedFeatureConfig) HasUnsupportdFeatures() bool`

HasUnsupportdFeatures returns a boolean if a field has been set.

### SetUnsupportdFeaturesNil

`func (o *CapabilityUnsupportedFeatureConfig) SetUnsupportdFeaturesNil(b bool)`

 SetUnsupportdFeaturesNil sets the value for UnsupportdFeatures to be an explicit nil

### UnsetUnsupportdFeatures
`func (o *CapabilityUnsupportedFeatureConfig) UnsetUnsupportdFeatures()`

UnsetUnsupportdFeatures ensures that no value is present for UnsupportdFeatures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


