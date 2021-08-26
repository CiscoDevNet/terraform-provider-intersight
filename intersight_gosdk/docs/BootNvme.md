# BootNvme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Nvme"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Nvme"]
**Bootloader** | Pointer to [**NullableBootBootloader**](BootBootloader.md) |  | [optional] 

## Methods

### NewBootNvme

`func NewBootNvme(classId string, objectType string, ) *BootNvme`

NewBootNvme instantiates a new BootNvme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootNvmeWithDefaults

`func NewBootNvmeWithDefaults() *BootNvme`

NewBootNvmeWithDefaults instantiates a new BootNvme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootNvme) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootNvme) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootNvme) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootNvme) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootNvme) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootNvme) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootloader

`func (o *BootNvme) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootNvme) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootNvme) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootNvme) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### SetBootloaderNil

`func (o *BootNvme) SetBootloaderNil(b bool)`

 SetBootloaderNil sets the value for Bootloader to be an explicit nil

### UnsetBootloader
`func (o *BootNvme) UnsetBootloader()`

UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


