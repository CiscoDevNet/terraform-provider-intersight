# BlueprintBaseOdataQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Rules** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBlueprintBaseOdataQualifier

`func NewBlueprintBaseOdataQualifier(classId string, objectType string, ) *BlueprintBaseOdataQualifier`

NewBlueprintBaseOdataQualifier instantiates a new BlueprintBaseOdataQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBaseOdataQualifierWithDefaults

`func NewBlueprintBaseOdataQualifierWithDefaults() *BlueprintBaseOdataQualifier`

NewBlueprintBaseOdataQualifierWithDefaults instantiates a new BlueprintBaseOdataQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BlueprintBaseOdataQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BlueprintBaseOdataQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BlueprintBaseOdataQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BlueprintBaseOdataQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BlueprintBaseOdataQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BlueprintBaseOdataQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRules

`func (o *BlueprintBaseOdataQualifier) GetRules() []string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BlueprintBaseOdataQualifier) GetRulesOk() (*[]string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BlueprintBaseOdataQualifier) SetRules(v []string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BlueprintBaseOdataQualifier) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *BlueprintBaseOdataQualifier) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BlueprintBaseOdataQualifier) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


