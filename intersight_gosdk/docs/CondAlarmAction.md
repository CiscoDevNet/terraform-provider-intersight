# CondAlarmAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmAction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmAction"]
**Message** | Pointer to [**NullableIssueMessage**](IssueMessage.md) |  | [optional] 
**Spec** | Pointer to [**NullableCondAlarmSpec**](CondAlarmSpec.md) |  | [optional] 

## Methods

### NewCondAlarmAction

`func NewCondAlarmAction(classId string, objectType string, ) *CondAlarmAction`

NewCondAlarmAction instantiates a new CondAlarmAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmActionWithDefaults

`func NewCondAlarmActionWithDefaults() *CondAlarmAction`

NewCondAlarmActionWithDefaults instantiates a new CondAlarmAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmAction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmAction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmAction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMessage

`func (o *CondAlarmAction) GetMessage() IssueMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CondAlarmAction) GetMessageOk() (*IssueMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CondAlarmAction) SetMessage(v IssueMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CondAlarmAction) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CondAlarmAction) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CondAlarmAction) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSpec

`func (o *CondAlarmAction) GetSpec() CondAlarmSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CondAlarmAction) GetSpecOk() (*CondAlarmSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CondAlarmAction) SetSpec(v CondAlarmSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CondAlarmAction) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *CondAlarmAction) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *CondAlarmAction) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


