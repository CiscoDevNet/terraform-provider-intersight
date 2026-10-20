# NiaapiMaintainedRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**MaintainedRelease** | Pointer to **string** | Lastest maintained release. | [optional] 
**SoftwareRelease** | Pointer to **string** | Software release version string. | [optional] 
**VersionTag** | Pointer to **string** | Long lived version or short lived version. | [optional] 

## Methods

### NewNiaapiMaintainedRelease

`func NewNiaapiMaintainedRelease(classId string, objectType string, ) *NiaapiMaintainedRelease`

NewNiaapiMaintainedRelease instantiates a new NiaapiMaintainedRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiMaintainedReleaseWithDefaults

`func NewNiaapiMaintainedReleaseWithDefaults() *NiaapiMaintainedRelease`

NewNiaapiMaintainedReleaseWithDefaults instantiates a new NiaapiMaintainedRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiMaintainedRelease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiMaintainedRelease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiMaintainedRelease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiMaintainedRelease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiMaintainedRelease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiMaintainedRelease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMaintainedRelease

`func (o *NiaapiMaintainedRelease) GetMaintainedRelease() string`

GetMaintainedRelease returns the MaintainedRelease field if non-nil, zero value otherwise.

### GetMaintainedReleaseOk

`func (o *NiaapiMaintainedRelease) GetMaintainedReleaseOk() (*string, bool)`

GetMaintainedReleaseOk returns a tuple with the MaintainedRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainedRelease

`func (o *NiaapiMaintainedRelease) SetMaintainedRelease(v string)`

SetMaintainedRelease sets MaintainedRelease field to given value.

### HasMaintainedRelease

`func (o *NiaapiMaintainedRelease) HasMaintainedRelease() bool`

HasMaintainedRelease returns a boolean if a field has been set.

### GetSoftwareRelease

`func (o *NiaapiMaintainedRelease) GetSoftwareRelease() string`

GetSoftwareRelease returns the SoftwareRelease field if non-nil, zero value otherwise.

### GetSoftwareReleaseOk

`func (o *NiaapiMaintainedRelease) GetSoftwareReleaseOk() (*string, bool)`

GetSoftwareReleaseOk returns a tuple with the SoftwareRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareRelease

`func (o *NiaapiMaintainedRelease) SetSoftwareRelease(v string)`

SetSoftwareRelease sets SoftwareRelease field to given value.

### HasSoftwareRelease

`func (o *NiaapiMaintainedRelease) HasSoftwareRelease() bool`

HasSoftwareRelease returns a boolean if a field has been set.

### GetVersionTag

`func (o *NiaapiMaintainedRelease) GetVersionTag() string`

GetVersionTag returns the VersionTag field if non-nil, zero value otherwise.

### GetVersionTagOk

`func (o *NiaapiMaintainedRelease) GetVersionTagOk() (*string, bool)`

GetVersionTagOk returns a tuple with the VersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionTag

`func (o *NiaapiMaintainedRelease) SetVersionTag(v string)`

SetVersionTag sets VersionTag field to given value.

### HasVersionTag

`func (o *NiaapiMaintainedRelease) HasVersionTag() bool`

HasVersionTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


