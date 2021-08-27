# IppoolIpLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.IpLease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.IpLease"]
**IpType** | Pointer to **string** | Type of the IP address requested. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Address** | Pointer to **string** | IPv4 Address given as a lease to an external entity like server profiles. | [optional] 
**IpV4Config** | Pointer to [**NullableIppoolIpV4Config**](IppoolIpV4Config.md) |  | [optional] 
**IpV6Address** | Pointer to **string** | IPv6 Address given as a lease to an external entity like server profiles. | [optional] 
**IpV6Config** | Pointer to [**NullableIppoolIpV6Config**](IppoolIpV6Config.md) |  | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockLease** | Pointer to [**IppoolBlockLeaseRelationship**](IppoolBlockLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IppoolPoolMemberRelationship**](IppoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**IppoolUniverseRelationship**](IppoolUniverseRelationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolIpLease

`func NewIppoolIpLease(classId string, objectType string, ) *IppoolIpLease`

NewIppoolIpLease instantiates a new IppoolIpLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpLeaseWithDefaults

`func NewIppoolIpLeaseWithDefaults() *IppoolIpLease`

NewIppoolIpLeaseWithDefaults instantiates a new IppoolIpLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolIpLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolIpLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolIpLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolIpLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolIpLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolIpLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolIpLease) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolIpLease) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolIpLease) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolIpLease) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Address

`func (o *IppoolIpLease) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolIpLease) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolIpLease) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolIpLease) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolIpLease) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolIpLease) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolIpLease) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolIpLease) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### SetIpV4ConfigNil

`func (o *IppoolIpLease) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *IppoolIpLease) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
### GetIpV6Address

`func (o *IppoolIpLease) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *IppoolIpLease) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *IppoolIpLease) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *IppoolIpLease) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetIpV6Config

`func (o *IppoolIpLease) GetIpV6Config() IppoolIpV6Config`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *IppoolIpLease) GetIpV6ConfigOk() (*IppoolIpV6Config, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *IppoolIpLease) SetIpV6Config(v IppoolIpV6Config)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *IppoolIpLease) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *IppoolIpLease) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *IppoolIpLease) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
### GetAssignedToEntity

`func (o *IppoolIpLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolIpLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolIpLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolIpLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockLease

`func (o *IppoolIpLease) GetBlockLease() IppoolBlockLeaseRelationship`

GetBlockLease returns the BlockLease field if non-nil, zero value otherwise.

### GetBlockLeaseOk

`func (o *IppoolIpLease) GetBlockLeaseOk() (*IppoolBlockLeaseRelationship, bool)`

GetBlockLeaseOk returns a tuple with the BlockLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLease

`func (o *IppoolIpLease) SetBlockLease(v IppoolBlockLeaseRelationship)`

SetBlockLease sets BlockLease field to given value.

### HasBlockLease

`func (o *IppoolIpLease) HasBlockLease() bool`

HasBlockLease returns a boolean if a field has been set.

### GetPool

`func (o *IppoolIpLease) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolIpLease) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolIpLease) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolIpLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IppoolIpLease) GetPoolMember() IppoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IppoolIpLease) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IppoolIpLease) SetPoolMember(v IppoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IppoolIpLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolIpLease) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolIpLease) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolIpLease) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolIpLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolIpLease) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolIpLease) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolIpLease) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolIpLease) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


