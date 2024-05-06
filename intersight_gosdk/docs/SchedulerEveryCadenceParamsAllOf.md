# SchedulerEveryCadenceParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "scheduler.EveryCadenceParams"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "scheduler.EveryCadenceParams"]
**Interval** | Pointer to **string** | An interval specified as string where valid time units are \&quot;ns\&quot;, \&quot;us\&quot;, \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. The calender calculations use a gregorian calendar with no leap seconds. The default is 24h. | [optional] 

## Methods

### NewSchedulerEveryCadenceParamsAllOf

`func NewSchedulerEveryCadenceParamsAllOf(classId string, objectType string, ) *SchedulerEveryCadenceParamsAllOf`

NewSchedulerEveryCadenceParamsAllOf instantiates a new SchedulerEveryCadenceParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerEveryCadenceParamsAllOfWithDefaults

`func NewSchedulerEveryCadenceParamsAllOfWithDefaults() *SchedulerEveryCadenceParamsAllOf`

NewSchedulerEveryCadenceParamsAllOfWithDefaults instantiates a new SchedulerEveryCadenceParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SchedulerEveryCadenceParamsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SchedulerEveryCadenceParamsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SchedulerEveryCadenceParamsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SchedulerEveryCadenceParamsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SchedulerEveryCadenceParamsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SchedulerEveryCadenceParamsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterval

`func (o *SchedulerEveryCadenceParamsAllOf) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SchedulerEveryCadenceParamsAllOf) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SchedulerEveryCadenceParamsAllOf) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SchedulerEveryCadenceParamsAllOf) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


