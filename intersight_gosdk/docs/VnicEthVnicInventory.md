# VnicEthVnicInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.EthVnicInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.EthVnicInventory"]
**Cos** | Pointer to **int64** | Class of Service to be associated to the traffic on the virtual interface. | [optional] [readonly] 
**Mtu** | Pointer to **int64** | The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the virtual ethernet interface. | [optional] [readonly] 
**TrustHostCos** | Pointer to **bool** | Enables usage of the Class of Service provided by the operating system. | [optional] [readonly] [default to false]
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewVnicEthVnicInventory

`func NewVnicEthVnicInventory(classId string, objectType string, ) *VnicEthVnicInventory`

NewVnicEthVnicInventory instantiates a new VnicEthVnicInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicEthVnicInventoryWithDefaults

`func NewVnicEthVnicInventoryWithDefaults() *VnicEthVnicInventory`

NewVnicEthVnicInventoryWithDefaults instantiates a new VnicEthVnicInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicEthVnicInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicEthVnicInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicEthVnicInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicEthVnicInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicEthVnicInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicEthVnicInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCos

`func (o *VnicEthVnicInventory) GetCos() int64`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *VnicEthVnicInventory) GetCosOk() (*int64, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *VnicEthVnicInventory) SetCos(v int64)`

SetCos sets Cos field to given value.

### HasCos

`func (o *VnicEthVnicInventory) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetMtu

`func (o *VnicEthVnicInventory) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VnicEthVnicInventory) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VnicEthVnicInventory) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VnicEthVnicInventory) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *VnicEthVnicInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnicEthVnicInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnicEthVnicInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnicEthVnicInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrustHostCos

`func (o *VnicEthVnicInventory) GetTrustHostCos() bool`

GetTrustHostCos returns the TrustHostCos field if non-nil, zero value otherwise.

### GetTrustHostCosOk

`func (o *VnicEthVnicInventory) GetTrustHostCosOk() (*bool, bool)`

GetTrustHostCosOk returns a tuple with the TrustHostCos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustHostCos

`func (o *VnicEthVnicInventory) SetTrustHostCos(v bool)`

SetTrustHostCos sets TrustHostCos field to given value.

### HasTrustHostCos

`func (o *VnicEthVnicInventory) HasTrustHostCos() bool`

HasTrustHostCos returns a boolean if a field has been set.

### GetTargetMo

`func (o *VnicEthVnicInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *VnicEthVnicInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *VnicEthVnicInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *VnicEthVnicInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *VnicEthVnicInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *VnicEthVnicInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


