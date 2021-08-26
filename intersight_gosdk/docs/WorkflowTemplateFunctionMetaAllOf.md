# WorkflowTemplateFunctionMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TemplateFunctionMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TemplateFunctionMeta"]
**Comments** | Pointer to [**NullableWorkflowComments**](WorkflowComments.md) |  | [optional] 
**Inputs** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**IsGuidedModeSupported** | Pointer to **bool** | The flag indicates whether a guided mode template is supported for it or not. | [optional] [readonly] 
**Name** | Pointer to **string** | The template function name. | [optional] [readonly] 
**Outputs** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 

## Methods

### NewWorkflowTemplateFunctionMetaAllOf

`func NewWorkflowTemplateFunctionMetaAllOf(classId string, objectType string, ) *WorkflowTemplateFunctionMetaAllOf`

NewWorkflowTemplateFunctionMetaAllOf instantiates a new WorkflowTemplateFunctionMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateFunctionMetaAllOfWithDefaults

`func NewWorkflowTemplateFunctionMetaAllOfWithDefaults() *WorkflowTemplateFunctionMetaAllOf`

NewWorkflowTemplateFunctionMetaAllOfWithDefaults instantiates a new WorkflowTemplateFunctionMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTemplateFunctionMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTemplateFunctionMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTemplateFunctionMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTemplateFunctionMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComments

`func (o *WorkflowTemplateFunctionMetaAllOf) GetComments() WorkflowComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetCommentsOk() (*WorkflowComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WorkflowTemplateFunctionMetaAllOf) SetComments(v WorkflowComments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WorkflowTemplateFunctionMetaAllOf) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *WorkflowTemplateFunctionMetaAllOf) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *WorkflowTemplateFunctionMetaAllOf) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetInputs

`func (o *WorkflowTemplateFunctionMetaAllOf) GetInputs() []WorkflowBaseDataType`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetInputsOk() (*[]WorkflowBaseDataType, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *WorkflowTemplateFunctionMetaAllOf) SetInputs(v []WorkflowBaseDataType)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *WorkflowTemplateFunctionMetaAllOf) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *WorkflowTemplateFunctionMetaAllOf) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *WorkflowTemplateFunctionMetaAllOf) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMetaAllOf) GetIsGuidedModeSupported() bool`

GetIsGuidedModeSupported returns the IsGuidedModeSupported field if non-nil, zero value otherwise.

### GetIsGuidedModeSupportedOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetIsGuidedModeSupportedOk() (*bool, bool)`

GetIsGuidedModeSupportedOk returns a tuple with the IsGuidedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMetaAllOf) SetIsGuidedModeSupported(v bool)`

SetIsGuidedModeSupported sets IsGuidedModeSupported field to given value.

### HasIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMetaAllOf) HasIsGuidedModeSupported() bool`

HasIsGuidedModeSupported returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTemplateFunctionMetaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateFunctionMetaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTemplateFunctionMetaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTemplateFunctionMetaAllOf) GetOutputs() []WorkflowBaseDataType`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTemplateFunctionMetaAllOf) GetOutputsOk() (*[]WorkflowBaseDataType, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTemplateFunctionMetaAllOf) SetOutputs(v []WorkflowBaseDataType)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTemplateFunctionMetaAllOf) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowTemplateFunctionMetaAllOf) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowTemplateFunctionMetaAllOf) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


