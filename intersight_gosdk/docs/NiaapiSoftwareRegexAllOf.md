# NiaapiSoftwareRegexAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niaapi.SoftwareRegex"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niaapi.SoftwareRegex"]
**Regex** | Pointer to **string** | Regular Expression pattern used to reconginze the version string. | [optional] 
**SoftwareVersion** | Pointer to **string** | Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern. | [optional] 

## Methods

### NewNiaapiSoftwareRegexAllOf

`func NewNiaapiSoftwareRegexAllOf(classId string, objectType string, ) *NiaapiSoftwareRegexAllOf`

NewNiaapiSoftwareRegexAllOf instantiates a new NiaapiSoftwareRegexAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiSoftwareRegexAllOfWithDefaults

`func NewNiaapiSoftwareRegexAllOfWithDefaults() *NiaapiSoftwareRegexAllOf`

NewNiaapiSoftwareRegexAllOfWithDefaults instantiates a new NiaapiSoftwareRegexAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiSoftwareRegexAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiSoftwareRegexAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiSoftwareRegexAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiSoftwareRegexAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiSoftwareRegexAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiSoftwareRegexAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


