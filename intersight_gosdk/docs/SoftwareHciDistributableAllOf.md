# SoftwareHciDistributableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.HciDistributable"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.HciDistributable"]
**MetaInfo** | Pointer to **interface{}** | Additional information associated with the distributable object. This data is provided as a JSON blob by Nutanix, a partner vendor, and the format is not fixed. | [optional] [readonly] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwareHciDistributableAllOf

`func NewSoftwareHciDistributableAllOf(classId string, objectType string, ) *SoftwareHciDistributableAllOf`

NewSoftwareHciDistributableAllOf instantiates a new SoftwareHciDistributableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHciDistributableAllOfWithDefaults

`func NewSoftwareHciDistributableAllOfWithDefaults() *SoftwareHciDistributableAllOf`

NewSoftwareHciDistributableAllOfWithDefaults instantiates a new SoftwareHciDistributableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareHciDistributableAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareHciDistributableAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareHciDistributableAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareHciDistributableAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareHciDistributableAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareHciDistributableAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMetaInfo

`func (o *SoftwareHciDistributableAllOf) GetMetaInfo() interface{}`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *SoftwareHciDistributableAllOf) GetMetaInfoOk() (*interface{}, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *SoftwareHciDistributableAllOf) SetMetaInfo(v interface{})`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *SoftwareHciDistributableAllOf) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### SetMetaInfoNil

`func (o *SoftwareHciDistributableAllOf) SetMetaInfoNil(b bool)`

 SetMetaInfoNil sets the value for MetaInfo to be an explicit nil

### UnsetMetaInfo
`func (o *SoftwareHciDistributableAllOf) UnsetMetaInfo()`

UnsetMetaInfo ensures that no value is present for MetaInfo, not even an explicit nil
### GetCatalog

`func (o *SoftwareHciDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHciDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHciDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHciDistributableAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


