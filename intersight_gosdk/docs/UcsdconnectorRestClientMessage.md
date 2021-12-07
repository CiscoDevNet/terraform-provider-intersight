# UcsdconnectorRestClientMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ucsdconnector.RestClientMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ucsdconnector.RestClientMessage"]
**Body** | Pointer to **string** | Payload which is sent along with the request. Most applicable to POST methods. | [optional] 
**Header** | Pointer to **interface{}** | Headers to be passed with the HTTP rest request. | [optional] 
**Method** | Pointer to **string** | REST Method, should be set to one of [HTTP.MethodGet, HTTP.MethodPost]. | [optional] 
**RestUrl** | Pointer to **string** | REST URL endpoint to which the HTTP request is sent. | [optional] 

## Methods

### NewUcsdconnectorRestClientMessage

`func NewUcsdconnectorRestClientMessage(classId string, objectType string, ) *UcsdconnectorRestClientMessage`

NewUcsdconnectorRestClientMessage instantiates a new UcsdconnectorRestClientMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcsdconnectorRestClientMessageWithDefaults

`func NewUcsdconnectorRestClientMessageWithDefaults() *UcsdconnectorRestClientMessage`

NewUcsdconnectorRestClientMessageWithDefaults instantiates a new UcsdconnectorRestClientMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UcsdconnectorRestClientMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UcsdconnectorRestClientMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UcsdconnectorRestClientMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UcsdconnectorRestClientMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UcsdconnectorRestClientMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UcsdconnectorRestClientMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBody

`func (o *UcsdconnectorRestClientMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UcsdconnectorRestClientMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UcsdconnectorRestClientMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UcsdconnectorRestClientMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeader

`func (o *UcsdconnectorRestClientMessage) GetHeader() interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *UcsdconnectorRestClientMessage) GetHeaderOk() (*interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *UcsdconnectorRestClientMessage) SetHeader(v interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *UcsdconnectorRestClientMessage) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *UcsdconnectorRestClientMessage) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *UcsdconnectorRestClientMessage) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetMethod

`func (o *UcsdconnectorRestClientMessage) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UcsdconnectorRestClientMessage) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UcsdconnectorRestClientMessage) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *UcsdconnectorRestClientMessage) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRestUrl

`func (o *UcsdconnectorRestClientMessage) GetRestUrl() string`

GetRestUrl returns the RestUrl field if non-nil, zero value otherwise.

### GetRestUrlOk

`func (o *UcsdconnectorRestClientMessage) GetRestUrlOk() (*string, bool)`

GetRestUrlOk returns a tuple with the RestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestUrl

`func (o *UcsdconnectorRestClientMessage) SetRestUrl(v string)`

SetRestUrl sets RestUrl field to given value.

### HasRestUrl

`func (o *UcsdconnectorRestClientMessage) HasRestUrl() bool`

HasRestUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


