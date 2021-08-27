# BootSdCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.SdCard"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.SdCard"]
**Bootloader** | Pointer to [**NullableBootBootloader**](BootBootloader.md) |  | [optional] 
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. | [optional] [default to 0]
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for SD card boot device. * &#x60;flex-util&#x60; - Use of FlexUtil (microSD) card as sub type for SD card boot device. * &#x60;flex-flash&#x60; - Use of FlexFlash (SD) card as sub type for SD card boot device. * &#x60;SDCARD&#x60; - Use of SD card as sub type for the SD Card boot device. | [optional] [default to "None"]

## Methods

### NewBootSdCard

`func NewBootSdCard(classId string, objectType string, ) *BootSdCard`

NewBootSdCard instantiates a new BootSdCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootSdCardWithDefaults

`func NewBootSdCardWithDefaults() *BootSdCard`

NewBootSdCardWithDefaults instantiates a new BootSdCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootSdCard) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootSdCard) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootSdCard) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootSdCard) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootSdCard) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootSdCard) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootloader

`func (o *BootSdCard) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootSdCard) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootSdCard) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootSdCard) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### SetBootloaderNil

`func (o *BootSdCard) SetBootloaderNil(b bool)`

 SetBootloaderNil sets the value for Bootloader to be an explicit nil

### UnsetBootloader
`func (o *BootSdCard) UnsetBootloader()`

UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
### GetLun

`func (o *BootSdCard) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *BootSdCard) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *BootSdCard) SetLun(v int64)`

SetLun sets Lun field to given value.

### HasLun

`func (o *BootSdCard) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetSubtype

`func (o *BootSdCard) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootSdCard) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootSdCard) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootSdCard) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


