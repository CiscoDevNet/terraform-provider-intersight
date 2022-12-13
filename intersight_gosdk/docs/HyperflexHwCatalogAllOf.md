# HyperflexHwCatalogAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HwCatalog"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HwCatalog"]
**Version** | Pointer to **string** | Hardware catalog version for HyperFlex hardware catalog. | [optional] 
**CatalogInfos** | Pointer to [**[]HclHwCatalogInfoRelationship**](HclHwCatalogInfoRelationship.md) | An array of relationships to hclHwCatalogInfo resources. | [optional] 

## Methods

### NewHyperflexHwCatalogAllOf

`func NewHyperflexHwCatalogAllOf(classId string, objectType string, ) *HyperflexHwCatalogAllOf`

NewHyperflexHwCatalogAllOf instantiates a new HyperflexHwCatalogAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHwCatalogAllOfWithDefaults

`func NewHyperflexHwCatalogAllOfWithDefaults() *HyperflexHwCatalogAllOf`

NewHyperflexHwCatalogAllOfWithDefaults instantiates a new HyperflexHwCatalogAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHwCatalogAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHwCatalogAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHwCatalogAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHwCatalogAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHwCatalogAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHwCatalogAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVersion

`func (o *HyperflexHwCatalogAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexHwCatalogAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexHwCatalogAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexHwCatalogAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalogInfos

`func (o *HyperflexHwCatalogAllOf) GetCatalogInfos() []HclHwCatalogInfoRelationship`

GetCatalogInfos returns the CatalogInfos field if non-nil, zero value otherwise.

### GetCatalogInfosOk

`func (o *HyperflexHwCatalogAllOf) GetCatalogInfosOk() (*[]HclHwCatalogInfoRelationship, bool)`

GetCatalogInfosOk returns a tuple with the CatalogInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogInfos

`func (o *HyperflexHwCatalogAllOf) SetCatalogInfos(v []HclHwCatalogInfoRelationship)`

SetCatalogInfos sets CatalogInfos field to given value.

### HasCatalogInfos

`func (o *HyperflexHwCatalogAllOf) HasCatalogInfos() bool`

HasCatalogInfos returns a boolean if a field has been set.

### SetCatalogInfosNil

`func (o *HyperflexHwCatalogAllOf) SetCatalogInfosNil(b bool)`

 SetCatalogInfosNil sets the value for CatalogInfos to be an explicit nil

### UnsetCatalogInfos
`func (o *HyperflexHwCatalogAllOf) UnsetCatalogInfos()`

UnsetCatalogInfos ensures that no value is present for CatalogInfos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


