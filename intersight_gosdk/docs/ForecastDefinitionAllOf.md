# ForecastDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertThresholdInPercentage** | Pointer to **int64** | Threshold above which user needs to be indicated through alarm/alert. | [optional] [readonly] 
**DataSource** | Pointer to **string** | Data source from where we get the data for the metrics to compute regression model. For example Druid. | [optional] [readonly] 
**MetricName** | Pointer to **string** | Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage. | [optional] [readonly] 
**MinNumOfDaysOfData** | Pointer to **int64** | Minimum number of days of data required for computing forecast model. | [optional] [readonly] 
**NumOfDaysOfHistoricalData** | Pointer to **int64** | Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement. | [optional] [readonly] 
**Catalog** | Pointer to [**ForecastCatalogRelationship**](forecast.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewForecastDefinitionAllOf

`func NewForecastDefinitionAllOf() *ForecastDefinitionAllOf`

NewForecastDefinitionAllOf instantiates a new ForecastDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastDefinitionAllOfWithDefaults

`func NewForecastDefinitionAllOfWithDefaults() *ForecastDefinitionAllOf`

NewForecastDefinitionAllOfWithDefaults instantiates a new ForecastDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertThresholdInPercentage

`func (o *ForecastDefinitionAllOf) GetAlertThresholdInPercentage() int64`

GetAlertThresholdInPercentage returns the AlertThresholdInPercentage field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentageOk

`func (o *ForecastDefinitionAllOf) GetAlertThresholdInPercentageOk() (*int64, bool)`

GetAlertThresholdInPercentageOk returns a tuple with the AlertThresholdInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercentage

`func (o *ForecastDefinitionAllOf) SetAlertThresholdInPercentage(v int64)`

SetAlertThresholdInPercentage sets AlertThresholdInPercentage field to given value.

### HasAlertThresholdInPercentage

`func (o *ForecastDefinitionAllOf) HasAlertThresholdInPercentage() bool`

HasAlertThresholdInPercentage returns a boolean if a field has been set.

### GetDataSource

`func (o *ForecastDefinitionAllOf) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ForecastDefinitionAllOf) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ForecastDefinitionAllOf) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ForecastDefinitionAllOf) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetMetricName

`func (o *ForecastDefinitionAllOf) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ForecastDefinitionAllOf) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ForecastDefinitionAllOf) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ForecastDefinitionAllOf) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMinNumOfDaysOfData

`func (o *ForecastDefinitionAllOf) GetMinNumOfDaysOfData() int64`

GetMinNumOfDaysOfData returns the MinNumOfDaysOfData field if non-nil, zero value otherwise.

### GetMinNumOfDaysOfDataOk

`func (o *ForecastDefinitionAllOf) GetMinNumOfDaysOfDataOk() (*int64, bool)`

GetMinNumOfDaysOfDataOk returns a tuple with the MinNumOfDaysOfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumOfDaysOfData

`func (o *ForecastDefinitionAllOf) SetMinNumOfDaysOfData(v int64)`

SetMinNumOfDaysOfData sets MinNumOfDaysOfData field to given value.

### HasMinNumOfDaysOfData

`func (o *ForecastDefinitionAllOf) HasMinNumOfDaysOfData() bool`

HasMinNumOfDaysOfData returns a boolean if a field has been set.

### GetNumOfDaysOfHistoricalData

`func (o *ForecastDefinitionAllOf) GetNumOfDaysOfHistoricalData() int64`

GetNumOfDaysOfHistoricalData returns the NumOfDaysOfHistoricalData field if non-nil, zero value otherwise.

### GetNumOfDaysOfHistoricalDataOk

`func (o *ForecastDefinitionAllOf) GetNumOfDaysOfHistoricalDataOk() (*int64, bool)`

GetNumOfDaysOfHistoricalDataOk returns a tuple with the NumOfDaysOfHistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfDaysOfHistoricalData

`func (o *ForecastDefinitionAllOf) SetNumOfDaysOfHistoricalData(v int64)`

SetNumOfDaysOfHistoricalData sets NumOfDaysOfHistoricalData field to given value.

### HasNumOfDaysOfHistoricalData

`func (o *ForecastDefinitionAllOf) HasNumOfDaysOfHistoricalData() bool`

HasNumOfDaysOfHistoricalData returns a boolean if a field has been set.

### GetPlatformType

`func (o *ForecastDefinitionAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ForecastDefinitionAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ForecastDefinitionAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ForecastDefinitionAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetCatalog

`func (o *ForecastDefinitionAllOf) GetCatalog() ForecastCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ForecastDefinitionAllOf) GetCatalogOk() (*ForecastCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ForecastDefinitionAllOf) SetCatalog(v ForecastCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ForecastDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


