# TelemetryDruidGroupByResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 timestamp. | [optional] 
**Version** | Pointer to **string** | The version of the Druid GroupBy Engine used in query | [optional] 
**Event** | Pointer to **map[string]interface{}** | Grouped aggregate dimension(s) with values | [optional] 

## Methods

### NewTelemetryDruidGroupByResult

`func NewTelemetryDruidGroupByResult() *TelemetryDruidGroupByResult`

NewTelemetryDruidGroupByResult instantiates a new TelemetryDruidGroupByResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidGroupByResultWithDefaults

`func NewTelemetryDruidGroupByResultWithDefaults() *TelemetryDruidGroupByResult`

NewTelemetryDruidGroupByResultWithDefaults instantiates a new TelemetryDruidGroupByResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TelemetryDruidGroupByResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TelemetryDruidGroupByResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TelemetryDruidGroupByResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TelemetryDruidGroupByResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *TelemetryDruidGroupByResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TelemetryDruidGroupByResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TelemetryDruidGroupByResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TelemetryDruidGroupByResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvent

`func (o *TelemetryDruidGroupByResult) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TelemetryDruidGroupByResult) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TelemetryDruidGroupByResult) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TelemetryDruidGroupByResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


