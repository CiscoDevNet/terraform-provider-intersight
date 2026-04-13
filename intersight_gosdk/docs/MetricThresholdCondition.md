# MetricThresholdCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metric.ThresholdCondition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metric.ThresholdCondition"]
**Filters** | Pointer to [**[]CondFilterRule**](CondFilterRule.md) |  | [optional] 
**Granularity** | Pointer to **string** | Defines the interval at which alarm condition is evaluated. * &#x60;PT10M&#x60; - Duration of 10 minutes in the ISO 8601 duration format. * &#x60;PT1M&#x60; - Duration of 1 minute in the ISO 8601 duration format. * &#x60;PT30M&#x60; - Duration of 30 minutes in the ISO 8601 duration format. * &#x60;PT1H&#x60; - Duration of 1 hour in the ISO 8601 duration format. | [optional] [default to "PT10M"]
**GroupBy** | Pointer to **string** | A set of metric attributes used for grouping the metric data. Currently, this can be the &#39;id&#39; or &#39;host.id&#39; attributes, as alarms must be generated for Intersight-managed object. * &#x60;host.id&#x60; - The API path of the host associated with the resource component for which the metric is reported. * &#x60;id&#x60; - The API path of the resource component for which the metric is reported. | [optional] [default to "host.id"]
**InstrumentName** | Pointer to **string** | The instrument to which the metric belongs to. | [optional] 
**MetricName** | Pointer to **string** | The metric on which condition is specified. | [optional] 
**RollUpAggregate** | Pointer to **string** | Aggregation function to be applied to the grouped metric values. Example - \&quot;Avg\&quot; network utilization across all uplink ports in FI. * &#x60;Last&#x60; - Last value observed within a time duration - for example, last temperature reading. * &#x60;Average&#x60; - Arithmetic mean of the metric measurements collected over time, calculated as the sum of all metric measurements divided by the count. For example - Average temperature over time. * &#x60;AverageRate&#x60; - Rate of the metric measurements collected over time, calculated as the sum of all metric measurements divided by the duration of collection. For example - Average rate of network CRC errors over time. * &#x60;Minimum&#x60; - Smallest of all the metric values collected over time. * &#x60;MinimumRate&#x60; - Smallest of all the metric rate values collected over time. * &#x60;Maximum&#x60; - Maximum of all metric values collected over time. * &#x60;MaximumRate&#x60; - Maximum of all the metric rate values collected over time. * &#x60;Sum&#x60; - Sum of the metric values, like energy consumed over time. * &#x60;SumRate&#x60; - Sum of the metric rate values over time. | [optional] [default to "Last"]
**SeverityCriteria** | Pointer to [**[]MetricSeverityCriteria**](MetricSeverityCriteria.md) |  | [optional] 
**TotalDataPoints** | Pointer to **int64** | Total number of data points for which the alarm should be evaluated for this condition. | [optional] 
**TriggerCount** | Pointer to **int64** | The minimum number of instances that must satisfy the alarm criteria. | [optional] 

## Methods

### NewMetricThresholdCondition

`func NewMetricThresholdCondition(classId string, objectType string, ) *MetricThresholdCondition`

NewMetricThresholdCondition instantiates a new MetricThresholdCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricThresholdConditionWithDefaults

`func NewMetricThresholdConditionWithDefaults() *MetricThresholdCondition`

NewMetricThresholdConditionWithDefaults instantiates a new MetricThresholdCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricThresholdCondition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricThresholdCondition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricThresholdCondition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricThresholdCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricThresholdCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricThresholdCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilters

`func (o *MetricThresholdCondition) GetFilters() []CondFilterRule`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricThresholdCondition) GetFiltersOk() (*[]CondFilterRule, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricThresholdCondition) SetFilters(v []CondFilterRule)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricThresholdCondition) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *MetricThresholdCondition) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *MetricThresholdCondition) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetGranularity

`func (o *MetricThresholdCondition) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricThresholdCondition) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricThresholdCondition) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricThresholdCondition) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetGroupBy

`func (o *MetricThresholdCondition) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricThresholdCondition) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricThresholdCondition) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricThresholdCondition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetInstrumentName

`func (o *MetricThresholdCondition) GetInstrumentName() string`

GetInstrumentName returns the InstrumentName field if non-nil, zero value otherwise.

### GetInstrumentNameOk

`func (o *MetricThresholdCondition) GetInstrumentNameOk() (*string, bool)`

GetInstrumentNameOk returns a tuple with the InstrumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentName

`func (o *MetricThresholdCondition) SetInstrumentName(v string)`

SetInstrumentName sets InstrumentName field to given value.

### HasInstrumentName

`func (o *MetricThresholdCondition) HasInstrumentName() bool`

HasInstrumentName returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricThresholdCondition) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricThresholdCondition) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricThresholdCondition) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MetricThresholdCondition) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetRollUpAggregate

`func (o *MetricThresholdCondition) GetRollUpAggregate() string`

GetRollUpAggregate returns the RollUpAggregate field if non-nil, zero value otherwise.

### GetRollUpAggregateOk

`func (o *MetricThresholdCondition) GetRollUpAggregateOk() (*string, bool)`

GetRollUpAggregateOk returns a tuple with the RollUpAggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollUpAggregate

`func (o *MetricThresholdCondition) SetRollUpAggregate(v string)`

SetRollUpAggregate sets RollUpAggregate field to given value.

### HasRollUpAggregate

`func (o *MetricThresholdCondition) HasRollUpAggregate() bool`

HasRollUpAggregate returns a boolean if a field has been set.

### GetSeverityCriteria

`func (o *MetricThresholdCondition) GetSeverityCriteria() []MetricSeverityCriteria`

GetSeverityCriteria returns the SeverityCriteria field if non-nil, zero value otherwise.

### GetSeverityCriteriaOk

`func (o *MetricThresholdCondition) GetSeverityCriteriaOk() (*[]MetricSeverityCriteria, bool)`

GetSeverityCriteriaOk returns a tuple with the SeverityCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCriteria

`func (o *MetricThresholdCondition) SetSeverityCriteria(v []MetricSeverityCriteria)`

SetSeverityCriteria sets SeverityCriteria field to given value.

### HasSeverityCriteria

`func (o *MetricThresholdCondition) HasSeverityCriteria() bool`

HasSeverityCriteria returns a boolean if a field has been set.

### SetSeverityCriteriaNil

`func (o *MetricThresholdCondition) SetSeverityCriteriaNil(b bool)`

 SetSeverityCriteriaNil sets the value for SeverityCriteria to be an explicit nil

### UnsetSeverityCriteria
`func (o *MetricThresholdCondition) UnsetSeverityCriteria()`

UnsetSeverityCriteria ensures that no value is present for SeverityCriteria, not even an explicit nil
### GetTotalDataPoints

`func (o *MetricThresholdCondition) GetTotalDataPoints() int64`

GetTotalDataPoints returns the TotalDataPoints field if non-nil, zero value otherwise.

### GetTotalDataPointsOk

`func (o *MetricThresholdCondition) GetTotalDataPointsOk() (*int64, bool)`

GetTotalDataPointsOk returns a tuple with the TotalDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataPoints

`func (o *MetricThresholdCondition) SetTotalDataPoints(v int64)`

SetTotalDataPoints sets TotalDataPoints field to given value.

### HasTotalDataPoints

`func (o *MetricThresholdCondition) HasTotalDataPoints() bool`

HasTotalDataPoints returns a boolean if a field has been set.

### GetTriggerCount

`func (o *MetricThresholdCondition) GetTriggerCount() int64`

GetTriggerCount returns the TriggerCount field if non-nil, zero value otherwise.

### GetTriggerCountOk

`func (o *MetricThresholdCondition) GetTriggerCountOk() (*int64, bool)`

GetTriggerCountOk returns a tuple with the TriggerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCount

`func (o *MetricThresholdCondition) SetTriggerCount(v int64)`

SetTriggerCount sets TriggerCount field to given value.

### HasTriggerCount

`func (o *MetricThresholdCondition) HasTriggerCount() bool`

HasTriggerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


