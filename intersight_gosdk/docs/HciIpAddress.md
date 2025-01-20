# HciIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.IpAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.IpAddress"]
**Ipv4Address** | Pointer to **string** | An IPv4 address in this IpAddress. | [optional] [readonly] 
**Ipv4PrefixLength** | Pointer to **int32** | The prefix length of the IPv4 address. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | An IPv6 address in this IpAddress. | [optional] [readonly] 
**Ipv6PrefixLength** | Pointer to **int32** | The prefix length of the IPv6 address. | [optional] [readonly] 

## Methods

### NewHciIpAddress

`func NewHciIpAddress(classId string, objectType string, ) *HciIpAddress`

NewHciIpAddress instantiates a new HciIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciIpAddressWithDefaults

`func NewHciIpAddressWithDefaults() *HciIpAddress`

NewHciIpAddressWithDefaults instantiates a new HciIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciIpAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciIpAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciIpAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciIpAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciIpAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciIpAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpv4Address

`func (o *HciIpAddress) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *HciIpAddress) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *HciIpAddress) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *HciIpAddress) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv4PrefixLength

`func (o *HciIpAddress) GetIpv4PrefixLength() int32`

GetIpv4PrefixLength returns the Ipv4PrefixLength field if non-nil, zero value otherwise.

### GetIpv4PrefixLengthOk

`func (o *HciIpAddress) GetIpv4PrefixLengthOk() (*int32, bool)`

GetIpv4PrefixLengthOk returns a tuple with the Ipv4PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4PrefixLength

`func (o *HciIpAddress) SetIpv4PrefixLength(v int32)`

SetIpv4PrefixLength sets Ipv4PrefixLength field to given value.

### HasIpv4PrefixLength

`func (o *HciIpAddress) HasIpv4PrefixLength() bool`

HasIpv4PrefixLength returns a boolean if a field has been set.

### GetIpv6Address

`func (o *HciIpAddress) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *HciIpAddress) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *HciIpAddress) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *HciIpAddress) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetIpv6PrefixLength

`func (o *HciIpAddress) GetIpv6PrefixLength() int32`

GetIpv6PrefixLength returns the Ipv6PrefixLength field if non-nil, zero value otherwise.

### GetIpv6PrefixLengthOk

`func (o *HciIpAddress) GetIpv6PrefixLengthOk() (*int32, bool)`

GetIpv6PrefixLengthOk returns a tuple with the Ipv6PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixLength

`func (o *HciIpAddress) SetIpv6PrefixLength(v int32)`

SetIpv6PrefixLength sets Ipv6PrefixLength field to given value.

### HasIpv6PrefixLength

`func (o *HciIpAddress) HasIpv6PrefixLength() bool`

HasIpv6PrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


