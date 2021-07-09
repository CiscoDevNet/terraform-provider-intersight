# IamLdapPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapPolicy"]
**BaseProperties** | Pointer to [**NullableIamLdapBaseProperties**](iam.LdapBaseProperties.md) |  | [optional] 
**DnsParameters** | Pointer to [**NullableIamLdapDnsParameters**](iam.LdapDnsParameters.md) |  | [optional] 
**EnableDns** | Pointer to **bool** | Enables DNS to access LDAP servers. | [optional] 
**Enabled** | Pointer to **bool** | LDAP server performs authentication. | [optional] [default to true]
**UserSearchPrecedence** | Pointer to **string** | Search precedence between local user database and LDAP user database. * &#x60;LocalUserDb&#x60; - Precedence is given to local user database while searching. * &#x60;LDAPUserDb&#x60; - Precedence is given to LADP user database while searching. | [optional] [default to "LocalUserDb"]
**ApplianceAccount** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Groups** | Pointer to [**[]IamLdapGroupRelationship**](IamLdapGroupRelationship.md) | An array of relationships to iamLdapGroup resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 
**Providers** | Pointer to [**[]IamLdapProviderRelationship**](IamLdapProviderRelationship.md) | An array of relationships to iamLdapProvider resources. | [optional] 

## Methods

### NewIamLdapPolicy

`func NewIamLdapPolicy(classId string, objectType string, ) *IamLdapPolicy`

NewIamLdapPolicy instantiates a new IamLdapPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapPolicyWithDefaults

`func NewIamLdapPolicyWithDefaults() *IamLdapPolicy`

NewIamLdapPolicyWithDefaults instantiates a new IamLdapPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaseProperties

`func (o *IamLdapPolicy) GetBaseProperties() IamLdapBaseProperties`

GetBaseProperties returns the BaseProperties field if non-nil, zero value otherwise.

### GetBasePropertiesOk

`func (o *IamLdapPolicy) GetBasePropertiesOk() (*IamLdapBaseProperties, bool)`

GetBasePropertiesOk returns a tuple with the BaseProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProperties

`func (o *IamLdapPolicy) SetBaseProperties(v IamLdapBaseProperties)`

SetBaseProperties sets BaseProperties field to given value.

### HasBaseProperties

`func (o *IamLdapPolicy) HasBaseProperties() bool`

HasBaseProperties returns a boolean if a field has been set.

### SetBasePropertiesNil

`func (o *IamLdapPolicy) SetBasePropertiesNil(b bool)`

 SetBasePropertiesNil sets the value for BaseProperties to be an explicit nil

### UnsetBaseProperties
`func (o *IamLdapPolicy) UnsetBaseProperties()`

UnsetBaseProperties ensures that no value is present for BaseProperties, not even an explicit nil
### GetDnsParameters

`func (o *IamLdapPolicy) GetDnsParameters() IamLdapDnsParameters`

GetDnsParameters returns the DnsParameters field if non-nil, zero value otherwise.

### GetDnsParametersOk

`func (o *IamLdapPolicy) GetDnsParametersOk() (*IamLdapDnsParameters, bool)`

GetDnsParametersOk returns a tuple with the DnsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsParameters

`func (o *IamLdapPolicy) SetDnsParameters(v IamLdapDnsParameters)`

SetDnsParameters sets DnsParameters field to given value.

### HasDnsParameters

`func (o *IamLdapPolicy) HasDnsParameters() bool`

HasDnsParameters returns a boolean if a field has been set.

### SetDnsParametersNil

`func (o *IamLdapPolicy) SetDnsParametersNil(b bool)`

 SetDnsParametersNil sets the value for DnsParameters to be an explicit nil

### UnsetDnsParameters
`func (o *IamLdapPolicy) UnsetDnsParameters()`

UnsetDnsParameters ensures that no value is present for DnsParameters, not even an explicit nil
### GetEnableDns

`func (o *IamLdapPolicy) GetEnableDns() bool`

GetEnableDns returns the EnableDns field if non-nil, zero value otherwise.

### GetEnableDnsOk

`func (o *IamLdapPolicy) GetEnableDnsOk() (*bool, bool)`

GetEnableDnsOk returns a tuple with the EnableDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns

`func (o *IamLdapPolicy) SetEnableDns(v bool)`

SetEnableDns sets EnableDns field to given value.

### HasEnableDns

`func (o *IamLdapPolicy) HasEnableDns() bool`

HasEnableDns returns a boolean if a field has been set.

### GetEnabled

`func (o *IamLdapPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamLdapPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamLdapPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamLdapPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserSearchPrecedence

`func (o *IamLdapPolicy) GetUserSearchPrecedence() string`

GetUserSearchPrecedence returns the UserSearchPrecedence field if non-nil, zero value otherwise.

### GetUserSearchPrecedenceOk

`func (o *IamLdapPolicy) GetUserSearchPrecedenceOk() (*string, bool)`

GetUserSearchPrecedenceOk returns a tuple with the UserSearchPrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchPrecedence

`func (o *IamLdapPolicy) SetUserSearchPrecedence(v string)`

SetUserSearchPrecedence sets UserSearchPrecedence field to given value.

### HasUserSearchPrecedence

`func (o *IamLdapPolicy) HasUserSearchPrecedence() bool`

HasUserSearchPrecedence returns a boolean if a field has been set.

### GetApplianceAccount

`func (o *IamLdapPolicy) GetApplianceAccount() IamAccountRelationship`

GetApplianceAccount returns the ApplianceAccount field if non-nil, zero value otherwise.

### GetApplianceAccountOk

`func (o *IamLdapPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool)`

GetApplianceAccountOk returns a tuple with the ApplianceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceAccount

`func (o *IamLdapPolicy) SetApplianceAccount(v IamAccountRelationship)`

SetApplianceAccount sets ApplianceAccount field to given value.

### HasApplianceAccount

`func (o *IamLdapPolicy) HasApplianceAccount() bool`

HasApplianceAccount returns a boolean if a field has been set.

### GetGroups

`func (o *IamLdapPolicy) GetGroups() []IamLdapGroupRelationship`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamLdapPolicy) GetGroupsOk() (*[]IamLdapGroupRelationship, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamLdapPolicy) SetGroups(v []IamLdapGroupRelationship)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamLdapPolicy) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IamLdapPolicy) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IamLdapPolicy) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOrganization

`func (o *IamLdapPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamLdapPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamLdapPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamLdapPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *IamLdapPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *IamLdapPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *IamLdapPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *IamLdapPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *IamLdapPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *IamLdapPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil
### GetProviders

`func (o *IamLdapPolicy) GetProviders() []IamLdapProviderRelationship`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamLdapPolicy) GetProvidersOk() (*[]IamLdapProviderRelationship, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamLdapPolicy) SetProviders(v []IamLdapProviderRelationship)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamLdapPolicy) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *IamLdapPolicy) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *IamLdapPolicy) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


