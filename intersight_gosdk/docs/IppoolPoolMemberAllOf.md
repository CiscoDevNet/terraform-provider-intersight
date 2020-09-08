# IppoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpV4Address** | Pointer to **string** | IPv4 Address of this pool member. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**IpV4Block** | Pointer to [**IppoolShadowBlockRelationship**](ippool.ShadowBlock.Relationship.md) |  | [optional] 
**Peer** | Pointer to [**IppoolIpLeaseRelationship**](ippool.IpLease.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolShadowPoolRelationship**](ippool.ShadowPool.Relationship.md) |  | [optional] 

## Methods

### NewIppoolPoolMemberAllOf

`func NewIppoolPoolMemberAllOf() *IppoolPoolMemberAllOf`

NewIppoolPoolMemberAllOf instantiates a new IppoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolMemberAllOfWithDefaults

`func NewIppoolPoolMemberAllOfWithDefaults() *IppoolPoolMemberAllOf`

NewIppoolPoolMemberAllOfWithDefaults instantiates a new IppoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpV4Address

`func (o *IppoolPoolMemberAllOf) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolPoolMemberAllOf) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolPoolMemberAllOf) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolPoolMemberAllOf) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetIpV4Block

`func (o *IppoolPoolMemberAllOf) GetIpV4Block() IppoolShadowBlockRelationship`

GetIpV4Block returns the IpV4Block field if non-nil, zero value otherwise.

### GetIpV4BlockOk

`func (o *IppoolPoolMemberAllOf) GetIpV4BlockOk() (*IppoolShadowBlockRelationship, bool)`

GetIpV4BlockOk returns a tuple with the IpV4Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Block

`func (o *IppoolPoolMemberAllOf) SetIpV4Block(v IppoolShadowBlockRelationship)`

SetIpV4Block sets IpV4Block field to given value.

### HasIpV4Block

`func (o *IppoolPoolMemberAllOf) HasIpV4Block() bool`

HasIpV4Block returns a boolean if a field has been set.

### GetPeer

`func (o *IppoolPoolMemberAllOf) GetPeer() IppoolIpLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IppoolPoolMemberAllOf) GetPeerOk() (*IppoolIpLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IppoolPoolMemberAllOf) SetPeer(v IppoolIpLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IppoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *IppoolPoolMemberAllOf) GetPool() IppoolShadowPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolPoolMemberAllOf) GetPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolPoolMemberAllOf) SetPool(v IppoolShadowPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


