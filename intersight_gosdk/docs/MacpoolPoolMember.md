# MacpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.PoolMember"]
**MacAddress** | Pointer to **string** | MAC Address of this pool member. | [optional] [readonly] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**NullableMacpoolIdBlockRelationship**](MacpoolIdBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**NullableMacpoolLeaseRelationship**](MacpoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableMacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 
**Reservation** | Pointer to [**NullableMacpoolReservationRelationship**](MacpoolReservationRelationship.md) |  | [optional] 

## Methods

### NewMacpoolPoolMember

`func NewMacpoolPoolMember(classId string, objectType string, ) *MacpoolPoolMember`

NewMacpoolPoolMember instantiates a new MacpoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolPoolMemberWithDefaults

`func NewMacpoolPoolMemberWithDefaults() *MacpoolPoolMember`

NewMacpoolPoolMemberWithDefaults instantiates a new MacpoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAddress

`func (o *MacpoolPoolMember) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MacpoolPoolMember) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MacpoolPoolMember) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MacpoolPoolMember) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *MacpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *MacpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *MacpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *MacpoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### SetAssignedToEntityNil

`func (o *MacpoolPoolMember) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *MacpoolPoolMember) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
### GetBlockHead

`func (o *MacpoolPoolMember) GetBlockHead() MacpoolIdBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *MacpoolPoolMember) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *MacpoolPoolMember) SetBlockHead(v MacpoolIdBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *MacpoolPoolMember) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### SetBlockHeadNil

`func (o *MacpoolPoolMember) SetBlockHeadNil(b bool)`

 SetBlockHeadNil sets the value for BlockHead to be an explicit nil

### UnsetBlockHead
`func (o *MacpoolPoolMember) UnsetBlockHead()`

UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
### GetPeer

`func (o *MacpoolPoolMember) GetPeer() MacpoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *MacpoolPoolMember) GetPeerOk() (*MacpoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *MacpoolPoolMember) SetPeer(v MacpoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *MacpoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### SetPeerNil

`func (o *MacpoolPoolMember) SetPeerNil(b bool)`

 SetPeerNil sets the value for Peer to be an explicit nil

### UnsetPeer
`func (o *MacpoolPoolMember) UnsetPeer()`

UnsetPeer ensures that no value is present for Peer, not even an explicit nil
### GetPool

`func (o *MacpoolPoolMember) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolPoolMember) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolPoolMember) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *MacpoolPoolMember) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *MacpoolPoolMember) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservation

`func (o *MacpoolPoolMember) GetReservation() MacpoolReservationRelationship`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *MacpoolPoolMember) GetReservationOk() (*MacpoolReservationRelationship, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *MacpoolPoolMember) SetReservation(v MacpoolReservationRelationship)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *MacpoolPoolMember) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *MacpoolPoolMember) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *MacpoolPoolMember) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


