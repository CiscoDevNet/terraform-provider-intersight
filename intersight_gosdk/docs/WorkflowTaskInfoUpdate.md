# WorkflowTaskInfoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskInfoUpdate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskInfoUpdate"]
**Input** | Pointer to **interface{}** | Inputs for the specified TaskInfo. Inputs must only be provided for tasks which has included an input definition and the inputs must match the constraints specified in the input definition. | [optional] 
**Name** | Pointer to **string** | Name of the task being updated and this name must match the task instance name included inside the workflow definition. This name is also captured in the RefName property of the TaskInfo object for the task. | [optional] 
**Status** | Pointer to **string** | New status of the task being updated, only Failed and Completed statuses are supported. * &#x60;Scheduled&#x60; - The enum represents the status when task is in scheduled state. * &#x60;InProgress&#x60; - The enum represents the status when task is in-progress state. * &#x60;NoOp&#x60; - The enum represents the status when task is a noop. * &#x60;Timeout&#x60; - The enum represents the status when task has timed out. * &#x60;Completed&#x60; - The enum represents the status when task has completed. * &#x60;Failed&#x60; - The enum represents the status when task has failed. | [optional] [default to "Scheduled"]

## Methods

### NewWorkflowTaskInfoUpdate

`func NewWorkflowTaskInfoUpdate(classId string, objectType string, ) *WorkflowTaskInfoUpdate`

NewWorkflowTaskInfoUpdate instantiates a new WorkflowTaskInfoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskInfoUpdateWithDefaults

`func NewWorkflowTaskInfoUpdateWithDefaults() *WorkflowTaskInfoUpdate`

NewWorkflowTaskInfoUpdateWithDefaults instantiates a new WorkflowTaskInfoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskInfoUpdate) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskInfoUpdate) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskInfoUpdate) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskInfoUpdate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskInfoUpdate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskInfoUpdate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInput

`func (o *WorkflowTaskInfoUpdate) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowTaskInfoUpdate) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowTaskInfoUpdate) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowTaskInfoUpdate) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowTaskInfoUpdate) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowTaskInfoUpdate) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *WorkflowTaskInfoUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskInfoUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskInfoUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskInfoUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowTaskInfoUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTaskInfoUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTaskInfoUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTaskInfoUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


