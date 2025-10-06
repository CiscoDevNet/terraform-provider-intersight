# SchedulerMonthlyCadenceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.MonthlyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.MonthlyCadenceParams"]
**CustomDayOfMonth** | Pointer to **string** | Significant business days, such as days when reports are generated for analysis. * &#x60;None&#x60; - Placeholder. One of the following two fields must be selected. * &#x60;FirstWeekDay&#x60; - First week day of the month. * &#x60;MonthLastDay&#x60; - The last day of the month. * &#x60;FirstWeek&#x60; - Selected weekdays on first week of the month. * &#x60;SecondWeek&#x60; - Selected weekdays on second week of the month. * &#x60;ThirdWeek&#x60; - Selected weekdays on third week of the month. * &#x60;FourthWeek&#x60; - Selected weekdays on fourth week of the month. * &#x60;FifthWeek&#x60; - Selected weekdays on fifth week of the month. | [optional] [default to "None"]
**DayOfMonth** | Pointer to **[]int64** |  | [optional] 
**DayOfWeek** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSchedulerMonthlyCadenceParams

`func NewSchedulerMonthlyCadenceParams(classId string, objectType string, ) *SchedulerMonthlyCadenceParams`

NewSchedulerMonthlyCadenceParams instantiates a new SchedulerMonthlyCadenceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerMonthlyCadenceParamsWithDefaults

`func NewSchedulerMonthlyCadenceParamsWithDefaults() *SchedulerMonthlyCadenceParams`

NewSchedulerMonthlyCadenceParamsWithDefaults instantiates a new SchedulerMonthlyCadenceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerMonthlyCadenceParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerMonthlyCadenceParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerMonthlyCadenceParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerMonthlyCadenceParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerMonthlyCadenceParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerMonthlyCadenceParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCustomDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) GetCustomDayOfMonth() string`

GetCustomDayOfMonth returns the CustomDayOfMonth field if non-nil, zero value otherwise.

### GetCustomDayOfMonthOk

`func (o *SchedulerMonthlyCadenceParams) GetCustomDayOfMonthOk() (*string, bool)`

GetCustomDayOfMonthOk returns a tuple with the CustomDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) SetCustomDayOfMonth(v string)`

SetCustomDayOfMonth sets CustomDayOfMonth field to given value.

### HasCustomDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) HasCustomDayOfMonth() bool`

HasCustomDayOfMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) GetDayOfMonth() []int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *SchedulerMonthlyCadenceParams) GetDayOfMonthOk() (*[]int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) SetDayOfMonth(v []int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *SchedulerMonthlyCadenceParams) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonthNil

`func (o *SchedulerMonthlyCadenceParams) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *SchedulerMonthlyCadenceParams) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetDayOfWeek

`func (o *SchedulerMonthlyCadenceParams) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SchedulerMonthlyCadenceParams) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SchedulerMonthlyCadenceParams) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SchedulerMonthlyCadenceParams) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *SchedulerMonthlyCadenceParams) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *SchedulerMonthlyCadenceParams) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


