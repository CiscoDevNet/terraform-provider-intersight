# CondAlarmSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmSuppression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmSuppression"]
**AlarmRules** | Pointer to [**[]CondAlarmRuleExpression**](CondAlarmRuleExpression.md) |  | [optional] 
**Description** | Pointer to **string** | User given description on why the suppression is enabled at this entity. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the suppression is enabled by the user or not. The user should be able to toggle this between true and false. The property is set to true when the suppression is created. The user can set this to false to disable the suppression. The suppression rule should be active only if both systemEnabled and enabled are true. | [optional] [default to true]
**EndDate** | Pointer to **time.Time** | The end date for this alarm suppression rule. The date must follow the RFC 3339 format for date and time representation. | [optional] 
**Name** | Pointer to **string** | The name that identifies the alarm suppression. | [optional] 
**OdataFilterInternal** | Pointer to **string** | Odata filter string managed internally. It is built by combining all the rules. | [optional] [readonly] 
**RulesOperator** | Pointer to **string** | Operation that binds all the different rules together. * &#x60;All&#x60; - All is an AND condition applied against the individual conditions. * &#x60;Any&#x60; - Any is an OR condition applied against the individual conditions. | [optional] [default to "All"]
**StartDate** | Pointer to **time.Time** | The start date for enabling this alarm suppression rule. The date must follow the RFC 3339 format for date and time representation. If this date more than 60 seconds in the past, the suppression rule will be rejected. If the date is within 60 seconds of the present time (plus or minus), the suppression will be started immediately. Otherwise, the suppression will be scheduled to start at the requested time. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Classifications** | Pointer to [**[]CondAlarmClassificationRelationship**](CondAlarmClassificationRelationship.md) | An array of relationships to condAlarmClassification resources. | [optional] 
**Entity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewCondAlarmSuppression

`func NewCondAlarmSuppression(classId string, objectType string, ) *CondAlarmSuppression`

NewCondAlarmSuppression instantiates a new CondAlarmSuppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmSuppressionWithDefaults

`func NewCondAlarmSuppressionWithDefaults() *CondAlarmSuppression`

NewCondAlarmSuppressionWithDefaults instantiates a new CondAlarmSuppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmSuppression) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmSuppression) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmSuppression) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmSuppression) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmSuppression) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmSuppression) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmRules

`func (o *CondAlarmSuppression) GetAlarmRules() []CondAlarmRuleExpression`

GetAlarmRules returns the AlarmRules field if non-nil, zero value otherwise.

### GetAlarmRulesOk

`func (o *CondAlarmSuppression) GetAlarmRulesOk() (*[]CondAlarmRuleExpression, bool)`

GetAlarmRulesOk returns a tuple with the AlarmRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmRules

`func (o *CondAlarmSuppression) SetAlarmRules(v []CondAlarmRuleExpression)`

SetAlarmRules sets AlarmRules field to given value.

### HasAlarmRules

`func (o *CondAlarmSuppression) HasAlarmRules() bool`

HasAlarmRules returns a boolean if a field has been set.

### SetAlarmRulesNil

`func (o *CondAlarmSuppression) SetAlarmRulesNil(b bool)`

 SetAlarmRulesNil sets the value for AlarmRules to be an explicit nil

### UnsetAlarmRules
`func (o *CondAlarmSuppression) UnsetAlarmRules()`

UnsetAlarmRules ensures that no value is present for AlarmRules, not even an explicit nil
### GetDescription

`func (o *CondAlarmSuppression) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarmSuppression) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarmSuppression) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarmSuppression) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CondAlarmSuppression) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CondAlarmSuppression) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CondAlarmSuppression) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CondAlarmSuppression) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndDate

`func (o *CondAlarmSuppression) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CondAlarmSuppression) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CondAlarmSuppression) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CondAlarmSuppression) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *CondAlarmSuppression) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarmSuppression) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarmSuppression) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarmSuppression) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOdataFilterInternal

`func (o *CondAlarmSuppression) GetOdataFilterInternal() string`

GetOdataFilterInternal returns the OdataFilterInternal field if non-nil, zero value otherwise.

### GetOdataFilterInternalOk

`func (o *CondAlarmSuppression) GetOdataFilterInternalOk() (*string, bool)`

GetOdataFilterInternalOk returns a tuple with the OdataFilterInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataFilterInternal

`func (o *CondAlarmSuppression) SetOdataFilterInternal(v string)`

SetOdataFilterInternal sets OdataFilterInternal field to given value.

### HasOdataFilterInternal

`func (o *CondAlarmSuppression) HasOdataFilterInternal() bool`

HasOdataFilterInternal returns a boolean if a field has been set.

### GetRulesOperator

`func (o *CondAlarmSuppression) GetRulesOperator() string`

GetRulesOperator returns the RulesOperator field if non-nil, zero value otherwise.

### GetRulesOperatorOk

`func (o *CondAlarmSuppression) GetRulesOperatorOk() (*string, bool)`

GetRulesOperatorOk returns a tuple with the RulesOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesOperator

`func (o *CondAlarmSuppression) SetRulesOperator(v string)`

SetRulesOperator sets RulesOperator field to given value.

### HasRulesOperator

`func (o *CondAlarmSuppression) HasRulesOperator() bool`

HasRulesOperator returns a boolean if a field has been set.

### GetStartDate

`func (o *CondAlarmSuppression) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CondAlarmSuppression) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CondAlarmSuppression) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CondAlarmSuppression) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetAccount

`func (o *CondAlarmSuppression) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CondAlarmSuppression) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CondAlarmSuppression) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CondAlarmSuppression) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CondAlarmSuppression) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CondAlarmSuppression) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetClassifications

`func (o *CondAlarmSuppression) GetClassifications() []CondAlarmClassificationRelationship`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *CondAlarmSuppression) GetClassificationsOk() (*[]CondAlarmClassificationRelationship, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *CondAlarmSuppression) SetClassifications(v []CondAlarmClassificationRelationship)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *CondAlarmSuppression) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### SetClassificationsNil

`func (o *CondAlarmSuppression) SetClassificationsNil(b bool)`

 SetClassificationsNil sets the value for Classifications to be an explicit nil

### UnsetClassifications
`func (o *CondAlarmSuppression) UnsetClassifications()`

UnsetClassifications ensures that no value is present for Classifications, not even an explicit nil
### GetEntity

`func (o *CondAlarmSuppression) GetEntity() MoBaseMoRelationship`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CondAlarmSuppression) GetEntityOk() (*MoBaseMoRelationship, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CondAlarmSuppression) SetEntity(v MoBaseMoRelationship)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CondAlarmSuppression) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *CondAlarmSuppression) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *CondAlarmSuppression) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


