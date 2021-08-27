# IqnpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.PoolMember"]
**IqnAddress** | Pointer to **string** | IQN Address of this pool member. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**IqnpoolBlockRelationship**](IqnpoolBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**IqnpoolLeaseRelationship**](IqnpoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolPoolMember

`func NewIqnpoolPoolMember(classId string, objectType string, ) *IqnpoolPoolMember`

NewIqnpoolPoolMember instantiates a new IqnpoolPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolPoolMemberWithDefaults

`func NewIqnpoolPoolMemberWithDefaults() *IqnpoolPoolMember`

NewIqnpoolPoolMemberWithDefaults instantiates a new IqnpoolPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolPoolMember) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolPoolMember) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolPoolMember) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnAddress

`func (o *IqnpoolPoolMember) GetIqnAddress() string`

GetIqnAddress returns the IqnAddress field if non-nil, zero value otherwise.

### GetIqnAddressOk

`func (o *IqnpoolPoolMember) GetIqnAddressOk() (*string, bool)`

GetIqnAddressOk returns a tuple with the IqnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnAddress

`func (o *IqnpoolPoolMember) SetIqnAddress(v string)`

SetIqnAddress sets IqnAddress field to given value.

### HasIqnAddress

`func (o *IqnpoolPoolMember) HasIqnAddress() bool`

HasIqnAddress returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IqnpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IqnpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IqnpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IqnpoolPoolMember) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *IqnpoolPoolMember) GetBlockHead() IqnpoolBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *IqnpoolPoolMember) GetBlockHeadOk() (*IqnpoolBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *IqnpoolPoolMember) SetBlockHead(v IqnpoolBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *IqnpoolPoolMember) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *IqnpoolPoolMember) GetPeer() IqnpoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IqnpoolPoolMember) GetPeerOk() (*IqnpoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IqnpoolPoolMember) SetPeer(v IqnpoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IqnpoolPoolMember) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolPoolMember) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolPoolMember) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolPoolMember) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolPoolMember) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


