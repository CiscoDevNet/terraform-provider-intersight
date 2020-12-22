# TelemetryDruidSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp. | [optional] 
**Result** | Pointer to **[]map[string]interface{}** | A list of dimension values | [optional] 

## Methods

### NewTelemetryDruidSearchResult

`func NewTelemetryDruidSearchResult() *TelemetryDruidSearchResult`

NewTelemetryDruidSearchResult instantiates a new TelemetryDruidSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSearchResultWithDefaults

`func NewTelemetryDruidSearchResultWithDefaults() *TelemetryDruidSearchResult`

NewTelemetryDruidSearchResultWithDefaults instantiates a new TelemetryDruidSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TelemetryDruidSearchResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TelemetryDruidSearchResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TelemetryDruidSearchResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TelemetryDruidSearchResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResult

`func (o *TelemetryDruidSearchResult) GetResult() []map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TelemetryDruidSearchResult) GetResultOk() (*[]map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TelemetryDruidSearchResult) SetResult(v []map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TelemetryDruidSearchResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


