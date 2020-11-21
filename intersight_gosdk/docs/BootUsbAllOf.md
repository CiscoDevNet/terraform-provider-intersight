# BootUsbAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Usb"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Usb"]
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for USB boot device. * &#x60;usb-cd&#x60; - Use of Compact Disk (CD) as sub-type for the USB boot device. * &#x60;usb-fdd&#x60; - Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device. * &#x60;usb-hdd&#x60; - Use of Hard Disk Drive (HDD) as sub-type for the USB boot device. | [optional] [default to "None"]

## Methods

### NewBootUsbAllOf

`func NewBootUsbAllOf(classId string, objectType string, ) *BootUsbAllOf`

NewBootUsbAllOf instantiates a new BootUsbAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootUsbAllOfWithDefaults

`func NewBootUsbAllOfWithDefaults() *BootUsbAllOf`

NewBootUsbAllOfWithDefaults instantiates a new BootUsbAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootUsbAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootUsbAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootUsbAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootUsbAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootUsbAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootUsbAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSubtype

`func (o *BootUsbAllOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootUsbAllOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootUsbAllOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootUsbAllOf) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


