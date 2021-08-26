# GraphicsControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "graphics.Controller"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "graphics.Controller"]
**ControllerId** | Pointer to **int64** | The id of the graphics controller. | [optional] [readonly] 
**PciAddr** | Pointer to **string** | The PCI address of the graphics controller. | [optional] [readonly] 
**PciSlot** | Pointer to **string** | The PCI slot information of the graphics controller. | [optional] [readonly] 
**GraphicsCard** | Pointer to [**GraphicsCardRelationship**](GraphicsCardRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewGraphicsControllerAllOf

`func NewGraphicsControllerAllOf(classId string, objectType string, ) *GraphicsControllerAllOf`

NewGraphicsControllerAllOf instantiates a new GraphicsControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphicsControllerAllOfWithDefaults

`func NewGraphicsControllerAllOfWithDefaults() *GraphicsControllerAllOf`

NewGraphicsControllerAllOfWithDefaults instantiates a new GraphicsControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *GraphicsControllerAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *GraphicsControllerAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *GraphicsControllerAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *GraphicsControllerAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GraphicsControllerAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GraphicsControllerAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetControllerId

`func (o *GraphicsControllerAllOf) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *GraphicsControllerAllOf) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *GraphicsControllerAllOf) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *GraphicsControllerAllOf) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetPciAddr

`func (o *GraphicsControllerAllOf) GetPciAddr() string`

GetPciAddr returns the PciAddr field if non-nil, zero value otherwise.

### GetPciAddrOk

`func (o *GraphicsControllerAllOf) GetPciAddrOk() (*string, bool)`

GetPciAddrOk returns a tuple with the PciAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddr

`func (o *GraphicsControllerAllOf) SetPciAddr(v string)`

SetPciAddr sets PciAddr field to given value.

### HasPciAddr

`func (o *GraphicsControllerAllOf) HasPciAddr() bool`

HasPciAddr returns a boolean if a field has been set.

### GetPciSlot

`func (o *GraphicsControllerAllOf) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *GraphicsControllerAllOf) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *GraphicsControllerAllOf) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *GraphicsControllerAllOf) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetGraphicsCard

`func (o *GraphicsControllerAllOf) GetGraphicsCard() GraphicsCardRelationship`

GetGraphicsCard returns the GraphicsCard field if non-nil, zero value otherwise.

### GetGraphicsCardOk

`func (o *GraphicsControllerAllOf) GetGraphicsCardOk() (*GraphicsCardRelationship, bool)`

GetGraphicsCardOk returns a tuple with the GraphicsCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicsCard

`func (o *GraphicsControllerAllOf) SetGraphicsCard(v GraphicsCardRelationship)`

SetGraphicsCard sets GraphicsCard field to given value.

### HasGraphicsCard

`func (o *GraphicsControllerAllOf) HasGraphicsCard() bool`

HasGraphicsCard returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *GraphicsControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *GraphicsControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *GraphicsControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *GraphicsControllerAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *GraphicsControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *GraphicsControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *GraphicsControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *GraphicsControllerAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


