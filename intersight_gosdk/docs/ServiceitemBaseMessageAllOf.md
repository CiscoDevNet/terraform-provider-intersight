# ServiceitemBaseMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ActionOperation** | Pointer to **string** | The type of action instance operation, such as Validate, Start, Retry, RetryFailed, Cancel, Stop etc. * &#x60;None&#x60; - No action is set, this is the default value for action field. * &#x60;Validate&#x60; - Validate the action instance inputs and run the validation workflows. * &#x60;Start&#x60; - Start a new execution of the action instance. * &#x60;Rerun&#x60; - Rerun the service item action instance from the beginning. * &#x60;Retry&#x60; - Retry the workflow that has failed from the failure point. * &#x60;Cancel&#x60; - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * &#x60;Stop&#x60; - Stop the action instance which is in progress and didn&#39;t complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped. | [optional] [readonly] [default to "None"]
**CreateTime** | Pointer to **time.Time** | The timestamp when the message was created. | [optional] [readonly] 
**Message** | Pointer to **string** | An i18n message which can be localized and conveys status on the action. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the message such as error, warning, info etc. * &#x60;Info&#x60; - The enum represents the log level to be used to convey info message. * &#x60;Warning&#x60; - The enum represents the log level to be used to convey warning message. * &#x60;Debug&#x60; - The enum represents the log level to be used to convey debug message. * &#x60;Error&#x60; - The enum represents the log level to be used to convey error message. | [optional] [readonly] [default to "Info"]

## Methods

### NewServiceitemBaseMessageAllOf

`func NewServiceitemBaseMessageAllOf(classId string, objectType string, ) *ServiceitemBaseMessageAllOf`

NewServiceitemBaseMessageAllOf instantiates a new ServiceitemBaseMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceitemBaseMessageAllOfWithDefaults

`func NewServiceitemBaseMessageAllOfWithDefaults() *ServiceitemBaseMessageAllOf`

NewServiceitemBaseMessageAllOfWithDefaults instantiates a new ServiceitemBaseMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ServiceitemBaseMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServiceitemBaseMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServiceitemBaseMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ServiceitemBaseMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServiceitemBaseMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServiceitemBaseMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActionOperation

`func (o *ServiceitemBaseMessageAllOf) GetActionOperation() string`

GetActionOperation returns the ActionOperation field if non-nil, zero value otherwise.

### GetActionOperationOk

`func (o *ServiceitemBaseMessageAllOf) GetActionOperationOk() (*string, bool)`

GetActionOperationOk returns a tuple with the ActionOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOperation

`func (o *ServiceitemBaseMessageAllOf) SetActionOperation(v string)`

SetActionOperation sets ActionOperation field to given value.

### HasActionOperation

`func (o *ServiceitemBaseMessageAllOf) HasActionOperation() bool`

HasActionOperation returns a boolean if a field has been set.

### GetCreateTime

`func (o *ServiceitemBaseMessageAllOf) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ServiceitemBaseMessageAllOf) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ServiceitemBaseMessageAllOf) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ServiceitemBaseMessageAllOf) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceitemBaseMessageAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceitemBaseMessageAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceitemBaseMessageAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceitemBaseMessageAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *ServiceitemBaseMessageAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServiceitemBaseMessageAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServiceitemBaseMessageAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServiceitemBaseMessageAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


