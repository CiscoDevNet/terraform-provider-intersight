# BootNvmeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Nvme"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Nvme"]
**Bootloader** | Pointer to [**NullableBootBootloader**](BootBootloader.md) |  | [optional] 

## Methods

### NewBootNvmeAllOf

`func NewBootNvmeAllOf(classId string, objectType string, ) *BootNvmeAllOf`

NewBootNvmeAllOf instantiates a new BootNvmeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootNvmeAllOfWithDefaults

`func NewBootNvmeAllOfWithDefaults() *BootNvmeAllOf`

NewBootNvmeAllOfWithDefaults instantiates a new BootNvmeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootNvmeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootNvmeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootNvmeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootNvmeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootNvmeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootNvmeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootloader

`func (o *BootNvmeAllOf) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootNvmeAllOf) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootNvmeAllOf) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootNvmeAllOf) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### SetBootloaderNil

`func (o *BootNvmeAllOf) SetBootloaderNil(b bool)`

 SetBootloaderNil sets the value for Bootloader to be an explicit nil

### UnsetBootloader
`func (o *BootNvmeAllOf) UnsetBootloader()`

UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


