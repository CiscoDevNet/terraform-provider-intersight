# ConfigExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadPath** | Pointer to **string** | Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired. | [optional] [readonly] 
**Items** | Pointer to [**[]ConfigMoRef**](config.MoRef.md) |  | [optional] 
**Name** | Pointer to **string** | An identifier for the exporter instance. | [optional] 
**Status** | Pointer to **string** | Status of the export operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**ExportedItems** | Pointer to [**[]ConfigExportedItemRelationship**](config.ExportedItem.Relationship.md) | An array of relationships to configExportedItem resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewConfigExporter

`func NewConfigExporter() *ConfigExporter`

NewConfigExporter instantiates a new ConfigExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExporterWithDefaults

`func NewConfigExporterWithDefaults() *ConfigExporter`

NewConfigExporterWithDefaults instantiates a new ConfigExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadPath

`func (o *ConfigExporter) GetDownloadPath() string`

GetDownloadPath returns the DownloadPath field if non-nil, zero value otherwise.

### GetDownloadPathOk

`func (o *ConfigExporter) GetDownloadPathOk() (*string, bool)`

GetDownloadPathOk returns a tuple with the DownloadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPath

`func (o *ConfigExporter) SetDownloadPath(v string)`

SetDownloadPath sets DownloadPath field to given value.

### HasDownloadPath

`func (o *ConfigExporter) HasDownloadPath() bool`

HasDownloadPath returns a boolean if a field has been set.

### GetItems

`func (o *ConfigExporter) GetItems() []ConfigMoRef`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConfigExporter) GetItemsOk() (*[]ConfigMoRef, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConfigExporter) SetItems(v []ConfigMoRef)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConfigExporter) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *ConfigExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigExporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigExporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigExporter) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigExporter) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigExporter) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigExporter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigExporter) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigExporter) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigExporter) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigExporter) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExportedItems

`func (o *ConfigExporter) GetExportedItems() []ConfigExportedItemRelationship`

GetExportedItems returns the ExportedItems field if non-nil, zero value otherwise.

### GetExportedItemsOk

`func (o *ConfigExporter) GetExportedItemsOk() (*[]ConfigExportedItemRelationship, bool)`

GetExportedItemsOk returns a tuple with the ExportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedItems

`func (o *ConfigExporter) SetExportedItems(v []ConfigExportedItemRelationship)`

SetExportedItems sets ExportedItems field to given value.

### HasExportedItems

`func (o *ConfigExporter) HasExportedItems() bool`

HasExportedItems returns a boolean if a field has been set.

### SetExportedItemsNil

`func (o *ConfigExporter) SetExportedItemsNil(b bool)`

 SetExportedItemsNil sets the value for ExportedItems to be an explicit nil

### UnsetExportedItems
`func (o *ConfigExporter) UnsetExportedItems()`

UnsetExportedItems ensures that no value is present for ExportedItems, not even an explicit nil
### GetOrganization

`func (o *ConfigExporter) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConfigExporter) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConfigExporter) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConfigExporter) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


