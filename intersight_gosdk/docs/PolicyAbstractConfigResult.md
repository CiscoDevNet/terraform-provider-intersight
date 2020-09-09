# PolicyAbstractConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigStage** | Pointer to **string** | The current running stage of the configuration or workflow. | [optional] 
**ConfigState** | Pointer to **string** | Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**ValidationState** | Pointer to **string** | Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. | [optional] 

## Methods

### NewPolicyAbstractConfigResult

`func NewPolicyAbstractConfigResult() *PolicyAbstractConfigResult`

NewPolicyAbstractConfigResult instantiates a new PolicyAbstractConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigResultWithDefaults

`func NewPolicyAbstractConfigResultWithDefaults() *PolicyAbstractConfigResult`

NewPolicyAbstractConfigResultWithDefaults instantiates a new PolicyAbstractConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


