# InlineResponseDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | A well-defined error code. | [optional] 
**ErrorMessage** | Pointer to **string** | A free-form message with more information about the error. May be null. | [optional] 
**ErrorClass** | Pointer to **string** | The class of the exception that caused this error. May be null. | [optional] 
**Host** | Pointer to **string** | The host on which this error occurred. May be null. | [optional] 

## Methods

### NewInlineResponseDefault

`func NewInlineResponseDefault() *InlineResponseDefault`

NewInlineResponseDefault instantiates a new InlineResponseDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponseDefaultWithDefaults

`func NewInlineResponseDefaultWithDefaults() *InlineResponseDefault`

NewInlineResponseDefaultWithDefaults instantiates a new InlineResponseDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InlineResponseDefault) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponseDefault) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponseDefault) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponseDefault) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorMessage

`func (o *InlineResponseDefault) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InlineResponseDefault) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InlineResponseDefault) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InlineResponseDefault) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorClass

`func (o *InlineResponseDefault) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *InlineResponseDefault) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *InlineResponseDefault) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.

### HasErrorClass

`func (o *InlineResponseDefault) HasErrorClass() bool`

HasErrorClass returns a boolean if a field has been set.

### GetHost

`func (o *InlineResponseDefault) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InlineResponseDefault) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InlineResponseDefault) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InlineResponseDefault) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


