# CondAlarmSuppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmSuppression"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmSuppression"]
**Description** | Pointer to **string** | User given description on why the suppression is enabled at this entity. | [optional] 
**Classifications** | Pointer to [**[]CondAlarmClassificationRelationship**](CondAlarmClassificationRelationship.md) | An array of relationships to condAlarmClassification resources. | [optional] 
**Entity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


