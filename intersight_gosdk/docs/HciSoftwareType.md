# HciSoftwareType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.SoftwareType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.SoftwareType"]
**SoftwareType** | Pointer to **string** | The type of the software. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the software. | [optional] [readonly] 

## Methods

### NewHciSoftwareType

`func NewHciSoftwareType(classId string, objectType string, ) *HciSoftwareType`

NewHciSoftwareType instantiates a new HciSoftwareType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciSoftwareTypeWithDefaults

`func NewHciSoftwareTypeWithDefaults() *HciSoftwareType`

NewHciSoftwareTypeWithDefaults instantiates a new HciSoftwareType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciSoftwareType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciSoftwareType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciSoftwareType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciSoftwareType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciSoftwareType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciSoftwareType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSoftwareType

`func (o *HciSoftwareType) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *HciSoftwareType) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *HciSoftwareType) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.

### HasSoftwareType

`func (o *HciSoftwareType) HasSoftwareType() bool`

HasSoftwareType returns a boolean if a field has been set.

### GetVersion

`func (o *HciSoftwareType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HciSoftwareType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HciSoftwareType) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HciSoftwareType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


