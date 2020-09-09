# NiaapiSoftwareRegexAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | Pointer to **string** | Regular Expression pattern used to reconginze the version string. | [optional] 
**SoftwareVersion** | Pointer to **string** | Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern. | [optional] 

## Methods

### NewNiaapiSoftwareRegexAllOf

`func NewNiaapiSoftwareRegexAllOf() *NiaapiSoftwareRegexAllOf`

NewNiaapiSoftwareRegexAllOf instantiates a new NiaapiSoftwareRegexAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiSoftwareRegexAllOfWithDefaults

`func NewNiaapiSoftwareRegexAllOfWithDefaults() *NiaapiSoftwareRegexAllOf`

NewNiaapiSoftwareRegexAllOfWithDefaults instantiates a new NiaapiSoftwareRegexAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *NiaapiSoftwareRegexAllOf) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *NiaapiSoftwareRegexAllOf) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *NiaapiSoftwareRegexAllOf) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *NiaapiSoftwareRegexAllOf) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *NiaapiSoftwareRegexAllOf) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NiaapiSoftwareRegexAllOf) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NiaapiSoftwareRegexAllOf) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NiaapiSoftwareRegexAllOf) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


