# MetricsMetricCriterion

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

### NewMetricsMetricCriterion

`func NewMetricsMetricCriterion(classId string, objectType string, ) *MetricsMetricCriterion`

NewMetricsMetricCriterion instantiates a new MetricsMetricCriterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetricCriterionWithDefaults

`func NewMetricsMetricCriterionWithDefaults() *MetricsMetricCriterion`

NewMetricsMetricCriterionWithDefaults instantiates a new MetricsMetricCriterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsMetricCriterion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsMetricCriterion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsMetricCriterion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsMetricCriterion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsMetricCriterion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsMetricCriterion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregation

`func (o *MetricsMetricCriterion) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricsMetricCriterion) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricsMetricCriterion) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MetricsMetricCriterion) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetFilters

`func (o *MetricsMetricCriterion) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricsMetricCriterion) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricsMetricCriterion) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricsMetricCriterion) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *MetricsMetricCriterion) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *MetricsMetricCriterion) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetGroups

`func (o *MetricsMetricCriterion) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MetricsMetricCriterion) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MetricsMetricCriterion) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MetricsMetricCriterion) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *MetricsMetricCriterion) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *MetricsMetricCriterion) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetInstrument

`func (o *MetricsMetricCriterion) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *MetricsMetricCriterion) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *MetricsMetricCriterion) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *MetricsMetricCriterion) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MetricsMetricCriterion) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MetricsMetricCriterion) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MetricsMetricCriterion) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MetricsMetricCriterion) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMetric

`func (o *MetricsMetricCriterion) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsMetricCriterion) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsMetricCriterion) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricsMetricCriterion) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMetricAggregation

`func (o *MetricsMetricCriterion) GetMetricAggregation() string`

GetMetricAggregation returns the MetricAggregation field if non-nil, zero value otherwise.

### GetMetricAggregationOk

`func (o *MetricsMetricCriterion) GetMetricAggregationOk() (*string, bool)`

GetMetricAggregationOk returns a tuple with the MetricAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricAggregation

`func (o *MetricsMetricCriterion) SetMetricAggregation(v string)`

SetMetricAggregation sets MetricAggregation field to given value.

### HasMetricAggregation

`func (o *MetricsMetricCriterion) HasMetricAggregation() bool`

HasMetricAggregation returns a boolean if a field has been set.

### GetTopLimit

`func (o *MetricsMetricCriterion) GetTopLimit() int64`

GetTopLimit returns the TopLimit field if non-nil, zero value otherwise.

### GetTopLimitOk

`func (o *MetricsMetricCriterion) GetTopLimitOk() (*int64, bool)`

GetTopLimitOk returns a tuple with the TopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLimit

`func (o *MetricsMetricCriterion) SetTopLimit(v int64)`

SetTopLimit sets TopLimit field to given value.

### HasTopLimit

`func (o *MetricsMetricCriterion) HasTopLimit() bool`

HasTopLimit returns a boolean if a field has been set.

### GetTopSort

`func (o *MetricsMetricCriterion) GetTopSort() string`

GetTopSort returns the TopSort field if non-nil, zero value otherwise.

### GetTopSortOk

`func (o *MetricsMetricCriterion) GetTopSortOk() (*string, bool)`

GetTopSortOk returns a tuple with the TopSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSort

`func (o *MetricsMetricCriterion) SetTopSort(v string)`

SetTopSort sets TopSort field to given value.

### HasTopSort

`func (o *MetricsMetricCriterion) HasTopSort() bool`

HasTopSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


