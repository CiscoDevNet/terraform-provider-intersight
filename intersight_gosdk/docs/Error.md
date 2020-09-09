# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A value that is used to determine the nature of the error, and why it occurred. | 
**MessageId** | Pointer to **string** | A language-independent identifier of the specific error message. | [optional] 
**Message** | **string** | A human-readable description of the error, localized with the i18n standard. The language is determined from the \&quot;Accept-Language\&quot; request-header. The Accept-Language request-header restricts the set of natural languages that are preferred as a response to the request. See RFC 2616 for more details. Codes: 1. **InternalServerError**   An internal error occurred. 1. **InvalidMethod**         The HTTP method (POST, PUT...) is invalid for this API path. For example, a POST request was sent but this API path only supports GET. 1. **InvalidUrl**            The HTTP request contains an invalid URL. 1. **InvalidRequest**        The HTTP request contains an invalid or malformed message body. 1. **NotSupported**          The request is not supported for the specified REST resource. 1. **NotImplemented**        This API path is experimental and not implemented yet. 1. **NotFound**              The requested REST resource does not exist. 1. **AuthenticationFailure** The request lacks valid authentication credentials. 1. **UnauthorizedOperation** The client is not authorized to perform the operation, such as when the user has insufficient privileges. 1. **ValidationConflict**    The request contains conflicting attributes, such as two mutually exclusive attribute values. 1. **ServiceUnavailable**    See RFC 7231, status 503. | 
**Cause** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewError

`func NewError(code string, message string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessageId

`func (o *Error) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Error) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Error) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Error) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCause

`func (o *Error) GetCause() Error`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *Error) GetCauseOk() (*Error, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *Error) SetCause(v Error)`

SetCause sets Cause field to given value.

### HasCause

`func (o *Error) HasCause() bool`

HasCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


