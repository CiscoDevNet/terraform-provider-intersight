# VnicFcQosPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.FcQosPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.FcQosPolicyInventory"]
**Burst** | Pointer to **int64** | The burst traffic, in bytes, allowed on the vHBA. | [optional] [readonly] [default to 10240]
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] [readonly] [default to 3]
**MaxDataFieldSize** | Pointer to **int64** | The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports. | [optional] [readonly] [default to 2112]
**Priority** | Pointer to **string** | The priortity matching the System QoS specified in the fabric profile. * &#x60;Best Effort&#x60; - QoS Priority for Best-effort traffic. * &#x60;FC&#x60; - QoS Priority for FC traffic. * &#x60;Platinum&#x60; - QoS Priority for Platinum traffic. * &#x60;Gold&#x60; - QoS Priority for Gold traffic. * &#x60;Silver&#x60; - QoS Priority for Silver traffic. * &#x60;Bronze&#x60; - QoS Priority for Bronze traffic. | [optional] [readonly] [default to "Best Effort"]
**RateLimit** | Pointer to **int64** | The value in Mbps to use for limiting the data rate on the virtual interface. | [optional] [readonly] [default to 0]
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicFcQosPolicyInventory

`func NewVnicFcQosPolicyInventory(classId string, objectType string, ) *VnicFcQosPolicyInventory`

NewVnicFcQosPolicyInventory instantiates a new VnicFcQosPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicFcQosPolicyInventoryWithDefaults

`func NewVnicFcQosPolicyInventoryWithDefaults() *VnicFcQosPolicyInventory`

NewVnicFcQosPolicyInventoryWithDefaults instantiates a new VnicFcQosPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicFcQosPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicFcQosPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicFcQosPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicFcQosPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicFcQosPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicFcQosPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBurst

`func (o *VnicFcQosPolicyInventory) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *VnicFcQosPolicyInventory) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *VnicFcQosPolicyInventory) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *VnicFcQosPolicyInventory) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetCos

`func (o *VnicFcQosPolicyInventory) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicFcQosPolicyInventory) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicFcQosPolicyInventory) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicFcQosPolicyInventory) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMaxDataFieldSize

`func (o *VnicFcQosPolicyInventory) GetMaxDataFieldSize() int64`

GetMaxDataFieldSize returns the MaxDataFieldSize field if non-nil, zero value otherwise.

### GetMaxDataFieldSizeOk

`func (o *VnicFcQosPolicyInventory) GetMaxDataFieldSizeOk() (*int64, bool)`

GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataFieldSize

`func (o *VnicFcQosPolicyInventory) SetMaxDataFieldSize(v int64)`

SetMaxDataFieldSize sets MaxDataFieldSize field to given value.

### HasMaxDataFieldSize

`func (o *VnicFcQosPolicyInventory) HasMaxDataFieldSize() bool`

HasMaxDataFieldSize returns a boolean if a field has been set.

### GetPriority

`func (o *VnicFcQosPolicyInventory) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VnicFcQosPolicyInventory) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VnicFcQosPolicyInventory) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VnicFcQosPolicyInventory) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRateLimit

`func (o *VnicFcQosPolicyInventory) GetRateLimit() int64`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *VnicFcQosPolicyInventory) GetRateLimitOk() (*int64, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *VnicFcQosPolicyInventory) SetRateLimit(v int64)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *VnicFcQosPolicyInventory) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicFcQosPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicFcQosPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicFcQosPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicFcQosPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *VnicFcQosPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *VnicFcQosPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


