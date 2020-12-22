# NiaapiNiaMetadataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.NiaMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.NiaMetadata"]
**Content** | Pointer to [**[]NiaapiDetail**](NiaapiDetail.md) |  | [optional] 
**Date** | Pointer to **time.Time** | The date when this package is generated. | [optional] 
**MetadataChksum** | Pointer to **string** | Chksum used to check the integrity of the Metadata file downloaded. | [optional] 
**MetadataFilename** | Pointer to **string** | The Filename of this Metadata package. | [optional] 
**Version** | Pointer to **int64** | The version number of the Metadata package. | [optional] 

## Methods

### NewNiaapiNiaMetadataAllOf

`func NewNiaapiNiaMetadataAllOf(classId string, objectType string, ) *NiaapiNiaMetadataAllOf`

NewNiaapiNiaMetadataAllOf instantiates a new NiaapiNiaMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiNiaMetadataAllOfWithDefaults

`func NewNiaapiNiaMetadataAllOfWithDefaults() *NiaapiNiaMetadataAllOf`

NewNiaapiNiaMetadataAllOfWithDefaults instantiates a new NiaapiNiaMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiNiaMetadataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiNiaMetadataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiNiaMetadataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiNiaMetadataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiNiaMetadataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiNiaMetadataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContent

`func (o *NiaapiNiaMetadataAllOf) GetContent() []NiaapiDetail`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NiaapiNiaMetadataAllOf) GetContentOk() (*[]NiaapiDetail, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NiaapiNiaMetadataAllOf) SetContent(v []NiaapiDetail)`

SetContent sets Content field to given value.

### HasContent

`func (o *NiaapiNiaMetadataAllOf) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *NiaapiNiaMetadataAllOf) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *NiaapiNiaMetadataAllOf) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDate

`func (o *NiaapiNiaMetadataAllOf) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NiaapiNiaMetadataAllOf) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NiaapiNiaMetadataAllOf) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *NiaapiNiaMetadataAllOf) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMetadataChksum

`func (o *NiaapiNiaMetadataAllOf) GetMetadataChksum() string`

GetMetadataChksum returns the MetadataChksum field if non-nil, zero value otherwise.

### GetMetadataChksumOk

`func (o *NiaapiNiaMetadataAllOf) GetMetadataChksumOk() (*string, bool)`

GetMetadataChksumOk returns a tuple with the MetadataChksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataChksum

`func (o *NiaapiNiaMetadataAllOf) SetMetadataChksum(v string)`

SetMetadataChksum sets MetadataChksum field to given value.

### HasMetadataChksum

`func (o *NiaapiNiaMetadataAllOf) HasMetadataChksum() bool`

HasMetadataChksum returns a boolean if a field has been set.

### GetMetadataFilename

`func (o *NiaapiNiaMetadataAllOf) GetMetadataFilename() string`

GetMetadataFilename returns the MetadataFilename field if non-nil, zero value otherwise.

### GetMetadataFilenameOk

`func (o *NiaapiNiaMetadataAllOf) GetMetadataFilenameOk() (*string, bool)`

GetMetadataFilenameOk returns a tuple with the MetadataFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilename

`func (o *NiaapiNiaMetadataAllOf) SetMetadataFilename(v string)`

SetMetadataFilename sets MetadataFilename field to given value.

### HasMetadataFilename

`func (o *NiaapiNiaMetadataAllOf) HasMetadataFilename() bool`

HasMetadataFilename returns a boolean if a field has been set.

### GetVersion

`func (o *NiaapiNiaMetadataAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiNiaMetadataAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiNiaMetadataAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiNiaMetadataAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


