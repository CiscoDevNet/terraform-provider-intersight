# WorkflowServiceItemHealthCheckExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemHealthCheckExecution"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemHealthCheckExecution"]
**ErrorElements** | Pointer to [**[]ServiceitemHealthCheckErrorElement**](ServiceitemHealthCheckErrorElement.md) |  | [optional] 
**ErrorSummary** | Pointer to **string** | Error summary of a health check execution failure. | [optional] [readonly] 
**LastPassedTimestamp** | Pointer to **time.Time** | Timestamp of the last passed health check execution. | [optional] [readonly] 
**Result** | Pointer to **string** | Health check execution result. * &#x60;Unknown&#x60; - Indicates that the health check results could not be determined. * &#x60;Pass&#x60; - Indicates that the health check has passed. * &#x60;Fail&#x60; - Indicates that the health check has failed. * &#x60;Warning&#x60; - Indicates that the health check completed with a warning. * &#x60;NotApplicable&#x60; - Indicates that the health check is either unsupported, or not applicable for the service item. | [optional] [readonly] [default to "Unknown"]
**Summary** | Pointer to **string** | A brief summary of health check execution result. | [optional] [readonly] 
**WorkflowStatus** | Pointer to **string** | Status of the workflow that is executed as a part of health check execution. | [optional] [readonly] 
**HealthCheckDefinition** | Pointer to [**WorkflowServiceItemHealthCheckDefinitionRelationship**](WorkflowServiceItemHealthCheckDefinitionRelationship.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**WorkflowServiceItemInstanceRelationship**](WorkflowServiceItemInstanceRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemHealthCheckExecution

`func NewWorkflowServiceItemHealthCheckExecution(classId string, objectType string, ) *WorkflowServiceItemHealthCheckExecution`

NewWorkflowServiceItemHealthCheckExecution instantiates a new WorkflowServiceItemHealthCheckExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemHealthCheckExecutionWithDefaults

`func NewWorkflowServiceItemHealthCheckExecutionWithDefaults() *WorkflowServiceItemHealthCheckExecution`

NewWorkflowServiceItemHealthCheckExecutionWithDefaults instantiates a new WorkflowServiceItemHealthCheckExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemHealthCheckExecution) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemHealthCheckExecution) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemHealthCheckExecution) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemHealthCheckExecution) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorElements

`func (o *WorkflowServiceItemHealthCheckExecution) GetErrorElements() []ServiceitemHealthCheckErrorElement`

GetErrorElements returns the ErrorElements field if non-nil, zero value otherwise.

### GetErrorElementsOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetErrorElementsOk() (*[]ServiceitemHealthCheckErrorElement, bool)`

GetErrorElementsOk returns a tuple with the ErrorElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorElements

`func (o *WorkflowServiceItemHealthCheckExecution) SetErrorElements(v []ServiceitemHealthCheckErrorElement)`

SetErrorElements sets ErrorElements field to given value.

### HasErrorElements

`func (o *WorkflowServiceItemHealthCheckExecution) HasErrorElements() bool`

HasErrorElements returns a boolean if a field has been set.

### SetErrorElementsNil

`func (o *WorkflowServiceItemHealthCheckExecution) SetErrorElementsNil(b bool)`

 SetErrorElementsNil sets the value for ErrorElements to be an explicit nil

### UnsetErrorElements
`func (o *WorkflowServiceItemHealthCheckExecution) UnsetErrorElements()`

UnsetErrorElements ensures that no value is present for ErrorElements, not even an explicit nil
### GetErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecution) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecution) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecution) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecution) GetLastPassedTimestamp() time.Time`

GetLastPassedTimestamp returns the LastPassedTimestamp field if non-nil, zero value otherwise.

### GetLastPassedTimestampOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetLastPassedTimestampOk() (*time.Time, bool)`

GetLastPassedTimestampOk returns a tuple with the LastPassedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecution) SetLastPassedTimestamp(v time.Time)`

SetLastPassedTimestamp sets LastPassedTimestamp field to given value.

### HasLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecution) HasLastPassedTimestamp() bool`

HasLastPassedTimestamp returns a boolean if a field has been set.

### GetResult

`func (o *WorkflowServiceItemHealthCheckExecution) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowServiceItemHealthCheckExecution) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *WorkflowServiceItemHealthCheckExecution) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSummary

`func (o *WorkflowServiceItemHealthCheckExecution) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WorkflowServiceItemHealthCheckExecution) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *WorkflowServiceItemHealthCheckExecution) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecution) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecution) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecution) GetHealthCheckDefinition() WorkflowServiceItemHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetHealthCheckDefinitionOk() (*WorkflowServiceItemHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecution) SetHealthCheckDefinition(v WorkflowServiceItemHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecution) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecution) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecution) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecution) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecution) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecution) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


