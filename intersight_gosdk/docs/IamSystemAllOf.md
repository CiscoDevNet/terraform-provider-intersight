# IamSystemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.System"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.System"]
**EndPointPrivileges** | Pointer to [**[]IamEndPointPrivilegeRelationship**](IamEndPointPrivilegeRelationship.md) | An array of relationships to iamEndPointPrivilege resources. | [optional] [readonly] 
**EndPointRoles** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**Idp** | Pointer to [**IamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](IamPrivilegeRelationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**Roles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] [readonly] 
**ServiceProvider** | Pointer to [**IamServiceProviderRelationship**](IamServiceProviderRelationship.md) |  | [optional] 

## Methods

### NewIamSystemAllOf

`func NewIamSystemAllOf(classId string, objectType string, ) *IamSystemAllOf`

NewIamSystemAllOf instantiates a new IamSystemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSystemAllOfWithDefaults

`func NewIamSystemAllOfWithDefaults() *IamSystemAllOf`

NewIamSystemAllOfWithDefaults instantiates a new IamSystemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSystemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSystemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSystemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSystemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSystemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSystemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEndPointPrivileges

`func (o *IamSystemAllOf) GetEndPointPrivileges() []IamEndPointPrivilegeRelationship`

GetEndPointPrivileges returns the EndPointPrivileges field if non-nil, zero value otherwise.

### GetEndPointPrivilegesOk

`func (o *IamSystemAllOf) GetEndPointPrivilegesOk() (*[]IamEndPointPrivilegeRelationship, bool)`

GetEndPointPrivilegesOk returns a tuple with the EndPointPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointPrivileges

`func (o *IamSystemAllOf) SetEndPointPrivileges(v []IamEndPointPrivilegeRelationship)`

SetEndPointPrivileges sets EndPointPrivileges field to given value.

### HasEndPointPrivileges

`func (o *IamSystemAllOf) HasEndPointPrivileges() bool`

HasEndPointPrivileges returns a boolean if a field has been set.

### SetEndPointPrivilegesNil

`func (o *IamSystemAllOf) SetEndPointPrivilegesNil(b bool)`

 SetEndPointPrivilegesNil sets the value for EndPointPrivileges to be an explicit nil

### UnsetEndPointPrivileges
`func (o *IamSystemAllOf) UnsetEndPointPrivileges()`

UnsetEndPointPrivileges ensures that no value is present for EndPointPrivileges, not even an explicit nil
### GetEndPointRoles

`func (o *IamSystemAllOf) GetEndPointRoles() []IamEndPointRoleRelationship`

GetEndPointRoles returns the EndPointRoles field if non-nil, zero value otherwise.

### GetEndPointRolesOk

`func (o *IamSystemAllOf) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRolesOk returns a tuple with the EndPointRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRoles

`func (o *IamSystemAllOf) SetEndPointRoles(v []IamEndPointRoleRelationship)`

SetEndPointRoles sets EndPointRoles field to given value.

### HasEndPointRoles

`func (o *IamSystemAllOf) HasEndPointRoles() bool`

HasEndPointRoles returns a boolean if a field has been set.

### SetEndPointRolesNil

`func (o *IamSystemAllOf) SetEndPointRolesNil(b bool)`

 SetEndPointRolesNil sets the value for EndPointRoles to be an explicit nil

### UnsetEndPointRoles
`func (o *IamSystemAllOf) UnsetEndPointRoles()`

UnsetEndPointRoles ensures that no value is present for EndPointRoles, not even an explicit nil
### GetIdp

`func (o *IamSystemAllOf) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamSystemAllOf) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamSystemAllOf) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamSystemAllOf) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPrivilegeSets

`func (o *IamSystemAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamSystemAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamSystemAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamSystemAllOf) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamSystemAllOf) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamSystemAllOf) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetPrivileges

`func (o *IamSystemAllOf) GetPrivileges() []IamPrivilegeRelationship`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *IamSystemAllOf) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *IamSystemAllOf) SetPrivileges(v []IamPrivilegeRelationship)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *IamSystemAllOf) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### SetPrivilegesNil

`func (o *IamSystemAllOf) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *IamSystemAllOf) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil
### GetRoles

`func (o *IamSystemAllOf) GetRoles() []IamRoleRelationship`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamSystemAllOf) GetRolesOk() (*[]IamRoleRelationship, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamSystemAllOf) SetRoles(v []IamRoleRelationship)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamSystemAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *IamSystemAllOf) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *IamSystemAllOf) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetServiceProvider

`func (o *IamSystemAllOf) GetServiceProvider() IamServiceProviderRelationship`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *IamSystemAllOf) GetServiceProviderOk() (*IamServiceProviderRelationship, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *IamSystemAllOf) SetServiceProvider(v IamServiceProviderRelationship)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *IamSystemAllOf) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


