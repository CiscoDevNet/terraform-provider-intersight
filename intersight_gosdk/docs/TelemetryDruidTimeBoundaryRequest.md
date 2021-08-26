# TelemetryDruidTimeBoundaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | null | 
**DataSource** | [**TelemetryDruidDataSource**](TelemetryDruidDataSource.md) |  | 
**Bound** | Pointer to **string** | Optional, set to maxTime or minTime to return only the latest or earliest timestamp. Default to returning both if not set. | [optional] 
**Filter** | Pointer to [**TelemetryDruidFilter**](TelemetryDruidFilter.md) |  | [optional] 
**Context** | Pointer to [**TelemetryDruidQueryContext**](TelemetryDruidQueryContext.md) |  | [optional] 

## Methods

### NewTelemetryDruidTimeBoundaryRequest

`func NewTelemetryDruidTimeBoundaryRequest(queryType string, dataSource TelemetryDruidDataSource, ) *TelemetryDruidTimeBoundaryRequest`

NewTelemetryDruidTimeBoundaryRequest instantiates a new TelemetryDruidTimeBoundaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidTimeBoundaryRequestWithDefaults

`func NewTelemetryDruidTimeBoundaryRequestWithDefaults() *TelemetryDruidTimeBoundaryRequest`

NewTelemetryDruidTimeBoundaryRequestWithDefaults instantiates a new TelemetryDruidTimeBoundaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *TelemetryDruidTimeBoundaryRequest) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *TelemetryDruidTimeBoundaryRequest) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *TelemetryDruidTimeBoundaryRequest) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetDataSource

`func (o *TelemetryDruidTimeBoundaryRequest) GetDataSource() TelemetryDruidDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *TelemetryDruidTimeBoundaryRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *TelemetryDruidTimeBoundaryRequest) SetDataSource(v TelemetryDruidDataSource)`

SetDataSource sets DataSource field to given value.


### GetBound

`func (o *TelemetryDruidTimeBoundaryRequest) GetBound() string`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *TelemetryDruidTimeBoundaryRequest) GetBoundOk() (*string, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *TelemetryDruidTimeBoundaryRequest) SetBound(v string)`

SetBound sets Bound field to given value.

### HasBound

`func (o *TelemetryDruidTimeBoundaryRequest) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetFilter

`func (o *TelemetryDruidTimeBoundaryRequest) GetFilter() TelemetryDruidFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TelemetryDruidTimeBoundaryRequest) GetFilterOk() (*TelemetryDruidFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TelemetryDruidTimeBoundaryRequest) SetFilter(v TelemetryDruidFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TelemetryDruidTimeBoundaryRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetContext

`func (o *TelemetryDruidTimeBoundaryRequest) GetContext() TelemetryDruidQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TelemetryDruidTimeBoundaryRequest) GetContextOk() (*TelemetryDruidQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TelemetryDruidTimeBoundaryRequest) SetContext(v TelemetryDruidQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TelemetryDruidTimeBoundaryRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


