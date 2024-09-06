# SchedulerDailyCadenceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.DailyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.DailyCadenceParams"]
**RunEvery** | Pointer to **int64** | Run every day by default if not specified. | [optional] [default to 1]

## Methods

### NewSchedulerDailyCadenceParams

`func NewSchedulerDailyCadenceParams(classId string, objectType string, ) *SchedulerDailyCadenceParams`

NewSchedulerDailyCadenceParams instantiates a new SchedulerDailyCadenceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerDailyCadenceParamsWithDefaults

`func NewSchedulerDailyCadenceParamsWithDefaults() *SchedulerDailyCadenceParams`

NewSchedulerDailyCadenceParamsWithDefaults instantiates a new SchedulerDailyCadenceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerDailyCadenceParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerDailyCadenceParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerDailyCadenceParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerDailyCadenceParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerDailyCadenceParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerDailyCadenceParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunEvery

`func (o *SchedulerDailyCadenceParams) GetRunEvery() int64`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *SchedulerDailyCadenceParams) GetRunEveryOk() (*int64, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *SchedulerDailyCadenceParams) SetRunEvery(v int64)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *SchedulerDailyCadenceParams) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


