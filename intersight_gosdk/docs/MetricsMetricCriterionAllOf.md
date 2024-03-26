# MetricsMetricCriterionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metrics.MetricCriterion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metrics.MetricCriterion"]
**Aggregation** | Pointer to **string** | Function name which used to combine the group buckets into a single timeseries. | [optional] 
**Filters** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Instrument** | Pointer to **string** | Instrument name used to collect measurements for the query. | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates if this criterion should be used for the query. | [optional] 
**Metric** | Pointer to **string** | Measurement name that is collected by the instrument for the query. | [optional] 
**MetricAggregation** | Pointer to **string** | Function name which used to combine the metrics into granularity buckets. | [optional] 
**TopLimit** | Pointer to **int64** | The maximum number of result rows. | [optional] 
**TopSort** | Pointer to **string** | Method on how to sort the result rows. | [optional] 

## Methods

### NewMetricsMetricCriterionAllOf

`func NewMetricsMetricCriterionAllOf(classId string, objectType string, ) *MetricsMetricCriterionAllOf`

NewMetricsMetricCriterionAllOf instantiates a new MetricsMetricCriterionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetricCriterionAllOfWithDefaults

`func NewMetricsMetricCriterionAllOfWithDefaults() *MetricsMetricCriterionAllOf`

NewMetricsMetricCriterionAllOfWithDefaults instantiates a new MetricsMetricCriterionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsMetricCriterionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsMetricCriterionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsMetricCriterionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsMetricCriterionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsMetricCriterionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsMetricCriterionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregation

`func (o *MetricsMetricCriterionAllOf) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsMetricCriterionAllOf) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsMetricCriterionAllOf) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsMetricCriterionAllOf) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetFilters

`func (o *MetricsMetricCriterionAllOf) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricsMetricCriterionAllOf) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricsMetricCriterionAllOf) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricsMetricCriterionAllOf) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *MetricsMetricCriterionAllOf) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *MetricsMetricCriterionAllOf) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetGroups

`func (o *MetricsMetricCriterionAllOf) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MetricsMetricCriterionAllOf) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MetricsMetricCriterionAllOf) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MetricsMetricCriterionAllOf) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *MetricsMetricCriterionAllOf) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *MetricsMetricCriterionAllOf) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetInstrument

`func (o *MetricsMetricCriterionAllOf) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *MetricsMetricCriterionAllOf) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *MetricsMetricCriterionAllOf) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *MetricsMetricCriterionAllOf) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MetricsMetricCriterionAllOf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MetricsMetricCriterionAllOf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MetricsMetricCriterionAllOf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MetricsMetricCriterionAllOf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetric

`func (o *MetricsMetricCriterionAllOf) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsMetricCriterionAllOf) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsMetricCriterionAllOf) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsMetricCriterionAllOf) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMetricAggregation

`func (o *MetricsMetricCriterionAllOf) GetMetricAggregation() string`

GetMetricAggregation returns the MetricAggregation field if non-nil, zero value otherwise.

### GetMetricAggregationOk

`func (o *MetricsMetricCriterionAllOf) GetMetricAggregationOk() (*string, bool)`

GetMetricAggregationOk returns a tuple with the MetricAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricAggregation

`func (o *MetricsMetricCriterionAllOf) SetMetricAggregation(v string)`

SetMetricAggregation sets MetricAggregation field to given value.

### HasMetricAggregation

`func (o *MetricsMetricCriterionAllOf) HasMetricAggregation() bool`

HasMetricAggregation returns a boolean if a field has been set.

### GetTopLimit

`func (o *MetricsMetricCriterionAllOf) GetTopLimit() int64`

GetTopLimit returns the TopLimit field if non-nil, zero value otherwise.

### GetTopLimitOk

`func (o *MetricsMetricCriterionAllOf) GetTopLimitOk() (*int64, bool)`

GetTopLimitOk returns a tuple with the TopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLimit

`func (o *MetricsMetricCriterionAllOf) SetTopLimit(v int64)`

SetTopLimit sets TopLimit field to given value.

### HasTopLimit

`func (o *MetricsMetricCriterionAllOf) HasTopLimit() bool`

HasTopLimit returns a boolean if a field has been set.

### GetTopSort

`func (o *MetricsMetricCriterionAllOf) GetTopSort() string`

GetTopSort returns the TopSort field if non-nil, zero value otherwise.

### GetTopSortOk

`func (o *MetricsMetricCriterionAllOf) GetTopSortOk() (*string, bool)`

GetTopSortOk returns a tuple with the TopSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSort

`func (o *MetricsMetricCriterionAllOf) SetTopSort(v string)`

SetTopSort sets TopSort field to given value.

### HasTopSort

`func (o *MetricsMetricCriterionAllOf) HasTopSort() bool`

HasTopSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


