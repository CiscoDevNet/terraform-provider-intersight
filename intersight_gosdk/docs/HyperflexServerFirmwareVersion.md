# HyperflexServerFirmwareVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareVersionEntries** | Pointer to [**[]HyperflexServerFirmwareVersionEntry**](hyperflex.ServerFirmwareVersionEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](hyperflex.AppCatalog.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexServerFirmwareVersion

`func NewHyperflexServerFirmwareVersion() *HyperflexServerFirmwareVersion`

NewHyperflexServerFirmwareVersion instantiates a new HyperflexServerFirmwareVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionWithDefaults

`func NewHyperflexServerFirmwareVersionWithDefaults() *HyperflexServerFirmwareVersion`

NewHyperflexServerFirmwareVersionWithDefaults instantiates a new HyperflexServerFirmwareVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntries() []HyperflexServerFirmwareVersionEntry`

GetServerFirmwareVersionEntries returns the ServerFirmwareVersionEntries field if non-nil, zero value otherwise.

### GetServerFirmwareVersionEntriesOk

`func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntriesOk() (*[]HyperflexServerFirmwareVersionEntry, bool)`

GetServerFirmwareVersionEntriesOk returns a tuple with the ServerFirmwareVersionEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) SetServerFirmwareVersionEntries(v []HyperflexServerFirmwareVersionEntry)`

SetServerFirmwareVersionEntries sets ServerFirmwareVersionEntries field to given value.

### HasServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) HasServerFirmwareVersionEntries() bool`

HasServerFirmwareVersionEntries returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HyperflexServerFirmwareVersion) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexServerFirmwareVersion) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexServerFirmwareVersion) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexServerFirmwareVersion) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


