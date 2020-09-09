# SoftwareHclMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The type of content that the Json file holds (Incremental or full dump). * &#x60;Full&#x60; - Indicates that the JSON File does have full content for HCL metadata. * &#x60;Incremental&#x60; - Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded. | [optional] [default to "Full"]
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwareHclMeta

`func NewSoftwareHclMeta() *SoftwareHclMeta`

NewSoftwareHclMeta instantiates a new SoftwareHclMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareHclMetaWithDefaults

`func NewSoftwareHclMetaWithDefaults() *SoftwareHclMeta`

NewSoftwareHclMetaWithDefaults instantiates a new SoftwareHclMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *SoftwareHclMeta) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SoftwareHclMeta) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SoftwareHclMeta) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SoftwareHclMeta) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareHclMeta) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareHclMeta) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareHclMeta) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareHclMeta) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


