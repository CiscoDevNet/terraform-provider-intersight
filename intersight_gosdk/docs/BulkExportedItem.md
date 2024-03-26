# BulkExportedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.ExportedItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.ExportedItem"]
**ExportTags** | Pointer to **bool** | Specifies whether tags must be exported for item MO. | [optional] [readonly] [default to false]
**FileName** | Pointer to **string** | Name of the file corresponding to item MO. | [optional] [readonly] 
**Item** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | MO item identity (the moref corresponding to item) expressed as a string. | [optional] [readonly] 
**ServiceName** | Pointer to **string** | Name of the target service that owns the item MO. Service responsible for handling exported item mo notifications. | [optional] [readonly] 
**ServiceVersion** | Pointer to **string** | Version of the service that owns the item MO. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the item&#39;s export operation. * &#x60;&#x60; - The operation has not started. * &#x60;ValidationInProgress&#x60; - The validation operation is in progress. * &#x60;Valid&#x60; - The content to be imported is valid. * &#x60;InValid&#x60; - The content to be imported is not valid and the status message will have the reason. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. * &#x60;OperationCancelled&#x60; - The operation has been canceled. * &#x60;CancelInProgress&#x60; - The operation is being canceled. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Progress or error message for the MO&#39;s export operation. | [optional] [readonly] 
**Export** | Pointer to [**BulkExportRelationship**](BulkExportRelationship.md) |  | [optional] 
**ParentItem** | Pointer to [**BulkExportedItemRelationship**](BulkExportedItemRelationship.md) |  | [optional] 
**RelatedItems** | Pointer to [**[]BulkExportedItemRelationship**](BulkExportedItemRelationship.md) | An array of relationships to bulkExportedItem resources. | [optional] [readonly] 

## Methods

### NewBulkExportedItem

`func NewBulkExportedItem(classId string, objectType string, ) *BulkExportedItem`

NewBulkExportedItem instantiates a new BulkExportedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkExportedItemWithDefaults

`func NewBulkExportedItemWithDefaults() *BulkExportedItem`

NewBulkExportedItemWithDefaults instantiates a new BulkExportedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkExportedItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkExportedItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkExportedItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkExportedItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkExportedItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkExportedItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExportTags

`func (o *BulkExportedItem) GetExportTags() bool`

GetExportTags returns the ExportTags field if non-nil, zero value otherwise.

### GetExportTagsOk

`func (o *BulkExportedItem) GetExportTagsOk() (*bool, bool)`

GetExportTagsOk returns a tuple with the ExportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTags

`func (o *BulkExportedItem) SetExportTags(v bool)`

SetExportTags sets ExportTags field to given value.

### HasExportTags

`func (o *BulkExportedItem) HasExportTags() bool`

HasExportTags returns a boolean if a field has been set.

### GetFileName

`func (o *BulkExportedItem) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BulkExportedItem) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BulkExportedItem) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BulkExportedItem) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetItem

`func (o *BulkExportedItem) GetItem() MoMoRef`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BulkExportedItem) GetItemOk() (*MoMoRef, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BulkExportedItem) SetItem(v MoMoRef)`

SetItem sets Item field to given value.

### HasItem

`func (o *BulkExportedItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetName

`func (o *BulkExportedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkExportedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkExportedItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkExportedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceName

`func (o *BulkExportedItem) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BulkExportedItem) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BulkExportedItem) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BulkExportedItem) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *BulkExportedItem) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *BulkExportedItem) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *BulkExportedItem) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *BulkExportedItem) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetStatus

`func (o *BulkExportedItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkExportedItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkExportedItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkExportedItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *BulkExportedItem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BulkExportedItem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BulkExportedItem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *BulkExportedItem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExport

`func (o *BulkExportedItem) GetExport() BulkExportRelationship`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *BulkExportedItem) GetExportOk() (*BulkExportRelationship, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *BulkExportedItem) SetExport(v BulkExportRelationship)`

SetExport sets Export field to given value.

### HasExport

`func (o *BulkExportedItem) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetParentItem

`func (o *BulkExportedItem) GetParentItem() BulkExportedItemRelationship`

GetParentItem returns the ParentItem field if non-nil, zero value otherwise.

### GetParentItemOk

`func (o *BulkExportedItem) GetParentItemOk() (*BulkExportedItemRelationship, bool)`

GetParentItemOk returns a tuple with the ParentItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItem

`func (o *BulkExportedItem) SetParentItem(v BulkExportedItemRelationship)`

SetParentItem sets ParentItem field to given value.

### HasParentItem

`func (o *BulkExportedItem) HasParentItem() bool`

HasParentItem returns a boolean if a field has been set.

### GetRelatedItems

`func (o *BulkExportedItem) GetRelatedItems() []BulkExportedItemRelationship`

GetRelatedItems returns the RelatedItems field if non-nil, zero value otherwise.

### GetRelatedItemsOk

`func (o *BulkExportedItem) GetRelatedItemsOk() (*[]BulkExportedItemRelationship, bool)`

GetRelatedItemsOk returns a tuple with the RelatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedItems

`func (o *BulkExportedItem) SetRelatedItems(v []BulkExportedItemRelationship)`

SetRelatedItems sets RelatedItems field to given value.

### HasRelatedItems

`func (o *BulkExportedItem) HasRelatedItems() bool`

HasRelatedItems returns a boolean if a field has been set.

### SetRelatedItemsNil

`func (o *BulkExportedItem) SetRelatedItemsNil(b bool)`

 SetRelatedItemsNil sets the value for RelatedItems to be an explicit nil

### UnsetRelatedItems
`func (o *BulkExportedItem) UnsetRelatedItems()`

UnsetRelatedItems ensures that no value is present for RelatedItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


