# FcpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WwnId** | Pointer to **string** | WWN ID of this pool member. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**BlockHead** | Pointer to [**FcpoolFcBlockRelationship**](fcpool.FcBlock.Relationship.md) |  | [optional] 
**Peer** | Pointer to [**FcpoolLeaseRelationship**](fcpool.Lease.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**FcpoolPoolRelationship**](fcpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewFcpoolPoolMember

`func NewFcpoolPoolMember() *FcpoolPoolMember`

NewFcpoolPoolMember instantiates a new FcpoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolPoolMemberWithDefaults

`func NewFcpoolPoolMemberWithDefaults() *FcpoolPoolMember`

NewFcpoolPoolMemberWithDefaults instantiates a new FcpoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWwnId

`func (o *FcpoolPoolMember) GetWwnId() string`

GetWwnId returns the WwnId field if non-nil, zero value otherwise.

### GetWwnIdOk

`func (o *FcpoolPoolMember) GetWwnIdOk() (*string, bool)`

GetWwnIdOk returns a tuple with the WwnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnId

`func (o *FcpoolPoolMember) SetWwnId(v string)`

SetWwnId sets WwnId field to given value.

### HasWwnId

`func (o *FcpoolPoolMember) HasWwnId() bool`

HasWwnId returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *FcpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *FcpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *FcpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *FcpoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *FcpoolPoolMember) GetBlockHead() FcpoolFcBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *FcpoolPoolMember) GetBlockHeadOk() (*FcpoolFcBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *FcpoolPoolMember) SetBlockHead(v FcpoolFcBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *FcpoolPoolMember) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *FcpoolPoolMember) GetPeer() FcpoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *FcpoolPoolMember) GetPeerOk() (*FcpoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *FcpoolPoolMember) SetPeer(v FcpoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *FcpoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *FcpoolPoolMember) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolPoolMember) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolPoolMember) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


