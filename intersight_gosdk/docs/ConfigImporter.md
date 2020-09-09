# ConfigImporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportPath** | Pointer to **string** | The path to the archive in Intersight storage that has all the MOs to be imported. | [optional] 
**ImportSource** | Pointer to **string** | The source of the archive in Intersight storage that has all the MOs to be imported. * &#x60;ImageRepo&#x60; - The &#39;ImageRepo&#39; source if the source of exporter archive is image repository. * &#x60;URL&#x60; - The &#39;URL&#39; source if the source of exported archive is a URL. | [optional] [default to "ImageRepo"]
**Name** | Pointer to **string** | An identifier for the importer instance. | [optional] 
**SkipIntegrityChecks** | Pointer to **bool** | Specifies whether integrity checks must be skipped. | [optional] 
**Status** | Pointer to **string** | Status of the import operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**ImportedItems** | Pointer to [**[]ConfigImportedItemRelationship**](config.ImportedItem.Relationship.md) | An array of relationships to configImportedItem resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewConfigImporter

`func NewConfigImporter() *ConfigImporter`

NewConfigImporter instantiates a new ConfigImporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigImporterWithDefaults

`func NewConfigImporterWithDefaults() *ConfigImporter`

NewConfigImporterWithDefaults instantiates a new ConfigImporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportPath

`func (o *ConfigImporter) GetImportPath() string`

GetImportPath returns the ImportPath field if non-nil, zero value otherwise.

### GetImportPathOk

`func (o *ConfigImporter) GetImportPathOk() (*string, bool)`

GetImportPathOk returns a tuple with the ImportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPath

`func (o *ConfigImporter) SetImportPath(v string)`

SetImportPath sets ImportPath field to given value.

### HasImportPath

`func (o *ConfigImporter) HasImportPath() bool`

HasImportPath returns a boolean if a field has been set.

### GetImportSource

`func (o *ConfigImporter) GetImportSource() string`

GetImportSource returns the ImportSource field if non-nil, zero value otherwise.

### GetImportSourceOk

`func (o *ConfigImporter) GetImportSourceOk() (*string, bool)`

GetImportSourceOk returns a tuple with the ImportSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSource

`func (o *ConfigImporter) SetImportSource(v string)`

SetImportSource sets ImportSource field to given value.

### HasImportSource

`func (o *ConfigImporter) HasImportSource() bool`

HasImportSource returns a boolean if a field has been set.

### GetName

`func (o *ConfigImporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigImporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigImporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigImporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipIntegrityChecks

`func (o *ConfigImporter) GetSkipIntegrityChecks() bool`

GetSkipIntegrityChecks returns the SkipIntegrityChecks field if non-nil, zero value otherwise.

### GetSkipIntegrityChecksOk

`func (o *ConfigImporter) GetSkipIntegrityChecksOk() (*bool, bool)`

GetSkipIntegrityChecksOk returns a tuple with the SkipIntegrityChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntegrityChecks

`func (o *ConfigImporter) SetSkipIntegrityChecks(v bool)`

SetSkipIntegrityChecks sets SkipIntegrityChecks field to given value.

### HasSkipIntegrityChecks

`func (o *ConfigImporter) HasSkipIntegrityChecks() bool`

HasSkipIntegrityChecks returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigImporter) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigImporter) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigImporter) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigImporter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigImporter) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigImporter) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigImporter) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigImporter) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetImportedItems

`func (o *ConfigImporter) GetImportedItems() []ConfigImportedItemRelationship`

GetImportedItems returns the ImportedItems field if non-nil, zero value otherwise.

### GetImportedItemsOk

`func (o *ConfigImporter) GetImportedItemsOk() (*[]ConfigImportedItemRelationship, bool)`

GetImportedItemsOk returns a tuple with the ImportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedItems

`func (o *ConfigImporter) SetImportedItems(v []ConfigImportedItemRelationship)`

SetImportedItems sets ImportedItems field to given value.

### HasImportedItems

`func (o *ConfigImporter) HasImportedItems() bool`

HasImportedItems returns a boolean if a field has been set.

### SetImportedItemsNil

`func (o *ConfigImporter) SetImportedItemsNil(b bool)`

 SetImportedItemsNil sets the value for ImportedItems to be an explicit nil

### UnsetImportedItems
`func (o *ConfigImporter) UnsetImportedItems()`

UnsetImportedItems ensures that no value is present for ImportedItems, not even an explicit nil
### GetOrganization

`func (o *ConfigImporter) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConfigImporter) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConfigImporter) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConfigImporter) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


