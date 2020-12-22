# ComputeIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.IpAddress"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.IpAddress"]
**Address** | Pointer to **string** | IP Address to be used for KVM. | [optional] [readonly] 
**Category** | Pointer to **string** | Category of the Kvm IP Address. * &#x60;Equipment&#x60; - Ip Address assigned to an equipment. * &#x60;ServiceProfile&#x60; - Ip Address assigned to a Service Profile. | [optional] [readonly] [default to "Equipment"]
**DefaultGateway** | Pointer to **string** | Default gateway property of KVM IP Address. | [optional] [readonly] 
**Dn** | Pointer to **string** | The distinguished name for this managed object. | [optional] [readonly] 
**HttpPort** | Pointer to **int64** | HTTP port of an IP Address. | [optional] [readonly] [default to 80]
**HttpsPort** | Pointer to **int64** | Secured HTTP port of an IP Address. | [optional] [readonly] [default to 443]
**KvmPort** | Pointer to **int64** | Port number on which the KVM is running and used for connecting to KVM console. | [optional] [readonly] [default to 2068]
**KvmVlan** | Pointer to **int64** | VLAN Identifier of Inband IP Address. | [optional] [readonly] 
**Name** | Pointer to **string** | Name to identify the KVM IP Address. * &#x60;Outband&#x60; - The user assigned Out of band IP Address. * &#x60;Inband&#x60; - The user assigned Inband IP Address. | [optional] [readonly] [default to "Outband"]
**Subnet** | Pointer to **string** | Subnet detail of a KVM IP Address. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the KVM IP Address. * &#x60;MgmtInterface&#x60; - Ip Address of a Management Interface. * &#x60;VnicIpV4StaticAddr&#x60; - Static Ipv4 Address of a Virtual Network Interface. * &#x60;VnicIpV4PooledAddr&#x60; - Ipv4 Address of a Virtual Network Interface from an address pool. * &#x60;VnicIpV4MgmtPooledAddr&#x60; - Ipv4 Address of a Virtual Network Interface from a Managed address pool. * &#x60;VnicIpV6StaticAddr&#x60; - Static Ipv6 Address of a Virtual Network Interface. * &#x60;VnicIpV6MgmtPooledAddr&#x60; - Ipv6 Address of a Virtual Network Interface from a Managed address pool. * &#x60;VnicIpV4ProfDerivedAddr&#x60; - Server Profile derived Ipv4 Address of a Virtual Network Interface. * &#x60;MgmtIpV6ProfDerivedAddr&#x60; - Server Profile derived Ipv6 Address used for accessing server management services. | [optional] [readonly] [default to "MgmtInterface"]

## Methods

### NewComputeIpAddress

`func NewComputeIpAddress(classId string, objectType string, ) *ComputeIpAddress`

NewComputeIpAddress instantiates a new ComputeIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeIpAddressWithDefaults

`func NewComputeIpAddressWithDefaults() *ComputeIpAddress`

NewComputeIpAddressWithDefaults instantiates a new ComputeIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeIpAddress) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeIpAddress) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeIpAddress) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeIpAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeIpAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeIpAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *ComputeIpAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ComputeIpAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ComputeIpAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ComputeIpAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCategory

`func (o *ComputeIpAddress) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ComputeIpAddress) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ComputeIpAddress) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ComputeIpAddress) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *ComputeIpAddress) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *ComputeIpAddress) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *ComputeIpAddress) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *ComputeIpAddress) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetDn

`func (o *ComputeIpAddress) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ComputeIpAddress) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ComputeIpAddress) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ComputeIpAddress) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetHttpPort

`func (o *ComputeIpAddress) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ComputeIpAddress) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ComputeIpAddress) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *ComputeIpAddress) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *ComputeIpAddress) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ComputeIpAddress) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ComputeIpAddress) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *ComputeIpAddress) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetKvmPort

`func (o *ComputeIpAddress) GetKvmPort() int64`

GetKvmPort returns the KvmPort field if non-nil, zero value otherwise.

### GetKvmPortOk

`func (o *ComputeIpAddress) GetKvmPortOk() (*int64, bool)`

GetKvmPortOk returns a tuple with the KvmPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmPort

`func (o *ComputeIpAddress) SetKvmPort(v int64)`

SetKvmPort sets KvmPort field to given value.

### HasKvmPort

`func (o *ComputeIpAddress) HasKvmPort() bool`

HasKvmPort returns a boolean if a field has been set.

### GetKvmVlan

`func (o *ComputeIpAddress) GetKvmVlan() int64`

GetKvmVlan returns the KvmVlan field if non-nil, zero value otherwise.

### GetKvmVlanOk

`func (o *ComputeIpAddress) GetKvmVlanOk() (*int64, bool)`

GetKvmVlanOk returns a tuple with the KvmVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmVlan

`func (o *ComputeIpAddress) SetKvmVlan(v int64)`

SetKvmVlan sets KvmVlan field to given value.

### HasKvmVlan

`func (o *ComputeIpAddress) HasKvmVlan() bool`

HasKvmVlan returns a boolean if a field has been set.

### GetName

`func (o *ComputeIpAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeIpAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeIpAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeIpAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *ComputeIpAddress) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ComputeIpAddress) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ComputeIpAddress) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ComputeIpAddress) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetType

`func (o *ComputeIpAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputeIpAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputeIpAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputeIpAddress) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


