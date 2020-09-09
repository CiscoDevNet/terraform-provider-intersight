# NetworkconfigPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateIpv4dnsServer** | Pointer to **string** | IP address of the secondary DNS server. | [optional] 
**AlternateIpv6dnsServer** | Pointer to **string** | IP address of the secondary DNS server. | [optional] 
**DynamicDnsDomain** | Pointer to **string** | The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request. | [optional] 
**EnableDynamicDns** | Pointer to **bool** | If enabled, updates the resource records to the DNS from Cisco IMC. | [optional] 
**EnableIpv4dnsFromDhcp** | Pointer to **bool** | If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it. | [optional] 
**EnableIpv6** | Pointer to **bool** | If enabled, allows to configure IPv6 properties. | [optional] 
**EnableIpv6dnsFromDhcp** | Pointer to **bool** | If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it. | [optional] 
**PreferredIpv4dnsServer** | Pointer to **string** | IP address of the primary DNS server. | [optional] 
**PreferredIpv6dnsServer** | Pointer to **string** | IP address of the primary DNS server. | [optional] 
**ApplianceAccount** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewNetworkconfigPolicyAllOf

`func NewNetworkconfigPolicyAllOf() *NetworkconfigPolicyAllOf`

NewNetworkconfigPolicyAllOf instantiates a new NetworkconfigPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkconfigPolicyAllOfWithDefaults

`func NewNetworkconfigPolicyAllOfWithDefaults() *NetworkconfigPolicyAllOf`

NewNetworkconfigPolicyAllOfWithDefaults instantiates a new NetworkconfigPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) GetAlternateIpv4dnsServer() string`

GetAlternateIpv4dnsServer returns the AlternateIpv4dnsServer field if non-nil, zero value otherwise.

### GetAlternateIpv4dnsServerOk

`func (o *NetworkconfigPolicyAllOf) GetAlternateIpv4dnsServerOk() (*string, bool)`

GetAlternateIpv4dnsServerOk returns a tuple with the AlternateIpv4dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) SetAlternateIpv4dnsServer(v string)`

SetAlternateIpv4dnsServer sets AlternateIpv4dnsServer field to given value.

### HasAlternateIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) HasAlternateIpv4dnsServer() bool`

HasAlternateIpv4dnsServer returns a boolean if a field has been set.

### GetAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) GetAlternateIpv6dnsServer() string`

GetAlternateIpv6dnsServer returns the AlternateIpv6dnsServer field if non-nil, zero value otherwise.

### GetAlternateIpv6dnsServerOk

`func (o *NetworkconfigPolicyAllOf) GetAlternateIpv6dnsServerOk() (*string, bool)`

GetAlternateIpv6dnsServerOk returns a tuple with the AlternateIpv6dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) SetAlternateIpv6dnsServer(v string)`

SetAlternateIpv6dnsServer sets AlternateIpv6dnsServer field to given value.

### HasAlternateIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) HasAlternateIpv6dnsServer() bool`

HasAlternateIpv6dnsServer returns a boolean if a field has been set.

### GetDynamicDnsDomain

`func (o *NetworkconfigPolicyAllOf) GetDynamicDnsDomain() string`

GetDynamicDnsDomain returns the DynamicDnsDomain field if non-nil, zero value otherwise.

### GetDynamicDnsDomainOk

`func (o *NetworkconfigPolicyAllOf) GetDynamicDnsDomainOk() (*string, bool)`

GetDynamicDnsDomainOk returns a tuple with the DynamicDnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicDnsDomain

`func (o *NetworkconfigPolicyAllOf) SetDynamicDnsDomain(v string)`

SetDynamicDnsDomain sets DynamicDnsDomain field to given value.

### HasDynamicDnsDomain

`func (o *NetworkconfigPolicyAllOf) HasDynamicDnsDomain() bool`

HasDynamicDnsDomain returns a boolean if a field has been set.

### GetEnableDynamicDns

`func (o *NetworkconfigPolicyAllOf) GetEnableDynamicDns() bool`

GetEnableDynamicDns returns the EnableDynamicDns field if non-nil, zero value otherwise.

### GetEnableDynamicDnsOk

`func (o *NetworkconfigPolicyAllOf) GetEnableDynamicDnsOk() (*bool, bool)`

GetEnableDynamicDnsOk returns a tuple with the EnableDynamicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDynamicDns

`func (o *NetworkconfigPolicyAllOf) SetEnableDynamicDns(v bool)`

SetEnableDynamicDns sets EnableDynamicDns field to given value.

### HasEnableDynamicDns

`func (o *NetworkconfigPolicyAllOf) HasEnableDynamicDns() bool`

HasEnableDynamicDns returns a boolean if a field has been set.

### GetEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv4dnsFromDhcp() bool`

GetEnableIpv4dnsFromDhcp returns the EnableIpv4dnsFromDhcp field if non-nil, zero value otherwise.

### GetEnableIpv4dnsFromDhcpOk

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv4dnsFromDhcpOk() (*bool, bool)`

GetEnableIpv4dnsFromDhcpOk returns a tuple with the EnableIpv4dnsFromDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) SetEnableIpv4dnsFromDhcp(v bool)`

SetEnableIpv4dnsFromDhcp sets EnableIpv4dnsFromDhcp field to given value.

### HasEnableIpv4dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) HasEnableIpv4dnsFromDhcp() bool`

HasEnableIpv4dnsFromDhcp returns a boolean if a field has been set.

### GetEnableIpv6

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv6() bool`

GetEnableIpv6 returns the EnableIpv6 field if non-nil, zero value otherwise.

### GetEnableIpv6Ok

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv6Ok() (*bool, bool)`

GetEnableIpv6Ok returns a tuple with the EnableIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv6

`func (o *NetworkconfigPolicyAllOf) SetEnableIpv6(v bool)`

SetEnableIpv6 sets EnableIpv6 field to given value.

### HasEnableIpv6

`func (o *NetworkconfigPolicyAllOf) HasEnableIpv6() bool`

HasEnableIpv6 returns a boolean if a field has been set.

### GetEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv6dnsFromDhcp() bool`

GetEnableIpv6dnsFromDhcp returns the EnableIpv6dnsFromDhcp field if non-nil, zero value otherwise.

### GetEnableIpv6dnsFromDhcpOk

`func (o *NetworkconfigPolicyAllOf) GetEnableIpv6dnsFromDhcpOk() (*bool, bool)`

GetEnableIpv6dnsFromDhcpOk returns a tuple with the EnableIpv6dnsFromDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) SetEnableIpv6dnsFromDhcp(v bool)`

SetEnableIpv6dnsFromDhcp sets EnableIpv6dnsFromDhcp field to given value.

### HasEnableIpv6dnsFromDhcp

`func (o *NetworkconfigPolicyAllOf) HasEnableIpv6dnsFromDhcp() bool`

HasEnableIpv6dnsFromDhcp returns a boolean if a field has been set.

### GetPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) GetPreferredIpv4dnsServer() string`

GetPreferredIpv4dnsServer returns the PreferredIpv4dnsServer field if non-nil, zero value otherwise.

### GetPreferredIpv4dnsServerOk

`func (o *NetworkconfigPolicyAllOf) GetPreferredIpv4dnsServerOk() (*string, bool)`

GetPreferredIpv4dnsServerOk returns a tuple with the PreferredIpv4dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) SetPreferredIpv4dnsServer(v string)`

SetPreferredIpv4dnsServer sets PreferredIpv4dnsServer field to given value.

### HasPreferredIpv4dnsServer

`func (o *NetworkconfigPolicyAllOf) HasPreferredIpv4dnsServer() bool`

HasPreferredIpv4dnsServer returns a boolean if a field has been set.

### GetPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) GetPreferredIpv6dnsServer() string`

GetPreferredIpv6dnsServer returns the PreferredIpv6dnsServer field if non-nil, zero value otherwise.

### GetPreferredIpv6dnsServerOk

`func (o *NetworkconfigPolicyAllOf) GetPreferredIpv6dnsServerOk() (*string, bool)`

GetPreferredIpv6dnsServerOk returns a tuple with the PreferredIpv6dnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) SetPreferredIpv6dnsServer(v string)`

SetPreferredIpv6dnsServer sets PreferredIpv6dnsServer field to given value.

### HasPreferredIpv6dnsServer

`func (o *NetworkconfigPolicyAllOf) HasPreferredIpv6dnsServer() bool`

HasPreferredIpv6dnsServer returns a boolean if a field has been set.

### GetApplianceAccount

`func (o *NetworkconfigPolicyAllOf) GetApplianceAccount() IamAccountRelationship`

GetApplianceAccount returns the ApplianceAccount field if non-nil, zero value otherwise.

### GetApplianceAccountOk

`func (o *NetworkconfigPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool)`

GetApplianceAccountOk returns a tuple with the ApplianceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceAccount

`func (o *NetworkconfigPolicyAllOf) SetApplianceAccount(v IamAccountRelationship)`

SetApplianceAccount sets ApplianceAccount field to given value.

### HasApplianceAccount

`func (o *NetworkconfigPolicyAllOf) HasApplianceAccount() bool`

HasApplianceAccount returns a boolean if a field has been set.

### GetOrganization

`func (o *NetworkconfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NetworkconfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NetworkconfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *NetworkconfigPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *NetworkconfigPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *NetworkconfigPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *NetworkconfigPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *NetworkconfigPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *NetworkconfigPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *NetworkconfigPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


