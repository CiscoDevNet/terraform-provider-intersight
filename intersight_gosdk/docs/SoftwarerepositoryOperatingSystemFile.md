# SoftwarerepositoryOperatingSystemFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryOperatingSystemFile

`func NewSoftwarerepositoryOperatingSystemFile() *SoftwarerepositoryOperatingSystemFile`

NewSoftwarerepositoryOperatingSystemFile instantiates a new SoftwarerepositoryOperatingSystemFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryOperatingSystemFileWithDefaults

`func NewSoftwarerepositoryOperatingSystemFileWithDefaults() *SoftwarerepositoryOperatingSystemFile`

NewSoftwarerepositoryOperatingSystemFileWithDefaults instantiates a new SoftwarerepositoryOperatingSystemFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SoftwarerepositoryOperatingSystemFile) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SoftwarerepositoryOperatingSystemFile) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SoftwarerepositoryOperatingSystemFile) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwarerepositoryOperatingSystemFile) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwarerepositoryOperatingSystemFile) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwarerepositoryOperatingSystemFile) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


