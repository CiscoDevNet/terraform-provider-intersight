# WorkspaceFolderAllOf

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

### NewWorkspaceFolderAllOf

`func NewWorkspaceFolderAllOf(classId string, objectType string, ) *WorkspaceFolderAllOf`

NewWorkspaceFolderAllOf instantiates a new WorkspaceFolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceFolderAllOfWithDefaults

`func NewWorkspaceFolderAllOfWithDefaults() *WorkspaceFolderAllOf`

NewWorkspaceFolderAllOfWithDefaults instantiates a new WorkspaceFolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkspaceFolderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkspaceFolderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkspaceFolderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkspaceFolderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkspaceFolderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkspaceFolderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArchived

`func (o *WorkspaceFolderAllOf) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkspaceFolderAllOf) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkspaceFolderAllOf) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WorkspaceFolderAllOf) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreateUser

`func (o *WorkspaceFolderAllOf) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *WorkspaceFolderAllOf) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *WorkspaceFolderAllOf) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *WorkspaceFolderAllOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetModUser

`func (o *WorkspaceFolderAllOf) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *WorkspaceFolderAllOf) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *WorkspaceFolderAllOf) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *WorkspaceFolderAllOf) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceFolderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceFolderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceFolderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceFolderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssets

`func (o *WorkspaceFolderAllOf) GetAssets() []MoBaseMoRelationship`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *WorkspaceFolderAllOf) GetAssetsOk() (*[]MoBaseMoRelationship, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *WorkspaceFolderAllOf) SetAssets(v []MoBaseMoRelationship)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *WorkspaceFolderAllOf) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### SetAssetsNil

`func (o *WorkspaceFolderAllOf) SetAssetsNil(b bool)`

 SetAssetsNil sets the value for Assets to be an explicit nil

### UnsetAssets
`func (o *WorkspaceFolderAllOf) UnsetAssets()`

UnsetAssets ensures that no value is present for Assets, not even an explicit nil
### GetCatalog

`func (o *WorkspaceFolderAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkspaceFolderAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkspaceFolderAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkspaceFolderAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetParentFolder

`func (o *WorkspaceFolderAllOf) GetParentFolder() WorkspaceFolderRelationship`

GetParentFolder returns the ParentFolder field if non-nil, zero value otherwise.

### GetParentFolderOk

`func (o *WorkspaceFolderAllOf) GetParentFolderOk() (*WorkspaceFolderRelationship, bool)`

GetParentFolderOk returns a tuple with the ParentFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolder

`func (o *WorkspaceFolderAllOf) SetParentFolder(v WorkspaceFolderRelationship)`

SetParentFolder sets ParentFolder field to given value.

### HasParentFolder

`func (o *WorkspaceFolderAllOf) HasParentFolder() bool`

HasParentFolder returns a boolean if a field has been set.

### GetSubFolders

`func (o *WorkspaceFolderAllOf) GetSubFolders() []WorkspaceFolderRelationship`

GetSubFolders returns the SubFolders field if non-nil, zero value otherwise.

### GetSubFoldersOk

`func (o *WorkspaceFolderAllOf) GetSubFoldersOk() (*[]WorkspaceFolderRelationship, bool)`

GetSubFoldersOk returns a tuple with the SubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFolders

`func (o *WorkspaceFolderAllOf) SetSubFolders(v []WorkspaceFolderRelationship)`

SetSubFolders sets SubFolders field to given value.

### HasSubFolders

`func (o *WorkspaceFolderAllOf) HasSubFolders() bool`

HasSubFolders returns a boolean if a field has been set.

### SetSubFoldersNil

`func (o *WorkspaceFolderAllOf) SetSubFoldersNil(b bool)`

 SetSubFoldersNil sets the value for SubFolders to be an explicit nil

### UnsetSubFolders
`func (o *WorkspaceFolderAllOf) UnsetSubFolders()`

UnsetSubFolders ensures that no value is present for SubFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


