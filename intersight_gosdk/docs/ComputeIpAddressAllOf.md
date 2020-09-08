# ComputeIpAddressAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IP Address to be used for KVM. | [optional] [readonly] 
**Category** | Pointer to **string** | Category of the Kvm IP Address. * &#x60;Equipment&#x60; - Ip Address assigned to an equipment. * &#x60;ServiceProfile&#x60; - Ip Address assigned to a Service Profile. | [optional] [readonly] [default to "Equipment"]
**DefaultGateway** | Pointer to **string** | Default gateway property of KVM IP Address. | [optional] [readonly] 
**Dn** | Pointer to **string** | The distinguished name for this managed object. | [optional] [readonly] 
**HttpPort** | Pointer to **int64** | HTTP port of an IP Address. | [optional] [readonly] 
**HttpsPort** | Pointer to **int64** | Secured HTTP port of an IP Address. | [optional] [readonly] 
**KvmPort** | Pointer to **int64** | Port number on which the KVM is running and used for connecting to KVM console. | [optional] [readonly] 
**KvmVlan** | Pointer to **int64** | VLAN Identifier of Inband IP Address. | [optional] [readonly] 
**Name** | Pointer to **string** | Name to identify the KVM IP Address. * &#x60;Outband&#x60; - The user assigned Out of band IP Address. * &#x60;Inband&#x60; - The user assigned Inband IP Address. | [optional] [readonly] [default to "Outband"]
**Subnet** | Pointer to **string** | Subnet detail of a KVM IP Address. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the KVM IP Address. * &#x60;MgmtInterface&#x60; - Ip Address of a Management Interface. * &#x60;VnicIpV4StaticAddr&#x60; - Static Ipv4 Address of a Virtual Network Interface. * &#x60;VnicIpV4PooledAddr&#x60; - Ipv4 Address of a Virtual Network Interface from an address pool. * &#x60;VnicIpV4MgmtPooledAddr&#x60; - Ipv4 Address of a Virtual Network Interface from a Managed address pool. * &#x60;VnicIpV6StaticAddr&#x60; - Static Ipv6 Address of a Virtual Network Interface. * &#x60;VnicIpV6MgmtPooledAddr&#x60; - Ipv6 Address of a Virtual Network Interface from a Managed address pool. * &#x60;VnicIpV4ProfDerivedAddr&#x60; - Server Profile derived Ipv4 Address of a Virtual Network Interface. | [optional] [readonly] [default to "MgmtInterface"]

## Methods

### NewComputeIpAddressAllOf

`func NewComputeIpAddressAllOf() *ComputeIpAddressAllOf`

NewComputeIpAddressAllOf instantiates a new ComputeIpAddressAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeIpAddressAllOfWithDefaults

`func NewComputeIpAddressAllOfWithDefaults() *ComputeIpAddressAllOf`

NewComputeIpAddressAllOfWithDefaults instantiates a new ComputeIpAddressAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ComputeIpAddressAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ComputeIpAddressAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ComputeIpAddressAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ComputeIpAddressAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCategory

`func (o *ComputeIpAddressAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ComputeIpAddressAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ComputeIpAddressAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ComputeIpAddressAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *ComputeIpAddressAllOf) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *ComputeIpAddressAllOf) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *ComputeIpAddressAllOf) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *ComputeIpAddressAllOf) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetDn

`func (o *ComputeIpAddressAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputeIpAddressAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputeIpAddressAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputeIpAddressAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHttpPort

`func (o *ComputeIpAddressAllOf) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ComputeIpAddressAllOf) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ComputeIpAddressAllOf) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *ComputeIpAddressAllOf) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *ComputeIpAddressAllOf) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ComputeIpAddressAllOf) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ComputeIpAddressAllOf) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *ComputeIpAddressAllOf) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetKvmPort

`func (o *ComputeIpAddressAllOf) GetKvmPort() int64`

GetKvmPort returns the KvmPort field if non-nil, zero value otherwise.

### GetKvmPortOk

`func (o *ComputeIpAddressAllOf) GetKvmPortOk() (*int64, bool)`

GetKvmPortOk returns a tuple with the KvmPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmPort

`func (o *ComputeIpAddressAllOf) SetKvmPort(v int64)`

SetKvmPort sets KvmPort field to given value.

### HasKvmPort

`func (o *ComputeIpAddressAllOf) HasKvmPort() bool`

HasKvmPort returns a boolean if a field has been set.

### GetKvmVlan

`func (o *ComputeIpAddressAllOf) GetKvmVlan() int64`

GetKvmVlan returns the KvmVlan field if non-nil, zero value otherwise.

### GetKvmVlanOk

`func (o *ComputeIpAddressAllOf) GetKvmVlanOk() (*int64, bool)`

GetKvmVlanOk returns a tuple with the KvmVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmVlan

`func (o *ComputeIpAddressAllOf) SetKvmVlan(v int64)`

SetKvmVlan sets KvmVlan field to given value.

### HasKvmVlan

`func (o *ComputeIpAddressAllOf) HasKvmVlan() bool`

HasKvmVlan returns a boolean if a field has been set.

### GetName

`func (o *ComputeIpAddressAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeIpAddressAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeIpAddressAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeIpAddressAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *ComputeIpAddressAllOf) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ComputeIpAddressAllOf) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ComputeIpAddressAllOf) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ComputeIpAddressAllOf) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetType

`func (o *ComputeIpAddressAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputeIpAddressAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputeIpAddressAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputeIpAddressAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


