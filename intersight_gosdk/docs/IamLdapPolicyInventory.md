# IamLdapPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapPolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapPolicyInventory"]
**BaseProperties** | Pointer to [**NullableIamLdapBaseProperties**](IamLdapBaseProperties.md) |  | [optional] 
**DnsParameters** | Pointer to [**NullableIamLdapDnsParameters**](IamLdapDnsParameters.md) |  | [optional] 
**EnableDns** | Pointer to **bool** | Enables DNS to access LDAP servers. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | LDAP server performs authentication. | [optional] [readonly] [default to true]
**UserSearchPrecedence** | Pointer to **string** | Search precedence between local user database and LDAP user database. * &#x60;LocalUserDb&#x60; - Precedence is given to local user database while searching. * &#x60;LDAPUserDb&#x60; - Precedence is given to LADP user database while searching. | [optional] [readonly] [default to "LocalUserDb"]
**Groups** | Pointer to [**[]IamLdapGroupInventoryRelationship**](IamLdapGroupInventoryRelationship.md) | An array of relationships to iamLdapGroupInventory resources. | [optional] [readonly] 
**Providers** | Pointer to [**[]IamLdapProviderInventoryRelationship**](IamLdapProviderInventoryRelationship.md) | An array of relationships to iamLdapProviderInventory resources. | [optional] [readonly] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewIamLdapPolicyInventory

`func NewIamLdapPolicyInventory(classId string, objectType string, ) *IamLdapPolicyInventory`

NewIamLdapPolicyInventory instantiates a new IamLdapPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapPolicyInventoryWithDefaults

`func NewIamLdapPolicyInventoryWithDefaults() *IamLdapPolicyInventory`

NewIamLdapPolicyInventoryWithDefaults instantiates a new IamLdapPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBaseProperties

`func (o *IamLdapPolicyInventory) GetBaseProperties() IamLdapBaseProperties`

GetBaseProperties returns the BaseProperties field if non-nil, zero value otherwise.

### GetBasePropertiesOk

`func (o *IamLdapPolicyInventory) GetBasePropertiesOk() (*IamLdapBaseProperties, bool)`

GetBasePropertiesOk returns a tuple with the BaseProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProperties

`func (o *IamLdapPolicyInventory) SetBaseProperties(v IamLdapBaseProperties)`

SetBaseProperties sets BaseProperties field to given value.

### HasBaseProperties

`func (o *IamLdapPolicyInventory) HasBaseProperties() bool`

HasBaseProperties returns a boolean if a field has been set.

### SetBasePropertiesNil

`func (o *IamLdapPolicyInventory) SetBasePropertiesNil(b bool)`

 SetBasePropertiesNil sets the value for BaseProperties to be an explicit nil

### UnsetBaseProperties
`func (o *IamLdapPolicyInventory) UnsetBaseProperties()`

UnsetBaseProperties ensures that no value is present for BaseProperties, not even an explicit nil
### GetDnsParameters

`func (o *IamLdapPolicyInventory) GetDnsParameters() IamLdapDnsParameters`

GetDnsParameters returns the DnsParameters field if non-nil, zero value otherwise.

### GetDnsParametersOk

`func (o *IamLdapPolicyInventory) GetDnsParametersOk() (*IamLdapDnsParameters, bool)`

GetDnsParametersOk returns a tuple with the DnsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsParameters

`func (o *IamLdapPolicyInventory) SetDnsParameters(v IamLdapDnsParameters)`

SetDnsParameters sets DnsParameters field to given value.

### HasDnsParameters

`func (o *IamLdapPolicyInventory) HasDnsParameters() bool`

HasDnsParameters returns a boolean if a field has been set.

### SetDnsParametersNil

`func (o *IamLdapPolicyInventory) SetDnsParametersNil(b bool)`

 SetDnsParametersNil sets the value for DnsParameters to be an explicit nil

### UnsetDnsParameters
`func (o *IamLdapPolicyInventory) UnsetDnsParameters()`

UnsetDnsParameters ensures that no value is present for DnsParameters, not even an explicit nil
### GetEnableDns

`func (o *IamLdapPolicyInventory) GetEnableDns() bool`

GetEnableDns returns the EnableDns field if non-nil, zero value otherwise.

### GetEnableDnsOk

`func (o *IamLdapPolicyInventory) GetEnableDnsOk() (*bool, bool)`

GetEnableDnsOk returns a tuple with the EnableDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDns

`func (o *IamLdapPolicyInventory) SetEnableDns(v bool)`

SetEnableDns sets EnableDns field to given value.

### HasEnableDns

`func (o *IamLdapPolicyInventory) HasEnableDns() bool`

HasEnableDns returns a boolean if a field has been set.

### GetEnabled

`func (o *IamLdapPolicyInventory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamLdapPolicyInventory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamLdapPolicyInventory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamLdapPolicyInventory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserSearchPrecedence

`func (o *IamLdapPolicyInventory) GetUserSearchPrecedence() string`

GetUserSearchPrecedence returns the UserSearchPrecedence field if non-nil, zero value otherwise.

### GetUserSearchPrecedenceOk

`func (o *IamLdapPolicyInventory) GetUserSearchPrecedenceOk() (*string, bool)`

GetUserSearchPrecedenceOk returns a tuple with the UserSearchPrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchPrecedence

`func (o *IamLdapPolicyInventory) SetUserSearchPrecedence(v string)`

SetUserSearchPrecedence sets UserSearchPrecedence field to given value.

### HasUserSearchPrecedence

`func (o *IamLdapPolicyInventory) HasUserSearchPrecedence() bool`

HasUserSearchPrecedence returns a boolean if a field has been set.

### GetGroups

`func (o *IamLdapPolicyInventory) GetGroups() []IamLdapGroupInventoryRelationship`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamLdapPolicyInventory) GetGroupsOk() (*[]IamLdapGroupInventoryRelationship, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamLdapPolicyInventory) SetGroups(v []IamLdapGroupInventoryRelationship)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamLdapPolicyInventory) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IamLdapPolicyInventory) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IamLdapPolicyInventory) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetProviders

`func (o *IamLdapPolicyInventory) GetProviders() []IamLdapProviderInventoryRelationship`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IamLdapPolicyInventory) GetProvidersOk() (*[]IamLdapProviderInventoryRelationship, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IamLdapPolicyInventory) SetProviders(v []IamLdapProviderInventoryRelationship)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IamLdapPolicyInventory) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *IamLdapPolicyInventory) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *IamLdapPolicyInventory) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetTargetMo

`func (o *IamLdapPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *IamLdapPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *IamLdapPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *IamLdapPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *IamLdapPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *IamLdapPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


