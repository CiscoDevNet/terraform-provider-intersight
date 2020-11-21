# BootVirtualMediaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.VirtualMedia"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.VirtualMedia"]
**Subtype** | Pointer to **string** | The subtype for the selected device type. * &#x60;None&#x60; - No sub type for virtual media. * &#x60;cimc-mapped-dvd&#x60; - The virtual media device is mapped to a virtual DVD device. * &#x60;cimc-mapped-hdd&#x60; - The virtual media device is mapped to a virtual HDD device. * &#x60;kvm-mapped-dvd&#x60; - A KVM mapped DVD virtual media device. * &#x60;kvm-mapped-hdd&#x60; - A KVM mapped HDD virtual media device. * &#x60;kvm-mapped-fdd&#x60; - A KVM mapped FDD virtual media device. | [optional] [default to "None"]

## Methods

### NewBootVirtualMediaAllOf

`func NewBootVirtualMediaAllOf(classId string, objectType string, ) *BootVirtualMediaAllOf`

NewBootVirtualMediaAllOf instantiates a new BootVirtualMediaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootVirtualMediaAllOfWithDefaults

`func NewBootVirtualMediaAllOfWithDefaults() *BootVirtualMediaAllOf`

NewBootVirtualMediaAllOfWithDefaults instantiates a new BootVirtualMediaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootVirtualMediaAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootVirtualMediaAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootVirtualMediaAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootVirtualMediaAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootVirtualMediaAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootVirtualMediaAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSubtype

`func (o *BootVirtualMediaAllOf) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BootVirtualMediaAllOf) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BootVirtualMediaAllOf) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BootVirtualMediaAllOf) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


