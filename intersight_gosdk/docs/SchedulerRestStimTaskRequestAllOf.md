# SchedulerRestStimTaskRequestAllOf

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

### NewSchedulerRestStimTaskRequestAllOf

`func NewSchedulerRestStimTaskRequestAllOf(classId string, objectType string, ) *SchedulerRestStimTaskRequestAllOf`

NewSchedulerRestStimTaskRequestAllOf instantiates a new SchedulerRestStimTaskRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerRestStimTaskRequestAllOfWithDefaults

`func NewSchedulerRestStimTaskRequestAllOfWithDefaults() *SchedulerRestStimTaskRequestAllOf`

NewSchedulerRestStimTaskRequestAllOfWithDefaults instantiates a new SchedulerRestStimTaskRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerRestStimTaskRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerRestStimTaskRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerRestStimTaskRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerRestStimTaskRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *SchedulerRestStimTaskRequestAllOf) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SchedulerRestStimTaskRequestAllOf) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *SchedulerRestStimTaskRequestAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SchedulerRestStimTaskRequestAllOf) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SchedulerRestStimTaskRequestAllOf) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetHeaders

`func (o *SchedulerRestStimTaskRequestAllOf) GetHeaders() interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetHeadersOk() (*interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SchedulerRestStimTaskRequestAllOf) SetHeaders(v interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SchedulerRestStimTaskRequestAllOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *SchedulerRestStimTaskRequestAllOf) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *SchedulerRestStimTaskRequestAllOf) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *SchedulerRestStimTaskRequestAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SchedulerRestStimTaskRequestAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SchedulerRestStimTaskRequestAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetProtocol

`func (o *SchedulerRestStimTaskRequestAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SchedulerRestStimTaskRequestAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SchedulerRestStimTaskRequestAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetResponse

`func (o *SchedulerRestStimTaskRequestAllOf) GetResponse() interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetResponseOk() (*interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SchedulerRestStimTaskRequestAllOf) SetResponse(v interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SchedulerRestStimTaskRequestAllOf) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *SchedulerRestStimTaskRequestAllOf) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *SchedulerRestStimTaskRequestAllOf) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetTimeout

`func (o *SchedulerRestStimTaskRequestAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SchedulerRestStimTaskRequestAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SchedulerRestStimTaskRequestAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *SchedulerRestStimTaskRequestAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SchedulerRestStimTaskRequestAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SchedulerRestStimTaskRequestAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SchedulerRestStimTaskRequestAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


