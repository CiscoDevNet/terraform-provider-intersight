# BootIscsiAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Iscsi"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Iscsi"]
**Bootloader** | Pointer to [**NullableBootBootloader**](BootBootloader.md) |  | [optional] 
**InterfaceName** | Pointer to **string** | The name of the underlying virtual ethernet interface used by the iSCSI boot device. | [optional] 
**Port** | Pointer to **int64** | Port ID of the ISCSI boot device. | [optional] [default to 0]
**Slot** | Pointer to **string** | The slot id of the device. Supported values are (1 - 255, \&quot;MLOM\&quot;, \&quot;L\&quot;, \&quot;L1\&quot;, \&quot;L2\&quot;, \&quot;OCP\&quot;). | [optional] 

## Methods

### NewBootIscsiAllOf

`func NewBootIscsiAllOf(classId string, objectType string, ) *BootIscsiAllOf`

NewBootIscsiAllOf instantiates a new BootIscsiAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootIscsiAllOfWithDefaults

`func NewBootIscsiAllOfWithDefaults() *BootIscsiAllOf`

NewBootIscsiAllOfWithDefaults instantiates a new BootIscsiAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootIscsiAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootIscsiAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootIscsiAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootIscsiAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootIscsiAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootIscsiAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBootloader

`func (o *BootIscsiAllOf) GetBootloader() BootBootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *BootIscsiAllOf) GetBootloaderOk() (*BootBootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *BootIscsiAllOf) SetBootloader(v BootBootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *BootIscsiAllOf) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### SetBootloaderNil

`func (o *BootIscsiAllOf) SetBootloaderNil(b bool)`

 SetBootloaderNil sets the value for Bootloader to be an explicit nil

### UnsetBootloader
`func (o *BootIscsiAllOf) UnsetBootloader()`

UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
### GetInterfaceName

`func (o *BootIscsiAllOf) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *BootIscsiAllOf) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *BootIscsiAllOf) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *BootIscsiAllOf) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetPort

`func (o *BootIscsiAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BootIscsiAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BootIscsiAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *BootIscsiAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSlot

`func (o *BootIscsiAllOf) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootIscsiAllOf) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootIscsiAllOf) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootIscsiAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


