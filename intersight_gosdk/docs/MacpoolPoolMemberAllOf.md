# MacpoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | MAC Address of this pool member. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**BlockHead** | Pointer to [**MacpoolIdBlockRelationship**](macpool.IdBlock.Relationship.md) |  | [optional] 
**Peer** | Pointer to [**MacpoolLeaseRelationship**](macpool.Lease.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**MacpoolPoolRelationship**](macpool.Pool.Relationship.md) |  | [optional] 

## Methods

### NewMacpoolPoolMemberAllOf

`func NewMacpoolPoolMemberAllOf() *MacpoolPoolMemberAllOf`

NewMacpoolPoolMemberAllOf instantiates a new MacpoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolPoolMemberAllOfWithDefaults

`func NewMacpoolPoolMemberAllOfWithDefaults() *MacpoolPoolMemberAllOf`

NewMacpoolPoolMemberAllOfWithDefaults instantiates a new MacpoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *MacpoolPoolMemberAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MacpoolPoolMemberAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MacpoolPoolMemberAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MacpoolPoolMemberAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *MacpoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *MacpoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *MacpoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *MacpoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *MacpoolPoolMemberAllOf) GetBlockHead() MacpoolIdBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *MacpoolPoolMemberAllOf) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *MacpoolPoolMemberAllOf) SetBlockHead(v MacpoolIdBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *MacpoolPoolMemberAllOf) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *MacpoolPoolMemberAllOf) GetPeer() MacpoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *MacpoolPoolMemberAllOf) GetPeerOk() (*MacpoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *MacpoolPoolMemberAllOf) SetPeer(v MacpoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *MacpoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *MacpoolPoolMemberAllOf) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolPoolMemberAllOf) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolPoolMemberAllOf) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


