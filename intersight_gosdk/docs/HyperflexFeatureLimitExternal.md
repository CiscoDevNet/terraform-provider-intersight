# HyperflexFeatureLimitExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.FeatureLimitExternal"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.FeatureLimitExternal"]
**FeatureLimitEntries** | Pointer to [**[]HyperflexFeatureLimitEntry**](HyperflexFeatureLimitEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHyperflexFeatureLimitExternal

`func NewHyperflexFeatureLimitExternal(classId string, objectType string, ) *HyperflexFeatureLimitExternal`

NewHyperflexFeatureLimitExternal instantiates a new HyperflexFeatureLimitExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexFeatureLimitExternalWithDefaults

`func NewHyperflexFeatureLimitExternalWithDefaults() *HyperflexFeatureLimitExternal`

NewHyperflexFeatureLimitExternalWithDefaults instantiates a new HyperflexFeatureLimitExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexFeatureLimitExternal) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexFeatureLimitExternal) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexFeatureLimitExternal) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexFeatureLimitExternal) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexFeatureLimitExternal) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexFeatureLimitExternal) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureLimitEntries

`func (o *HyperflexFeatureLimitExternal) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry`

GetFeatureLimitEntries returns the FeatureLimitEntries field if non-nil, zero value otherwise.

### GetFeatureLimitEntriesOk

`func (o *HyperflexFeatureLimitExternal) GetFeatureLimitEntriesOk() (*[]HyperflexFeatureLimitEntry, bool)`

GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitEntries

`func (o *HyperflexFeatureLimitExternal) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry)`

SetFeatureLimitEntries sets FeatureLimitEntries field to given value.

### HasFeatureLimitEntries

`func (o *HyperflexFeatureLimitExternal) HasFeatureLimitEntries() bool`

HasFeatureLimitEntries returns a boolean if a field has been set.

### SetFeatureLimitEntriesNil

`func (o *HyperflexFeatureLimitExternal) SetFeatureLimitEntriesNil(b bool)`

 SetFeatureLimitEntriesNil sets the value for FeatureLimitEntries to be an explicit nil

### UnsetFeatureLimitEntries
`func (o *HyperflexFeatureLimitExternal) UnsetFeatureLimitEntries()`

UnsetFeatureLimitEntries ensures that no value is present for FeatureLimitEntries, not even an explicit nil
### GetAppCatalog

`func (o *HyperflexFeatureLimitExternal) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexFeatureLimitExternal) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexFeatureLimitExternal) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexFeatureLimitExternal) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


