# NiaapiVersionRegex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.VersionRegex"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.VersionRegex"]
**Apic** | Pointer to [**NullableNiaapiVersionRegexPlatform**](NiaapiVersionRegexPlatform.md) |  | [optional] 
**Dcnm** | Pointer to [**NullableNiaapiVersionRegexPlatform**](NiaapiVersionRegexPlatform.md) |  | [optional] 
**Version** | Pointer to **string** | Version number for the Version Regex data, also used as identity. | [optional] 

## Methods

### NewNiaapiVersionRegex

`func NewNiaapiVersionRegex(classId string, objectType string, ) *NiaapiVersionRegex`

NewNiaapiVersionRegex instantiates a new NiaapiVersionRegex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiVersionRegexWithDefaults

`func NewNiaapiVersionRegexWithDefaults() *NiaapiVersionRegex`

NewNiaapiVersionRegexWithDefaults instantiates a new NiaapiVersionRegex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiVersionRegex) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiVersionRegex) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiVersionRegex) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiVersionRegex) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiVersionRegex) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiVersionRegex) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApic

`func (o *NiaapiVersionRegex) GetApic() NiaapiVersionRegexPlatform`

GetApic returns the Apic field if non-nil, zero value otherwise.

### GetApicOk

`func (o *NiaapiVersionRegex) GetApicOk() (*NiaapiVersionRegexPlatform, bool)`

GetApicOk returns a tuple with the Apic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApic

`func (o *NiaapiVersionRegex) SetApic(v NiaapiVersionRegexPlatform)`

SetApic sets Apic field to given value.

### HasApic

`func (o *NiaapiVersionRegex) HasApic() bool`

HasApic returns a boolean if a field has been set.

### SetApicNil

`func (o *NiaapiVersionRegex) SetApicNil(b bool)`

 SetApicNil sets the value for Apic to be an explicit nil

### UnsetApic
`func (o *NiaapiVersionRegex) UnsetApic()`

UnsetApic ensures that no value is present for Apic, not even an explicit nil
### GetDcnm

`func (o *NiaapiVersionRegex) GetDcnm() NiaapiVersionRegexPlatform`

GetDcnm returns the Dcnm field if non-nil, zero value otherwise.

### GetDcnmOk

`func (o *NiaapiVersionRegex) GetDcnmOk() (*NiaapiVersionRegexPlatform, bool)`

GetDcnmOk returns a tuple with the Dcnm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnm

`func (o *NiaapiVersionRegex) SetDcnm(v NiaapiVersionRegexPlatform)`

SetDcnm sets Dcnm field to given value.

### HasDcnm

`func (o *NiaapiVersionRegex) HasDcnm() bool`

HasDcnm returns a boolean if a field has been set.

### SetDcnmNil

`func (o *NiaapiVersionRegex) SetDcnmNil(b bool)`

 SetDcnmNil sets the value for Dcnm to be an explicit nil

### UnsetDcnm
`func (o *NiaapiVersionRegex) UnsetDcnm()`

UnsetDcnm ensures that no value is present for Dcnm, not even an explicit nil
### GetVersion

`func (o *NiaapiVersionRegex) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiVersionRegex) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiVersionRegex) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiVersionRegex) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


