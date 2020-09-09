# IamLdapPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseProperties** | Pointer to [**IamLdapBaseProperties**](iam.LdapBaseProperties.md) |  | [optional] 
**DnsParameters** | Pointer to [**IamLdapDnsParameters**](iam.LdapDnsParameters.md) |  | [optional] 
**EnableDns** | Pointer to **bool** | Enables DNS to access LDAP servers. | [optional] 
**Enabled** | Pointer to **bool** | LDAP server performs authentication. | [optional] 
**UserSearchPrecedence** | Pointer to **string** | Search precedence between local user database and LDAP user database. * &#x60;LocalUserDb&#x60; - Precedence is given to local user database while searching. * &#x60;LDAPUserDb&#x60; - Precedence is given to LADP user database while searching. | [optional] [default to "LocalUserDb"]
**Var0Idp** | Pointer to [**IamIdpRelationship**](iam.Idp.Relationship.md) |  | [optional] 
**ApplianceAccount** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Groups** | Pointer to [**[]IamLdapGroupRelationship**](iam.LdapGroup.Relationship.md) | An array of relationships to iamLdapGroup resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](policy.AbstractConfigProfile.Relationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 
**Providers** | Pointer to [**[]IamLdapProviderRelationship**](iam.LdapProvider.Relationship.md) | An array of relationships to iamLdapProvider resources. | [optional] 

## Methods

### NewIamLdapPolicyAllOf

`func NewIamLdapPolicyAllOf() *IamLdapPolicyAllOf`

NewIamLdapPolicyAllOf instantiates a new IamLdapPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapPolicyAllOfWithDefaults

`func NewIamLdapPolicyAllOfWithDefaults() *IamLdapPolicyAllOf`

NewIamLdapPolicyAllOfWithDefaults instantiates a new IamLdapPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseProperties

`func (o *IamLdapPolicyAllOf) GetBaseProperties() IamLdapBaseProperties`

GetBaseProperties returns the BaseProperties field if non-nil, zero value otherwise.

### GetBasePropertiesOk

`func (o *IamLdapPolicyAllOf) GetBasePropertiesOk() (*IamLdapBaseProperties, bool)`

GetBasePropertiesOk returns a tuple with the BaseProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProperties

`func (o *IamLdapPolicyAllOf) SetBaseProperties(v IamLdapBaseProperties)`

SetBaseProperties sets BaseProperties field to given value.

### HasBaseProperties

`func (o *IamLdapPolicyAllOf) HasBaseProperties() bool`

HasBaseProperties returns a boolean if a field has been set.

### GetDnsParameters

`func (o *IamLdapPolicyAllOf) GetDnsParameters() IamLdapDnsParameters`

GetDnsParameters returns the DnsParameters field if non-nil, zero value otherwise.

### GetDnsParametersOk

`func (o *IamLdapPolicyAllOf) GetDnsParametersOk() (*IamLdapDnsParameters, bool)`

GetDnsParametersOk returns a tuple with the DnsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsParameters

`func (o *IamLdapPolicyAllOf) SetDnsParameters(v IamLdapDnsParameters)`

SetDnsParameters sets DnsParameters field to given value.

### HasDnsParameters

`func (o *IamLdapPolicyAllOf) HasDnsParameters() bool`

HasDnsParameters returns a boolean if a field has been set.

### GetEnableDns

`func (o *IamLdapPolicyAllOf) GetEnableDns() bool`

GetEnableDns returns the EnableDns field if non-nil, zero value otherwise.

### GetEnableDnsOk

`func (o *IamLdapPolicyAllOf) GetEnableDnsOk() (*bool, bool)`

GetEnableDnsOk returns a tuple with the EnableDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns

`func (o *IamLdapPolicyAllOf) SetEnableDns(v bool)`

SetEnableDns sets EnableDns field to given value.

### HasEnableDns

`func (o *IamLdapPolicyAllOf) HasEnableDns() bool`

HasEnableDns returns a boolean if a field has been set.

### GetEnabled

`func (o *IamLdapPolicyAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamLdapPolicyAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamLdapPolicyAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamLdapPolicyAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserSearchPrecedence

`func (o *IamLdapPolicyAllOf) GetUserSearchPrecedence() string`

GetUserSearchPrecedence returns the UserSearchPrecedence field if non-nil, zero value otherwise.

### GetUserSearchPrecedenceOk

`func (o *IamLdapPolicyAllOf) GetUserSearchPrecedenceOk() (*string, bool)`

GetUserSearchPrecedenceOk returns a tuple with the UserSearchPrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchPrecedence

`func (o *IamLdapPolicyAllOf) SetUserSearchPrecedence(v string)`

SetUserSearchPrecedence sets UserSearchPrecedence field to given value.

### HasUserSearchPrecedence

`func (o *IamLdapPolicyAllOf) HasUserSearchPrecedence() bool`

HasUserSearchPrecedence returns a boolean if a field has been set.

### GetVar0Idp

`func (o *IamLdapPolicyAllOf) GetVar0Idp() IamIdpRelationship`

GetVar0Idp returns the Var0Idp field if non-nil, zero value otherwise.

### GetVar0IdpOk

`func (o *IamLdapPolicyAllOf) GetVar0IdpOk() (*IamIdpRelationship, bool)`

GetVar0IdpOk returns a tuple with the Var0Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0Idp

`func (o *IamLdapPolicyAllOf) SetVar0Idp(v IamIdpRelationship)`

SetVar0Idp sets Var0Idp field to given value.

### HasVar0Idp

`func (o *IamLdapPolicyAllOf) HasVar0Idp() bool`

HasVar0Idp returns a boolean if a field has been set.

### GetApplianceAccount

`func (o *IamLdapPolicyAllOf) GetApplianceAccount() IamAccountRelationship`

GetApplianceAccount returns the ApplianceAccount field if non-nil, zero value otherwise.

### GetApplianceAccountOk

`func (o *IamLdapPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool)`

GetApplianceAccountOk returns a tuple with the ApplianceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceAccount

`func (o *IamLdapPolicyAllOf) SetApplianceAccount(v IamAccountRelationship)`

SetApplianceAccount sets ApplianceAccount field to given value.

### HasApplianceAccount

`func (o *IamLdapPolicyAllOf) HasApplianceAccount() bool`

HasApplianceAccount returns a boolean if a field has been set.

### GetGroups

`func (o *IamLdapPolicyAllOf) GetGroups() []IamLdapGroupRelationship`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamLdapPolicyAllOf) GetGroupsOk() (*[]IamLdapGroupRelationship, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamLdapPolicyAllOf) SetGroups(v []IamLdapGroupRelationship)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamLdapPolicyAllOf) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IamLdapPolicyAllOf) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IamLdapPolicyAllOf) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOrganization

`func (o *IamLdapPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamLdapPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamLdapPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamLdapPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *IamLdapPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IamLdapPolicyAllOf) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IamLdapPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IamLdapPolicyAllOf) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *IamLdapPolicyAllOf) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *IamLdapPolicyAllOf) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetProviders

`func (o *IamLdapPolicyAllOf) GetProviders() []IamLdapProviderRelationship`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamLdapPolicyAllOf) GetProvidersOk() (*[]IamLdapProviderRelationship, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamLdapPolicyAllOf) SetProviders(v []IamLdapProviderRelationship)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamLdapPolicyAllOf) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *IamLdapPolicyAllOf) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *IamLdapPolicyAllOf) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


