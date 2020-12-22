# TelemetryDruidTopNResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp. | [optional] 
**Result** | Pointer to **[]map[string]interface{}** | A sorted list of dimension values | [optional] 

## Methods

### NewTelemetryDruidTopNResult

`func NewTelemetryDruidTopNResult() *TelemetryDruidTopNResult`

NewTelemetryDruidTopNResult instantiates a new TelemetryDruidTopNResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTopNResultWithDefaults

`func NewTelemetryDruidTopNResultWithDefaults() *TelemetryDruidTopNResult`

NewTelemetryDruidTopNResultWithDefaults instantiates a new TelemetryDruidTopNResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TelemetryDruidTopNResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TelemetryDruidTopNResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TelemetryDruidTopNResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TelemetryDruidTopNResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResult

`func (o *TelemetryDruidTopNResult) GetResult() []map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TelemetryDruidTopNResult) GetResultOk() (*[]map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TelemetryDruidTopNResult) SetResult(v []map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TelemetryDruidTopNResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


