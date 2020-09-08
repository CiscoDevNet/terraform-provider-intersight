# VnicRssHashSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Hash** | Pointer to **bool** | When enabled, the IPv4 address is used for traffic distribution. | [optional] 
**Ipv6ExtHash** | Pointer to **bool** | When enabled, the IPv6 extensions are used for traffic distribution. | [optional] 
**Ipv6Hash** | Pointer to **bool** | When enabled, the IPv6 address is used for traffic distribution. | [optional] 
**TcpIpv4Hash** | Pointer to **bool** | When enabled, both the IPv4 address and TCP port number are used for traffic distribution. | [optional] 
**TcpIpv6ExtHash** | Pointer to **bool** | When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution. | [optional] 
**TcpIpv6Hash** | Pointer to **bool** | When enabled, both the IPv6 address and TCP port number are used for traffic distribution. | [optional] 
**UdpIpv4Hash** | Pointer to **bool** | When enabled, both the IPv4 address and UDP port number are used for traffic distribution. | [optional] 
**UdpIpv6Hash** | Pointer to **bool** | When enabled, both the IPv6 address and UDP port number are used for traffic distribution. | [optional] 

## Methods

### NewVnicRssHashSettings

`func NewVnicRssHashSettings() *VnicRssHashSettings`

NewVnicRssHashSettings instantiates a new VnicRssHashSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicRssHashSettingsWithDefaults

`func NewVnicRssHashSettingsWithDefaults() *VnicRssHashSettings`

NewVnicRssHashSettingsWithDefaults instantiates a new VnicRssHashSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Hash

`func (o *VnicRssHashSettings) GetIpv4Hash() bool`

GetIpv4Hash returns the Ipv4Hash field if non-nil, zero value otherwise.

### GetIpv4HashOk

`func (o *VnicRssHashSettings) GetIpv4HashOk() (*bool, bool)`

GetIpv4HashOk returns a tuple with the Ipv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Hash

`func (o *VnicRssHashSettings) SetIpv4Hash(v bool)`

SetIpv4Hash sets Ipv4Hash field to given value.

### HasIpv4Hash

`func (o *VnicRssHashSettings) HasIpv4Hash() bool`

HasIpv4Hash returns a boolean if a field has been set.

### GetIpv6ExtHash

`func (o *VnicRssHashSettings) GetIpv6ExtHash() bool`

GetIpv6ExtHash returns the Ipv6ExtHash field if non-nil, zero value otherwise.

### GetIpv6ExtHashOk

`func (o *VnicRssHashSettings) GetIpv6ExtHashOk() (*bool, bool)`

GetIpv6ExtHashOk returns a tuple with the Ipv6ExtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6ExtHash

`func (o *VnicRssHashSettings) SetIpv6ExtHash(v bool)`

SetIpv6ExtHash sets Ipv6ExtHash field to given value.

### HasIpv6ExtHash

`func (o *VnicRssHashSettings) HasIpv6ExtHash() bool`

HasIpv6ExtHash returns a boolean if a field has been set.

### GetIpv6Hash

`func (o *VnicRssHashSettings) GetIpv6Hash() bool`

GetIpv6Hash returns the Ipv6Hash field if non-nil, zero value otherwise.

### GetIpv6HashOk

`func (o *VnicRssHashSettings) GetIpv6HashOk() (*bool, bool)`

GetIpv6HashOk returns a tuple with the Ipv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Hash

`func (o *VnicRssHashSettings) SetIpv6Hash(v bool)`

SetIpv6Hash sets Ipv6Hash field to given value.

### HasIpv6Hash

`func (o *VnicRssHashSettings) HasIpv6Hash() bool`

HasIpv6Hash returns a boolean if a field has been set.

### GetTcpIpv4Hash

`func (o *VnicRssHashSettings) GetTcpIpv4Hash() bool`

GetTcpIpv4Hash returns the TcpIpv4Hash field if non-nil, zero value otherwise.

### GetTcpIpv4HashOk

`func (o *VnicRssHashSettings) GetTcpIpv4HashOk() (*bool, bool)`

GetTcpIpv4HashOk returns a tuple with the TcpIpv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv4Hash

`func (o *VnicRssHashSettings) SetTcpIpv4Hash(v bool)`

SetTcpIpv4Hash sets TcpIpv4Hash field to given value.

### HasTcpIpv4Hash

`func (o *VnicRssHashSettings) HasTcpIpv4Hash() bool`

HasTcpIpv4Hash returns a boolean if a field has been set.

### GetTcpIpv6ExtHash

`func (o *VnicRssHashSettings) GetTcpIpv6ExtHash() bool`

GetTcpIpv6ExtHash returns the TcpIpv6ExtHash field if non-nil, zero value otherwise.

### GetTcpIpv6ExtHashOk

`func (o *VnicRssHashSettings) GetTcpIpv6ExtHashOk() (*bool, bool)`

GetTcpIpv6ExtHashOk returns a tuple with the TcpIpv6ExtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv6ExtHash

`func (o *VnicRssHashSettings) SetTcpIpv6ExtHash(v bool)`

SetTcpIpv6ExtHash sets TcpIpv6ExtHash field to given value.

### HasTcpIpv6ExtHash

`func (o *VnicRssHashSettings) HasTcpIpv6ExtHash() bool`

HasTcpIpv6ExtHash returns a boolean if a field has been set.

### GetTcpIpv6Hash

`func (o *VnicRssHashSettings) GetTcpIpv6Hash() bool`

GetTcpIpv6Hash returns the TcpIpv6Hash field if non-nil, zero value otherwise.

### GetTcpIpv6HashOk

`func (o *VnicRssHashSettings) GetTcpIpv6HashOk() (*bool, bool)`

GetTcpIpv6HashOk returns a tuple with the TcpIpv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpIpv6Hash

`func (o *VnicRssHashSettings) SetTcpIpv6Hash(v bool)`

SetTcpIpv6Hash sets TcpIpv6Hash field to given value.

### HasTcpIpv6Hash

`func (o *VnicRssHashSettings) HasTcpIpv6Hash() bool`

HasTcpIpv6Hash returns a boolean if a field has been set.

### GetUdpIpv4Hash

`func (o *VnicRssHashSettings) GetUdpIpv4Hash() bool`

GetUdpIpv4Hash returns the UdpIpv4Hash field if non-nil, zero value otherwise.

### GetUdpIpv4HashOk

`func (o *VnicRssHashSettings) GetUdpIpv4HashOk() (*bool, bool)`

GetUdpIpv4HashOk returns a tuple with the UdpIpv4Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpIpv4Hash

`func (o *VnicRssHashSettings) SetUdpIpv4Hash(v bool)`

SetUdpIpv4Hash sets UdpIpv4Hash field to given value.

### HasUdpIpv4Hash

`func (o *VnicRssHashSettings) HasUdpIpv4Hash() bool`

HasUdpIpv4Hash returns a boolean if a field has been set.

### GetUdpIpv6Hash

`func (o *VnicRssHashSettings) GetUdpIpv6Hash() bool`

GetUdpIpv6Hash returns the UdpIpv6Hash field if non-nil, zero value otherwise.

### GetUdpIpv6HashOk

`func (o *VnicRssHashSettings) GetUdpIpv6HashOk() (*bool, bool)`

GetUdpIpv6HashOk returns a tuple with the UdpIpv6Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpIpv6Hash

`func (o *VnicRssHashSettings) SetUdpIpv6Hash(v bool)`

SetUdpIpv6Hash sets UdpIpv6Hash field to given value.

### HasUdpIpv6Hash

`func (o *VnicRssHashSettings) HasUdpIpv6Hash() bool`

HasUdpIpv6Hash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


