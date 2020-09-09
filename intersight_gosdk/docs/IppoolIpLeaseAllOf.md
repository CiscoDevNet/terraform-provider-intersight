# IppoolIpLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpV4Address** | Pointer to **string** | IPv4 Address given as a lease to an external entity like server profiles. | [optional] [readonly] 
**IpV4Config** | Pointer to [**IppoolIpV4Config**](ippool.IpV4Config.md) |  | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolPoolRelationship**](ippool.Pool.Relationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IppoolPoolMemberRelationship**](ippool.PoolMember.Relationship.md) |  | [optional] 
**Universe** | Pointer to [**IppoolUniverseRelationship**](ippool.Universe.Relationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](vrf.Vrf.Relationship.md) |  | [optional] 

## Methods

### NewIppoolIpLeaseAllOf

`func NewIppoolIpLeaseAllOf() *IppoolIpLeaseAllOf`

NewIppoolIpLeaseAllOf instantiates a new IppoolIpLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpLeaseAllOfWithDefaults

`func NewIppoolIpLeaseAllOfWithDefaults() *IppoolIpLeaseAllOf`

NewIppoolIpLeaseAllOfWithDefaults instantiates a new IppoolIpLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV4Address

`func (o *IppoolIpLeaseAllOf) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolIpLeaseAllOf) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolIpLeaseAllOf) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolIpLeaseAllOf) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV4Config

`func (o *IppoolIpLeaseAllOf) GetIpV4Config() IppoolIpV4Config`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *IppoolIpLeaseAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *IppoolIpLeaseAllOf) SetIpV4Config(v IppoolIpV4Config)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *IppoolIpLeaseAllOf) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolIpLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolIpLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolIpLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolIpLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *IppoolIpLeaseAllOf) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolIpLeaseAllOf) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolIpLeaseAllOf) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolIpLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IppoolIpLeaseAllOf) GetPoolMember() IppoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IppoolIpLeaseAllOf) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IppoolIpLeaseAllOf) SetPoolMember(v IppoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IppoolIpLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolIpLeaseAllOf) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolIpLeaseAllOf) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolIpLeaseAllOf) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolIpLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolIpLeaseAllOf) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolIpLeaseAllOf) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolIpLeaseAllOf) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolIpLeaseAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


