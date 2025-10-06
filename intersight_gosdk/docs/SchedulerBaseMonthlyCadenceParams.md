# SchedulerBaseMonthlyCadenceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "scheduler.MonthlyCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "scheduler.MonthlyCadenceParams"]
**RunEvery** | Pointer to **int64** | Run every month by default if not specified. | [optional] [default to 1]

## Methods

### NewSchedulerBaseMonthlyCadenceParams

`func NewSchedulerBaseMonthlyCadenceParams(classId string, objectType string, ) *SchedulerBaseMonthlyCadenceParams`

NewSchedulerBaseMonthlyCadenceParams instantiates a new SchedulerBaseMonthlyCadenceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerBaseMonthlyCadenceParamsWithDefaults

`func NewSchedulerBaseMonthlyCadenceParamsWithDefaults() *SchedulerBaseMonthlyCadenceParams`

NewSchedulerBaseMonthlyCadenceParamsWithDefaults instantiates a new SchedulerBaseMonthlyCadenceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerBaseMonthlyCadenceParams) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerBaseMonthlyCadenceParams) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerBaseMonthlyCadenceParams) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerBaseMonthlyCadenceParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerBaseMonthlyCadenceParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerBaseMonthlyCadenceParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunEvery

`func (o *SchedulerBaseMonthlyCadenceParams) GetRunEvery() int64`

GetRunEvery returns the RunEvery field if non-nil, zero value otherwise.

### GetRunEveryOk

`func (o *SchedulerBaseMonthlyCadenceParams) GetRunEveryOk() (*int64, bool)`

GetRunEveryOk returns a tuple with the RunEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvery

`func (o *SchedulerBaseMonthlyCadenceParams) SetRunEvery(v int64)`

SetRunEvery sets RunEvery field to given value.

### HasRunEvery

`func (o *SchedulerBaseMonthlyCadenceParams) HasRunEvery() bool`

HasRunEvery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


