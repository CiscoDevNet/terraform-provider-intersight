# MacpoolReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.Reservation"]
**Identity** | Pointer to **string** | MAC identity to be reserved. | [optional] 
**MemberOf** | Pointer to [**[]MacpoolMemberOf**](MacpoolMemberOf.md) |  | [optional] 
**BlockHead** | Pointer to [**NullableMacpoolIdBlockRelationship**](MacpoolIdBlockRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableMacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**NullableMacpoolPoolMemberRelationship**](MacpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**NullableMacpoolUniverseRelationship**](MacpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewMacpoolReservation

`func NewMacpoolReservation(classId string, objectType string, ) *MacpoolReservation`

NewMacpoolReservation instantiates a new MacpoolReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolReservationWithDefaults

`func NewMacpoolReservationWithDefaults() *MacpoolReservation`

NewMacpoolReservationWithDefaults instantiates a new MacpoolReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *MacpoolReservation) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MacpoolReservation) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MacpoolReservation) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MacpoolReservation) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMemberOf

`func (o *MacpoolReservation) GetMemberOf() []MacpoolMemberOf`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MacpoolReservation) GetMemberOfOk() (*[]MacpoolMemberOf, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *MacpoolReservation) SetMemberOf(v []MacpoolMemberOf)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *MacpoolReservation) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOfNil

`func (o *MacpoolReservation) SetMemberOfNil(b bool)`

 SetMemberOfNil sets the value for MemberOf to be an explicit nil

### UnsetMemberOf
`func (o *MacpoolReservation) UnsetMemberOf()`

UnsetMemberOf ensures that no value is present for MemberOf, not even an explicit nil
### GetBlockHead

`func (o *MacpoolReservation) GetBlockHead() MacpoolIdBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *MacpoolReservation) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *MacpoolReservation) SetBlockHead(v MacpoolIdBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *MacpoolReservation) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### SetBlockHeadNil

`func (o *MacpoolReservation) SetBlockHeadNil(b bool)`

 SetBlockHeadNil sets the value for BlockHead to be an explicit nil

### UnsetBlockHead
`func (o *MacpoolReservation) UnsetBlockHead()`

UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
### GetOrganization

`func (o *MacpoolReservation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MacpoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MacpoolReservation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MacpoolReservation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MacpoolReservation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MacpoolReservation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPool

`func (o *MacpoolReservation) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolReservation) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolReservation) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolReservation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *MacpoolReservation) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *MacpoolReservation) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPoolMember

`func (o *MacpoolReservation) GetPoolMember() MacpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *MacpoolReservation) GetPoolMemberOk() (*MacpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *MacpoolReservation) SetPoolMember(v MacpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *MacpoolReservation) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### SetPoolMemberNil

`func (o *MacpoolReservation) SetPoolMemberNil(b bool)`

 SetPoolMemberNil sets the value for PoolMember to be an explicit nil

### UnsetPoolMember
`func (o *MacpoolReservation) UnsetPoolMember()`

UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
### GetUniverse

`func (o *MacpoolReservation) GetUniverse() MacpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *MacpoolReservation) GetUniverseOk() (*MacpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *MacpoolReservation) SetUniverse(v MacpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *MacpoolReservation) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### SetUniverseNil

`func (o *MacpoolReservation) SetUniverseNil(b bool)`

 SetUniverseNil sets the value for Universe to be an explicit nil

### UnsetUniverse
`func (o *MacpoolReservation) UnsetUniverse()`

UnsetUniverse ensures that no value is present for Universe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


