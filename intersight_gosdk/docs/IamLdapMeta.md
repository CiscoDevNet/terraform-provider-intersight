# IamLdapMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapMeta"]
**Connectors** | Pointer to **[]string** |  | [optional] 
**Domain** | Pointer to **string** | The domain of the Ldap IdP corresponding to the appliance Ldap configuration. | [optional] [readonly] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 

## Methods

### NewIamLdapMeta

`func NewIamLdapMeta(classId string, objectType string, ) *IamLdapMeta`

NewIamLdapMeta instantiates a new IamLdapMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapMetaWithDefaults

`func NewIamLdapMetaWithDefaults() *IamLdapMeta`

NewIamLdapMetaWithDefaults instantiates a new IamLdapMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectors

`func (o *IamLdapMeta) GetConnectors() []string`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *IamLdapMeta) GetConnectorsOk() (*[]string, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *IamLdapMeta) SetConnectors(v []string)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *IamLdapMeta) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### SetConnectorsNil

`func (o *IamLdapMeta) SetConnectorsNil(b bool)`

 SetConnectorsNil sets the value for Connectors to be an explicit nil

### UnsetConnectors
`func (o *IamLdapMeta) UnsetConnectors()`

UnsetConnectors ensures that no value is present for Connectors, not even an explicit nil
### GetDomain

`func (o *IamLdapMeta) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IamLdapMeta) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IamLdapMeta) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IamLdapMeta) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIdp

`func (o *IamLdapMeta) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamLdapMeta) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamLdapMeta) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamLdapMeta) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *IamLdapMeta) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *IamLdapMeta) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


