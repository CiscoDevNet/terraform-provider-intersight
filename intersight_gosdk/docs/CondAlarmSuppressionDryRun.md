# CondAlarmSuppressionDryRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmSuppressionDryRun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmSuppressionDryRun"]
**AlarmRules** | Pointer to [**[]CondAlarmRuleExpression**](CondAlarmRuleExpression.md) |  | [optional] 
**RulesOperator** | Pointer to **string** | Operation that binds all the different rules together. * &#x60;All&#x60; - All is an AND condition applied against the individual conditions. * &#x60;Any&#x60; - Any is an OR condition applied against the individual conditions. | [optional] [default to "All"]
**Summary** | Pointer to [**NullableCondAlarmSuppressionDryRunSummary**](CondAlarmSuppressionDryRunSummary.md) |  | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewCondAlarmSuppressionDryRun

`func NewCondAlarmSuppressionDryRun(classId string, objectType string, ) *CondAlarmSuppressionDryRun`

NewCondAlarmSuppressionDryRun instantiates a new CondAlarmSuppressionDryRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmSuppressionDryRunWithDefaults

`func NewCondAlarmSuppressionDryRunWithDefaults() *CondAlarmSuppressionDryRun`

NewCondAlarmSuppressionDryRunWithDefaults instantiates a new CondAlarmSuppressionDryRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmSuppressionDryRun) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmSuppressionDryRun) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmSuppressionDryRun) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmSuppressionDryRun) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmSuppressionDryRun) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmSuppressionDryRun) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmRules

`func (o *CondAlarmSuppressionDryRun) GetAlarmRules() []CondAlarmRuleExpression`

GetAlarmRules returns the AlarmRules field if non-nil, zero value otherwise.

### GetAlarmRulesOk

`func (o *CondAlarmSuppressionDryRun) GetAlarmRulesOk() (*[]CondAlarmRuleExpression, bool)`

GetAlarmRulesOk returns a tuple with the AlarmRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRules

`func (o *CondAlarmSuppressionDryRun) SetAlarmRules(v []CondAlarmRuleExpression)`

SetAlarmRules sets AlarmRules field to given value.

### HasAlarmRules

`func (o *CondAlarmSuppressionDryRun) HasAlarmRules() bool`

HasAlarmRules returns a boolean if a field has been set.

### SetAlarmRulesNil

`func (o *CondAlarmSuppressionDryRun) SetAlarmRulesNil(b bool)`

 SetAlarmRulesNil sets the value for AlarmRules to be an explicit nil

### UnsetAlarmRules
`func (o *CondAlarmSuppressionDryRun) UnsetAlarmRules()`

UnsetAlarmRules ensures that no value is present for AlarmRules, not even an explicit nil
### GetRulesOperator

`func (o *CondAlarmSuppressionDryRun) GetRulesOperator() string`

GetRulesOperator returns the RulesOperator field if non-nil, zero value otherwise.

### GetRulesOperatorOk

`func (o *CondAlarmSuppressionDryRun) GetRulesOperatorOk() (*string, bool)`

GetRulesOperatorOk returns a tuple with the RulesOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesOperator

`func (o *CondAlarmSuppressionDryRun) SetRulesOperator(v string)`

SetRulesOperator sets RulesOperator field to given value.

### HasRulesOperator

`func (o *CondAlarmSuppressionDryRun) HasRulesOperator() bool`

HasRulesOperator returns a boolean if a field has been set.

### GetSummary

`func (o *CondAlarmSuppressionDryRun) GetSummary() CondAlarmSuppressionDryRunSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CondAlarmSuppressionDryRun) GetSummaryOk() (*CondAlarmSuppressionDryRunSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CondAlarmSuppressionDryRun) SetSummary(v CondAlarmSuppressionDryRunSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CondAlarmSuppressionDryRun) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *CondAlarmSuppressionDryRun) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *CondAlarmSuppressionDryRun) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetAccount

`func (o *CondAlarmSuppressionDryRun) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CondAlarmSuppressionDryRun) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CondAlarmSuppressionDryRun) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CondAlarmSuppressionDryRun) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CondAlarmSuppressionDryRun) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CondAlarmSuppressionDryRun) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


