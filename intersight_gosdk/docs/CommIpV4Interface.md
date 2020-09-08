# CommIpV4Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | The IPv4 address of the default gateway. | [optional] 
**IpAddress** | Pointer to **string** | The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3. | [optional] 
**Netmask** | Pointer to **string** | The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0. | [optional] 

## Methods

### NewCommIpV4Interface

`func NewCommIpV4Interface() *CommIpV4Interface`

NewCommIpV4Interface instantiates a new CommIpV4Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommIpV4InterfaceWithDefaults

`func NewCommIpV4InterfaceWithDefaults() *CommIpV4Interface`

NewCommIpV4InterfaceWithDefaults instantiates a new CommIpV4Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CommIpV4Interface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CommIpV4Interface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CommIpV4Interface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CommIpV4Interface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *CommIpV4Interface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CommIpV4Interface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CommIpV4Interface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CommIpV4Interface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *CommIpV4Interface) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CommIpV4Interface) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CommIpV4Interface) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CommIpV4Interface) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


