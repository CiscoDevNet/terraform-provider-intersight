# BootUsb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for USB boot device. * &#x60;usb-cd&#x60; - Use of Compact Disk (CD) as sub-type for the USB boot device. * &#x60;usb-fdd&#x60; - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device. * &#x60;usb-hdd&#x60; - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device. | [optional] [default to "None"]

## Methods

### NewBootUsb

`func NewBootUsb() *BootUsb`

NewBootUsb instantiates a new BootUsb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootUsbWithDefaults

`func NewBootUsbWithDefaults() *BootUsb`

NewBootUsbWithDefaults instantiates a new BootUsb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtype

`func (o *BootUsb) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootUsb) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootUsb) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootUsb) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


