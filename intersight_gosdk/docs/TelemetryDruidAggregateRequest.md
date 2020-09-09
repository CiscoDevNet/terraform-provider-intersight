# TelemetryDruidAggregateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | null | 
**DataSource** | [**TelemetryDruidDataSource**](telemetry.DruidDataSource.md) |  | 
**Descending** | Pointer to **bool** | Whether to make descending ordered result. Default is false(ascending). | [optional] 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. If an interval is not specified, the query will use a default interval that spans a configurable period before the end time of the most recent segment. | 
**Granularity** | [**TelemetryDruidGranularity**](telemetry.DruidGranularity.md) |  | 
**Filter** | Pointer to [**TelemetryDruidFilter**](telemetry.DruidFilter.md) |  | [optional] 
**Aggregations** | Pointer to [**[]TelemetryDruidAggregator**](telemetry.DruidAggregator.md) | Aggregation functions are used to summarize data in buckets. Summarization functions include counting rows, calculating the min/max/sum of metrics and retrieving the first/last value of metrics for each bucket. Additional summarization functions are available with extensions. If no aggregator is provided, the results will be empty for each bucket. | [optional] 
**PostAggregations** | Pointer to [**[]TelemetryDruidPostAggregator**](telemetry.DruidPostAggregator.md) | Post-aggregations are specifications of processing that should happen on aggregated values as they come out of Apache Druid. If you include a post aggregation as part of a query, make sure to include all aggregators the post-aggregator requires. | [optional] 
**Limit** | Pointer to **int32** | How many rows to return. If not specified, all rows will be returned. | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](telemetry.DruidQueryContext.md) |  | [optional] 
**Dimension** | [**TelemetryDruidDimensionSpec**](telemetry.DruidDimensionSpec.md) |  | 
**Threshold** | **int32** | An integer defining the N in the topN (i.e. how many results you want in the top list). | 
**Metric** | [**TelemetryDruidTopNMetricSpec**](telemetry.DruidTopNMetricSpec.md) |  | 
**Dimensions** | [**[]TelemetryDruidDimensionSpec**](telemetry.DruidDimensionSpec.md) | A JSON list of dimensions to do the groupBy over; or see DimensionSpec for ways to extract dimensions.. | 
**LimitSpec** | Pointer to [**TelemetryDruidDefaultLimitSpec**](telemetry.DruidDefaultLimitSpec.md) |  | [optional] 
**Having** | Pointer to [**TelemetryDruidHavingFilter**](telemetry.DruidHavingFilter.md) |  | [optional] 
**SubtotalsSpec** | Pointer to **map[string]interface{}** | A JSON array of arrays to return additional result sets for groupings of subsets of top level dimensions. The subtotals feature allows computation of multiple sub-groupings in a single query. To use this feature, add a \&quot;subtotalsSpec\&quot; to your query, which should be a list of subgroup dimension sets. It should contain the \&quot;outputName\&quot; from dimensions in your \&quot;dimensions\&quot; attribute, in the same order as they appear in the \&quot;dimensions\&quot; attribute. | [optional] 
**ResultFormat** | Pointer to **string** | How the results are represented, list, compactedList or valueVector. Currently only list and compactedList are supported. | [optional] [default to "list"]
**Columns** | Pointer to **[]string** | A String array of dimensions and metrics to scan. If left empty, all dimensions and metrics are returned. | [optional] 
**BatchSize** | Pointer to **int32** | The maximum number of rows buffered before being returned to the client. | [optional] [default to 20480]
**Order** | Pointer to **string** | The ordering of returned rows based on timestamp. \&quot;ascending\&quot;, \&quot;descending\&quot;, and \&quot;none\&quot; (default) are supported. Currently, \&quot;ascending\&quot; and \&quot;descending\&quot; are only supported for queries where the __time column is included in the columns field and the requirements outlined in the time ordering section are met. | [optional] [default to "none"]
**Legacy** | Pointer to **bool** | Return results consistent with the legacy \&quot;scan-query\&quot; contrib extension. Defaults to the value set by druid.query.scan.legacy, which in turn defaults to false. | [optional] [default to false]
**Bound** | Pointer to **string** | Optional, set to maxTime or minTime to return only the latest or earliest timestamp. Default to returning both if not set. | [optional] 
**ToInclude** | Pointer to **map[string]interface{}** | A JSON Object representing what columns should be included in the result. Defaults to \&quot;all\&quot;. | [optional] 
**Merge** | Pointer to **bool** | Merge all individual segment metadata results into a single result. | [optional] 
**AnalysisTypes** | Pointer to **[]string** | A list of Strings specifying what column properties (e.g. cardinality, size) should be calculated and returned in the result. Defaults to [\&quot;cardinality\&quot;, \&quot;interval\&quot;, \&quot;minmax\&quot;], but can be overridden with using the segment metadata query config. * cardinality - in the result will return the estimated floor of cardinality for each column. Only relevant for dimension columns. * minmax - Estimated min/max values for each column. Only relevant for dimension columns. * size - in the result will contain the estimated total segment byte size as if the data were stored in text format. * intervals - in the result will contain the list of intervals associated with the queried segments. * timestampSpec - in the result will contain timestampSpec of data stored in segments. This can be null if timestampSpec of segments was unknown or unmergeable (if merging is enabled). * queryGranularity - in the result will contain query granularity of data stored in segments. This can be null if query granularity of segments was unknown or unmergeable (if merging is enabled). * aggregators - in the result will contain the list of aggregators usable for querying metric columns. This may be null if the aggregators are unknown or unmergeable (if merging is enabled). Merging can be strict or lenient. The form of the result is a map of column name to aggregator. * rollup - in the result is true/false/null. When merging is enabled, if some are rollup, others are not, result is null. | [optional] 
**LenientAggregatorMerge** | Pointer to **bool** | If true, and if the \&quot;aggregators\&quot; analysisType is enabled, aggregators will be merged leniently. | [optional] 

## Methods

### NewTelemetryDruidAggregateRequest

`func NewTelemetryDruidAggregateRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, dimension TelemetryDruidDimensionSpec, threshold int32, metric TelemetryDruidTopNMetricSpec, dimensions []TelemetryDruidDimensionSpec, ) *TelemetryDruidAggregateRequest`

NewTelemetryDruidAggregateRequest instantiates a new TelemetryDruidAggregateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidAggregateRequestWithDefaults

`func NewTelemetryDruidAggregateRequestWithDefaults() *TelemetryDruidAggregateRequest`

NewTelemetryDruidAggregateRequestWithDefaults instantiates a new TelemetryDruidAggregateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *TelemetryDruidAggregateRequest) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *TelemetryDruidAggregateRequest) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *TelemetryDruidAggregateRequest) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetDataSource

`func (o *TelemetryDruidAggregateRequest) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidAggregateRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidAggregateRequest) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetDescending

`func (o *TelemetryDruidAggregateRequest) GetDescending() bool`

GetDescending returns the Descending field if non-nil, zero value otherwise.

### GetDescendingOk

`func (o *TelemetryDruidAggregateRequest) GetDescendingOk() (*bool, bool)`

GetDescendingOk returns a tuple with the Descending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescending

`func (o *TelemetryDruidAggregateRequest) SetDescending(v bool)`

SetDescending sets Descending field to given value.

### HasDescending

`func (o *TelemetryDruidAggregateRequest) HasDescending() bool`

HasDescending returns a boolean if a field has been set.

### GetIntervals

`func (o *TelemetryDruidAggregateRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidAggregateRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidAggregateRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetGranularity

`func (o *TelemetryDruidAggregateRequest) GetGranularity() TelemetryDruidGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *TelemetryDruidAggregateRequest) GetGranularityOk() (*TelemetryDruidGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *TelemetryDruidAggregateRequest) SetGranularity(v TelemetryDruidGranularity)`

SetGranularity sets Granularity field to given value.


### GetFilter

`func (o *TelemetryDruidAggregateRequest) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidAggregateRequest) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidAggregateRequest) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidAggregateRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAggregations

`func (o *TelemetryDruidAggregateRequest) GetAggregations() []TelemetryDruidAggregator`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TelemetryDruidAggregateRequest) GetAggregationsOk() (*[]TelemetryDruidAggregator, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TelemetryDruidAggregateRequest) SetAggregations(v []TelemetryDruidAggregator)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TelemetryDruidAggregateRequest) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetPostAggregations

`func (o *TelemetryDruidAggregateRequest) GetPostAggregations() []TelemetryDruidPostAggregator`

GetPostAggregations returns the PostAggregations field if non-nil, zero value otherwise.

### GetPostAggregationsOk

`func (o *TelemetryDruidAggregateRequest) GetPostAggregationsOk() (*[]TelemetryDruidPostAggregator, bool)`

GetPostAggregationsOk returns a tuple with the PostAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAggregations

`func (o *TelemetryDruidAggregateRequest) SetPostAggregations(v []TelemetryDruidPostAggregator)`

SetPostAggregations sets PostAggregations field to given value.

### HasPostAggregations

`func (o *TelemetryDruidAggregateRequest) HasPostAggregations() bool`

HasPostAggregations returns a boolean if a field has been set.

### GetLimit

`func (o *TelemetryDruidAggregateRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidAggregateRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidAggregateRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TelemetryDruidAggregateRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidAggregateRequest) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidAggregateRequest) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidAggregateRequest) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidAggregateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDimension

`func (o *TelemetryDruidAggregateRequest) GetDimension() TelemetryDruidDimensionSpec`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TelemetryDruidAggregateRequest) GetDimensionOk() (*TelemetryDruidDimensionSpec, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TelemetryDruidAggregateRequest) SetDimension(v TelemetryDruidDimensionSpec)`

SetDimension sets Dimension field to given value.


### GetThreshold

`func (o *TelemetryDruidAggregateRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TelemetryDruidAggregateRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TelemetryDruidAggregateRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetMetric

`func (o *TelemetryDruidAggregateRequest) GetMetric() TelemetryDruidTopNMetricSpec`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TelemetryDruidAggregateRequest) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TelemetryDruidAggregateRequest) SetMetric(v TelemetryDruidTopNMetricSpec)`

SetMetric sets Metric field to given value.


### GetDimensions

`func (o *TelemetryDruidAggregateRequest) GetDimensions() []TelemetryDruidDimensionSpec`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *TelemetryDruidAggregateRequest) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *TelemetryDruidAggregateRequest) SetDimensions(v []TelemetryDruidDimensionSpec)`

SetDimensions sets Dimensions field to given value.


### GetLimitSpec

`func (o *TelemetryDruidAggregateRequest) GetLimitSpec() TelemetryDruidDefaultLimitSpec`

GetLimitSpec returns the LimitSpec field if non-nil, zero value otherwise.

### GetLimitSpecOk

`func (o *TelemetryDruidAggregateRequest) GetLimitSpecOk() (*TelemetryDruidDefaultLimitSpec, bool)`

GetLimitSpecOk returns a tuple with the LimitSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSpec

`func (o *TelemetryDruidAggregateRequest) SetLimitSpec(v TelemetryDruidDefaultLimitSpec)`

SetLimitSpec sets LimitSpec field to given value.

### HasLimitSpec

`func (o *TelemetryDruidAggregateRequest) HasLimitSpec() bool`

HasLimitSpec returns a boolean if a field has been set.

### GetHaving

`func (o *TelemetryDruidAggregateRequest) GetHaving() TelemetryDruidHavingFilter`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *TelemetryDruidAggregateRequest) GetHavingOk() (*TelemetryDruidHavingFilter, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *TelemetryDruidAggregateRequest) SetHaving(v TelemetryDruidHavingFilter)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *TelemetryDruidAggregateRequest) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetSubtotalsSpec

`func (o *TelemetryDruidAggregateRequest) GetSubtotalsSpec() map[string]interface{}`

GetSubtotalsSpec returns the SubtotalsSpec field if non-nil, zero value otherwise.

### GetSubtotalsSpecOk

`func (o *TelemetryDruidAggregateRequest) GetSubtotalsSpecOk() (*map[string]interface{}, bool)`

GetSubtotalsSpecOk returns a tuple with the SubtotalsSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalsSpec

`func (o *TelemetryDruidAggregateRequest) SetSubtotalsSpec(v map[string]interface{})`

SetSubtotalsSpec sets SubtotalsSpec field to given value.

### HasSubtotalsSpec

`func (o *TelemetryDruidAggregateRequest) HasSubtotalsSpec() bool`

HasSubtotalsSpec returns a boolean if a field has been set.

### GetResultFormat

`func (o *TelemetryDruidAggregateRequest) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *TelemetryDruidAggregateRequest) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *TelemetryDruidAggregateRequest) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.

### HasResultFormat

`func (o *TelemetryDruidAggregateRequest) HasResultFormat() bool`

HasResultFormat returns a boolean if a field has been set.

### GetColumns

`func (o *TelemetryDruidAggregateRequest) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidAggregateRequest) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidAggregateRequest) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryDruidAggregateRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetBatchSize

`func (o *TelemetryDruidAggregateRequest) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *TelemetryDruidAggregateRequest) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *TelemetryDruidAggregateRequest) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *TelemetryDruidAggregateRequest) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetOrder

`func (o *TelemetryDruidAggregateRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TelemetryDruidAggregateRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TelemetryDruidAggregateRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TelemetryDruidAggregateRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetLegacy

`func (o *TelemetryDruidAggregateRequest) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *TelemetryDruidAggregateRequest) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *TelemetryDruidAggregateRequest) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *TelemetryDruidAggregateRequest) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetBound

`func (o *TelemetryDruidAggregateRequest) GetBound() string`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *TelemetryDruidAggregateRequest) GetBoundOk() (*string, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *TelemetryDruidAggregateRequest) SetBound(v string)`

SetBound sets Bound field to given value.

### HasBound

`func (o *TelemetryDruidAggregateRequest) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetToInclude

`func (o *TelemetryDruidAggregateRequest) GetToInclude() map[string]interface{}`

GetToInclude returns the ToInclude field if non-nil, zero value otherwise.

### GetToIncludeOk

`func (o *TelemetryDruidAggregateRequest) GetToIncludeOk() (*map[string]interface{}, bool)`

GetToIncludeOk returns a tuple with the ToInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInclude

`func (o *TelemetryDruidAggregateRequest) SetToInclude(v map[string]interface{})`

SetToInclude sets ToInclude field to given value.

### HasToInclude

`func (o *TelemetryDruidAggregateRequest) HasToInclude() bool`

HasToInclude returns a boolean if a field has been set.

### GetMerge

`func (o *TelemetryDruidAggregateRequest) GetMerge() bool`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *TelemetryDruidAggregateRequest) GetMergeOk() (*bool, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *TelemetryDruidAggregateRequest) SetMerge(v bool)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *TelemetryDruidAggregateRequest) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### GetAnalysisTypes

`func (o *TelemetryDruidAggregateRequest) GetAnalysisTypes() []string`

GetAnalysisTypes returns the AnalysisTypes field if non-nil, zero value otherwise.

### GetAnalysisTypesOk

`func (o *TelemetryDruidAggregateRequest) GetAnalysisTypesOk() (*[]string, bool)`

GetAnalysisTypesOk returns a tuple with the AnalysisTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisTypes

`func (o *TelemetryDruidAggregateRequest) SetAnalysisTypes(v []string)`

SetAnalysisTypes sets AnalysisTypes field to given value.

### HasAnalysisTypes

`func (o *TelemetryDruidAggregateRequest) HasAnalysisTypes() bool`

HasAnalysisTypes returns a boolean if a field has been set.

### GetLenientAggregatorMerge

`func (o *TelemetryDruidAggregateRequest) GetLenientAggregatorMerge() bool`

GetLenientAggregatorMerge returns the LenientAggregatorMerge field if non-nil, zero value otherwise.

### GetLenientAggregatorMergeOk

`func (o *TelemetryDruidAggregateRequest) GetLenientAggregatorMergeOk() (*bool, bool)`

GetLenientAggregatorMergeOk returns a tuple with the LenientAggregatorMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLenientAggregatorMerge

`func (o *TelemetryDruidAggregateRequest) SetLenientAggregatorMerge(v bool)`

SetLenientAggregatorMerge sets LenientAggregatorMerge field to given value.

### HasLenientAggregatorMerge

`func (o *TelemetryDruidAggregateRequest) HasLenientAggregatorMerge() bool`

HasLenientAggregatorMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


