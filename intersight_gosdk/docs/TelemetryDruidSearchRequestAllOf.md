# TelemetryDruidSearchRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. | 
**Granularity** | [**TelemetryDruidGranularity**](TelemetryDruidGranularity.md) |  | 
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Aggregations** | Pointer to [**[]TelemetryDruidAggregator**](TelemetryDruidAggregator.md) | Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket. | [optional] 
**SearchDimensions** | Pointer to **[]string** | The list of dimensions to run the search over. Excluding this means the search is run over all dimensions. | [optional] 
**Query** | Pointer to [**TelemetryDruidAggregateSearchSpec**](TelemetryDruidAggregateSearchSpec.md) |  | [optional] 
**Limit** | Pointer to **int32** | An integer that limits the number of results. The default is unlimited. | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidSearchRequestAllOf

`func NewTelemetryDruidSearchRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, ) *TelemetryDruidSearchRequestAllOf`

NewTelemetryDruidSearchRequestAllOf instantiates a new TelemetryDruidSearchRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSearchRequestAllOfWithDefaults

`func NewTelemetryDruidSearchRequestAllOfWithDefaults() *TelemetryDruidSearchRequestAllOf`

NewTelemetryDruidSearchRequestAllOfWithDefaults instantiates a new TelemetryDruidSearchRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *TelemetryDruidSearchRequestAllOf) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidSearchRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidSearchRequestAllOf) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetIntervals

`func (o *TelemetryDruidSearchRequestAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidSearchRequestAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidSearchRequestAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetGranularity

`func (o *TelemetryDruidSearchRequestAllOf) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidSearchRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidSearchRequestAllOf) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *TelemetryDruidSearchRequestAllOf) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidSearchRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidSearchRequestAllOf) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidSearchRequestAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAggregations

`func (o *TelemetryDruidSearchRequestAllOf) GetAggregations() []TelemetryDruidAggregator`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TelemetryDruidSearchRequestAllOf) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TelemetryDruidSearchRequestAllOf) SetAggregations(v []TelemetryDruidAggregator)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TelemetryDruidSearchRequestAllOf) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetSearchDimensions

`func (o *TelemetryDruidSearchRequestAllOf) GetSearchDimensions() []string`

GetSearchDimensions returns the SearchDimensions field if non-nil, zero value otherwise.

### GetSearchDimensionsOk

`func (o *TelemetryDruidSearchRequestAllOf) GetSearchDimensionsOk() (*[]string, bool)`

GetSearchDimensionsOk returns a tuple with the SearchDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDimensions

`func (o *TelemetryDruidSearchRequestAllOf) SetSearchDimensions(v []string)`

SetSearchDimensions sets SearchDimensions field to given value.

### HasSearchDimensions

`func (o *TelemetryDruidSearchRequestAllOf) HasSearchDimensions() bool`

HasSearchDimensions returns a boolean if a field has been set.

### GetQuery

`func (o *TelemetryDruidSearchRequestAllOf) GetQuery() TelemetryDruidAggregateSearchSpec`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidSearchRequestAllOf) GetQueryOk() (*TelemetryDruidAggregateSearchSpec, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidSearchRequestAllOf) SetQuery(v TelemetryDruidAggregateSearchSpec)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TelemetryDruidSearchRequestAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetLimit

`func (o *TelemetryDruidSearchRequestAllOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidSearchRequestAllOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidSearchRequestAllOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TelemetryDruidSearchRequestAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidSearchRequestAllOf) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidSearchRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidSearchRequestAllOf) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidSearchRequestAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


