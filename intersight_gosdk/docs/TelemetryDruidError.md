# TelemetryDruidError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | A well-defined error code. | [optional] 
**ErrorMessage** | Pointer to **string** | A free-form message with more information about the error. May be null. | [optional] 
**ErrorClass** | Pointer to **interface{}** | The class of the exception that caused this error. May be null. | [optional] 

## Methods

### NewTelemetryDruidError

`func NewTelemetryDruidError() *TelemetryDruidError`

NewTelemetryDruidError instantiates a new TelemetryDruidError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidErrorWithDefaults

`func NewTelemetryDruidErrorWithDefaults() *TelemetryDruidError`

NewTelemetryDruidErrorWithDefaults instantiates a new TelemetryDruidError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TelemetryDruidError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TelemetryDruidError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TelemetryDruidError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TelemetryDruidError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TelemetryDruidError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TelemetryDruidError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TelemetryDruidError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TelemetryDruidError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorClass

`func (o *TelemetryDruidError) GetErrorClass() interface{}`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *TelemetryDruidError) GetErrorClassOk() (*interface{}, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *TelemetryDruidError) SetErrorClass(v interface{})`

SetErrorClass sets ErrorClass field to given value.

### HasErrorClass

`func (o *TelemetryDruidError) HasErrorClass() bool`

HasErrorClass returns a boolean if a field has been set.

### SetErrorClassNil

`func (o *TelemetryDruidError) SetErrorClassNil(b bool)`

 SetErrorClassNil sets the value for ErrorClass to be an explicit nil

### UnsetErrorClass
`func (o *TelemetryDruidError) UnsetErrorClass()`

UnsetErrorClass ensures that no value is present for ErrorClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


