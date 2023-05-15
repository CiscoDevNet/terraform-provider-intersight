# PartnerintegrationFileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "partnerintegration.File"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "partnerintegration.File"]
**FilePath** | Pointer to **string** | Path of the file being uploaded. | [optional] [readonly] 
**FileType** | Pointer to **string** | Type of the file being uploaded. * &#x60;None&#x60; - Invalid file type for partnerIntegration appliance. * &#x60;Model&#x60; - Model file of Generic Device. * &#x60;Etl&#x60; - ETL file of Generic Device. * &#x60;Ui&#x60; - UI file of Generic Device. * &#x60;DeviceConnector&#x60; - Generic Device Connector file. | [optional] [default to "None"]
**WorkspaceName** | Pointer to **string** | The partner integration workspace to use to upload the File. | [optional] [readonly] [default to "default"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewPartnerintegrationFileAllOf

`func NewPartnerintegrationFileAllOf(classId string, objectType string, ) *PartnerintegrationFileAllOf`

NewPartnerintegrationFileAllOf instantiates a new PartnerintegrationFileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerintegrationFileAllOfWithDefaults

`func NewPartnerintegrationFileAllOfWithDefaults() *PartnerintegrationFileAllOf`

NewPartnerintegrationFileAllOfWithDefaults instantiates a new PartnerintegrationFileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PartnerintegrationFileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PartnerintegrationFileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PartnerintegrationFileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PartnerintegrationFileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PartnerintegrationFileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PartnerintegrationFileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFilePath

`func (o *PartnerintegrationFileAllOf) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *PartnerintegrationFileAllOf) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *PartnerintegrationFileAllOf) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *PartnerintegrationFileAllOf) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileType

`func (o *PartnerintegrationFileAllOf) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *PartnerintegrationFileAllOf) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *PartnerintegrationFileAllOf) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *PartnerintegrationFileAllOf) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetWorkspaceName

`func (o *PartnerintegrationFileAllOf) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *PartnerintegrationFileAllOf) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *PartnerintegrationFileAllOf) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *PartnerintegrationFileAllOf) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.

### GetCatalog

`func (o *PartnerintegrationFileAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *PartnerintegrationFileAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *PartnerintegrationFileAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *PartnerintegrationFileAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


