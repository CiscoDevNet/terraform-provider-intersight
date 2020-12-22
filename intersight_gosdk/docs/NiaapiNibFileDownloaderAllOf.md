# NiaapiNibFileDownloaderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.NibFileDownloader"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.NibFileDownloader"]
**FileName** | Pointer to **string** | Filename of this Metadata package file, folder will be handled by api. | [optional] 
**PresignedUrl** | Pointer to **string** | The presigned URL from server to download this file. | [optional] 

## Methods

### NewNiaapiNibFileDownloaderAllOf

`func NewNiaapiNibFileDownloaderAllOf(classId string, objectType string, ) *NiaapiNibFileDownloaderAllOf`

NewNiaapiNibFileDownloaderAllOf instantiates a new NiaapiNibFileDownloaderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNibFileDownloaderAllOfWithDefaults

`func NewNiaapiNibFileDownloaderAllOfWithDefaults() *NiaapiNibFileDownloaderAllOf`

NewNiaapiNibFileDownloaderAllOfWithDefaults instantiates a new NiaapiNibFileDownloaderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNibFileDownloaderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNibFileDownloaderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNibFileDownloaderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNibFileDownloaderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNibFileDownloaderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNibFileDownloaderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *NiaapiNibFileDownloaderAllOf) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *NiaapiNibFileDownloaderAllOf) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *NiaapiNibFileDownloaderAllOf) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *NiaapiNibFileDownloaderAllOf) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetPresignedUrl

`func (o *NiaapiNibFileDownloaderAllOf) GetPresignedUrl() string`

GetPresignedUrl returns the PresignedUrl field if non-nil, zero value otherwise.

### GetPresignedUrlOk

`func (o *NiaapiNibFileDownloaderAllOf) GetPresignedUrlOk() (*string, bool)`

GetPresignedUrlOk returns a tuple with the PresignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedUrl

`func (o *NiaapiNibFileDownloaderAllOf) SetPresignedUrl(v string)`

SetPresignedUrl sets PresignedUrl field to given value.

### HasPresignedUrl

`func (o *NiaapiNibFileDownloaderAllOf) HasPresignedUrl() bool`

HasPresignedUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


