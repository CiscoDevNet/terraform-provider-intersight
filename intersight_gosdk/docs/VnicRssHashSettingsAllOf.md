# VnicRssHashSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.RssHashSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.RssHashSettings"]
**Ipv4Hash** | Pointer to **bool** | When enabled, the IPv4 address is used for traffic distribution. | [optional] [default to true]
**Ipv6ExtHash** | Pointer to **bool** | When enabled, the IPv6 extensions are used for traffic distribution. | [optional] [default to false]
**Ipv6Hash** | Pointer to **bool** | When enabled, the IPv6 address is used for traffic distribution. | [optional] [default to true]
**TcpIpv4Hash** | Pointer to **bool** | When enabled, both the IPv4 address and TCP port number are used for traffic distribution. | [optional] [default to true]
**TcpIpv6ExtHash** | Pointer to **bool** | When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution. | [optional] [default to false]
**TcpIpv6Hash** | Pointer to **bool** | When enabled, both the IPv6 address and TCP port number are used for traffic distribution. | [optional] [default to true]
**UdpIpv4Hash** | Pointer to **bool** | When enabled, both the IPv4 address and UDP port number are used for traffic distribution. | [optional] [default to false]
**UdpIpv6Hash** | Pointer to **bool** | When enabled, both the IPv6 address and UDP port number are used for traffic distribution. | [optional] [default to false]

## Methods

### NewVnicRssHashSettingsAllOf

`func NewVnicRssHashSettingsAllOf(classId string, objectType string, ) *VnicRssHashSettingsAllOf`

NewVnicRssHashSettingsAllOf instantiates a new VnicRssHashSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicRssHashSettingsAllOfWithDefaults

`func NewVnicRssHashSettingsAllOfWithDefaults() *VnicRssHashSettingsAllOf`

NewVnicRssHashSettingsAllOfWithDefaults instantiates a new VnicRssHashSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicRssHashSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicRssHashSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicRssHashSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicRssHashSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicRssHashSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicRssHashSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpv4Hash

`func (o *VnicRssHashSettingsAllOf) GetIpv4Hash() bool`

GetIpv4Hash returns the Ipv4Hash field if non-nil, zero value otherwise.

### GetIpv4HashOk

`func (o *VnicRssHashSettingsAllOf) GetIpv4HashOk() (*bool, bool)`

GetIpv4HashOk returns a tuple with the Ipv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Hash

`func (o *VnicRssHashSettingsAllOf) SetIpv4Hash(v bool)`

SetIpv4Hash sets Ipv4Hash field to given value.

### HasIpv4Hash

`func (o *VnicRssHashSettingsAllOf) HasIpv4Hash() bool`

HasIpv4Hash returns a boolean if a field has been set.

### GetIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) GetIpv6ExtHash() bool`

GetIpv6ExtHash returns the Ipv6ExtHash field if non-nil, zero value otherwise.

### GetIpv6ExtHashOk

`func (o *VnicRssHashSettingsAllOf) GetIpv6ExtHashOk() (*bool, bool)`

GetIpv6ExtHashOk returns a tuple with the Ipv6ExtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) SetIpv6ExtHash(v bool)`

SetIpv6ExtHash sets Ipv6ExtHash field to given value.

### HasIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) HasIpv6ExtHash() bool`

HasIpv6ExtHash returns a boolean if a field has been set.

### GetIpv6Hash

`func (o *VnicRssHashSettingsAllOf) GetIpv6Hash() bool`

GetIpv6Hash returns the Ipv6Hash field if non-nil, zero value otherwise.

### GetIpv6HashOk

`func (o *VnicRssHashSettingsAllOf) GetIpv6HashOk() (*bool, bool)`

GetIpv6HashOk returns a tuple with the Ipv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Hash

`func (o *VnicRssHashSettingsAllOf) SetIpv6Hash(v bool)`

SetIpv6Hash sets Ipv6Hash field to given value.

### HasIpv6Hash

`func (o *VnicRssHashSettingsAllOf) HasIpv6Hash() bool`

HasIpv6Hash returns a boolean if a field has been set.

### GetTcpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv4Hash() bool`

GetTcpIpv4Hash returns the TcpIpv4Hash field if non-nil, zero value otherwise.

### GetTcpIpv4HashOk

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv4HashOk() (*bool, bool)`

GetTcpIpv4HashOk returns a tuple with the TcpIpv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) SetTcpIpv4Hash(v bool)`

SetTcpIpv4Hash sets TcpIpv4Hash field to given value.

### HasTcpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) HasTcpIpv4Hash() bool`

HasTcpIpv4Hash returns a boolean if a field has been set.

### GetTcpIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv6ExtHash() bool`

GetTcpIpv6ExtHash returns the TcpIpv6ExtHash field if non-nil, zero value otherwise.

### GetTcpIpv6ExtHashOk

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv6ExtHashOk() (*bool, bool)`

GetTcpIpv6ExtHashOk returns a tuple with the TcpIpv6ExtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) SetTcpIpv6ExtHash(v bool)`

SetTcpIpv6ExtHash sets TcpIpv6ExtHash field to given value.

### HasTcpIpv6ExtHash

`func (o *VnicRssHashSettingsAllOf) HasTcpIpv6ExtHash() bool`

HasTcpIpv6ExtHash returns a boolean if a field has been set.

### GetTcpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv6Hash() bool`

GetTcpIpv6Hash returns the TcpIpv6Hash field if non-nil, zero value otherwise.

### GetTcpIpv6HashOk

`func (o *VnicRssHashSettingsAllOf) GetTcpIpv6HashOk() (*bool, bool)`

GetTcpIpv6HashOk returns a tuple with the TcpIpv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) SetTcpIpv6Hash(v bool)`

SetTcpIpv6Hash sets TcpIpv6Hash field to given value.

### HasTcpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) HasTcpIpv6Hash() bool`

HasTcpIpv6Hash returns a boolean if a field has been set.

### GetUdpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) GetUdpIpv4Hash() bool`

GetUdpIpv4Hash returns the UdpIpv4Hash field if non-nil, zero value otherwise.

### GetUdpIpv4HashOk

`func (o *VnicRssHashSettingsAllOf) GetUdpIpv4HashOk() (*bool, bool)`

GetUdpIpv4HashOk returns a tuple with the UdpIpv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) SetUdpIpv4Hash(v bool)`

SetUdpIpv4Hash sets UdpIpv4Hash field to given value.

### HasUdpIpv4Hash

`func (o *VnicRssHashSettingsAllOf) HasUdpIpv4Hash() bool`

HasUdpIpv4Hash returns a boolean if a field has been set.

### GetUdpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) GetUdpIpv6Hash() bool`

GetUdpIpv6Hash returns the UdpIpv6Hash field if non-nil, zero value otherwise.

### GetUdpIpv6HashOk

`func (o *VnicRssHashSettingsAllOf) GetUdpIpv6HashOk() (*bool, bool)`

GetUdpIpv6HashOk returns a tuple with the UdpIpv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) SetUdpIpv6Hash(v bool)`

SetUdpIpv6Hash sets UdpIpv6Hash field to given value.

### HasUdpIpv6Hash

`func (o *VnicRssHashSettingsAllOf) HasUdpIpv6Hash() bool`

HasUdpIpv6Hash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


