# NiaapiNibMetadataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.NibMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.NibMetadata"]
**Content** | Pointer to [**[]NiaapiDetail**](NiaapiDetail.md) |  | [optional] 
**Date** | Pointer to **time.Time** | The date when the package was generated. | [optional] 
**MetadataChksum** | Pointer to **string** | Chksum used to check the integrity of the metadata file downloaded. | [optional] 
**MetadataFilename** | Pointer to **string** | The filename of the metadata package. | [optional] 
**Version** | Pointer to **int64** | The version number of the metadata package. | [optional] 

## Methods

### NewNiaapiNibMetadataAllOf

`func NewNiaapiNibMetadataAllOf(classId string, objectType string, ) *NiaapiNibMetadataAllOf`

NewNiaapiNibMetadataAllOf instantiates a new NiaapiNibMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNibMetadataAllOfWithDefaults

`func NewNiaapiNibMetadataAllOfWithDefaults() *NiaapiNibMetadataAllOf`

NewNiaapiNibMetadataAllOfWithDefaults instantiates a new NiaapiNibMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNibMetadataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNibMetadataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNibMetadataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNibMetadataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNibMetadataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNibMetadataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContent

`func (o *NiaapiNibMetadataAllOf) GetContent() []NiaapiDetail`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NiaapiNibMetadataAllOf) GetContentOk() (*[]NiaapiDetail, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NiaapiNibMetadataAllOf) SetContent(v []NiaapiDetail)`

SetContent sets Content field to given value.

### HasContent

`func (o *NiaapiNibMetadataAllOf) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *NiaapiNibMetadataAllOf) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *NiaapiNibMetadataAllOf) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDate

`func (o *NiaapiNibMetadataAllOf) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NiaapiNibMetadataAllOf) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NiaapiNibMetadataAllOf) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *NiaapiNibMetadataAllOf) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMetadataChksum

`func (o *NiaapiNibMetadataAllOf) GetMetadataChksum() string`

GetMetadataChksum returns the MetadataChksum field if non-nil, zero value otherwise.

### GetMetadataChksumOk

`func (o *NiaapiNibMetadataAllOf) GetMetadataChksumOk() (*string, bool)`

GetMetadataChksumOk returns a tuple with the MetadataChksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChksum

`func (o *NiaapiNibMetadataAllOf) SetMetadataChksum(v string)`

SetMetadataChksum sets MetadataChksum field to given value.

### HasMetadataChksum

`func (o *NiaapiNibMetadataAllOf) HasMetadataChksum() bool`

HasMetadataChksum returns a boolean if a field has been set.

### GetMetadataFilename

`func (o *NiaapiNibMetadataAllOf) GetMetadataFilename() string`

GetMetadataFilename returns the MetadataFilename field if non-nil, zero value otherwise.

### GetMetadataFilenameOk

`func (o *NiaapiNibMetadataAllOf) GetMetadataFilenameOk() (*string, bool)`

GetMetadataFilenameOk returns a tuple with the MetadataFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilename

`func (o *NiaapiNibMetadataAllOf) SetMetadataFilename(v string)`

SetMetadataFilename sets MetadataFilename field to given value.

### HasMetadataFilename

`func (o *NiaapiNibMetadataAllOf) HasMetadataFilename() bool`

HasMetadataFilename returns a boolean if a field has been set.

### GetVersion

`func (o *NiaapiNibMetadataAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiNibMetadataAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiNibMetadataAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiNibMetadataAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


