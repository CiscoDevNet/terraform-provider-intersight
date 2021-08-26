# SoftwareHclMetaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.HclMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.HclMeta"]
**ContentType** | Pointer to **string** | The type of content that the Json file holds (Incremental or full dump). * &#x60;Full&#x60; - Indicates that the JSON File does have full content for HCL metadata. * &#x60;Incremental&#x60; - Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded. | [optional] [default to "Full"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwareHclMetaAllOf

`func NewSoftwareHclMetaAllOf(classId string, objectType string, ) *SoftwareHclMetaAllOf`

NewSoftwareHclMetaAllOf instantiates a new SoftwareHclMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHclMetaAllOfWithDefaults

`func NewSoftwareHclMetaAllOfWithDefaults() *SoftwareHclMetaAllOf`

NewSoftwareHclMetaAllOfWithDefaults instantiates a new SoftwareHclMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareHclMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareHclMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareHclMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareHclMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareHclMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareHclMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContentType

`func (o *SoftwareHclMetaAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SoftwareHclMetaAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SoftwareHclMetaAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SoftwareHclMetaAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareHclMetaAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHclMetaAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHclMetaAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHclMetaAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


