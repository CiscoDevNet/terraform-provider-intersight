# MetricsResourceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metrics.ResourceConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metrics.ResourceConfiguration"]
**Enabled** | Pointer to **bool** | Metric collection is enabled for this resource, when enabled all available metrics are collected from the resource into Intersight. | [optional] 
**Domain** | Pointer to [**NullableAssetTargetRelationship**](AssetTargetRelationship.md) |  | [optional] 
**Resource** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewMetricsResourceConfiguration

`func NewMetricsResourceConfiguration(classId string, objectType string, ) *MetricsResourceConfiguration`

NewMetricsResourceConfiguration instantiates a new MetricsResourceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsResourceConfigurationWithDefaults

`func NewMetricsResourceConfigurationWithDefaults() *MetricsResourceConfiguration`

NewMetricsResourceConfigurationWithDefaults instantiates a new MetricsResourceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsResourceConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsResourceConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsResourceConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsResourceConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsResourceConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsResourceConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *MetricsResourceConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsResourceConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsResourceConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsResourceConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDomain

`func (o *MetricsResourceConfiguration) GetDomain() AssetTargetRelationship`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MetricsResourceConfiguration) GetDomainOk() (*AssetTargetRelationship, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MetricsResourceConfiguration) SetDomain(v AssetTargetRelationship)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *MetricsResourceConfiguration) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *MetricsResourceConfiguration) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *MetricsResourceConfiguration) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetResource

`func (o *MetricsResourceConfiguration) GetResource() MoBaseMoRelationship`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MetricsResourceConfiguration) GetResourceOk() (*MoBaseMoRelationship, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MetricsResourceConfiguration) SetResource(v MoBaseMoRelationship)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MetricsResourceConfiguration) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MetricsResourceConfiguration) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MetricsResourceConfiguration) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


