# NetworkconfigPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "networkconfig.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "networkconfig.PolicyInventory"]
**AlternateIpv4dnsServer** | Pointer to **string** | IP address of the secondary DNS server. | [optional] [readonly] 
**AlternateIpv6dnsServer** | Pointer to **string** | IP address of the secondary DNS server. | [optional] [readonly] 
**DynamicDnsDomain** | Pointer to **string** | The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request. | [optional] [readonly] 
**EnableDynamicDns** | Pointer to **bool** | If enabled, updates the resource records to the DNS from Cisco IMC. | [optional] [readonly] 
**EnableIpv4dnsFromDhcp** | Pointer to **bool** | If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it. | [optional] [readonly] 
**EnableIpv6** | Pointer to **bool** | If enabled, allows to configure IPv6 properties. | [optional] [readonly] 
**EnableIpv6dnsFromDhcp** | Pointer to **bool** | If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it. | [optional] [readonly] 
**PreferredIpv4dnsServer** | Pointer to **string** | IP address of the primary DNS server. | [optional] [readonly] 
**PreferredIpv6dnsServer** | Pointer to **string** | IP address of the primary DNS server. | [optional] [readonly] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewNetworkconfigPolicyInventory

`func NewNetworkconfigPolicyInventory(classId string, objectType string, ) *NetworkconfigPolicyInventory`

NewNetworkconfigPolicyInventory instantiates a new NetworkconfigPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkconfigPolicyInventoryWithDefaults

`func NewNetworkconfigPolicyInventoryWithDefaults() *NetworkconfigPolicyInventory`

NewNetworkconfigPolicyInventoryWithDefaults instantiates a new NetworkconfigPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkconfigPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkconfigPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkconfigPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkconfigPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkconfigPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkconfigPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) GetAlternateIpv4dnsServer() string`

GetAlternateIpv4dnsServer returns the AlternateIpv4dnsServer field if non-nil, zero value otherwise.

### GetAlternateIpv4dnsServerOk

`func (o *NetworkconfigPolicyInventory) GetAlternateIpv4dnsServerOk() (*string, bool)`

GetAlternateIpv4dnsServerOk returns a tuple with the AlternateIpv4dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) SetAlternateIpv4dnsServer(v string)`

SetAlternateIpv4dnsServer sets AlternateIpv4dnsServer field to given value.

### HasAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) HasAlternateIpv4dnsServer() bool`

HasAlternateIpv4dnsServer returns a boolean if a field has been set.

### GetAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) GetAlternateIpv6dnsServer() string`

GetAlternateIpv6dnsServer returns the AlternateIpv6dnsServer field if non-nil, zero value otherwise.

### GetAlternateIpv6dnsServerOk

`func (o *NetworkconfigPolicyInventory) GetAlternateIpv6dnsServerOk() (*string, bool)`

GetAlternateIpv6dnsServerOk returns a tuple with the AlternateIpv6dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) SetAlternateIpv6dnsServer(v string)`

SetAlternateIpv6dnsServer sets AlternateIpv6dnsServer field to given value.

### HasAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) HasAlternateIpv6dnsServer() bool`

HasAlternateIpv6dnsServer returns a boolean if a field has been set.

### GetDynamicDnsDomain

`func (o *NetworkconfigPolicyInventory) GetDynamicDnsDomain() string`

GetDynamicDnsDomain returns the DynamicDnsDomain field if non-nil, zero value otherwise.

### GetDynamicDnsDomainOk

`func (o *NetworkconfigPolicyInventory) GetDynamicDnsDomainOk() (*string, bool)`

GetDynamicDnsDomainOk returns a tuple with the DynamicDnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicDnsDomain

`func (o *NetworkconfigPolicyInventory) SetDynamicDnsDomain(v string)`

SetDynamicDnsDomain sets DynamicDnsDomain field to given value.

### HasDynamicDnsDomain

`func (o *NetworkconfigPolicyInventory) HasDynamicDnsDomain() bool`

HasDynamicDnsDomain returns a boolean if a field has been set.

### GetEnableDynamicDns

`func (o *NetworkconfigPolicyInventory) GetEnableDynamicDns() bool`

GetEnableDynamicDns returns the EnableDynamicDns field if non-nil, zero value otherwise.

### GetEnableDynamicDnsOk

`func (o *NetworkconfigPolicyInventory) GetEnableDynamicDnsOk() (*bool, bool)`

GetEnableDynamicDnsOk returns a tuple with the EnableDynamicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDynamicDns

`func (o *NetworkconfigPolicyInventory) SetEnableDynamicDns(v bool)`

SetEnableDynamicDns sets EnableDynamicDns field to given value.

### HasEnableDynamicDns

`func (o *NetworkconfigPolicyInventory) HasEnableDynamicDns() bool`

HasEnableDynamicDns returns a boolean if a field has been set.

### GetEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) GetEnableIpv4dnsFromDhcp() bool`

GetEnableIpv4dnsFromDhcp returns the EnableIpv4dnsFromDhcp field if non-nil, zero value otherwise.

### GetEnableIpv4dnsFromDhcpOk

`func (o *NetworkconfigPolicyInventory) GetEnableIpv4dnsFromDhcpOk() (*bool, bool)`

GetEnableIpv4dnsFromDhcpOk returns a tuple with the EnableIpv4dnsFromDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) SetEnableIpv4dnsFromDhcp(v bool)`

SetEnableIpv4dnsFromDhcp sets EnableIpv4dnsFromDhcp field to given value.

### HasEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) HasEnableIpv4dnsFromDhcp() bool`

HasEnableIpv4dnsFromDhcp returns a boolean if a field has been set.

### GetEnableIpv6

`func (o *NetworkconfigPolicyInventory) GetEnableIpv6() bool`

GetEnableIpv6 returns the EnableIpv6 field if non-nil, zero value otherwise.

### GetEnableIpv6Ok

`func (o *NetworkconfigPolicyInventory) GetEnableIpv6Ok() (*bool, bool)`

GetEnableIpv6Ok returns a tuple with the EnableIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv6

`func (o *NetworkconfigPolicyInventory) SetEnableIpv6(v bool)`

SetEnableIpv6 sets EnableIpv6 field to given value.

### HasEnableIpv6

`func (o *NetworkconfigPolicyInventory) HasEnableIpv6() bool`

HasEnableIpv6 returns a boolean if a field has been set.

### GetEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) GetEnableIpv6dnsFromDhcp() bool`

GetEnableIpv6dnsFromDhcp returns the EnableIpv6dnsFromDhcp field if non-nil, zero value otherwise.

### GetEnableIpv6dnsFromDhcpOk

`func (o *NetworkconfigPolicyInventory) GetEnableIpv6dnsFromDhcpOk() (*bool, bool)`

GetEnableIpv6dnsFromDhcpOk returns a tuple with the EnableIpv6dnsFromDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) SetEnableIpv6dnsFromDhcp(v bool)`

SetEnableIpv6dnsFromDhcp sets EnableIpv6dnsFromDhcp field to given value.

### HasEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyInventory) HasEnableIpv6dnsFromDhcp() bool`

HasEnableIpv6dnsFromDhcp returns a boolean if a field has been set.

### GetPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) GetPreferredIpv4dnsServer() string`

GetPreferredIpv4dnsServer returns the PreferredIpv4dnsServer field if non-nil, zero value otherwise.

### GetPreferredIpv4dnsServerOk

`func (o *NetworkconfigPolicyInventory) GetPreferredIpv4dnsServerOk() (*string, bool)`

GetPreferredIpv4dnsServerOk returns a tuple with the PreferredIpv4dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) SetPreferredIpv4dnsServer(v string)`

SetPreferredIpv4dnsServer sets PreferredIpv4dnsServer field to given value.

### HasPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyInventory) HasPreferredIpv4dnsServer() bool`

HasPreferredIpv4dnsServer returns a boolean if a field has been set.

### GetPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) GetPreferredIpv6dnsServer() string`

GetPreferredIpv6dnsServer returns the PreferredIpv6dnsServer field if non-nil, zero value otherwise.

### GetPreferredIpv6dnsServerOk

`func (o *NetworkconfigPolicyInventory) GetPreferredIpv6dnsServerOk() (*string, bool)`

GetPreferredIpv6dnsServerOk returns a tuple with the PreferredIpv6dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) SetPreferredIpv6dnsServer(v string)`

SetPreferredIpv6dnsServer sets PreferredIpv6dnsServer field to given value.

### HasPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyInventory) HasPreferredIpv6dnsServer() bool`

HasPreferredIpv6dnsServer returns a boolean if a field has been set.

### GetTargetMo

`func (o *NetworkconfigPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *NetworkconfigPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *NetworkconfigPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *NetworkconfigPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


