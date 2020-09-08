# NiaapiVersionRegex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apic** | Pointer to [**NiaapiVersionRegexPlatform**](niaapi.VersionRegexPlatform.md) |  | [optional] 
**Dcnm** | Pointer to [**NiaapiVersionRegexPlatform**](niaapi.VersionRegexPlatform.md) |  | [optional] 
**Version** | Pointer to **string** | Version number for the Version Regex data, also used as identity. | [optional] 

## Methods

### NewNiaapiVersionRegex

`func NewNiaapiVersionRegex() *NiaapiVersionRegex`

NewNiaapiVersionRegex instantiates a new NiaapiVersionRegex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiVersionRegexWithDefaults

`func NewNiaapiVersionRegexWithDefaults() *NiaapiVersionRegex`

NewNiaapiVersionRegexWithDefaults instantiates a new NiaapiVersionRegex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


