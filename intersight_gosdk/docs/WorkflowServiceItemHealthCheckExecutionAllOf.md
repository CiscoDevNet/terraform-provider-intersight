# WorkflowServiceItemHealthCheckExecutionAllOf

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

### NewWorkflowServiceItemHealthCheckExecutionAllOf

`func NewWorkflowServiceItemHealthCheckExecutionAllOf(classId string, objectType string, ) *WorkflowServiceItemHealthCheckExecutionAllOf`

NewWorkflowServiceItemHealthCheckExecutionAllOf instantiates a new WorkflowServiceItemHealthCheckExecutionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemHealthCheckExecutionAllOfWithDefaults

`func NewWorkflowServiceItemHealthCheckExecutionAllOfWithDefaults() *WorkflowServiceItemHealthCheckExecutionAllOf`

NewWorkflowServiceItemHealthCheckExecutionAllOfWithDefaults instantiates a new WorkflowServiceItemHealthCheckExecutionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorElements

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorElements() []ServiceitemHealthCheckErrorElement`

GetErrorElements returns the ErrorElements field if non-nil, zero value otherwise.

### GetErrorElementsOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorElementsOk() (*[]ServiceitemHealthCheckErrorElement, bool)`

GetErrorElementsOk returns a tuple with the ErrorElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorElements

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetErrorElements(v []ServiceitemHealthCheckErrorElement)`

SetErrorElements sets ErrorElements field to given value.

### HasErrorElements

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasErrorElements() bool`

HasErrorElements returns a boolean if a field has been set.

### SetErrorElementsNil

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetErrorElementsNil(b bool)`

 SetErrorElementsNil sets the value for ErrorElements to be an explicit nil

### UnsetErrorElements
`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) UnsetErrorElements()`

UnsetErrorElements ensures that no value is present for ErrorElements, not even an explicit nil
### GetErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetLastPassedTimestamp() time.Time`

GetLastPassedTimestamp returns the LastPassedTimestamp field if non-nil, zero value otherwise.

### GetLastPassedTimestampOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetLastPassedTimestampOk() (*time.Time, bool)`

GetLastPassedTimestampOk returns a tuple with the LastPassedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetLastPassedTimestamp(v time.Time)`

SetLastPassedTimestamp sets LastPassedTimestamp field to given value.

### HasLastPassedTimestamp

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasLastPassedTimestamp() bool`

HasLastPassedTimestamp returns a boolean if a field has been set.

### GetResult

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowStatus() string`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowStatusOk() (*string, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetWorkflowStatus(v string)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### GetHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetHealthCheckDefinition() WorkflowServiceItemHealthCheckDefinitionRelationship`

GetHealthCheckDefinition returns the HealthCheckDefinition field if non-nil, zero value otherwise.

### GetHealthCheckDefinitionOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetHealthCheckDefinitionOk() (*WorkflowServiceItemHealthCheckDefinitionRelationship, bool)`

GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetHealthCheckDefinition(v WorkflowServiceItemHealthCheckDefinitionRelationship)`

SetHealthCheckDefinition sets HealthCheckDefinition field to given value.

### HasHealthCheckDefinition

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasHealthCheckDefinition() bool`

HasHealthCheckDefinition returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


