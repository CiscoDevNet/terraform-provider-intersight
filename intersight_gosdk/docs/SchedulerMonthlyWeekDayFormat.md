# SchedulerMonthlyWeekDayFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **[]string** |  | [optional] 
**WeekOfMonth** | Pointer to **int64** | The week of the month, 1 through 5. | [optional] 

## Methods

### NewSchedulerMonthlyWeekDayFormat

`func NewSchedulerMonthlyWeekDayFormat() *SchedulerMonthlyWeekDayFormat`

NewSchedulerMonthlyWeekDayFormat instantiates a new SchedulerMonthlyWeekDayFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerMonthlyWeekDayFormatWithDefaults

`func NewSchedulerMonthlyWeekDayFormatWithDefaults() *SchedulerMonthlyWeekDayFormat`

NewSchedulerMonthlyWeekDayFormatWithDefaults instantiates a new SchedulerMonthlyWeekDayFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *SchedulerMonthlyWeekDayFormat) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SchedulerMonthlyWeekDayFormat) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SchedulerMonthlyWeekDayFormat) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SchedulerMonthlyWeekDayFormat) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *SchedulerMonthlyWeekDayFormat) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *SchedulerMonthlyWeekDayFormat) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetWeekOfMonth

`func (o *SchedulerMonthlyWeekDayFormat) GetWeekOfMonth() int64`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *SchedulerMonthlyWeekDayFormat) GetWeekOfMonthOk() (*int64, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *SchedulerMonthlyWeekDayFormat) SetWeekOfMonth(v int64)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *SchedulerMonthlyWeekDayFormat) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


