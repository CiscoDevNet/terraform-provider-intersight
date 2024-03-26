# CondAlarmClassificationEligibilityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmClassificationEligibility"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmClassificationEligibility"]
**EntityType** | Pointer to **string** | The Intersight MO type on which a classification can be suppressed. | [optional] 

## Methods

### NewCondAlarmClassificationEligibilityAllOf

`func NewCondAlarmClassificationEligibilityAllOf(classId string, objectType string, ) *CondAlarmClassificationEligibilityAllOf`

NewCondAlarmClassificationEligibilityAllOf instantiates a new CondAlarmClassificationEligibilityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmClassificationEligibilityAllOfWithDefaults

`func NewCondAlarmClassificationEligibilityAllOfWithDefaults() *CondAlarmClassificationEligibilityAllOf`

NewCondAlarmClassificationEligibilityAllOfWithDefaults instantiates a new CondAlarmClassificationEligibilityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmClassificationEligibilityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmClassificationEligibilityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmClassificationEligibilityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmClassificationEligibilityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmClassificationEligibilityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmClassificationEligibilityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEntityType

`func (o *CondAlarmClassificationEligibilityAllOf) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CondAlarmClassificationEligibilityAllOf) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CondAlarmClassificationEligibilityAllOf) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CondAlarmClassificationEligibilityAllOf) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


