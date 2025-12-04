# ApplianceUtilizationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "appliance.MetricsIngestionUtilization"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "appliance.MetricsIngestionUtilization"]
**Utilization** | Pointer to **float64** | The percent utilization of the metric. | [optional] [readonly] 

## Methods

### NewApplianceUtilizationMetric

`func NewApplianceUtilizationMetric(classId string, objectType string, ) *ApplianceUtilizationMetric`

NewApplianceUtilizationMetric instantiates a new ApplianceUtilizationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUtilizationMetricWithDefaults

`func NewApplianceUtilizationMetricWithDefaults() *ApplianceUtilizationMetric`

NewApplianceUtilizationMetricWithDefaults instantiates a new ApplianceUtilizationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceUtilizationMetric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceUtilizationMetric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceUtilizationMetric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceUtilizationMetric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceUtilizationMetric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceUtilizationMetric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUtilization

`func (o *ApplianceUtilizationMetric) GetUtilization() float64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *ApplianceUtilizationMetric) GetUtilizationOk() (*float64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *ApplianceUtilizationMetric) SetUtilization(v float64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *ApplianceUtilizationMetric) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


