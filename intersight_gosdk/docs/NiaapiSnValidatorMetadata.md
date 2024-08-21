# NiaapiSnValidatorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.SnValidatorMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.SnValidatorMetadata"]
**Checksum** | Pointer to **string** | Checksum of SnValidatorData file. | [optional] 
**FileName** | Pointer to **string** | The filename of sn metadata file. | [optional] 
**Version** | Pointer to **int64** | The version number of the SerialNumber Metadata. | [optional] 

## Methods

### NewNiaapiSnValidatorMetadata

`func NewNiaapiSnValidatorMetadata(classId string, objectType string, ) *NiaapiSnValidatorMetadata`

NewNiaapiSnValidatorMetadata instantiates a new NiaapiSnValidatorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiSnValidatorMetadataWithDefaults

`func NewNiaapiSnValidatorMetadataWithDefaults() *NiaapiSnValidatorMetadata`

NewNiaapiSnValidatorMetadataWithDefaults instantiates a new NiaapiSnValidatorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiSnValidatorMetadata) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiSnValidatorMetadata) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiSnValidatorMetadata) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiSnValidatorMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiSnValidatorMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiSnValidatorMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChecksum

`func (o *NiaapiSnValidatorMetadata) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *NiaapiSnValidatorMetadata) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *NiaapiSnValidatorMetadata) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *NiaapiSnValidatorMetadata) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetFileName

`func (o *NiaapiSnValidatorMetadata) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *NiaapiSnValidatorMetadata) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *NiaapiSnValidatorMetadata) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *NiaapiSnValidatorMetadata) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetVersion

`func (o *NiaapiSnValidatorMetadata) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiSnValidatorMetadata) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiSnValidatorMetadata) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiSnValidatorMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


