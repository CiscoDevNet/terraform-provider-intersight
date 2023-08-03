# IqnpoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iqnpool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iqnpool.PoolMember"]
**IqnAddress** | Pointer to **string** | IQN Address of this pool member. It is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnNumber** | Pointer to **int64** | Number of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnPrefix** | Pointer to **string** | Prefix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**IqnSuffix** | Pointer to **string** | Suffix of the IQN address. IQN Address is constructed as &lt;prefix&gt;:&lt;suffix&gt;:&lt;number&gt;. | [optional] [readonly] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**BlockHead** | Pointer to [**IqnpoolBlockRelationship**](IqnpoolBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**IqnpoolLeaseRelationship**](IqnpoolLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IqnpoolPoolRelationship**](IqnpoolPoolRelationship.md) |  | [optional] 
**Reservation** | Pointer to [**IqnpoolReservationRelationship**](IqnpoolReservationRelationship.md) |  | [optional] 

## Methods

### NewIqnpoolPoolMemberAllOf

`func NewIqnpoolPoolMemberAllOf(classId string, objectType string, ) *IqnpoolPoolMemberAllOf`

NewIqnpoolPoolMemberAllOf instantiates a new IqnpoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqnpoolPoolMemberAllOfWithDefaults

`func NewIqnpoolPoolMemberAllOfWithDefaults() *IqnpoolPoolMemberAllOf`

NewIqnpoolPoolMemberAllOfWithDefaults instantiates a new IqnpoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IqnpoolPoolMemberAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IqnpoolPoolMemberAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IqnpoolPoolMemberAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IqnpoolPoolMemberAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IqnpoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IqnpoolPoolMemberAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnAddress

`func (o *IqnpoolPoolMemberAllOf) GetIqnAddress() string`

GetIqnAddress returns the IqnAddress field if non-nil, zero value otherwise.

### GetIqnAddressOk

`func (o *IqnpoolPoolMemberAllOf) GetIqnAddressOk() (*string, bool)`

GetIqnAddressOk returns a tuple with the IqnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnAddress

`func (o *IqnpoolPoolMemberAllOf) SetIqnAddress(v string)`

SetIqnAddress sets IqnAddress field to given value.

### HasIqnAddress

`func (o *IqnpoolPoolMemberAllOf) HasIqnAddress() bool`

HasIqnAddress returns a boolean if a field has been set.

### GetIqnNumber

`func (o *IqnpoolPoolMemberAllOf) GetIqnNumber() int64`

GetIqnNumber returns the IqnNumber field if non-nil, zero value otherwise.

### GetIqnNumberOk

`func (o *IqnpoolPoolMemberAllOf) GetIqnNumberOk() (*int64, bool)`

GetIqnNumberOk returns a tuple with the IqnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnNumber

`func (o *IqnpoolPoolMemberAllOf) SetIqnNumber(v int64)`

SetIqnNumber sets IqnNumber field to given value.

### HasIqnNumber

`func (o *IqnpoolPoolMemberAllOf) HasIqnNumber() bool`

HasIqnNumber returns a boolean if a field has been set.

### GetIqnPrefix

`func (o *IqnpoolPoolMemberAllOf) GetIqnPrefix() string`

GetIqnPrefix returns the IqnPrefix field if non-nil, zero value otherwise.

### GetIqnPrefixOk

`func (o *IqnpoolPoolMemberAllOf) GetIqnPrefixOk() (*string, bool)`

GetIqnPrefixOk returns a tuple with the IqnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnPrefix

`func (o *IqnpoolPoolMemberAllOf) SetIqnPrefix(v string)`

SetIqnPrefix sets IqnPrefix field to given value.

### HasIqnPrefix

`func (o *IqnpoolPoolMemberAllOf) HasIqnPrefix() bool`

HasIqnPrefix returns a boolean if a field has been set.

### GetIqnSuffix

`func (o *IqnpoolPoolMemberAllOf) GetIqnSuffix() string`

GetIqnSuffix returns the IqnSuffix field if non-nil, zero value otherwise.

### GetIqnSuffixOk

`func (o *IqnpoolPoolMemberAllOf) GetIqnSuffixOk() (*string, bool)`

GetIqnSuffixOk returns a tuple with the IqnSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnSuffix

`func (o *IqnpoolPoolMemberAllOf) SetIqnSuffix(v string)`

SetIqnSuffix sets IqnSuffix field to given value.

### HasIqnSuffix

`func (o *IqnpoolPoolMemberAllOf) HasIqnSuffix() bool`

HasIqnSuffix returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IqnpoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IqnpoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IqnpoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IqnpoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetBlockHead

`func (o *IqnpoolPoolMemberAllOf) GetBlockHead() IqnpoolBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *IqnpoolPoolMemberAllOf) GetBlockHeadOk() (*IqnpoolBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *IqnpoolPoolMemberAllOf) SetBlockHead(v IqnpoolBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *IqnpoolPoolMemberAllOf) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetPeer

`func (o *IqnpoolPoolMemberAllOf) GetPeer() IqnpoolLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IqnpoolPoolMemberAllOf) GetPeerOk() (*IqnpoolLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IqnpoolPoolMemberAllOf) SetPeer(v IqnpoolLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IqnpoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *IqnpoolPoolMemberAllOf) GetPool() IqnpoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IqnpoolPoolMemberAllOf) GetPoolOk() (*IqnpoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IqnpoolPoolMemberAllOf) SetPool(v IqnpoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IqnpoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetReservation

`func (o *IqnpoolPoolMemberAllOf) GetReservation() IqnpoolReservationRelationship`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *IqnpoolPoolMemberAllOf) GetReservationOk() (*IqnpoolReservationRelationship, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *IqnpoolPoolMemberAllOf) SetReservation(v IqnpoolReservationRelationship)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *IqnpoolPoolMemberAllOf) HasReservation() bool`

HasReservation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


