# HyperflexHealthCheckSchedulePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HealthCheckSchedulePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HealthCheckSchedulePolicy"]
**LastScheduledOn** | Pointer to **time.Time** | The date and time when this HealthCheck Policy was last enabled. | [optional] 
**LastUnscheduledOn** | Pointer to **time.Time** | The date and time when this HealthCheck Policy was last disabled. | [optional] 
**NextExpectedExecution** | Pointer to **time.Time** | The date and time when the next health check execution is expected. | [optional] 
**PolicyEnabled** | Pointer to **bool** | Indicates whether HealthCheck schedule policy is enabled on the HyperFlex cluster. | [optional] 
**ScheduleInterval** | Pointer to **int32** | The frequency at which the health checks are run on the HyperFlex cluster. * &#x60;86400&#x60; - Execute the health check every 24 hours. * &#x60;43200&#x60; - Execute the health check every 12 hours. * &#x60;21600&#x60; - Execute the health check every 6 hours. * &#x60;10800&#x60; - Execute the health check every 3 hours. * &#x60;3600&#x60; - Execute the health check every 1 hours. * &#x60;300&#x60; - Execute the health check every 5 minutes. * &#x60;0&#x60; - Disable the continuous health check. | [optional] [default to 86400]
**Uuid** | Pointer to **string** | The unique identifier of the health check policy. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHealthCheckSchedulePolicyAllOf

`func NewHyperflexHealthCheckSchedulePolicyAllOf(classId string, objectType string, ) *HyperflexHealthCheckSchedulePolicyAllOf`

NewHyperflexHealthCheckSchedulePolicyAllOf instantiates a new HyperflexHealthCheckSchedulePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHealthCheckSchedulePolicyAllOfWithDefaults

`func NewHyperflexHealthCheckSchedulePolicyAllOfWithDefaults() *HyperflexHealthCheckSchedulePolicyAllOf`

NewHyperflexHealthCheckSchedulePolicyAllOfWithDefaults instantiates a new HyperflexHealthCheckSchedulePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastScheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetLastScheduledOn() time.Time`

GetLastScheduledOn returns the LastScheduledOn field if non-nil, zero value otherwise.

### GetLastScheduledOnOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetLastScheduledOnOk() (*time.Time, bool)`

GetLastScheduledOnOk returns a tuple with the LastScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetLastScheduledOn(v time.Time)`

SetLastScheduledOn sets LastScheduledOn field to given value.

### HasLastScheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasLastScheduledOn() bool`

HasLastScheduledOn returns a boolean if a field has been set.

### GetLastUnscheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetLastUnscheduledOn() time.Time`

GetLastUnscheduledOn returns the LastUnscheduledOn field if non-nil, zero value otherwise.

### GetLastUnscheduledOnOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetLastUnscheduledOnOk() (*time.Time, bool)`

GetLastUnscheduledOnOk returns a tuple with the LastUnscheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUnscheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetLastUnscheduledOn(v time.Time)`

SetLastUnscheduledOn sets LastUnscheduledOn field to given value.

### HasLastUnscheduledOn

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasLastUnscheduledOn() bool`

HasLastUnscheduledOn returns a boolean if a field has been set.

### GetNextExpectedExecution

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetNextExpectedExecution() time.Time`

GetNextExpectedExecution returns the NextExpectedExecution field if non-nil, zero value otherwise.

### GetNextExpectedExecutionOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetNextExpectedExecutionOk() (*time.Time, bool)`

GetNextExpectedExecutionOk returns a tuple with the NextExpectedExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExpectedExecution

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetNextExpectedExecution(v time.Time)`

SetNextExpectedExecution sets NextExpectedExecution field to given value.

### HasNextExpectedExecution

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasNextExpectedExecution() bool`

HasNextExpectedExecution returns a boolean if a field has been set.

### GetPolicyEnabled

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetPolicyEnabled() bool`

GetPolicyEnabled returns the PolicyEnabled field if non-nil, zero value otherwise.

### GetPolicyEnabledOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetPolicyEnabledOk() (*bool, bool)`

GetPolicyEnabledOk returns a tuple with the PolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnabled

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetPolicyEnabled(v bool)`

SetPolicyEnabled sets PolicyEnabled field to given value.

### HasPolicyEnabled

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasPolicyEnabled() bool`

HasPolicyEnabled returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetScheduleInterval() int32`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetScheduleIntervalOk() (*int32, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetScheduleInterval(v int32)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HyperflexHealthCheckSchedulePolicyAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


