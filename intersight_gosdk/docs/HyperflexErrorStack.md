# HyperflexErrorStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ErrorStack"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ErrorStack"]
**Message** | Pointer to **string** | The error message string for this error stack. | [optional] [readonly] 
**MessageId** | Pointer to **int64** | The error message ID for this error stack. | [optional] [readonly] 

## Methods

### NewHyperflexErrorStack

`func NewHyperflexErrorStack(classId string, objectType string, ) *HyperflexErrorStack`

NewHyperflexErrorStack instantiates a new HyperflexErrorStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexErrorStackWithDefaults

`func NewHyperflexErrorStackWithDefaults() *HyperflexErrorStack`

NewHyperflexErrorStackWithDefaults instantiates a new HyperflexErrorStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexErrorStack) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexErrorStack) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexErrorStack) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexErrorStack) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexErrorStack) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexErrorStack) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *HyperflexErrorStack) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HyperflexErrorStack) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HyperflexErrorStack) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HyperflexErrorStack) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageId

`func (o *HyperflexErrorStack) GetMessageId() int64`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *HyperflexErrorStack) GetMessageIdOk() (*int64, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *HyperflexErrorStack) SetMessageId(v int64)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *HyperflexErrorStack) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


