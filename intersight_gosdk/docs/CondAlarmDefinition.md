# CondAlarmDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmDefinition"]
**Actions** | Pointer to [**[]CondAlarmAction**](CondAlarmAction.md) |  | [optional] 
**Status** | Pointer to **string** | Controls the behavior of alarm processing depending upon the status. If Enabled, alarm is evaluated for any condition that meets the criteria. If Disabled or SystemDisabled, alarm is not evaluated and the existing alarms for this AlarmDefinition is cleared. If Inactive, alarm is not evaluated and the existing alarms for this AlarmDefinition is deleted. * &#x60;Enabled&#x60; - Enables alarm evaluation for any condition that meets the criteria. * &#x60;Disabled&#x60; - Disables alarm evaluation temporarily and clears the existing alarms. * &#x60;SystemDisabled&#x60; - Disables alarm evaluation temporarily and clears the existing alarms once alarm limit per alarm rule is reached. * &#x60;Inactive&#x60; - Stops alarm evaluation permanently and deletes the existing alarms. | [optional] [default to "Enabled"]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Device** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**ThresholdDefinition** | Pointer to [**NullableCondThresholdDefinitionRelationship**](CondThresholdDefinitionRelationship.md) |  | [optional] 

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
### GetStatus

`func (o *CondAlarmDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondAlarmDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondAlarmDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondAlarmDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *CondAlarmDefinition) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CondAlarmDefinition) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CondAlarmDefinition) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CondAlarmDefinition) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CondAlarmDefinition) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CondAlarmDefinition) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetDevice

`func (o *CondAlarmDefinition) GetDevice() MoBaseMoRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CondAlarmDefinition) GetDeviceOk() (*MoBaseMoRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CondAlarmDefinition) SetDevice(v MoBaseMoRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CondAlarmDefinition) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *CondAlarmDefinition) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *CondAlarmDefinition) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetThresholdDefinition

`func (o *CondAlarmDefinition) GetThresholdDefinition() CondThresholdDefinitionRelationship`

GetThresholdDefinition returns the ThresholdDefinition field if non-nil, zero value otherwise.

### GetThresholdDefinitionOk

`func (o *CondAlarmDefinition) GetThresholdDefinitionOk() (*CondThresholdDefinitionRelationship, bool)`

GetThresholdDefinitionOk returns a tuple with the ThresholdDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDefinition

`func (o *CondAlarmDefinition) SetThresholdDefinition(v CondThresholdDefinitionRelationship)`

SetThresholdDefinition sets ThresholdDefinition field to given value.

### HasThresholdDefinition

`func (o *CondAlarmDefinition) HasThresholdDefinition() bool`

HasThresholdDefinition returns a boolean if a field has been set.

### SetThresholdDefinitionNil

`func (o *CondAlarmDefinition) SetThresholdDefinitionNil(b bool)`

 SetThresholdDefinitionNil sets the value for ThresholdDefinition to be an explicit nil

### UnsetThresholdDefinition
`func (o *CondAlarmDefinition) UnsetThresholdDefinition()`

UnsetThresholdDefinition ensures that no value is present for ThresholdDefinition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


