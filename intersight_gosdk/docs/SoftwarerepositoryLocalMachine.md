# SoftwarerepositoryLocalMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | Pointer to **string** | When the import action in the file MO is updated with &#39;GeneratePreSignedDownloadUrl&#39;, Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL. | [optional] [readonly] 
**PartSize** | Pointer to **int64** | The chunk size (in bytes) for each part of the file to be uploaded. | [optional] 
**UploadId** | Pointer to **string** | When the import action in file MO is updated with &#39;GeneratePreSignedUploadUrl&#39;, Intersight shall return a upload Id in this property as part of the PATCH response. | [optional] 
**UploadUrl** | Pointer to **string** | When a file MO is created with &#39;LocalMachine&#39; as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object&#39;s transfer state to success. | [optional] [readonly] 
**UploadUrls** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSoftwarerepositoryLocalMachine

`func NewSoftwarerepositoryLocalMachine() *SoftwarerepositoryLocalMachine`

NewSoftwarerepositoryLocalMachine instantiates a new SoftwarerepositoryLocalMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryLocalMachineWithDefaults

`func NewSoftwarerepositoryLocalMachineWithDefaults() *SoftwarerepositoryLocalMachine`

NewSoftwarerepositoryLocalMachineWithDefaults instantiates a new SoftwarerepositoryLocalMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadUrl

`func (o *SoftwarerepositoryLocalMachine) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SoftwarerepositoryLocalMachine) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SoftwarerepositoryLocalMachine) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *SoftwarerepositoryLocalMachine) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetPartSize

`func (o *SoftwarerepositoryLocalMachine) GetPartSize() int64`

GetPartSize returns the PartSize field if non-nil, zero value otherwise.

### GetPartSizeOk

`func (o *SoftwarerepositoryLocalMachine) GetPartSizeOk() (*int64, bool)`

GetPartSizeOk returns a tuple with the PartSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartSize

`func (o *SoftwarerepositoryLocalMachine) SetPartSize(v int64)`

SetPartSize sets PartSize field to given value.

### HasPartSize

`func (o *SoftwarerepositoryLocalMachine) HasPartSize() bool`

HasPartSize returns a boolean if a field has been set.

### GetUploadId

`func (o *SoftwarerepositoryLocalMachine) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *SoftwarerepositoryLocalMachine) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *SoftwarerepositoryLocalMachine) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *SoftwarerepositoryLocalMachine) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *SoftwarerepositoryLocalMachine) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *SoftwarerepositoryLocalMachine) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *SoftwarerepositoryLocalMachine) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *SoftwarerepositoryLocalMachine) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadUrls

`func (o *SoftwarerepositoryLocalMachine) GetUploadUrls() []string`

GetUploadUrls returns the UploadUrls field if non-nil, zero value otherwise.

### GetUploadUrlsOk

`func (o *SoftwarerepositoryLocalMachine) GetUploadUrlsOk() (*[]string, bool)`

GetUploadUrlsOk returns a tuple with the UploadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrls

`func (o *SoftwarerepositoryLocalMachine) SetUploadUrls(v []string)`

SetUploadUrls sets UploadUrls field to given value.

### HasUploadUrls

`func (o *SoftwarerepositoryLocalMachine) HasUploadUrls() bool`

HasUploadUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


