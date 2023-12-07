# AaaAbstractAuditRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "aaa.AuditRecord"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "aaa.AuditRecord"]
**Event** | Pointer to **string** | The operation that was performed on this Managed Object. The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type. | [optional] [readonly] 
**MoDisplayNames** | Pointer to **interface{}** | The user-friendly names of the changed MO. | [optional] [readonly] 
**MoType** | Pointer to **string** | The object type of the REST resource that was created, modified or deleted. | [optional] [readonly] 
**ObjectMoid** | Pointer to **string** | The Moid of the REST resource that was created, modified or deleted. | [optional] [readonly] 
**Request** | Pointer to **interface{}** | The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document. | [optional] [readonly] 
**TraceId** | Pointer to **string** | The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team. | [optional] [readonly] 
**UserAgent** | Pointer to [**NullableAaaUserAgent**](AaaUserAgent.md) |  | [optional] 
**UserAgentString** | Pointer to **string** | The raw, string representation of the user agent of the request from the user-agent http request header. | [optional] [readonly] 

## Methods

### NewAaaAbstractAuditRecord

`func NewAaaAbstractAuditRecord(classId string, objectType string, ) *AaaAbstractAuditRecord`

NewAaaAbstractAuditRecord instantiates a new AaaAbstractAuditRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaAbstractAuditRecordWithDefaults

`func NewAaaAbstractAuditRecordWithDefaults() *AaaAbstractAuditRecord`

NewAaaAbstractAuditRecordWithDefaults instantiates a new AaaAbstractAuditRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AaaAbstractAuditRecord) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AaaAbstractAuditRecord) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AaaAbstractAuditRecord) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AaaAbstractAuditRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AaaAbstractAuditRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AaaAbstractAuditRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEvent

`func (o *AaaAbstractAuditRecord) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AaaAbstractAuditRecord) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AaaAbstractAuditRecord) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AaaAbstractAuditRecord) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMoDisplayNames

`func (o *AaaAbstractAuditRecord) GetMoDisplayNames() interface{}`

GetMoDisplayNames returns the MoDisplayNames field if non-nil, zero value otherwise.

### GetMoDisplayNamesOk

`func (o *AaaAbstractAuditRecord) GetMoDisplayNamesOk() (*interface{}, bool)`

GetMoDisplayNamesOk returns a tuple with the MoDisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoDisplayNames

`func (o *AaaAbstractAuditRecord) SetMoDisplayNames(v interface{})`

SetMoDisplayNames sets MoDisplayNames field to given value.

### HasMoDisplayNames

`func (o *AaaAbstractAuditRecord) HasMoDisplayNames() bool`

HasMoDisplayNames returns a boolean if a field has been set.

### SetMoDisplayNamesNil

`func (o *AaaAbstractAuditRecord) SetMoDisplayNamesNil(b bool)`

 SetMoDisplayNamesNil sets the value for MoDisplayNames to be an explicit nil

### UnsetMoDisplayNames
`func (o *AaaAbstractAuditRecord) UnsetMoDisplayNames()`

UnsetMoDisplayNames ensures that no value is present for MoDisplayNames, not even an explicit nil
### GetMoType

`func (o *AaaAbstractAuditRecord) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *AaaAbstractAuditRecord) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *AaaAbstractAuditRecord) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *AaaAbstractAuditRecord) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetObjectMoid

`func (o *AaaAbstractAuditRecord) GetObjectMoid() string`

GetObjectMoid returns the ObjectMoid field if non-nil, zero value otherwise.

### GetObjectMoidOk

`func (o *AaaAbstractAuditRecord) GetObjectMoidOk() (*string, bool)`

GetObjectMoidOk returns a tuple with the ObjectMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMoid

`func (o *AaaAbstractAuditRecord) SetObjectMoid(v string)`

SetObjectMoid sets ObjectMoid field to given value.

### HasObjectMoid

`func (o *AaaAbstractAuditRecord) HasObjectMoid() bool`

HasObjectMoid returns a boolean if a field has been set.

### GetRequest

`func (o *AaaAbstractAuditRecord) GetRequest() interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AaaAbstractAuditRecord) GetRequestOk() (*interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AaaAbstractAuditRecord) SetRequest(v interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AaaAbstractAuditRecord) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *AaaAbstractAuditRecord) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *AaaAbstractAuditRecord) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetTraceId

`func (o *AaaAbstractAuditRecord) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *AaaAbstractAuditRecord) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *AaaAbstractAuditRecord) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *AaaAbstractAuditRecord) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetUserAgent

`func (o *AaaAbstractAuditRecord) GetUserAgent() AaaUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AaaAbstractAuditRecord) GetUserAgentOk() (*AaaUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AaaAbstractAuditRecord) SetUserAgent(v AaaUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AaaAbstractAuditRecord) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *AaaAbstractAuditRecord) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *AaaAbstractAuditRecord) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetUserAgentString

`func (o *AaaAbstractAuditRecord) GetUserAgentString() string`

GetUserAgentString returns the UserAgentString field if non-nil, zero value otherwise.

### GetUserAgentStringOk

`func (o *AaaAbstractAuditRecord) GetUserAgentStringOk() (*string, bool)`

GetUserAgentStringOk returns a tuple with the UserAgentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentString

`func (o *AaaAbstractAuditRecord) SetUserAgentString(v string)`

SetUserAgentString sets UserAgentString field to given value.

### HasUserAgentString

`func (o *AaaAbstractAuditRecord) HasUserAgentString() bool`

HasUserAgentString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


