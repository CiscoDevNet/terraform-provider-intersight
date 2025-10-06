# ResourceOdataRuleSetQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.OdataRuleSetQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.OdataRuleSetQualifier"]
**Rules** | Pointer to [**[]ResourceSelector**](ResourceSelector.md) |  | [optional] 

## Methods

### NewResourceOdataRuleSetQualifier

`func NewResourceOdataRuleSetQualifier(classId string, objectType string, ) *ResourceOdataRuleSetQualifier`

NewResourceOdataRuleSetQualifier instantiates a new ResourceOdataRuleSetQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOdataRuleSetQualifierWithDefaults

`func NewResourceOdataRuleSetQualifierWithDefaults() *ResourceOdataRuleSetQualifier`

NewResourceOdataRuleSetQualifierWithDefaults instantiates a new ResourceOdataRuleSetQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceOdataRuleSetQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceOdataRuleSetQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceOdataRuleSetQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceOdataRuleSetQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceOdataRuleSetQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceOdataRuleSetQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRules

`func (o *ResourceOdataRuleSetQualifier) GetRules() []ResourceSelector`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ResourceOdataRuleSetQualifier) GetRulesOk() (*[]ResourceSelector, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ResourceOdataRuleSetQualifier) SetRules(v []ResourceSelector)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ResourceOdataRuleSetQualifier) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *ResourceOdataRuleSetQualifier) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *ResourceOdataRuleSetQualifier) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


