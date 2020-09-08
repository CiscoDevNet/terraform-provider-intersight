# BootLocalDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**BootBootloader**](boot.Bootloader.md) |  | [optional] 
**Slot** | Pointer to **string** | The slot id of the local disk device. Supported values are (1-255, \&quot;M\&quot;, \&quot;HBA\&quot;, \&quot;SAS\&quot;, \&quot;RAID\&quot;, \&quot;MRAID\&quot;, \&quot;MSTOR-RAID\&quot;). | [optional] 

## Methods

### NewBootLocalDisk

`func NewBootLocalDisk() *BootLocalDisk`

NewBootLocalDisk instantiates a new BootLocalDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootLocalDiskWithDefaults

`func NewBootLocalDiskWithDefaults() *BootLocalDisk`

NewBootLocalDiskWithDefaults instantiates a new BootLocalDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootloader

`func (o *BootLocalDisk) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootLocalDisk) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootLocalDisk) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootLocalDisk) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetSlot

`func (o *BootLocalDisk) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootLocalDisk) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootLocalDisk) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootLocalDisk) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


