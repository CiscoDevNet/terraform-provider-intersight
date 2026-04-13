# MetricThresholdDefinitionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metric.ThresholdDefinitionState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metric.ThresholdDefinitionState"]
**Enabled** | Pointer to **bool** | Controls the behavior of alarm processing based on the criteria set for this metric. If set to true which is also the default behavior, the alarm is evaluated against the criteria set for the metric. If set to false, the alarm is not evaluated, and any existing alarms triggered by the criteria set for the metric are cleared. | [optional] [default to true]
**InstrumentName** | Pointer to **string** | The instrument to which the metric belongs to. | [optional] 
**MetricName** | Pointer to **string** | The metric for which rule is specified. | [optional] 

## Methods

### NewMetricThresholdDefinitionState

`func NewMetricThresholdDefinitionState(classId string, objectType string, ) *MetricThresholdDefinitionState`

NewMetricThresholdDefinitionState instantiates a new MetricThresholdDefinitionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricThresholdDefinitionStateWithDefaults

`func NewMetricThresholdDefinitionStateWithDefaults() *MetricThresholdDefinitionState`

NewMetricThresholdDefinitionStateWithDefaults instantiates a new MetricThresholdDefinitionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricThresholdDefinitionState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricThresholdDefinitionState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricThresholdDefinitionState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricThresholdDefinitionState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricThresholdDefinitionState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricThresholdDefinitionState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *MetricThresholdDefinitionState) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricThresholdDefinitionState) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricThresholdDefinitionState) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricThresholdDefinitionState) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstrumentName

`func (o *MetricThresholdDefinitionState) GetInstrumentName() string`

GetInstrumentName returns the InstrumentName field if non-nil, zero value otherwise.

### GetInstrumentNameOk

`func (o *MetricThresholdDefinitionState) GetInstrumentNameOk() (*string, bool)`

GetInstrumentNameOk returns a tuple with the InstrumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentName

`func (o *MetricThresholdDefinitionState) SetInstrumentName(v string)`

SetInstrumentName sets InstrumentName field to given value.

### HasInstrumentName

`func (o *MetricThresholdDefinitionState) HasInstrumentName() bool`

HasInstrumentName returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricThresholdDefinitionState) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricThresholdDefinitionState) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricThresholdDefinitionState) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MetricThresholdDefinitionState) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


