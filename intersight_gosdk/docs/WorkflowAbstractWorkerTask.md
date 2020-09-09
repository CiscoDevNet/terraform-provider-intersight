# WorkflowAbstractWorkerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputParameters** | Pointer to **interface{}** | JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is &#39;${Source.input/output.JsonPath}&#39;. &#39;Source&#39; can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source&#39;s input/output. | [optional] 
**OnFailure** | Pointer to **string** | This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node. | [optional] 

## Methods

### NewWorkflowAbstractWorkerTask

`func NewWorkflowAbstractWorkerTask() *WorkflowAbstractWorkerTask`

NewWorkflowAbstractWorkerTask instantiates a new WorkflowAbstractWorkerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAbstractWorkerTaskWithDefaults

`func NewWorkflowAbstractWorkerTaskWithDefaults() *WorkflowAbstractWorkerTask`

NewWorkflowAbstractWorkerTaskWithDefaults instantiates a new WorkflowAbstractWorkerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


