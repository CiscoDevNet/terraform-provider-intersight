# ConfigImportedItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "config.ImportedItem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "config.ImportedItem"]
**IsShared** | Pointer to **bool** | Specifies whether this item MO was in shared scope or user scope when exported. | [optional] [readonly] 
**IsUpdated** | Pointer to **bool** | Specifies whether this item MO was updated or created while importing in desired service. | [optional] [readonly] 
**Item** | Pointer to [**NullableConfigMoRef**](ConfigMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | MO item identity (the moref corresponding to item) expressed as a string. | [optional] [readonly] 
**NewMoid** | Pointer to **string** | Moid of the MO created/updated during import for the item. | [optional] [readonly] 
**ServiceVersion** | Pointer to **string** | Version of the service that owned the item MO when the item was exported. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the item&#39;s import operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Progress or error message for the MO&#39;s import operation. | [optional] [readonly] 
**Importer** | Pointer to [**ConfigImporterRelationship**](ConfigImporterRelationship.md) |  | [optional] 

## Methods

### NewConfigImportedItemAllOf

`func NewConfigImportedItemAllOf(classId string, objectType string, ) *ConfigImportedItemAllOf`

NewConfigImportedItemAllOf instantiates a new ConfigImportedItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigImportedItemAllOfWithDefaults

`func NewConfigImportedItemAllOfWithDefaults() *ConfigImportedItemAllOf`

NewConfigImportedItemAllOfWithDefaults instantiates a new ConfigImportedItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConfigImportedItemAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConfigImportedItemAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConfigImportedItemAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConfigImportedItemAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigImportedItemAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigImportedItemAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsShared

`func (o *ConfigImportedItemAllOf) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *ConfigImportedItemAllOf) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *ConfigImportedItemAllOf) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *ConfigImportedItemAllOf) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetIsUpdated

`func (o *ConfigImportedItemAllOf) GetIsUpdated() bool`

GetIsUpdated returns the IsUpdated field if non-nil, zero value otherwise.

### GetIsUpdatedOk

`func (o *ConfigImportedItemAllOf) GetIsUpdatedOk() (*bool, bool)`

GetIsUpdatedOk returns a tuple with the IsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpdated

`func (o *ConfigImportedItemAllOf) SetIsUpdated(v bool)`

SetIsUpdated sets IsUpdated field to given value.

### HasIsUpdated

`func (o *ConfigImportedItemAllOf) HasIsUpdated() bool`

HasIsUpdated returns a boolean if a field has been set.

### GetItem

`func (o *ConfigImportedItemAllOf) GetItem() ConfigMoRef`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ConfigImportedItemAllOf) GetItemOk() (*ConfigMoRef, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ConfigImportedItemAllOf) SetItem(v ConfigMoRef)`

SetItem sets Item field to given value.

### HasItem

`func (o *ConfigImportedItemAllOf) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *ConfigImportedItemAllOf) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *ConfigImportedItemAllOf) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetName

`func (o *ConfigImportedItemAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigImportedItemAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigImportedItemAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigImportedItemAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewMoid

`func (o *ConfigImportedItemAllOf) GetNewMoid() string`

GetNewMoid returns the NewMoid field if non-nil, zero value otherwise.

### GetNewMoidOk

`func (o *ConfigImportedItemAllOf) GetNewMoidOk() (*string, bool)`

GetNewMoidOk returns a tuple with the NewMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMoid

`func (o *ConfigImportedItemAllOf) SetNewMoid(v string)`

SetNewMoid sets NewMoid field to given value.

### HasNewMoid

`func (o *ConfigImportedItemAllOf) HasNewMoid() bool`

HasNewMoid returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ConfigImportedItemAllOf) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ConfigImportedItemAllOf) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ConfigImportedItemAllOf) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ConfigImportedItemAllOf) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigImportedItemAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigImportedItemAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigImportedItemAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigImportedItemAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigImportedItemAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigImportedItemAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigImportedItemAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigImportedItemAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetImporter

`func (o *ConfigImportedItemAllOf) GetImporter() ConfigImporterRelationship`

GetImporter returns the Importer field if non-nil, zero value otherwise.

### GetImporterOk

`func (o *ConfigImportedItemAllOf) GetImporterOk() (*ConfigImporterRelationship, bool)`

GetImporterOk returns a tuple with the Importer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporter

`func (o *ConfigImportedItemAllOf) SetImporter(v ConfigImporterRelationship)`

SetImporter sets Importer field to given value.

### HasImporter

`func (o *ConfigImportedItemAllOf) HasImporter() bool`

HasImporter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


