# PolicyScheduledAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ScheduledAction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ScheduledAction"]
**Action** | Pointer to **string** | Name of the action to be performed on the profile. | [optional] 
**ActionQualifier** | Pointer to [**NullablePolicyActionQualifier**](PolicyActionQualifier.md) |  | [optional] 
**ProceedOnReboot** | Pointer to **bool** | ProceedOnReboot can be used to acknowledge server reboot while triggering deploy/activate. | [optional] 

## Methods

### NewPolicyScheduledAction

`func NewPolicyScheduledAction(classId string, objectType string, ) *PolicyScheduledAction`

NewPolicyScheduledAction instantiates a new PolicyScheduledAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyScheduledActionWithDefaults

`func NewPolicyScheduledActionWithDefaults() *PolicyScheduledAction`

NewPolicyScheduledActionWithDefaults instantiates a new PolicyScheduledAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyScheduledAction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyScheduledAction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyScheduledAction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyScheduledAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyScheduledAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyScheduledAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *PolicyScheduledAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyScheduledAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyScheduledAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyScheduledAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionQualifier

`func (o *PolicyScheduledAction) GetActionQualifier() PolicyActionQualifier`

GetActionQualifier returns the ActionQualifier field if non-nil, zero value otherwise.

### GetActionQualifierOk

`func (o *PolicyScheduledAction) GetActionQualifierOk() (*PolicyActionQualifier, bool)`

GetActionQualifierOk returns a tuple with the ActionQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionQualifier

`func (o *PolicyScheduledAction) SetActionQualifier(v PolicyActionQualifier)`

SetActionQualifier sets ActionQualifier field to given value.

### HasActionQualifier

`func (o *PolicyScheduledAction) HasActionQualifier() bool`

HasActionQualifier returns a boolean if a field has been set.

### SetActionQualifierNil

`func (o *PolicyScheduledAction) SetActionQualifierNil(b bool)`

 SetActionQualifierNil sets the value for ActionQualifier to be an explicit nil

### UnsetActionQualifier
`func (o *PolicyScheduledAction) UnsetActionQualifier()`

UnsetActionQualifier ensures that no value is present for ActionQualifier, not even an explicit nil
### GetProceedOnReboot

`func (o *PolicyScheduledAction) GetProceedOnReboot() bool`

GetProceedOnReboot returns the ProceedOnReboot field if non-nil, zero value otherwise.

### GetProceedOnRebootOk

`func (o *PolicyScheduledAction) GetProceedOnRebootOk() (*bool, bool)`

GetProceedOnRebootOk returns a tuple with the ProceedOnReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceedOnReboot

`func (o *PolicyScheduledAction) SetProceedOnReboot(v bool)`

SetProceedOnReboot sets ProceedOnReboot field to given value.

### HasProceedOnReboot

`func (o *PolicyScheduledAction) HasProceedOnReboot() bool`

HasProceedOnReboot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


