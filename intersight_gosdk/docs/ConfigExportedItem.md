# ConfigExportedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Name of the file corresponding to item MO. | [optional] [readonly] 
**Item** | Pointer to [**ConfigMoRef**](config.MoRef.md) |  | [optional] 
**Name** | Pointer to **string** | MO item identity (the moref corresponding to item) expressed as a string. | [optional] [readonly] 
**ServiceVersion** | Pointer to **string** | Version of the service that owns the item MO. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the item&#39;s export operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Progress or error message for the MO&#39;s export operation. | [optional] [readonly] 
**Exporter** | Pointer to [**ConfigExporterRelationship**](config.Exporter.Relationship.md) |  | [optional] 
**ParentItem** | Pointer to [**ConfigExportedItemRelationship**](config.ExportedItem.Relationship.md) |  | [optional] 
**RelatedItems** | Pointer to [**[]ConfigExportedItemRelationship**](config.ExportedItem.Relationship.md) | An array of relationships to configExportedItem resources. | [optional] [readonly] 

## Methods

### NewConfigExportedItem

`func NewConfigExportedItem() *ConfigExportedItem`

NewConfigExportedItem instantiates a new ConfigExportedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExportedItemWithDefaults

`func NewConfigExportedItemWithDefaults() *ConfigExportedItem`

NewConfigExportedItemWithDefaults instantiates a new ConfigExportedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *ConfigExportedItem) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ConfigExportedItem) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ConfigExportedItem) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ConfigExportedItem) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetItem

`func (o *ConfigExportedItem) GetItem() ConfigMoRef`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ConfigExportedItem) GetItemOk() (*ConfigMoRef, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ConfigExportedItem) SetItem(v ConfigMoRef)`

SetItem sets Item field to given value.

### HasItem

`func (o *ConfigExportedItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetName

`func (o *ConfigExportedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigExportedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigExportedItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigExportedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ConfigExportedItem) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ConfigExportedItem) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ConfigExportedItem) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ConfigExportedItem) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigExportedItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigExportedItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigExportedItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigExportedItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigExportedItem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigExportedItem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigExportedItem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigExportedItem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExporter

`func (o *ConfigExportedItem) GetExporter() ConfigExporterRelationship`

GetExporter returns the Exporter field if non-nil, zero value otherwise.

### GetExporterOk

`func (o *ConfigExportedItem) GetExporterOk() (*ConfigExporterRelationship, bool)`

GetExporterOk returns a tuple with the Exporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporter

`func (o *ConfigExportedItem) SetExporter(v ConfigExporterRelationship)`

SetExporter sets Exporter field to given value.

### HasExporter

`func (o *ConfigExportedItem) HasExporter() bool`

HasExporter returns a boolean if a field has been set.

### GetParentItem

`func (o *ConfigExportedItem) GetParentItem() ConfigExportedItemRelationship`

GetParentItem returns the ParentItem field if non-nil, zero value otherwise.

### GetParentItemOk

`func (o *ConfigExportedItem) GetParentItemOk() (*ConfigExportedItemRelationship, bool)`

GetParentItemOk returns a tuple with the ParentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItem

`func (o *ConfigExportedItem) SetParentItem(v ConfigExportedItemRelationship)`

SetParentItem sets ParentItem field to given value.

### HasParentItem

`func (o *ConfigExportedItem) HasParentItem() bool`

HasParentItem returns a boolean if a field has been set.

### GetRelatedItems

`func (o *ConfigExportedItem) GetRelatedItems() []ConfigExportedItemRelationship`

GetRelatedItems returns the RelatedItems field if non-nil, zero value otherwise.

### GetRelatedItemsOk

`func (o *ConfigExportedItem) GetRelatedItemsOk() (*[]ConfigExportedItemRelationship, bool)`

GetRelatedItemsOk returns a tuple with the RelatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedItems

`func (o *ConfigExportedItem) SetRelatedItems(v []ConfigExportedItemRelationship)`

SetRelatedItems sets RelatedItems field to given value.

### HasRelatedItems

`func (o *ConfigExportedItem) HasRelatedItems() bool`

HasRelatedItems returns a boolean if a field has been set.

### SetRelatedItemsNil

`func (o *ConfigExportedItem) SetRelatedItemsNil(b bool)`

 SetRelatedItemsNil sets the value for RelatedItems to be an explicit nil

### UnsetRelatedItems
`func (o *ConfigExportedItem) UnsetRelatedItems()`

UnsetRelatedItems ensures that no value is present for RelatedItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


