# WorkflowWorkflowNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowNotification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowNotification"]
**CorrelationId** | Pointer to **string** | The correlationId of the workflow. | [optional] 
**EndTime** | Pointer to **string** | The end time of the workflow. | [optional] 
**Event** | Pointer to **string** | The event of the workflow. | [optional] 
**ExecutionTime** | Pointer to **float32** | The execution time of the workflow. | [optional] 
**FailedReferenceTaskNames** | Pointer to **string** | The reference task names of the failed tasks. | [optional] 
**Input** | Pointer to **string** | The input of the workflow. | [optional] 
**Output** | Pointer to **string** | The output of the workflow. | [optional] 
**ReasonForIncompletion** | Pointer to **string** | The reason for incompletion status of the workflow. | [optional] 
**StartTime** | Pointer to **string** | The start time of the workflow. | [optional] 
**Status** | Pointer to **string** | The final status of the workflow. | [optional] 
**UpdateTime** | Pointer to **string** | The last update time of the workflow. | [optional] 
**Version** | Pointer to **int64** | The version of the workflow. | [optional] 
**WorkflowId** | Pointer to **string** | The unique id of the workflow. | [optional] 
**WorkflowType** | Pointer to **string** | The type of the workflow. | [optional] 

## Methods

### NewWorkflowWorkflowNotification

`func NewWorkflowWorkflowNotification(classId string, objectType string, ) *WorkflowWorkflowNotification`

NewWorkflowWorkflowNotification instantiates a new WorkflowWorkflowNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowNotificationWithDefaults

`func NewWorkflowWorkflowNotificationWithDefaults() *WorkflowWorkflowNotification`

NewWorkflowWorkflowNotificationWithDefaults instantiates a new WorkflowWorkflowNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowNotification) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowNotification) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowNotification) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowNotification) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowNotification) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowNotification) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCorrelationId

`func (o *WorkflowWorkflowNotification) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowWorkflowNotification) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowWorkflowNotification) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowWorkflowNotification) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowNotification) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowNotification) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowNotification) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowNotification) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEvent

`func (o *WorkflowWorkflowNotification) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowWorkflowNotification) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowWorkflowNotification) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowWorkflowNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExecutionTime

`func (o *WorkflowWorkflowNotification) GetExecutionTime() float32`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *WorkflowWorkflowNotification) GetExecutionTimeOk() (*float32, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *WorkflowWorkflowNotification) SetExecutionTime(v float32)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *WorkflowWorkflowNotification) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotification) GetFailedReferenceTaskNames() string`

GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field if non-nil, zero value otherwise.

### GetFailedReferenceTaskNamesOk

`func (o *WorkflowWorkflowNotification) GetFailedReferenceTaskNamesOk() (*string, bool)`

GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotification) SetFailedReferenceTaskNames(v string)`

SetFailedReferenceTaskNames sets FailedReferenceTaskNames field to given value.

### HasFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotification) HasFailedReferenceTaskNames() bool`

HasFailedReferenceTaskNames returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowNotification) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowNotification) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowNotification) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowNotification) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowNotification) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowNotification) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowNotification) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowNotification) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *WorkflowWorkflowNotification) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *WorkflowWorkflowNotification) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *WorkflowWorkflowNotification) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *WorkflowWorkflowNotification) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowNotification) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowNotification) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowNotification) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowNotification) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowNotification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowNotification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowNotification) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowNotification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WorkflowWorkflowNotification) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowWorkflowNotification) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowWorkflowNotification) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowWorkflowNotification) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkflowNotification) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowNotification) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowNotification) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowNotification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowWorkflowNotification) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowWorkflowNotification) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowWorkflowNotification) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowWorkflowNotification) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowWorkflowNotification) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowWorkflowNotification) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowWorkflowNotification) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowWorkflowNotification) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


