# ApplianceNodeUtilizationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Utilization** | Pointer to **float64** | The percent utilization of the metric. | [optional] [readonly] 

## Methods

### NewApplianceNodeUtilizationMetric

`func NewApplianceNodeUtilizationMetric(classId string, objectType string, ) *ApplianceNodeUtilizationMetric`

NewApplianceNodeUtilizationMetric instantiates a new ApplianceNodeUtilizationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeUtilizationMetricWithDefaults

`func NewApplianceNodeUtilizationMetricWithDefaults() *ApplianceNodeUtilizationMetric`

NewApplianceNodeUtilizationMetricWithDefaults instantiates a new ApplianceNodeUtilizationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeUtilizationMetric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeUtilizationMetric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeUtilizationMetric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeUtilizationMetric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeUtilizationMetric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeUtilizationMetric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUtilization

`func (o *ApplianceNodeUtilizationMetric) GetUtilization() float64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *ApplianceNodeUtilizationMetric) GetUtilizationOk() (*float64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *ApplianceNodeUtilizationMetric) SetUtilization(v float64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *ApplianceNodeUtilizationMetric) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


