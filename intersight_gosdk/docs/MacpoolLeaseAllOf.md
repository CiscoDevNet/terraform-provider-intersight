# MacpoolLeaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "macpool.Lease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "macpool.Lease"]
**MacAddress** | Pointer to **string** | MAC address allocated for pool-based allocation. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**Pool** | Pointer to [**MacpoolPoolRelationship**](MacpoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**MacpoolPoolMemberRelationship**](MacpoolPoolMemberRelationship.md) |  | [optional] 
**Universe** | Pointer to [**MacpoolUniverseRelationship**](MacpoolUniverseRelationship.md) |  | [optional] 

## Methods

### NewMacpoolLeaseAllOf

`func NewMacpoolLeaseAllOf(classId string, objectType string, ) *MacpoolLeaseAllOf`

NewMacpoolLeaseAllOf instantiates a new MacpoolLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacpoolLeaseAllOfWithDefaults

`func NewMacpoolLeaseAllOfWithDefaults() *MacpoolLeaseAllOf`

NewMacpoolLeaseAllOfWithDefaults instantiates a new MacpoolLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MacpoolLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MacpoolLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MacpoolLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MacpoolLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MacpoolLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MacpoolLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMacAddress

`func (o *MacpoolLeaseAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MacpoolLeaseAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MacpoolLeaseAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MacpoolLeaseAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *MacpoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *MacpoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *MacpoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *MacpoolLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetPool

`func (o *MacpoolLeaseAllOf) GetPool() MacpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *MacpoolLeaseAllOf) GetPoolOk() (*MacpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *MacpoolLeaseAllOf) SetPool(v MacpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *MacpoolLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *MacpoolLeaseAllOf) GetPoolMember() MacpoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *MacpoolLeaseAllOf) GetPoolMemberOk() (*MacpoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *MacpoolLeaseAllOf) SetPoolMember(v MacpoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *MacpoolLeaseAllOf) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetUniverse

`func (o *MacpoolLeaseAllOf) GetUniverse() MacpoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *MacpoolLeaseAllOf) GetUniverseOk() (*MacpoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *MacpoolLeaseAllOf) SetUniverse(v MacpoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *MacpoolLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


