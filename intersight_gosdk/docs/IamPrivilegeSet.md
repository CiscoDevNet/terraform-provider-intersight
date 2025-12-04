# IamPrivilegeSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.PrivilegeSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.PrivilegeSet"]
**AllowFutureUpdates** | Pointer to **bool** | Flag used by UI to keep track of the user selection option for future updates of privilege sets. | [optional] [default to true]
**Description** | Pointer to **string** | Description of the privilege set. | [optional] 
**IsPrivilegeNamesUpdated** | Pointer to **bool** | Flag to indicate if the privilege names are updated. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the privilege set. | [optional] 
**PrivilegeNames** | Pointer to **[]string** |  | [optional] 
**PrivilegeSetType** | Pointer to **string** | Type of the privilege set. * &#x60;Internal&#x60; - Privilege set is internal to the system. * &#x60;SystemPackaged&#x60; - Privilege set is packaged by the system and user can use it as a reference for custom privilege set creation. * &#x60;SystemDefined&#x60; - Privilege set is defined by the system. * &#x60;TreeNode&#x60; - Privilege set is a tree node in the custom privilege set creation hierarchy. * &#x60;TreeRoot&#x60; - Privilege set is a tree root in the custom privilege set creation hierarchy. * &#x60;TreeLeaf&#x60; - Privilege set is a tree leaf in the custom privilege set creation hierarchy. * &#x60;UserCreated&#x60; - Privilege set is created by the user. | [optional] [readonly] [default to "Internal"]
**Scope** | Pointer to **string** | The scope of the privilege set. * &#x60;Generic&#x60; - Privilege set can be added to account wide permission or organization permissions. * &#x60;Account&#x60; - Privilege set can be added to account wide permission only. | [optional] [readonly] [default to "Generic"]
**Uuid** | Pointer to **string** | UUID of the privilege set. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**AssociatedPrivilegeSets** | Pointer to [**[]IamPrivilegeSetRelationship**](IamPrivilegeSetRelationship.md) | An array of relationships to iamPrivilegeSet resources. | [optional] 
**Privileges** | Pointer to [**[]IamPrivilegeRelationship**](IamPrivilegeRelationship.md) | An array of relationships to iamPrivilege resources. | [optional] [readonly] 
**System** | Pointer to [**NullableIamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamPrivilegeSet

`func NewIamPrivilegeSet(classId string, objectType string, ) *IamPrivilegeSet`

NewIamPrivilegeSet instantiates a new IamPrivilegeSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamPrivilegeSetWithDefaults

`func NewIamPrivilegeSetWithDefaults() *IamPrivilegeSet`

NewIamPrivilegeSetWithDefaults instantiates a new IamPrivilegeSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamPrivilegeSet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamPrivilegeSet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamPrivilegeSet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamPrivilegeSet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamPrivilegeSet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamPrivilegeSet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowFutureUpdates

`func (o *IamPrivilegeSet) GetAllowFutureUpdates() bool`

GetAllowFutureUpdates returns the AllowFutureUpdates field if non-nil, zero value otherwise.

### GetAllowFutureUpdatesOk

`func (o *IamPrivilegeSet) GetAllowFutureUpdatesOk() (*bool, bool)`

GetAllowFutureUpdatesOk returns a tuple with the AllowFutureUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFutureUpdates

`func (o *IamPrivilegeSet) SetAllowFutureUpdates(v bool)`

SetAllowFutureUpdates sets AllowFutureUpdates field to given value.

### HasAllowFutureUpdates

`func (o *IamPrivilegeSet) HasAllowFutureUpdates() bool`

HasAllowFutureUpdates returns a boolean if a field has been set.

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

### GetIsPrivilegeNamesUpdated

`func (o *IamPrivilegeSet) GetIsPrivilegeNamesUpdated() bool`

GetIsPrivilegeNamesUpdated returns the IsPrivilegeNamesUpdated field if non-nil, zero value otherwise.

### GetIsPrivilegeNamesUpdatedOk

`func (o *IamPrivilegeSet) GetIsPrivilegeNamesUpdatedOk() (*bool, bool)`

GetIsPrivilegeNamesUpdatedOk returns a tuple with the IsPrivilegeNamesUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivilegeNamesUpdated

`func (o *IamPrivilegeSet) SetIsPrivilegeNamesUpdated(v bool)`

SetIsPrivilegeNamesUpdated sets IsPrivilegeNamesUpdated field to given value.

### HasIsPrivilegeNamesUpdated

`func (o *IamPrivilegeSet) HasIsPrivilegeNamesUpdated() bool`

HasIsPrivilegeNamesUpdated returns a boolean if a field has been set.

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

### SetPrivilegeNamesNil

`func (o *IamPrivilegeSet) SetPrivilegeNamesNil(b bool)`

 SetPrivilegeNamesNil sets the value for PrivilegeNames to be an explicit nil

### UnsetPrivilegeNames
`func (o *IamPrivilegeSet) UnsetPrivilegeNames()`

UnsetPrivilegeNames ensures that no value is present for PrivilegeNames, not even an explicit nil
### GetPrivilegeSetType

`func (o *IamPrivilegeSet) GetPrivilegeSetType() string`

GetPrivilegeSetType returns the PrivilegeSetType field if non-nil, zero value otherwise.

### GetPrivilegeSetTypeOk

`func (o *IamPrivilegeSet) GetPrivilegeSetTypeOk() (*string, bool)`

GetPrivilegeSetTypeOk returns a tuple with the PrivilegeSetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeSetType

`func (o *IamPrivilegeSet) SetPrivilegeSetType(v string)`

SetPrivilegeSetType sets PrivilegeSetType field to given value.

### HasPrivilegeSetType

`func (o *IamPrivilegeSet) HasPrivilegeSetType() bool`

HasPrivilegeSetType returns a boolean if a field has been set.

### GetScope

`func (o *IamPrivilegeSet) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamPrivilegeSet) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamPrivilegeSet) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamPrivilegeSet) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUuid

`func (o *IamPrivilegeSet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IamPrivilegeSet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IamPrivilegeSet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IamPrivilegeSet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

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

### SetAccountNil

`func (o *IamPrivilegeSet) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamPrivilegeSet) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
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

### SetSystemNil

`func (o *IamPrivilegeSet) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *IamPrivilegeSet) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


