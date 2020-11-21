# FirmwareDistributableMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.DistributableMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.DistributableMeta"]
**BucketName** | Pointer to **string** | The S3 bucket name where the images are present, if source is external cloud store. | [optional] 
**FileType** | Pointer to **string** | The type of distributable image, example huu, scu, driver, os. * &#x60;Distributable&#x60; - Stores firmware host utility images and fabric images. * &#x60;DriverDistributable&#x60; - Stores driver distributable images. * &#x60;ServerConfigurationUtilityDistributable&#x60; - Stores server configuration utility images. * &#x60;OperatingSystemFile&#x60; - Stores operating system iso images. * &#x60;HyperflexDistributable&#x60; - It stores HyperFlex images. | [optional] [default to "Distributable"]
**FromVersion** | Pointer to **string** | The version from which user can download images from amazon store, if source is external cloud store. | [optional] 
**Mdfid** | Pointer to **string** | The mdfid of the image provided by cisco.com. | [optional] 
**SoftwareTypeId** | Pointer to **string** | The software type id provided by cisco.com. | [optional] 
**Source** | Pointer to **string** | The image can be downloaded from cisco.com or external cloud store. * &#x60;Cisco&#x60; - External repository hosted on cisco.com. * &#x60;IntersightCloud&#x60; - Repository hosted by the Intersight Cloud. * &#x60;LocalMachine&#x60; - The file is available on the local client machine. Used as an upload source type. * &#x60;NetworkShare&#x60; - External repository in the customer datacenter. This will typically be a file server. | [optional] [default to "Cisco"]
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**ToVersion** | Pointer to **string** | The version till which user can download images from amazon store, if source is external cloud store. | [optional] 

## Methods

### NewFirmwareDistributableMeta

`func NewFirmwareDistributableMeta(classId string, objectType string, ) *FirmwareDistributableMeta`

NewFirmwareDistributableMeta instantiates a new FirmwareDistributableMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareDistributableMetaWithDefaults

`func NewFirmwareDistributableMetaWithDefaults() *FirmwareDistributableMeta`

NewFirmwareDistributableMetaWithDefaults instantiates a new FirmwareDistributableMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareDistributableMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareDistributableMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareDistributableMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareDistributableMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareDistributableMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareDistributableMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBucketName

`func (o *FirmwareDistributableMeta) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *FirmwareDistributableMeta) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *FirmwareDistributableMeta) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *FirmwareDistributableMeta) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFileType

`func (o *FirmwareDistributableMeta) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FirmwareDistributableMeta) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FirmwareDistributableMeta) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FirmwareDistributableMeta) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFromVersion

`func (o *FirmwareDistributableMeta) GetFromVersion() string`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *FirmwareDistributableMeta) GetFromVersionOk() (*string, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *FirmwareDistributableMeta) SetFromVersion(v string)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *FirmwareDistributableMeta) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.

### GetMdfid

`func (o *FirmwareDistributableMeta) GetMdfid() string`

GetMdfid returns the Mdfid field if non-nil, zero value otherwise.

### GetMdfidOk

`func (o *FirmwareDistributableMeta) GetMdfidOk() (*string, bool)`

GetMdfidOk returns a tuple with the Mdfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfid

`func (o *FirmwareDistributableMeta) SetMdfid(v string)`

SetMdfid sets Mdfid field to given value.

### HasMdfid

`func (o *FirmwareDistributableMeta) HasMdfid() bool`

HasMdfid returns a boolean if a field has been set.

### GetSoftwareTypeId

`func (o *FirmwareDistributableMeta) GetSoftwareTypeId() string`

GetSoftwareTypeId returns the SoftwareTypeId field if non-nil, zero value otherwise.

### GetSoftwareTypeIdOk

`func (o *FirmwareDistributableMeta) GetSoftwareTypeIdOk() (*string, bool)`

GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTypeId

`func (o *FirmwareDistributableMeta) SetSoftwareTypeId(v string)`

SetSoftwareTypeId sets SoftwareTypeId field to given value.

### HasSoftwareTypeId

`func (o *FirmwareDistributableMeta) HasSoftwareTypeId() bool`

HasSoftwareTypeId returns a boolean if a field has been set.

### GetSource

`func (o *FirmwareDistributableMeta) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FirmwareDistributableMeta) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FirmwareDistributableMeta) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FirmwareDistributableMeta) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSupportedModels

`func (o *FirmwareDistributableMeta) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *FirmwareDistributableMeta) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *FirmwareDistributableMeta) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *FirmwareDistributableMeta) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *FirmwareDistributableMeta) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *FirmwareDistributableMeta) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil
### GetToVersion

`func (o *FirmwareDistributableMeta) GetToVersion() string`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *FirmwareDistributableMeta) GetToVersionOk() (*string, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *FirmwareDistributableMeta) SetToVersion(v string)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *FirmwareDistributableMeta) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


