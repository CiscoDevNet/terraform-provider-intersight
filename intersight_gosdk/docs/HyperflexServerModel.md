# HyperflexServerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerModelEntries** | Pointer to [**[]HyperflexServerModelEntry**](hyperflex.ServerModelEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](hyperflex.AppCatalog.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexServerModel

`func NewHyperflexServerModel() *HyperflexServerModel`

NewHyperflexServerModel instantiates a new HyperflexServerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerModelWithDefaults

`func NewHyperflexServerModelWithDefaults() *HyperflexServerModel`

NewHyperflexServerModelWithDefaults instantiates a new HyperflexServerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerModelEntries

`func (o *HyperflexServerModel) GetServerModelEntries() []HyperflexServerModelEntry`

GetServerModelEntries returns the ServerModelEntries field if non-nil, zero value otherwise.

### GetServerModelEntriesOk

`func (o *HyperflexServerModel) GetServerModelEntriesOk() (*[]HyperflexServerModelEntry, bool)`

GetServerModelEntriesOk returns a tuple with the ServerModelEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModelEntries

`func (o *HyperflexServerModel) SetServerModelEntries(v []HyperflexServerModelEntry)`

SetServerModelEntries sets ServerModelEntries field to given value.

### HasServerModelEntries

`func (o *HyperflexServerModel) HasServerModelEntries() bool`

HasServerModelEntries returns a boolean if a field has been set.

### GetAppCatalog

`func (o *HyperflexServerModel) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexServerModel) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexServerModel) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexServerModel) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


