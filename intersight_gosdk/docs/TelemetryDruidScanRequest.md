# TelemetryDruidScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | null | 
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Intervals** | **[]string** | A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over. | 
**ResultFormat** | Pointer to **string** | How the results are represented, list, compactedList or valueVector. Currently only list is supported. | [optional] [default to "list"]
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Columns** | Pointer to **[]string** | A String array of dimensions and metrics to scan. If left empty, all dimensions and metrics are returned. | [optional] 
**BatchSize** | Pointer to **int32** | The maximum number of rows buffered before being returned to the client. | [optional] [default to 20480]
**Limit** | Pointer to **int32** | How many rows to return. If not specified, all rows will be returned. | [optional] 
**Order** | Pointer to **string** | The ordering of returned rows based on timestamp. \&quot;ascending\&quot;, \&quot;descending\&quot;, and \&quot;none\&quot; (default) are supported. Currently, \&quot;ascending\&quot; and \&quot;descending\&quot; are only supported for queries where the __time column is included in the columns field and the requirements outlined in the time ordering section are met. | [optional] [default to "none"]
**Legacy** | Pointer to **bool** | Return results consistent with the legacy \&quot;scan-query\&quot; contrib extension. Defaults to the value set by druid.query.scan.legacy, which in turn defaults to false. | [optional] [default to false]
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidScanRequest

`func NewTelemetryDruidScanRequest(queryType string, dataSource TelemetryDruidDataSource, intervals []string, ) *TelemetryDruidScanRequest`

NewTelemetryDruidScanRequest instantiates a new TelemetryDruidScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidScanRequestWithDefaults

`func NewTelemetryDruidScanRequestWithDefaults() *TelemetryDruidScanRequest`

NewTelemetryDruidScanRequestWithDefaults instantiates a new TelemetryDruidScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *TelemetryDruidScanRequest) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *TelemetryDruidScanRequest) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *TelemetryDruidScanRequest) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetDataSource

`func (o *TelemetryDruidScanRequest) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidScanRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidScanRequest) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetIntervals

`func (o *TelemetryDruidScanRequest) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidScanRequest) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidScanRequest) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetResultFormat

`func (o *TelemetryDruidScanRequest) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *TelemetryDruidScanRequest) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *TelemetryDruidScanRequest) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.

### HasResultFormat

`func (o *TelemetryDruidScanRequest) HasResultFormat() bool`

HasResultFormat returns a boolean if a field has been set.

### GetFilter

`func (o *TelemetryDruidScanRequest) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidScanRequest) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidScanRequest) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidScanRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetColumns

`func (o *TelemetryDruidScanRequest) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidScanRequest) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidScanRequest) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryDruidScanRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetBatchSize

`func (o *TelemetryDruidScanRequest) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *TelemetryDruidScanRequest) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *TelemetryDruidScanRequest) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *TelemetryDruidScanRequest) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetLimit

`func (o *TelemetryDruidScanRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidScanRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidScanRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TelemetryDruidScanRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrder

`func (o *TelemetryDruidScanRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TelemetryDruidScanRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TelemetryDruidScanRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TelemetryDruidScanRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetLegacy

`func (o *TelemetryDruidScanRequest) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *TelemetryDruidScanRequest) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *TelemetryDruidScanRequest) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *TelemetryDruidScanRequest) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidScanRequest) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidScanRequest) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidScanRequest) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidScanRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


