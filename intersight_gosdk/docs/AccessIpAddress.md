# AccessIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "access.IpAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "access.IpAddress"]
**Ipv4Address** | Pointer to **string** | IPv4 Address leased to this Server Profile for In-Band Deployment. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | IPv4 Address leased to this Server Profile for In-Band Deployment. | [optional] [readonly] 
**OobIpv4Address** | Pointer to **string** | IPv4 Address leased to this Server Profile for Out-Of-Band deployment. | [optional] [readonly] 
**Ipv4Lease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**Ipv6Lease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**OobIpv4Lease** | Pointer to [**IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) |  | [optional] 
**Profile** | Pointer to [**PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 

## Methods

### NewAccessIpAddress

`func NewAccessIpAddress(classId string, objectType string, ) *AccessIpAddress`

NewAccessIpAddress instantiates a new AccessIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessIpAddressWithDefaults

`func NewAccessIpAddressWithDefaults() *AccessIpAddress`

NewAccessIpAddressWithDefaults instantiates a new AccessIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AccessIpAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AccessIpAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AccessIpAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AccessIpAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessIpAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessIpAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpv4Address

`func (o *AccessIpAddress) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *AccessIpAddress) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *AccessIpAddress) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *AccessIpAddress) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *AccessIpAddress) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *AccessIpAddress) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *AccessIpAddress) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *AccessIpAddress) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetOobIpv4Address

`func (o *AccessIpAddress) GetOobIpv4Address() string`

GetOobIpv4Address returns the OobIpv4Address field if non-nil, zero value otherwise.

### GetOobIpv4AddressOk

`func (o *AccessIpAddress) GetOobIpv4AddressOk() (*string, bool)`

GetOobIpv4AddressOk returns a tuple with the OobIpv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpv4Address

`func (o *AccessIpAddress) SetOobIpv4Address(v string)`

SetOobIpv4Address sets OobIpv4Address field to given value.

### HasOobIpv4Address

`func (o *AccessIpAddress) HasOobIpv4Address() bool`

HasOobIpv4Address returns a boolean if a field has been set.

### GetIpv4Lease

`func (o *AccessIpAddress) GetIpv4Lease() IppoolIpLeaseRelationship`

GetIpv4Lease returns the Ipv4Lease field if non-nil, zero value otherwise.

### GetIpv4LeaseOk

`func (o *AccessIpAddress) GetIpv4LeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpv4LeaseOk returns a tuple with the Ipv4Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Lease

`func (o *AccessIpAddress) SetIpv4Lease(v IppoolIpLeaseRelationship)`

SetIpv4Lease sets Ipv4Lease field to given value.

### HasIpv4Lease

`func (o *AccessIpAddress) HasIpv4Lease() bool`

HasIpv4Lease returns a boolean if a field has been set.

### GetIpv6Lease

`func (o *AccessIpAddress) GetIpv6Lease() IppoolIpLeaseRelationship`

GetIpv6Lease returns the Ipv6Lease field if non-nil, zero value otherwise.

### GetIpv6LeaseOk

`func (o *AccessIpAddress) GetIpv6LeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetIpv6LeaseOk returns a tuple with the Ipv6Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Lease

`func (o *AccessIpAddress) SetIpv6Lease(v IppoolIpLeaseRelationship)`

SetIpv6Lease sets Ipv6Lease field to given value.

### HasIpv6Lease

`func (o *AccessIpAddress) HasIpv6Lease() bool`

HasIpv6Lease returns a boolean if a field has been set.

### GetOobIpv4Lease

`func (o *AccessIpAddress) GetOobIpv4Lease() IppoolIpLeaseRelationship`

GetOobIpv4Lease returns the OobIpv4Lease field if non-nil, zero value otherwise.

### GetOobIpv4LeaseOk

`func (o *AccessIpAddress) GetOobIpv4LeaseOk() (*IppoolIpLeaseRelationship, bool)`

GetOobIpv4LeaseOk returns a tuple with the OobIpv4Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpv4Lease

`func (o *AccessIpAddress) SetOobIpv4Lease(v IppoolIpLeaseRelationship)`

SetOobIpv4Lease sets OobIpv4Lease field to given value.

### HasOobIpv4Lease

`func (o *AccessIpAddress) HasOobIpv4Lease() bool`

HasOobIpv4Lease returns a boolean if a field has been set.

### GetProfile

`func (o *AccessIpAddress) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AccessIpAddress) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AccessIpAddress) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AccessIpAddress) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


