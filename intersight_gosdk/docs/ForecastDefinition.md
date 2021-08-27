# ForecastDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "forecast.Definition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "forecast.Definition"]
**AlertThresholdInPercentage** | Pointer to **int64** | Threshold above which user needs to be indicated through alarm/alert. | [optional] [readonly] 
**DataSource** | Pointer to **string** | Data source from where we get the data for the metrics to compute regression model. For example Druid. | [optional] [readonly] 
**MetricName** | Pointer to **string** | Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage. | [optional] [readonly] 
**MinNumOfDaysOfData** | Pointer to **int64** | Minimum number of days of data required for computing forecast model. | [optional] [readonly] 
**NumOfDaysOfHistoricalData** | Pointer to **int64** | Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model. | [optional] [readonly] 
**PlatformType** | Pointer to **string** | The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement. | [optional] [readonly] 
**Catalog** | Pointer to [**ForecastCatalogRelationship**](ForecastCatalogRelationship.md) |  | [optional] 

## Methods

### NewForecastDefinition

`func NewForecastDefinition(classId string, objectType string, ) *ForecastDefinition`

NewForecastDefinition instantiates a new ForecastDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastDefinitionWithDefaults

`func NewForecastDefinitionWithDefaults() *ForecastDefinition`

NewForecastDefinitionWithDefaults instantiates a new ForecastDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ForecastDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ForecastDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ForecastDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ForecastDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ForecastDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ForecastDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlertThresholdInPercentage

`func (o *ForecastDefinition) GetAlertThresholdInPercentage() int64`

GetAlertThresholdInPercentage returns the AlertThresholdInPercentage field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentageOk

`func (o *ForecastDefinition) GetAlertThresholdInPercentageOk() (*int64, bool)`

GetAlertThresholdInPercentageOk returns a tuple with the AlertThresholdInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercentage

`func (o *ForecastDefinition) SetAlertThresholdInPercentage(v int64)`

SetAlertThresholdInPercentage sets AlertThresholdInPercentage field to given value.

### HasAlertThresholdInPercentage

`func (o *ForecastDefinition) HasAlertThresholdInPercentage() bool`

HasAlertThresholdInPercentage returns a boolean if a field has been set.

### GetDataSource

`func (o *ForecastDefinition) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ForecastDefinition) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ForecastDefinition) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ForecastDefinition) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetMetricName

`func (o *ForecastDefinition) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ForecastDefinition) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ForecastDefinition) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ForecastDefinition) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMinNumOfDaysOfData

`func (o *ForecastDefinition) GetMinNumOfDaysOfData() int64`

GetMinNumOfDaysOfData returns the MinNumOfDaysOfData field if non-nil, zero value otherwise.

### GetMinNumOfDaysOfDataOk

`func (o *ForecastDefinition) GetMinNumOfDaysOfDataOk() (*int64, bool)`

GetMinNumOfDaysOfDataOk returns a tuple with the MinNumOfDaysOfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumOfDaysOfData

`func (o *ForecastDefinition) SetMinNumOfDaysOfData(v int64)`

SetMinNumOfDaysOfData sets MinNumOfDaysOfData field to given value.

### HasMinNumOfDaysOfData

`func (o *ForecastDefinition) HasMinNumOfDaysOfData() bool`

HasMinNumOfDaysOfData returns a boolean if a field has been set.

### GetNumOfDaysOfHistoricalData

`func (o *ForecastDefinition) GetNumOfDaysOfHistoricalData() int64`

GetNumOfDaysOfHistoricalData returns the NumOfDaysOfHistoricalData field if non-nil, zero value otherwise.

### GetNumOfDaysOfHistoricalDataOk

`func (o *ForecastDefinition) GetNumOfDaysOfHistoricalDataOk() (*int64, bool)`

GetNumOfDaysOfHistoricalDataOk returns a tuple with the NumOfDaysOfHistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfDaysOfHistoricalData

`func (o *ForecastDefinition) SetNumOfDaysOfHistoricalData(v int64)`

SetNumOfDaysOfHistoricalData sets NumOfDaysOfHistoricalData field to given value.

### HasNumOfDaysOfHistoricalData

`func (o *ForecastDefinition) HasNumOfDaysOfHistoricalData() bool`

HasNumOfDaysOfHistoricalData returns a boolean if a field has been set.

### GetPlatformType

`func (o *ForecastDefinition) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *ForecastDefinition) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *ForecastDefinition) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *ForecastDefinition) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetCatalog

`func (o *ForecastDefinition) GetCatalog() ForecastCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ForecastDefinition) GetCatalogOk() (*ForecastCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ForecastDefinition) SetCatalog(v ForecastCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ForecastDefinition) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


