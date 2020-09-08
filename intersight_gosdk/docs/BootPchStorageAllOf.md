# BootPchStorageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**BootBootloader**](boot.Bootloader.md) |  | [optional] 
**Lun** | Pointer to **int64** | The Logical Unit Number (LUN) of the device. It is the Virtual Drive number for Cisco UCS C-Series Servers. Virtual Drive number is found in storage inventory. | [optional] 

## Methods

### NewBootPchStorageAllOf

`func NewBootPchStorageAllOf() *BootPchStorageAllOf`

NewBootPchStorageAllOf instantiates a new BootPchStorageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPchStorageAllOfWithDefaults

`func NewBootPchStorageAllOfWithDefaults() *BootPchStorageAllOf`

NewBootPchStorageAllOfWithDefaults instantiates a new BootPchStorageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootloader

`func (o *BootPchStorageAllOf) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootPchStorageAllOf) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootPchStorageAllOf) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootPchStorageAllOf) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetLun

`func (o *BootPchStorageAllOf) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *BootPchStorageAllOf) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *BootPchStorageAllOf) SetLun(v int64)`

SetLun sets Lun field to given value.

### HasLun

`func (o *BootPchStorageAllOf) HasLun() bool`

HasLun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


