# IppoolIpV4ConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | IP address of the default IPv4 gateway. | [optional] [readonly] 
**Netmask** | Pointer to **string** | A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address. | [optional] [readonly] 
**PrimaryDns** | Pointer to **string** | IP Address of the primary Domain Name System (DNS) server. | [optional] [readonly] 
**SecondaryDns** | Pointer to **string** | IP Address of the secondary Domain Name System (DNS) server. | [optional] [readonly] 

## Methods

### NewIppoolIpV4ConfigAllOf

`func NewIppoolIpV4ConfigAllOf() *IppoolIpV4ConfigAllOf`

NewIppoolIpV4ConfigAllOf instantiates a new IppoolIpV4ConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolIpV4ConfigAllOfWithDefaults

`func NewIppoolIpV4ConfigAllOfWithDefaults() *IppoolIpV4ConfigAllOf`

NewIppoolIpV4ConfigAllOfWithDefaults instantiates a new IppoolIpV4ConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *IppoolIpV4ConfigAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IppoolIpV4ConfigAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IppoolIpV4ConfigAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IppoolIpV4ConfigAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *IppoolIpV4ConfigAllOf) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IppoolIpV4ConfigAllOf) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IppoolIpV4ConfigAllOf) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IppoolIpV4ConfigAllOf) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *IppoolIpV4ConfigAllOf) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *IppoolIpV4ConfigAllOf) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *IppoolIpV4ConfigAllOf) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *IppoolIpV4ConfigAllOf) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *IppoolIpV4ConfigAllOf) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *IppoolIpV4ConfigAllOf) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *IppoolIpV4ConfigAllOf) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *IppoolIpV4ConfigAllOf) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


