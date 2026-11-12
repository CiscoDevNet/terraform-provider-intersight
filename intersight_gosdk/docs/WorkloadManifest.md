# WorkloadManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.Manifest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.Manifest"]
**ImportStatus** | Pointer to **string** | The import status for the current manifest. If all items in the manifest items list are imported, the current manifest will be marked as imported. * &#x60;ImportPending&#x60; - The item from the manifest file is pending an import into the shared catalog. * &#x60;Imported&#x60; - The item from the manifest file has been imported in the shared catalog. | [optional] [default to "ImportPending"]
**ManifestItems** | Pointer to [**[]WorkloadManifestItem**](WorkloadManifestItem.md) |  | [optional] 
**Version** | Pointer to **string** | The version number for the manifest. | [optional] [readonly] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkloadManifest

`func NewWorkloadManifest(classId string, objectType string, ) *WorkloadManifest`

NewWorkloadManifest instantiates a new WorkloadManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadManifestWithDefaults

`func NewWorkloadManifestWithDefaults() *WorkloadManifest`

NewWorkloadManifestWithDefaults instantiates a new WorkloadManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadManifest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadManifest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadManifest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadManifest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadManifest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadManifest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImportStatus

`func (o *WorkloadManifest) GetImportStatus() string`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *WorkloadManifest) GetImportStatusOk() (*string, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *WorkloadManifest) SetImportStatus(v string)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *WorkloadManifest) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### GetManifestItems

`func (o *WorkloadManifest) GetManifestItems() []WorkloadManifestItem`

GetManifestItems returns the ManifestItems field if non-nil, zero value otherwise.

### GetManifestItemsOk

`func (o *WorkloadManifest) GetManifestItemsOk() (*[]WorkloadManifestItem, bool)`

GetManifestItemsOk returns a tuple with the ManifestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestItems

`func (o *WorkloadManifest) SetManifestItems(v []WorkloadManifestItem)`

SetManifestItems sets ManifestItems field to given value.

### HasManifestItems

`func (o *WorkloadManifest) HasManifestItems() bool`

HasManifestItems returns a boolean if a field has been set.

### SetManifestItemsNil

`func (o *WorkloadManifest) SetManifestItemsNil(b bool)`

 SetManifestItemsNil sets the value for ManifestItems to be an explicit nil

### UnsetManifestItems
`func (o *WorkloadManifest) UnsetManifestItems()`

UnsetManifestItems ensures that no value is present for ManifestItems, not even an explicit nil
### GetVersion

`func (o *WorkloadManifest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadManifest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadManifest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkloadManifest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkloadManifest) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkloadManifest) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkloadManifest) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkloadManifest) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *WorkloadManifest) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *WorkloadManifest) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


