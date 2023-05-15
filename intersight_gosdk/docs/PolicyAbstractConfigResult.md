# PolicyAbstractConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AppliedPolicies** | Pointer to [**[]PolicyPolicyStatus**](PolicyPolicyStatus.md) |  | [optional] 
**ConfigStage** | Pointer to **string** | The current running stage of the configuration or workflow. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. | [optional] [readonly] 
**ValidationState** | Pointer to **string** | Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. | [optional] [readonly] 

## Methods

### NewPolicyAbstractConfigResult

`func NewPolicyAbstractConfigResult(classId string, objectType string, ) *PolicyAbstractConfigResult`

NewPolicyAbstractConfigResult instantiates a new PolicyAbstractConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigResultWithDefaults

`func NewPolicyAbstractConfigResultWithDefaults() *PolicyAbstractConfigResult`

NewPolicyAbstractConfigResultWithDefaults instantiates a new PolicyAbstractConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractConfigResult) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractConfigResult) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractConfigResult) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractConfigResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractConfigResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractConfigResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppliedPolicies

`func (o *PolicyAbstractConfigResult) GetAppliedPolicies() []PolicyPolicyStatus`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *PolicyAbstractConfigResult) GetAppliedPoliciesOk() (*[]PolicyPolicyStatus, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *PolicyAbstractConfigResult) SetAppliedPolicies(v []PolicyPolicyStatus)`

SetAppliedPolicies sets AppliedPolicies field to given value.

### HasAppliedPolicies

`func (o *PolicyAbstractConfigResult) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### SetAppliedPoliciesNil

`func (o *PolicyAbstractConfigResult) SetAppliedPoliciesNil(b bool)`

 SetAppliedPoliciesNil sets the value for AppliedPolicies to be an explicit nil

### UnsetAppliedPolicies
`func (o *PolicyAbstractConfigResult) UnsetAppliedPolicies()`

UnsetAppliedPolicies ensures that no value is present for AppliedPolicies, not even an explicit nil
### GetConfigStage

`func (o *PolicyAbstractConfigResult) GetConfigStage() string`

GetConfigStage returns the ConfigStage field if non-nil, zero value otherwise.

### GetConfigStageOk

`func (o *PolicyAbstractConfigResult) GetConfigStageOk() (*string, bool)`

GetConfigStageOk returns a tuple with the ConfigStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStage

`func (o *PolicyAbstractConfigResult) SetConfigStage(v string)`

SetConfigStage sets ConfigStage field to given value.

### HasConfigStage

`func (o *PolicyAbstractConfigResult) HasConfigStage() bool`

HasConfigStage returns a boolean if a field has been set.

### GetConfigState

`func (o *PolicyAbstractConfigResult) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyAbstractConfigResult) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyAbstractConfigResult) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyAbstractConfigResult) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetValidationState

`func (o *PolicyAbstractConfigResult) GetValidationState() string`

GetValidationState returns the ValidationState field if non-nil, zero value otherwise.

### GetValidationStateOk

`func (o *PolicyAbstractConfigResult) GetValidationStateOk() (*string, bool)`

GetValidationStateOk returns a tuple with the ValidationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationState

`func (o *PolicyAbstractConfigResult) SetValidationState(v string)`

SetValidationState sets ValidationState field to given value.

### HasValidationState

`func (o *PolicyAbstractConfigResult) HasValidationState() bool`

HasValidationState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


