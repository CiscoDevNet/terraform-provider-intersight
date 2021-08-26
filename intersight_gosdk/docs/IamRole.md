# IamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.Role"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.Role"]
**Description** | Pointer to **string** | Informative description about each role. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the role which has to be granted to user. | [optional] 
**PrivilegeNames** | Pointer to **[]string** |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**PrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] [readonly] 
**System** | Pointer to [**IamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamRole

`func NewIamRole(classId string, objectType string, ) *IamRole`

NewIamRole instantiates a new IamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleWithDefaults

`func NewIamRoleWithDefaults() *IamRole`

NewIamRoleWithDefaults instantiates a new IamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *IamRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivilegeNames

`func (o *IamRole) GetPrivilegeNames() []string`

GetPrivilegeNames returns the PrivilegeNames field if non-nil, zero value otherwise.

### GetPrivilegeNamesOk

`func (o *IamRole) GetPrivilegeNamesOk() (*[]string, bool)`

GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeNames

`func (o *IamRole) SetPrivilegeNames(v []string)`

SetPrivilegeNames sets PrivilegeNames field to given value.

### HasPrivilegeNames

`func (o *IamRole) HasPrivilegeNames() bool`

HasPrivilegeNames returns a boolean if a field has been set.

### SetPrivilegeNamesNil

`func (o *IamRole) SetPrivilegeNamesNil(b bool)`

 SetPrivilegeNamesNil sets the value for PrivilegeNames to be an explicit nil

### UnsetPrivilegeNames
`func (o *IamRole) UnsetPrivilegeNames()`

UnsetPrivilegeNames ensures that no value is present for PrivilegeNames, not even an explicit nil
### GetAccount

`func (o *IamRole) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamRole) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamRole) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamRole) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPrivilegeSets

`func (o *IamRole) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamRole) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamRole) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamRole) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamRole) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamRole) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetSystem

`func (o *IamRole) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamRole) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamRole) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamRole) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


