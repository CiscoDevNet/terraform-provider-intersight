# NotificationHttpHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "notification.HttpHeader"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "notification.HttpHeader"]
**ClearTextValue** | Pointer to **string** | The plaintext value of the HTTP header. Use this field when the header value is not sensitive. Used only when encrypt is false. | [optional] 
**Encrypt** | Pointer to **bool** | When false (default), clearTextValue is attached to the outgoing request as-is. When true, encryptedValue is decrypted and attached to the outgoing request. | [optional] 
**EncryptedValue** | Pointer to **string** | The sensitive value of the HTTP header, stored encrypted. Use this field when the header value is sensitive, for example an API key or token. Used only when encrypt is true. Not returned in GET API responses; use IsEncryptedValueSet to check whether a value has been set. | [optional] 
**IsEncryptedValueSet** | Pointer to **bool** | Indicates whether the value of the &#39;encryptedValue&#39; property has been set. | [optional] [readonly] [default to false]
**Name** | Pointer to **string** | Name of the HTTP header, for example X-Custom-Header. | [optional] 

## Methods

### NewNotificationHttpHeader

`func NewNotificationHttpHeader(classId string, objectType string, ) *NotificationHttpHeader`

NewNotificationHttpHeader instantiates a new NotificationHttpHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationHttpHeaderWithDefaults

`func NewNotificationHttpHeaderWithDefaults() *NotificationHttpHeader`

NewNotificationHttpHeaderWithDefaults instantiates a new NotificationHttpHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NotificationHttpHeader) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NotificationHttpHeader) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NotificationHttpHeader) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NotificationHttpHeader) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationHttpHeader) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationHttpHeader) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClearTextValue

`func (o *NotificationHttpHeader) GetClearTextValue() string`

GetClearTextValue returns the ClearTextValue field if non-nil, zero value otherwise.

### GetClearTextValueOk

`func (o *NotificationHttpHeader) GetClearTextValueOk() (*string, bool)`

GetClearTextValueOk returns a tuple with the ClearTextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearTextValue

`func (o *NotificationHttpHeader) SetClearTextValue(v string)`

SetClearTextValue sets ClearTextValue field to given value.

### HasClearTextValue

`func (o *NotificationHttpHeader) HasClearTextValue() bool`

HasClearTextValue returns a boolean if a field has been set.

### GetEncrypt

`func (o *NotificationHttpHeader) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *NotificationHttpHeader) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *NotificationHttpHeader) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *NotificationHttpHeader) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptedValue

`func (o *NotificationHttpHeader) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *NotificationHttpHeader) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *NotificationHttpHeader) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *NotificationHttpHeader) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetIsEncryptedValueSet

`func (o *NotificationHttpHeader) GetIsEncryptedValueSet() bool`

GetIsEncryptedValueSet returns the IsEncryptedValueSet field if non-nil, zero value otherwise.

### GetIsEncryptedValueSetOk

`func (o *NotificationHttpHeader) GetIsEncryptedValueSetOk() (*bool, bool)`

GetIsEncryptedValueSetOk returns a tuple with the IsEncryptedValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptedValueSet

`func (o *NotificationHttpHeader) SetIsEncryptedValueSet(v bool)`

SetIsEncryptedValueSet sets IsEncryptedValueSet field to given value.

### HasIsEncryptedValueSet

`func (o *NotificationHttpHeader) HasIsEncryptedValueSet() bool`

HasIsEncryptedValueSet returns a boolean if a field has been set.

### GetName

`func (o *NotificationHttpHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationHttpHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationHttpHeader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationHttpHeader) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


