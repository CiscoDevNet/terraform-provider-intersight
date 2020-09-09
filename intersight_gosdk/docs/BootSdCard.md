# BootSdCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**BootBootloader**](boot.Bootloader.md) |  | [optional] 
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. | [optional] 
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for SD card boot device. * &#x60;flex-util&#x60; - Use of FlexUtil (microSD) card as sub type for SD card boot device. * &#x60;flex-flash&#x60; - Use of FlexFlash (SD) card as sub type for SD card boot device. * &#x60;SDCARD&#x60; - Use of SD card as sub type for the SD Card boot device. | [optional] [default to "None"]

## Methods

### NewBootSdCard

`func NewBootSdCard() *BootSdCard`

NewBootSdCard instantiates a new BootSdCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootSdCardWithDefaults

`func NewBootSdCardWithDefaults() *BootSdCard`

NewBootSdCardWithDefaults instantiates a new BootSdCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


