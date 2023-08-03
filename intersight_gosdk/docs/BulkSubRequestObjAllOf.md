# BulkSubRequestObjAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.SubRequestObj"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.SubRequestObj"]
**Body** | Pointer to [**MoBaseMo**](MoBaseMo.md) |  | [optional] 
**BodyString** | Pointer to **string** | The body of the sub-request in string format. | [optional] [readonly] 
**ExecutionCompletionTime** | Pointer to **string** | The time at which processing of this request completed. | [optional] [readonly] 
**ExecutionStartTime** | Pointer to **string** | The time at which processing of this request started. | [optional] [readonly] 
**IsBulkMoOp** | Pointer to **bool** | For Async Bulk Mo Operations this flag will be set to true. | [optional] 
**IsObjectPresent** | Pointer to **bool** | This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence. | [optional] [readonly] 
**Result** | Pointer to [**NullableBulkApiResult**](BulkApiResult.md) |  | [optional] 
**SkipDuplicates** | Pointer to **bool** | Skip the already present objects. The value from the Request. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the request. * &#x60;Pending&#x60; - Indicates that the request is yet to be processed. * &#x60;ObjPresenceCheckInProgress&#x60; - Indicates that the checking for object presence is in progress. * &#x60;ObjPresenceCheckInComplete&#x60; - Indicates that the request is being processed. * &#x60;ObjPresenceCheckFailed&#x60; - Indicates that the checking for object presence failed. * &#x60;Processing&#x60; - Indicates that the request is being processed. * &#x60;TimedOut&#x60; - Indicates that the request processing timed out. * &#x60;Failed&#x60; - Indicates that the request processing failed. * &#x60;Completed&#x60; - Indicates that the request processing is complete. * &#x60;Skipped&#x60; - Indicates that the request was skipped. | [optional] [readonly] [default to "Pending"]
**SystemDefinedObjectDetected** | Pointer to **bool** | This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence. | [optional] [readonly] 
**TargetMoid** | Pointer to **string** | Used with PATCH &amp; DELETE actions. The moid of an existing object instance. | [optional] 
**Uri** | Pointer to **string** | The URI on which this bulk action is to be performed. | [optional] 
**Verb** | Pointer to **string** | The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * &#x60;POST&#x60; - Used to create a REST resource. * &#x60;PATCH&#x60; - Used to update a REST resource. * &#x60;DELETE&#x60; - Used to delete a REST resource. | [optional] [default to "POST"]
**AsyncRequest** | Pointer to [**BulkResultRelationship**](BulkResultRelationship.md) |  | [optional] 
**Request** | Pointer to [**BulkRequestRelationship**](BulkRequestRelationship.md) |  | [optional] 

## Methods

### NewBulkSubRequestObjAllOf

`func NewBulkSubRequestObjAllOf(classId string, objectType string, ) *BulkSubRequestObjAllOf`

NewBulkSubRequestObjAllOf instantiates a new BulkSubRequestObjAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkSubRequestObjAllOfWithDefaults

`func NewBulkSubRequestObjAllOfWithDefaults() *BulkSubRequestObjAllOf`

NewBulkSubRequestObjAllOfWithDefaults instantiates a new BulkSubRequestObjAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkSubRequestObjAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkSubRequestObjAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkSubRequestObjAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkSubRequestObjAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkSubRequestObjAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkSubRequestObjAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *BulkSubRequestObjAllOf) GetBody() MoBaseMo`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BulkSubRequestObjAllOf) GetBodyOk() (*MoBaseMo, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BulkSubRequestObjAllOf) SetBody(v MoBaseMo)`

SetBody sets Body field to given value.

### HasBody

`func (o *BulkSubRequestObjAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyString

`func (o *BulkSubRequestObjAllOf) GetBodyString() string`

GetBodyString returns the BodyString field if non-nil, zero value otherwise.

### GetBodyStringOk

`func (o *BulkSubRequestObjAllOf) GetBodyStringOk() (*string, bool)`

GetBodyStringOk returns a tuple with the BodyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyString

`func (o *BulkSubRequestObjAllOf) SetBodyString(v string)`

SetBodyString sets BodyString field to given value.

### HasBodyString

`func (o *BulkSubRequestObjAllOf) HasBodyString() bool`

HasBodyString returns a boolean if a field has been set.

### GetExecutionCompletionTime

`func (o *BulkSubRequestObjAllOf) GetExecutionCompletionTime() string`

GetExecutionCompletionTime returns the ExecutionCompletionTime field if non-nil, zero value otherwise.

### GetExecutionCompletionTimeOk

`func (o *BulkSubRequestObjAllOf) GetExecutionCompletionTimeOk() (*string, bool)`

GetExecutionCompletionTimeOk returns a tuple with the ExecutionCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCompletionTime

`func (o *BulkSubRequestObjAllOf) SetExecutionCompletionTime(v string)`

SetExecutionCompletionTime sets ExecutionCompletionTime field to given value.

### HasExecutionCompletionTime

`func (o *BulkSubRequestObjAllOf) HasExecutionCompletionTime() bool`

HasExecutionCompletionTime returns a boolean if a field has been set.

### GetExecutionStartTime

`func (o *BulkSubRequestObjAllOf) GetExecutionStartTime() string`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *BulkSubRequestObjAllOf) GetExecutionStartTimeOk() (*string, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *BulkSubRequestObjAllOf) SetExecutionStartTime(v string)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *BulkSubRequestObjAllOf) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### GetIsBulkMoOp

`func (o *BulkSubRequestObjAllOf) GetIsBulkMoOp() bool`

GetIsBulkMoOp returns the IsBulkMoOp field if non-nil, zero value otherwise.

### GetIsBulkMoOpOk

`func (o *BulkSubRequestObjAllOf) GetIsBulkMoOpOk() (*bool, bool)`

GetIsBulkMoOpOk returns a tuple with the IsBulkMoOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBulkMoOp

`func (o *BulkSubRequestObjAllOf) SetIsBulkMoOp(v bool)`

SetIsBulkMoOp sets IsBulkMoOp field to given value.

### HasIsBulkMoOp

`func (o *BulkSubRequestObjAllOf) HasIsBulkMoOp() bool`

HasIsBulkMoOp returns a boolean if a field has been set.

### GetIsObjectPresent

`func (o *BulkSubRequestObjAllOf) GetIsObjectPresent() bool`

GetIsObjectPresent returns the IsObjectPresent field if non-nil, zero value otherwise.

### GetIsObjectPresentOk

`func (o *BulkSubRequestObjAllOf) GetIsObjectPresentOk() (*bool, bool)`

GetIsObjectPresentOk returns a tuple with the IsObjectPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObjectPresent

`func (o *BulkSubRequestObjAllOf) SetIsObjectPresent(v bool)`

SetIsObjectPresent sets IsObjectPresent field to given value.

### HasIsObjectPresent

`func (o *BulkSubRequestObjAllOf) HasIsObjectPresent() bool`

HasIsObjectPresent returns a boolean if a field has been set.

### GetResult

`func (o *BulkSubRequestObjAllOf) GetResult() BulkApiResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BulkSubRequestObjAllOf) GetResultOk() (*BulkApiResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BulkSubRequestObjAllOf) SetResult(v BulkApiResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *BulkSubRequestObjAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *BulkSubRequestObjAllOf) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *BulkSubRequestObjAllOf) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetSkipDuplicates

`func (o *BulkSubRequestObjAllOf) GetSkipDuplicates() bool`

GetSkipDuplicates returns the SkipDuplicates field if non-nil, zero value otherwise.

### GetSkipDuplicatesOk

`func (o *BulkSubRequestObjAllOf) GetSkipDuplicatesOk() (*bool, bool)`

GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDuplicates

`func (o *BulkSubRequestObjAllOf) SetSkipDuplicates(v bool)`

SetSkipDuplicates sets SkipDuplicates field to given value.

### HasSkipDuplicates

`func (o *BulkSubRequestObjAllOf) HasSkipDuplicates() bool`

HasSkipDuplicates returns a boolean if a field has been set.

### GetStatus

`func (o *BulkSubRequestObjAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkSubRequestObjAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkSubRequestObjAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkSubRequestObjAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemDefinedObjectDetected

`func (o *BulkSubRequestObjAllOf) GetSystemDefinedObjectDetected() bool`

GetSystemDefinedObjectDetected returns the SystemDefinedObjectDetected field if non-nil, zero value otherwise.

### GetSystemDefinedObjectDetectedOk

`func (o *BulkSubRequestObjAllOf) GetSystemDefinedObjectDetectedOk() (*bool, bool)`

GetSystemDefinedObjectDetectedOk returns a tuple with the SystemDefinedObjectDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefinedObjectDetected

`func (o *BulkSubRequestObjAllOf) SetSystemDefinedObjectDetected(v bool)`

SetSystemDefinedObjectDetected sets SystemDefinedObjectDetected field to given value.

### HasSystemDefinedObjectDetected

`func (o *BulkSubRequestObjAllOf) HasSystemDefinedObjectDetected() bool`

HasSystemDefinedObjectDetected returns a boolean if a field has been set.

### GetTargetMoid

`func (o *BulkSubRequestObjAllOf) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *BulkSubRequestObjAllOf) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *BulkSubRequestObjAllOf) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *BulkSubRequestObjAllOf) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.

### GetUri

`func (o *BulkSubRequestObjAllOf) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkSubRequestObjAllOf) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkSubRequestObjAllOf) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkSubRequestObjAllOf) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVerb

`func (o *BulkSubRequestObjAllOf) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *BulkSubRequestObjAllOf) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *BulkSubRequestObjAllOf) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *BulkSubRequestObjAllOf) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetAsyncRequest

`func (o *BulkSubRequestObjAllOf) GetAsyncRequest() BulkResultRelationship`

GetAsyncRequest returns the AsyncRequest field if non-nil, zero value otherwise.

### GetAsyncRequestOk

`func (o *BulkSubRequestObjAllOf) GetAsyncRequestOk() (*BulkResultRelationship, bool)`

GetAsyncRequestOk returns a tuple with the AsyncRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncRequest

`func (o *BulkSubRequestObjAllOf) SetAsyncRequest(v BulkResultRelationship)`

SetAsyncRequest sets AsyncRequest field to given value.

### HasAsyncRequest

`func (o *BulkSubRequestObjAllOf) HasAsyncRequest() bool`

HasAsyncRequest returns a boolean if a field has been set.

### GetRequest

`func (o *BulkSubRequestObjAllOf) GetRequest() BulkRequestRelationship`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *BulkSubRequestObjAllOf) GetRequestOk() (*BulkRequestRelationship, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *BulkSubRequestObjAllOf) SetRequest(v BulkRequestRelationship)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *BulkSubRequestObjAllOf) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


