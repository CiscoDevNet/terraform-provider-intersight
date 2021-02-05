# WorkflowWebApiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WebApi"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WebApi"]
**Cookies** | Pointer to **interface{}** | Collection of key value pairs to set in the request header as Cookie list. | [optional] 
**EndpointRequestType** | Pointer to **string** | If the target type is Endpoint, this property determines whether the request is to be handled as internal request or external request by the device connector. * &#x60;Internal&#x60; - The endpoint API executed is an internal request handled by the device connector plugin. * &#x60;External&#x60; - The endpoint API request is passed through by the device connector. | [optional] [default to "Internal"]
**Headers** | Pointer to **interface{}** | Collection of key value pairs to set in the request header. | [optional] 
**Method** | Pointer to **string** | The HTTP method to be executed in the given URL (GET, POST, PUT, etc). If the value is not specified, GET will be used as default. The supported values are GET, POST, PUT, DELETE, PATCH, HEAD. | [optional] 
**MoType** | Pointer to **string** | The type of the intersight object for which API request is to be made. The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value. | [optional] 
**Protocol** | Pointer to **string** | The accepted web protocol values are http and https. | [optional] 
**TargetType** | Pointer to **string** | If the web API is to be executed in a remote device connected to the Intersight through device connector, &#39;Endpoint&#39; is expected as the value whereas if the API is an Intersight API, &#39;Local&#39; is expected as the value. | [optional] 
**Url** | Pointer to **string** | The URL of the resource in the target to which the API request is made. | [optional] 

## Methods

### NewWorkflowWebApiAllOf

`func NewWorkflowWebApiAllOf(classId string, objectType string, ) *WorkflowWebApiAllOf`

NewWorkflowWebApiAllOf instantiates a new WorkflowWebApiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWebApiAllOfWithDefaults

`func NewWorkflowWebApiAllOfWithDefaults() *WorkflowWebApiAllOf`

NewWorkflowWebApiAllOfWithDefaults instantiates a new WorkflowWebApiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWebApiAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWebApiAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWebApiAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWebApiAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWebApiAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWebApiAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCookies

`func (o *WorkflowWebApiAllOf) GetCookies() interface{}`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *WorkflowWebApiAllOf) GetCookiesOk() (*interface{}, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *WorkflowWebApiAllOf) SetCookies(v interface{})`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *WorkflowWebApiAllOf) HasCookies() bool`

HasCookies returns a boolean if a field has been set.

### SetCookiesNil

`func (o *WorkflowWebApiAllOf) SetCookiesNil(b bool)`

 SetCookiesNil sets the value for Cookies to be an explicit nil

### UnsetCookies
`func (o *WorkflowWebApiAllOf) UnsetCookies()`

UnsetCookies ensures that no value is present for Cookies, not even an explicit nil
### GetEndpointRequestType

`func (o *WorkflowWebApiAllOf) GetEndpointRequestType() string`

GetEndpointRequestType returns the EndpointRequestType field if non-nil, zero value otherwise.

### GetEndpointRequestTypeOk

`func (o *WorkflowWebApiAllOf) GetEndpointRequestTypeOk() (*string, bool)`

GetEndpointRequestTypeOk returns a tuple with the EndpointRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointRequestType

`func (o *WorkflowWebApiAllOf) SetEndpointRequestType(v string)`

SetEndpointRequestType sets EndpointRequestType field to given value.

### HasEndpointRequestType

`func (o *WorkflowWebApiAllOf) HasEndpointRequestType() bool`

HasEndpointRequestType returns a boolean if a field has been set.

### GetHeaders

`func (o *WorkflowWebApiAllOf) GetHeaders() interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WorkflowWebApiAllOf) GetHeadersOk() (*interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WorkflowWebApiAllOf) SetHeaders(v interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WorkflowWebApiAllOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *WorkflowWebApiAllOf) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *WorkflowWebApiAllOf) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetMethod

`func (o *WorkflowWebApiAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WorkflowWebApiAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WorkflowWebApiAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WorkflowWebApiAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMoType

`func (o *WorkflowWebApiAllOf) GetMoType() string`

GetMoType returns the MoType field if non-nil, zero value otherwise.

### GetMoTypeOk

`func (o *WorkflowWebApiAllOf) GetMoTypeOk() (*string, bool)`

GetMoTypeOk returns a tuple with the MoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoType

`func (o *WorkflowWebApiAllOf) SetMoType(v string)`

SetMoType sets MoType field to given value.

### HasMoType

`func (o *WorkflowWebApiAllOf) HasMoType() bool`

HasMoType returns a boolean if a field has been set.

### GetProtocol

`func (o *WorkflowWebApiAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WorkflowWebApiAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WorkflowWebApiAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WorkflowWebApiAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTargetType

`func (o *WorkflowWebApiAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *WorkflowWebApiAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *WorkflowWebApiAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *WorkflowWebApiAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetUrl

`func (o *WorkflowWebApiAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkflowWebApiAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkflowWebApiAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WorkflowWebApiAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


