# MetaRelationshipDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAccess** | Pointer to **string** | API access definition for this relationship. * &#x60;NoAccess&#x60; - The property is not accessible from the API. * &#x60;ReadOnly&#x60; - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * &#x60;CreateOnly&#x60; - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * &#x60;ReadWrite&#x60; - The property has read/write access. * &#x60;WriteOnly&#x60; - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * &#x60;ReadOnCreate&#x60; - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created. | [optional] [readonly] [default to "NoAccess"]
**Collection** | Pointer to **bool** | Specifies whether the relationship is a collection. | [optional] [readonly] 
**Export** | Pointer to **bool** | When turned off, the peer MO is not exported when the local MO is exported. | [optional] [readonly] 
**ExportWithPeer** | Pointer to **bool** | When turned on, the local MO is exported when the peer is exported. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the relationship. | [optional] [readonly] 
**Type** | Pointer to **string** | Fully qualified type of the foreign managed object. | [optional] [readonly] 

## Methods

### NewMetaRelationshipDefinition

`func NewMetaRelationshipDefinition() *MetaRelationshipDefinition`

NewMetaRelationshipDefinition instantiates a new MetaRelationshipDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaRelationshipDefinitionWithDefaults

`func NewMetaRelationshipDefinitionWithDefaults() *MetaRelationshipDefinition`

NewMetaRelationshipDefinitionWithDefaults instantiates a new MetaRelationshipDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAccess

`func (o *MetaRelationshipDefinition) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaRelationshipDefinition) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaRelationshipDefinition) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaRelationshipDefinition) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetCollection

`func (o *MetaRelationshipDefinition) GetCollection() bool`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *MetaRelationshipDefinition) GetCollectionOk() (*bool, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *MetaRelationshipDefinition) SetCollection(v bool)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *MetaRelationshipDefinition) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetExport

`func (o *MetaRelationshipDefinition) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *MetaRelationshipDefinition) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *MetaRelationshipDefinition) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *MetaRelationshipDefinition) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetExportWithPeer

`func (o *MetaRelationshipDefinition) GetExportWithPeer() bool`

GetExportWithPeer returns the ExportWithPeer field if non-nil, zero value otherwise.

### GetExportWithPeerOk

`func (o *MetaRelationshipDefinition) GetExportWithPeerOk() (*bool, bool)`

GetExportWithPeerOk returns a tuple with the ExportWithPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportWithPeer

`func (o *MetaRelationshipDefinition) SetExportWithPeer(v bool)`

SetExportWithPeer sets ExportWithPeer field to given value.

### HasExportWithPeer

`func (o *MetaRelationshipDefinition) HasExportWithPeer() bool`

HasExportWithPeer returns a boolean if a field has been set.

### GetName

`func (o *MetaRelationshipDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaRelationshipDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaRelationshipDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaRelationshipDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MetaRelationshipDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaRelationshipDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaRelationshipDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaRelationshipDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


