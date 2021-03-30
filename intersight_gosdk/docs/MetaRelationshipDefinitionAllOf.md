# MetaRelationshipDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.RelationshipDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.RelationshipDefinition"]
**ApiAccess** | Pointer to **string** | API access definition for this relationship. * &#x60;NoAccess&#x60; - The property is not accessible from the API. * &#x60;ReadOnly&#x60; - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * &#x60;CreateOnly&#x60; - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * &#x60;ReadWrite&#x60; - The property has read/write access. * &#x60;WriteOnly&#x60; - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * &#x60;ReadOnCreate&#x60; - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created. | [optional] [readonly] [default to "NoAccess"]
**Collection** | Pointer to **bool** | Specifies whether the relationship is a collection. | [optional] [readonly] 
**Export** | Pointer to **bool** | When turned off, the peer MO is not exported when the local MO is exported. | [optional] [readonly] 
**ExportWithPeer** | Pointer to **bool** | When turned on, the local MO is exported when the peer is exported. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the relationship. | [optional] [readonly] 
**PeerRelName** | Pointer to **string** | Name of relationship in peer managed object. | [optional] [readonly] 
**PeerSync** | Pointer to **bool** | When turned on, peer MO corresponding to the reference provided in relation is updated with a reference to the current MO. | [optional] [readonly] 
**Type** | Pointer to **string** | Fully qualified type of the peer managed object. | [optional] [readonly] 

## Methods

### NewMetaRelationshipDefinitionAllOf

`func NewMetaRelationshipDefinitionAllOf(classId string, objectType string, ) *MetaRelationshipDefinitionAllOf`

NewMetaRelationshipDefinitionAllOf instantiates a new MetaRelationshipDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaRelationshipDefinitionAllOfWithDefaults

`func NewMetaRelationshipDefinitionAllOfWithDefaults() *MetaRelationshipDefinitionAllOf`

NewMetaRelationshipDefinitionAllOfWithDefaults instantiates a new MetaRelationshipDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaRelationshipDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaRelationshipDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaRelationshipDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaRelationshipDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaRelationshipDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaRelationshipDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApiAccess

`func (o *MetaRelationshipDefinitionAllOf) GetApiAccess() string`

GetApiAccess returns the ApiAccess field if non-nil, zero value otherwise.

### GetApiAccessOk

`func (o *MetaRelationshipDefinitionAllOf) GetApiAccessOk() (*string, bool)`

GetApiAccessOk returns a tuple with the ApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccess

`func (o *MetaRelationshipDefinitionAllOf) SetApiAccess(v string)`

SetApiAccess sets ApiAccess field to given value.

### HasApiAccess

`func (o *MetaRelationshipDefinitionAllOf) HasApiAccess() bool`

HasApiAccess returns a boolean if a field has been set.

### GetCollection

`func (o *MetaRelationshipDefinitionAllOf) GetCollection() bool`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *MetaRelationshipDefinitionAllOf) GetCollectionOk() (*bool, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *MetaRelationshipDefinitionAllOf) SetCollection(v bool)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *MetaRelationshipDefinitionAllOf) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetExport

`func (o *MetaRelationshipDefinitionAllOf) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *MetaRelationshipDefinitionAllOf) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *MetaRelationshipDefinitionAllOf) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *MetaRelationshipDefinitionAllOf) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetExportWithPeer

`func (o *MetaRelationshipDefinitionAllOf) GetExportWithPeer() bool`

GetExportWithPeer returns the ExportWithPeer field if non-nil, zero value otherwise.

### GetExportWithPeerOk

`func (o *MetaRelationshipDefinitionAllOf) GetExportWithPeerOk() (*bool, bool)`

GetExportWithPeerOk returns a tuple with the ExportWithPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportWithPeer

`func (o *MetaRelationshipDefinitionAllOf) SetExportWithPeer(v bool)`

SetExportWithPeer sets ExportWithPeer field to given value.

### HasExportWithPeer

`func (o *MetaRelationshipDefinitionAllOf) HasExportWithPeer() bool`

HasExportWithPeer returns a boolean if a field has been set.

### GetName

`func (o *MetaRelationshipDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaRelationshipDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaRelationshipDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaRelationshipDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeerRelName

`func (o *MetaRelationshipDefinitionAllOf) GetPeerRelName() string`

GetPeerRelName returns the PeerRelName field if non-nil, zero value otherwise.

### GetPeerRelNameOk

`func (o *MetaRelationshipDefinitionAllOf) GetPeerRelNameOk() (*string, bool)`

GetPeerRelNameOk returns a tuple with the PeerRelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRelName

`func (o *MetaRelationshipDefinitionAllOf) SetPeerRelName(v string)`

SetPeerRelName sets PeerRelName field to given value.

### HasPeerRelName

`func (o *MetaRelationshipDefinitionAllOf) HasPeerRelName() bool`

HasPeerRelName returns a boolean if a field has been set.

### GetPeerSync

`func (o *MetaRelationshipDefinitionAllOf) GetPeerSync() bool`

GetPeerSync returns the PeerSync field if non-nil, zero value otherwise.

### GetPeerSyncOk

`func (o *MetaRelationshipDefinitionAllOf) GetPeerSyncOk() (*bool, bool)`

GetPeerSyncOk returns a tuple with the PeerSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSync

`func (o *MetaRelationshipDefinitionAllOf) SetPeerSync(v bool)`

SetPeerSync sets PeerSync field to given value.

### HasPeerSync

`func (o *MetaRelationshipDefinitionAllOf) HasPeerSync() bool`

HasPeerSync returns a boolean if a field has been set.

### GetType

`func (o *MetaRelationshipDefinitionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaRelationshipDefinitionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaRelationshipDefinitionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaRelationshipDefinitionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


