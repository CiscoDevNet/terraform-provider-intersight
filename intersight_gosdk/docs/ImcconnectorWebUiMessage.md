# ImcconnectorWebUiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "imcconnector.WebUiMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "imcconnector.WebUiMessage"]
**WebUiRequest** | Pointer to **string** | The body content of the UI HTTP request to send to the BMC platform. | [optional] 

## Methods

### NewImcconnectorWebUiMessage

`func NewImcconnectorWebUiMessage(classId string, objectType string, ) *ImcconnectorWebUiMessage`

NewImcconnectorWebUiMessage instantiates a new ImcconnectorWebUiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImcconnectorWebUiMessageWithDefaults

`func NewImcconnectorWebUiMessageWithDefaults() *ImcconnectorWebUiMessage`

NewImcconnectorWebUiMessageWithDefaults instantiates a new ImcconnectorWebUiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ImcconnectorWebUiMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ImcconnectorWebUiMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ImcconnectorWebUiMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ImcconnectorWebUiMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ImcconnectorWebUiMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ImcconnectorWebUiMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetWebUiRequest

`func (o *ImcconnectorWebUiMessage) GetWebUiRequest() string`

GetWebUiRequest returns the WebUiRequest field if non-nil, zero value otherwise.

### GetWebUiRequestOk

`func (o *ImcconnectorWebUiMessage) GetWebUiRequestOk() (*string, bool)`

GetWebUiRequestOk returns a tuple with the WebUiRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUiRequest

`func (o *ImcconnectorWebUiMessage) SetWebUiRequest(v string)`

SetWebUiRequest sets WebUiRequest field to given value.

### HasWebUiRequest

`func (o *ImcconnectorWebUiMessage) HasWebUiRequest() bool`

HasWebUiRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


