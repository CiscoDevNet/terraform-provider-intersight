# CondAlarmRuleCompoundExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmRuleCompoundExpression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmRuleCompoundExpression"]
**Filter** | Pointer to [**NullableFilterexprFilterExpression**](FilterexprFilterExpression.md) |  | [optional] 

## Methods

### NewCondAlarmRuleCompoundExpression

`func NewCondAlarmRuleCompoundExpression(classId string, objectType string, ) *CondAlarmRuleCompoundExpression`

NewCondAlarmRuleCompoundExpression instantiates a new CondAlarmRuleCompoundExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmRuleCompoundExpressionWithDefaults

`func NewCondAlarmRuleCompoundExpressionWithDefaults() *CondAlarmRuleCompoundExpression`

NewCondAlarmRuleCompoundExpressionWithDefaults instantiates a new CondAlarmRuleCompoundExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmRuleCompoundExpression) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmRuleCompoundExpression) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmRuleCompoundExpression) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmRuleCompoundExpression) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmRuleCompoundExpression) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmRuleCompoundExpression) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilter

`func (o *CondAlarmRuleCompoundExpression) GetFilter() FilterexprFilterExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CondAlarmRuleCompoundExpression) GetFilterOk() (*FilterexprFilterExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CondAlarmRuleCompoundExpression) SetFilter(v FilterexprFilterExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CondAlarmRuleCompoundExpression) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CondAlarmRuleCompoundExpression) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CondAlarmRuleCompoundExpression) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


