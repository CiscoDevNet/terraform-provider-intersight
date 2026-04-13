# CondAlarmRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmRule"]
**Description** | Pointer to **string** | Informative description of AlarmRule. | [optional] 
**Name** | Pointer to **string** | A unique name assigned by the user to AlarmRule. This user-defined name acts as the identity field, ensuring that AlarmRule is distinctly identifiable within the account. | [optional] 
**State** | Pointer to **string** | Controls the behavior of alarm processing depending upon the state. If Enabled, which is also the default behavior, the alarm is evaluated for the device based on the conditions specified in the ThresholdDefinition objects attached to it. If Disabled or SystemDisabled, alarm is not evaluated for the device and the existing alarms raised against the device is cleared. * &#x60;Enabled&#x60; - User initiated action which is also the default action that enables alarm evaluation for any condition that meets the criteria. * &#x60;Disabled&#x60; - User initiated action that disables alarm evaluation temporarily and clears the existing alarms. * &#x60;SystemDisabled&#x60; - System initiated action that disables alarm evaluation temporarily and clears the existing alarms once alarm limit per alarm rule is reached. | [optional] [default to "Enabled"]
**SystemDisabledTime** | Pointer to **time.Time** | The time at which AlarmRule object is system-disabled. | [optional] [readonly] 
**ThresholdDefinitionStates** | Pointer to [**[]CondThresholdDefinitionState**](CondThresholdDefinitionState.md) |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Devices** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 
**ThresholdDefinitions** | Pointer to [**[]CondThresholdDefinitionRelationship**](CondThresholdDefinitionRelationship.md) | An array of relationships to condThresholdDefinition resources. | [optional] 

## Methods

### NewCondAlarmRule

`func NewCondAlarmRule(classId string, objectType string, ) *CondAlarmRule`

NewCondAlarmRule instantiates a new CondAlarmRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmRuleWithDefaults

`func NewCondAlarmRuleWithDefaults() *CondAlarmRule`

NewCondAlarmRuleWithDefaults instantiates a new CondAlarmRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *CondAlarmRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarmRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarmRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarmRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CondAlarmRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarmRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarmRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarmRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *CondAlarmRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CondAlarmRule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CondAlarmRule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CondAlarmRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemDisabledTime

`func (o *CondAlarmRule) GetSystemDisabledTime() time.Time`

GetSystemDisabledTime returns the SystemDisabledTime field if non-nil, zero value otherwise.

### GetSystemDisabledTimeOk

`func (o *CondAlarmRule) GetSystemDisabledTimeOk() (*time.Time, bool)`

GetSystemDisabledTimeOk returns a tuple with the SystemDisabledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDisabledTime

`func (o *CondAlarmRule) SetSystemDisabledTime(v time.Time)`

SetSystemDisabledTime sets SystemDisabledTime field to given value.

### HasSystemDisabledTime

`func (o *CondAlarmRule) HasSystemDisabledTime() bool`

HasSystemDisabledTime returns a boolean if a field has been set.

### GetThresholdDefinitionStates

`func (o *CondAlarmRule) GetThresholdDefinitionStates() []CondThresholdDefinitionState`

GetThresholdDefinitionStates returns the ThresholdDefinitionStates field if non-nil, zero value otherwise.

### GetThresholdDefinitionStatesOk

`func (o *CondAlarmRule) GetThresholdDefinitionStatesOk() (*[]CondThresholdDefinitionState, bool)`

GetThresholdDefinitionStatesOk returns a tuple with the ThresholdDefinitionStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDefinitionStates

`func (o *CondAlarmRule) SetThresholdDefinitionStates(v []CondThresholdDefinitionState)`

SetThresholdDefinitionStates sets ThresholdDefinitionStates field to given value.

### HasThresholdDefinitionStates

`func (o *CondAlarmRule) HasThresholdDefinitionStates() bool`

HasThresholdDefinitionStates returns a boolean if a field has been set.

### SetThresholdDefinitionStatesNil

`func (o *CondAlarmRule) SetThresholdDefinitionStatesNil(b bool)`

 SetThresholdDefinitionStatesNil sets the value for ThresholdDefinitionStates to be an explicit nil

### UnsetThresholdDefinitionStates
`func (o *CondAlarmRule) UnsetThresholdDefinitionStates()`

UnsetThresholdDefinitionStates ensures that no value is present for ThresholdDefinitionStates, not even an explicit nil
### GetAccount

`func (o *CondAlarmRule) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CondAlarmRule) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CondAlarmRule) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CondAlarmRule) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CondAlarmRule) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CondAlarmRule) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetDevices

`func (o *CondAlarmRule) GetDevices() []MoBaseMoRelationship`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CondAlarmRule) GetDevicesOk() (*[]MoBaseMoRelationship, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CondAlarmRule) SetDevices(v []MoBaseMoRelationship)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CondAlarmRule) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### SetDevicesNil

`func (o *CondAlarmRule) SetDevicesNil(b bool)`

 SetDevicesNil sets the value for Devices to be an explicit nil

### UnsetDevices
`func (o *CondAlarmRule) UnsetDevices()`

UnsetDevices ensures that no value is present for Devices, not even an explicit nil
### GetThresholdDefinitions

`func (o *CondAlarmRule) GetThresholdDefinitions() []CondThresholdDefinitionRelationship`

GetThresholdDefinitions returns the ThresholdDefinitions field if non-nil, zero value otherwise.

### GetThresholdDefinitionsOk

`func (o *CondAlarmRule) GetThresholdDefinitionsOk() (*[]CondThresholdDefinitionRelationship, bool)`

GetThresholdDefinitionsOk returns a tuple with the ThresholdDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdDefinitions

`func (o *CondAlarmRule) SetThresholdDefinitions(v []CondThresholdDefinitionRelationship)`

SetThresholdDefinitions sets ThresholdDefinitions field to given value.

### HasThresholdDefinitions

`func (o *CondAlarmRule) HasThresholdDefinitions() bool`

HasThresholdDefinitions returns a boolean if a field has been set.

### SetThresholdDefinitionsNil

`func (o *CondAlarmRule) SetThresholdDefinitionsNil(b bool)`

 SetThresholdDefinitionsNil sets the value for ThresholdDefinitions to be an explicit nil

### UnsetThresholdDefinitions
`func (o *CondAlarmRule) UnsetThresholdDefinitions()`

UnsetThresholdDefinitions ensures that no value is present for ThresholdDefinitions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


