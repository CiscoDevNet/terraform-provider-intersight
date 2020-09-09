# UuidpoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID Prefix+Suffix of this PoolMember. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**BlockHead** | Pointer to [**UuidpoolBlockRelationship**](uuidpool.Block.Relationship.md) |  | [optional] 
**Peer** | Pointer to [**UuidpoolUuidLeaseRelationship**](uuidpool.UuidLease.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](uuidpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewUuidpoolPoolMemberAllOf

`func NewUuidpoolPoolMemberAllOf() *UuidpoolPoolMemberAllOf`

NewUuidpoolPoolMemberAllOf instantiates a new UuidpoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolPoolMemberAllOfWithDefaults

`func NewUuidpoolPoolMemberAllOfWithDefaults() *UuidpoolPoolMemberAllOf`

NewUuidpoolPoolMemberAllOfWithDefaults instantiates a new UuidpoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UuidpoolPoolMemberAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UuidpoolPoolMemberAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UuidpoolPoolMemberAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UuidpoolPoolMemberAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *UuidpoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *UuidpoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *UuidpoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *UuidpoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *UuidpoolPoolMemberAllOf) GetBlockHead() UuidpoolBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *UuidpoolPoolMemberAllOf) GetBlockHeadOk() (*UuidpoolBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *UuidpoolPoolMemberAllOf) SetBlockHead(v UuidpoolBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *UuidpoolPoolMemberAllOf) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *UuidpoolPoolMemberAllOf) GetPeer() UuidpoolUuidLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *UuidpoolPoolMemberAllOf) GetPeerOk() (*UuidpoolUuidLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *UuidpoolPoolMemberAllOf) SetPeer(v UuidpoolUuidLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *UuidpoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolPoolMemberAllOf) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolPoolMemberAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolPoolMemberAllOf) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


