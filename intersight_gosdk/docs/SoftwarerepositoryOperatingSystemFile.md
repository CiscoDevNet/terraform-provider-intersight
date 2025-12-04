# SoftwarerepositoryOperatingSystemFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.OperatingSystemFile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.OperatingSystemFile"]
**ImportProgress** | Pointer to **int64** | The progress percentage for the import operation. | [optional] 
**SampleHashes** | Pointer to **string** | File sample hashes at deterministic positions for efficient duplicate detection of large files. | [optional] 
**Vendor** | Pointer to **string** | The vendor or publisher of this file. | [optional] 
**Catalog** | Pointer to [**NullableSoftwarerepositoryCatalogRelationship**](SoftwarerepositoryCatalogRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryOperatingSystemFile

`func NewSoftwarerepositoryOperatingSystemFile(classId string, objectType string, ) *SoftwarerepositoryOperatingSystemFile`

NewSoftwarerepositoryOperatingSystemFile instantiates a new SoftwarerepositoryOperatingSystemFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryOperatingSystemFileWithDefaults

`func NewSoftwarerepositoryOperatingSystemFileWithDefaults() *SoftwarerepositoryOperatingSystemFile`

NewSoftwarerepositoryOperatingSystemFileWithDefaults instantiates a new SoftwarerepositoryOperatingSystemFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryOperatingSystemFile) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryOperatingSystemFile) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryOperatingSystemFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryOperatingSystemFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImportProgress

`func (o *SoftwarerepositoryOperatingSystemFile) GetImportProgress() int64`

GetImportProgress returns the ImportProgress field if non-nil, zero value otherwise.

### GetImportProgressOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetImportProgressOk() (*int64, bool)`

GetImportProgressOk returns a tuple with the ImportProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportProgress

`func (o *SoftwarerepositoryOperatingSystemFile) SetImportProgress(v int64)`

SetImportProgress sets ImportProgress field to given value.

### HasImportProgress

`func (o *SoftwarerepositoryOperatingSystemFile) HasImportProgress() bool`

HasImportProgress returns a boolean if a field has been set.

### GetSampleHashes

`func (o *SoftwarerepositoryOperatingSystemFile) GetSampleHashes() string`

GetSampleHashes returns the SampleHashes field if non-nil, zero value otherwise.

### GetSampleHashesOk

`func (o *SoftwarerepositoryOperatingSystemFile) GetSampleHashesOk() (*string, bool)`

GetSampleHashesOk returns a tuple with the SampleHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleHashes

`func (o *SoftwarerepositoryOperatingSystemFile) SetSampleHashes(v string)`

SetSampleHashes sets SampleHashes field to given value.

### HasSampleHashes

`func (o *SoftwarerepositoryOperatingSystemFile) HasSampleHashes() bool`

HasSampleHashes returns a boolean if a field has been set.

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

### SetCatalogNil

`func (o *SoftwarerepositoryOperatingSystemFile) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *SoftwarerepositoryOperatingSystemFile) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


