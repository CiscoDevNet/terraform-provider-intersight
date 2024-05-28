# BulkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.Result"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.Result"]
**ActionOnError** | Pointer to **string** | The action that will be performed when an error occurs during processing of the request. * &#x60;Stop&#x60; - Stop the processing of the request after the first error. * &#x60;Proceed&#x60; - Proceed with the processing of the request even when an error occurs. | [optional] [readonly] [default to "Stop"]
**CompletionTime** | Pointer to **time.Time** | The timestamp in UTC when the request processing is completed. | [optional] [readonly] 
**NumSubRequests** | Pointer to **int64** | The number of subrequests received in this request. | [optional] [readonly] 
**Request** | Pointer to **interface{}** | The individual request to be executed asynchronously. | [optional] 
**RequestReceivedTime** | Pointer to **time.Time** | The timestamp in UTC when the request was received. | [optional] [readonly] 
**Status** | Pointer to **string** | The processing status of the request. * &#x60;NotStarted&#x60; - Indicates that the request processing has not begun yet. * &#x60;ObjPresenceCheckInProgress&#x60; - Indicates that the object presence check is in progress for this request. * &#x60;ObjPresenceCheckComplete&#x60; - Indicates that the object presence check is complete. * &#x60;ExecutionInProgress&#x60; - Indicates that the request processing is in progress. * &#x60;Completed&#x60; - Indicates that the request processing has been completed successfully. * &#x60;CompletedWithErrors&#x60; - Indicates that the request processing has one or more failed subrequests. * &#x60;Failed&#x60; - Indicates that the processing of this request failed. * &#x60;TimedOut&#x60; - Indicates that the request processing timed out. | [optional] [readonly] [default to "NotStarted"]
**StatusMessage** | Pointer to **string** | The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case. | [optional] [readonly] 
**Uri** | Pointer to **string** | The URI on which this async operation is being performed. | [optional] [readonly] 
**MoCloner** | Pointer to [**NullableBulkMoClonerRelationship**](BulkMoClonerRelationship.md) |  | [optional] 
**MoDeepCloner** | Pointer to [**NullableBulkMoDeepClonerRelationship**](BulkMoDeepClonerRelationship.md) |  | [optional] 
**MoMerger** | Pointer to [**NullableBulkMoMergerRelationship**](BulkMoMergerRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Results** | Pointer to [**[]BulkSubRequestObjRelationship**](BulkSubRequestObjRelationship.md) | An array of relationships to bulkSubRequestObj resources. | [optional] [readonly] 
**WorkflowInfo** | Pointer to [**NullableWorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewBulkResult

`func NewBulkResult(classId string, objectType string, ) *BulkResult`

NewBulkResult instantiates a new BulkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResultWithDefaults

`func NewBulkResultWithDefaults() *BulkResult`

NewBulkResultWithDefaults instantiates a new BulkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionOnError

`func (o *BulkResult) GetActionOnError() string`

GetActionOnError returns the ActionOnError field if non-nil, zero value otherwise.

### GetActionOnErrorOk

`func (o *BulkResult) GetActionOnErrorOk() (*string, bool)`

GetActionOnErrorOk returns a tuple with the ActionOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnError

`func (o *BulkResult) SetActionOnError(v string)`

SetActionOnError sets ActionOnError field to given value.

### HasActionOnError

`func (o *BulkResult) HasActionOnError() bool`

HasActionOnError returns a boolean if a field has been set.

### GetCompletionTime

`func (o *BulkResult) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *BulkResult) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *BulkResult) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *BulkResult) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetNumSubRequests

`func (o *BulkResult) GetNumSubRequests() int64`

GetNumSubRequests returns the NumSubRequests field if non-nil, zero value otherwise.

### GetNumSubRequestsOk

`func (o *BulkResult) GetNumSubRequestsOk() (*int64, bool)`

GetNumSubRequestsOk returns a tuple with the NumSubRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubRequests

`func (o *BulkResult) SetNumSubRequests(v int64)`

SetNumSubRequests sets NumSubRequests field to given value.

### HasNumSubRequests

`func (o *BulkResult) HasNumSubRequests() bool`

HasNumSubRequests returns a boolean if a field has been set.

### GetRequest

`func (o *BulkResult) GetRequest() interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BulkResult) GetRequestOk() (*interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BulkResult) SetRequest(v interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *BulkResult) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *BulkResult) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *BulkResult) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetRequestReceivedTime

`func (o *BulkResult) GetRequestReceivedTime() time.Time`

GetRequestReceivedTime returns the RequestReceivedTime field if non-nil, zero value otherwise.

### GetRequestReceivedTimeOk

`func (o *BulkResult) GetRequestReceivedTimeOk() (*time.Time, bool)`

GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReceivedTime

`func (o *BulkResult) SetRequestReceivedTime(v time.Time)`

SetRequestReceivedTime sets RequestReceivedTime field to given value.

### HasRequestReceivedTime

`func (o *BulkResult) HasRequestReceivedTime() bool`

HasRequestReceivedTime returns a boolean if a field has been set.

### GetStatus

`func (o *BulkResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *BulkResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BulkResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BulkResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *BulkResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUri

`func (o *BulkResult) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkResult) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkResult) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkResult) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetMoCloner

`func (o *BulkResult) GetMoCloner() BulkMoClonerRelationship`

GetMoCloner returns the MoCloner field if non-nil, zero value otherwise.

### GetMoClonerOk

`func (o *BulkResult) GetMoClonerOk() (*BulkMoClonerRelationship, bool)`

GetMoClonerOk returns a tuple with the MoCloner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoCloner

`func (o *BulkResult) SetMoCloner(v BulkMoClonerRelationship)`

SetMoCloner sets MoCloner field to given value.

### HasMoCloner

`func (o *BulkResult) HasMoCloner() bool`

HasMoCloner returns a boolean if a field has been set.

### SetMoClonerNil

`func (o *BulkResult) SetMoClonerNil(b bool)`

 SetMoClonerNil sets the value for MoCloner to be an explicit nil

### UnsetMoCloner
`func (o *BulkResult) UnsetMoCloner()`

UnsetMoCloner ensures that no value is present for MoCloner, not even an explicit nil
### GetMoDeepCloner

`func (o *BulkResult) GetMoDeepCloner() BulkMoDeepClonerRelationship`

GetMoDeepCloner returns the MoDeepCloner field if non-nil, zero value otherwise.

### GetMoDeepClonerOk

`func (o *BulkResult) GetMoDeepClonerOk() (*BulkMoDeepClonerRelationship, bool)`

GetMoDeepClonerOk returns a tuple with the MoDeepCloner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoDeepCloner

`func (o *BulkResult) SetMoDeepCloner(v BulkMoDeepClonerRelationship)`

SetMoDeepCloner sets MoDeepCloner field to given value.

### HasMoDeepCloner

`func (o *BulkResult) HasMoDeepCloner() bool`

HasMoDeepCloner returns a boolean if a field has been set.

### SetMoDeepClonerNil

`func (o *BulkResult) SetMoDeepClonerNil(b bool)`

 SetMoDeepClonerNil sets the value for MoDeepCloner to be an explicit nil

### UnsetMoDeepCloner
`func (o *BulkResult) UnsetMoDeepCloner()`

UnsetMoDeepCloner ensures that no value is present for MoDeepCloner, not even an explicit nil
### GetMoMerger

`func (o *BulkResult) GetMoMerger() BulkMoMergerRelationship`

GetMoMerger returns the MoMerger field if non-nil, zero value otherwise.

### GetMoMergerOk

`func (o *BulkResult) GetMoMergerOk() (*BulkMoMergerRelationship, bool)`

GetMoMergerOk returns a tuple with the MoMerger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoMerger

`func (o *BulkResult) SetMoMerger(v BulkMoMergerRelationship)`

SetMoMerger sets MoMerger field to given value.

### HasMoMerger

`func (o *BulkResult) HasMoMerger() bool`

HasMoMerger returns a boolean if a field has been set.

### SetMoMergerNil

`func (o *BulkResult) SetMoMergerNil(b bool)`

 SetMoMergerNil sets the value for MoMerger to be an explicit nil

### UnsetMoMerger
`func (o *BulkResult) UnsetMoMerger()`

UnsetMoMerger ensures that no value is present for MoMerger, not even an explicit nil
### GetOrganization

`func (o *BulkResult) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkResult) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkResult) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkResult) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *BulkResult) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *BulkResult) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetResults

`func (o *BulkResult) GetResults() []BulkSubRequestObjRelationship`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkResult) GetResultsOk() (*[]BulkSubRequestObjRelationship, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkResult) SetResults(v []BulkSubRequestObjRelationship)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *BulkResult) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BulkResult) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetWorkflowInfo

`func (o *BulkResult) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *BulkResult) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *BulkResult) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *BulkResult) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.

### SetWorkflowInfoNil

`func (o *BulkResult) SetWorkflowInfoNil(b bool)`

 SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil

### UnsetWorkflowInfo
`func (o *BulkResult) UnsetWorkflowInfo()`

UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


