# \TelemetryApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryTelemetryDatasourceMetadata**](TelemetryApi.md#QueryTelemetryDatasourceMetadata) | **Post** /api/v1/telemetry/DataSourceMetadata | Perform a Druid DatasourceMetadata request.
[**QueryTelemetryGroupBy**](TelemetryApi.md#QueryTelemetryGroupBy) | **Post** /api/v1/telemetry/GroupBys | Perform a Druid GroupBy request.
[**QueryTelemetryScan**](TelemetryApi.md#QueryTelemetryScan) | **Post** /api/v1/telemetry/Scans | Perform a Druid Scan request.
[**QueryTelemetrySearch**](TelemetryApi.md#QueryTelemetrySearch) | **Post** /api/v1/telemetry/Searches | Perform a Druid Search request.
[**QueryTelemetrySegmentMetadata**](TelemetryApi.md#QueryTelemetrySegmentMetadata) | **Post** /api/v1/telemetry/SegmentMetadata | Perform a Druid SegmentMetadata request.
[**QueryTelemetryTimeBoundary**](TelemetryApi.md#QueryTelemetryTimeBoundary) | **Post** /api/v1/telemetry/TimeBoundaries | Perform a Druid TimeBoundary request.
[**QueryTelemetryTimeSeries**](TelemetryApi.md#QueryTelemetryTimeSeries) | **Post** /api/v1/telemetry/TimeSeries | Perform a Druid TimeSeries request.
[**QueryTelemetryTopN**](TelemetryApi.md#QueryTelemetryTopN) | **Post** /api/v1/telemetry/Topns | Perform a Druid TopN request.



## QueryTelemetryDatasourceMetadata

> []TelemetryDruidDataSourceMetadataResult QueryTelemetryDatasourceMetadata(ctx).TelemetryDruidDataSourceMetadataRequest(telemetryDruidDataSourceMetadataRequest).Execute()

Perform a Druid DatasourceMetadata request.



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
    telemetryDruidDataSourceMetadataRequest := *openapiclient.NewTelemetryDruidDataSourceMetadataRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}) // TelemetryDruidDataSourceMetadataRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryDatasourceMetadata(context.Background()).TelemetryDruidDataSourceMetadataRequest(telemetryDruidDataSourceMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryDatasourceMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryDatasourceMetadata`: []TelemetryDruidDataSourceMetadataResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryDatasourceMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryDatasourceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidDataSourceMetadataRequest** | [**TelemetryDruidDataSourceMetadataRequest**](TelemetryDruidDataSourceMetadataRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidDataSourceMetadataResult**](TelemetryDruidDataSourceMetadataResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryGroupBy

> []TelemetryDruidGroupByResult QueryTelemetryGroupBy(ctx).TelemetryDruidGroupByRequest(telemetryDruidGroupByRequest).Execute()

Perform a Druid GroupBy request.



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
    telemetryDruidGroupByRequest := *openapiclient.NewTelemetryDruidGroupByRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}, []openapiclient.TelemetryDruidDimensionSpec{openapiclient.telemetry.DruidDimensionSpec{TelemetryDruidDefaultDimensionSpec: openapiclient.NewTelemetryDruidDefaultDimensionSpec("Type_example", "Dimension_example", "OutputName_example", "OutputType_example")}}, openapiclient.telemetry.DruidGranularity{TelemetryDruidDurationGranularity: openapiclient.NewTelemetryDruidDurationGranularity("Type_example", int64(123))}, []string{"Intervals_example"}) // TelemetryDruidGroupByRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryGroupBy(context.Background()).TelemetryDruidGroupByRequest(telemetryDruidGroupByRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryGroupBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryGroupBy`: []TelemetryDruidGroupByResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryGroupBy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryGroupByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidGroupByRequest** | [**TelemetryDruidGroupByRequest**](TelemetryDruidGroupByRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidGroupByResult**](TelemetryDruidGroupByResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryScan

> []TelemetryDruidScanResult QueryTelemetryScan(ctx).TelemetryDruidScanRequest(telemetryDruidScanRequest).Execute()

Perform a Druid Scan request.



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
    telemetryDruidScanRequest := *openapiclient.NewTelemetryDruidScanRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}, []string{"Intervals_example"}) // TelemetryDruidScanRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryScan(context.Background()).TelemetryDruidScanRequest(telemetryDruidScanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryScan`: []TelemetryDruidScanResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidScanRequest** | [**TelemetryDruidScanRequest**](TelemetryDruidScanRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidScanResult**](TelemetryDruidScanResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetrySearch

> []TelemetryDruidSearchResult QueryTelemetrySearch(ctx).TelemetryDruidSearchRequest(telemetryDruidSearchRequest).Execute()

Perform a Druid Search request.



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
    telemetryDruidSearchRequest := *openapiclient.NewTelemetryDruidSearchRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}, []string{"Intervals_example"}, openapiclient.telemetry.DruidGranularity{TelemetryDruidDurationGranularity: openapiclient.NewTelemetryDruidDurationGranularity("Type_example", int64(123))}) // TelemetryDruidSearchRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetrySearch(context.Background()).TelemetryDruidSearchRequest(telemetryDruidSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetrySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetrySearch`: []TelemetryDruidSearchResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetrySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetrySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidSearchRequest** | [**TelemetryDruidSearchRequest**](TelemetryDruidSearchRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidSearchResult**](TelemetryDruidSearchResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetrySegmentMetadata

> []TelemetryDruidSegmentMetadataResult QueryTelemetrySegmentMetadata(ctx).TelemetryDruidSegmentMetadataRequest(telemetryDruidSegmentMetadataRequest).Execute()

Perform a Druid SegmentMetadata request.



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
    telemetryDruidSegmentMetadataRequest := *openapiclient.NewTelemetryDruidSegmentMetadataRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}) // TelemetryDruidSegmentMetadataRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetrySegmentMetadata(context.Background()).TelemetryDruidSegmentMetadataRequest(telemetryDruidSegmentMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetrySegmentMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetrySegmentMetadata`: []TelemetryDruidSegmentMetadataResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetrySegmentMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetrySegmentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidSegmentMetadataRequest** | [**TelemetryDruidSegmentMetadataRequest**](TelemetryDruidSegmentMetadataRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidSegmentMetadataResult**](TelemetryDruidSegmentMetadataResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTimeBoundary

> []TelemetryDruidTimeBoundaryResult QueryTelemetryTimeBoundary(ctx).TelemetryDruidTimeBoundaryRequest(telemetryDruidTimeBoundaryRequest).Execute()

Perform a Druid TimeBoundary request.



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
    telemetryDruidTimeBoundaryRequest := *openapiclient.NewTelemetryDruidTimeBoundaryRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}) // TelemetryDruidTimeBoundaryRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTimeBoundary(context.Background()).TelemetryDruidTimeBoundaryRequest(telemetryDruidTimeBoundaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTimeBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTimeBoundary`: []TelemetryDruidTimeBoundaryResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTimeBoundary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTimeBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidTimeBoundaryRequest** | [**TelemetryDruidTimeBoundaryRequest**](TelemetryDruidTimeBoundaryRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidTimeBoundaryResult**](TelemetryDruidTimeBoundaryResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTimeSeries

> []TelemetryDruidIntervalResult QueryTelemetryTimeSeries(ctx).TelemetryDruidTimeSeriesRequest(telemetryDruidTimeSeriesRequest).Execute()

Perform a Druid TimeSeries request.



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
    telemetryDruidTimeSeriesRequest := *openapiclient.NewTelemetryDruidTimeSeriesRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}, []string{"Intervals_example"}, openapiclient.telemetry.DruidGranularity{TelemetryDruidDurationGranularity: openapiclient.NewTelemetryDruidDurationGranularity("Type_example", int64(123))}) // TelemetryDruidTimeSeriesRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTimeSeries(context.Background()).TelemetryDruidTimeSeriesRequest(telemetryDruidTimeSeriesRequest).Execute()
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
 **telemetryDruidTimeSeriesRequest** | [**TelemetryDruidTimeSeriesRequest**](TelemetryDruidTimeSeriesRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidIntervalResult**](TelemetryDruidIntervalResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTopN

> []TelemetryDruidTopNResult QueryTelemetryTopN(ctx).TelemetryDruidTopNRequest(telemetryDruidTopNRequest).Execute()

Perform a Druid TopN request.



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
    telemetryDruidTopNRequest := *openapiclient.NewTelemetryDruidTopNRequest("QueryType_example", openapiclient.telemetry.DruidDataSource{TelemetryDruidInlineDataSource: openapiclient.NewTelemetryDruidInlineDataSource("Type_example", []string{"ColumnNames_example"}, [][]string{[]string{"Rows_example"}})}, []string{"Intervals_example"}, openapiclient.telemetry.DruidGranularity{TelemetryDruidDurationGranularity: openapiclient.NewTelemetryDruidDurationGranularity("Type_example", int64(123))}, openapiclient.telemetry.DruidDimensionSpec{TelemetryDruidDefaultDimensionSpec: openapiclient.NewTelemetryDruidDefaultDimensionSpec("Type_example", "Dimension_example", "OutputName_example", "OutputType_example")}, int32(123), openapiclient.telemetry.DruidTopNMetricSpec{TelemetryDruidDimensionTopNMetricSpec: openapiclient.NewTelemetryDruidDimensionTopNMetricSpec("Type_example")}) // TelemetryDruidTopNRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTopN(context.Background()).TelemetryDruidTopNRequest(telemetryDruidTopNRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTopN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTopN`: []TelemetryDruidTopNResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTopN`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTopNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidTopNRequest** | [**TelemetryDruidTopNRequest**](TelemetryDruidTopNRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidTopNResult**](TelemetryDruidTopNResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

