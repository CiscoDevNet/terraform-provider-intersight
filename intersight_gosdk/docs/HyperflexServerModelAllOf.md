# HyperflexServerModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerModel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerModel"]
**ServerModelEntries** | Pointer to [**[]HyperflexServerModelEntry**](HyperflexServerModelEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHyperflexServerModelAllOf

`func NewHyperflexServerModelAllOf(classId string, objectType string, ) *HyperflexServerModelAllOf`

NewHyperflexServerModelAllOf instantiates a new HyperflexServerModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerModelAllOfWithDefaults

`func NewHyperflexServerModelAllOfWithDefaults() *HyperflexServerModelAllOf`

NewHyperflexServerModelAllOfWithDefaults instantiates a new HyperflexServerModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerModelAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerModelAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerModelAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerModelAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerModelAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerModelAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServerModelEntries

`func (o *HyperflexServerModelAllOf) GetServerModelEntries() []HyperflexServerModelEntry`

GetServerModelEntries returns the ServerModelEntries field if non-nil, zero value otherwise.

### GetServerModelEntriesOk

`func (o *HyperflexServerModelAllOf) GetServerModelEntriesOk() (*[]HyperflexServerModelEntry, bool)`

GetServerModelEntriesOk returns a tuple with the ServerModelEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModelEntries

`func (o *HyperflexServerModelAllOf) SetServerModelEntries(v []HyperflexServerModelEntry)`

SetServerModelEntries sets ServerModelEntries field to given value.

### HasServerModelEntries

`func (o *HyperflexServerModelAllOf) HasServerModelEntries() bool`

HasServerModelEntries returns a boolean if a field has been set.

### SetServerModelEntriesNil

`func (o *HyperflexServerModelAllOf) SetServerModelEntriesNil(b bool)`

 SetServerModelEntriesNil sets the value for ServerModelEntries to be an explicit nil

### UnsetServerModelEntries
`func (o *HyperflexServerModelAllOf) UnsetServerModelEntries()`

UnsetServerModelEntries ensures that no value is present for ServerModelEntries, not even an explicit nil
### GetAppCatalog

`func (o *HyperflexServerModelAllOf) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexServerModelAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexServerModelAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexServerModelAllOf) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


