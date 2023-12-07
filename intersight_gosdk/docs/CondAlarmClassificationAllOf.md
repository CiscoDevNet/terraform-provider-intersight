# CondAlarmClassificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmClassification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmClassification"]
**AppliesTo** | Pointer to [**[]CondAlarmClassificationEligibility**](CondAlarmClassificationEligibility.md) |  | [optional] 
**Description** | Pointer to **string** | A description about the alarms group that usually gives what kind of alarms are part of this classification. | [optional] [readonly] 
**Name** | Pointer to **string** | The name that uniquely identifies the alarm classificaton. | [optional] [readonly] 
**AlarmDefinitions** | Pointer to [**[]CondAlarmDefinitionRelationship**](CondAlarmDefinitionRelationship.md) | An array of relationships to condAlarmDefinition resources. | [optional] [readonly] 

## Methods

### NewCondAlarmClassificationAllOf

`func NewCondAlarmClassificationAllOf(classId string, objectType string, ) *CondAlarmClassificationAllOf`

NewCondAlarmClassificationAllOf instantiates a new CondAlarmClassificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmClassificationAllOfWithDefaults

`func NewCondAlarmClassificationAllOfWithDefaults() *CondAlarmClassificationAllOf`

NewCondAlarmClassificationAllOfWithDefaults instantiates a new CondAlarmClassificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmClassificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmClassificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmClassificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmClassificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmClassificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmClassificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppliesTo

`func (o *CondAlarmClassificationAllOf) GetAppliesTo() []CondAlarmClassificationEligibility`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CondAlarmClassificationAllOf) GetAppliesToOk() (*[]CondAlarmClassificationEligibility, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CondAlarmClassificationAllOf) SetAppliesTo(v []CondAlarmClassificationEligibility)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CondAlarmClassificationAllOf) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *CondAlarmClassificationAllOf) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *CondAlarmClassificationAllOf) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetDescription

`func (o *CondAlarmClassificationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarmClassificationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarmClassificationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarmClassificationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CondAlarmClassificationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarmClassificationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarmClassificationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarmClassificationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlarmDefinitions

`func (o *CondAlarmClassificationAllOf) GetAlarmDefinitions() []CondAlarmDefinitionRelationship`

GetAlarmDefinitions returns the AlarmDefinitions field if non-nil, zero value otherwise.

### GetAlarmDefinitionsOk

`func (o *CondAlarmClassificationAllOf) GetAlarmDefinitionsOk() (*[]CondAlarmDefinitionRelationship, bool)`

GetAlarmDefinitionsOk returns a tuple with the AlarmDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmDefinitions

`func (o *CondAlarmClassificationAllOf) SetAlarmDefinitions(v []CondAlarmDefinitionRelationship)`

SetAlarmDefinitions sets AlarmDefinitions field to given value.

### HasAlarmDefinitions

`func (o *CondAlarmClassificationAllOf) HasAlarmDefinitions() bool`

HasAlarmDefinitions returns a boolean if a field has been set.

### SetAlarmDefinitionsNil

`func (o *CondAlarmClassificationAllOf) SetAlarmDefinitionsNil(b bool)`

 SetAlarmDefinitionsNil sets the value for AlarmDefinitions to be an explicit nil

### UnsetAlarmDefinitions
`func (o *CondAlarmClassificationAllOf) UnsetAlarmDefinitions()`

UnsetAlarmDefinitions ensures that no value is present for AlarmDefinitions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


