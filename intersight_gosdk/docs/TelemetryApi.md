# \TelemetryApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryTelemetryTimeSeries**](TelemetryApi.md#QueryTelemetryTimeSeries) | **Post** /api/v1/telemetry/TimeSeries | Perform a Druid time series aggregation request.



## QueryTelemetryTimeSeries

> []TelemetryDruidIntervalResult QueryTelemetryTimeSeries(ctx).TelemetryDruidAggregateRequest(telemetryDruidAggregateRequest).Execute()

Perform a Druid time series aggregation request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidAggregateRequest := openapiclient.telemetry.DruidAggregateRequest{QueryType: "QueryType_example", DataSource: openapiclient.telemetry.DruidDataSource{Type: "Type_example", Name: "Name_example", DataSources: []string{"DataSources_example"), Query: openapiclient.telemetry.DruidGroupByRequest{QueryType: "QueryType_example", DataSource: openapiclient.telemetry.DruidDataSource{Type: "Type_example", Name: "Name_example", DataSources: []string{"DataSources_example"), Query: openapiclient.telemetry.DruidGroupByRequest{QueryType: "QueryType_example", DataSource: , Dimensions: []TelemetryDruidDimensionSpec{openapiclient.telemetry.DruidDimensionSpec{Type: "Type_example", Dimension: "Dimension_example", OutputName: "OutputName_example", OutputType: "OutputType_example", ExtractionFn: 123}), LimitSpec: openapiclient.telemetry.DruidDefaultLimitSpec{Type: "Type_example", Limit: 123, Columns: []TelemetryDruidOrderByColumnSpec{openapiclient.telemetry.DruidOrderByColumnSpec{Dimension: "Dimension_example", Direction: "Direction_example", DimensionOrder: "DimensionOrder_example"})}, Having: openapiclient.telemetry.DruidHavingFilter{Type: "Type_example", Filter: openapiclient.telemetry.DruidFilter{Type: "Type_example", ExtractionFn: 123, Dimension: "Dimension_example", Value: "Value_example", Dimensions: []TelemetryDruidDimensionSpec{openapiclient.telemetry.DruidDimensionSpec{Type: "Type_example", Dimension: "Dimension_example", OutputName: "OutputName_example", OutputType: "OutputType_example", ExtractionFn: 123}), Pattern: "Pattern_example", Fields: []TelemetryDruidFilter{openapiclient.telemetry.DruidFilter{Type: "Type_example", ExtractionFn: 123, Dimension: "Dimension_example", Value: "Value_example", Dimensions: []TelemetryDruidDimensionSpec{), Pattern: "Pattern_example", Fields: []TelemetryDruidFilter{), Field: }), Field: }, Aggregation: "Aggregation_example", Value: 123, Dimension: "Dimension_example"}, Granularity: openapiclient.telemetry.DruidGranularity{Type: "Type_example", Duration: int64(123), Origin: "TODO", Period: "Period_example", TimeZone: "TimeZone_example"}, Filter: , Aggregations: []TelemetryDruidAggregator{openapiclient.telemetry.DruidAggregator{Type: "Type_example", Name: "Name_example", FieldName: "FieldName_example", MaxStringBytes: 123, Size: 123, Filter: , Aggregator: openapiclient.telemetry.DruidAggregator{Type: "Type_example", Name: "Name_example", FieldName: "FieldName_example", MaxStringBytes: 123, Size: 123, Filter: , Aggregator: }}), PostAggregations: []TelemetryDruidPostAggregator{openapiclient.telemetry.DruidPostAggregator{Type: "Type_example", Name: "Name_example", Fn: "Fn_example", Fields: []string{"Fields_example"), Ordering: "Ordering_example", FieldName: "FieldName_example", Value: 123, Field: "Field_example", Func: "Func_example", Size: 123}), Intervals: []string{"Intervals_example"), SubtotalsSpec: 123, Context: openapiclient.telemetry.DruidQueryContext{GrandTotal: false, SkipEmptyBuckets: false, Timeout: 123, Priority: 123, QueryId: "QueryId_example", UseCache: false, PopulateCache: false, UseResultLevelCache: false, PopulateResultLevelCache: false, BySegment: false, Finalize: false, ChunkPeriod: "ChunkPeriod_example", MaxScatterGatherBytes: 123, MaxQueuedBytes: 123, SerializeDateTimeAsLong: false, SerializeDateTimeAsLongInner: false, EnableParallelMerge: false, ParallelMergeParallelism: 123, ParallelMergeInitialYieldRows: 123, ParallelMergeSmallBatchRows: 123}}, Lookup: "Lookup_example", Left: "Left_example", Right: "Right_example", RightPrefix: "RightPrefix_example", Condition: "Condition_example", JoinType: "JoinType_example", ColumnNames: []string{"ColumnNames_example"), Rows: [][]string{[]string{"Rows_example"))}, Dimensions: []TelemetryDruidDimensionSpec{), LimitSpec: openapiclient.telemetry.DruidDefaultLimitSpec{Type: "Type_example", Limit: 123, Columns: []TelemetryDruidOrderByColumnSpec{openapiclient.telemetry.DruidOrderByColumnSpec{Dimension: "Dimension_example", Direction: "Direction_example", DimensionOrder: "DimensionOrder_example"})}, Having: openapiclient.telemetry.DruidHavingFilter{Type: "Type_example", Filter: , Aggregation: "Aggregation_example", Value: 123, Dimension: "Dimension_example"}, Granularity: openapiclient.telemetry.DruidGranularity{Type: "Type_example", Duration: int64(123), Origin: "TODO", Period: "Period_example", TimeZone: "TimeZone_example"}, Filter: , Aggregations: []TelemetryDruidAggregator{), PostAggregations: []TelemetryDruidPostAggregator{openapiclient.telemetry.DruidPostAggregator{Type: "Type_example", Name: "Name_example", Fn: "Fn_example", Fields: []string{"Fields_example"), Ordering: "Ordering_example", FieldName: "FieldName_example", Value: 123, Field: "Field_example", Func: "Func_example", Size: 123}), Intervals: []string{"Intervals_example"), SubtotalsSpec: 123, Context: openapiclient.telemetry.DruidQueryContext{GrandTotal: false, SkipEmptyBuckets: false, Timeout: 123, Priority: 123, QueryId: "QueryId_example", UseCache: false, PopulateCache: false, UseResultLevelCache: false, PopulateResultLevelCache: false, BySegment: false, Finalize: false, ChunkPeriod: "ChunkPeriod_example", MaxScatterGatherBytes: 123, MaxQueuedBytes: 123, SerializeDateTimeAsLong: false, SerializeDateTimeAsLongInner: false, EnableParallelMerge: false, ParallelMergeParallelism: 123, ParallelMergeInitialYieldRows: 123, ParallelMergeSmallBatchRows: 123}}, Lookup: "Lookup_example", Left: "Left_example", Right: "Right_example", RightPrefix: "RightPrefix_example", Condition: "Condition_example", JoinType: "JoinType_example", ColumnNames: []string{"ColumnNames_example"), Rows: [][]string{[]string{"Rows_example"))}, Descending: false, Intervals: []string{"Intervals_example"), Granularity: , Filter: , Aggregations: []TelemetryDruidAggregator{), PostAggregations: []TelemetryDruidPostAggregator{), Limit: 123, Context: , Dimension: , Threshold: 123, Metric: openapiclient.telemetry.DruidTopNMetricSpec{Type: "Type_example", Metric: openapiclient.telemetry.DruidTopNMetricSpec{Type: "Type_example", Metric: , Ordering: "Ordering_example", PreviousStop: "PreviousStop_example"}, Ordering: "Ordering_example", PreviousStop: "PreviousStop_example"}, Dimensions: []TelemetryDruidDimensionSpec{), LimitSpec: , Having: , SubtotalsSpec: 123, ResultFormat: "ResultFormat_example", Columns: []string{"Columns_example"), BatchSize: 123, Order: "Order_example", Legacy: false, Bound: "Bound_example", ToInclude: 123, Merge: false, AnalysisTypes: []string{"AnalysisTypes_example"), LenientAggregatorMerge: false} // TelemetryDruidAggregateRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTimeSeries(context.Background(), telemetryDruidAggregateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTimeSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTimeSeries`: []TelemetryDruidIntervalResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTimeSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTimeSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidAggregateRequest** | [**TelemetryDruidAggregateRequest**](TelemetryDruidAggregateRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidIntervalResult**](telemetry.DruidIntervalResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

