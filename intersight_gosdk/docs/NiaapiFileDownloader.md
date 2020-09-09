# NiaapiFileDownloader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Filename of this Metadata package file, folder will be handled by api. | [optional] 
**PresignedUrl** | Pointer to **string** | The presigned URL from server to download this file. | [optional] 

## Methods

### NewNiaapiFileDownloader

`func NewNiaapiFileDownloader() *NiaapiFileDownloader`

NewNiaapiFileDownloader instantiates a new NiaapiFileDownloader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiFileDownloaderWithDefaults

`func NewNiaapiFileDownloaderWithDefaults() *NiaapiFileDownloader`

NewNiaapiFileDownloaderWithDefaults instantiates a new NiaapiFileDownloader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *NiaapiFileDownloader) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *NiaapiFileDownloader) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *NiaapiFileDownloader) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *NiaapiFileDownloader) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetPresignedUrl

`func (o *NiaapiFileDownloader) GetPresignedUrl() string`

GetPresignedUrl returns the PresignedUrl field if non-nil, zero value otherwise.

### GetPresignedUrlOk

`func (o *NiaapiFileDownloader) GetPresignedUrlOk() (*string, bool)`

GetPresignedUrlOk returns a tuple with the PresignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedUrl

`func (o *NiaapiFileDownloader) SetPresignedUrl(v string)`

SetPresignedUrl sets PresignedUrl field to given value.

### HasPresignedUrl

`func (o *NiaapiFileDownloader) HasPresignedUrl() bool`

HasPresignedUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


