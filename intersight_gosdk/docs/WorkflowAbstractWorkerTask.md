# WorkflowAbstractWorkerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**InputParameters** | Pointer to **interface{}** | JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is &#39;${Source.input/output.JsonPath}&#39;. &#39;Source&#39; can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source&#39;s input/output. | [optional] 
**OnFailure** | Pointer to **string** | This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**RollbackDisabled** | Pointer to **bool** | The task is disabled/enabled for rollback operation in this workflow if the task has rollback support. | [optional] [default to false]
**UseDefault** | Pointer to **bool** | UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution. | [optional] [default to false]

## Methods

### NewWorkflowAbstractWorkerTask

`func NewWorkflowAbstractWorkerTask(classId string, objectType string, ) *WorkflowAbstractWorkerTask`

NewWorkflowAbstractWorkerTask instantiates a new WorkflowAbstractWorkerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAbstractWorkerTaskWithDefaults

`func NewWorkflowAbstractWorkerTaskWithDefaults() *WorkflowAbstractWorkerTask`

NewWorkflowAbstractWorkerTaskWithDefaults instantiates a new WorkflowAbstractWorkerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowAbstractWorkerTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowAbstractWorkerTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowAbstractWorkerTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowAbstractWorkerTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAbstractWorkerTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAbstractWorkerTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInputParameters

`func (o *WorkflowAbstractWorkerTask) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowAbstractWorkerTask) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowAbstractWorkerTask) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowAbstractWorkerTask) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowAbstractWorkerTask) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowAbstractWorkerTask) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetOnFailure

`func (o *WorkflowAbstractWorkerTask) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *WorkflowAbstractWorkerTask) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *WorkflowAbstractWorkerTask) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *WorkflowAbstractWorkerTask) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowAbstractWorkerTask) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowAbstractWorkerTask) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowAbstractWorkerTask) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowAbstractWorkerTask) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetRollbackDisabled

`func (o *WorkflowAbstractWorkerTask) GetRollbackDisabled() bool`

GetRollbackDisabled returns the RollbackDisabled field if non-nil, zero value otherwise.

### GetRollbackDisabledOk

`func (o *WorkflowAbstractWorkerTask) GetRollbackDisabledOk() (*bool, bool)`

GetRollbackDisabledOk returns a tuple with the RollbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackDisabled

`func (o *WorkflowAbstractWorkerTask) SetRollbackDisabled(v bool)`

SetRollbackDisabled sets RollbackDisabled field to given value.

### HasRollbackDisabled

`func (o *WorkflowAbstractWorkerTask) HasRollbackDisabled() bool`

HasRollbackDisabled returns a boolean if a field has been set.

### GetUseDefault

`func (o *WorkflowAbstractWorkerTask) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *WorkflowAbstractWorkerTask) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *WorkflowAbstractWorkerTask) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.

### HasUseDefault

`func (o *WorkflowAbstractWorkerTask) HasUseDefault() bool`

HasUseDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


