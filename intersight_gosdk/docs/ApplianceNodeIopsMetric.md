# ApplianceNodeIopsMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.NodeIopsMetric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.NodeIopsMetric"]
**Iops** | Pointer to **float64** | Aggregate input/output operations per second. | [optional] [readonly] 

## Methods

### NewApplianceNodeIopsMetric

`func NewApplianceNodeIopsMetric(classId string, objectType string, ) *ApplianceNodeIopsMetric`

NewApplianceNodeIopsMetric instantiates a new ApplianceNodeIopsMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceNodeIopsMetricWithDefaults

`func NewApplianceNodeIopsMetricWithDefaults() *ApplianceNodeIopsMetric`

NewApplianceNodeIopsMetricWithDefaults instantiates a new ApplianceNodeIopsMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceNodeIopsMetric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceNodeIopsMetric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceNodeIopsMetric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceNodeIopsMetric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceNodeIopsMetric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceNodeIopsMetric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIops

`func (o *ApplianceNodeIopsMetric) GetIops() float64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *ApplianceNodeIopsMetric) GetIopsOk() (*float64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *ApplianceNodeIopsMetric) SetIops(v float64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *ApplianceNodeIopsMetric) HasIops() bool`

HasIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


