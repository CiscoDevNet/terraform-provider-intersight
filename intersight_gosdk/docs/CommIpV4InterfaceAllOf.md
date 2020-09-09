# CommIpV4InterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | The IPv4 address of the default gateway. | [optional] 
**IpAddress** | Pointer to **string** | The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3. | [optional] 
**Netmask** | Pointer to **string** | The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0. | [optional] 

## Methods

### NewCommIpV4InterfaceAllOf

`func NewCommIpV4InterfaceAllOf() *CommIpV4InterfaceAllOf`

NewCommIpV4InterfaceAllOf instantiates a new CommIpV4InterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommIpV4InterfaceAllOfWithDefaults

`func NewCommIpV4InterfaceAllOfWithDefaults() *CommIpV4InterfaceAllOf`

NewCommIpV4InterfaceAllOfWithDefaults instantiates a new CommIpV4InterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CommIpV4InterfaceAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CommIpV4InterfaceAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CommIpV4InterfaceAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CommIpV4InterfaceAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *CommIpV4InterfaceAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CommIpV4InterfaceAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CommIpV4InterfaceAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CommIpV4InterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *CommIpV4InterfaceAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CommIpV4InterfaceAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CommIpV4InterfaceAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CommIpV4InterfaceAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


