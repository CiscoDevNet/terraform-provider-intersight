# ConnectorHttpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Contents of the request body to send for PUT/PATCH/POST requests. | [optional] 
**DialTimeout** | Pointer to **int64** | The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used. | [optional] 
**Header** | Pointer to **interface{}** | Collection of key value pairs to set in the request header. | [optional] 
**Internal** | Pointer to **bool** | The request is for an internal platform API that requires authentication to be inserted by the platform implementation. | [optional] 
**Method** | Pointer to **string** | Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used. | [optional] 
**Url** | Pointer to [**ConnectorUrl**](connector.Url.md) |  | [optional] 

## Methods

### NewConnectorHttpRequest

`func NewConnectorHttpRequest() *ConnectorHttpRequest`

NewConnectorHttpRequest instantiates a new ConnectorHttpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorHttpRequestWithDefaults

`func NewConnectorHttpRequestWithDefaults() *ConnectorHttpRequest`

NewConnectorHttpRequestWithDefaults instantiates a new ConnectorHttpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ConnectorHttpRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConnectorHttpRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConnectorHttpRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ConnectorHttpRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDialTimeout

`func (o *ConnectorHttpRequest) GetDialTimeout() int64`

GetDialTimeout returns the DialTimeout field if non-nil, zero value otherwise.

### GetDialTimeoutOk

`func (o *ConnectorHttpRequest) GetDialTimeoutOk() (*int64, bool)`

GetDialTimeoutOk returns a tuple with the DialTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialTimeout

`func (o *ConnectorHttpRequest) SetDialTimeout(v int64)`

SetDialTimeout sets DialTimeout field to given value.

### HasDialTimeout

`func (o *ConnectorHttpRequest) HasDialTimeout() bool`

HasDialTimeout returns a boolean if a field has been set.

### GetHeader

`func (o *ConnectorHttpRequest) GetHeader() interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ConnectorHttpRequest) GetHeaderOk() (*interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ConnectorHttpRequest) SetHeader(v interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ConnectorHttpRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ConnectorHttpRequest) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ConnectorHttpRequest) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetInternal

`func (o *ConnectorHttpRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ConnectorHttpRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ConnectorHttpRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ConnectorHttpRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMethod

`func (o *ConnectorHttpRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConnectorHttpRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConnectorHttpRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConnectorHttpRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorHttpRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorHttpRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorHttpRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorHttpRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ConnectorHttpRequest) GetUrl() ConnectorUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConnectorHttpRequest) GetUrlOk() (*ConnectorUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConnectorHttpRequest) SetUrl(v ConnectorUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConnectorHttpRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


