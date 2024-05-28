# AccessPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.PolicyInventory"]
**AddressType** | Pointer to [**NullableAccessAddressType**](AccessAddressType.md) |  | [optional] 
**ConfigurationType** | Pointer to [**NullableAccessConfigurationType**](AccessConfigurationType.md) |  | [optional] 
**InbandVlan** | Pointer to **int64** | VLAN to be used for server access over Inband network. | [optional] [readonly] 
**InbandIpPool** | Pointer to [**NullableIppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**InbandVrf** | Pointer to [**NullableVrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**OutOfBandIpPool** | Pointer to [**NullableIppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**OutOfBandVrf** | Pointer to [**NullableVrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewAccessPolicyInventory

`func NewAccessPolicyInventory(classId string, objectType string, ) *AccessPolicyInventory`

NewAccessPolicyInventory instantiates a new AccessPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyInventoryWithDefaults

`func NewAccessPolicyInventoryWithDefaults() *AccessPolicyInventory`

NewAccessPolicyInventoryWithDefaults instantiates a new AccessPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddressType

`func (o *AccessPolicyInventory) GetAddressType() AccessAddressType`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AccessPolicyInventory) GetAddressTypeOk() (*AccessAddressType, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AccessPolicyInventory) SetAddressType(v AccessAddressType)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *AccessPolicyInventory) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### SetAddressTypeNil

`func (o *AccessPolicyInventory) SetAddressTypeNil(b bool)`

 SetAddressTypeNil sets the value for AddressType to be an explicit nil

### UnsetAddressType
`func (o *AccessPolicyInventory) UnsetAddressType()`

UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
### GetConfigurationType

`func (o *AccessPolicyInventory) GetConfigurationType() AccessConfigurationType`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *AccessPolicyInventory) GetConfigurationTypeOk() (*AccessConfigurationType, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *AccessPolicyInventory) SetConfigurationType(v AccessConfigurationType)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *AccessPolicyInventory) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### SetConfigurationTypeNil

`func (o *AccessPolicyInventory) SetConfigurationTypeNil(b bool)`

 SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil

### UnsetConfigurationType
`func (o *AccessPolicyInventory) UnsetConfigurationType()`

UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
### GetInbandVlan

`func (o *AccessPolicyInventory) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *AccessPolicyInventory) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *AccessPolicyInventory) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *AccessPolicyInventory) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInbandIpPool

`func (o *AccessPolicyInventory) GetInbandIpPool() IppoolPoolRelationship`

GetInbandIpPool returns the InbandIpPool field if non-nil, zero value otherwise.

### GetInbandIpPoolOk

`func (o *AccessPolicyInventory) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInbandIpPoolOk returns a tuple with the InbandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpPool

`func (o *AccessPolicyInventory) SetInbandIpPool(v IppoolPoolRelationship)`

SetInbandIpPool sets InbandIpPool field to given value.

### HasInbandIpPool

`func (o *AccessPolicyInventory) HasInbandIpPool() bool`

HasInbandIpPool returns a boolean if a field has been set.

### SetInbandIpPoolNil

`func (o *AccessPolicyInventory) SetInbandIpPoolNil(b bool)`

 SetInbandIpPoolNil sets the value for InbandIpPool to be an explicit nil

### UnsetInbandIpPool
`func (o *AccessPolicyInventory) UnsetInbandIpPool()`

UnsetInbandIpPool ensures that no value is present for InbandIpPool, not even an explicit nil
### GetInbandVrf

`func (o *AccessPolicyInventory) GetInbandVrf() VrfVrfRelationship`

GetInbandVrf returns the InbandVrf field if non-nil, zero value otherwise.

### GetInbandVrfOk

`func (o *AccessPolicyInventory) GetInbandVrfOk() (*VrfVrfRelationship, bool)`

GetInbandVrfOk returns a tuple with the InbandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVrf

`func (o *AccessPolicyInventory) SetInbandVrf(v VrfVrfRelationship)`

SetInbandVrf sets InbandVrf field to given value.

### HasInbandVrf

`func (o *AccessPolicyInventory) HasInbandVrf() bool`

HasInbandVrf returns a boolean if a field has been set.

### SetInbandVrfNil

`func (o *AccessPolicyInventory) SetInbandVrfNil(b bool)`

 SetInbandVrfNil sets the value for InbandVrf to be an explicit nil

### UnsetInbandVrf
`func (o *AccessPolicyInventory) UnsetInbandVrf()`

UnsetInbandVrf ensures that no value is present for InbandVrf, not even an explicit nil
### GetOutOfBandIpPool

`func (o *AccessPolicyInventory) GetOutOfBandIpPool() IppoolPoolRelationship`

GetOutOfBandIpPool returns the OutOfBandIpPool field if non-nil, zero value otherwise.

### GetOutOfBandIpPoolOk

`func (o *AccessPolicyInventory) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpPool

`func (o *AccessPolicyInventory) SetOutOfBandIpPool(v IppoolPoolRelationship)`

SetOutOfBandIpPool sets OutOfBandIpPool field to given value.

### HasOutOfBandIpPool

`func (o *AccessPolicyInventory) HasOutOfBandIpPool() bool`

HasOutOfBandIpPool returns a boolean if a field has been set.

### SetOutOfBandIpPoolNil

`func (o *AccessPolicyInventory) SetOutOfBandIpPoolNil(b bool)`

 SetOutOfBandIpPoolNil sets the value for OutOfBandIpPool to be an explicit nil

### UnsetOutOfBandIpPool
`func (o *AccessPolicyInventory) UnsetOutOfBandIpPool()`

UnsetOutOfBandIpPool ensures that no value is present for OutOfBandIpPool, not even an explicit nil
### GetOutOfBandVrf

`func (o *AccessPolicyInventory) GetOutOfBandVrf() VrfVrfRelationship`

GetOutOfBandVrf returns the OutOfBandVrf field if non-nil, zero value otherwise.

### GetOutOfBandVrfOk

`func (o *AccessPolicyInventory) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool)`

GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandVrf

`func (o *AccessPolicyInventory) SetOutOfBandVrf(v VrfVrfRelationship)`

SetOutOfBandVrf sets OutOfBandVrf field to given value.

### HasOutOfBandVrf

`func (o *AccessPolicyInventory) HasOutOfBandVrf() bool`

HasOutOfBandVrf returns a boolean if a field has been set.

### SetOutOfBandVrfNil

`func (o *AccessPolicyInventory) SetOutOfBandVrfNil(b bool)`

 SetOutOfBandVrfNil sets the value for OutOfBandVrf to be an explicit nil

### UnsetOutOfBandVrf
`func (o *AccessPolicyInventory) UnsetOutOfBandVrf()`

UnsetOutOfBandVrf ensures that no value is present for OutOfBandVrf, not even an explicit nil
### GetTargetMo

`func (o *AccessPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *AccessPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *AccessPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *AccessPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *AccessPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *AccessPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


