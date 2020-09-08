# IamPrivilegeSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the privilege set. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the privilege set. | [optional] 
**PrivilegeNames** | Pointer to **[]string** |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**AssociatedPrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](iam.PrivilegeSet.Relationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](iam.Privilege.Relationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamPrivilegeSet

`func NewIamPrivilegeSet() *IamPrivilegeSet`

NewIamPrivilegeSet instantiates a new IamPrivilegeSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeSetWithDefaults

`func NewIamPrivilegeSetWithDefaults() *IamPrivilegeSet`

NewIamPrivilegeSetWithDefaults instantiates a new IamPrivilegeSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IamPrivilegeSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamPrivilegeSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamPrivilegeSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamPrivilegeSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IamPrivilegeSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamPrivilegeSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamPrivilegeSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamPrivilegeSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivilegeNames

`func (o *IamPrivilegeSet) GetPrivilegeNames() []string`

GetPrivilegeNames returns the PrivilegeNames field if non-nil, zero value otherwise.

### GetPrivilegeNamesOk

`func (o *IamPrivilegeSet) GetPrivilegeNamesOk() (*[]string, bool)`

GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeNames

`func (o *IamPrivilegeSet) SetPrivilegeNames(v []string)`

SetPrivilegeNames sets PrivilegeNames field to given value.

### HasPrivilegeNames

`func (o *IamPrivilegeSet) HasPrivilegeNames() bool`

HasPrivilegeNames returns a boolean if a field has been set.

### GetAccount

`func (o *IamPrivilegeSet) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamPrivilegeSet) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamPrivilegeSet) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamPrivilegeSet) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssociatedPrivilegeSets

`func (o *IamPrivilegeSet) GetAssociatedPrivilegeSets() []IamPrivilegeSetRelationship`

GetAssociatedPrivilegeSets returns the AssociatedPrivilegeSets field if non-nil, zero value otherwise.

### GetAssociatedPrivilegeSetsOk

`func (o *IamPrivilegeSet) GetAssociatedPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool)`

GetAssociatedPrivilegeSetsOk returns a tuple with the AssociatedPrivilegeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPrivilegeSets

`func (o *IamPrivilegeSet) SetAssociatedPrivilegeSets(v []IamPrivilegeSetRelationship)`

SetAssociatedPrivilegeSets sets AssociatedPrivilegeSets field to given value.

### HasAssociatedPrivilegeSets

`func (o *IamPrivilegeSet) HasAssociatedPrivilegeSets() bool`

HasAssociatedPrivilegeSets returns a boolean if a field has been set.

### SetAssociatedPrivilegeSetsNil

`func (o *IamPrivilegeSet) SetAssociatedPrivilegeSetsNil(b bool)`

 SetAssociatedPrivilegeSetsNil sets the value for AssociatedPrivilegeSets to be an explicit nil

### UnsetAssociatedPrivilegeSets
`func (o *IamPrivilegeSet) UnsetAssociatedPrivilegeSets()`

UnsetAssociatedPrivilegeSets ensures that no value is present for AssociatedPrivilegeSets, not even an explicit nil
### GetPrivileges

`func (o *IamPrivilegeSet) GetPrivileges() []IamPrivilegeRelationship`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *IamPrivilegeSet) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *IamPrivilegeSet) SetPrivileges(v []IamPrivilegeRelationship)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *IamPrivilegeSet) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### SetPrivilegesNil

`func (o *IamPrivilegeSet) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *IamPrivilegeSet) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil
### GetSystem

`func (o *IamPrivilegeSet) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamPrivilegeSet) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamPrivilegeSet) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamPrivilegeSet) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


