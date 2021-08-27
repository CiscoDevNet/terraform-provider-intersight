# UuidpoolUuidLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "uuidpool.UuidLease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "uuidpool.UuidLease"]
**Uuid** | Pointer to **string** | UUID Prefix+Suffix numbers. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**UuidpoolPoolRelationship**](UuidpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**UuidpoolPoolMemberRelationship**](UuidpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**UuidpoolUniverseRelationship**](UuidpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewUuidpoolUuidLease

`func NewUuidpoolUuidLease(classId string, objectType string, ) *UuidpoolUuidLease`

NewUuidpoolUuidLease instantiates a new UuidpoolUuidLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUuidpoolUuidLeaseWithDefaults

`func NewUuidpoolUuidLeaseWithDefaults() *UuidpoolUuidLease`

NewUuidpoolUuidLeaseWithDefaults instantiates a new UuidpoolUuidLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *UuidpoolUuidLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *UuidpoolUuidLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *UuidpoolUuidLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *UuidpoolUuidLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UuidpoolUuidLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UuidpoolUuidLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUuid

`func (o *UuidpoolUuidLease) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UuidpoolUuidLease) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UuidpoolUuidLease) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UuidpoolUuidLease) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *UuidpoolUuidLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *UuidpoolUuidLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *UuidpoolUuidLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *UuidpoolUuidLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *UuidpoolUuidLease) GetPool() UuidpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UuidpoolUuidLease) GetPoolOk() (*UuidpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UuidpoolUuidLease) SetPool(v UuidpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UuidpoolUuidLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *UuidpoolUuidLease) GetPoolMember() UuidpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *UuidpoolUuidLease) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *UuidpoolUuidLease) SetPoolMember(v UuidpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *UuidpoolUuidLease) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *UuidpoolUuidLease) GetUniverse() UuidpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *UuidpoolUuidLease) GetUniverseOk() (*UuidpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *UuidpoolUuidLease) SetUniverse(v UuidpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *UuidpoolUuidLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


