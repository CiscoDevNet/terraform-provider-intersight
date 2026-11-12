# WorkloadManifestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.ManifestItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.ManifestItem"]
**ImportObjectType** | Pointer to **string** | The type of the object that is imported. | [optional] [readonly] 
**ImportStatus** | Pointer to **string** | The import status for the item to indicate if the item is available in the shared catalog. * &#x60;ImportPending&#x60; - The item from the manifest file is pending an import into the shared catalog. * &#x60;Imported&#x60; - The item from the manifest file has been imported in the shared catalog. | [optional] [default to "ImportPending"]
**Name** | Pointer to **string** | The name for the item that is in the manifest. The value here corresponds to the name for the imported object in the shared catalog. | [optional] [readonly] 
**Revision** | Pointer to **int64** | The revision number for the item specified by the manifest. | [optional] [readonly] [default to 0]
**Version** | Pointer to **int64** | The version number for the item. The value here will match the version number set for the item in the file used for the import. For policies and templates, it will be the same as the number suffixed in the item name. | [optional] [readonly] 

## Methods

### NewWorkloadManifestItem

`func NewWorkloadManifestItem(classId string, objectType string, ) *WorkloadManifestItem`

NewWorkloadManifestItem instantiates a new WorkloadManifestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadManifestItemWithDefaults

`func NewWorkloadManifestItemWithDefaults() *WorkloadManifestItem`

NewWorkloadManifestItemWithDefaults instantiates a new WorkloadManifestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadManifestItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadManifestItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadManifestItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadManifestItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadManifestItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadManifestItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImportObjectType

`func (o *WorkloadManifestItem) GetImportObjectType() string`

GetImportObjectType returns the ImportObjectType field if non-nil, zero value otherwise.

### GetImportObjectTypeOk

`func (o *WorkloadManifestItem) GetImportObjectTypeOk() (*string, bool)`

GetImportObjectTypeOk returns a tuple with the ImportObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportObjectType

`func (o *WorkloadManifestItem) SetImportObjectType(v string)`

SetImportObjectType sets ImportObjectType field to given value.

### HasImportObjectType

`func (o *WorkloadManifestItem) HasImportObjectType() bool`

HasImportObjectType returns a boolean if a field has been set.

### GetImportStatus

`func (o *WorkloadManifestItem) GetImportStatus() string`

GetImportStatus returns the ImportStatus field if non-nil, zero value otherwise.

### GetImportStatusOk

`func (o *WorkloadManifestItem) GetImportStatusOk() (*string, bool)`

GetImportStatusOk returns a tuple with the ImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStatus

`func (o *WorkloadManifestItem) SetImportStatus(v string)`

SetImportStatus sets ImportStatus field to given value.

### HasImportStatus

`func (o *WorkloadManifestItem) HasImportStatus() bool`

HasImportStatus returns a boolean if a field has been set.

### GetName

`func (o *WorkloadManifestItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadManifestItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadManifestItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadManifestItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevision

`func (o *WorkloadManifestItem) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *WorkloadManifestItem) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *WorkloadManifestItem) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *WorkloadManifestItem) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetVersion

`func (o *WorkloadManifestItem) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadManifestItem) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadManifestItem) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkloadManifestItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


