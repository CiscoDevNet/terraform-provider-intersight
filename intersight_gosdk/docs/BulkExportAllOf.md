# BulkExportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.Export"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.Export"]
**Action** | Pointer to **string** | Action to be performed on the export operation. * &#x60;Start&#x60; - Starts the export operation. * &#x60;Cancel&#x60; - Cancels the export operation that is in progress. | [optional] [default to "Start"]
**ExportTags** | Pointer to **bool** | Specifies whether tags must be exported and will be considered for all the items MOs. | [optional] [default to true]
**ExportedObjects** | Pointer to [**[]BulkSubRequest**](BulkSubRequest.md) |  | [optional] 
**ImportOrder** | Pointer to **interface{}** | Contains the list of import order. | [optional] [readonly] 
**Items** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). | [optional] 
**Status** | Pointer to **string** | Status of the export operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;OrderInProgress&#x60; - The archive operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;OperationTimedOut&#x60; - The operation has timed out. * &#x60;OperationCancelled&#x60; - The operation has been cancelled. * &#x60;CancelInProgress&#x60; - The operation is being cancelled. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**ExportedItems** | Pointer to [**[]BulkExportedItemRelationship**](BulkExportedItemRelationship.md) | An array of relationships to bulkExportedItem resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewBulkExportAllOf

`func NewBulkExportAllOf(classId string, objectType string, ) *BulkExportAllOf`

NewBulkExportAllOf instantiates a new BulkExportAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkExportAllOfWithDefaults

`func NewBulkExportAllOfWithDefaults() *BulkExportAllOf`

NewBulkExportAllOfWithDefaults instantiates a new BulkExportAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkExportAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkExportAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkExportAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkExportAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkExportAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkExportAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *BulkExportAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkExportAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkExportAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkExportAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExportTags

`func (o *BulkExportAllOf) GetExportTags() bool`

GetExportTags returns the ExportTags field if non-nil, zero value otherwise.

### GetExportTagsOk

`func (o *BulkExportAllOf) GetExportTagsOk() (*bool, bool)`

GetExportTagsOk returns a tuple with the ExportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTags

`func (o *BulkExportAllOf) SetExportTags(v bool)`

SetExportTags sets ExportTags field to given value.

### HasExportTags

`func (o *BulkExportAllOf) HasExportTags() bool`

HasExportTags returns a boolean if a field has been set.

### GetExportedObjects

`func (o *BulkExportAllOf) GetExportedObjects() []BulkSubRequest`

GetExportedObjects returns the ExportedObjects field if non-nil, zero value otherwise.

### GetExportedObjectsOk

`func (o *BulkExportAllOf) GetExportedObjectsOk() (*[]BulkSubRequest, bool)`

GetExportedObjectsOk returns a tuple with the ExportedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedObjects

`func (o *BulkExportAllOf) SetExportedObjects(v []BulkSubRequest)`

SetExportedObjects sets ExportedObjects field to given value.

### HasExportedObjects

`func (o *BulkExportAllOf) HasExportedObjects() bool`

HasExportedObjects returns a boolean if a field has been set.

### SetExportedObjectsNil

`func (o *BulkExportAllOf) SetExportedObjectsNil(b bool)`

 SetExportedObjectsNil sets the value for ExportedObjects to be an explicit nil

### UnsetExportedObjects
`func (o *BulkExportAllOf) UnsetExportedObjects()`

UnsetExportedObjects ensures that no value is present for ExportedObjects, not even an explicit nil
### GetImportOrder

`func (o *BulkExportAllOf) GetImportOrder() interface{}`

GetImportOrder returns the ImportOrder field if non-nil, zero value otherwise.

### GetImportOrderOk

`func (o *BulkExportAllOf) GetImportOrderOk() (*interface{}, bool)`

GetImportOrderOk returns a tuple with the ImportOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOrder

`func (o *BulkExportAllOf) SetImportOrder(v interface{})`

SetImportOrder sets ImportOrder field to given value.

### HasImportOrder

`func (o *BulkExportAllOf) HasImportOrder() bool`

HasImportOrder returns a boolean if a field has been set.

### SetImportOrderNil

`func (o *BulkExportAllOf) SetImportOrderNil(b bool)`

 SetImportOrderNil sets the value for ImportOrder to be an explicit nil

### UnsetImportOrder
`func (o *BulkExportAllOf) UnsetImportOrder()`

UnsetImportOrder ensures that no value is present for ImportOrder, not even an explicit nil
### GetItems

`func (o *BulkExportAllOf) GetItems() []MoMoRef`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkExportAllOf) GetItemsOk() (*[]MoMoRef, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkExportAllOf) SetItems(v []MoMoRef)`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkExportAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BulkExportAllOf) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BulkExportAllOf) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetName

`func (o *BulkExportAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkExportAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkExportAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkExportAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BulkExportAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkExportAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkExportAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkExportAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *BulkExportAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BulkExportAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BulkExportAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *BulkExportAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExportedItems

`func (o *BulkExportAllOf) GetExportedItems() []BulkExportedItemRelationship`

GetExportedItems returns the ExportedItems field if non-nil, zero value otherwise.

### GetExportedItemsOk

`func (o *BulkExportAllOf) GetExportedItemsOk() (*[]BulkExportedItemRelationship, bool)`

GetExportedItemsOk returns a tuple with the ExportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedItems

`func (o *BulkExportAllOf) SetExportedItems(v []BulkExportedItemRelationship)`

SetExportedItems sets ExportedItems field to given value.

### HasExportedItems

`func (o *BulkExportAllOf) HasExportedItems() bool`

HasExportedItems returns a boolean if a field has been set.

### SetExportedItemsNil

`func (o *BulkExportAllOf) SetExportedItemsNil(b bool)`

 SetExportedItemsNil sets the value for ExportedItems to be an explicit nil

### UnsetExportedItems
`func (o *BulkExportAllOf) UnsetExportedItems()`

UnsetExportedItems ensures that no value is present for ExportedItems, not even an explicit nil
### GetOrganization

`func (o *BulkExportAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkExportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkExportAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkExportAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


