# IppoolReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.Reservation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.Reservation"]
**Identity** | Pointer to **string** | IP identity to be reserved. | [optional] 
**IpType** | Pointer to **string** | Type of the IP address that needs to be reserved. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**IpV4Address** | Pointer to **string** | IPv4 Address to be reserved. | [optional] [readonly] 
**IpV6Address** | Pointer to **string** | IPv6 Address to be reserved. | [optional] [readonly] 
**VrfMoid** | Pointer to **string** | The moid of the Virtual Routing and Forwarding MO. | [optional] 
**BlockHead** | Pointer to [**IppoolShadowBlockRelationship**](IppoolShadowBlockRelationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Pool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**PoolMember** | Pointer to [**IppoolPoolMemberRelationship**](IppoolPoolMemberRelationship.md) |  | [optional] 
**ShadowPool** | Pointer to [**IppoolShadowPoolRelationship**](IppoolShadowPoolRelationship.md) |  | [optional] 
**Universe** | Pointer to [**IppoolUniverseRelationship**](IppoolUniverseRelationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolReservation

`func NewIppoolReservation(classId string, objectType string, ) *IppoolReservation`

NewIppoolReservation instantiates a new IppoolReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolReservationWithDefaults

`func NewIppoolReservationWithDefaults() *IppoolReservation`

NewIppoolReservationWithDefaults instantiates a new IppoolReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolReservation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolReservation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolReservation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolReservation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolReservation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolReservation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *IppoolReservation) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IppoolReservation) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IppoolReservation) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IppoolReservation) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpType

`func (o *IppoolReservation) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolReservation) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolReservation) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolReservation) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpV4Address

`func (o *IppoolReservation) GetIpV4Address() string`

GetIpV4Address returns the IpV4Address field if non-nil, zero value otherwise.

### GetIpV4AddressOk

`func (o *IppoolReservation) GetIpV4AddressOk() (*string, bool)`

GetIpV4AddressOk returns a tuple with the IpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Address

`func (o *IppoolReservation) SetIpV4Address(v string)`

SetIpV4Address sets IpV4Address field to given value.

### HasIpV4Address

`func (o *IppoolReservation) HasIpV4Address() bool`

HasIpV4Address returns a boolean if a field has been set.

### GetIpV6Address

`func (o *IppoolReservation) GetIpV6Address() string`

GetIpV6Address returns the IpV6Address field if non-nil, zero value otherwise.

### GetIpV6AddressOk

`func (o *IppoolReservation) GetIpV6AddressOk() (*string, bool)`

GetIpV6AddressOk returns a tuple with the IpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Address

`func (o *IppoolReservation) SetIpV6Address(v string)`

SetIpV6Address sets IpV6Address field to given value.

### HasIpV6Address

`func (o *IppoolReservation) HasIpV6Address() bool`

HasIpV6Address returns a boolean if a field has been set.

### GetVrfMoid

`func (o *IppoolReservation) GetVrfMoid() string`

GetVrfMoid returns the VrfMoid field if non-nil, zero value otherwise.

### GetVrfMoidOk

`func (o *IppoolReservation) GetVrfMoidOk() (*string, bool)`

GetVrfMoidOk returns a tuple with the VrfMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfMoid

`func (o *IppoolReservation) SetVrfMoid(v string)`

SetVrfMoid sets VrfMoid field to given value.

### HasVrfMoid

`func (o *IppoolReservation) HasVrfMoid() bool`

HasVrfMoid returns a boolean if a field has been set.

### GetBlockHead

`func (o *IppoolReservation) GetBlockHead() IppoolShadowBlockRelationship`

GetBlockHead returns the BlockHead field if non-nil, zero value otherwise.

### GetBlockHeadOk

`func (o *IppoolReservation) GetBlockHeadOk() (*IppoolShadowBlockRelationship, bool)`

GetBlockHeadOk returns a tuple with the BlockHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHead

`func (o *IppoolReservation) SetBlockHead(v IppoolShadowBlockRelationship)`

SetBlockHead sets BlockHead field to given value.

### HasBlockHead

`func (o *IppoolReservation) HasBlockHead() bool`

HasBlockHead returns a boolean if a field has been set.

### GetOrganization

`func (o *IppoolReservation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IppoolReservation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IppoolReservation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IppoolReservation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPool

`func (o *IppoolReservation) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolReservation) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolReservation) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolReservation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolMember

`func (o *IppoolReservation) GetPoolMember() IppoolPoolMemberRelationship`

GetPoolMember returns the PoolMember field if non-nil, zero value otherwise.

### GetPoolMemberOk

`func (o *IppoolReservation) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool)`

GetPoolMemberOk returns a tuple with the PoolMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMember

`func (o *IppoolReservation) SetPoolMember(v IppoolPoolMemberRelationship)`

SetPoolMember sets PoolMember field to given value.

### HasPoolMember

`func (o *IppoolReservation) HasPoolMember() bool`

HasPoolMember returns a boolean if a field has been set.

### GetShadowPool

`func (o *IppoolReservation) GetShadowPool() IppoolShadowPoolRelationship`

GetShadowPool returns the ShadowPool field if non-nil, zero value otherwise.

### GetShadowPoolOk

`func (o *IppoolReservation) GetShadowPoolOk() (*IppoolShadowPoolRelationship, bool)`

GetShadowPoolOk returns a tuple with the ShadowPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowPool

`func (o *IppoolReservation) SetShadowPool(v IppoolShadowPoolRelationship)`

SetShadowPool sets ShadowPool field to given value.

### HasShadowPool

`func (o *IppoolReservation) HasShadowPool() bool`

HasShadowPool returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolReservation) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolReservation) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolReservation) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolReservation) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolReservation) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolReservation) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolReservation) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolReservation) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


