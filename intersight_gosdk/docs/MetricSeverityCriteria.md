# MetricSeverityCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metric.SeverityCriteria"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metric.SeverityCriteria"]
**Operator** | Pointer to **string** | The comparison operator that evaluates the metric value against the threshold. * &#x60;greaterThan&#x60; - This checks if the metric value is greater than the threshold value. * &#x60;lessThan&#x60; - This checks if the metric value is less than the threshold value. | [optional] [default to "greaterThan"]
**Severity** | Pointer to **string** | The severity of the alarm. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [default to "None"]
**ThresholdValue** | Pointer to **string** | The threshold value of the metric for the corresponding severity. | [optional] 

## Methods

### NewMetricSeverityCriteria

`func NewMetricSeverityCriteria(classId string, objectType string, ) *MetricSeverityCriteria`

NewMetricSeverityCriteria instantiates a new MetricSeverityCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSeverityCriteriaWithDefaults

`func NewMetricSeverityCriteriaWithDefaults() *MetricSeverityCriteria`

NewMetricSeverityCriteriaWithDefaults instantiates a new MetricSeverityCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricSeverityCriteria) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricSeverityCriteria) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricSeverityCriteria) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricSeverityCriteria) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricSeverityCriteria) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricSeverityCriteria) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperator

`func (o *MetricSeverityCriteria) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetricSeverityCriteria) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetricSeverityCriteria) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MetricSeverityCriteria) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetSeverity

`func (o *MetricSeverityCriteria) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MetricSeverityCriteria) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MetricSeverityCriteria) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MetricSeverityCriteria) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetThresholdValue

`func (o *MetricSeverityCriteria) GetThresholdValue() string`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *MetricSeverityCriteria) GetThresholdValueOk() (*string, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *MetricSeverityCriteria) SetThresholdValue(v string)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *MetricSeverityCriteria) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


