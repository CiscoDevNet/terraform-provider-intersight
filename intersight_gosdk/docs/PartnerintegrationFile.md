# PartnerintegrationFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.File"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.File"]
**FilePath** | Pointer to **string** | Path of the file being uploaded. | [optional] [readonly] 
**FileType** | Pointer to **string** | Type of the file being uploaded. * &#x60;None&#x60; - Invalid file type for partnerIntegration appliance. * &#x60;Model&#x60; - Model file of Generic Device. * &#x60;Etl&#x60; - ETL file of Generic Device. * &#x60;Ui&#x60; - UI file of Generic Device. * &#x60;DeviceConnector&#x60; - Generic Device Connector file. | [optional] [default to "None"]
**WorkspaceName** | Pointer to **string** | The partner integration workspace to use to upload the File. | [optional] [readonly] [default to "default"]
**Catalog** | Pointer to [**NullableSoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationFile

`func NewPartnerintegrationFile(classId string, objectType string, ) *PartnerintegrationFile`

NewPartnerintegrationFile instantiates a new PartnerintegrationFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationFileWithDefaults

`func NewPartnerintegrationFileWithDefaults() *PartnerintegrationFile`

NewPartnerintegrationFileWithDefaults instantiates a new PartnerintegrationFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *PartnerintegrationFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *PartnerintegrationFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *PartnerintegrationFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *PartnerintegrationFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileType

`func (o *PartnerintegrationFile) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *PartnerintegrationFile) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *PartnerintegrationFile) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *PartnerintegrationFile) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetWorkspaceName

`func (o *PartnerintegrationFile) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *PartnerintegrationFile) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *PartnerintegrationFile) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *PartnerintegrationFile) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.

### GetCatalog

`func (o *PartnerintegrationFile) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *PartnerintegrationFile) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *PartnerintegrationFile) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *PartnerintegrationFile) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *PartnerintegrationFile) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *PartnerintegrationFile) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


