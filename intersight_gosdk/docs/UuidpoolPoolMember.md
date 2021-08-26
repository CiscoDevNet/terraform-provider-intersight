# UuidpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.PoolMember"]
**Uuid** | Pointer to **string** | UUID Prefix+Suffix of this PoolMember. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**UuidpoolBlockRelationship**](UuidpoolBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**UuidpoolUuidLeaseRelationship**](UuidpoolUuidLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewUuidpoolPoolMember

`func NewUuidpoolPoolMember(classId string, objectType string, ) *UuidpoolPoolMember`

NewUuidpoolPoolMember instantiates a new UuidpoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolPoolMemberWithDefaults

`func NewUuidpoolPoolMemberWithDefaults() *UuidpoolPoolMember`

NewUuidpoolPoolMemberWithDefaults instantiates a new UuidpoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUuid

`func (o *UuidpoolPoolMember) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UuidpoolPoolMember) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UuidpoolPoolMember) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UuidpoolPoolMember) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *UuidpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *UuidpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *UuidpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *UuidpoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *UuidpoolPoolMember) GetBlockHead() UuidpoolBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *UuidpoolPoolMember) GetBlockHeadOk() (*UuidpoolBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *UuidpoolPoolMember) SetBlockHead(v UuidpoolBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *UuidpoolPoolMember) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *UuidpoolPoolMember) GetPeer() UuidpoolUuidLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *UuidpoolPoolMember) GetPeerOk() (*UuidpoolUuidLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *UuidpoolPoolMember) SetPeer(v UuidpoolUuidLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *UuidpoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolPoolMember) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolPoolMember) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolPoolMember) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


