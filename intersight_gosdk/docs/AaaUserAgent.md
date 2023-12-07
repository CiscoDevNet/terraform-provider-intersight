# AaaUserAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "aaa.UserAgent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "aaa.UserAgent"]
**OsFamily** | Pointer to **string** | The type of operating system that sent the request. Not applicable to Intersight SDK requests. | [optional] [readonly] 
**OsVersion** | Pointer to **string** | The version of the operating system that sent the request. Not applicable for Intersight SDK requests. | [optional] [readonly] 
**SoftwareFamily** | Pointer to **string** | The type of client that made the request. For browser requests, it is the specific browser that made the request (e.g. Chrome, Firefox, etc.). | [optional] [readonly] 
**SoftwareSubtype** | Pointer to **string** | The subtype of software that made the request. For SDK requests, this is the programming language used in the SDK. For browser requests, this provides additional context on the client (e.g. if the client is running on a mobile device). | [optional] [readonly] 
**SoftwareType** | Pointer to **string** | The type of application that made the request. This can be a browser or some other software, such as an SDK. | [optional] [readonly] 
**SoftwareVersion** | Pointer to **string** | The version of the client that made the request. | [optional] [readonly] 

## Methods

### NewAaaUserAgent

`func NewAaaUserAgent(classId string, objectType string, ) *AaaUserAgent`

NewAaaUserAgent instantiates a new AaaUserAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaaUserAgentWithDefaults

`func NewAaaUserAgentWithDefaults() *AaaUserAgent`

NewAaaUserAgentWithDefaults instantiates a new AaaUserAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AaaUserAgent) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AaaUserAgent) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AaaUserAgent) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AaaUserAgent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AaaUserAgent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AaaUserAgent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOsFamily

`func (o *AaaUserAgent) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *AaaUserAgent) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *AaaUserAgent) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *AaaUserAgent) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersion

`func (o *AaaUserAgent) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *AaaUserAgent) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *AaaUserAgent) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *AaaUserAgent) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetSoftwareFamily

`func (o *AaaUserAgent) GetSoftwareFamily() string`

GetSoftwareFamily returns the SoftwareFamily field if non-nil, zero value otherwise.

### GetSoftwareFamilyOk

`func (o *AaaUserAgent) GetSoftwareFamilyOk() (*string, bool)`

GetSoftwareFamilyOk returns a tuple with the SoftwareFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareFamily

`func (o *AaaUserAgent) SetSoftwareFamily(v string)`

SetSoftwareFamily sets SoftwareFamily field to given value.

### HasSoftwareFamily

`func (o *AaaUserAgent) HasSoftwareFamily() bool`

HasSoftwareFamily returns a boolean if a field has been set.

### GetSoftwareSubtype

`func (o *AaaUserAgent) GetSoftwareSubtype() string`

GetSoftwareSubtype returns the SoftwareSubtype field if non-nil, zero value otherwise.

### GetSoftwareSubtypeOk

`func (o *AaaUserAgent) GetSoftwareSubtypeOk() (*string, bool)`

GetSoftwareSubtypeOk returns a tuple with the SoftwareSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareSubtype

`func (o *AaaUserAgent) SetSoftwareSubtype(v string)`

SetSoftwareSubtype sets SoftwareSubtype field to given value.

### HasSoftwareSubtype

`func (o *AaaUserAgent) HasSoftwareSubtype() bool`

HasSoftwareSubtype returns a boolean if a field has been set.

### GetSoftwareType

`func (o *AaaUserAgent) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *AaaUserAgent) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *AaaUserAgent) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.

### HasSoftwareType

`func (o *AaaUserAgent) HasSoftwareType() bool`

HasSoftwareType returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *AaaUserAgent) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *AaaUserAgent) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *AaaUserAgent) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *AaaUserAgent) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


