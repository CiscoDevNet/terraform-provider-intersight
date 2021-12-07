# ConnectorHttpRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.HttpRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.HttpRequest"]
**AssetTargetMoid** | Pointer to **string** | The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight to create HTTP request along with any authentication info specifed. | [optional] 
**Body** | Pointer to **string** | Contents of the request body to send for PUT/PATCH/POST requests. | [optional] 
**DialTimeout** | Pointer to **int64** | The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used. | [optional] 
**EndpointMoid** | Pointer to **string** | The MO id of the asset.EndpointConnection this request is directed to. If set plugin will insert connection details into the request, including credentials if defined. | [optional] 
**Header** | Pointer to **interface{}** | Collection of key value pairs to set in the request header. | [optional] 
**Internal** | Pointer to **bool** | The request is for an internal platform API that requires authentication to be inserted by the platform implementation. | [optional] 
**Method** | Pointer to **string** | Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used. | [optional] 
**Url** | Pointer to [**NullableConnectorUrl**](ConnectorUrl.md) |  | [optional] 

## Methods

### NewConnectorHttpRequestAllOf

`func NewConnectorHttpRequestAllOf(classId string, objectType string, ) *ConnectorHttpRequestAllOf`

NewConnectorHttpRequestAllOf instantiates a new ConnectorHttpRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorHttpRequestAllOfWithDefaults

`func NewConnectorHttpRequestAllOfWithDefaults() *ConnectorHttpRequestAllOf`

NewConnectorHttpRequestAllOfWithDefaults instantiates a new ConnectorHttpRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorHttpRequestAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorHttpRequestAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorHttpRequestAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorHttpRequestAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorHttpRequestAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorHttpRequestAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssetTargetMoid

`func (o *ConnectorHttpRequestAllOf) GetAssetTargetMoid() string`

GetAssetTargetMoid returns the AssetTargetMoid field if non-nil, zero value otherwise.

### GetAssetTargetMoidOk

`func (o *ConnectorHttpRequestAllOf) GetAssetTargetMoidOk() (*string, bool)`

GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTargetMoid

`func (o *ConnectorHttpRequestAllOf) SetAssetTargetMoid(v string)`

SetAssetTargetMoid sets AssetTargetMoid field to given value.

### HasAssetTargetMoid

`func (o *ConnectorHttpRequestAllOf) HasAssetTargetMoid() bool`

HasAssetTargetMoid returns a boolean if a field has been set.

### GetBody

`func (o *ConnectorHttpRequestAllOf) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConnectorHttpRequestAllOf) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConnectorHttpRequestAllOf) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ConnectorHttpRequestAllOf) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDialTimeout

`func (o *ConnectorHttpRequestAllOf) GetDialTimeout() int64`

GetDialTimeout returns the DialTimeout field if non-nil, zero value otherwise.

### GetDialTimeoutOk

`func (o *ConnectorHttpRequestAllOf) GetDialTimeoutOk() (*int64, bool)`

GetDialTimeoutOk returns a tuple with the DialTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialTimeout

`func (o *ConnectorHttpRequestAllOf) SetDialTimeout(v int64)`

SetDialTimeout sets DialTimeout field to given value.

### HasDialTimeout

`func (o *ConnectorHttpRequestAllOf) HasDialTimeout() bool`

HasDialTimeout returns a boolean if a field has been set.

### GetEndpointMoid

`func (o *ConnectorHttpRequestAllOf) GetEndpointMoid() string`

GetEndpointMoid returns the EndpointMoid field if non-nil, zero value otherwise.

### GetEndpointMoidOk

`func (o *ConnectorHttpRequestAllOf) GetEndpointMoidOk() (*string, bool)`

GetEndpointMoidOk returns a tuple with the EndpointMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointMoid

`func (o *ConnectorHttpRequestAllOf) SetEndpointMoid(v string)`

SetEndpointMoid sets EndpointMoid field to given value.

### HasEndpointMoid

`func (o *ConnectorHttpRequestAllOf) HasEndpointMoid() bool`

HasEndpointMoid returns a boolean if a field has been set.

### GetHeader

`func (o *ConnectorHttpRequestAllOf) GetHeader() interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ConnectorHttpRequestAllOf) GetHeaderOk() (*interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ConnectorHttpRequestAllOf) SetHeader(v interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ConnectorHttpRequestAllOf) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *ConnectorHttpRequestAllOf) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *ConnectorHttpRequestAllOf) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetInternal

`func (o *ConnectorHttpRequestAllOf) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ConnectorHttpRequestAllOf) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ConnectorHttpRequestAllOf) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ConnectorHttpRequestAllOf) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetMethod

`func (o *ConnectorHttpRequestAllOf) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConnectorHttpRequestAllOf) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConnectorHttpRequestAllOf) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConnectorHttpRequestAllOf) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTimeout

`func (o *ConnectorHttpRequestAllOf) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectorHttpRequestAllOf) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectorHttpRequestAllOf) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectorHttpRequestAllOf) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ConnectorHttpRequestAllOf) GetUrl() ConnectorUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConnectorHttpRequestAllOf) GetUrlOk() (*ConnectorUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConnectorHttpRequestAllOf) SetUrl(v ConnectorUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConnectorHttpRequestAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConnectorHttpRequestAllOf) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConnectorHttpRequestAllOf) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


