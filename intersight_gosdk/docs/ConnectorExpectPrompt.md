# ConnectorExpectPrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.ExpectPrompt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.ExpectPrompt"]
**Expect** | Pointer to **string** | The regex of the expect prompt of the interactive command. | [optional] 
**ExpectTimeout** | Pointer to **int64** | The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used. | [optional] 
**Send** | Pointer to **string** | The answer string to the expect prompt. | [optional] 

## Methods

### NewConnectorExpectPrompt

`func NewConnectorExpectPrompt(classId string, objectType string, ) *ConnectorExpectPrompt`

NewConnectorExpectPrompt instantiates a new ConnectorExpectPrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorExpectPromptWithDefaults

`func NewConnectorExpectPromptWithDefaults() *ConnectorExpectPrompt`

NewConnectorExpectPromptWithDefaults instantiates a new ConnectorExpectPrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorExpectPrompt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorExpectPrompt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorExpectPrompt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorExpectPrompt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorExpectPrompt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorExpectPrompt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpect

`func (o *ConnectorExpectPrompt) GetExpect() string`

GetExpect returns the Expect field if non-nil, zero value otherwise.

### GetExpectOk

`func (o *ConnectorExpectPrompt) GetExpectOk() (*string, bool)`

GetExpectOk returns a tuple with the Expect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpect

`func (o *ConnectorExpectPrompt) SetExpect(v string)`

SetExpect sets Expect field to given value.

### HasExpect

`func (o *ConnectorExpectPrompt) HasExpect() bool`

HasExpect returns a boolean if a field has been set.

### GetExpectTimeout

`func (o *ConnectorExpectPrompt) GetExpectTimeout() int64`

GetExpectTimeout returns the ExpectTimeout field if non-nil, zero value otherwise.

### GetExpectTimeoutOk

`func (o *ConnectorExpectPrompt) GetExpectTimeoutOk() (*int64, bool)`

GetExpectTimeoutOk returns a tuple with the ExpectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectTimeout

`func (o *ConnectorExpectPrompt) SetExpectTimeout(v int64)`

SetExpectTimeout sets ExpectTimeout field to given value.

### HasExpectTimeout

`func (o *ConnectorExpectPrompt) HasExpectTimeout() bool`

HasExpectTimeout returns a boolean if a field has been set.

### GetSend

`func (o *ConnectorExpectPrompt) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *ConnectorExpectPrompt) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *ConnectorExpectPrompt) SetSend(v string)`

SetSend sets Send field to given value.

### HasSend

`func (o *ConnectorExpectPrompt) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


