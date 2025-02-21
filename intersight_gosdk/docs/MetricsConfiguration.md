# MetricsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metrics.Configuration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metrics.Configuration"]
**CollectNewDevices** | Pointer to **string** | The behavior of the system when new resources are added, controls whether metric collection are automatically enabled for the new resources. * &#x60;AutoEnable&#x60; - Automatically enable metric collection for new resources, up to the limit of resource collection. * &#x60;Disabled&#x60; - Metrics will not be enabled on new resources, to enable collection requires an explicit user enable. | [optional] [default to "AutoEnable"]
**CollectionGranularity** | Pointer to **string** | The current supported collection granularity by the system, defined as the lowest granularity supported, with the actual granularity per resource determined by the license tier of the resource. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Enables metric collection for the account, if disabled metrics will be stopped for all resources in the account. | [optional] [readonly] 
**Limit** | Pointer to **int64** | The total number of resources that can be enabled for metric collection in this account. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewMetricsConfiguration

`func NewMetricsConfiguration(classId string, objectType string, ) *MetricsConfiguration`

NewMetricsConfiguration instantiates a new MetricsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsConfigurationWithDefaults

`func NewMetricsConfigurationWithDefaults() *MetricsConfiguration`

NewMetricsConfigurationWithDefaults instantiates a new MetricsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetricsConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetricsConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetricsConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetricsConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricsConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricsConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectNewDevices

`func (o *MetricsConfiguration) GetCollectNewDevices() string`

GetCollectNewDevices returns the CollectNewDevices field if non-nil, zero value otherwise.

### GetCollectNewDevicesOk

`func (o *MetricsConfiguration) GetCollectNewDevicesOk() (*string, bool)`

GetCollectNewDevicesOk returns a tuple with the CollectNewDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectNewDevices

`func (o *MetricsConfiguration) SetCollectNewDevices(v string)`

SetCollectNewDevices sets CollectNewDevices field to given value.

### HasCollectNewDevices

`func (o *MetricsConfiguration) HasCollectNewDevices() bool`

HasCollectNewDevices returns a boolean if a field has been set.

### GetCollectionGranularity

`func (o *MetricsConfiguration) GetCollectionGranularity() string`

GetCollectionGranularity returns the CollectionGranularity field if non-nil, zero value otherwise.

### GetCollectionGranularityOk

`func (o *MetricsConfiguration) GetCollectionGranularityOk() (*string, bool)`

GetCollectionGranularityOk returns a tuple with the CollectionGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionGranularity

`func (o *MetricsConfiguration) SetCollectionGranularity(v string)`

SetCollectionGranularity sets CollectionGranularity field to given value.

### HasCollectionGranularity

`func (o *MetricsConfiguration) HasCollectionGranularity() bool`

HasCollectionGranularity returns a boolean if a field has been set.

### GetEnabled

`func (o *MetricsConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MetricsConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MetricsConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MetricsConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLimit

`func (o *MetricsConfiguration) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsConfiguration) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsConfiguration) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MetricsConfiguration) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetAccount

`func (o *MetricsConfiguration) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *MetricsConfiguration) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *MetricsConfiguration) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *MetricsConfiguration) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *MetricsConfiguration) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *MetricsConfiguration) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


