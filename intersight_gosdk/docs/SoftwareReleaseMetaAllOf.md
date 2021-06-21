# SoftwareReleaseMetaAllOf

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

### NewSoftwareReleaseMetaAllOf

`func NewSoftwareReleaseMetaAllOf(classId string, objectType string, ) *SoftwareReleaseMetaAllOf`

NewSoftwareReleaseMetaAllOf instantiates a new SoftwareReleaseMetaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareReleaseMetaAllOfWithDefaults

`func NewSoftwareReleaseMetaAllOfWithDefaults() *SoftwareReleaseMetaAllOf`

NewSoftwareReleaseMetaAllOfWithDefaults instantiates a new SoftwareReleaseMetaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwareReleaseMetaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwareReleaseMetaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwareReleaseMetaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwareReleaseMetaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareReleaseMetaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareReleaseMetaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageType

`func (o *SoftwareReleaseMetaAllOf) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *SoftwareReleaseMetaAllOf) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *SoftwareReleaseMetaAllOf) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *SoftwareReleaseMetaAllOf) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetLatestFileName

`func (o *SoftwareReleaseMetaAllOf) GetLatestFileName() string`

GetLatestFileName returns the LatestFileName field if non-nil, zero value otherwise.

### GetLatestFileNameOk

`func (o *SoftwareReleaseMetaAllOf) GetLatestFileNameOk() (*string, bool)`

GetLatestFileNameOk returns a tuple with the LatestFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestFileName

`func (o *SoftwareReleaseMetaAllOf) SetLatestFileName(v string)`

SetLatestFileName sets LatestFileName field to given value.

### HasLatestFileName

`func (o *SoftwareReleaseMetaAllOf) HasLatestFileName() bool`

HasLatestFileName returns a boolean if a field has been set.

### GetLatestVersion

`func (o *SoftwareReleaseMetaAllOf) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *SoftwareReleaseMetaAllOf) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *SoftwareReleaseMetaAllOf) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *SoftwareReleaseMetaAllOf) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *SoftwareReleaseMetaAllOf) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *SoftwareReleaseMetaAllOf) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *SoftwareReleaseMetaAllOf) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *SoftwareReleaseMetaAllOf) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetCatalog

`func (o *SoftwareReleaseMetaAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SoftwareReleaseMetaAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SoftwareReleaseMetaAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SoftwareReleaseMetaAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetImage

`func (o *SoftwareReleaseMetaAllOf) GetImage() FirmwareBaseDistributableRelationship`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SoftwareReleaseMetaAllOf) GetImageOk() (*FirmwareBaseDistributableRelationship, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SoftwareReleaseMetaAllOf) SetImage(v FirmwareBaseDistributableRelationship)`

SetImage sets Image field to given value.

### HasImage

`func (o *SoftwareReleaseMetaAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


