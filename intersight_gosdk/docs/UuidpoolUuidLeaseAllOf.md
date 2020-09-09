# UuidpoolUuidLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID Prefix+Suffix numbers. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](uuidpool.Pool.Relationship.md) |  | [optional] 
**PoolMember** | Pointer to [**UuidpoolPoolMemberRelationship**](uuidpool.PoolMember.Relationship.md) |  | [optional] 
**Universe** | Pointer to [**UuidpoolUniverseRelationship**](uuidpool.Universe.Relationship.md) |  | [optional] 

## Methods

### NewUuidpoolUuidLeaseAllOf

`func NewUuidpoolUuidLeaseAllOf() *UuidpoolUuidLeaseAllOf`

NewUuidpoolUuidLeaseAllOf instantiates a new UuidpoolUuidLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolUuidLeaseAllOfWithDefaults

`func NewUuidpoolUuidLeaseAllOfWithDefaults() *UuidpoolUuidLeaseAllOf`

NewUuidpoolUuidLeaseAllOfWithDefaults instantiates a new UuidpoolUuidLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UuidpoolUuidLeaseAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UuidpoolUuidLeaseAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UuidpoolUuidLeaseAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UuidpoolUuidLeaseAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *UuidpoolUuidLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *UuidpoolUuidLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *UuidpoolUuidLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *UuidpoolUuidLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolUuidLeaseAllOf) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolUuidLeaseAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolUuidLeaseAllOf) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolUuidLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *UuidpoolUuidLeaseAllOf) GetPoolMember() UuidpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *UuidpoolUuidLeaseAllOf) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *UuidpoolUuidLeaseAllOf) SetPoolMember(v UuidpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *UuidpoolUuidLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *UuidpoolUuidLeaseAllOf) GetUniverse() UuidpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *UuidpoolUuidLeaseAllOf) GetUniverseOk() (*UuidpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *UuidpoolUuidLeaseAllOf) SetUniverse(v UuidpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *UuidpoolUuidLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


