# WorkflowAbstractLoopTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Count** | Pointer to **string** | Count value for the loop, this can be a static value defined as a constant at design time or can be a dynamic value defined as an expression that will evaluate to an integer value at execution time. Dynamic values for count must be specified as a template. For example, if a loop must run for a count which matches the length of a workflow input called StringArray, then the count must be specified using a template &#39;{{ len .global.workflow.input.StringArray }}&#39;. The count must be less than or equal to 100. If count is given as a dynamic value, and during execution time if count evaluates to be a value greater than 100, then the loop task will fail. | [optional] 
**LoopStartTask** | Pointer to **string** | Start task where the list of tasks will be executed multiple times based on the count or condition value. | [optional] 
**OnFailure** | Pointer to **string** | This specifies the name of the next task to run if all iterations of the loop task do not succeed. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node. | [optional] 
**OnSuccess** | Pointer to **string** | This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node. | [optional] 

## Methods

### NewWorkflowAbstractLoopTaskAllOf

`func NewWorkflowAbstractLoopTaskAllOf(classId string, objectType string, ) *WorkflowAbstractLoopTaskAllOf`

NewWorkflowAbstractLoopTaskAllOf instantiates a new WorkflowAbstractLoopTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAbstractLoopTaskAllOfWithDefaults

`func NewWorkflowAbstractLoopTaskAllOfWithDefaults() *WorkflowAbstractLoopTaskAllOf`

NewWorkflowAbstractLoopTaskAllOfWithDefaults instantiates a new WorkflowAbstractLoopTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowAbstractLoopTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowAbstractLoopTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowAbstractLoopTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAbstractLoopTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *WorkflowAbstractLoopTaskAllOf) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkflowAbstractLoopTaskAllOf) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *WorkflowAbstractLoopTaskAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLoopStartTask

`func (o *WorkflowAbstractLoopTaskAllOf) GetLoopStartTask() string`

GetLoopStartTask returns the LoopStartTask field if non-nil, zero value otherwise.

### GetLoopStartTaskOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetLoopStartTaskOk() (*string, bool)`

GetLoopStartTaskOk returns a tuple with the LoopStartTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopStartTask

`func (o *WorkflowAbstractLoopTaskAllOf) SetLoopStartTask(v string)`

SetLoopStartTask sets LoopStartTask field to given value.

### HasLoopStartTask

`func (o *WorkflowAbstractLoopTaskAllOf) HasLoopStartTask() bool`

HasLoopStartTask returns a boolean if a field has been set.

### GetOnFailure

`func (o *WorkflowAbstractLoopTaskAllOf) GetOnFailure() string`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetOnFailureOk() (*string, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *WorkflowAbstractLoopTaskAllOf) SetOnFailure(v string)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *WorkflowAbstractLoopTaskAllOf) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowAbstractLoopTaskAllOf) GetOnSuccess() string`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowAbstractLoopTaskAllOf) GetOnSuccessOk() (*string, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowAbstractLoopTaskAllOf) SetOnSuccess(v string)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowAbstractLoopTaskAllOf) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


