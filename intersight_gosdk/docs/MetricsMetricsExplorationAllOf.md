# MetricsMetricsExplorationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metrics.MetricsExploration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metrics.MetricsExploration"]
**Description** | Pointer to **string** | User specified description of this MetricsExploration. | [optional] 
**Granularity** | Pointer to **string** | The time unit in which the metrics will be aggregated into. | [optional] 
**Intervals** | Pointer to **[]string** |  | [optional] 
**IsWidget** | Pointer to **bool** | True when this MetricsExploration is exposed as a Dashlet widget. | [optional] 
**MetricCriteria** | Pointer to [**[]MetricsMetricCriterion**](MetricsMetricCriterion.md) |  | [optional] 
**Name** | Pointer to **string** | User specified name of this MetricsExploration. | [optional] 
**RawQueries** | Pointer to **[]string** |  | [optional] 
**VisualConfig** | Pointer to **interface{}** | Chart configuration options. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMetricsMetricsExplorationAllOf

`func NewMetricsMetricsExplorationAllOf(classId string, objectType string, ) *MetricsMetricsExplorationAllOf`

NewMetricsMetricsExplorationAllOf instantiates a new MetricsMetricsExplorationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetricsExplorationAllOfWithDefaults

`func NewMetricsMetricsExplorationAllOfWithDefaults() *MetricsMetricsExplorationAllOf`

NewMetricsMetricsExplorationAllOfWithDefaults instantiates a new MetricsMetricsExplorationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsMetricsExplorationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsMetricsExplorationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsMetricsExplorationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsMetricsExplorationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsMetricsExplorationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsMetricsExplorationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *MetricsMetricsExplorationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsMetricsExplorationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsMetricsExplorationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsMetricsExplorationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGranularity

`func (o *MetricsMetricsExplorationAllOf) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsMetricsExplorationAllOf) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsMetricsExplorationAllOf) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricsMetricsExplorationAllOf) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetIntervals

`func (o *MetricsMetricsExplorationAllOf) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *MetricsMetricsExplorationAllOf) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *MetricsMetricsExplorationAllOf) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *MetricsMetricsExplorationAllOf) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### SetIntervalsNil

`func (o *MetricsMetricsExplorationAllOf) SetIntervalsNil(b bool)`

 SetIntervalsNil sets the value for Intervals to be an explicit nil

### UnsetIntervals
`func (o *MetricsMetricsExplorationAllOf) UnsetIntervals()`

UnsetIntervals ensures that no value is present for Intervals, not even an explicit nil
### GetIsWidget

`func (o *MetricsMetricsExplorationAllOf) GetIsWidget() bool`

GetIsWidget returns the IsWidget field if non-nil, zero value otherwise.

### GetIsWidgetOk

`func (o *MetricsMetricsExplorationAllOf) GetIsWidgetOk() (*bool, bool)`

GetIsWidgetOk returns a tuple with the IsWidget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWidget

`func (o *MetricsMetricsExplorationAllOf) SetIsWidget(v bool)`

SetIsWidget sets IsWidget field to given value.

### HasIsWidget

`func (o *MetricsMetricsExplorationAllOf) HasIsWidget() bool`

HasIsWidget returns a boolean if a field has been set.

### GetMetricCriteria

`func (o *MetricsMetricsExplorationAllOf) GetMetricCriteria() []MetricsMetricCriterion`

GetMetricCriteria returns the MetricCriteria field if non-nil, zero value otherwise.

### GetMetricCriteriaOk

`func (o *MetricsMetricsExplorationAllOf) GetMetricCriteriaOk() (*[]MetricsMetricCriterion, bool)`

GetMetricCriteriaOk returns a tuple with the MetricCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCriteria

`func (o *MetricsMetricsExplorationAllOf) SetMetricCriteria(v []MetricsMetricCriterion)`

SetMetricCriteria sets MetricCriteria field to given value.

### HasMetricCriteria

`func (o *MetricsMetricsExplorationAllOf) HasMetricCriteria() bool`

HasMetricCriteria returns a boolean if a field has been set.

### SetMetricCriteriaNil

`func (o *MetricsMetricsExplorationAllOf) SetMetricCriteriaNil(b bool)`

 SetMetricCriteriaNil sets the value for MetricCriteria to be an explicit nil

### UnsetMetricCriteria
`func (o *MetricsMetricsExplorationAllOf) UnsetMetricCriteria()`

UnsetMetricCriteria ensures that no value is present for MetricCriteria, not even an explicit nil
### GetName

`func (o *MetricsMetricsExplorationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsMetricsExplorationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsMetricsExplorationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsMetricsExplorationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawQueries

`func (o *MetricsMetricsExplorationAllOf) GetRawQueries() []string`

GetRawQueries returns the RawQueries field if non-nil, zero value otherwise.

### GetRawQueriesOk

`func (o *MetricsMetricsExplorationAllOf) GetRawQueriesOk() (*[]string, bool)`

GetRawQueriesOk returns a tuple with the RawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQueries

`func (o *MetricsMetricsExplorationAllOf) SetRawQueries(v []string)`

SetRawQueries sets RawQueries field to given value.

### HasRawQueries

`func (o *MetricsMetricsExplorationAllOf) HasRawQueries() bool`

HasRawQueries returns a boolean if a field has been set.

### SetRawQueriesNil

`func (o *MetricsMetricsExplorationAllOf) SetRawQueriesNil(b bool)`

 SetRawQueriesNil sets the value for RawQueries to be an explicit nil

### UnsetRawQueries
`func (o *MetricsMetricsExplorationAllOf) UnsetRawQueries()`

UnsetRawQueries ensures that no value is present for RawQueries, not even an explicit nil
### GetVisualConfig

`func (o *MetricsMetricsExplorationAllOf) GetVisualConfig() interface{}`

GetVisualConfig returns the VisualConfig field if non-nil, zero value otherwise.

### GetVisualConfigOk

`func (o *MetricsMetricsExplorationAllOf) GetVisualConfigOk() (*interface{}, bool)`

GetVisualConfigOk returns a tuple with the VisualConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualConfig

`func (o *MetricsMetricsExplorationAllOf) SetVisualConfig(v interface{})`

SetVisualConfig sets VisualConfig field to given value.

### HasVisualConfig

`func (o *MetricsMetricsExplorationAllOf) HasVisualConfig() bool`

HasVisualConfig returns a boolean if a field has been set.

### SetVisualConfigNil

`func (o *MetricsMetricsExplorationAllOf) SetVisualConfigNil(b bool)`

 SetVisualConfigNil sets the value for VisualConfig to be an explicit nil

### UnsetVisualConfig
`func (o *MetricsMetricsExplorationAllOf) UnsetVisualConfig()`

UnsetVisualConfig ensures that no value is present for VisualConfig, not even an explicit nil
### GetOrganization

`func (o *MetricsMetricsExplorationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MetricsMetricsExplorationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MetricsMetricsExplorationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MetricsMetricsExplorationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


