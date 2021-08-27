# FcpoolLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fcpool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fcpool.Lease"]
**PoolPurpose** | Pointer to **string** | Purpose of this WWN pool. | [optional] 
**WwnId** | Pointer to **string** | WWN ID allocated for pool based allocation. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**FcpoolPoolRelationship**](FcpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**FcpoolPoolMemberRelationship**](FcpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**FcpoolUniverseRelationship**](FcpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewFcpoolLeaseAllOf

`func NewFcpoolLeaseAllOf(classId string, objectType string, ) *FcpoolLeaseAllOf`

NewFcpoolLeaseAllOf instantiates a new FcpoolLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpoolLeaseAllOfWithDefaults

`func NewFcpoolLeaseAllOfWithDefaults() *FcpoolLeaseAllOf`

NewFcpoolLeaseAllOfWithDefaults instantiates a new FcpoolLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FcpoolLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FcpoolLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FcpoolLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FcpoolLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FcpoolLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FcpoolLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPoolPurpose

`func (o *FcpoolLeaseAllOf) GetPoolPurpose() string`

GetPoolPurpose returns the PoolPurpose field if non-nil, zero value otherwise.

### GetPoolPurposeOk

`func (o *FcpoolLeaseAllOf) GetPoolPurposeOk() (*string, bool)`

GetPoolPurposeOk returns a tuple with the PoolPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPurpose

`func (o *FcpoolLeaseAllOf) SetPoolPurpose(v string)`

SetPoolPurpose sets PoolPurpose field to given value.

### HasPoolPurpose

`func (o *FcpoolLeaseAllOf) HasPoolPurpose() bool`

HasPoolPurpose returns a boolean if a field has been set.

### GetWwnId

`func (o *FcpoolLeaseAllOf) GetWwnId() string`

GetWwnId returns the WwnId field if non-nil, zero value otherwise.

### GetWwnIdOk

`func (o *FcpoolLeaseAllOf) GetWwnIdOk() (*string, bool)`

GetWwnIdOk returns a tuple with the WwnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnId

`func (o *FcpoolLeaseAllOf) SetWwnId(v string)`

SetWwnId sets WwnId field to given value.

### HasWwnId

`func (o *FcpoolLeaseAllOf) HasWwnId() bool`

HasWwnId returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *FcpoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *FcpoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *FcpoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *FcpoolLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *FcpoolLeaseAllOf) GetPool() FcpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *FcpoolLeaseAllOf) GetPoolOk() (*FcpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *FcpoolLeaseAllOf) SetPool(v FcpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *FcpoolLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *FcpoolLeaseAllOf) GetPoolMember() FcpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *FcpoolLeaseAllOf) GetPoolMemberOk() (*FcpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *FcpoolLeaseAllOf) SetPoolMember(v FcpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *FcpoolLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *FcpoolLeaseAllOf) GetUniverse() FcpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *FcpoolLeaseAllOf) GetUniverseOk() (*FcpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *FcpoolLeaseAllOf) SetUniverse(v FcpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *FcpoolLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


