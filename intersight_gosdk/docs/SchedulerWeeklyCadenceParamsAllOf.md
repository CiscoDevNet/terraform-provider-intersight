# SchedulerWeeklyCadenceParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.WeeklyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.WeeklyCadenceParams"]
**DayOfWeek** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSchedulerWeeklyCadenceParamsAllOf

`func NewSchedulerWeeklyCadenceParamsAllOf(classId string, objectType string, ) *SchedulerWeeklyCadenceParamsAllOf`

NewSchedulerWeeklyCadenceParamsAllOf instantiates a new SchedulerWeeklyCadenceParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerWeeklyCadenceParamsAllOfWithDefaults

`func NewSchedulerWeeklyCadenceParamsAllOfWithDefaults() *SchedulerWeeklyCadenceParamsAllOf`

NewSchedulerWeeklyCadenceParamsAllOfWithDefaults instantiates a new SchedulerWeeklyCadenceParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerWeeklyCadenceParamsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerWeeklyCadenceParamsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDayOfWeek

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SchedulerWeeklyCadenceParamsAllOf) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SchedulerWeeklyCadenceParamsAllOf) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SchedulerWeeklyCadenceParamsAllOf) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *SchedulerWeeklyCadenceParamsAllOf) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *SchedulerWeeklyCadenceParamsAllOf) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


