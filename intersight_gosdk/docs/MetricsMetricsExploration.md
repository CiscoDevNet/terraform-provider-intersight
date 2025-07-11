# MetricsMetricsExploration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metrics.MetricsExploration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metrics.MetricsExploration"]
**Description** | Pointer to **string** | User specified description of the MetricsExploration. | [optional] 
**Granularity** | Pointer to **string** | Time unit for aggregating the metrics. | [optional] 
**Intervals** | Pointer to **[]string** |  | [optional] 
**IsWidget** | Pointer to **bool** | Set to true when the MetricsExploration is presented as a Dashlet widget. | [optional] 
**MetricCriteria** | Pointer to [**[]MetricsMetricCriterion**](MetricsMetricCriterion.md) |  | [optional] 
**Name** | Pointer to **string** | User specified name of the MetricsExploration. | [optional] 
**RawQueries** | Pointer to **[]string** |  | [optional] 
**VisualConfig** | Pointer to **interface{}** | Chart configuration options. | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMetricsMetricsExploration

`func NewMetricsMetricsExploration(classId string, objectType string, ) *MetricsMetricsExploration`

NewMetricsMetricsExploration instantiates a new MetricsMetricsExploration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMetricsExplorationWithDefaults

`func NewMetricsMetricsExplorationWithDefaults() *MetricsMetricsExploration`

NewMetricsMetricsExplorationWithDefaults instantiates a new MetricsMetricsExploration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsMetricsExploration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsMetricsExploration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsMetricsExploration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsMetricsExploration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsMetricsExploration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsMetricsExploration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *MetricsMetricsExploration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsMetricsExploration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsMetricsExploration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsMetricsExploration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGranularity

`func (o *MetricsMetricsExploration) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MetricsMetricsExploration) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MetricsMetricsExploration) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MetricsMetricsExploration) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetIntervals

`func (o *MetricsMetricsExploration) GetIntervals() []string`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *MetricsMetricsExploration) GetIntervalsOk() (*[]string, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *MetricsMetricsExploration) SetIntervals(v []string)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *MetricsMetricsExploration) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### SetIntervalsNil

`func (o *MetricsMetricsExploration) SetIntervalsNil(b bool)`

 SetIntervalsNil sets the value for Intervals to be an explicit nil

### UnsetIntervals
`func (o *MetricsMetricsExploration) UnsetIntervals()`

UnsetIntervals ensures that no value is present for Intervals, not even an explicit nil
### GetIsWidget

`func (o *MetricsMetricsExploration) GetIsWidget() bool`

GetIsWidget returns the IsWidget field if non-nil, zero value otherwise.

### GetIsWidgetOk

`func (o *MetricsMetricsExploration) GetIsWidgetOk() (*bool, bool)`

GetIsWidgetOk returns a tuple with the IsWidget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWidget

`func (o *MetricsMetricsExploration) SetIsWidget(v bool)`

SetIsWidget sets IsWidget field to given value.

### HasIsWidget

`func (o *MetricsMetricsExploration) HasIsWidget() bool`

HasIsWidget returns a boolean if a field has been set.

### GetMetricCriteria

`func (o *MetricsMetricsExploration) GetMetricCriteria() []MetricsMetricCriterion`

GetMetricCriteria returns the MetricCriteria field if non-nil, zero value otherwise.

### GetMetricCriteriaOk

`func (o *MetricsMetricsExploration) GetMetricCriteriaOk() (*[]MetricsMetricCriterion, bool)`

GetMetricCriteriaOk returns a tuple with the MetricCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCriteria

`func (o *MetricsMetricsExploration) SetMetricCriteria(v []MetricsMetricCriterion)`

SetMetricCriteria sets MetricCriteria field to given value.

### HasMetricCriteria

`func (o *MetricsMetricsExploration) HasMetricCriteria() bool`

HasMetricCriteria returns a boolean if a field has been set.

### SetMetricCriteriaNil

`func (o *MetricsMetricsExploration) SetMetricCriteriaNil(b bool)`

 SetMetricCriteriaNil sets the value for MetricCriteria to be an explicit nil

### UnsetMetricCriteria
`func (o *MetricsMetricsExploration) UnsetMetricCriteria()`

UnsetMetricCriteria ensures that no value is present for MetricCriteria, not even an explicit nil
### GetName

`func (o *MetricsMetricsExploration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricsMetricsExploration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricsMetricsExploration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricsMetricsExploration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawQueries

`func (o *MetricsMetricsExploration) GetRawQueries() []string`

GetRawQueries returns the RawQueries field if non-nil, zero value otherwise.

### GetRawQueriesOk

`func (o *MetricsMetricsExploration) GetRawQueriesOk() (*[]string, bool)`

GetRawQueriesOk returns a tuple with the RawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQueries

`func (o *MetricsMetricsExploration) SetRawQueries(v []string)`

SetRawQueries sets RawQueries field to given value.

### HasRawQueries

`func (o *MetricsMetricsExploration) HasRawQueries() bool`

HasRawQueries returns a boolean if a field has been set.

### SetRawQueriesNil

`func (o *MetricsMetricsExploration) SetRawQueriesNil(b bool)`

 SetRawQueriesNil sets the value for RawQueries to be an explicit nil

### UnsetRawQueries
`func (o *MetricsMetricsExploration) UnsetRawQueries()`

UnsetRawQueries ensures that no value is present for RawQueries, not even an explicit nil
### GetVisualConfig

`func (o *MetricsMetricsExploration) GetVisualConfig() interface{}`

GetVisualConfig returns the VisualConfig field if non-nil, zero value otherwise.

### GetVisualConfigOk

`func (o *MetricsMetricsExploration) GetVisualConfigOk() (*interface{}, bool)`

GetVisualConfigOk returns a tuple with the VisualConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualConfig

`func (o *MetricsMetricsExploration) SetVisualConfig(v interface{})`

SetVisualConfig sets VisualConfig field to given value.

### HasVisualConfig

`func (o *MetricsMetricsExploration) HasVisualConfig() bool`

HasVisualConfig returns a boolean if a field has been set.

### SetVisualConfigNil

`func (o *MetricsMetricsExploration) SetVisualConfigNil(b bool)`

 SetVisualConfigNil sets the value for VisualConfig to be an explicit nil

### UnsetVisualConfig
`func (o *MetricsMetricsExploration) UnsetVisualConfig()`

UnsetVisualConfig ensures that no value is present for VisualConfig, not even an explicit nil
### GetOrganization

`func (o *MetricsMetricsExploration) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MetricsMetricsExploration) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MetricsMetricsExploration) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MetricsMetricsExploration) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MetricsMetricsExploration) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MetricsMetricsExploration) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


