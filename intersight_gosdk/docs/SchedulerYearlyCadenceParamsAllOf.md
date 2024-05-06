# SchedulerYearlyCadenceParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.YearlyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.YearlyCadenceParams"]
**MonthOfYear** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSchedulerYearlyCadenceParamsAllOf

`func NewSchedulerYearlyCadenceParamsAllOf(classId string, objectType string, ) *SchedulerYearlyCadenceParamsAllOf`

NewSchedulerYearlyCadenceParamsAllOf instantiates a new SchedulerYearlyCadenceParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerYearlyCadenceParamsAllOfWithDefaults

`func NewSchedulerYearlyCadenceParamsAllOfWithDefaults() *SchedulerYearlyCadenceParamsAllOf`

NewSchedulerYearlyCadenceParamsAllOfWithDefaults instantiates a new SchedulerYearlyCadenceParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerYearlyCadenceParamsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerYearlyCadenceParamsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerYearlyCadenceParamsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerYearlyCadenceParamsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerYearlyCadenceParamsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerYearlyCadenceParamsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMonthOfYear

`func (o *SchedulerYearlyCadenceParamsAllOf) GetMonthOfYear() []string`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *SchedulerYearlyCadenceParamsAllOf) GetMonthOfYearOk() (*[]string, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *SchedulerYearlyCadenceParamsAllOf) SetMonthOfYear(v []string)`

SetMonthOfYear sets MonthOfYear field to given value.

### HasMonthOfYear

`func (o *SchedulerYearlyCadenceParamsAllOf) HasMonthOfYear() bool`

HasMonthOfYear returns a boolean if a field has been set.

### SetMonthOfYearNil

`func (o *SchedulerYearlyCadenceParamsAllOf) SetMonthOfYearNil(b bool)`

 SetMonthOfYearNil sets the value for MonthOfYear to be an explicit nil

### UnsetMonthOfYear
`func (o *SchedulerYearlyCadenceParamsAllOf) UnsetMonthOfYear()`

UnsetMonthOfYear ensures that no value is present for MonthOfYear, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


