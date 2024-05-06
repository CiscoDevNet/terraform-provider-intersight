# WorkspaceFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workspace.Folder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workspace.Folder"]
**Archived** | Pointer to **bool** | It is to define if folder is archived or not. | [optional] [default to false]
**CreateUser** | Pointer to **string** | The UserID or email who created this folder. | [optional] [readonly] 
**ModUser** | Pointer to **string** | The UserID or email who last modified this folder. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for this folder. You can have multiple versions of the folder with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**Assets** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 
**ParentFolder** | Pointer to [**WorkspaceFolderRelationship**](WorkspaceFolderRelationship.md) |  | [optional] 
**SubFolders** | Pointer to [**[]WorkspaceFolderRelationship**](WorkspaceFolderRelationship.md) | An array of relationships to workspaceFolder resources. | [optional] 

## Methods

### NewWorkspaceFolder

`func NewWorkspaceFolder(classId string, objectType string, ) *WorkspaceFolder`

NewWorkspaceFolder instantiates a new WorkspaceFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceFolderWithDefaults

`func NewWorkspaceFolderWithDefaults() *WorkspaceFolder`

NewWorkspaceFolderWithDefaults instantiates a new WorkspaceFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkspaceFolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkspaceFolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkspaceFolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkspaceFolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkspaceFolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkspaceFolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchived

`func (o *WorkspaceFolder) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkspaceFolder) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkspaceFolder) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WorkspaceFolder) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreateUser

`func (o *WorkspaceFolder) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *WorkspaceFolder) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *WorkspaceFolder) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *WorkspaceFolder) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetModUser

`func (o *WorkspaceFolder) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *WorkspaceFolder) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *WorkspaceFolder) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *WorkspaceFolder) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssets

`func (o *WorkspaceFolder) GetAssets() []MoBaseMoRelationship`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *WorkspaceFolder) GetAssetsOk() (*[]MoBaseMoRelationship, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *WorkspaceFolder) SetAssets(v []MoBaseMoRelationship)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *WorkspaceFolder) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### SetAssetsNil

`func (o *WorkspaceFolder) SetAssetsNil(b bool)`

 SetAssetsNil sets the value for Assets to be an explicit nil

### UnsetAssets
`func (o *WorkspaceFolder) UnsetAssets()`

UnsetAssets ensures that no value is present for Assets, not even an explicit nil
### GetCatalog

`func (o *WorkspaceFolder) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkspaceFolder) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkspaceFolder) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkspaceFolder) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetParentFolder

`func (o *WorkspaceFolder) GetParentFolder() WorkspaceFolderRelationship`

GetParentFolder returns the ParentFolder field if non-nil, zero value otherwise.

### GetParentFolderOk

`func (o *WorkspaceFolder) GetParentFolderOk() (*WorkspaceFolderRelationship, bool)`

GetParentFolderOk returns a tuple with the ParentFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolder

`func (o *WorkspaceFolder) SetParentFolder(v WorkspaceFolderRelationship)`

SetParentFolder sets ParentFolder field to given value.

### HasParentFolder

`func (o *WorkspaceFolder) HasParentFolder() bool`

HasParentFolder returns a boolean if a field has been set.

### GetSubFolders

`func (o *WorkspaceFolder) GetSubFolders() []WorkspaceFolderRelationship`

GetSubFolders returns the SubFolders field if non-nil, zero value otherwise.

### GetSubFoldersOk

`func (o *WorkspaceFolder) GetSubFoldersOk() (*[]WorkspaceFolderRelationship, bool)`

GetSubFoldersOk returns a tuple with the SubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFolders

`func (o *WorkspaceFolder) SetSubFolders(v []WorkspaceFolderRelationship)`

SetSubFolders sets SubFolders field to given value.

### HasSubFolders

`func (o *WorkspaceFolder) HasSubFolders() bool`

HasSubFolders returns a boolean if a field has been set.

### SetSubFoldersNil

`func (o *WorkspaceFolder) SetSubFoldersNil(b bool)`

 SetSubFoldersNil sets the value for SubFolders to be an explicit nil

### UnsetSubFolders
`func (o *WorkspaceFolder) UnsetSubFolders()`

UnsetSubFolders ensures that no value is present for SubFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


