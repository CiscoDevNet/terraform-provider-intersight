# IppoolPoolMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.PoolMember"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.PoolMember"]
**IpType** | Pointer to **string** | Type of the IP address requested. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Address** | Pointer to **string** | IPv4 Address of this pool member. | [optional] 
**IpV6Address** | Pointer to **string** | IPv6 Address of this pool member. | [optional] 
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**IpBlock** | Pointer to [**IppoolShadowBlockRelationship**](IppoolShadowBlockRelationship.md) |  | [optional] 
**Peer** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) |  | [optional] 

## Methods

### NewIppoolPoolMemberAllOf

`func NewIppoolPoolMemberAllOf(classId string, objectType string, ) *IppoolPoolMemberAllOf`

NewIppoolPoolMemberAllOf instantiates a new IppoolPoolMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolPoolMemberAllOfWithDefaults

`func NewIppoolPoolMemberAllOfWithDefaults() *IppoolPoolMemberAllOf`

NewIppoolPoolMemberAllOfWithDefaults instantiates a new IppoolPoolMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolPoolMemberAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolPoolMemberAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolPoolMemberAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolPoolMemberAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolPoolMemberAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolPoolMemberAllOf) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolPoolMemberAllOf) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolPoolMemberAllOf) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolPoolMemberAllOf) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Address

`func (o *IppoolPoolMemberAllOf) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolPoolMemberAllOf) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolPoolMemberAllOf) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolPoolMemberAllOf) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV6Address

`func (o *IppoolPoolMemberAllOf) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *IppoolPoolMemberAllOf) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *IppoolPoolMemberAllOf) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *IppoolPoolMemberAllOf) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolPoolMemberAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetIpBlock

`func (o *IppoolPoolMemberAllOf) GetIpBlock() IppoolShadowBlockRelationship`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *IppoolPoolMemberAllOf) GetIpBlockOk() (*IppoolShadowBlockRelationship, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *IppoolPoolMemberAllOf) SetIpBlock(v IppoolShadowBlockRelationship)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *IppoolPoolMemberAllOf) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetPeer

`func (o *IppoolPoolMemberAllOf) GetPeer() IppoolIpLeaseRelationship`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *IppoolPoolMemberAllOf) GetPeerOk() (*IppoolIpLeaseRelationship, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *IppoolPoolMemberAllOf) SetPeer(v IppoolIpLeaseRelationship)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *IppoolPoolMemberAllOf) HasPeer() bool`

HasPeer returns a boolean if a field has been set.

### GetPool

`func (o *IppoolPoolMemberAllOf) GetPool() IppoolShadowPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolPoolMemberAllOf) GetPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolPoolMemberAllOf) SetPool(v IppoolShadowPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolPoolMemberAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


