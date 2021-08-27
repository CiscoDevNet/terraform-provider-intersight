# WorkflowTemplateFunctionMeta

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

### NewWorkflowTemplateFunctionMeta

`func NewWorkflowTemplateFunctionMeta(classId string, objectType string, ) *WorkflowTemplateFunctionMeta`

NewWorkflowTemplateFunctionMeta instantiates a new WorkflowTemplateFunctionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateFunctionMetaWithDefaults

`func NewWorkflowTemplateFunctionMetaWithDefaults() *WorkflowTemplateFunctionMeta`

NewWorkflowTemplateFunctionMetaWithDefaults instantiates a new WorkflowTemplateFunctionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTemplateFunctionMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTemplateFunctionMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTemplateFunctionMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTemplateFunctionMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTemplateFunctionMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTemplateFunctionMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComments

`func (o *WorkflowTemplateFunctionMeta) GetComments() WorkflowComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WorkflowTemplateFunctionMeta) GetCommentsOk() (*WorkflowComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WorkflowTemplateFunctionMeta) SetComments(v WorkflowComments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WorkflowTemplateFunctionMeta) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *WorkflowTemplateFunctionMeta) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *WorkflowTemplateFunctionMeta) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetInputs

`func (o *WorkflowTemplateFunctionMeta) GetInputs() []WorkflowBaseDataType`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *WorkflowTemplateFunctionMeta) GetInputsOk() (*[]WorkflowBaseDataType, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *WorkflowTemplateFunctionMeta) SetInputs(v []WorkflowBaseDataType)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *WorkflowTemplateFunctionMeta) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *WorkflowTemplateFunctionMeta) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *WorkflowTemplateFunctionMeta) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMeta) GetIsGuidedModeSupported() bool`

GetIsGuidedModeSupported returns the IsGuidedModeSupported field if non-nil, zero value otherwise.

### GetIsGuidedModeSupportedOk

`func (o *WorkflowTemplateFunctionMeta) GetIsGuidedModeSupportedOk() (*bool, bool)`

GetIsGuidedModeSupportedOk returns a tuple with the IsGuidedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMeta) SetIsGuidedModeSupported(v bool)`

SetIsGuidedModeSupported sets IsGuidedModeSupported field to given value.

### HasIsGuidedModeSupported

`func (o *WorkflowTemplateFunctionMeta) HasIsGuidedModeSupported() bool`

HasIsGuidedModeSupported returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTemplateFunctionMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateFunctionMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateFunctionMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTemplateFunctionMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTemplateFunctionMeta) GetOutputs() []WorkflowBaseDataType`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTemplateFunctionMeta) GetOutputsOk() (*[]WorkflowBaseDataType, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTemplateFunctionMeta) SetOutputs(v []WorkflowBaseDataType)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTemplateFunctionMeta) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *WorkflowTemplateFunctionMeta) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *WorkflowTemplateFunctionMeta) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


