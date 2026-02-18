# CondAlarmRuleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmRuleExpression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmRuleExpression"]
**Operator** | Pointer to **string** | The operator to apply. Operators supported are: eq, contains, in. | [optional] 
**Property** | Pointer to **string** | The property name keyword to filter on. For a list of supported property keywords see the Intersight Help Center. | [optional] 
**Value** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCondAlarmRuleExpression

`func NewCondAlarmRuleExpression(classId string, objectType string, ) *CondAlarmRuleExpression`

NewCondAlarmRuleExpression instantiates a new CondAlarmRuleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmRuleExpressionWithDefaults

`func NewCondAlarmRuleExpressionWithDefaults() *CondAlarmRuleExpression`

NewCondAlarmRuleExpressionWithDefaults instantiates a new CondAlarmRuleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmRuleExpression) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmRuleExpression) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmRuleExpression) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmRuleExpression) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmRuleExpression) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmRuleExpression) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperator

`func (o *CondAlarmRuleExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CondAlarmRuleExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CondAlarmRuleExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CondAlarmRuleExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetProperty

`func (o *CondAlarmRuleExpression) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CondAlarmRuleExpression) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CondAlarmRuleExpression) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CondAlarmRuleExpression) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *CondAlarmRuleExpression) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CondAlarmRuleExpression) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CondAlarmRuleExpression) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CondAlarmRuleExpression) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CondAlarmRuleExpression) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CondAlarmRuleExpression) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


