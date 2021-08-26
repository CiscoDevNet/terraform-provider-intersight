# TelemetryDruidScanRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTelemetryDruidScanRequestAllOf

`func NewTelemetryDruidScanRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, ) *TelemetryDruidScanRequestAllOf`

NewTelemetryDruidScanRequestAllOf instantiates a new TelemetryDruidScanRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidScanRequestAllOfWithDefaults

`func NewTelemetryDruidScanRequestAllOfWithDefaults() *TelemetryDruidScanRequestAllOf`

NewTelemetryDruidScanRequestAllOfWithDefaults instantiates a new TelemetryDruidScanRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *TelemetryDruidScanRequestAllOf) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidScanRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidScanRequestAllOf) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetIntervals

`func (o *TelemetryDruidScanRequestAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *TelemetryDruidScanRequestAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *TelemetryDruidScanRequestAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.


### GetResultFormat

`func (o *TelemetryDruidScanRequestAllOf) GetResultFormat() string`

GetResultFormat returns the ResultFormat field if non-nil, zero value otherwise.

### GetResultFormatOk

`func (o *TelemetryDruidScanRequestAllOf) GetResultFormatOk() (*string, bool)`

GetResultFormatOk returns a tuple with the ResultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFormat

`func (o *TelemetryDruidScanRequestAllOf) SetResultFormat(v string)`

SetResultFormat sets ResultFormat field to given value.

### HasResultFormat

`func (o *TelemetryDruidScanRequestAllOf) HasResultFormat() bool`

HasResultFormat returns a boolean if a field has been set.

### GetFilter

`func (o *TelemetryDruidScanRequestAllOf) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidScanRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidScanRequestAllOf) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidScanRequestAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetColumns

`func (o *TelemetryDruidScanRequestAllOf) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryDruidScanRequestAllOf) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryDruidScanRequestAllOf) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryDruidScanRequestAllOf) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetBatchSize

`func (o *TelemetryDruidScanRequestAllOf) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *TelemetryDruidScanRequestAllOf) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *TelemetryDruidScanRequestAllOf) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *TelemetryDruidScanRequestAllOf) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetLimit

`func (o *TelemetryDruidScanRequestAllOf) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TelemetryDruidScanRequestAllOf) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TelemetryDruidScanRequestAllOf) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TelemetryDruidScanRequestAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrder

`func (o *TelemetryDruidScanRequestAllOf) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TelemetryDruidScanRequestAllOf) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TelemetryDruidScanRequestAllOf) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TelemetryDruidScanRequestAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetLegacy

`func (o *TelemetryDruidScanRequestAllOf) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *TelemetryDruidScanRequestAllOf) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *TelemetryDruidScanRequestAllOf) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *TelemetryDruidScanRequestAllOf) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidScanRequestAllOf) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidScanRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidScanRequestAllOf) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidScanRequestAllOf) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


