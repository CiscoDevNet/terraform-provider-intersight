# WorkflowSolutionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionOutput"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionOutput"]
**Name** | Pointer to **string** | Output name which is used in the output definition of the solution. | [optional] 
**Output** | Pointer to **interface{}** | Solution output for a solution instance and the format is specified by output definition of the solution definition. | [optional] 
**SolutionInstance** | Pointer to [**WorkflowSolutionInstanceRelationship**](WorkflowSolutionInstanceRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionOutput

`func NewWorkflowSolutionOutput(classId string, objectType string, ) *WorkflowSolutionOutput`

NewWorkflowSolutionOutput instantiates a new WorkflowSolutionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionOutputWithDefaults

`func NewWorkflowSolutionOutputWithDefaults() *WorkflowSolutionOutput`

NewWorkflowSolutionOutputWithDefaults instantiates a new WorkflowSolutionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionOutput) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionOutput) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionOutput) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionOutput) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionOutput) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionOutput) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *WorkflowSolutionOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowSolutionOutput) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowSolutionOutput) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowSolutionOutput) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowSolutionOutput) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowSolutionOutput) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowSolutionOutput) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetSolutionInstance

`func (o *WorkflowSolutionOutput) GetSolutionInstance() WorkflowSolutionInstanceRelationship`

GetSolutionInstance returns the SolutionInstance field if non-nil, zero value otherwise.

### GetSolutionInstanceOk

`func (o *WorkflowSolutionOutput) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool)`

GetSolutionInstanceOk returns a tuple with the SolutionInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionInstance

`func (o *WorkflowSolutionOutput) SetSolutionInstance(v WorkflowSolutionInstanceRelationship)`

SetSolutionInstance sets SolutionInstance field to given value.

### HasSolutionInstance

`func (o *WorkflowSolutionOutput) HasSolutionInstance() bool`

HasSolutionInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


