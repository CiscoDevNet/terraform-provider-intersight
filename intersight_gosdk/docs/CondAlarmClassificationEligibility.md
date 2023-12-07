# CondAlarmClassificationEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmClassificationEligibility"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmClassificationEligibility"]
**EntityType** | Pointer to **string** | The Intersight MO type on which a classification can be suppressed. | [optional] 

## Methods

### NewCondAlarmClassificationEligibility

`func NewCondAlarmClassificationEligibility(classId string, objectType string, ) *CondAlarmClassificationEligibility`

NewCondAlarmClassificationEligibility instantiates a new CondAlarmClassificationEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmClassificationEligibilityWithDefaults

`func NewCondAlarmClassificationEligibilityWithDefaults() *CondAlarmClassificationEligibility`

NewCondAlarmClassificationEligibilityWithDefaults instantiates a new CondAlarmClassificationEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmClassificationEligibility) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmClassificationEligibility) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmClassificationEligibility) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmClassificationEligibility) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmClassificationEligibility) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmClassificationEligibility) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEntityType

`func (o *CondAlarmClassificationEligibility) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CondAlarmClassificationEligibility) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CondAlarmClassificationEligibility) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CondAlarmClassificationEligibility) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


