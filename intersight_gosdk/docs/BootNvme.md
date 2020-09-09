# BootNvme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**BootBootloader**](boot.Bootloader.md) |  | [optional] 

## Methods

### NewBootNvme

`func NewBootNvme() *BootNvme`

NewBootNvme instantiates a new BootNvme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootNvmeWithDefaults

`func NewBootNvmeWithDefaults() *BootNvme`

NewBootNvmeWithDefaults instantiates a new BootNvme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


