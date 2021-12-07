# ImcconnectorWebUiMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "imcconnector.WebUiMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "imcconnector.WebUiMessage"]
**WebUiRequest** | Pointer to **string** | The body content of the UI HTTP request to send to the BMC platform. | [optional] 

## Methods

### NewImcconnectorWebUiMessageAllOf

`func NewImcconnectorWebUiMessageAllOf(classId string, objectType string, ) *ImcconnectorWebUiMessageAllOf`

NewImcconnectorWebUiMessageAllOf instantiates a new ImcconnectorWebUiMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImcconnectorWebUiMessageAllOfWithDefaults

`func NewImcconnectorWebUiMessageAllOfWithDefaults() *ImcconnectorWebUiMessageAllOf`

NewImcconnectorWebUiMessageAllOfWithDefaults instantiates a new ImcconnectorWebUiMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ImcconnectorWebUiMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ImcconnectorWebUiMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ImcconnectorWebUiMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ImcconnectorWebUiMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ImcconnectorWebUiMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ImcconnectorWebUiMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWebUiRequest

`func (o *ImcconnectorWebUiMessageAllOf) GetWebUiRequest() string`

GetWebUiRequest returns the WebUiRequest field if non-nil, zero value otherwise.

### GetWebUiRequestOk

`func (o *ImcconnectorWebUiMessageAllOf) GetWebUiRequestOk() (*string, bool)`

GetWebUiRequestOk returns a tuple with the WebUiRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUiRequest

`func (o *ImcconnectorWebUiMessageAllOf) SetWebUiRequest(v string)`

SetWebUiRequest sets WebUiRequest field to given value.

### HasWebUiRequest

`func (o *ImcconnectorWebUiMessageAllOf) HasWebUiRequest() bool`

HasWebUiRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


