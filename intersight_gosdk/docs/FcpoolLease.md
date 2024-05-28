# FcpoolLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Lease"]
**PoolPurpose** | Pointer to **string** | Purpose of this WWN pool. | [optional] 
**Reservation** | Pointer to [**FcpoolReservationReference**](FcpoolReservationReference.md) |  | [optional] 
**WwnId** | Pointer to **string** | WWN ID allocated for pool based allocation. | [optional] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableFcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**NullableFcpoolPoolMemberRelationship**](FcpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**NullableFcpoolUniverseRelationship**](FcpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewFcpoolLease

`func NewFcpoolLease(classId string, objectType string, ) *FcpoolLease`

NewFcpoolLease instantiates a new FcpoolLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolLeaseWithDefaults

`func NewFcpoolLeaseWithDefaults() *FcpoolLease`

NewFcpoolLeaseWithDefaults instantiates a new FcpoolLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoolPurpose

`func (o *FcpoolLease) GetPoolPurpose() string`

GetPoolPurpose returns the PoolPurpose field if non-nil, zero value otherwise.

### GetPoolPurposeOk

`func (o *FcpoolLease) GetPoolPurposeOk() (*string, bool)`

GetPoolPurposeOk returns a tuple with the PoolPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPurpose

`func (o *FcpoolLease) SetPoolPurpose(v string)`

SetPoolPurpose sets PoolPurpose field to given value.

### HasPoolPurpose

`func (o *FcpoolLease) HasPoolPurpose() bool`

HasPoolPurpose returns a boolean if a field has been set.

### GetReservation

`func (o *FcpoolLease) GetReservation() FcpoolReservationReference`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *FcpoolLease) GetReservationOk() (*FcpoolReservationReference, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *FcpoolLease) SetReservation(v FcpoolReservationReference)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *FcpoolLease) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetWwnId

`func (o *FcpoolLease) GetWwnId() string`

GetWwnId returns the WwnId field if non-nil, zero value otherwise.

### GetWwnIdOk

`func (o *FcpoolLease) GetWwnIdOk() (*string, bool)`

GetWwnIdOk returns a tuple with the WwnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnId

`func (o *FcpoolLease) SetWwnId(v string)`

SetWwnId sets WwnId field to given value.

### HasWwnId

`func (o *FcpoolLease) HasWwnId() bool`

HasWwnId returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *FcpoolLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *FcpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *FcpoolLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *FcpoolLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### SetAssignedToEntityNil

`func (o *FcpoolLease) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *FcpoolLease) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
### GetPool

`func (o *FcpoolLease) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolLease) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolLease) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *FcpoolLease) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *FcpoolLease) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPoolMember

`func (o *FcpoolLease) GetPoolMember() FcpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *FcpoolLease) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *FcpoolLease) SetPoolMember(v FcpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *FcpoolLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### SetPoolMemberNil

`func (o *FcpoolLease) SetPoolMemberNil(b bool)`

 SetPoolMemberNil sets the value for PoolMember to be an explicit nil

### UnsetPoolMember
`func (o *FcpoolLease) UnsetPoolMember()`

UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
### GetUniverse

`func (o *FcpoolLease) GetUniverse() FcpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *FcpoolLease) GetUniverseOk() (*FcpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *FcpoolLease) SetUniverse(v FcpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *FcpoolLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### SetUniverseNil

`func (o *FcpoolLease) SetUniverseNil(b bool)`

 SetUniverseNil sets the value for Universe to be an explicit nil

### UnsetUniverse
`func (o *FcpoolLease) UnsetUniverse()`

UnsetUniverse ensures that no value is present for Universe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


