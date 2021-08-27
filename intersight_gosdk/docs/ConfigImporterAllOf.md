# ConfigImporterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "config.Importer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "config.Importer"]
**ImportPath** | Pointer to **string** | The path to the archive in Intersight storage that has all the MOs to be imported. | [optional] 
**ImportSource** | Pointer to **string** | The source of the archive in Intersight storage that has all the MOs to be imported. * &#x60;ImageRepo&#x60; - The &#39;ImageRepo&#39; source if the source of exporter archive is image repository. * &#x60;URL&#x60; - The &#39;URL&#39; source if the source of exported archive is a URL. | [optional] [default to "ImageRepo"]
**Name** | Pointer to **string** | An identifier for the importer instance. | [optional] 
**SkipIntegrityChecks** | Pointer to **bool** | Specifies whether integrity checks must be skipped. | [optional] [default to false]
**Status** | Pointer to **string** | Status of the import operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**ImportedItems** | Pointer to [**[]ConfigImportedItemRelationship**](ConfigImportedItemRelationship.md) | An array of relationships to configImportedItem resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewConfigImporterAllOf

`func NewConfigImporterAllOf(classId string, objectType string, ) *ConfigImporterAllOf`

NewConfigImporterAllOf instantiates a new ConfigImporterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigImporterAllOfWithDefaults

`func NewConfigImporterAllOfWithDefaults() *ConfigImporterAllOf`

NewConfigImporterAllOfWithDefaults instantiates a new ConfigImporterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConfigImporterAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConfigImporterAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConfigImporterAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConfigImporterAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigImporterAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigImporterAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImportPath

`func (o *ConfigImporterAllOf) GetImportPath() string`

GetImportPath returns the ImportPath field if non-nil, zero value otherwise.

### GetImportPathOk

`func (o *ConfigImporterAllOf) GetImportPathOk() (*string, bool)`

GetImportPathOk returns a tuple with the ImportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPath

`func (o *ConfigImporterAllOf) SetImportPath(v string)`

SetImportPath sets ImportPath field to given value.

### HasImportPath

`func (o *ConfigImporterAllOf) HasImportPath() bool`

HasImportPath returns a boolean if a field has been set.

### GetImportSource

`func (o *ConfigImporterAllOf) GetImportSource() string`

GetImportSource returns the ImportSource field if non-nil, zero value otherwise.

### GetImportSourceOk

`func (o *ConfigImporterAllOf) GetImportSourceOk() (*string, bool)`

GetImportSourceOk returns a tuple with the ImportSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSource

`func (o *ConfigImporterAllOf) SetImportSource(v string)`

SetImportSource sets ImportSource field to given value.

### HasImportSource

`func (o *ConfigImporterAllOf) HasImportSource() bool`

HasImportSource returns a boolean if a field has been set.

### GetName

`func (o *ConfigImporterAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigImporterAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigImporterAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigImporterAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkipIntegrityChecks

`func (o *ConfigImporterAllOf) GetSkipIntegrityChecks() bool`

GetSkipIntegrityChecks returns the SkipIntegrityChecks field if non-nil, zero value otherwise.

### GetSkipIntegrityChecksOk

`func (o *ConfigImporterAllOf) GetSkipIntegrityChecksOk() (*bool, bool)`

GetSkipIntegrityChecksOk returns a tuple with the SkipIntegrityChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntegrityChecks

`func (o *ConfigImporterAllOf) SetSkipIntegrityChecks(v bool)`

SetSkipIntegrityChecks sets SkipIntegrityChecks field to given value.

### HasSkipIntegrityChecks

`func (o *ConfigImporterAllOf) HasSkipIntegrityChecks() bool`

HasSkipIntegrityChecks returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigImporterAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigImporterAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigImporterAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigImporterAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigImporterAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigImporterAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigImporterAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigImporterAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetImportedItems

`func (o *ConfigImporterAllOf) GetImportedItems() []ConfigImportedItemRelationship`

GetImportedItems returns the ImportedItems field if non-nil, zero value otherwise.

### GetImportedItemsOk

`func (o *ConfigImporterAllOf) GetImportedItemsOk() (*[]ConfigImportedItemRelationship, bool)`

GetImportedItemsOk returns a tuple with the ImportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedItems

`func (o *ConfigImporterAllOf) SetImportedItems(v []ConfigImportedItemRelationship)`

SetImportedItems sets ImportedItems field to given value.

### HasImportedItems

`func (o *ConfigImporterAllOf) HasImportedItems() bool`

HasImportedItems returns a boolean if a field has been set.

### SetImportedItemsNil

`func (o *ConfigImporterAllOf) SetImportedItemsNil(b bool)`

 SetImportedItemsNil sets the value for ImportedItems to be an explicit nil

### UnsetImportedItems
`func (o *ConfigImporterAllOf) UnsetImportedItems()`

UnsetImportedItems ensures that no value is present for ImportedItems, not even an explicit nil
### GetOrganization

`func (o *ConfigImporterAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConfigImporterAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConfigImporterAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConfigImporterAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


