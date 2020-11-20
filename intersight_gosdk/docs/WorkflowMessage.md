# WorkflowMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.Message"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.Message"]
**Message** | Pointer to **string** | An i18n message that can be translated in multiple languages to support internationalization. | [optional] 
**Severity** | Pointer to **string** | The severity of the Task or Workflow message warning/error/info etc. * &#x60;Info&#x60; - The enum represents the log level to be used to convey info message. * &#x60;Warning&#x60; - The enum represents the log level to be used to convey warning message. * &#x60;Debug&#x60; - The enum represents the log level to be used to convey debug message. * &#x60;Error&#x60; - The enum represents the log level to be used to convey error message. | [optional] [default to "Info"]

## Methods

### NewWorkflowMessage

`func NewWorkflowMessage(classId string, objectType string, ) *WorkflowMessage`

NewWorkflowMessage instantiates a new WorkflowMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMessageWithDefaults

`func NewWorkflowMessageWithDefaults() *WorkflowMessage`

NewWorkflowMessageWithDefaults instantiates a new WorkflowMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *WorkflowMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *WorkflowMessage) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WorkflowMessage) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WorkflowMessage) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WorkflowMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


