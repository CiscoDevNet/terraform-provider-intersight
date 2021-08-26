# HyperflexServerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerModel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerModel"]
**ServerModelEntries** | Pointer to [**[]HyperflexServerModelEntry**](HyperflexServerModelEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHyperflexServerModel

`func NewHyperflexServerModel(classId string, objectType string, ) *HyperflexServerModel`

NewHyperflexServerModel instantiates a new HyperflexServerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerModelWithDefaults

`func NewHyperflexServerModelWithDefaults() *HyperflexServerModel`

NewHyperflexServerModelWithDefaults instantiates a new HyperflexServerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerModel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerModel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerModel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerModel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerModel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerModel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetServerModelEntriesNil

`func (o *HyperflexServerModel) SetServerModelEntriesNil(b bool)`

 SetServerModelEntriesNil sets the value for ServerModelEntries to be an explicit nil

### UnsetServerModelEntries
`func (o *HyperflexServerModel) UnsetServerModelEntries()`

UnsetServerModelEntries ensures that no value is present for ServerModelEntries, not even an explicit nil
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


