# SchedulerBaseMonthlyCadenceParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**CustomDayOfMonth** | Pointer to **string** | Significant business days, such as days when reports are generated for analysis. * &#x60;None&#x60; - Placeholder. One of the following two fields must be selected. * &#x60;FirstWeekDay&#x60; - First week day of the month. * &#x60;MonthLastDay&#x60; - The last day of the month. | [optional] [default to "None"]
**DayOfMonth** | Pointer to **[]int64** |  | [optional] 
**WeekDayFormat** | Pointer to [**NullableSchedulerMonthlyWeekDayFormat**](SchedulerMonthlyWeekDayFormat.md) |  | [optional] 

## Methods

### NewSchedulerBaseMonthlyCadenceParamsAllOf

`func NewSchedulerBaseMonthlyCadenceParamsAllOf(classId string, objectType string, ) *SchedulerBaseMonthlyCadenceParamsAllOf`

NewSchedulerBaseMonthlyCadenceParamsAllOf instantiates a new SchedulerBaseMonthlyCadenceParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerBaseMonthlyCadenceParamsAllOfWithDefaults

`func NewSchedulerBaseMonthlyCadenceParamsAllOfWithDefaults() *SchedulerBaseMonthlyCadenceParamsAllOf`

NewSchedulerBaseMonthlyCadenceParamsAllOfWithDefaults instantiates a new SchedulerBaseMonthlyCadenceParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCustomDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetCustomDayOfMonth() string`

GetCustomDayOfMonth returns the CustomDayOfMonth field if non-nil, zero value otherwise.

### GetCustomDayOfMonthOk

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetCustomDayOfMonthOk() (*string, bool)`

GetCustomDayOfMonthOk returns a tuple with the CustomDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetCustomDayOfMonth(v string)`

SetCustomDayOfMonth sets CustomDayOfMonth field to given value.

### HasCustomDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) HasCustomDayOfMonth() bool`

HasCustomDayOfMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetDayOfMonth() []int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetDayOfMonthOk() (*[]int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetDayOfMonth(v []int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonthNil

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetWeekDayFormat

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetWeekDayFormat() SchedulerMonthlyWeekDayFormat`

GetWeekDayFormat returns the WeekDayFormat field if non-nil, zero value otherwise.

### GetWeekDayFormatOk

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) GetWeekDayFormatOk() (*SchedulerMonthlyWeekDayFormat, bool)`

GetWeekDayFormatOk returns a tuple with the WeekDayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDayFormat

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetWeekDayFormat(v SchedulerMonthlyWeekDayFormat)`

SetWeekDayFormat sets WeekDayFormat field to given value.

### HasWeekDayFormat

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) HasWeekDayFormat() bool`

HasWeekDayFormat returns a boolean if a field has been set.

### SetWeekDayFormatNil

`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) SetWeekDayFormatNil(b bool)`

 SetWeekDayFormatNil sets the value for WeekDayFormat to be an explicit nil

### UnsetWeekDayFormat
`func (o *SchedulerBaseMonthlyCadenceParamsAllOf) UnsetWeekDayFormat()`

UnsetWeekDayFormat ensures that no value is present for WeekDayFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


