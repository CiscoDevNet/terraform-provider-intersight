# VnicEthQosPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthQosPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthQosPolicyInventory"]
**Burst** | Pointer to **int64** | The burst traffic, in bytes, allowed on the vNIC. | [optional] [readonly] [default to 10240]
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] [readonly] [default to 0]
**Mtu** | Pointer to **int64** | The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. | [optional] [readonly] [default to 1500]
**Priority** | Pointer to **string** | The priortity matching the System QoS specified in the fabric profile. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [readonly] [default to "Best Effort"]
**RateLimit** | Pointer to **int64** | The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. | [optional] [readonly] [default to 0]
**TrustHostCos** | Pointer to **bool** | Enables usage of the Class of Service provided by the operating system. | [optional] [readonly] [default to false]
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicEthQosPolicyInventoryAllOf

`func NewVnicEthQosPolicyInventoryAllOf(classId string, objectType string, ) *VnicEthQosPolicyInventoryAllOf`

NewVnicEthQosPolicyInventoryAllOf instantiates a new VnicEthQosPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthQosPolicyInventoryAllOfWithDefaults

`func NewVnicEthQosPolicyInventoryAllOfWithDefaults() *VnicEthQosPolicyInventoryAllOf`

NewVnicEthQosPolicyInventoryAllOfWithDefaults instantiates a new VnicEthQosPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthQosPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthQosPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthQosPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthQosPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBurst

`func (o *VnicEthQosPolicyInventoryAllOf) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *VnicEthQosPolicyInventoryAllOf) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *VnicEthQosPolicyInventoryAllOf) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetCos

`func (o *VnicEthQosPolicyInventoryAllOf) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicEthQosPolicyInventoryAllOf) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicEthQosPolicyInventoryAllOf) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMtu

`func (o *VnicEthQosPolicyInventoryAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VnicEthQosPolicyInventoryAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VnicEthQosPolicyInventoryAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetPriority

`func (o *VnicEthQosPolicyInventoryAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VnicEthQosPolicyInventoryAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VnicEthQosPolicyInventoryAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRateLimit

`func (o *VnicEthQosPolicyInventoryAllOf) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VnicEthQosPolicyInventoryAllOf) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VnicEthQosPolicyInventoryAllOf) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetTrustHostCos

`func (o *VnicEthQosPolicyInventoryAllOf) GetTrustHostCos() bool`

GetTrustHostCos returns the TrustHostCos field if non-nil, zero value otherwise.

### GetTrustHostCosOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetTrustHostCosOk() (*bool, bool)`

GetTrustHostCosOk returns a tuple with the TrustHostCos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustHostCos

`func (o *VnicEthQosPolicyInventoryAllOf) SetTrustHostCos(v bool)`

SetTrustHostCos sets TrustHostCos field to given value.

### HasTrustHostCos

`func (o *VnicEthQosPolicyInventoryAllOf) HasTrustHostCos() bool`

HasTrustHostCos returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicEthQosPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicEthQosPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicEthQosPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicEthQosPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


