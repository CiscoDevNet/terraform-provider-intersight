# NiaapiVersionRegexPlatformAllOf

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

### NewNiaapiVersionRegexPlatformAllOf

`func NewNiaapiVersionRegexPlatformAllOf(classId string, objectType string, ) *NiaapiVersionRegexPlatformAllOf`

NewNiaapiVersionRegexPlatformAllOf instantiates a new NiaapiVersionRegexPlatformAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiVersionRegexPlatformAllOfWithDefaults

`func NewNiaapiVersionRegexPlatformAllOfWithDefaults() *NiaapiVersionRegexPlatformAllOf`

NewNiaapiVersionRegexPlatformAllOfWithDefaults instantiates a new NiaapiVersionRegexPlatformAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiVersionRegexPlatformAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiVersionRegexPlatformAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiVersionRegexPlatformAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiVersionRegexPlatformAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnyllregex

`func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregex() string`

GetAnyllregex returns the Anyllregex field if non-nil, zero value otherwise.

### GetAnyllregexOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetAnyllregexOk() (*string, bool)`

GetAnyllregexOk returns a tuple with the Anyllregex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyllregex

`func (o *NiaapiVersionRegexPlatformAllOf) SetAnyllregex(v string)`

SetAnyllregex sets Anyllregex field to given value.

### HasAnyllregex

`func (o *NiaapiVersionRegexPlatformAllOf) HasAnyllregex() bool`

HasAnyllregex returns a boolean if a field has been set.

### GetCurrentlltrain

`func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrain() NiaapiSoftwareRegex`

GetCurrentlltrain returns the Currentlltrain field if non-nil, zero value otherwise.

### GetCurrentlltrainOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetCurrentlltrainOk() (*NiaapiSoftwareRegex, bool)`

GetCurrentlltrainOk returns a tuple with the Currentlltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlltrain

`func (o *NiaapiVersionRegexPlatformAllOf) SetCurrentlltrain(v NiaapiSoftwareRegex)`

SetCurrentlltrain sets Currentlltrain field to given value.

### HasCurrentlltrain

`func (o *NiaapiVersionRegexPlatformAllOf) HasCurrentlltrain() bool`

HasCurrentlltrain returns a boolean if a field has been set.

### SetCurrentlltrainNil

`func (o *NiaapiVersionRegexPlatformAllOf) SetCurrentlltrainNil(b bool)`

 SetCurrentlltrainNil sets the value for Currentlltrain to be an explicit nil

### UnsetCurrentlltrain
`func (o *NiaapiVersionRegexPlatformAllOf) UnsetCurrentlltrain()`

UnsetCurrentlltrain ensures that no value is present for Currentlltrain, not even an explicit nil
### GetLatestsltrain

`func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrain() NiaapiSoftwareRegex`

GetLatestsltrain returns the Latestsltrain field if non-nil, zero value otherwise.

### GetLatestsltrainOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetLatestsltrainOk() (*NiaapiSoftwareRegex, bool)`

GetLatestsltrainOk returns a tuple with the Latestsltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestsltrain

`func (o *NiaapiVersionRegexPlatformAllOf) SetLatestsltrain(v NiaapiSoftwareRegex)`

SetLatestsltrain sets Latestsltrain field to given value.

### HasLatestsltrain

`func (o *NiaapiVersionRegexPlatformAllOf) HasLatestsltrain() bool`

HasLatestsltrain returns a boolean if a field has been set.

### SetLatestsltrainNil

`func (o *NiaapiVersionRegexPlatformAllOf) SetLatestsltrainNil(b bool)`

 SetLatestsltrainNil sets the value for Latestsltrain to be an explicit nil

### UnsetLatestsltrain
`func (o *NiaapiVersionRegexPlatformAllOf) UnsetLatestsltrain()`

UnsetLatestsltrain ensures that no value is present for Latestsltrain, not even an explicit nil
### GetSltrain

`func (o *NiaapiVersionRegexPlatformAllOf) GetSltrain() []NiaapiSoftwareRegex`

GetSltrain returns the Sltrain field if non-nil, zero value otherwise.

### GetSltrainOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetSltrainOk() (*[]NiaapiSoftwareRegex, bool)`

GetSltrainOk returns a tuple with the Sltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltrain

`func (o *NiaapiVersionRegexPlatformAllOf) SetSltrain(v []NiaapiSoftwareRegex)`

SetSltrain sets Sltrain field to given value.

### HasSltrain

`func (o *NiaapiVersionRegexPlatformAllOf) HasSltrain() bool`

HasSltrain returns a boolean if a field has been set.

### SetSltrainNil

`func (o *NiaapiVersionRegexPlatformAllOf) SetSltrainNil(b bool)`

 SetSltrainNil sets the value for Sltrain to be an explicit nil

### UnsetSltrain
`func (o *NiaapiVersionRegexPlatformAllOf) UnsetSltrain()`

UnsetSltrain ensures that no value is present for Sltrain, not even an explicit nil
### GetUpcominglltrain

`func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrain() NiaapiSoftwareRegex`

GetUpcominglltrain returns the Upcominglltrain field if non-nil, zero value otherwise.

### GetUpcominglltrainOk

`func (o *NiaapiVersionRegexPlatformAllOf) GetUpcominglltrainOk() (*NiaapiSoftwareRegex, bool)`

GetUpcominglltrainOk returns a tuple with the Upcominglltrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcominglltrain

`func (o *NiaapiVersionRegexPlatformAllOf) SetUpcominglltrain(v NiaapiSoftwareRegex)`

SetUpcominglltrain sets Upcominglltrain field to given value.

### HasUpcominglltrain

`func (o *NiaapiVersionRegexPlatformAllOf) HasUpcominglltrain() bool`

HasUpcominglltrain returns a boolean if a field has been set.

### SetUpcominglltrainNil

`func (o *NiaapiVersionRegexPlatformAllOf) SetUpcominglltrainNil(b bool)`

 SetUpcominglltrainNil sets the value for Upcominglltrain to be an explicit nil

### UnsetUpcominglltrain
`func (o *NiaapiVersionRegexPlatformAllOf) UnsetUpcominglltrain()`

UnsetUpcominglltrain ensures that no value is present for Upcominglltrain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


