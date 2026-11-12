# WorkloadBlueprintReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.BlueprintReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.BlueprintReference"]
**BlueprintVersionUpdateState** | Pointer to [**NullableWorkloadBlueprintVersionUpdateState**](WorkloadBlueprintVersionUpdateState.md) |  | [optional] 
**InputOverride** | Pointer to **[]string** |  | [optional] 
**InputSourceReferences** | Pointer to [**[]WorkloadInputSourceReference**](WorkloadInputSourceReference.md) |  | [optional] 

## Methods

### NewWorkloadBlueprintReference

`func NewWorkloadBlueprintReference(classId string, objectType string, ) *WorkloadBlueprintReference`

NewWorkloadBlueprintReference instantiates a new WorkloadBlueprintReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBlueprintReferenceWithDefaults

`func NewWorkloadBlueprintReferenceWithDefaults() *WorkloadBlueprintReference`

NewWorkloadBlueprintReferenceWithDefaults instantiates a new WorkloadBlueprintReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadBlueprintReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadBlueprintReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadBlueprintReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadBlueprintReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadBlueprintReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadBlueprintReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintVersionUpdateState

`func (o *WorkloadBlueprintReference) GetBlueprintVersionUpdateState() WorkloadBlueprintVersionUpdateState`

GetBlueprintVersionUpdateState returns the BlueprintVersionUpdateState field if non-nil, zero value otherwise.

### GetBlueprintVersionUpdateStateOk

`func (o *WorkloadBlueprintReference) GetBlueprintVersionUpdateStateOk() (*WorkloadBlueprintVersionUpdateState, bool)`

GetBlueprintVersionUpdateStateOk returns a tuple with the BlueprintVersionUpdateState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintVersionUpdateState

`func (o *WorkloadBlueprintReference) SetBlueprintVersionUpdateState(v WorkloadBlueprintVersionUpdateState)`

SetBlueprintVersionUpdateState sets BlueprintVersionUpdateState field to given value.

### HasBlueprintVersionUpdateState

`func (o *WorkloadBlueprintReference) HasBlueprintVersionUpdateState() bool`

HasBlueprintVersionUpdateState returns a boolean if a field has been set.

### SetBlueprintVersionUpdateStateNil

`func (o *WorkloadBlueprintReference) SetBlueprintVersionUpdateStateNil(b bool)`

 SetBlueprintVersionUpdateStateNil sets the value for BlueprintVersionUpdateState to be an explicit nil

### UnsetBlueprintVersionUpdateState
`func (o *WorkloadBlueprintReference) UnsetBlueprintVersionUpdateState()`

UnsetBlueprintVersionUpdateState ensures that no value is present for BlueprintVersionUpdateState, not even an explicit nil
### GetInputOverride

`func (o *WorkloadBlueprintReference) GetInputOverride() []string`

GetInputOverride returns the InputOverride field if non-nil, zero value otherwise.

### GetInputOverrideOk

`func (o *WorkloadBlueprintReference) GetInputOverrideOk() (*[]string, bool)`

GetInputOverrideOk returns a tuple with the InputOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputOverride

`func (o *WorkloadBlueprintReference) SetInputOverride(v []string)`

SetInputOverride sets InputOverride field to given value.

### HasInputOverride

`func (o *WorkloadBlueprintReference) HasInputOverride() bool`

HasInputOverride returns a boolean if a field has been set.

### SetInputOverrideNil

`func (o *WorkloadBlueprintReference) SetInputOverrideNil(b bool)`

 SetInputOverrideNil sets the value for InputOverride to be an explicit nil

### UnsetInputOverride
`func (o *WorkloadBlueprintReference) UnsetInputOverride()`

UnsetInputOverride ensures that no value is present for InputOverride, not even an explicit nil
### GetInputSourceReferences

`func (o *WorkloadBlueprintReference) GetInputSourceReferences() []WorkloadInputSourceReference`

GetInputSourceReferences returns the InputSourceReferences field if non-nil, zero value otherwise.

### GetInputSourceReferencesOk

`func (o *WorkloadBlueprintReference) GetInputSourceReferencesOk() (*[]WorkloadInputSourceReference, bool)`

GetInputSourceReferencesOk returns a tuple with the InputSourceReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSourceReferences

`func (o *WorkloadBlueprintReference) SetInputSourceReferences(v []WorkloadInputSourceReference)`

SetInputSourceReferences sets InputSourceReferences field to given value.

### HasInputSourceReferences

`func (o *WorkloadBlueprintReference) HasInputSourceReferences() bool`

HasInputSourceReferences returns a boolean if a field has been set.

### SetInputSourceReferencesNil

`func (o *WorkloadBlueprintReference) SetInputSourceReferencesNil(b bool)`

 SetInputSourceReferencesNil sets the value for InputSourceReferences to be an explicit nil

### UnsetInputSourceReferences
`func (o *WorkloadBlueprintReference) UnsetInputSourceReferences()`

UnsetInputSourceReferences ensures that no value is present for InputSourceReferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


