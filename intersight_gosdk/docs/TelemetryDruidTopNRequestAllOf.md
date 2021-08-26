# TelemetryDruidTopNRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. | 
**Granularity** | [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | 
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Aggregations** | Pointer to [**[]TelemetryDruidAggregator**](TelemetryDruidAggregator.md) | Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket. | [optional] 
**PostAggregations** | Pointer to [**[]TelemetryDruidPostAggregator**](TelemetryDruidPostAggregator.md) | Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires. | [optional] 
**Dimension** | [**TelemetryDruidDimensionSpec**](TelemetryDruidDimensionSpec.md) |  | 
**Threshold** | **int32** | An integer defining the N in the topN (i.e. how many results you want in the top list). | 
**Metric** | [**TelemetryDruidTopNMetricSpec**](TelemetryDruidTopNMetricSpec.md) |  | 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidTopNRequestAllOf

`func NewTelemetryDruidTopNRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, dimension TelemetryDruidDimensionSpec, threshold int32, metric TelemetryDruidTopNMetricSpec, ) *TelemetryDruidTopNRequestAllOf`

NewTelemetryDruidTopNRequestAllOf instantiates a new TelemetryDruidTopNRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTopNRequestAllOfWithDefaults

`func NewTelemetryDruidTopNRequestAllOfWithDefaults() *TelemetryDruidTopNRequestAllOf`

NewTelemetryDruidTopNRequestAllOfWithDefaults instantiates a new TelemetryDruidTopNRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *TelemetryDruidTopNRequestAllOf) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidTopNRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidTopNRequestAllOf) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetIntervals

`func (o *TelemetryDruidTopNRequestAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidTopNRequestAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidTopNRequestAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetGranularity

`func (o *TelemetryDruidTopNRequestAllOf) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidTopNRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidTopNRequestAllOf) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *TelemetryDruidTopNRequestAllOf) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidTopNRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidTopNRequestAllOf) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidTopNRequestAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAggregations

`func (o *TelemetryDruidTopNRequestAllOf) GetAggregations() []TelemetryDruidAggregator`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TelemetryDruidTopNRequestAllOf) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TelemetryDruidTopNRequestAllOf) SetAggregations(v []TelemetryDruidAggregator)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TelemetryDruidTopNRequestAllOf) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetPostAggregations

`func (o *TelemetryDruidTopNRequestAllOf) GetPostAggregations() []TelemetryDruidPostAggregator`

GetPostAggregations returns the PostAggregations field if non-nil, zero value otherwise.

### GetPostAggregationsOk

`func (o *TelemetryDruidTopNRequestAllOf) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetPostAggregationsOk returns a tuple with the PostAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAggregations

`func (o *TelemetryDruidTopNRequestAllOf) SetPostAggregations(v []TelemetryDruidPostAggregator)`

SetPostAggregations sets PostAggregations field to given value.

### HasPostAggregations

`func (o *TelemetryDruidTopNRequestAllOf) HasPostAggregations() bool`

HasPostAggregations returns a boolean if a field has been set.

### GetDimension

`func (o *TelemetryDruidTopNRequestAllOf) GetDimension() TelemetryDruidDimensionSpec`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidTopNRequestAllOf) GetDimensionOk() (*TelemetryDruidDimensionSpec, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidTopNRequestAllOf) SetDimension(v TelemetryDruidDimensionSpec)`

SetDimension sets Dimension field to given value.


### GetThreshold

`func (o *TelemetryDruidTopNRequestAllOf) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TelemetryDruidTopNRequestAllOf) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TelemetryDruidTopNRequestAllOf) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetMetric

`func (o *TelemetryDruidTopNRequestAllOf) GetMetric() TelemetryDruidTopNMetricSpec`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TelemetryDruidTopNRequestAllOf) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TelemetryDruidTopNRequestAllOf) SetMetric(v TelemetryDruidTopNMetricSpec)`

SetMetric sets Metric field to given value.


### GetContext

`func (o *TelemetryDruidTopNRequestAllOf) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidTopNRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidTopNRequestAllOf) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidTopNRequestAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


