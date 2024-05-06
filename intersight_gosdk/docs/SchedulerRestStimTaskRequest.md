# SchedulerRestStimTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.RestStimTaskRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.RestStimTaskRequest"]
**Body** | Pointer to **interface{}** | The request body that is sent as part of this API request. | [optional] 
**Headers** | Pointer to **interface{}** | Collection of key value pairs to set in the request header. | [optional] 
**Method** | Pointer to **string** | The supported values are POST, PUT, DELETE, PATCH. | [optional] 
**Protocol** | Pointer to **string** | The accepted web protocol values are http and https. | [optional] 
**Response** | Pointer to **interface{}** | The response obtained for the scheduled API service. | [optional] 
**Timeout** | Pointer to **int64** | Upper limit on the execution time of a scheduled task. Helps purge run-away scheduled tasks. | [optional] 
**Url** | Pointer to **string** | The URL of the resource in the target to which the API request is made. | [optional] 

## Methods

### NewSchedulerRestStimTaskRequest

`func NewSchedulerRestStimTaskRequest(classId string, objectType string, ) *SchedulerRestStimTaskRequest`

NewSchedulerRestStimTaskRequest instantiates a new SchedulerRestStimTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerRestStimTaskRequestWithDefaults

`func NewSchedulerRestStimTaskRequestWithDefaults() *SchedulerRestStimTaskRequest`

NewSchedulerRestStimTaskRequestWithDefaults instantiates a new SchedulerRestStimTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerRestStimTaskRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerRestStimTaskRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerRestStimTaskRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerRestStimTaskRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerRestStimTaskRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerRestStimTaskRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *SchedulerRestStimTaskRequest) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SchedulerRestStimTaskRequest) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SchedulerRestStimTaskRequest) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *SchedulerRestStimTaskRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SchedulerRestStimTaskRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SchedulerRestStimTaskRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetHeaders

`func (o *SchedulerRestStimTaskRequest) GetHeaders() interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SchedulerRestStimTaskRequest) GetHeadersOk() (*interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SchedulerRestStimTaskRequest) SetHeaders(v interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SchedulerRestStimTaskRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *SchedulerRestStimTaskRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *SchedulerRestStimTaskRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *SchedulerRestStimTaskRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SchedulerRestStimTaskRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SchedulerRestStimTaskRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SchedulerRestStimTaskRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetProtocol

`func (o *SchedulerRestStimTaskRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SchedulerRestStimTaskRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SchedulerRestStimTaskRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SchedulerRestStimTaskRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetResponse

`func (o *SchedulerRestStimTaskRequest) GetResponse() interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SchedulerRestStimTaskRequest) GetResponseOk() (*interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SchedulerRestStimTaskRequest) SetResponse(v interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SchedulerRestStimTaskRequest) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *SchedulerRestStimTaskRequest) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *SchedulerRestStimTaskRequest) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetTimeout

`func (o *SchedulerRestStimTaskRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SchedulerRestStimTaskRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SchedulerRestStimTaskRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SchedulerRestStimTaskRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *SchedulerRestStimTaskRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SchedulerRestStimTaskRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SchedulerRestStimTaskRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SchedulerRestStimTaskRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


