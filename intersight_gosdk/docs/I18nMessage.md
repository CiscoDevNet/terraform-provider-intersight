# I18nMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The default (en-US) localized message. Default localized message will be stored and directly retrieved when the user&#39;s locale setting is en-US. | [optional] [readonly] 
**MessageId** | Pointer to **string** | The unique message identitifer used to lookup text templates in a multi-language message catalog. | [optional] [readonly] 
**MessageParams** | Pointer to [**[]I18nMessageParam**](i18n.MessageParam.md) |  | [optional] 

## Methods

### NewI18nMessage

`func NewI18nMessage() *I18nMessage`

NewI18nMessage instantiates a new I18nMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewI18nMessageWithDefaults

`func NewI18nMessageWithDefaults() *I18nMessage`

NewI18nMessageWithDefaults instantiates a new I18nMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *I18nMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *I18nMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *I18nMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *I18nMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageId

`func (o *I18nMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *I18nMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *I18nMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *I18nMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageParams

`func (o *I18nMessage) GetMessageParams() []I18nMessageParam`

GetMessageParams returns the MessageParams field if non-nil, zero value otherwise.

### GetMessageParamsOk

`func (o *I18nMessage) GetMessageParamsOk() (*[]I18nMessageParam, bool)`

GetMessageParamsOk returns a tuple with the MessageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageParams

`func (o *I18nMessage) SetMessageParams(v []I18nMessageParam)`

SetMessageParams sets MessageParams field to given value.

### HasMessageParams

`func (o *I18nMessage) HasMessageParams() bool`

HasMessageParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


