# HyperflexServerFirmwareVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.ServerFirmwareVersion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.ServerFirmwareVersion"]
**AppCatalog** | Pointer to [**HyperflexAppCatalogRelationship**](HyperflexAppCatalogRelationship.md) |  | [optional] 
**ServerFirmwareVersionEntries** | Pointer to [**[]HyperflexServerFirmwareVersionEntryRelationship**](HyperflexServerFirmwareVersionEntryRelationship.md) | An array of relationships to hyperflexServerFirmwareVersionEntry resources. | [optional] 

## Methods

### NewHyperflexServerFirmwareVersion

`func NewHyperflexServerFirmwareVersion(classId string, objectType string, ) *HyperflexServerFirmwareVersion`

NewHyperflexServerFirmwareVersion instantiates a new HyperflexServerFirmwareVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexServerFirmwareVersionWithDefaults

`func NewHyperflexServerFirmwareVersionWithDefaults() *HyperflexServerFirmwareVersion`

NewHyperflexServerFirmwareVersionWithDefaults instantiates a new HyperflexServerFirmwareVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexServerFirmwareVersion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexServerFirmwareVersion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexServerFirmwareVersion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexServerFirmwareVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexServerFirmwareVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexServerFirmwareVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntries() []HyperflexServerFirmwareVersionEntryRelationship`

GetServerFirmwareVersionEntries returns the ServerFirmwareVersionEntries field if non-nil, zero value otherwise.

### GetServerFirmwareVersionEntriesOk

`func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntriesOk() (*[]HyperflexServerFirmwareVersionEntryRelationship, bool)`

GetServerFirmwareVersionEntriesOk returns a tuple with the ServerFirmwareVersionEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) SetServerFirmwareVersionEntries(v []HyperflexServerFirmwareVersionEntryRelationship)`

SetServerFirmwareVersionEntries sets ServerFirmwareVersionEntries field to given value.

### HasServerFirmwareVersionEntries

`func (o *HyperflexServerFirmwareVersion) HasServerFirmwareVersionEntries() bool`

HasServerFirmwareVersionEntries returns a boolean if a field has been set.

### SetServerFirmwareVersionEntriesNil

`func (o *HyperflexServerFirmwareVersion) SetServerFirmwareVersionEntriesNil(b bool)`

 SetServerFirmwareVersionEntriesNil sets the value for ServerFirmwareVersionEntries to be an explicit nil

### UnsetServerFirmwareVersionEntries
`func (o *HyperflexServerFirmwareVersion) UnsetServerFirmwareVersionEntries()`

UnsetServerFirmwareVersionEntries ensures that no value is present for ServerFirmwareVersionEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


