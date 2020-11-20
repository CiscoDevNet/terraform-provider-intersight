# TelemetryDruidSegmentMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An identifier for the metadata | [optional] 
**Intervals** | Pointer to **string** | String representation of the interval queried | [optional] 
**Columns** | Pointer to **map[string]interface{}** | A map of columns and their properties | [optional] 
**Aggregators** | Pointer to **map[string]interface{}** | A map of metrics and their properties | [optional] 
**QueryGranularity** | Pointer to **map[string]interface{}** | query granularity of data stored in segments, may be &#39;null&#39; | [optional] 
**Size** | Pointer to **int32** | estimated total segment byte size as if stored as text | [optional] 
**NumRows** | Pointer to **int32** | number of rows stored in segment | [optional] 

## Methods

### NewTelemetryDruidSegmentMetadataResult

`func NewTelemetryDruidSegmentMetadataResult() *TelemetryDruidSegmentMetadataResult`

NewTelemetryDruidSegmentMetadataResult instantiates a new TelemetryDruidSegmentMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSegmentMetadataResultWithDefaults

`func NewTelemetryDruidSegmentMetadataResultWithDefaults() *TelemetryDruidSegmentMetadataResult`

NewTelemetryDruidSegmentMetadataResultWithDefaults instantiates a new TelemetryDruidSegmentMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TelemetryDruidSegmentMetadataResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TelemetryDruidSegmentMetadataResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TelemetryDruidSegmentMetadataResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TelemetryDruidSegmentMetadataResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntervals

`func (o *TelemetryDruidSegmentMetadataResult) GetIntervals() string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidSegmentMetadataResult) GetIntervalsOk() (*string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidSegmentMetadataResult) SetIntervals(v string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *TelemetryDruidSegmentMetadataResult) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetColumns

`func (o *TelemetryDruidSegmentMetadataResult) GetColumns() map[string]interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidSegmentMetadataResult) GetColumnsOk() (*map[string]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidSegmentMetadataResult) SetColumns(v map[string]interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryDruidSegmentMetadataResult) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetAggregators

`func (o *TelemetryDruidSegmentMetadataResult) GetAggregators() map[string]interface{}`

GetAggregators returns the Aggregators field if non-nil, zero value otherwise.

### GetAggregatorsOk

`func (o *TelemetryDruidSegmentMetadataResult) GetAggregatorsOk() (*map[string]interface{}, bool)`

GetAggregatorsOk returns a tuple with the Aggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregators

`func (o *TelemetryDruidSegmentMetadataResult) SetAggregators(v map[string]interface{})`

SetAggregators sets Aggregators field to given value.

### HasAggregators

`func (o *TelemetryDruidSegmentMetadataResult) HasAggregators() bool`

HasAggregators returns a boolean if a field has been set.

### GetQueryGranularity

`func (o *TelemetryDruidSegmentMetadataResult) GetQueryGranularity() map[string]interface{}`

GetQueryGranularity returns the QueryGranularity field if non-nil, zero value otherwise.

### GetQueryGranularityOk

`func (o *TelemetryDruidSegmentMetadataResult) GetQueryGranularityOk() (*map[string]interface{}, bool)`

GetQueryGranularityOk returns a tuple with the QueryGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryGranularity

`func (o *TelemetryDruidSegmentMetadataResult) SetQueryGranularity(v map[string]interface{})`

SetQueryGranularity sets QueryGranularity field to given value.

### HasQueryGranularity

`func (o *TelemetryDruidSegmentMetadataResult) HasQueryGranularity() bool`

HasQueryGranularity returns a boolean if a field has been set.

### GetSize

`func (o *TelemetryDruidSegmentMetadataResult) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TelemetryDruidSegmentMetadataResult) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TelemetryDruidSegmentMetadataResult) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TelemetryDruidSegmentMetadataResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetNumRows

`func (o *TelemetryDruidSegmentMetadataResult) GetNumRows() int32`

GetNumRows returns the NumRows field if non-nil, zero value otherwise.

### GetNumRowsOk

`func (o *TelemetryDruidSegmentMetadataResult) GetNumRowsOk() (*int32, bool)`

GetNumRowsOk returns a tuple with the NumRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRows

`func (o *TelemetryDruidSegmentMetadataResult) SetNumRows(v int32)`

SetNumRows sets NumRows field to given value.

### HasNumRows

`func (o *TelemetryDruidSegmentMetadataResult) HasNumRows() bool`

HasNumRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


