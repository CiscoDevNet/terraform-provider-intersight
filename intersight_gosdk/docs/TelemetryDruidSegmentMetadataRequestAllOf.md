# TelemetryDruidSegmentMetadataRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Intervals** | Pointer to **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. If an interval is not specified, the query will use a default interval that spans a configurable period before the end time of the most recent segment. | [optional] 
**ToInclude** | Pointer to **map[string]interface{}** | A JSON Object representing what columns should be included in the result. Defaults to \&quot;all\&quot;. | [optional] 
**Merge** | Pointer to **bool** | Merge all individual segment metadata results into a single result. | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 
**AnalysisTypes** | Pointer to **[]string** | A list of Strings specifying what column properties (e.g. cardinality, size) should be calculated and returned in the result. Defaults to [\&quot;cardinality\&quot;, \&quot;interval\&quot;, \&quot;minmax\&quot;], but can be overridden with using the segment metadata query config. * cardinality - in the result will return the estimated floor of cardinality for each column. Only relevant for dimension columns. * minmax - Estimated min/max values for each column. Only relevant for dimension columns. * size - in the result will contain the estimated total segment byte size as if the data were stored in text format. * intervals - in the result will contain the list of intervals associated with the queried segments. * timestampSpec - in the result will contain timestampSpec of data stored in segments. This can be null if timestampSpec of segments was unknown or unmergeable (if merging is enabled). * queryGranularity - in the result will contain query granularity of data stored in segments. This can be null if query granularity of segments was unknown or unmergeable (if merging is enabled). * aggregators - in the result will contain the list of aggregators usable for querying metric columns. This may be null if the aggregators are unknown or unmergeable (if merging is enabled). Merging can be strict or lenient. The form of the result is a map of column name to aggregator. * rollup - in the result is true/false/null. When merging is enabled, if some are rollup, others are not, result is null. | [optional] 
**LenientAggregatorMerge** | Pointer to **bool** | If true, and if the \&quot;aggregators\&quot; analysisType is enabled, aggregators will be merged leniently. | [optional] 

## Methods

### NewTelemetryDruidSegmentMetadataRequestAllOf

`func NewTelemetryDruidSegmentMetadataRequestAllOf(dataSource TelemetryDruidDataSource, ) *TelemetryDruidSegmentMetadataRequestAllOf`

NewTelemetryDruidSegmentMetadataRequestAllOf instantiates a new TelemetryDruidSegmentMetadataRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidSegmentMetadataRequestAllOfWithDefaults

`func NewTelemetryDruidSegmentMetadataRequestAllOfWithDefaults() *TelemetryDruidSegmentMetadataRequestAllOf`

NewTelemetryDruidSegmentMetadataRequestAllOfWithDefaults instantiates a new TelemetryDruidSegmentMetadataRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetIntervals

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetToInclude

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetToInclude() map[string]interface{}`

GetToInclude returns the ToInclude field if non-nil, zero value otherwise.

### GetToIncludeOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetToIncludeOk() (*map[string]interface{}, bool)`

GetToIncludeOk returns a tuple with the ToInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToInclude

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetToInclude(v map[string]interface{})`

SetToInclude sets ToInclude field to given value.

### HasToInclude

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasToInclude() bool`

HasToInclude returns a boolean if a field has been set.

### GetMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetMerge() bool`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetMergeOk() (*bool, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetMerge(v bool)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAnalysisTypes

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetAnalysisTypes() []string`

GetAnalysisTypes returns the AnalysisTypes field if non-nil, zero value otherwise.

### GetAnalysisTypesOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetAnalysisTypesOk() (*[]string, bool)`

GetAnalysisTypesOk returns a tuple with the AnalysisTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisTypes

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetAnalysisTypes(v []string)`

SetAnalysisTypes sets AnalysisTypes field to given value.

### HasAnalysisTypes

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasAnalysisTypes() bool`

HasAnalysisTypes returns a boolean if a field has been set.

### GetLenientAggregatorMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetLenientAggregatorMerge() bool`

GetLenientAggregatorMerge returns the LenientAggregatorMerge field if non-nil, zero value otherwise.

### GetLenientAggregatorMergeOk

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) GetLenientAggregatorMergeOk() (*bool, bool)`

GetLenientAggregatorMergeOk returns a tuple with the LenientAggregatorMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLenientAggregatorMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) SetLenientAggregatorMerge(v bool)`

SetLenientAggregatorMerge sets LenientAggregatorMerge field to given value.

### HasLenientAggregatorMerge

`func (o *TelemetryDruidSegmentMetadataRequestAllOf) HasLenientAggregatorMerge() bool`

HasLenientAggregatorMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


