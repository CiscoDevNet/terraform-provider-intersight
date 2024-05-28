# SoftwareHciDistributable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.HciDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.HciDistributable"]
**MetaInfo** | Pointer to **interface{}** | Additional information associated with the distributable object. This data is provided as a JSON blob by Nutanix, a partner vendor, and the format is not fixed. | [optional] [readonly] 
**Catalog** | Pointer to [**NullableSoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwareHciDistributable

`func NewSoftwareHciDistributable(classId string, objectType string, ) *SoftwareHciDistributable`

NewSoftwareHciDistributable instantiates a new SoftwareHciDistributable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHciDistributableWithDefaults

`func NewSoftwareHciDistributableWithDefaults() *SoftwareHciDistributable`

NewSoftwareHciDistributableWithDefaults instantiates a new SoftwareHciDistributable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareHciDistributable) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareHciDistributable) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareHciDistributable) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareHciDistributable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareHciDistributable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareHciDistributable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMetaInfo

`func (o *SoftwareHciDistributable) GetMetaInfo() interface{}`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *SoftwareHciDistributable) GetMetaInfoOk() (*interface{}, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *SoftwareHciDistributable) SetMetaInfo(v interface{})`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *SoftwareHciDistributable) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### SetMetaInfoNil

`func (o *SoftwareHciDistributable) SetMetaInfoNil(b bool)`

 SetMetaInfoNil sets the value for MetaInfo to be an explicit nil

### UnsetMetaInfo
`func (o *SoftwareHciDistributable) UnsetMetaInfo()`

UnsetMetaInfo ensures that no value is present for MetaInfo, not even an explicit nil
### GetCatalog

`func (o *SoftwareHciDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHciDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHciDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHciDistributable) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *SoftwareHciDistributable) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *SoftwareHciDistributable) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


