# BulkRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.Request"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.Request"]
**ActionOnError** | Pointer to **string** | The action to be taken when an error occurs during processing of the request. * &#x60;Stop&#x60; - Stop the processing of the request after the first error. * &#x60;Proceed&#x60; - Proceed with the processing of the request even when an error occurs. | [optional] [default to "Stop"]
**Actions** | Pointer to **[]string** |  | [optional] 
**CompletionTime** | Pointer to **string** | The timestamp when the request processing completed. | [optional] [readonly] 
**Headers** | Pointer to [**[]BulkHttpHeader**](BulkHttpHeader.md) |  | [optional] 
**NumSubRequests** | Pointer to **int64** | The number of sub requests received in this request. | [optional] [readonly] 
**OrgMoid** | Pointer to **string** | The moid of the organization under which this request was issued. | [optional] [readonly] 
**RequestReceivedTime** | Pointer to **string** | The timestamp when the request was received. | [optional] [readonly] 
**Requests** | Pointer to [**[]BulkSubRequest**](BulkSubRequest.md) |  | [optional] 
**Results** | Pointer to [**[]BulkApiResult**](BulkApiResult.md) |  | [optional] 
**SkipDuplicates** | Pointer to **bool** | Skip the already present objects. | [optional] 
**Status** | Pointer to **string** | The processing status of the Request. * &#x60;NotStarted&#x60; - Indicates that the request processing has not begun yet. * &#x60;ObjPresenceCheckInProgress&#x60; - Indicates that the object presence check is in progress for this request. * &#x60;ObjPresenceCheckComplete&#x60; - Indicates that the object presence check is complete. * &#x60;ExecutionInProgress&#x60; - Indicates that the request processing is in progress. * &#x60;Completed&#x60; - Indicates that the request processing has been completed successfully. * &#x60;Failed&#x60; - Indicates that the processing of this request failed. | [optional] [readonly] [default to "NotStarted"]
**StatusMessage** | Pointer to **string** | The status message corresponding to the status. | [optional] [readonly] 
**Uri** | Pointer to **string** | The URI on which this bulk action is to be performed. The value will be used when there is no override in the SubRequest. | [optional] 
**Verb** | Pointer to **string** | The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value will be used when there is no override in the SubRequest. * &#x60;POST&#x60; - Used to create a REST resource. * &#x60;PATCH&#x60; - Used to update a REST resource. * &#x60;DELETE&#x60; - Used to delete a REST resource. | [optional] [default to "POST"]
**AsyncResults** | Pointer to [**[]BulkSubRequestObjRelationship**](BulkSubRequestObjRelationship.md) | An array of relationships to bulkSubRequestObj resources. | [optional] [readonly] 
**AsyncResultsFailed** | Pointer to [**[]BulkSubRequestObjRelationship**](BulkSubRequestObjRelationship.md) | An array of relationships to bulkSubRequestObj resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewBulkRequestAllOf

`func NewBulkRequestAllOf(classId string, objectType string, ) *BulkRequestAllOf`

NewBulkRequestAllOf instantiates a new BulkRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRequestAllOfWithDefaults

`func NewBulkRequestAllOfWithDefaults() *BulkRequestAllOf`

NewBulkRequestAllOfWithDefaults instantiates a new BulkRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionOnError

`func (o *BulkRequestAllOf) GetActionOnError() string`

GetActionOnError returns the ActionOnError field if non-nil, zero value otherwise.

### GetActionOnErrorOk

`func (o *BulkRequestAllOf) GetActionOnErrorOk() (*string, bool)`

GetActionOnErrorOk returns a tuple with the ActionOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnError

`func (o *BulkRequestAllOf) SetActionOnError(v string)`

SetActionOnError sets ActionOnError field to given value.

### HasActionOnError

`func (o *BulkRequestAllOf) HasActionOnError() bool`

HasActionOnError returns a boolean if a field has been set.

### GetActions

`func (o *BulkRequestAllOf) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BulkRequestAllOf) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BulkRequestAllOf) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BulkRequestAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *BulkRequestAllOf) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *BulkRequestAllOf) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetCompletionTime

`func (o *BulkRequestAllOf) GetCompletionTime() string`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *BulkRequestAllOf) GetCompletionTimeOk() (*string, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *BulkRequestAllOf) SetCompletionTime(v string)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *BulkRequestAllOf) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetHeaders

`func (o *BulkRequestAllOf) GetHeaders() []BulkHttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BulkRequestAllOf) GetHeadersOk() (*[]BulkHttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BulkRequestAllOf) SetHeaders(v []BulkHttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BulkRequestAllOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *BulkRequestAllOf) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *BulkRequestAllOf) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetNumSubRequests

`func (o *BulkRequestAllOf) GetNumSubRequests() int64`

GetNumSubRequests returns the NumSubRequests field if non-nil, zero value otherwise.

### GetNumSubRequestsOk

`func (o *BulkRequestAllOf) GetNumSubRequestsOk() (*int64, bool)`

GetNumSubRequestsOk returns a tuple with the NumSubRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubRequests

`func (o *BulkRequestAllOf) SetNumSubRequests(v int64)`

SetNumSubRequests sets NumSubRequests field to given value.

### HasNumSubRequests

`func (o *BulkRequestAllOf) HasNumSubRequests() bool`

HasNumSubRequests returns a boolean if a field has been set.

### GetOrgMoid

`func (o *BulkRequestAllOf) GetOrgMoid() string`

GetOrgMoid returns the OrgMoid field if non-nil, zero value otherwise.

### GetOrgMoidOk

`func (o *BulkRequestAllOf) GetOrgMoidOk() (*string, bool)`

GetOrgMoidOk returns a tuple with the OrgMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMoid

`func (o *BulkRequestAllOf) SetOrgMoid(v string)`

SetOrgMoid sets OrgMoid field to given value.

### HasOrgMoid

`func (o *BulkRequestAllOf) HasOrgMoid() bool`

HasOrgMoid returns a boolean if a field has been set.

### GetRequestReceivedTime

`func (o *BulkRequestAllOf) GetRequestReceivedTime() string`

GetRequestReceivedTime returns the RequestReceivedTime field if non-nil, zero value otherwise.

### GetRequestReceivedTimeOk

`func (o *BulkRequestAllOf) GetRequestReceivedTimeOk() (*string, bool)`

GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReceivedTime

`func (o *BulkRequestAllOf) SetRequestReceivedTime(v string)`

SetRequestReceivedTime sets RequestReceivedTime field to given value.

### HasRequestReceivedTime

`func (o *BulkRequestAllOf) HasRequestReceivedTime() bool`

HasRequestReceivedTime returns a boolean if a field has been set.

### GetRequests

`func (o *BulkRequestAllOf) GetRequests() []BulkSubRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BulkRequestAllOf) GetRequestsOk() (*[]BulkSubRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BulkRequestAllOf) SetRequests(v []BulkSubRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BulkRequestAllOf) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequestsNil

`func (o *BulkRequestAllOf) SetRequestsNil(b bool)`

 SetRequestsNil sets the value for Requests to be an explicit nil

### UnsetRequests
`func (o *BulkRequestAllOf) UnsetRequests()`

UnsetRequests ensures that no value is present for Requests, not even an explicit nil
### GetResults

`func (o *BulkRequestAllOf) GetResults() []BulkApiResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkRequestAllOf) GetResultsOk() (*[]BulkApiResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkRequestAllOf) SetResults(v []BulkApiResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkRequestAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *BulkRequestAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *BulkRequestAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSkipDuplicates

`func (o *BulkRequestAllOf) GetSkipDuplicates() bool`

GetSkipDuplicates returns the SkipDuplicates field if non-nil, zero value otherwise.

### GetSkipDuplicatesOk

`func (o *BulkRequestAllOf) GetSkipDuplicatesOk() (*bool, bool)`

GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDuplicates

`func (o *BulkRequestAllOf) SetSkipDuplicates(v bool)`

SetSkipDuplicates sets SkipDuplicates field to given value.

### HasSkipDuplicates

`func (o *BulkRequestAllOf) HasSkipDuplicates() bool`

HasSkipDuplicates returns a boolean if a field has been set.

### GetStatus

`func (o *BulkRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *BulkRequestAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BulkRequestAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BulkRequestAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *BulkRequestAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUri

`func (o *BulkRequestAllOf) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkRequestAllOf) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkRequestAllOf) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkRequestAllOf) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVerb

`func (o *BulkRequestAllOf) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *BulkRequestAllOf) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *BulkRequestAllOf) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *BulkRequestAllOf) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetAsyncResults

`func (o *BulkRequestAllOf) GetAsyncResults() []BulkSubRequestObjRelationship`

GetAsyncResults returns the AsyncResults field if non-nil, zero value otherwise.

### GetAsyncResultsOk

`func (o *BulkRequestAllOf) GetAsyncResultsOk() (*[]BulkSubRequestObjRelationship, bool)`

GetAsyncResultsOk returns a tuple with the AsyncResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResults

`func (o *BulkRequestAllOf) SetAsyncResults(v []BulkSubRequestObjRelationship)`

SetAsyncResults sets AsyncResults field to given value.

### HasAsyncResults

`func (o *BulkRequestAllOf) HasAsyncResults() bool`

HasAsyncResults returns a boolean if a field has been set.

### SetAsyncResultsNil

`func (o *BulkRequestAllOf) SetAsyncResultsNil(b bool)`

 SetAsyncResultsNil sets the value for AsyncResults to be an explicit nil

### UnsetAsyncResults
`func (o *BulkRequestAllOf) UnsetAsyncResults()`

UnsetAsyncResults ensures that no value is present for AsyncResults, not even an explicit nil
### GetAsyncResultsFailed

`func (o *BulkRequestAllOf) GetAsyncResultsFailed() []BulkSubRequestObjRelationship`

GetAsyncResultsFailed returns the AsyncResultsFailed field if non-nil, zero value otherwise.

### GetAsyncResultsFailedOk

`func (o *BulkRequestAllOf) GetAsyncResultsFailedOk() (*[]BulkSubRequestObjRelationship, bool)`

GetAsyncResultsFailedOk returns a tuple with the AsyncResultsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncResultsFailed

`func (o *BulkRequestAllOf) SetAsyncResultsFailed(v []BulkSubRequestObjRelationship)`

SetAsyncResultsFailed sets AsyncResultsFailed field to given value.

### HasAsyncResultsFailed

`func (o *BulkRequestAllOf) HasAsyncResultsFailed() bool`

HasAsyncResultsFailed returns a boolean if a field has been set.

### SetAsyncResultsFailedNil

`func (o *BulkRequestAllOf) SetAsyncResultsFailedNil(b bool)`

 SetAsyncResultsFailedNil sets the value for AsyncResultsFailed to be an explicit nil

### UnsetAsyncResultsFailed
`func (o *BulkRequestAllOf) UnsetAsyncResultsFailed()`

UnsetAsyncResultsFailed ensures that no value is present for AsyncResultsFailed, not even an explicit nil
### GetOrganization

`func (o *BulkRequestAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkRequestAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *BulkRequestAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *BulkRequestAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *BulkRequestAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *BulkRequestAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


