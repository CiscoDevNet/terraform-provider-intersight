# WorkflowAbstractWorkerTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**InputParameters** | Pointer to **interface{}** | JSON formatted key-value pairs that define the inputs given to the task. Mapping for task inputs can be provided as either static values, direct mapping or advanced mapping using templates. The direct mapping can be specified as &#39;${Source.&lt; input | output | variable&gt;.&lt;JSONPath&gt;}&#39;. &#39;Source&#39; can be either workflow or the name of an earlier task within the workflow. You can map the task input to either a workflow input, a task output or a variable. Golang template syntax is supported for advanced mapping. A simple flattened example is \&quot;InputParameters\&quot;:{ \&quot;input1\&quot;:\&quot;${workflow.variable.var1}\&quot;, \&quot;input2\&quot;:\&quot;prefixStr_{{.global.workflow.input.input1}}\&quot; } where task input1 is mapped directly to variable var1 and task input2 is using a template to prefix a string to workflow input1 and then assign that value. | [optional] 
**OnFailure** | Pointer to **string** | This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**RollbackDisabled** | Pointer to **bool** | The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. | [optional] [default to false]
**VariableParameters** | Pointer to **interface{}** | JSON formatted key-value pairs that perform variable update at the end of the task execution. Mapping for variables can be provided as either static values, direct mapping or advanced mapping using templates. The direct mapping can be specified as &#39;${Source.&lt; input | output | variable&gt;.&lt;JSONPath&gt;}&#39;. &#39;Source&#39; can be either workflow or the name of the current or an earlier task within the workflow. You can map the variable to either a workflow input, a task output or another variable. Golang template syntax is supported for advanced mapping. A simple flattened example is \&quot;VariableParameters\&quot;:{ \&quot;var1\&quot;:\&quot;${task1.output.output1}\&quot;, \&quot;var2\&quot;:\&quot;{{ Itoa .global.workflow.variable.varInt}}\&quot; } where variable var1 is mapped directly to output1 of task1 and variable var2 is using a template to convert another variable varInt to string and assign that value. | [optional] 

## Methods

### NewWorkflowAbstractWorkerTaskAllOf

`func NewWorkflowAbstractWorkerTaskAllOf(classId string, objectType string, ) *WorkflowAbstractWorkerTaskAllOf`

NewWorkflowAbstractWorkerTaskAllOf instantiates a new WorkflowAbstractWorkerTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAbstractWorkerTaskAllOfWithDefaults

`func NewWorkflowAbstractWorkerTaskAllOfWithDefaults() *WorkflowAbstractWorkerTaskAllOf`

NewWorkflowAbstractWorkerTaskAllOfWithDefaults instantiates a new WorkflowAbstractWorkerTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowAbstractWorkerTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowAbstractWorkerTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowAbstractWorkerTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAbstractWorkerTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInputParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowAbstractWorkerTaskAllOf) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowAbstractWorkerTaskAllOf) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetOnFailure

`func (o *WorkflowAbstractWorkerTaskAllOf) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *WorkflowAbstractWorkerTaskAllOf) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *WorkflowAbstractWorkerTaskAllOf) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowAbstractWorkerTaskAllOf) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowAbstractWorkerTaskAllOf) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowAbstractWorkerTaskAllOf) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetRollbackDisabled

`func (o *WorkflowAbstractWorkerTaskAllOf) GetRollbackDisabled() bool`

GetRollbackDisabled returns the RollbackDisabled field if non-nil, zero value otherwise.

### GetRollbackDisabledOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetRollbackDisabledOk() (*bool, bool)`

GetRollbackDisabledOk returns a tuple with the RollbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackDisabled

`func (o *WorkflowAbstractWorkerTaskAllOf) SetRollbackDisabled(v bool)`

SetRollbackDisabled sets RollbackDisabled field to given value.

### HasRollbackDisabled

`func (o *WorkflowAbstractWorkerTaskAllOf) HasRollbackDisabled() bool`

HasRollbackDisabled returns a boolean if a field has been set.

### GetVariableParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) GetVariableParameters() interface{}`

GetVariableParameters returns the VariableParameters field if non-nil, zero value otherwise.

### GetVariableParametersOk

`func (o *WorkflowAbstractWorkerTaskAllOf) GetVariableParametersOk() (*interface{}, bool)`

GetVariableParametersOk returns a tuple with the VariableParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) SetVariableParameters(v interface{})`

SetVariableParameters sets VariableParameters field to given value.

### HasVariableParameters

`func (o *WorkflowAbstractWorkerTaskAllOf) HasVariableParameters() bool`

HasVariableParameters returns a boolean if a field has been set.

### SetVariableParametersNil

`func (o *WorkflowAbstractWorkerTaskAllOf) SetVariableParametersNil(b bool)`

 SetVariableParametersNil sets the value for VariableParameters to be an explicit nil

### UnsetVariableParameters
`func (o *WorkflowAbstractWorkerTaskAllOf) UnsetVariableParameters()`

UnsetVariableParameters ensures that no value is present for VariableParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


