# SoftwarerepositoryDownloadSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "softwarerepository.DownloadSpec"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "softwarerepository.DownloadSpec"]
**AuthToken** | Pointer to **string** | The OAuth2 token that will be used during image download by the endpoint to authenticate with file server. | [optional] 
**Certificate** | Pointer to **string** | The certificate of file server that will be used by the endpoint to validate the server before starting the file download. | [optional] 
**Filename** | Pointer to **string** | The name of the firmware image. | [optional] 
**Md5sum** | Pointer to **string** | MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download. | [optional] 
**Size** | Pointer to **int64** | The size (in bytes) of the firmware image. | [optional] 
**Url** | Pointer to **string** | The URL of this file in file server. The endpoint uses this URL to download the file from the file server. | [optional] 
**File** | Pointer to [**SoftwarerepositoryFileRelationship**](SoftwarerepositoryFileRelationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryDownloadSpecAllOf

`func NewSoftwarerepositoryDownloadSpecAllOf(classId string, objectType string, ) *SoftwarerepositoryDownloadSpecAllOf`

NewSoftwarerepositoryDownloadSpecAllOf instantiates a new SoftwarerepositoryDownloadSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryDownloadSpecAllOfWithDefaults

`func NewSoftwarerepositoryDownloadSpecAllOfWithDefaults() *SoftwarerepositoryDownloadSpecAllOf`

NewSoftwarerepositoryDownloadSpecAllOfWithDefaults instantiates a new SoftwarerepositoryDownloadSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthToken

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetCertificate

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetFilename

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetSize

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFile

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetFile() SoftwarerepositoryFileRelationship`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SoftwarerepositoryDownloadSpecAllOf) GetFileOk() (*SoftwarerepositoryFileRelationship, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SoftwarerepositoryDownloadSpecAllOf) SetFile(v SoftwarerepositoryFileRelationship)`

SetFile sets File field to given value.

### HasFile

`func (o *SoftwarerepositoryDownloadSpecAllOf) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


