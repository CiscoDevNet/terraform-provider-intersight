# BulkRelatedTypeExportOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.RelatedTypeExportOption"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.RelatedTypeExportOption"]
**ExcludeRelations** | Pointer to **bool** | Do not export relationships. | [optional] 
**ExcludedPeers** | Pointer to **[]string** |  | [optional] 
**RelatedType** | Pointer to **string** | The fully qualified related type name. | [optional] 

## Methods

### NewBulkRelatedTypeExportOption

`func NewBulkRelatedTypeExportOption(classId string, objectType string, ) *BulkRelatedTypeExportOption`

NewBulkRelatedTypeExportOption instantiates a new BulkRelatedTypeExportOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRelatedTypeExportOptionWithDefaults

`func NewBulkRelatedTypeExportOptionWithDefaults() *BulkRelatedTypeExportOption`

NewBulkRelatedTypeExportOptionWithDefaults instantiates a new BulkRelatedTypeExportOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkRelatedTypeExportOption) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkRelatedTypeExportOption) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkRelatedTypeExportOption) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkRelatedTypeExportOption) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkRelatedTypeExportOption) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkRelatedTypeExportOption) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExcludeRelations

`func (o *BulkRelatedTypeExportOption) GetExcludeRelations() bool`

GetExcludeRelations returns the ExcludeRelations field if non-nil, zero value otherwise.

### GetExcludeRelationsOk

`func (o *BulkRelatedTypeExportOption) GetExcludeRelationsOk() (*bool, bool)`

GetExcludeRelationsOk returns a tuple with the ExcludeRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRelations

`func (o *BulkRelatedTypeExportOption) SetExcludeRelations(v bool)`

SetExcludeRelations sets ExcludeRelations field to given value.

### HasExcludeRelations

`func (o *BulkRelatedTypeExportOption) HasExcludeRelations() bool`

HasExcludeRelations returns a boolean if a field has been set.

### GetExcludedPeers

`func (o *BulkRelatedTypeExportOption) GetExcludedPeers() []string`

GetExcludedPeers returns the ExcludedPeers field if non-nil, zero value otherwise.

### GetExcludedPeersOk

`func (o *BulkRelatedTypeExportOption) GetExcludedPeersOk() (*[]string, bool)`

GetExcludedPeersOk returns a tuple with the ExcludedPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPeers

`func (o *BulkRelatedTypeExportOption) SetExcludedPeers(v []string)`

SetExcludedPeers sets ExcludedPeers field to given value.

### HasExcludedPeers

`func (o *BulkRelatedTypeExportOption) HasExcludedPeers() bool`

HasExcludedPeers returns a boolean if a field has been set.

### SetExcludedPeersNil

`func (o *BulkRelatedTypeExportOption) SetExcludedPeersNil(b bool)`

 SetExcludedPeersNil sets the value for ExcludedPeers to be an explicit nil

### UnsetExcludedPeers
`func (o *BulkRelatedTypeExportOption) UnsetExcludedPeers()`

UnsetExcludedPeers ensures that no value is present for ExcludedPeers, not even an explicit nil
### GetRelatedType

`func (o *BulkRelatedTypeExportOption) GetRelatedType() string`

GetRelatedType returns the RelatedType field if non-nil, zero value otherwise.

### GetRelatedTypeOk

`func (o *BulkRelatedTypeExportOption) GetRelatedTypeOk() (*string, bool)`

GetRelatedTypeOk returns a tuple with the RelatedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedType

`func (o *BulkRelatedTypeExportOption) SetRelatedType(v string)`

SetRelatedType sets RelatedType field to given value.

### HasRelatedType

`func (o *BulkRelatedTypeExportOption) HasRelatedType() bool`

HasRelatedType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


