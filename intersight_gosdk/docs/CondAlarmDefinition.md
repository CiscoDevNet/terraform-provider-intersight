# CondAlarmDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmDefinition"]
**Actions** | Pointer to [**[]CondAlarmAction**](CondAlarmAction.md) |  | [optional] 

## Methods

### NewCondAlarmDefinition

`func NewCondAlarmDefinition(classId string, objectType string, ) *CondAlarmDefinition`

NewCondAlarmDefinition instantiates a new CondAlarmDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmDefinitionWithDefaults

`func NewCondAlarmDefinitionWithDefaults() *CondAlarmDefinition`

NewCondAlarmDefinitionWithDefaults instantiates a new CondAlarmDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActions

`func (o *CondAlarmDefinition) GetActions() []CondAlarmAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CondAlarmDefinition) GetActionsOk() (*[]CondAlarmAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CondAlarmDefinition) SetActions(v []CondAlarmAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CondAlarmDefinition) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *CondAlarmDefinition) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *CondAlarmDefinition) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


