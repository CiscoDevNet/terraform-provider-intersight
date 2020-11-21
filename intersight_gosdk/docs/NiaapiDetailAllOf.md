# NiaapiDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.Detail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.Detail"]
**Chksum** | Pointer to **string** | Checksum of this part of Content. | [optional] 
**Filename** | Pointer to **string** | The file name within this Metadata file. | [optional] 
**Name** | Pointer to **string** | The name of this Content. | [optional] 

## Methods

### NewNiaapiDetailAllOf

`func NewNiaapiDetailAllOf(classId string, objectType string, ) *NiaapiDetailAllOf`

NewNiaapiDetailAllOf instantiates a new NiaapiDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiDetailAllOfWithDefaults

`func NewNiaapiDetailAllOfWithDefaults() *NiaapiDetailAllOf`

NewNiaapiDetailAllOfWithDefaults instantiates a new NiaapiDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChksum

`func (o *NiaapiDetailAllOf) GetChksum() string`

GetChksum returns the Chksum field if non-nil, zero value otherwise.

### GetChksumOk

`func (o *NiaapiDetailAllOf) GetChksumOk() (*string, bool)`

GetChksumOk returns a tuple with the Chksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChksum

`func (o *NiaapiDetailAllOf) SetChksum(v string)`

SetChksum sets Chksum field to given value.

### HasChksum

`func (o *NiaapiDetailAllOf) HasChksum() bool`

HasChksum returns a boolean if a field has been set.

### GetFilename

`func (o *NiaapiDetailAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *NiaapiDetailAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *NiaapiDetailAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *NiaapiDetailAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *NiaapiDetailAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiaapiDetailAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiaapiDetailAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiaapiDetailAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


