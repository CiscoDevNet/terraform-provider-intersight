# SchedulerRecurringScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.RecurringScheduleParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.RecurringScheduleParams"]
**Cadence** | Pointer to **string** | Allowed values for a recurring schedule cadence. * &#x60;None&#x60; - No value set for the cadence type (Enum value None). * &#x60;Every&#x60; - Use the &#39;Every&#39; cadence for tasks that need to be run frequently and are relatively small or quick to execute. This could include tasks such as checking the status of a service every 15 minutes, or updating a counter. * &#x60;Daily&#x60; - A Daily cadence allows for a scheduled task to be run every day or every n-interval days. * &#x60;Weekly&#x60; - A Weekly cadence allows for a scheduled task to be run every week or every n-interval weeks on specific days. * &#x60;Monthly&#x60; - A Montly cadence allows for a scheduled task to be run every month on specific days. | [optional] [default to "None"]
**EndAfterOccurrences** | Pointer to **int64** | Specify the number of occurrences (instead of an end-time) for a recurring schedule. | [optional] 
**EndTime** | Pointer to **time.Time** | End time for the recurring schedule. The schedule will not run beyond this time. If using the endAfterOccurrences parameter instead, this field should be set to zero time, i.e, 0001-01-01T00:00:00Z. | [optional] 
**FailureThreshold** | Pointer to **int64** | The maximum number of consecutive failures until the recurring scheduled task is suspended by the system. The default is 1. | [optional] [default to 1]
**Params** | Pointer to [**NullableSchedulerBaseCadenceParams**](SchedulerBaseCadenceParams.md) |  | [optional] 

## Methods

### NewSchedulerRecurringScheduleParams

`func NewSchedulerRecurringScheduleParams(classId string, objectType string, ) *SchedulerRecurringScheduleParams`

NewSchedulerRecurringScheduleParams instantiates a new SchedulerRecurringScheduleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerRecurringScheduleParamsWithDefaults

`func NewSchedulerRecurringScheduleParamsWithDefaults() *SchedulerRecurringScheduleParams`

NewSchedulerRecurringScheduleParamsWithDefaults instantiates a new SchedulerRecurringScheduleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerRecurringScheduleParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerRecurringScheduleParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerRecurringScheduleParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerRecurringScheduleParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerRecurringScheduleParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerRecurringScheduleParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCadence

`func (o *SchedulerRecurringScheduleParams) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *SchedulerRecurringScheduleParams) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *SchedulerRecurringScheduleParams) SetCadence(v string)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *SchedulerRecurringScheduleParams) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetEndAfterOccurrences

`func (o *SchedulerRecurringScheduleParams) GetEndAfterOccurrences() int64`

GetEndAfterOccurrences returns the EndAfterOccurrences field if non-nil, zero value otherwise.

### GetEndAfterOccurrencesOk

`func (o *SchedulerRecurringScheduleParams) GetEndAfterOccurrencesOk() (*int64, bool)`

GetEndAfterOccurrencesOk returns a tuple with the EndAfterOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAfterOccurrences

`func (o *SchedulerRecurringScheduleParams) SetEndAfterOccurrences(v int64)`

SetEndAfterOccurrences sets EndAfterOccurrences field to given value.

### HasEndAfterOccurrences

`func (o *SchedulerRecurringScheduleParams) HasEndAfterOccurrences() bool`

HasEndAfterOccurrences returns a boolean if a field has been set.

### GetEndTime

`func (o *SchedulerRecurringScheduleParams) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SchedulerRecurringScheduleParams) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SchedulerRecurringScheduleParams) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SchedulerRecurringScheduleParams) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *SchedulerRecurringScheduleParams) GetFailureThreshold() int64`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *SchedulerRecurringScheduleParams) GetFailureThresholdOk() (*int64, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *SchedulerRecurringScheduleParams) SetFailureThreshold(v int64)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *SchedulerRecurringScheduleParams) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetParams

`func (o *SchedulerRecurringScheduleParams) GetParams() SchedulerBaseCadenceParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SchedulerRecurringScheduleParams) GetParamsOk() (*SchedulerBaseCadenceParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SchedulerRecurringScheduleParams) SetParams(v SchedulerBaseCadenceParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *SchedulerRecurringScheduleParams) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *SchedulerRecurringScheduleParams) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *SchedulerRecurringScheduleParams) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


