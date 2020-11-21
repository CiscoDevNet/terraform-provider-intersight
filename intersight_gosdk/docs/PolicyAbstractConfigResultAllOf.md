# PolicyAbstractConfigResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ConfigStage** | Pointer to **string** | The current running stage of the configuration or workflow. | [optional] 
**ConfigState** | Pointer to **string** | Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored. | [optional] 
**ValidationState** | Pointer to **string** | Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored. | [optional] 

## Methods

### NewPolicyAbstractConfigResultAllOf

`func NewPolicyAbstractConfigResultAllOf(classId string, objectType string, ) *PolicyAbstractConfigResultAllOf`

NewPolicyAbstractConfigResultAllOf instantiates a new PolicyAbstractConfigResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigResultAllOfWithDefaults

`func NewPolicyAbstractConfigResultAllOfWithDefaults() *PolicyAbstractConfigResultAllOf`

NewPolicyAbstractConfigResultAllOfWithDefaults instantiates a new PolicyAbstractConfigResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractConfigResultAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractConfigResultAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractConfigResultAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractConfigResultAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractConfigResultAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractConfigResultAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigStage

`func (o *PolicyAbstractConfigResultAllOf) GetConfigStage() string`

GetConfigStage returns the ConfigStage field if non-nil, zero value otherwise.

### GetConfigStageOk

`func (o *PolicyAbstractConfigResultAllOf) GetConfigStageOk() (*string, bool)`

GetConfigStageOk returns a tuple with the ConfigStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStage

`func (o *PolicyAbstractConfigResultAllOf) SetConfigStage(v string)`

SetConfigStage sets ConfigStage field to given value.

### HasConfigStage

`func (o *PolicyAbstractConfigResultAllOf) HasConfigStage() bool`

HasConfigStage returns a boolean if a field has been set.

### GetConfigState

`func (o *PolicyAbstractConfigResultAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyAbstractConfigResultAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyAbstractConfigResultAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyAbstractConfigResultAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetValidationState

`func (o *PolicyAbstractConfigResultAllOf) GetValidationState() string`

GetValidationState returns the ValidationState field if non-nil, zero value otherwise.

### GetValidationStateOk

`func (o *PolicyAbstractConfigResultAllOf) GetValidationStateOk() (*string, bool)`

GetValidationStateOk returns a tuple with the ValidationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationState

`func (o *PolicyAbstractConfigResultAllOf) SetValidationState(v string)`

SetValidationState sets ValidationState field to given value.

### HasValidationState

`func (o *PolicyAbstractConfigResultAllOf) HasValidationState() bool`

HasValidationState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


