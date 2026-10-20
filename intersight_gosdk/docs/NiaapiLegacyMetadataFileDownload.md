# NiaapiLegacyMetadataFileDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.LegacyMetadataFileDownload"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.LegacyMetadataFileDownload"]
**FileName** | Pointer to **string** | Name of the release-signed legacy metadata package file. | [optional] [readonly] 

## Methods

### NewNiaapiLegacyMetadataFileDownload

`func NewNiaapiLegacyMetadataFileDownload(classId string, objectType string, ) *NiaapiLegacyMetadataFileDownload`

NewNiaapiLegacyMetadataFileDownload instantiates a new NiaapiLegacyMetadataFileDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiLegacyMetadataFileDownloadWithDefaults

`func NewNiaapiLegacyMetadataFileDownloadWithDefaults() *NiaapiLegacyMetadataFileDownload`

NewNiaapiLegacyMetadataFileDownloadWithDefaults instantiates a new NiaapiLegacyMetadataFileDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiLegacyMetadataFileDownload) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiLegacyMetadataFileDownload) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiLegacyMetadataFileDownload) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiLegacyMetadataFileDownload) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiLegacyMetadataFileDownload) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiLegacyMetadataFileDownload) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *NiaapiLegacyMetadataFileDownload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *NiaapiLegacyMetadataFileDownload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *NiaapiLegacyMetadataFileDownload) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *NiaapiLegacyMetadataFileDownload) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


