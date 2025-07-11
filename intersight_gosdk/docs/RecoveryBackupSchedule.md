# RecoveryBackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recovery.BackupSchedule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recovery.BackupSchedule"]
**ExecutionTime** | Pointer to **time.Time** | The daily backup run time. | [optional] 
**FrequencyUnit** | Pointer to **string** | Backup schedule frequency. * &#x60;Daily&#x60; - Daily backup scheduling option. * &#x60;Periodic&#x60; - Hourly backup scheduling option. | [optional] [default to "Daily"]
**Hours** | Pointer to **int32** | Back schedule frequency (in hours). * &#x60;8&#x60; - The backup interval is 8 hours. * &#x60;4&#x60; - The backup interval is 4 hours. * &#x60;12&#x60; - The backup interval is 12 hours. * &#x60;16&#x60; - The backup interval is 16 hours. * &#x60;20&#x60; - The backup interval is 20 hours. | [optional] [default to 8]

## Methods

### NewRecoveryBackupSchedule

`func NewRecoveryBackupSchedule(classId string, objectType string, ) *RecoveryBackupSchedule`

NewRecoveryBackupSchedule instantiates a new RecoveryBackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryBackupScheduleWithDefaults

`func NewRecoveryBackupScheduleWithDefaults() *RecoveryBackupSchedule`

NewRecoveryBackupScheduleWithDefaults instantiates a new RecoveryBackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecoveryBackupSchedule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecoveryBackupSchedule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecoveryBackupSchedule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecoveryBackupSchedule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecoveryBackupSchedule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecoveryBackupSchedule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


