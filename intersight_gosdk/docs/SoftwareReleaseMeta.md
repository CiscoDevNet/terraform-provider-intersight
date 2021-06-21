# SoftwareReleaseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "software.ReleaseMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "software.ReleaseMeta"]
**ImageType** | Pointer to **string** | The subtype of the distributable image. For e.g. the firmware distributable is categorized according to the component it can upgrade - Standalone server, Intersight managed server or UCS Managed Fabric Interconnect. | [optional] 
**LatestFileName** | Pointer to **string** | The name of the latest image file uploaded for this software type. It is populated as part of the image import operation. | [optional] 
**LatestVersion** | Pointer to **string** | Latest version of the image avaiable for a specific software. | [optional] [readonly] 
**SoftwareTypeId** | Pointer to **string** | The software type id of the image (For e.g. firmware.Distributable, software.ApplianceDistributable, software.HyperflexBundleDistributable, software.UcsdBundleDistributable). | [optional] 
**Catalog** | Pointer to [**SoftwarerepositoryCatalogRelationship**](softwarerepository.Catalog.Relationship.md) |  | [optional] 
**Image** | Pointer to [**FirmwareBaseDistributableRelationship**](firmware.BaseDistributable.Relationship.md) |  | [optional] 

## Methods

### NewSoftwareReleaseMeta

`func NewSoftwareReleaseMeta(classId string, objectType string, ) *SoftwareReleaseMeta`

NewSoftwareReleaseMeta instantiates a new SoftwareReleaseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareReleaseMetaWithDefaults

`func NewSoftwareReleaseMetaWithDefaults() *SoftwareReleaseMeta`

NewSoftwareReleaseMetaWithDefaults instantiates a new SoftwareReleaseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareReleaseMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareReleaseMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareReleaseMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareReleaseMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareReleaseMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareReleaseMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageType

`func (o *SoftwareReleaseMeta) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *SoftwareReleaseMeta) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *SoftwareReleaseMeta) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *SoftwareReleaseMeta) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetLatestFileName

`func (o *SoftwareReleaseMeta) GetLatestFileName() string`

GetLatestFileName returns the LatestFileName field if non-nil, zero value otherwise.

### GetLatestFileNameOk

`func (o *SoftwareReleaseMeta) GetLatestFileNameOk() (*string, bool)`

GetLatestFileNameOk returns a tuple with the LatestFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFileName

`func (o *SoftwareReleaseMeta) SetLatestFileName(v string)`

SetLatestFileName sets LatestFileName field to given value.

### HasLatestFileName

`func (o *SoftwareReleaseMeta) HasLatestFileName() bool`

HasLatestFileName returns a boolean if a field has been set.

### GetLatestVersion

`func (o *SoftwareReleaseMeta) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *SoftwareReleaseMeta) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *SoftwareReleaseMeta) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *SoftwareReleaseMeta) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *SoftwareReleaseMeta) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *SoftwareReleaseMeta) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *SoftwareReleaseMeta) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *SoftwareReleaseMeta) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareReleaseMeta) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareReleaseMeta) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareReleaseMeta) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareReleaseMeta) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetImage

`func (o *SoftwareReleaseMeta) GetImage() FirmwareBaseDistributableRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SoftwareReleaseMeta) GetImageOk() (*FirmwareBaseDistributableRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SoftwareReleaseMeta) SetImage(v FirmwareBaseDistributableRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *SoftwareReleaseMeta) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


