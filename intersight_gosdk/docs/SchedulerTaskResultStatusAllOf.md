# SchedulerTaskResultStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.TaskResultStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.TaskResultStatus"]
**Reason** | Pointer to **string** | The status reason, used when a scheduled task fails or was suspended by the system. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the current scheduled task. * &#x60;None&#x60; - No status is set (default). * &#x60;Scheduled&#x60; - The status is set when a task is scheduled. * &#x60;Running&#x60; - The status is set when a task is running. * &#x60;Completed&#x60; - The status is set when a task is complete. * &#x60;Failed&#x60; - The status is set when a task fails. * &#x60;Suspended&#x60; - The status is set when a task is suspended. * &#x60;Skipped&#x60; - The status is set when a task is skipped because the previous task is still running. | [optional] [readonly] [default to "None"]

## Methods

### NewSchedulerTaskResultStatusAllOf

`func NewSchedulerTaskResultStatusAllOf(classId string, objectType string, ) *SchedulerTaskResultStatusAllOf`

NewSchedulerTaskResultStatusAllOf instantiates a new SchedulerTaskResultStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTaskResultStatusAllOfWithDefaults

`func NewSchedulerTaskResultStatusAllOfWithDefaults() *SchedulerTaskResultStatusAllOf`

NewSchedulerTaskResultStatusAllOfWithDefaults instantiates a new SchedulerTaskResultStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerTaskResultStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerTaskResultStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerTaskResultStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerTaskResultStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerTaskResultStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerTaskResultStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReason

`func (o *SchedulerTaskResultStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SchedulerTaskResultStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SchedulerTaskResultStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SchedulerTaskResultStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *SchedulerTaskResultStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchedulerTaskResultStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchedulerTaskResultStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchedulerTaskResultStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


