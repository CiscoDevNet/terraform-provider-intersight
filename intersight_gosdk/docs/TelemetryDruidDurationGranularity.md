# TelemetryDruidDurationGranularity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the type of granularity. | 
**Duration** | **int64** | The duration in milliseconds. | 
**Origin** | Pointer to **time.Time** | An optional value specifying when to start counting time buckets from. The default value is 1970-01-01T00:00:00Z. | [optional] 

## Methods

### NewTelemetryDruidDurationGranularity

`func NewTelemetryDruidDurationGranularity(type_ string, duration int64, ) *TelemetryDruidDurationGranularity`

NewTelemetryDruidDurationGranularity instantiates a new TelemetryDruidDurationGranularity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDurationGranularityWithDefaults

`func NewTelemetryDruidDurationGranularityWithDefaults() *TelemetryDruidDurationGranularity`

NewTelemetryDruidDurationGranularityWithDefaults instantiates a new TelemetryDruidDurationGranularity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidDurationGranularity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidDurationGranularity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidDurationGranularity) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *TelemetryDruidDurationGranularity) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TelemetryDruidDurationGranularity) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TelemetryDruidDurationGranularity) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetOrigin

`func (o *TelemetryDruidDurationGranularity) GetOrigin() time.Time`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *TelemetryDruidDurationGranularity) GetOriginOk() (*time.Time, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *TelemetryDruidDurationGranularity) SetOrigin(v time.Time)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *TelemetryDruidDurationGranularity) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


