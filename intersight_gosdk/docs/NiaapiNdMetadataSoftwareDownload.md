# NiaapiNdMetadataSoftwareDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.NdMetadataSoftwareDownload"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.NdMetadataSoftwareDownload"]
**Description** | Pointer to **string** | Description of metadata package. | [optional] [readonly] 
**Name** | Pointer to **string** | Metadata package filename. | [optional] [readonly] 
**ReleaseDate** | Pointer to **string** | Release date of metadata packages. | [optional] [readonly] 
**Size** | Pointer to **int64** | Size of Metadata package. | [optional] [readonly] 
**Type** | Pointer to **string** | Type label for the package. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of Metadata package. | [optional] [readonly] 

## Methods

### NewNiaapiNdMetadataSoftwareDownload

`func NewNiaapiNdMetadataSoftwareDownload(classId string, objectType string, ) *NiaapiNdMetadataSoftwareDownload`

NewNiaapiNdMetadataSoftwareDownload instantiates a new NiaapiNdMetadataSoftwareDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNdMetadataSoftwareDownloadWithDefaults

`func NewNiaapiNdMetadataSoftwareDownloadWithDefaults() *NiaapiNdMetadataSoftwareDownload`

NewNiaapiNdMetadataSoftwareDownloadWithDefaults instantiates a new NiaapiNdMetadataSoftwareDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNdMetadataSoftwareDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNdMetadataSoftwareDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNdMetadataSoftwareDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNdMetadataSoftwareDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *NiaapiNdMetadataSoftwareDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiNdMetadataSoftwareDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiNdMetadataSoftwareDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *NiaapiNdMetadataSoftwareDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiaapiNdMetadataSoftwareDownload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiaapiNdMetadataSoftwareDownload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *NiaapiNdMetadataSoftwareDownload) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *NiaapiNdMetadataSoftwareDownload) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *NiaapiNdMetadataSoftwareDownload) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSize

`func (o *NiaapiNdMetadataSoftwareDownload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NiaapiNdMetadataSoftwareDownload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *NiaapiNdMetadataSoftwareDownload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *NiaapiNdMetadataSoftwareDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NiaapiNdMetadataSoftwareDownload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NiaapiNdMetadataSoftwareDownload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *NiaapiNdMetadataSoftwareDownload) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiNdMetadataSoftwareDownload) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiNdMetadataSoftwareDownload) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiNdMetadataSoftwareDownload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


