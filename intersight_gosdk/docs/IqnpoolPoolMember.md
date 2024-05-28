# IqnpoolPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.PoolMember"]
**IqnAddress** | Pointer to **string** | IQN Address of this pool member. It is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnNumber** | Pointer to **int64** | Number of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnPrefix** | Pointer to **string** | Prefix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnSuffix** | Pointer to **string** | Suffix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**AssignedToEntity** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**NullableIqnpoolBlockRelationship**](IqnpoolBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**NullableIqnpoolLeaseRelationship**](IqnpoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**NullableIqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**Reservation** | Pointer to [**NullableIqnpoolReservationRelationship**](IqnpoolReservationRelationship.md) |  | [optional] 

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

### GetIqnNumber

`func (o *IqnpoolPoolMember) GetIqnNumber() int64`

GetIqnNumber returns the IqnNumber field if non-nil, zero value otherwise.

### GetIqnNumberOk

`func (o *IqnpoolPoolMember) GetIqnNumberOk() (*int64, bool)`

GetIqnNumberOk returns a tuple with the IqnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnNumber

`func (o *IqnpoolPoolMember) SetIqnNumber(v int64)`

SetIqnNumber sets IqnNumber field to given value.

### HasIqnNumber

`func (o *IqnpoolPoolMember) HasIqnNumber() bool`

HasIqnNumber returns a boolean if a field has been set.

### GetIqnPrefix

`func (o *IqnpoolPoolMember) GetIqnPrefix() string`

GetIqnPrefix returns the IqnPrefix field if non-nil, zero value otherwise.

### GetIqnPrefixOk

`func (o *IqnpoolPoolMember) GetIqnPrefixOk() (*string, bool)`

GetIqnPrefixOk returns a tuple with the IqnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPrefix

`func (o *IqnpoolPoolMember) SetIqnPrefix(v string)`

SetIqnPrefix sets IqnPrefix field to given value.

### HasIqnPrefix

`func (o *IqnpoolPoolMember) HasIqnPrefix() bool`

HasIqnPrefix returns a boolean if a field has been set.

### GetIqnSuffix

`func (o *IqnpoolPoolMember) GetIqnSuffix() string`

GetIqnSuffix returns the IqnSuffix field if non-nil, zero value otherwise.

### GetIqnSuffixOk

`func (o *IqnpoolPoolMember) GetIqnSuffixOk() (*string, bool)`

GetIqnSuffixOk returns a tuple with the IqnSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffix

`func (o *IqnpoolPoolMember) SetIqnSuffix(v string)`

SetIqnSuffix sets IqnSuffix field to given value.

### HasIqnSuffix

`func (o *IqnpoolPoolMember) HasIqnSuffix() bool`

HasIqnSuffix returns a boolean if a field has been set.

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

### SetAssignedToEntityNil

`func (o *IqnpoolPoolMember) SetAssignedToEntityNil(b bool)`

 SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil

### UnsetAssignedToEntity
`func (o *IqnpoolPoolMember) UnsetAssignedToEntity()`

UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
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

### SetBlockHeadNil

`func (o *IqnpoolPoolMember) SetBlockHeadNil(b bool)`

 SetBlockHeadNil sets the value for BlockHead to be an explicit nil

### UnsetBlockHead
`func (o *IqnpoolPoolMember) UnsetBlockHead()`

UnsetBlockHead ensures that no value is present for BlockHead, not even an explicit nil
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

### SetPeerNil

`func (o *IqnpoolPoolMember) SetPeerNil(b bool)`

 SetPeerNil sets the value for Peer to be an explicit nil

### UnsetPeer
`func (o *IqnpoolPoolMember) UnsetPeer()`

UnsetPeer ensures that no value is present for Peer, not even an explicit nil
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

### SetPoolNil

`func (o *IqnpoolPoolMember) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *IqnpoolPoolMember) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetReservation

`func (o *IqnpoolPoolMember) GetReservation() IqnpoolReservationRelationship`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *IqnpoolPoolMember) GetReservationOk() (*IqnpoolReservationRelationship, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *IqnpoolPoolMember) SetReservation(v IqnpoolReservationRelationship)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *IqnpoolPoolMember) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservationNil

`func (o *IqnpoolPoolMember) SetReservationNil(b bool)`

 SetReservationNil sets the value for Reservation to be an explicit nil

### UnsetReservation
`func (o *IqnpoolPoolMember) UnsetReservation()`

UnsetReservation ensures that no value is present for Reservation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


