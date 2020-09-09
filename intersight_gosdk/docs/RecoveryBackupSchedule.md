# RecoveryBackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionTime** | Pointer to [**time.Time**](time.Time.md) | The time at which the backup is to be run on a given day. This is used when the frequency unit is daily. | [optional] 
**FrequencyUnit** | Pointer to **string** | The frequency at which the backup schedule must run. * &#x60;Daily&#x60; - Allows the user to run the backup daily at a given time. * &#x60;Periodic&#x60; - Allows the user to run the backup after a certain number of hours. | [optional] [default to "Daily"]
**Hours** | Pointer to **int32** | The frequency, in hours, at which the backup schedule runs. * &#x60;8&#x60; -  * &#x60;4&#x60; -  * &#x60;12&#x60; -  * &#x60;16&#x60; -  * &#x60;20&#x60; - | [optional] [default to 8]

## Methods

### NewRecoveryBackupSchedule

`func NewRecoveryBackupSchedule() *RecoveryBackupSchedule`

NewRecoveryBackupSchedule instantiates a new RecoveryBackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupScheduleWithDefaults

`func NewRecoveryBackupScheduleWithDefaults() *RecoveryBackupSchedule`

NewRecoveryBackupScheduleWithDefaults instantiates a new RecoveryBackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionTime

`func (o *RecoveryBackupSchedule) GetExecutionTime() time.Time`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *RecoveryBackupSchedule) GetExecutionTimeOk() (*time.Time, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *RecoveryBackupSchedule) SetExecutionTime(v time.Time)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *RecoveryBackupSchedule) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetFrequencyUnit

`func (o *RecoveryBackupSchedule) GetFrequencyUnit() string`

GetFrequencyUnit returns the FrequencyUnit field if non-nil, zero value otherwise.

### GetFrequencyUnitOk

`func (o *RecoveryBackupSchedule) GetFrequencyUnitOk() (*string, bool)`

GetFrequencyUnitOk returns a tuple with the FrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyUnit

`func (o *RecoveryBackupSchedule) SetFrequencyUnit(v string)`

SetFrequencyUnit sets FrequencyUnit field to given value.

### HasFrequencyUnit

`func (o *RecoveryBackupSchedule) HasFrequencyUnit() bool`

HasFrequencyUnit returns a boolean if a field has been set.

### GetHours

`func (o *RecoveryBackupSchedule) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *RecoveryBackupSchedule) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *RecoveryBackupSchedule) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *RecoveryBackupSchedule) HasHours() bool`

HasHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


