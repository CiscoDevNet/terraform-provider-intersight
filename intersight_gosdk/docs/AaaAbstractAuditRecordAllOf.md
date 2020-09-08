# AaaAbstractAuditRecordAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The operation that was performed on this Managed Object. The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type. | [optional] 
**MoDisplayNames** | Pointer to **interface{}** | The user-friendly names of the changed MO. | [optional] 
**MoType** | Pointer to **string** | The object type of the REST resource that was created, modified or deleted. | [optional] 
**ObjectMoid** | Pointer to **string** | The Moid of the REST resource that was created, modified or deleted. | [optional] 
**Request** | Pointer to **interface{}** | The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document. | [optional] 
**TraceId** | Pointer to **string** | The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team. | [optional] 

## Methods

### NewAaaAbstractAuditRecordAllOf

`func NewAaaAbstractAuditRecordAllOf() *AaaAbstractAuditRecordAllOf`

NewAaaAbstractAuditRecordAllOf instantiates a new AaaAbstractAuditRecordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaAbstractAuditRecordAllOfWithDefaults

`func NewAaaAbstractAuditRecordAllOfWithDefaults() *AaaAbstractAuditRecordAllOf`

NewAaaAbstractAuditRecordAllOfWithDefaults instantiates a new AaaAbstractAuditRecordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *AaaAbstractAuditRecordAllOf) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AaaAbstractAuditRecordAllOf) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AaaAbstractAuditRecordAllOf) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AaaAbstractAuditRecordAllOf) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMoDisplayNames

`func (o *AaaAbstractAuditRecordAllOf) GetMoDisplayNames() interface{}`

GetMoDisplayNames returns the MoDisplayNames field if non-nil, zero value otherwise.

### GetMoDisplayNamesOk

`func (o *AaaAbstractAuditRecordAllOf) GetMoDisplayNamesOk() (*interface{}, bool)`

GetMoDisplayNamesOk returns a tuple with the MoDisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoDisplayNames

`func (o *AaaAbstractAuditRecordAllOf) SetMoDisplayNames(v interface{})`

SetMoDisplayNames sets MoDisplayNames field to given value.

### HasMoDisplayNames

`func (o *AaaAbstractAuditRecordAllOf) HasMoDisplayNames() bool`

HasMoDisplayNames returns a boolean if a field has been set.

### SetMoDisplayNamesNil

`func (o *AaaAbstractAuditRecordAllOf) SetMoDisplayNamesNil(b bool)`

 SetMoDisplayNamesNil sets the value for MoDisplayNames to be an explicit nil

### UnsetMoDisplayNames
`func (o *AaaAbstractAuditRecordAllOf) UnsetMoDisplayNames()`

UnsetMoDisplayNames ensures that no value is present for MoDisplayNames, not even an explicit nil
### GetMoType

`func (o *AaaAbstractAuditRecordAllOf) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *AaaAbstractAuditRecordAllOf) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *AaaAbstractAuditRecordAllOf) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *AaaAbstractAuditRecordAllOf) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetObjectMoid

`func (o *AaaAbstractAuditRecordAllOf) GetObjectMoid() string`

GetObjectMoid returns the ObjectMoid field if non-nil, zero value otherwise.

### GetObjectMoidOk

`func (o *AaaAbstractAuditRecordAllOf) GetObjectMoidOk() (*string, bool)`

GetObjectMoidOk returns a tuple with the ObjectMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMoid

`func (o *AaaAbstractAuditRecordAllOf) SetObjectMoid(v string)`

SetObjectMoid sets ObjectMoid field to given value.

### HasObjectMoid

`func (o *AaaAbstractAuditRecordAllOf) HasObjectMoid() bool`

HasObjectMoid returns a boolean if a field has been set.

### GetRequest

`func (o *AaaAbstractAuditRecordAllOf) GetRequest() interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AaaAbstractAuditRecordAllOf) GetRequestOk() (*interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AaaAbstractAuditRecordAllOf) SetRequest(v interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AaaAbstractAuditRecordAllOf) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequestNil

`func (o *AaaAbstractAuditRecordAllOf) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *AaaAbstractAuditRecordAllOf) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetTraceId

`func (o *AaaAbstractAuditRecordAllOf) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *AaaAbstractAuditRecordAllOf) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *AaaAbstractAuditRecordAllOf) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *AaaAbstractAuditRecordAllOf) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


