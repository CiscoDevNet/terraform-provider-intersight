# AccessPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.PolicyInventory"]
**AddressType** | Pointer to [**NullableAccessAddressType**](AccessAddressType.md) |  | [optional] 
**ConfigurationType** | Pointer to [**NullableAccessConfigurationType**](AccessConfigurationType.md) |  | [optional] 
**InbandVlan** | Pointer to **int64** | VLAN to be used for server access over Inband network. | [optional] [readonly] 
**InbandIpPool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**InbandVrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**OutOfBandIpPool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**OutOfBandVrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewAccessPolicyInventoryAllOf

`func NewAccessPolicyInventoryAllOf(classId string, objectType string, ) *AccessPolicyInventoryAllOf`

NewAccessPolicyInventoryAllOf instantiates a new AccessPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyInventoryAllOfWithDefaults

`func NewAccessPolicyInventoryAllOfWithDefaults() *AccessPolicyInventoryAllOf`

NewAccessPolicyInventoryAllOfWithDefaults instantiates a new AccessPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddressType

`func (o *AccessPolicyInventoryAllOf) GetAddressType() AccessAddressType`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AccessPolicyInventoryAllOf) GetAddressTypeOk() (*AccessAddressType, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AccessPolicyInventoryAllOf) SetAddressType(v AccessAddressType)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *AccessPolicyInventoryAllOf) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### SetAddressTypeNil

`func (o *AccessPolicyInventoryAllOf) SetAddressTypeNil(b bool)`

 SetAddressTypeNil sets the value for AddressType to be an explicit nil

### UnsetAddressType
`func (o *AccessPolicyInventoryAllOf) UnsetAddressType()`

UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
### GetConfigurationType

`func (o *AccessPolicyInventoryAllOf) GetConfigurationType() AccessConfigurationType`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *AccessPolicyInventoryAllOf) GetConfigurationTypeOk() (*AccessConfigurationType, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *AccessPolicyInventoryAllOf) SetConfigurationType(v AccessConfigurationType)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *AccessPolicyInventoryAllOf) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### SetConfigurationTypeNil

`func (o *AccessPolicyInventoryAllOf) SetConfigurationTypeNil(b bool)`

 SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil

### UnsetConfigurationType
`func (o *AccessPolicyInventoryAllOf) UnsetConfigurationType()`

UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
### GetInbandVlan

`func (o *AccessPolicyInventoryAllOf) GetInbandVlan() int64`

GetInbandVlan returns the InbandVlan field if non-nil, zero value otherwise.

### GetInbandVlanOk

`func (o *AccessPolicyInventoryAllOf) GetInbandVlanOk() (*int64, bool)`

GetInbandVlanOk returns a tuple with the InbandVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVlan

`func (o *AccessPolicyInventoryAllOf) SetInbandVlan(v int64)`

SetInbandVlan sets InbandVlan field to given value.

### HasInbandVlan

`func (o *AccessPolicyInventoryAllOf) HasInbandVlan() bool`

HasInbandVlan returns a boolean if a field has been set.

### GetInbandIpPool

`func (o *AccessPolicyInventoryAllOf) GetInbandIpPool() IppoolPoolRelationship`

GetInbandIpPool returns the InbandIpPool field if non-nil, zero value otherwise.

### GetInbandIpPoolOk

`func (o *AccessPolicyInventoryAllOf) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetInbandIpPoolOk returns a tuple with the InbandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpPool

`func (o *AccessPolicyInventoryAllOf) SetInbandIpPool(v IppoolPoolRelationship)`

SetInbandIpPool sets InbandIpPool field to given value.

### HasInbandIpPool

`func (o *AccessPolicyInventoryAllOf) HasInbandIpPool() bool`

HasInbandIpPool returns a boolean if a field has been set.

### GetInbandVrf

`func (o *AccessPolicyInventoryAllOf) GetInbandVrf() VrfVrfRelationship`

GetInbandVrf returns the InbandVrf field if non-nil, zero value otherwise.

### GetInbandVrfOk

`func (o *AccessPolicyInventoryAllOf) GetInbandVrfOk() (*VrfVrfRelationship, bool)`

GetInbandVrfOk returns a tuple with the InbandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandVrf

`func (o *AccessPolicyInventoryAllOf) SetInbandVrf(v VrfVrfRelationship)`

SetInbandVrf sets InbandVrf field to given value.

### HasInbandVrf

`func (o *AccessPolicyInventoryAllOf) HasInbandVrf() bool`

HasInbandVrf returns a boolean if a field has been set.

### GetOutOfBandIpPool

`func (o *AccessPolicyInventoryAllOf) GetOutOfBandIpPool() IppoolPoolRelationship`

GetOutOfBandIpPool returns the OutOfBandIpPool field if non-nil, zero value otherwise.

### GetOutOfBandIpPoolOk

`func (o *AccessPolicyInventoryAllOf) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool)`

GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpPool

`func (o *AccessPolicyInventoryAllOf) SetOutOfBandIpPool(v IppoolPoolRelationship)`

SetOutOfBandIpPool sets OutOfBandIpPool field to given value.

### HasOutOfBandIpPool

`func (o *AccessPolicyInventoryAllOf) HasOutOfBandIpPool() bool`

HasOutOfBandIpPool returns a boolean if a field has been set.

### GetOutOfBandVrf

`func (o *AccessPolicyInventoryAllOf) GetOutOfBandVrf() VrfVrfRelationship`

GetOutOfBandVrf returns the OutOfBandVrf field if non-nil, zero value otherwise.

### GetOutOfBandVrfOk

`func (o *AccessPolicyInventoryAllOf) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool)`

GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandVrf

`func (o *AccessPolicyInventoryAllOf) SetOutOfBandVrf(v VrfVrfRelationship)`

SetOutOfBandVrf sets OutOfBandVrf field to given value.

### HasOutOfBandVrf

`func (o *AccessPolicyInventoryAllOf) HasOutOfBandVrf() bool`

HasOutOfBandVrf returns a boolean if a field has been set.

### GetTargetMo

`func (o *AccessPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *AccessPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *AccessPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *AccessPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


