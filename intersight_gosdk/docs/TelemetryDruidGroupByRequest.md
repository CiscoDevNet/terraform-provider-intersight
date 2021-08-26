# TelemetryDruidGroupByRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | null | 
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Dimensions** | [**[]TelemetryDruidDimensionSpec**](TelemetryDruidDimensionSpec.md) | A JSON list of dimensions to do the groupBy over; or see DimensionSpec for ways to extract dimensions.. | 
**LimitSpec** | Pointer to [**TelemetryDruidDefaultLimitSpec**](TelemetryDruidDefaultLimitSpec.md) |  | [optional] 
**Having** | Pointer to [**TelemetryDruidHavingFilter**](TelemetryDruidHavingFilter.md) |  | [optional] 
**Granularity** | [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | 
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Aggregations** | Pointer to [**[]TelemetryDruidAggregator**](TelemetryDruidAggregator.md) | Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket. | [optional] 
**PostAggregations** | Pointer to [**[]TelemetryDruidPostAggregator**](TelemetryDruidPostAggregator.md) | Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires. | [optional] 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. | 
**SubtotalsSpec** | Pointer to **map[string]interface{}** | A JSON array of arrays to return additional result sets for groupings of subsets of top level dimensions. The subtotals feature allows computation of multiple sub-groupings in a single query. To use this feature, add a \&quot;subtotalsSpec\&quot; to your query, which should be a list of subgroup dimension sets. It should contain the \&quot;outputName\&quot; from dimensions in your \&quot;dimensions\&quot; attribute, in the same order as they appear in the \&quot;dimensions\&quot; attribute. | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidGroupByRequest

`func NewTelemetryDruidGroupByRequest(queryType string, dataSource TelemetryDruidDataSource, dimensions []TelemetryDruidDimensionSpec, granularity TelemetryDruidGranularity, intervals []string, ) *TelemetryDruidGroupByRequest`

NewTelemetryDruidGroupByRequest instantiates a new TelemetryDruidGroupByRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidGroupByRequestWithDefaults

`func NewTelemetryDruidGroupByRequestWithDefaults() *TelemetryDruidGroupByRequest`

NewTelemetryDruidGroupByRequestWithDefaults instantiates a new TelemetryDruidGroupByRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *TelemetryDruidGroupByRequest) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *TelemetryDruidGroupByRequest) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *TelemetryDruidGroupByRequest) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetDataSource

`func (o *TelemetryDruidGroupByRequest) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidGroupByRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidGroupByRequest) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetDimensions

`func (o *TelemetryDruidGroupByRequest) GetDimensions() []TelemetryDruidDimensionSpec`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TelemetryDruidGroupByRequest) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TelemetryDruidGroupByRequest) SetDimensions(v []TelemetryDruidDimensionSpec)`

SetDimensions sets Dimensions field to given value.


### GetLimitSpec

`func (o *TelemetryDruidGroupByRequest) GetLimitSpec() TelemetryDruidDefaultLimitSpec`

GetLimitSpec returns the LimitSpec field if non-nil, zero value otherwise.

### GetLimitSpecOk

`func (o *TelemetryDruidGroupByRequest) GetLimitSpecOk() (*TelemetryDruidDefaultLimitSpec, bool)`

GetLimitSpecOk returns a tuple with the LimitSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSpec

`func (o *TelemetryDruidGroupByRequest) SetLimitSpec(v TelemetryDruidDefaultLimitSpec)`

SetLimitSpec sets LimitSpec field to given value.

### HasLimitSpec

`func (o *TelemetryDruidGroupByRequest) HasLimitSpec() bool`

HasLimitSpec returns a boolean if a field has been set.

### GetHaving

`func (o *TelemetryDruidGroupByRequest) GetHaving() TelemetryDruidHavingFilter`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *TelemetryDruidGroupByRequest) GetHavingOk() (*TelemetryDruidHavingFilter, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *TelemetryDruidGroupByRequest) SetHaving(v TelemetryDruidHavingFilter)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *TelemetryDruidGroupByRequest) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetGranularity

`func (o *TelemetryDruidGroupByRequest) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidGroupByRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidGroupByRequest) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *TelemetryDruidGroupByRequest) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidGroupByRequest) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidGroupByRequest) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidGroupByRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAggregations

`func (o *TelemetryDruidGroupByRequest) GetAggregations() []TelemetryDruidAggregator`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TelemetryDruidGroupByRequest) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TelemetryDruidGroupByRequest) SetAggregations(v []TelemetryDruidAggregator)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TelemetryDruidGroupByRequest) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetPostAggregations

`func (o *TelemetryDruidGroupByRequest) GetPostAggregations() []TelemetryDruidPostAggregator`

GetPostAggregations returns the PostAggregations field if non-nil, zero value otherwise.

### GetPostAggregationsOk

`func (o *TelemetryDruidGroupByRequest) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetPostAggregationsOk returns a tuple with the PostAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAggregations

`func (o *TelemetryDruidGroupByRequest) SetPostAggregations(v []TelemetryDruidPostAggregator)`

SetPostAggregations sets PostAggregations field to given value.

### HasPostAggregations

`func (o *TelemetryDruidGroupByRequest) HasPostAggregations() bool`

HasPostAggregations returns a boolean if a field has been set.

### GetIntervals

`func (o *TelemetryDruidGroupByRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidGroupByRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidGroupByRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetSubtotalsSpec

`func (o *TelemetryDruidGroupByRequest) GetSubtotalsSpec() map[string]interface{}`

GetSubtotalsSpec returns the SubtotalsSpec field if non-nil, zero value otherwise.

### GetSubtotalsSpecOk

`func (o *TelemetryDruidGroupByRequest) GetSubtotalsSpecOk() (*map[string]interface{}, bool)`

GetSubtotalsSpecOk returns a tuple with the SubtotalsSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalsSpec

`func (o *TelemetryDruidGroupByRequest) SetSubtotalsSpec(v map[string]interface{})`

SetSubtotalsSpec sets SubtotalsSpec field to given value.

### HasSubtotalsSpec

`func (o *TelemetryDruidGroupByRequest) HasSubtotalsSpec() bool`

HasSubtotalsSpec returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidGroupByRequest) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidGroupByRequest) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidGroupByRequest) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidGroupByRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


