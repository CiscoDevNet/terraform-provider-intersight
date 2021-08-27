# NiaapiVersionRegexPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.VersionRegexPlatform"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.VersionRegexPlatform"]
**Anyllregex** | Pointer to **string** | All long live version Regex pattern. | [optional] 
**Currentlltrain** | Pointer to [**NullableNiaapiSoftwareRegex**](NiaapiSoftwareRegex.md) |  | [optional] 
**Latestsltrain** | Pointer to [**NullableNiaapiSoftwareRegex**](NiaapiSoftwareRegex.md) |  | [optional] 
**Sltrain** | Pointer to [**[]NiaapiSoftwareRegex**](NiaapiSoftwareRegex.md) |  | [optional] 
**Upcominglltrain** | Pointer to [**NullableNiaapiSoftwareRegex**](NiaapiSoftwareRegex.md) |  | [optional] 

## Methods

### NewNiaapiVersionRegexPlatform

`func NewNiaapiVersionRegexPlatform(classId string, objectType string, ) *NiaapiVersionRegexPlatform`

NewNiaapiVersionRegexPlatform instantiates a new NiaapiVersionRegexPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiVersionRegexPlatformWithDefaults

`func NewNiaapiVersionRegexPlatformWithDefaults() *NiaapiVersionRegexPlatform`

NewNiaapiVersionRegexPlatformWithDefaults instantiates a new NiaapiVersionRegexPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiVersionRegexPlatform) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiVersionRegexPlatform) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiVersionRegexPlatform) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiVersionRegexPlatform) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiVersionRegexPlatform) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiVersionRegexPlatform) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnyllregex

`func (o *NiaapiVersionRegexPlatform) GetAnyllregex() string`

GetAnyllregex returns the Anyllregex field if non-nil, zero value otherwise.

### GetAnyllregexOk

`func (o *NiaapiVersionRegexPlatform) GetAnyllregexOk() (*string, bool)`

GetAnyllregexOk returns a tuple with the Anyllregex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyllregex

`func (o *NiaapiVersionRegexPlatform) SetAnyllregex(v string)`

SetAnyllregex sets Anyllregex field to given value.

### HasAnyllregex

`func (o *NiaapiVersionRegexPlatform) HasAnyllregex() bool`

HasAnyllregex returns a boolean if a field has been set.

### GetCurrentlltrain

`func (o *NiaapiVersionRegexPlatform) GetCurrentlltrain() NiaapiSoftwareRegex`

GetCurrentlltrain returns the Currentlltrain field if non-nil, zero value otherwise.

### GetCurrentlltrainOk

`func (o *NiaapiVersionRegexPlatform) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool)`

GetCurrentlltrainOk returns a tuple with the Currentlltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlltrain

`func (o *NiaapiVersionRegexPlatform) SetCurrentlltrain(v NiaapiSoftwareRegex)`

SetCurrentlltrain sets Currentlltrain field to given value.

### HasCurrentlltrain

`func (o *NiaapiVersionRegexPlatform) HasCurrentlltrain() bool`

HasCurrentlltrain returns a boolean if a field has been set.

### SetCurrentlltrainNil

`func (o *NiaapiVersionRegexPlatform) SetCurrentlltrainNil(b bool)`

 SetCurrentlltrainNil sets the value for Currentlltrain to be an explicit nil

### UnsetCurrentlltrain
`func (o *NiaapiVersionRegexPlatform) UnsetCurrentlltrain()`

UnsetCurrentlltrain ensures that no value is present for Currentlltrain, not even an explicit nil
### GetLatestsltrain

`func (o *NiaapiVersionRegexPlatform) GetLatestsltrain() NiaapiSoftwareRegex`

GetLatestsltrain returns the Latestsltrain field if non-nil, zero value otherwise.

### GetLatestsltrainOk

`func (o *NiaapiVersionRegexPlatform) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool)`

GetLatestsltrainOk returns a tuple with the Latestsltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestsltrain

`func (o *NiaapiVersionRegexPlatform) SetLatestsltrain(v NiaapiSoftwareRegex)`

SetLatestsltrain sets Latestsltrain field to given value.

### HasLatestsltrain

`func (o *NiaapiVersionRegexPlatform) HasLatestsltrain() bool`

HasLatestsltrain returns a boolean if a field has been set.

### SetLatestsltrainNil

`func (o *NiaapiVersionRegexPlatform) SetLatestsltrainNil(b bool)`

 SetLatestsltrainNil sets the value for Latestsltrain to be an explicit nil

### UnsetLatestsltrain
`func (o *NiaapiVersionRegexPlatform) UnsetLatestsltrain()`

UnsetLatestsltrain ensures that no value is present for Latestsltrain, not even an explicit nil
### GetSltrain

`func (o *NiaapiVersionRegexPlatform) GetSltrain() []NiaapiSoftwareRegex`

GetSltrain returns the Sltrain field if non-nil, zero value otherwise.

### GetSltrainOk

`func (o *NiaapiVersionRegexPlatform) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool)`

GetSltrainOk returns a tuple with the Sltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltrain

`func (o *NiaapiVersionRegexPlatform) SetSltrain(v []NiaapiSoftwareRegex)`

SetSltrain sets Sltrain field to given value.

### HasSltrain

`func (o *NiaapiVersionRegexPlatform) HasSltrain() bool`

HasSltrain returns a boolean if a field has been set.

### SetSltrainNil

`func (o *NiaapiVersionRegexPlatform) SetSltrainNil(b bool)`

 SetSltrainNil sets the value for Sltrain to be an explicit nil

### UnsetSltrain
`func (o *NiaapiVersionRegexPlatform) UnsetSltrain()`

UnsetSltrain ensures that no value is present for Sltrain, not even an explicit nil
### GetUpcominglltrain

`func (o *NiaapiVersionRegexPlatform) GetUpcominglltrain() NiaapiSoftwareRegex`

GetUpcominglltrain returns the Upcominglltrain field if non-nil, zero value otherwise.

### GetUpcominglltrainOk

`func (o *NiaapiVersionRegexPlatform) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool)`

GetUpcominglltrainOk returns a tuple with the Upcominglltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcominglltrain

`func (o *NiaapiVersionRegexPlatform) SetUpcominglltrain(v NiaapiSoftwareRegex)`

SetUpcominglltrain sets Upcominglltrain field to given value.

### HasUpcominglltrain

`func (o *NiaapiVersionRegexPlatform) HasUpcominglltrain() bool`

HasUpcominglltrain returns a boolean if a field has been set.

### SetUpcominglltrainNil

`func (o *NiaapiVersionRegexPlatform) SetUpcominglltrainNil(b bool)`

 SetUpcominglltrainNil sets the value for Upcominglltrain to be an explicit nil

### UnsetUpcominglltrain
`func (o *NiaapiVersionRegexPlatform) UnsetUpcominglltrain()`

UnsetUpcominglltrain ensures that no value is present for Upcominglltrain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


