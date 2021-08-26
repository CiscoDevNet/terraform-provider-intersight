# TelemetryDruidTimeSeriesRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Descending** | Pointer to **bool** | Whether to make descending ordered result. Default is false(ascending). | [optional] 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. | 
**Granularity** | [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | 
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Aggregations** | Pointer to [**[]TelemetryDruidAggregator**](TelemetryDruidAggregator.md) | Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket. | [optional] 
**PostAggregations** | Pointer to [**[]TelemetryDruidPostAggregator**](TelemetryDruidPostAggregator.md) | Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires. | [optional] 
**Limit** | Pointer to **int32** | An integer that limits the number of results. The default is unlimited. | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidTimeSeriesRequestAllOf

`func NewTelemetryDruidTimeSeriesRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, ) *TelemetryDruidTimeSeriesRequestAllOf`

NewTelemetryDruidTimeSeriesRequestAllOf instantiates a new TelemetryDruidTimeSeriesRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTimeSeriesRequestAllOfWithDefaults

`func NewTelemetryDruidTimeSeriesRequestAllOfWithDefaults() *TelemetryDruidTimeSeriesRequestAllOf`

NewTelemetryDruidTimeSeriesRequestAllOfWithDefaults instantiates a new TelemetryDruidTimeSeriesRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetDescending

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetDescending() bool`

GetDescending returns the Descending field if non-nil, zero value otherwise.

### GetDescendingOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetDescendingOk() (*bool, bool)`

GetDescendingOk returns a tuple with the Descending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescending

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetDescending(v bool)`

SetDescending sets Descending field to given value.

### HasDescending

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasDescending() bool`

HasDescending returns a boolean if a field has been set.

### GetIntervals

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetGranularity

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetAggregations() []TelemetryDruidAggregator`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetAggregations(v []TelemetryDruidAggregator)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetPostAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetPostAggregations() []TelemetryDruidPostAggregator`

GetPostAggregations returns the PostAggregations field if non-nil, zero value otherwise.

### GetPostAggregationsOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetPostAggregationsOk returns a tuple with the PostAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetPostAggregations(v []TelemetryDruidPostAggregator)`

SetPostAggregations sets PostAggregations field to given value.

### HasPostAggregations

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasPostAggregations() bool`

HasPostAggregations returns a boolean if a field has been set.

### GetLimit

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidTimeSeriesRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidTimeSeriesRequestAllOf) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidTimeSeriesRequestAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


