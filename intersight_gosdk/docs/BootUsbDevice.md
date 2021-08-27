# BootUsbDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.UsbDevice"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.UsbDevice"]
**ComputePhysical** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewBootUsbDevice

`func NewBootUsbDevice(classId string, objectType string, ) *BootUsbDevice`

NewBootUsbDevice instantiates a new BootUsbDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootUsbDeviceWithDefaults

`func NewBootUsbDeviceWithDefaults() *BootUsbDevice`

NewBootUsbDeviceWithDefaults instantiates a new BootUsbDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootUsbDevice) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootUsbDevice) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootUsbDevice) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootUsbDevice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootUsbDevice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootUsbDevice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetComputePhysical

`func (o *BootUsbDevice) GetComputePhysical() ComputePhysicalRelationship`

GetComputePhysical returns the ComputePhysical field if non-nil, zero value otherwise.

### GetComputePhysicalOk

`func (o *BootUsbDevice) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool)`

GetComputePhysicalOk returns a tuple with the ComputePhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePhysical

`func (o *BootUsbDevice) SetComputePhysical(v ComputePhysicalRelationship)`

SetComputePhysical sets ComputePhysical field to given value.

### HasComputePhysical

`func (o *BootUsbDevice) HasComputePhysical() bool`

HasComputePhysical returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *BootUsbDevice) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *BootUsbDevice) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *BootUsbDevice) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *BootUsbDevice) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *BootUsbDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *BootUsbDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *BootUsbDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *BootUsbDevice) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


