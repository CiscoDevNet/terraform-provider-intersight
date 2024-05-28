# FcpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.PoolMember"]
**WwnId** | Pointer to **string** | WWN ID of this pool member. | [optional] [readonly] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**NullableFcpoolFcBlockRelationship**](FcpoolFcBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**NullableFcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableFcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**Reservation** | Pointer to [**NullableFcpoolReservationRelationship**](FcpoolReservationRelationship.md) |  | [optional] 

## Methods

### NewFcpoolPoolMember

`func NewFcpoolPoolMember(classId string, objectType string, ) *FcpoolPoolMember`

NewFcpoolPoolMember instantiates a new FcpoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolPoolMemberWithDefaults

`func NewFcpoolPoolMemberWithDefaults() *FcpoolPoolMember`

NewFcpoolPoolMemberWithDefaults instantiates a new FcpoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetAssignedToEntityNil

`func (o *FcpoolPoolMember) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *FcpoolPoolMember) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
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

### SetBlockHeadNil

`func (o *FcpoolPoolMember) SetBlockHeadNil(b bool)`

 SetBlockHeadNil sets the value for BlockHead to be an explicit nil

### UnsetBlockHead
`func (o *FcpoolPoolMember) UnsetBlockHead()`

UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
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

### SetPeerNil

`func (o *FcpoolPoolMember) SetPeerNil(b bool)`

 SetPeerNil sets the value for Peer to be an explicit nil

### UnsetPeer
`func (o *FcpoolPoolMember) UnsetPeer()`

UnsetPeer ensures that no value is present for Peer, not even an explicit nil
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

### SetPoolNil

`func (o *FcpoolPoolMember) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *FcpoolPoolMember) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservation

`func (o *FcpoolPoolMember) GetReservation() FcpoolReservationRelationship`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *FcpoolPoolMember) GetReservationOk() (*FcpoolReservationRelationship, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *FcpoolPoolMember) SetReservation(v FcpoolReservationRelationship)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *FcpoolPoolMember) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *FcpoolPoolMember) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *FcpoolPoolMember) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


