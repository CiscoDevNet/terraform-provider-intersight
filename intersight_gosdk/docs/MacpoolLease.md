# MacpoolLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | MAC address allocated for pool-based allocation. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**Pool** | Pointer to [**MacpoolPoolRelationship**](macpool.Pool.Relationship.md) |  | [optional] 
**PoolMember** | Pointer to [**MacpoolPoolMemberRelationship**](macpool.PoolMember.Relationship.md) |  | [optional] 
**Universe** | Pointer to [**MacpoolUniverseRelationship**](macpool.Universe.Relationship.md) |  | [optional] 

## Methods

### NewMacpoolLease

`func NewMacpoolLease() *MacpoolLease`

NewMacpoolLease instantiates a new MacpoolLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolLeaseWithDefaults

`func NewMacpoolLeaseWithDefaults() *MacpoolLease`

NewMacpoolLeaseWithDefaults instantiates a new MacpoolLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *MacpoolLease) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MacpoolLease) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MacpoolLease) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MacpoolLease) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *MacpoolLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *MacpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *MacpoolLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *MacpoolLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *MacpoolLease) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolLease) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolLease) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *MacpoolLease) GetPoolMember() MacpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *MacpoolLease) GetPoolMemberOk() (*MacpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *MacpoolLease) SetPoolMember(v MacpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *MacpoolLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *MacpoolLease) GetUniverse() MacpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *MacpoolLease) GetUniverseOk() (*MacpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *MacpoolLease) SetUniverse(v MacpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *MacpoolLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


