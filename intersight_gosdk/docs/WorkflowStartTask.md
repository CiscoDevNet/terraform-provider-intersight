# WorkflowStartTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.StartTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.StartTask"]
**NextTask** | Pointer to **string** | The name of the next task (Task names unique within workflow) to run.  In a graph model, denotes an edge to another Task Node. | [optional] 

## Methods

### NewWorkflowStartTask

`func NewWorkflowStartTask(classId string, objectType string, ) *WorkflowStartTask`

NewWorkflowStartTask instantiates a new WorkflowStartTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStartTaskWithDefaults

`func NewWorkflowStartTaskWithDefaults() *WorkflowStartTask`

NewWorkflowStartTaskWithDefaults instantiates a new WorkflowStartTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowStartTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowStartTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowStartTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowStartTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowStartTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowStartTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNextTask

`func (o *WorkflowStartTask) GetNextTask() string`

GetNextTask returns the NextTask field if non-nil, zero value otherwise.

### GetNextTaskOk

`func (o *WorkflowStartTask) GetNextTaskOk() (*string, bool)`

GetNextTaskOk returns a tuple with the NextTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTask

`func (o *WorkflowStartTask) SetNextTask(v string)`

SetNextTask sets NextTask field to given value.

### HasNextTask

`func (o *WorkflowStartTask) HasNextTask() bool`

HasNextTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


