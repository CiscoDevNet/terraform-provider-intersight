# BootSan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**BootBootloader**](boot.Bootloader.md) |  | [optional] 
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. | [optional] 
**Slot** | Pointer to **string** | Slot ID of the device. Supported values are ( 1 - 255, \&quot;MLOM\&quot;, \&quot;L1\&quot;, \&quot;L2\&quot; ). | [optional] 

## Methods

### NewBootSan

`func NewBootSan() *BootSan`

NewBootSan instantiates a new BootSan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootSanWithDefaults

`func NewBootSanWithDefaults() *BootSan`

NewBootSanWithDefaults instantiates a new BootSan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootloader

`func (o *BootSan) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootSan) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootSan) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootSan) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetLun

`func (o *BootSan) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *BootSan) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *BootSan) SetLun(v int64)`

SetLun sets Lun field to given value.

### HasLun

`func (o *BootSan) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetSlot

`func (o *BootSan) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootSan) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootSan) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootSan) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


