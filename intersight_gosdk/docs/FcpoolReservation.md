# FcpoolReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Reservation"]
**IdPurpose** | Pointer to **string** | Purpose of this WWN ID. Purpose can be WWPN or WWNN. | [optional] 
**Identity** | Pointer to **string** | WWN ID that needs to be reserved. | [optional] 
**Block** | Pointer to [**NullableFcpoolFcBlockRelationship**](FcpoolFcBlockRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableFcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**NullableFcpoolPoolMemberRelationship**](FcpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**NullableFcpoolUniverseRelationship**](FcpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewFcpoolReservation

`func NewFcpoolReservation(classId string, objectType string, ) *FcpoolReservation`

NewFcpoolReservation instantiates a new FcpoolReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolReservationWithDefaults

`func NewFcpoolReservationWithDefaults() *FcpoolReservation`

NewFcpoolReservationWithDefaults instantiates a new FcpoolReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdPurpose

`func (o *FcpoolReservation) GetIdPurpose() string`

GetIdPurpose returns the IdPurpose field if non-nil, zero value otherwise.

### GetIdPurposeOk

`func (o *FcpoolReservation) GetIdPurposeOk() (*string, bool)`

GetIdPurposeOk returns a tuple with the IdPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPurpose

`func (o *FcpoolReservation) SetIdPurpose(v string)`

SetIdPurpose sets IdPurpose field to given value.

### HasIdPurpose

`func (o *FcpoolReservation) HasIdPurpose() bool`

HasIdPurpose returns a boolean if a field has been set.

### GetIdentity

`func (o *FcpoolReservation) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *FcpoolReservation) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *FcpoolReservation) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *FcpoolReservation) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetBlock

`func (o *FcpoolReservation) GetBlock() FcpoolFcBlockRelationship`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *FcpoolReservation) GetBlockOk() (*FcpoolFcBlockRelationship, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *FcpoolReservation) SetBlock(v FcpoolFcBlockRelationship)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *FcpoolReservation) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### SetBlockNil

`func (o *FcpoolReservation) SetBlockNil(b bool)`

 SetBlockNil sets the value for Block to be an explicit nil

### UnsetBlock
`func (o *FcpoolReservation) UnsetBlock()`

UnsetBlock ensures that no value is present for Block, not even an explicit nil
### GetOrganization

`func (o *FcpoolReservation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FcpoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FcpoolReservation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FcpoolReservation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FcpoolReservation) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FcpoolReservation) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPool

`func (o *FcpoolReservation) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolReservation) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolReservation) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolReservation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *FcpoolReservation) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *FcpoolReservation) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPoolMember

`func (o *FcpoolReservation) GetPoolMember() FcpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *FcpoolReservation) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *FcpoolReservation) SetPoolMember(v FcpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *FcpoolReservation) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### SetPoolMemberNil

`func (o *FcpoolReservation) SetPoolMemberNil(b bool)`

 SetPoolMemberNil sets the value for PoolMember to be an explicit nil

### UnsetPoolMember
`func (o *FcpoolReservation) UnsetPoolMember()`

UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
### GetUniverse

`func (o *FcpoolReservation) GetUniverse() FcpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *FcpoolReservation) GetUniverseOk() (*FcpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *FcpoolReservation) SetUniverse(v FcpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *FcpoolReservation) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### SetUniverseNil

`func (o *FcpoolReservation) SetUniverseNil(b bool)`

 SetUniverseNil sets the value for Universe to be an explicit nil

### UnsetUniverse
`func (o *FcpoolReservation) UnsetUniverse()`

UnsetUniverse ensures that no value is present for Universe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


