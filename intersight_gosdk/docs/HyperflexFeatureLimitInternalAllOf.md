# HyperflexFeatureLimitInternalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.FeatureLimitInternal"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.FeatureLimitInternal"]
**FeatureLimitEntries** | Pointer to [**[]HyperflexFeatureLimitEntry**](HyperflexFeatureLimitEntry.md) |  | [optional] 
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 

## Methods

### NewHyperflexFeatureLimitInternalAllOf

`func NewHyperflexFeatureLimitInternalAllOf(classId string, objectType string, ) *HyperflexFeatureLimitInternalAllOf`

NewHyperflexFeatureLimitInternalAllOf instantiates a new HyperflexFeatureLimitInternalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexFeatureLimitInternalAllOfWithDefaults

`func NewHyperflexFeatureLimitInternalAllOfWithDefaults() *HyperflexFeatureLimitInternalAllOf`

NewHyperflexFeatureLimitInternalAllOfWithDefaults instantiates a new HyperflexFeatureLimitInternalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexFeatureLimitInternalAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexFeatureLimitInternalAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexFeatureLimitInternalAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexFeatureLimitInternalAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexFeatureLimitInternalAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexFeatureLimitInternalAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFeatureLimitEntries

`func (o *HyperflexFeatureLimitInternalAllOf) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry`

GetFeatureLimitEntries returns the FeatureLimitEntries field if non-nil, zero value otherwise.

### GetFeatureLimitEntriesOk

`func (o *HyperflexFeatureLimitInternalAllOf) GetFeatureLimitEntriesOk() (*[]HyperflexFeatureLimitEntry, bool)`

GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitEntries

`func (o *HyperflexFeatureLimitInternalAllOf) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry)`

SetFeatureLimitEntries sets FeatureLimitEntries field to given value.

### HasFeatureLimitEntries

`func (o *HyperflexFeatureLimitInternalAllOf) HasFeatureLimitEntries() bool`

HasFeatureLimitEntries returns a boolean if a field has been set.

### SetFeatureLimitEntriesNil

`func (o *HyperflexFeatureLimitInternalAllOf) SetFeatureLimitEntriesNil(b bool)`

 SetFeatureLimitEntriesNil sets the value for FeatureLimitEntries to be an explicit nil

### UnsetFeatureLimitEntries
`func (o *HyperflexFeatureLimitInternalAllOf) UnsetFeatureLimitEntries()`

UnsetFeatureLimitEntries ensures that no value is present for FeatureLimitEntries, not even an explicit nil
### GetAppCatalog

`func (o *HyperflexFeatureLimitInternalAllOf) GetAppCatalog() HyperflexAppCatalogRelationship`

GetAppCatalog returns the AppCatalog field if non-nil, zero value otherwise.

### GetAppCatalogOk

`func (o *HyperflexFeatureLimitInternalAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool)`

GetAppCatalogOk returns a tuple with the AppCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCatalog

`func (o *HyperflexFeatureLimitInternalAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship)`

SetAppCatalog sets AppCatalog field to given value.

### HasAppCatalog

`func (o *HyperflexFeatureLimitInternalAllOf) HasAppCatalog() bool`

HasAppCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


