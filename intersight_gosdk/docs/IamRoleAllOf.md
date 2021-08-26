# IamRoleAllOf

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

### NewIamRoleAllOf

`func NewIamRoleAllOf(classId string, objectType string, ) *IamRoleAllOf`

NewIamRoleAllOf instantiates a new IamRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleAllOfWithDefaults

`func NewIamRoleAllOfWithDefaults() *IamRoleAllOf`

NewIamRoleAllOfWithDefaults instantiates a new IamRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamRoleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamRoleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamRoleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamRoleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamRoleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamRoleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *IamRoleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRoleAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRoleAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamRoleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamRoleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamRoleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamRoleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamRoleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivilegeNames

`func (o *IamRoleAllOf) GetPrivilegeNames() []string`

GetPrivilegeNames returns the PrivilegeNames field if non-nil, zero value otherwise.

### GetPrivilegeNamesOk

`func (o *IamRoleAllOf) GetPrivilegeNamesOk() (*[]string, bool)`

GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeNames

`func (o *IamRoleAllOf) SetPrivilegeNames(v []string)`

SetPrivilegeNames sets PrivilegeNames field to given value.

### HasPrivilegeNames

`func (o *IamRoleAllOf) HasPrivilegeNames() bool`

HasPrivilegeNames returns a boolean if a field has been set.

### SetPrivilegeNamesNil

`func (o *IamRoleAllOf) SetPrivilegeNamesNil(b bool)`

 SetPrivilegeNamesNil sets the value for PrivilegeNames to be an explicit nil

### UnsetPrivilegeNames
`func (o *IamRoleAllOf) UnsetPrivilegeNames()`

UnsetPrivilegeNames ensures that no value is present for PrivilegeNames, not even an explicit nil
### GetAccount

`func (o *IamRoleAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamRoleAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamRoleAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamRoleAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPrivilegeSets

`func (o *IamRoleAllOf) GetPrivilegeSets() []IamPrivilegeSetRelationship`

GetPrivilegeSets returns the PrivilegeSets field if non-nil, zero value otherwise.

### GetPrivilegeSetsOk

`func (o *IamRoleAllOf) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSets

`func (o *IamRoleAllOf) SetPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetPrivilegeSets sets PrivilegeSets field to given value.

### HasPrivilegeSets

`func (o *IamRoleAllOf) HasPrivilegeSets() bool`

HasPrivilegeSets returns a boolean if a field has been set.

### SetPrivilegeSetsNil

`func (o *IamRoleAllOf) SetPrivilegeSetsNil(b bool)`

 SetPrivilegeSetsNil sets the value for PrivilegeSets to be an explicit nil

### UnsetPrivilegeSets
`func (o *IamRoleAllOf) UnsetPrivilegeSets()`

UnsetPrivilegeSets ensures that no value is present for PrivilegeSets, not even an explicit nil
### GetSystem

`func (o *IamRoleAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamRoleAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamRoleAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamRoleAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


