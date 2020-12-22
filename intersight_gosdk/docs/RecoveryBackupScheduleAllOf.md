# RecoveryBackupScheduleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.BackupSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.BackupSchedule"]
**ExecutionTime** | Pointer to **time.Time** | The time at which the backup is to be run on a given day. This is used when the frequency unit is daily. | [optional] 
**FrequencyUnit** | Pointer to **string** | The frequency at which the backup schedule must run. * &#x60;Daily&#x60; - Allows the user to run the backup daily at a given time. * &#x60;Periodic&#x60; - Allows the user to run the backup after a certain number of hours. | [optional] [default to "Daily"]
**Hours** | Pointer to **int32** | The frequency, in hours, at which the backup schedule runs. * &#x60;8&#x60; -  * &#x60;4&#x60; -  * &#x60;12&#x60; -  * &#x60;16&#x60; -  * &#x60;20&#x60; - | [optional] [default to 8]

## Methods

### NewRecoveryBackupScheduleAllOf

`func NewRecoveryBackupScheduleAllOf(classId string, objectType string, ) *RecoveryBackupScheduleAllOf`

NewRecoveryBackupScheduleAllOf instantiates a new RecoveryBackupScheduleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupScheduleAllOfWithDefaults

`func NewRecoveryBackupScheduleAllOfWithDefaults() *RecoveryBackupScheduleAllOf`

NewRecoveryBackupScheduleAllOfWithDefaults instantiates a new RecoveryBackupScheduleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryBackupScheduleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryBackupScheduleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryBackupScheduleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryBackupScheduleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryBackupScheduleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryBackupScheduleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExecutionTime

`func (o *RecoveryBackupScheduleAllOf) GetExecutionTime() time.Time`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *RecoveryBackupScheduleAllOf) GetExecutionTimeOk() (*time.Time, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *RecoveryBackupScheduleAllOf) SetExecutionTime(v time.Time)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *RecoveryBackupScheduleAllOf) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### GetFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnit() string`

GetFrequencyUnit returns the FrequencyUnit field if non-nil, zero value otherwise.

### GetFrequencyUnitOk

`func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnitOk() (*string, bool)`

GetFrequencyUnitOk returns a tuple with the FrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) SetFrequencyUnit(v string)`

SetFrequencyUnit sets FrequencyUnit field to given value.

### HasFrequencyUnit

`func (o *RecoveryBackupScheduleAllOf) HasFrequencyUnit() bool`

HasFrequencyUnit returns a boolean if a field has been set.

### GetHours

`func (o *RecoveryBackupScheduleAllOf) GetHours() int32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *RecoveryBackupScheduleAllOf) GetHoursOk() (*int32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *RecoveryBackupScheduleAllOf) SetHours(v int32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *RecoveryBackupScheduleAllOf) HasHours() bool`

HasHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


