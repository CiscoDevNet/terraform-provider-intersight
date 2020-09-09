# BootVirtualMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for virtual media. * &#x60;cimc-mapped-dvd&#x60; - The virtual media device is mapped to a virtual DVD device. * &#x60;cimc-mapped-hdd&#x60; - The virtual media device is mapped to a virtual HDD device. * &#x60;kvm-mapped-dvd&#x60; - A KVM mapped DVD virtual media device. * &#x60;kvm-mapped-hdd&#x60; - A KVM mapped HDD virtual media device. * &#x60;kvm-mapped-fdd&#x60; - A KVM mapped FDD virtual media device. | [optional] [default to "None"]

## Methods

### NewBootVirtualMedia

`func NewBootVirtualMedia() *BootVirtualMedia`

NewBootVirtualMedia instantiates a new BootVirtualMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootVirtualMediaWithDefaults

`func NewBootVirtualMediaWithDefaults() *BootVirtualMedia`

NewBootVirtualMediaWithDefaults instantiates a new BootVirtualMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtype

`func (o *BootVirtualMedia) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootVirtualMedia) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootVirtualMedia) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootVirtualMedia) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


