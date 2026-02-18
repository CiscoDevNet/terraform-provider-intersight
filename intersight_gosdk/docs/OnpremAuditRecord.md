# OnpremAuditRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "onprem.AuditRecord"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "onprem.AuditRecord"]
**Event** | Pointer to **string** | The type of event that occurred. Possible values are \&quot;Login\&quot;, \&quot;Logout\&quot;, \&quot;Created\&quot;, \&quot;Modified\&quot;, \&quot;Deleted\&quot;. | [optional] [readonly] 
**HttpOperation** | Pointer to **string** | It contains the http request type and URL for the operation. In case of authentication request, this field \&quot;POST /iam/adminlogin\&quot;. | [optional] [readonly] 
**HttpResponseCode** | Pointer to **int64** | The response code of the operation. If the operation is successful, this field will be 200. | [optional] [readonly] 
**HttpResponsePayload** | Pointer to **interface{}** | The body of the REST response that was sent to the client, represented as a JSON document. | [optional] [readonly] 
**MoType** | Pointer to **string** | The object type of the REST resource that was created, modified or deleted. | [optional] [readonly] 
**ObjectMoid** | Pointer to **string** | The Moid of the REST resource that was created, modified or deleted. | [optional] [readonly] 
**ObjectName** | Pointer to **interface{}** | The Name of the REST resource on which the operation was performed. | [optional] [readonly] 
**Request** | Pointer to **interface{}** | The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document. | [optional] [readonly] 
**SessionId** | Pointer to **string** | The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information. | [optional] [readonly] 
**SourceIp** | Pointer to **string** | The source IP of the client where the user action was performed. | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | The timestamp of the user action performed. | [optional] [readonly] 
**TraceId** | Pointer to **string** | The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team. | [optional] [readonly] 
**UserAgentString** | Pointer to **string** | The raw, string representation of the user agent of the request from the user-agent http request header. | [optional] [readonly] 
**UserIdOrEmail** | Pointer to **string** | The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information. | [optional] [readonly] 

## Methods

### NewOnpremAuditRecord

`func NewOnpremAuditRecord(classId string, objectType string, ) *OnpremAuditRecord`

NewOnpremAuditRecord instantiates a new OnpremAuditRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremAuditRecordWithDefaults

`func NewOnpremAuditRecordWithDefaults() *OnpremAuditRecord`

NewOnpremAuditRecordWithDefaults instantiates a new OnpremAuditRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OnpremAuditRecord) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OnpremAuditRecord) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OnpremAuditRecord) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OnpremAuditRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OnpremAuditRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OnpremAuditRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEvent

`func (o *OnpremAuditRecord) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OnpremAuditRecord) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OnpremAuditRecord) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OnpremAuditRecord) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetHttpOperation

`func (o *OnpremAuditRecord) GetHttpOperation() string`

GetHttpOperation returns the HttpOperation field if non-nil, zero value otherwise.

### GetHttpOperationOk

`func (o *OnpremAuditRecord) GetHttpOperationOk() (*string, bool)`

GetHttpOperationOk returns a tuple with the HttpOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperation

`func (o *OnpremAuditRecord) SetHttpOperation(v string)`

SetHttpOperation sets HttpOperation field to given value.

### HasHttpOperation

`func (o *OnpremAuditRecord) HasHttpOperation() bool`

HasHttpOperation returns a boolean if a field has been set.

### GetHttpResponseCode

`func (o *OnpremAuditRecord) GetHttpResponseCode() int64`

GetHttpResponseCode returns the HttpResponseCode field if non-nil, zero value otherwise.

### GetHttpResponseCodeOk

`func (o *OnpremAuditRecord) GetHttpResponseCodeOk() (*int64, bool)`

GetHttpResponseCodeOk returns a tuple with the HttpResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseCode

`func (o *OnpremAuditRecord) SetHttpResponseCode(v int64)`

SetHttpResponseCode sets HttpResponseCode field to given value.

### HasHttpResponseCode

`func (o *OnpremAuditRecord) HasHttpResponseCode() bool`

HasHttpResponseCode returns a boolean if a field has been set.

### GetHttpResponsePayload

`func (o *OnpremAuditRecord) GetHttpResponsePayload() interface{}`

GetHttpResponsePayload returns the HttpResponsePayload field if non-nil, zero value otherwise.

### GetHttpResponsePayloadOk

`func (o *OnpremAuditRecord) GetHttpResponsePayloadOk() (*interface{}, bool)`

GetHttpResponsePayloadOk returns a tuple with the HttpResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponsePayload

`func (o *OnpremAuditRecord) SetHttpResponsePayload(v interface{})`

SetHttpResponsePayload sets HttpResponsePayload field to given value.

### HasHttpResponsePayload

`func (o *OnpremAuditRecord) HasHttpResponsePayload() bool`

HasHttpResponsePayload returns a boolean if a field has been set.

### SetHttpResponsePayloadNil

`func (o *OnpremAuditRecord) SetHttpResponsePayloadNil(b bool)`

 SetHttpResponsePayloadNil sets the value for HttpResponsePayload to be an explicit nil

### UnsetHttpResponsePayload
`func (o *OnpremAuditRecord) UnsetHttpResponsePayload()`

UnsetHttpResponsePayload ensures that no value is present for HttpResponsePayload, not even an explicit nil
### GetMoType

`func (o *OnpremAuditRecord) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *OnpremAuditRecord) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *OnpremAuditRecord) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *OnpremAuditRecord) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetObjectMoid

`func (o *OnpremAuditRecord) GetObjectMoid() string`

GetObjectMoid returns the ObjectMoid field if non-nil, zero value otherwise.

### GetObjectMoidOk

`func (o *OnpremAuditRecord) GetObjectMoidOk() (*string, bool)`

GetObjectMoidOk returns a tuple with the ObjectMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMoid

`func (o *OnpremAuditRecord) SetObjectMoid(v string)`

SetObjectMoid sets ObjectMoid field to given value.

### HasObjectMoid

`func (o *OnpremAuditRecord) HasObjectMoid() bool`

HasObjectMoid returns a boolean if a field has been set.

### GetObjectName

`func (o *OnpremAuditRecord) GetObjectName() interface{}`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *OnpremAuditRecord) GetObjectNameOk() (*interface{}, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *OnpremAuditRecord) SetObjectName(v interface{})`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *OnpremAuditRecord) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *OnpremAuditRecord) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *OnpremAuditRecord) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetRequest

`func (o *OnpremAuditRecord) GetRequest() interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OnpremAuditRecord) GetRequestOk() (*interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OnpremAuditRecord) SetRequest(v interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OnpremAuditRecord) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *OnpremAuditRecord) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *OnpremAuditRecord) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetSessionId

`func (o *OnpremAuditRecord) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *OnpremAuditRecord) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *OnpremAuditRecord) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *OnpremAuditRecord) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceIp

`func (o *OnpremAuditRecord) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *OnpremAuditRecord) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *OnpremAuditRecord) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *OnpremAuditRecord) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetTimestamp

`func (o *OnpremAuditRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OnpremAuditRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OnpremAuditRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OnpremAuditRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *OnpremAuditRecord) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *OnpremAuditRecord) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *OnpremAuditRecord) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *OnpremAuditRecord) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUserAgentString

`func (o *OnpremAuditRecord) GetUserAgentString() string`

GetUserAgentString returns the UserAgentString field if non-nil, zero value otherwise.

### GetUserAgentStringOk

`func (o *OnpremAuditRecord) GetUserAgentStringOk() (*string, bool)`

GetUserAgentStringOk returns a tuple with the UserAgentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentString

`func (o *OnpremAuditRecord) SetUserAgentString(v string)`

SetUserAgentString sets UserAgentString field to given value.

### HasUserAgentString

`func (o *OnpremAuditRecord) HasUserAgentString() bool`

HasUserAgentString returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *OnpremAuditRecord) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *OnpremAuditRecord) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *OnpremAuditRecord) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *OnpremAuditRecord) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


