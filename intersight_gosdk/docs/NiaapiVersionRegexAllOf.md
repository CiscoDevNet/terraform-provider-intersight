# NiaapiVersionRegexAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.VersionRegex"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.VersionRegex"]
**Apic** | Pointer to [**NullableNiaapiVersionRegexPlatform**](NiaapiVersionRegexPlatform.md) |  | [optional] 
**Dcnm** | Pointer to [**NullableNiaapiVersionRegexPlatform**](NiaapiVersionRegexPlatform.md) |  | [optional] 
**Version** | Pointer to **string** | Version number for the Version Regex data, also used as identity. | [optional] 

## Methods

### NewNiaapiVersionRegexAllOf

`func NewNiaapiVersionRegexAllOf(classId string, objectType string, ) *NiaapiVersionRegexAllOf`

NewNiaapiVersionRegexAllOf instantiates a new NiaapiVersionRegexAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiVersionRegexAllOfWithDefaults

`func NewNiaapiVersionRegexAllOfWithDefaults() *NiaapiVersionRegexAllOf`

NewNiaapiVersionRegexAllOfWithDefaults instantiates a new NiaapiVersionRegexAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiVersionRegexAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiVersionRegexAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiVersionRegexAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiVersionRegexAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiVersionRegexAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiVersionRegexAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetApic

`func (o *NiaapiVersionRegexAllOf) GetApic() NiaapiVersionRegexPlatform`

GetApic returns the Apic field if non-nil, zero value otherwise.

### GetApicOk

`func (o *NiaapiVersionRegexAllOf) GetApicOk() (*NiaapiVersionRegexPlatform, bool)`

GetApicOk returns a tuple with the Apic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApic

`func (o *NiaapiVersionRegexAllOf) SetApic(v NiaapiVersionRegexPlatform)`

SetApic sets Apic field to given value.

### HasApic

`func (o *NiaapiVersionRegexAllOf) HasApic() bool`

HasApic returns a boolean if a field has been set.

### SetApicNil

`func (o *NiaapiVersionRegexAllOf) SetApicNil(b bool)`

 SetApicNil sets the value for Apic to be an explicit nil

### UnsetApic
`func (o *NiaapiVersionRegexAllOf) UnsetApic()`

UnsetApic ensures that no value is present for Apic, not even an explicit nil
### GetDcnm

`func (o *NiaapiVersionRegexAllOf) GetDcnm() NiaapiVersionRegexPlatform`

GetDcnm returns the Dcnm field if non-nil, zero value otherwise.

### GetDcnmOk

`func (o *NiaapiVersionRegexAllOf) GetDcnmOk() (*NiaapiVersionRegexPlatform, bool)`

GetDcnmOk returns a tuple with the Dcnm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnm

`func (o *NiaapiVersionRegexAllOf) SetDcnm(v NiaapiVersionRegexPlatform)`

SetDcnm sets Dcnm field to given value.

### HasDcnm

`func (o *NiaapiVersionRegexAllOf) HasDcnm() bool`

HasDcnm returns a boolean if a field has been set.

### SetDcnmNil

`func (o *NiaapiVersionRegexAllOf) SetDcnmNil(b bool)`

 SetDcnmNil sets the value for Dcnm to be an explicit nil

### UnsetDcnm
`func (o *NiaapiVersionRegexAllOf) UnsetDcnm()`

UnsetDcnm ensures that no value is present for Dcnm, not even an explicit nil
### GetVersion

`func (o *NiaapiVersionRegexAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiaapiVersionRegexAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiaapiVersionRegexAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiaapiVersionRegexAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


