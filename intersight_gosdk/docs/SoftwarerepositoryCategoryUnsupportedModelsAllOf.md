# SoftwarerepositoryCategoryUnsupportedModelsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.CategoryUnsupportedModels"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.CategoryUnsupportedModels"]
**MdfId** | Pointer to **string** | Cisco software repository image category identifier. | [optional] 
**ModelConstraint** | Pointer to [**[]SoftwarerepositoryUnsupportedModelConstraint**](SoftwarerepositoryUnsupportedModelConstraint.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCategoryUnsupportedModelsAllOf

`func NewSoftwarerepositoryCategoryUnsupportedModelsAllOf(classId string, objectType string, ) *SoftwarerepositoryCategoryUnsupportedModelsAllOf`

NewSoftwarerepositoryCategoryUnsupportedModelsAllOf instantiates a new SoftwarerepositoryCategoryUnsupportedModelsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCategoryUnsupportedModelsAllOfWithDefaults

`func NewSoftwarerepositoryCategoryUnsupportedModelsAllOfWithDefaults() *SoftwarerepositoryCategoryUnsupportedModelsAllOf`

NewSoftwarerepositoryCategoryUnsupportedModelsAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryUnsupportedModelsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMdfId

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetMdfId() string`

GetMdfId returns the MdfId field if non-nil, zero value otherwise.

### GetMdfIdOk

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetMdfIdOk() (*string, bool)`

GetMdfIdOk returns a tuple with the MdfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfId

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) SetMdfId(v string)`

SetMdfId sets MdfId field to given value.

### HasMdfId

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) HasMdfId() bool`

HasMdfId returns a boolean if a field has been set.

### GetModelConstraint

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetModelConstraint() []SoftwarerepositoryUnsupportedModelConstraint`

GetModelConstraint returns the ModelConstraint field if non-nil, zero value otherwise.

### GetModelConstraintOk

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) GetModelConstraintOk() (*[]SoftwarerepositoryUnsupportedModelConstraint, bool)`

GetModelConstraintOk returns a tuple with the ModelConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConstraint

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) SetModelConstraint(v []SoftwarerepositoryUnsupportedModelConstraint)`

SetModelConstraint sets ModelConstraint field to given value.

### HasModelConstraint

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) HasModelConstraint() bool`

HasModelConstraint returns a boolean if a field has been set.

### SetModelConstraintNil

`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) SetModelConstraintNil(b bool)`

 SetModelConstraintNil sets the value for ModelConstraint to be an explicit nil

### UnsetModelConstraint
`func (o *SoftwarerepositoryCategoryUnsupportedModelsAllOf) UnsetModelConstraint()`

UnsetModelConstraint ensures that no value is present for ModelConstraint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


