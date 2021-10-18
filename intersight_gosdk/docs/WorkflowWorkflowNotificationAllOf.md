# WorkflowWorkflowNotificationAllOf

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

### NewWorkflowWorkflowNotificationAllOf

`func NewWorkflowWorkflowNotificationAllOf(classId string, objectType string, ) *WorkflowWorkflowNotificationAllOf`

NewWorkflowWorkflowNotificationAllOf instantiates a new WorkflowWorkflowNotificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowNotificationAllOfWithDefaults

`func NewWorkflowWorkflowNotificationAllOfWithDefaults() *WorkflowWorkflowNotificationAllOf`

NewWorkflowWorkflowNotificationAllOfWithDefaults instantiates a new WorkflowWorkflowNotificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowNotificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowNotificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowNotificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowNotificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowNotificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCorrelationId

`func (o *WorkflowWorkflowNotificationAllOf) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *WorkflowWorkflowNotificationAllOf) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *WorkflowWorkflowNotificationAllOf) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *WorkflowWorkflowNotificationAllOf) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkflowWorkflowNotificationAllOf) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowWorkflowNotificationAllOf) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkflowWorkflowNotificationAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEvent

`func (o *WorkflowWorkflowNotificationAllOf) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WorkflowWorkflowNotificationAllOf) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WorkflowWorkflowNotificationAllOf) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WorkflowWorkflowNotificationAllOf) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExecutionTime

`func (o *WorkflowWorkflowNotificationAllOf) GetExecutionTime() float32`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetExecutionTimeOk() (*float32, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *WorkflowWorkflowNotificationAllOf) SetExecutionTime(v float32)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *WorkflowWorkflowNotificationAllOf) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotificationAllOf) GetFailedReferenceTaskNames() string`

GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field if non-nil, zero value otherwise.

### GetFailedReferenceTaskNamesOk

`func (o *WorkflowWorkflowNotificationAllOf) GetFailedReferenceTaskNamesOk() (*string, bool)`

GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotificationAllOf) SetFailedReferenceTaskNames(v string)`

SetFailedReferenceTaskNames sets FailedReferenceTaskNames field to given value.

### HasFailedReferenceTaskNames

`func (o *WorkflowWorkflowNotificationAllOf) HasFailedReferenceTaskNames() bool`

HasFailedReferenceTaskNames returns a boolean if a field has been set.

### GetInput

`func (o *WorkflowWorkflowNotificationAllOf) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowWorkflowNotificationAllOf) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowWorkflowNotificationAllOf) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowWorkflowNotificationAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowWorkflowNotificationAllOf) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowWorkflowNotificationAllOf) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowWorkflowNotificationAllOf) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowWorkflowNotificationAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetReasonForIncompletion

`func (o *WorkflowWorkflowNotificationAllOf) GetReasonForIncompletion() string`

GetReasonForIncompletion returns the ReasonForIncompletion field if non-nil, zero value otherwise.

### GetReasonForIncompletionOk

`func (o *WorkflowWorkflowNotificationAllOf) GetReasonForIncompletionOk() (*string, bool)`

GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForIncompletion

`func (o *WorkflowWorkflowNotificationAllOf) SetReasonForIncompletion(v string)`

SetReasonForIncompletion sets ReasonForIncompletion field to given value.

### HasReasonForIncompletion

`func (o *WorkflowWorkflowNotificationAllOf) HasReasonForIncompletion() bool`

HasReasonForIncompletion returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkflowWorkflowNotificationAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowWorkflowNotificationAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkflowWorkflowNotificationAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowWorkflowNotificationAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowWorkflowNotificationAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowWorkflowNotificationAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowWorkflowNotificationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WorkflowWorkflowNotificationAllOf) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WorkflowWorkflowNotificationAllOf) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WorkflowWorkflowNotificationAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkflowNotificationAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowNotificationAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowNotificationAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowNotificationAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowWorkflowNotificationAllOf) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowWorkflowNotificationAllOf) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowWorkflowNotificationAllOf) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowWorkflowNotificationAllOf) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


